/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

var (
	typeNames = flag.String("type", "", "comma-separated list of type names; must be set")
	output    = flag.String("output", "", "output file name; default srcdir/<type>_plc4xgen.go")
	buildTags = flag.String("tags", "", "comma-separated list of build tags to apply")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage of plc4xgenerator:\n")
	_, _ = fmt.Fprintf(os.Stderr, "\tplc4xgenerator [flags] -type T [directory]\n")
	_, _ = fmt.Fprintf(os.Stderr, "\tplc4xgenerator [flags] -type T files... # Must be a single package\n")
	_, _ = fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("plc4xgenerator: ")
	flag.Usage = Usage
	flag.Parse()
	if len(*typeNames) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	typeList := strings.Split(*typeNames, ",")
	var tags []string
	if len(*buildTags) > 0 {
		tags = strings.Split(*buildTags, ",")
	}

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"."}
	}

	// Parse the package once.
	var dir string
	generator := Generator{}
	if len(args) == 1 && isDirectory(args[0]) {
		dir = args[0]
	} else {
		if len(tags) != 0 {
			log.Fatal("-tags option applies only to directories, not when files are specified")
		}
		dir = filepath.Dir(args[0])
	}

	generator.parsePackage(args, tags)

	// Print the header and package clause.
	generator.Printf(asfHeader)
	generator.Printf("// Code generated by \"plc4xgenerator %s\"; DO NOT EDIT.\n", strings.Join(os.Args[1:], " "))
	generator.Printf("\n")
	generator.Printf("package %s", generator.pkg.name)
	generator.Printf("\n")
	generator.Printf("import (\n")
	generator.Printf("\t\"context\"\n")
	generator.Printf("\t\"encoding/binary\"\n")
	generator.Printf("")
	generator.Printf("\t\"github.com/apache/plc4x/plc4go/spi/utils\"\n")
	generator.Printf("\t\"fmt\"\n")
	generator.Printf(")\n")
	generator.Printf("\n")
	generator.Printf("var _ = fmt.Printf\n")

	// Run generate for each type.
	for _, typeName := range typeList {
		generator.generate(typeName)
	}

	// Format the output.
	src := generator.format()

	// Write to file.
	outputName := *output
	if outputName == "" {
		baseName := fmt.Sprintf("%s_plc4xgen.go", typeList[0])
		outputName = filepath.Join(dir, baseName)
	}
	err := os.WriteFile(outputName, src, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

// isDirectory reports whether the named file is a directory.
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	buf bytes.Buffer // Accumulated output.
	pkg *Package     // Package we are scanning.
}

func (g *Generator) Printf(format string, args ...any) {
	_, _ = fmt.Fprintf(&g.buf, format, args...)
}

// File holds a single parsed file and associated data.
type File struct {
	pkg  *Package  // Package to which this file belongs.
	file *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	typeName string  // Name of the type.
	fields   []Field // Accumulator for fields of that type.

	trimPrefix  string
	lineComment bool
}

type Package struct {
	name  string
	defs  map[*ast.Ident]types.Object
	files []*File
}

// parsePackage analyzes the single package constructed from the patterns and tags.
// parsePackage exits if there is an error.
func (g *Generator) parsePackage(patterns []string, tags []string) {
	cfg := &packages.Config{
		Mode:       packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
		Tests:      false,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
	}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}
	g.addPackage(pkgs[0])
}

// addPackage adds a type checked Package and its syntax files to the generator.
func (g *Generator) addPackage(pkg *packages.Package) {
	g.pkg = &Package{
		name:  pkg.Name,
		defs:  pkg.TypesInfo.Defs,
		files: make([]*File, len(pkg.Syntax)),
	}

	for i, file := range pkg.Syntax {
		g.pkg.files[i] = &File{
			file: file,
			pkg:  g.pkg,
		}
	}
}

