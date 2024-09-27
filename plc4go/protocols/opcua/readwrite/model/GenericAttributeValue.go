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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// GenericAttributeValue is the corresponding interface of GenericAttributeValue
type GenericAttributeValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetAttributeId returns AttributeId (property field)
	GetAttributeId() uint32
	// GetValue returns Value (property field)
	GetValue() Variant
	// IsGenericAttributeValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsGenericAttributeValue()
	// CreateBuilder creates a GenericAttributeValueBuilder
	CreateGenericAttributeValueBuilder() GenericAttributeValueBuilder
}

// _GenericAttributeValue is the data-structure of this message
type _GenericAttributeValue struct {
	ExtensionObjectDefinitionContract
	AttributeId uint32
	Value       Variant
}

var _ GenericAttributeValue = (*_GenericAttributeValue)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_GenericAttributeValue)(nil)

// NewGenericAttributeValue factory function for _GenericAttributeValue
func NewGenericAttributeValue(attributeId uint32, value Variant) *_GenericAttributeValue {
	if value == nil {
		panic("value of type Variant for GenericAttributeValue must not be nil")
	}
	_result := &_GenericAttributeValue{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		AttributeId:                       attributeId,
		Value:                             value,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// GenericAttributeValueBuilder is a builder for GenericAttributeValue
type GenericAttributeValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(attributeId uint32, value Variant) GenericAttributeValueBuilder
	// WithAttributeId adds AttributeId (property field)
	WithAttributeId(uint32) GenericAttributeValueBuilder
	// WithValue adds Value (property field)
	WithValue(Variant) GenericAttributeValueBuilder
	// Build builds the GenericAttributeValue or returns an error if something is wrong
	Build() (GenericAttributeValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() GenericAttributeValue
}

// NewGenericAttributeValueBuilder() creates a GenericAttributeValueBuilder
func NewGenericAttributeValueBuilder() GenericAttributeValueBuilder {
	return &_GenericAttributeValueBuilder{_GenericAttributeValue: new(_GenericAttributeValue)}
}

type _GenericAttributeValueBuilder struct {
	*_GenericAttributeValue

	err *utils.MultiError
}

var _ (GenericAttributeValueBuilder) = (*_GenericAttributeValueBuilder)(nil)

func (m *_GenericAttributeValueBuilder) WithMandatoryFields(attributeId uint32, value Variant) GenericAttributeValueBuilder {
	return m.WithAttributeId(attributeId).WithValue(value)
}

func (m *_GenericAttributeValueBuilder) WithAttributeId(attributeId uint32) GenericAttributeValueBuilder {
	m.AttributeId = attributeId
	return m
}

func (m *_GenericAttributeValueBuilder) WithValue(value Variant) GenericAttributeValueBuilder {
	m.Value = value
	return m
}

func (m *_GenericAttributeValueBuilder) Build() (GenericAttributeValue, error) {
	if m.Value == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'value' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._GenericAttributeValue.deepCopy(), nil
}

func (m *_GenericAttributeValueBuilder) MustBuild() GenericAttributeValue {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_GenericAttributeValueBuilder) DeepCopy() any {
	return m.CreateGenericAttributeValueBuilder()
}

// CreateGenericAttributeValueBuilder creates a GenericAttributeValueBuilder
func (m *_GenericAttributeValue) CreateGenericAttributeValueBuilder() GenericAttributeValueBuilder {
	if m == nil {
		return NewGenericAttributeValueBuilder()
	}
	return &_GenericAttributeValueBuilder{_GenericAttributeValue: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GenericAttributeValue) GetIdentifier() string {
	return "17608"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GenericAttributeValue) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GenericAttributeValue) GetAttributeId() uint32 {
	return m.AttributeId
}

func (m *_GenericAttributeValue) GetValue() Variant {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastGenericAttributeValue(structType any) GenericAttributeValue {
	if casted, ok := structType.(GenericAttributeValue); ok {
		return casted
	}
	if casted, ok := structType.(*GenericAttributeValue); ok {
		return *casted
	}
	return nil
}

func (m *_GenericAttributeValue) GetTypeName() string {
	return "GenericAttributeValue"
}

func (m *_GenericAttributeValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (attributeId)
	lengthInBits += 32

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_GenericAttributeValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_GenericAttributeValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__genericAttributeValue GenericAttributeValue, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GenericAttributeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GenericAttributeValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	attributeId, err := ReadSimpleField(ctx, "attributeId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'attributeId' field"))
	}
	m.AttributeId = attributeId

	value, err := ReadSimpleField[Variant](ctx, "value", ReadComplex[Variant](VariantParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("GenericAttributeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GenericAttributeValue")
	}

	return m, nil
}

func (m *_GenericAttributeValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GenericAttributeValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GenericAttributeValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GenericAttributeValue")
		}

		if err := WriteSimpleField[uint32](ctx, "attributeId", m.GetAttributeId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'attributeId' field")
		}

		if err := WriteSimpleField[Variant](ctx, "value", m.GetValue(), WriteComplex[Variant](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("GenericAttributeValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GenericAttributeValue")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GenericAttributeValue) IsGenericAttributeValue() {}

func (m *_GenericAttributeValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_GenericAttributeValue) deepCopy() *_GenericAttributeValue {
	if m == nil {
		return nil
	}
	_GenericAttributeValueCopy := &_GenericAttributeValue{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.AttributeId,
		m.Value.DeepCopy().(Variant),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _GenericAttributeValueCopy
}

func (m *_GenericAttributeValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
