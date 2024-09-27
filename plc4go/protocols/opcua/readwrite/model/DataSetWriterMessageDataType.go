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

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// DataSetWriterMessageDataType is the corresponding interface of DataSetWriterMessageDataType
type DataSetWriterMessageDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// IsDataSetWriterMessageDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDataSetWriterMessageDataType()
	// CreateBuilder creates a DataSetWriterMessageDataTypeBuilder
	CreateDataSetWriterMessageDataTypeBuilder() DataSetWriterMessageDataTypeBuilder
}

// _DataSetWriterMessageDataType is the data-structure of this message
type _DataSetWriterMessageDataType struct {
	ExtensionObjectDefinitionContract
}

var _ DataSetWriterMessageDataType = (*_DataSetWriterMessageDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DataSetWriterMessageDataType)(nil)

// NewDataSetWriterMessageDataType factory function for _DataSetWriterMessageDataType
func NewDataSetWriterMessageDataType() *_DataSetWriterMessageDataType {
	_result := &_DataSetWriterMessageDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DataSetWriterMessageDataTypeBuilder is a builder for DataSetWriterMessageDataType
type DataSetWriterMessageDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() DataSetWriterMessageDataTypeBuilder
	// Build builds the DataSetWriterMessageDataType or returns an error if something is wrong
	Build() (DataSetWriterMessageDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DataSetWriterMessageDataType
}

// NewDataSetWriterMessageDataTypeBuilder() creates a DataSetWriterMessageDataTypeBuilder
func NewDataSetWriterMessageDataTypeBuilder() DataSetWriterMessageDataTypeBuilder {
	return &_DataSetWriterMessageDataTypeBuilder{_DataSetWriterMessageDataType: new(_DataSetWriterMessageDataType)}
}

type _DataSetWriterMessageDataTypeBuilder struct {
	*_DataSetWriterMessageDataType

	err *utils.MultiError
}

var _ (DataSetWriterMessageDataTypeBuilder) = (*_DataSetWriterMessageDataTypeBuilder)(nil)

func (m *_DataSetWriterMessageDataTypeBuilder) WithMandatoryFields() DataSetWriterMessageDataTypeBuilder {
	return m
}

func (m *_DataSetWriterMessageDataTypeBuilder) Build() (DataSetWriterMessageDataType, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._DataSetWriterMessageDataType.deepCopy(), nil
}

func (m *_DataSetWriterMessageDataTypeBuilder) MustBuild() DataSetWriterMessageDataType {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_DataSetWriterMessageDataTypeBuilder) DeepCopy() any {
	return m.CreateDataSetWriterMessageDataTypeBuilder()
}

// CreateDataSetWriterMessageDataTypeBuilder creates a DataSetWriterMessageDataTypeBuilder
func (m *_DataSetWriterMessageDataType) CreateDataSetWriterMessageDataTypeBuilder() DataSetWriterMessageDataTypeBuilder {
	if m == nil {
		return NewDataSetWriterMessageDataTypeBuilder()
	}
	return &_DataSetWriterMessageDataTypeBuilder{_DataSetWriterMessageDataType: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataSetWriterMessageDataType) GetIdentifier() string {
	return "15607"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataSetWriterMessageDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastDataSetWriterMessageDataType(structType any) DataSetWriterMessageDataType {
	if casted, ok := structType.(DataSetWriterMessageDataType); ok {
		return casted
	}
	if casted, ok := structType.(*DataSetWriterMessageDataType); ok {
		return *casted
	}
	return nil
}

func (m *_DataSetWriterMessageDataType) GetTypeName() string {
	return "DataSetWriterMessageDataType"
}

func (m *_DataSetWriterMessageDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_DataSetWriterMessageDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DataSetWriterMessageDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__dataSetWriterMessageDataType DataSetWriterMessageDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataSetWriterMessageDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataSetWriterMessageDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("DataSetWriterMessageDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataSetWriterMessageDataType")
	}

	return m, nil
}

func (m *_DataSetWriterMessageDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataSetWriterMessageDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataSetWriterMessageDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataSetWriterMessageDataType")
		}

		if popErr := writeBuffer.PopContext("DataSetWriterMessageDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataSetWriterMessageDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataSetWriterMessageDataType) IsDataSetWriterMessageDataType() {}

func (m *_DataSetWriterMessageDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DataSetWriterMessageDataType) deepCopy() *_DataSetWriterMessageDataType {
	if m == nil {
		return nil
	}
	_DataSetWriterMessageDataTypeCopy := &_DataSetWriterMessageDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DataSetWriterMessageDataTypeCopy
}

func (m *_DataSetWriterMessageDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