// generate produces the String method for the named type.
func (g *Generator) generate(typeName string) {
	fields := make([]Field, 0, 100)
	for _, file := range g.pkg.files {
		// Set the state for this run of the walker.
		file.typeName = typeName
		if file.file != nil {
			ast.Inspect(file.file, file.genDecl)
			fields = append(fields, file.fields...)
		}
	}
	if len(fields) == 0 {
		log.Fatalf("no fields defined for type %s", typeName)
	}
	// TODO: for now we remove Default from the start (maybe move that to an option)
	logicalTypeName := "\"" + strings.TrimPrefix(typeName, "Default") + "\""

	// Generate code that will fail if the constants change value.
	g.Printf("func (d *%s) Serialize() ([]byte, error) {\n", typeName)
	g.Printf("wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))\n")
	g.Printf("\tif err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {\n")
	g.Printf("\t\treturn nil, err\n")
	g.Printf("\t}\n")
	g.Printf("\treturn wb.GetBytes(), nil\n")
	g.Printf("}\n\n")
	g.Printf("func (d *%s) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {\n", typeName)
	g.Printf("\tif err := writeBuffer.PushContext(%s); err != nil {\n", logicalTypeName)
	g.Printf("\t\treturn err\n")
	g.Printf("\t}\n")
	for _, field := range fields {
		fieldType := field.fieldType
		if field.isDelegate {
			g.Printf("\t\t\tif err := d.%s.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {\n", fieldType.(*ast.Ident).Name)
			g.Printf("\t\t\t\treturn err\n")
			g.Printf("\t\t\t}\n")
			continue
		}
		fieldName := field.name
		fieldNameUntitled := "\"" + unTitle(fieldName) + "\""
		if field.hasLocker != "" {
			g.Printf("if err := func()error {\n")
			g.Printf("\td." + field.hasLocker + ".Lock()\n")
			g.Printf("\tdefer d." + field.hasLocker + ".Unlock()\n")
		}
		needsDereference := false
		if starFieldType, ok := fieldType.(*ast.StarExpr); ok {
			fieldType = starFieldType.X
			needsDereference = true
		}
		if field.isStringer {
			if needsDereference {
				g.Printf("if d.%s != nil {", field.name)
			}
			g.Printf(stringFieldSerialize, "d."+field.name+".String()", fieldNameUntitled)
			if field.hasLocker != "" {
				g.Printf("\treturn nil\n")
				g.Printf("}(); err != nil {\n")
				g.Printf("\treturn err\n")
				g.Printf("}\n")
			}
			if needsDereference {
				g.Printf("}\n")
			}
			continue
		}
		switch fieldType := fieldType.(type) {
		case *ast.SelectorExpr:
			{
				// TODO: bit hacky but not sure how else we catch those ones
				x := fieldType.X
				sel := fieldType.Sel
				xIdent, xIsIdent := x.(*ast.Ident)
				if xIsIdent {
					if xIdent.Name == "atomic" {
						if sel.Name == "Uint32" {
							g.Printf(uint32FieldSerialize, "d."+field.name+".Load()", fieldNameUntitled)
							if field.hasLocker != "" {
								g.Printf("\treturn nil\n")
								g.Printf("}(); err != nil {\n")
								g.Printf("\treturn err\n")
								g.Printf("}\n")
							}
							continue
						}
						if sel.Name == "Uint64" {
							g.Printf(uint64FieldSerialize, "d."+field.name+".Load()", fieldNameUntitled)
							if field.hasLocker != "" {
								g.Printf("\treturn nil\n")
								g.Printf("}(); err != nil {\n")
								g.Printf("\treturn err\n")
								g.Printf("}\n")
							}
							continue
						}
						if sel.Name == "Int32" {
							g.Printf(int32FieldSerialize, "d."+field.name+".Load()", fieldNameUntitled)
							if field.hasLocker != "" {
								g.Printf("\treturn nil\n")
								g.Printf("}(); err != nil {\n")
								g.Printf("\treturn err\n")
								g.Printf("}\n")
							}
							continue
						}
						if sel.Name == "Bool" {
							g.Printf(boolFieldSerialize, "d."+field.name+".Load()", fieldNameUntitled)
							if field.hasLocker != "" {
								g.Printf("\treturn nil\n")
								g.Printf("}(); err != nil {\n")
								g.Printf("\treturn err\n")
								g.Printf("}\n")
							}
							continue
						}
						if sel.Name == "Value" {
							g.Printf(serializableFieldTemplate, "d."+field.name+".Load()", fieldNameUntitled)
							if field.hasLocker != "" {
								g.Printf("\treturn nil\n")
								g.Printf("}(); err != nil {\n")
								g.Printf("\treturn err\n")
								g.Printf("}\n")
							}
							continue
						}
					}
					if xIdent.Name == "sync" {
						fmt.Printf("\t skipping field %s because it is %v.%v\n", fieldName, x, sel)
						if field.hasLocker != "" {
							g.Printf("\treturn nil\n")
							g.Printf("}(); err != nil {\n")
							g.Printf("\treturn err\n")
							g.Printf("}\n")
						}
						continue
					}
				}
			}
			g.Printf(serializableFieldTemplate, "d."+field.name, fieldNameUntitled)
		case *ast.IndexExpr:
			x := fieldType.X
			if fieldType, isxFieldSelector := x.(*ast.SelectorExpr); isxFieldSelector { // TODO: we need to refactor this so we can reuse...
				xIdent, xIsIdent := fieldType.X.(*ast.Ident)
				sel := fieldType.Sel
				if xIsIdent && xIdent.Name == "atomic" && sel.Name == "Pointer" {
					g.Printf(atomicPointerFieldTemplate, "d."+field.name, field.name, fieldNameUntitled)
					if field.hasLocker != "" {
						g.Printf("\treturn nil\n")
						g.Printf("}(); err != nil {\n")
						g.Printf("\treturn err\n")
						g.Printf("}\n")
					}
					continue
				}
			}
			fmt.Printf("no support yet for %#q\n", fieldType)
			continue
		case *ast.Ident:
			switch fieldType.Name {
			case "byte":
				g.Printf(byteFieldSerialize, "d."+field.name, fieldNameUntitled)
			case "int":
				g.Printf(int64FieldSerialize, "int64(d."+field.name+")", fieldNameUntitled)
			case "int32":
				g.Printf(int32FieldSerialize, "int32(d."+field.name+")", fieldNameUntitled)
			case "uint32":
				g.Printf(uint32FieldSerialize, "d."+field.name, fieldNameUntitled)
			case "bool":
				g.Printf(boolFieldSerialize, "d."+field.name, fieldNameUntitled)
			case "string":
				g.Printf(stringFieldSerialize, "d."+field.name, fieldNameUntitled)
			case "error":
				g.Printf(errorFieldSerialize, "d."+field.name, fieldNameUntitled)
			default:
				fmt.Printf("\t no support implemented for Ident with type %v\n", fieldType)
				g.Printf("{\n")
				g.Printf("_value := fmt.Sprintf(\"%%v\", d.%s)\n", fieldName)
				g.Printf(stringFieldSerialize, "_value", fieldNameUntitled)
				g.Printf("}\n")
			}
		case *ast.ArrayType:
			if eltType, ok := fieldType.Elt.(*ast.Ident); ok && eltType.Name == "byte" {
				g.Printf("if err := writeBuffer.WriteByteArray(%s, d.%s); err != nil {\n", fieldNameUntitled, field.name)
				g.Printf("\treturn err\n")
				g.Printf("}\n")
			} else {
				g.Printf("if err := writeBuffer.PushContext(%s, utils.WithRenderAsList(true)); err != nil {\n\t\treturn err\n\t}\n", fieldNameUntitled)
				g.Printf("for _, elem := range d.%s {", field.name)
				switch eltType := fieldType.Elt.(type) {
				case *ast.SelectorExpr, *ast.StarExpr:
					g.Printf("\n\t\tvar elem any = elem\n")
					g.Printf(serializableFieldTemplate, "elem", "\"value\"")
				case *ast.Ident:
					switch eltType.Name {
					case "int":
						g.Printf(int64FieldSerialize, "int64(d."+field.name+")", fieldNameUntitled)
					case "uint32":
						g.Printf(uint32FieldSerialize, "d."+field.name, fieldNameUntitled)
					case "bool":
						g.Printf(boolFieldSerialize, "elem", "\"\"")
					case "string":
						g.Printf(stringFieldSerialize, "elem", "\"\"")
					case "error":
						g.Printf(errorFieldSerialize, "elem", "\"\"")
					default:
						fmt.Printf("\t no support implemented for Ident within ArrayType for %v\n", fieldType)
						g.Printf("_value := fmt.Sprintf(\"%%v\", elem)\n")
						g.Printf(stringFieldSerialize, "_value", fieldNameUntitled)
					}
				}
				g.Printf("}\n")
				g.Printf("if err := writeBuffer.PopContext(%s, utils.WithRenderAsList(true)); err != nil {\n\t\treturn err\n\t}\n", fieldNameUntitled)
			}
		case *ast.MapType:
			g.Printf("if err := writeBuffer.PushContext(%s, utils.WithRenderAsList(true)); err != nil {\n\t\treturn err\n\t}\n", fieldNameUntitled)
			// TODO: we use serializable or strings as we don't want to over-complex this
			g.Printf("for _name, elem := range d.%s {\n", fieldName)
			switch keyType := fieldType.Key.(type) {
			case *ast.Ident:
				switch keyType.Name {
				case "uint", "uint8", "uint16", "uint32", "uint64", "int", "int8", "int16", "int32", "int64": // TODO: add other types
					g.Printf("\t\tname := fmt.Sprintf(\"%s\", _name)\n", "%v")
				case "string":
					g.Printf("\t\tname := _name\n")
				default:
					g.Printf("\t\tname := fmt.Sprintf(\"%s\", &_name)\n", "%v")
				}
			default:
				g.Printf("\t\tname := fmt.Sprintf(\"%s\", &_name)\n", "%v")
			}
			switch eltType := fieldType.Value.(type) {
			case *ast.StarExpr, *ast.SelectorExpr:
				g.Printf("\n\t\tvar elem any = elem\n")
				g.Printf("\t\tif serializable, ok := elem.(utils.Serializable); ok {\n")
				g.Printf("\t\t\tif err := writeBuffer.PushContext(name); err != nil {\n")
				g.Printf("\t\t\t\treturn err\n")
				g.Printf("\t\t\t}\n")
				g.Printf("\t\t\tif err := serializable.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {\n")
				g.Printf("\t\t\t\treturn err\n")
				g.Printf("\t\t\t}\n")
				g.Printf("\t\t\tif err := writeBuffer.PopContext(name); err != nil {\n")
				g.Printf("\t\t\t\treturn err\n")
				g.Printf("\t\t\t}\n")
				g.Printf("\t\t} else {\n")
				g.Printf("\t\t\telemAsString := fmt.Sprintf(\"%%v\", elem)\n")
				g.Printf("\t\t\tif err := writeBuffer.WriteString(name, uint32(len(elemAsString)*8), elemAsString); err != nil {\n")
				g.Printf("\t\t\t\treturn err\n")
				g.Printf("\t\t\t}\n")
				g.Printf("\t\t}\n")
			case *ast.Ident:
				switch eltType.Name {
				case "bool":
					g.Printf(boolFieldSerialize, "elem", "name")
				case "string":
					g.Printf(stringFieldSerialize, "elem", "name")
				case "error":
					g.Printf(errorFieldSerialize, "elem", "name")
				default:
					fmt.Printf("\t no support implemented for Ident within MapType for %v\n", fieldType)
					g.Printf("\t\t_value := fmt.Sprintf(\"%%v\", elem)\n")
					g.Printf(stringFieldSerialize, "_value", "name")
				}
			default:
				fmt.Printf("\t no support implemented within MapType %v\n", fieldType.Value)
				g.Printf("\t\t_value := fmt.Sprintf(\"%%v\", elem)\n")
				g.Printf(stringFieldSerialize, "_value", "name")
			}
			g.Printf("\t}\n")
			g.Printf("if err := writeBuffer.PopContext(%s, utils.WithRenderAsList(true)); err != nil {\n\t\treturn err\n\t}\n", fieldNameUntitled)
		case *ast.ChanType:
			g.Printf(chanFieldSerialize, "d."+field.name, fieldNameUntitled, field.name)
		case *ast.FuncType:
			g.Printf(funcFieldSerialize, "d."+field.name, fieldNameUntitled)
		default:
			fmt.Printf("no support implemented %#v\n", fieldType)
		}
		if field.hasLocker != "" {
			g.Printf("\treturn nil\n")
			g.Printf("}(); err != nil {\n")
			g.Printf("\treturn err\n")
			g.Printf("}\n")
		}
	}
	g.Printf("\tif err := writeBuffer.PopContext(%s); err != nil {\n", logicalTypeName)
	g.Printf("\t\treturn err\n")
	g.Printf("\t}\n")
	g.Printf("\treturn nil\n")
	g.Printf("}\n")
	g.Printf("\n")
	g.Printf(stringerTemplate, typeName)
}

