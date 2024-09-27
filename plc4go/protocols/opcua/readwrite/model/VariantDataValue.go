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

// VariantDataValue is the corresponding interface of VariantDataValue
type VariantDataValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []DataValue
	// IsVariantDataValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVariantDataValue()
	// CreateBuilder creates a VariantDataValueBuilder
	CreateVariantDataValueBuilder() VariantDataValueBuilder
}

// _VariantDataValue is the data-structure of this message
type _VariantDataValue struct {
	VariantContract
	ArrayLength *int32
	Value       []DataValue
}

var _ VariantDataValue = (*_VariantDataValue)(nil)
var _ VariantRequirements = (*_VariantDataValue)(nil)

// NewVariantDataValue factory function for _VariantDataValue
func NewVariantDataValue(arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool, arrayLength *int32, value []DataValue) *_VariantDataValue {
	_result := &_VariantDataValue{
		VariantContract: NewVariant(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions),
		ArrayLength:     arrayLength,
		Value:           value,
	}
	_result.VariantContract.(*_Variant)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// VariantDataValueBuilder is a builder for VariantDataValue
type VariantDataValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value []DataValue) VariantDataValueBuilder
	// WithArrayLength adds ArrayLength (property field)
	WithOptionalArrayLength(int32) VariantDataValueBuilder
	// WithValue adds Value (property field)
	WithValue(...DataValue) VariantDataValueBuilder
	// Build builds the VariantDataValue or returns an error if something is wrong
	Build() (VariantDataValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() VariantDataValue
}

// NewVariantDataValueBuilder() creates a VariantDataValueBuilder
func NewVariantDataValueBuilder() VariantDataValueBuilder {
	return &_VariantDataValueBuilder{_VariantDataValue: new(_VariantDataValue)}
}

type _VariantDataValueBuilder struct {
	*_VariantDataValue

	err *utils.MultiError
}

var _ (VariantDataValueBuilder) = (*_VariantDataValueBuilder)(nil)

func (m *_VariantDataValueBuilder) WithMandatoryFields(value []DataValue) VariantDataValueBuilder {
	return m.WithValue(value...)
}

func (m *_VariantDataValueBuilder) WithOptionalArrayLength(arrayLength int32) VariantDataValueBuilder {
	m.ArrayLength = &arrayLength
	return m
}

func (m *_VariantDataValueBuilder) WithValue(value ...DataValue) VariantDataValueBuilder {
	m.Value = value
	return m
}

func (m *_VariantDataValueBuilder) Build() (VariantDataValue, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._VariantDataValue.deepCopy(), nil
}

func (m *_VariantDataValueBuilder) MustBuild() VariantDataValue {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_VariantDataValueBuilder) DeepCopy() any {
	return m.CreateVariantDataValueBuilder()
}

// CreateVariantDataValueBuilder creates a VariantDataValueBuilder
func (m *_VariantDataValue) CreateVariantDataValueBuilder() VariantDataValueBuilder {
	if m == nil {
		return NewVariantDataValueBuilder()
	}
	return &_VariantDataValueBuilder{_VariantDataValue: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantDataValue) GetVariantType() uint8 {
	return uint8(23)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantDataValue) GetParent() VariantContract {
	return m.VariantContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantDataValue) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantDataValue) GetValue() []DataValue {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastVariantDataValue(structType any) VariantDataValue {
	if casted, ok := structType.(VariantDataValue); ok {
		return casted
	}
	if casted, ok := structType.(*VariantDataValue); ok {
		return *casted
	}
	return nil
}

func (m *_VariantDataValue) GetTypeName() string {
	return "VariantDataValue"
}

func (m *_VariantDataValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.VariantContract.(*_Variant).getLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		for _curItem, element := range m.Value {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Value), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_VariantDataValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VariantDataValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Variant, arrayLengthSpecified bool) (__variantDataValue VariantDataValue, err error) {
	m.VariantContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VariantDataValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantDataValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var arrayLength *int32
	arrayLength, err = ReadOptionalField[int32](ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)), arrayLengthSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}
	m.ArrayLength = arrayLength

	value, err := ReadCountArrayField[DataValue](ctx, "value", ReadComplex[DataValue](DataValueParseWithBuffer, readBuffer), uint64(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return int32(int32(1)) }, func() any { return int32((*arrayLength)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("VariantDataValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantDataValue")
	}

	return m, nil
}

func (m *_VariantDataValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantDataValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantDataValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantDataValue")
		}

		if err := WriteOptionalField[int32](ctx, "arrayLength", m.GetArrayLength(), WriteSignedInt(writeBuffer, 32), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayLength' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "value", m.GetValue(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("VariantDataValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantDataValue")
		}
		return nil
	}
	return m.VariantContract.(*_Variant).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantDataValue) IsVariantDataValue() {}

func (m *_VariantDataValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_VariantDataValue) deepCopy() *_VariantDataValue {
	if m == nil {
		return nil
	}
	_VariantDataValueCopy := &_VariantDataValue{
		m.VariantContract.(*_Variant).deepCopy(),
		utils.CopyPtr[int32](m.ArrayLength),
		utils.DeepCopySlice[DataValue, DataValue](m.Value),
	}
	m.VariantContract.(*_Variant)._SubType = m
	return _VariantDataValueCopy
}

func (m *_VariantDataValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