// format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}

// Field represents a declared field.
type Field struct {
	name       string
	fieldType  ast.Expr
	isDelegate bool
	isStringer bool
	hasLocker  string
}

func (f *Field) String() string {
	return f.name
}

// genDecl processes one declaration clause.
func (f *File) genDecl(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if !ok || decl.Tok != token.TYPE {
		// We only care about type declarations.
		return true
	}
	for _, spec := range decl.Specs {
		typeSpec := spec.(*ast.TypeSpec)
		structDecl, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			continue
		}
		if typeSpec.Name.Name != f.typeName {
			continue
		}
		fmt.Printf("Handling %s\n", typeSpec.Name.Name)
		for _, field := range structDecl.Fields.List {
			if field.Tag != nil && field.Tag.Value == "`ignore:\"true\"`" {
				var name string
				if len(field.Names) != 0 {
					name = field.Names[0].Name
				} else {
					name = "<delegate>"
				}
				fmt.Printf("\t ignoring field %s %v\n", name, field.Type)
				continue
			}
			isStringer := false
			if field.Tag != nil && field.Tag.Value == "`stringer:\"true\"`" { // TODO: Check if we do that a bit smarter
				isStringer = true
			}
			hasLocker := ""
			if field.Tag != nil && strings.HasPrefix(field.Tag.Value, "`hasLocker:\"") { // TODO: Check if we do that a bit smarter
				hasLocker = strings.TrimPrefix(field.Tag.Value, "`hasLocker:\"")
				hasLocker = strings.TrimSuffix(hasLocker, "\"`")
			}
			if len(field.Names) == 0 {
				fmt.Printf("\t adding delegate\n")
				switch ft := field.Type.(type) {
				case *ast.Ident:
					f.fields = append(f.fields, Field{
						fieldType:  ft,
						isDelegate: true,
						isStringer: isStringer,
						hasLocker:  hasLocker,
					})
					continue
				case *ast.StarExpr:
					switch set := ft.X.(type) {
					case *ast.Ident:
						f.fields = append(f.fields, Field{
							fieldType:  set,
							isDelegate: true,
							isStringer: isStringer,
							hasLocker:  hasLocker,
						})
						continue
					default:
						panic(fmt.Sprintf("Only pointer to struct delegates supported now. Type %T", field.Type))
					}
				case *ast.SelectorExpr:
					f.fields = append(f.fields, Field{
						fieldType:  ft.Sel,
						isDelegate: true,
						isStringer: isStringer,
						hasLocker:  hasLocker,
					})
					continue
				default:
					panic(fmt.Sprintf("Only struct delegates supported now. Type %T", field.Type))
				}
			}
			fmt.Printf("\t adding field %s %v\n", field.Names[0].Name, field.Type)
			f.fields = append(f.fields, Field{
				name:       field.Names[0].Name,
				fieldType:  field.Type,
				isStringer: isStringer,
				hasLocker:  hasLocker,
			})
		}
	}
	return false
}

var asfHeader = `/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

`

var stringerTemplate = `
func (d *%s) String() string {
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
`

var serializableFieldTemplate = `
	if %[1]s != nil {
		if serializableField, ok := any(%[1]s).(utils.Serializable); ok {
			if err := writeBuffer.PushContext(%[2]s); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext(%[2]s); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%%v", %[1]s)
			if err := writeBuffer.WriteString(%[2]s, uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}
`

var atomicPointerFieldTemplate = `
	if %[2]sLoaded :=%[1]s.Load(); %[2]sLoaded != nil && *%[2]sLoaded != nil {
		%[2]s := *%[2]sLoaded
		if serializableField, ok := %[2]s.(utils.Serializable); ok {
			if err := writeBuffer.PushContext(%[3]s); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext(%[3]s); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%%v", %[2]s)
			if err := writeBuffer.WriteString(%[3]s, uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}
`

var byteFieldSerialize = `
	if err := writeBuffer.WriteByte(%[2]s, %[1]s); err != nil {
		return err
	}
`

var int32FieldSerialize = `
	if err := writeBuffer.WriteInt32(%[2]s, 32, %[1]s); err != nil {
		return err
	}
`

var int64FieldSerialize = `
	if err := writeBuffer.WriteInt64(%[2]s, 64, %[1]s); err != nil {
		return err
	}
`

var uint32FieldSerialize = `
	if err := writeBuffer.WriteUint32(%[2]s, 32, %[1]s); err != nil {
		return err
	}
`

var uint64FieldSerialize = `
	if err := writeBuffer.WriteUint64(%[2]s, 64, %[1]s); err != nil {
		return err
	}
`

var boolFieldSerialize = `
	if err := writeBuffer.WriteBit(%[2]s, %[1]s); err != nil {
		return err
	}
`

var stringFieldSerialize = `
	if err := writeBuffer.WriteString(%[2]s, uint32(len(%[1]s)*8), %[1]s); err != nil {
		return err
	}
`

var errorFieldSerialize = `
	if %[1]s != nil {
		_errString := %[1]s.Error()
		if err := writeBuffer.WriteString(%[2]s, uint32(len(_errString)*8), _errString); err != nil {
			return err
		}
	}
`

var chanFieldSerialize = `
	_%[3]s_plx4gen_description := fmt.Sprintf("%%d element(s)", len(%[1]s))
    if err := writeBuffer.WriteString(%[2]s, uint32(len(_%[3]s_plx4gen_description)*8), _%[3]s_plx4gen_description); err != nil {
		return err
	}
`

var funcFieldSerialize = `
	if err := writeBuffer.WriteBit(%[2]s, %[1]s != nil); err != nil {
		return err
	}
`

func unTitle(value string) string {
	return strings.ToLower(string(value[0])) + value[1:]
}
