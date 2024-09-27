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

// ByteStringArray is the corresponding interface of ByteStringArray
type ByteStringArray interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() int32
	// GetValue returns Value (property field)
	GetValue() []uint8
	// IsByteStringArray is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsByteStringArray()
	// CreateBuilder creates a ByteStringArrayBuilder
	CreateByteStringArrayBuilder() ByteStringArrayBuilder
}

// _ByteStringArray is the data-structure of this message
type _ByteStringArray struct {
	ArrayLength int32
	Value       []uint8
}

var _ ByteStringArray = (*_ByteStringArray)(nil)

// NewByteStringArray factory function for _ByteStringArray
func NewByteStringArray(arrayLength int32, value []uint8) *_ByteStringArray {
	return &_ByteStringArray{ArrayLength: arrayLength, Value: value}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ByteStringArrayBuilder is a builder for ByteStringArray
type ByteStringArrayBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(arrayLength int32, value []uint8) ByteStringArrayBuilder
	// WithArrayLength adds ArrayLength (property field)
	WithArrayLength(int32) ByteStringArrayBuilder
	// WithValue adds Value (property field)
	WithValue(...uint8) ByteStringArrayBuilder
	// Build builds the ByteStringArray or returns an error if something is wrong
	Build() (ByteStringArray, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ByteStringArray
}

// NewByteStringArrayBuilder() creates a ByteStringArrayBuilder
func NewByteStringArrayBuilder() ByteStringArrayBuilder {
	return &_ByteStringArrayBuilder{_ByteStringArray: new(_ByteStringArray)}
}

type _ByteStringArrayBuilder struct {
	*_ByteStringArray

	err *utils.MultiError
}

var _ (ByteStringArrayBuilder) = (*_ByteStringArrayBuilder)(nil)

func (m *_ByteStringArrayBuilder) WithMandatoryFields(arrayLength int32, value []uint8) ByteStringArrayBuilder {
	return m.WithArrayLength(arrayLength).WithValue(value...)
}

func (m *_ByteStringArrayBuilder) WithArrayLength(arrayLength int32) ByteStringArrayBuilder {
	m.ArrayLength = arrayLength
	return m
}

func (m *_ByteStringArrayBuilder) WithValue(value ...uint8) ByteStringArrayBuilder {
	m.Value = value
	return m
}

func (m *_ByteStringArrayBuilder) Build() (ByteStringArray, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ByteStringArray.deepCopy(), nil
}

func (m *_ByteStringArrayBuilder) MustBuild() ByteStringArray {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ByteStringArrayBuilder) DeepCopy() any {
	return m.CreateByteStringArrayBuilder()
}

// CreateByteStringArrayBuilder creates a ByteStringArrayBuilder
func (m *_ByteStringArray) CreateByteStringArrayBuilder() ByteStringArrayBuilder {
	if m == nil {
		return NewByteStringArrayBuilder()
	}
	return &_ByteStringArrayBuilder{_ByteStringArray: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ByteStringArray) GetArrayLength() int32 {
	return m.ArrayLength
}

func (m *_ByteStringArray) GetValue() []uint8 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastByteStringArray(structType any) ByteStringArray {
	if casted, ok := structType.(ByteStringArray); ok {
		return casted
	}
	if casted, ok := structType.(*ByteStringArray); ok {
		return *casted
	}
	return nil
}

func (m *_ByteStringArray) GetTypeName() string {
	return "ByteStringArray"
}

func (m *_ByteStringArray) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (arrayLength)
	lengthInBits += 32

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_ByteStringArray) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ByteStringArrayParse(ctx context.Context, theBytes []byte) (ByteStringArray, error) {
	return ByteStringArrayParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ByteStringArrayParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ByteStringArray, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ByteStringArray, error) {
		return ByteStringArrayParseWithBuffer(ctx, readBuffer)
	}
}

func ByteStringArrayParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ByteStringArray, error) {
	v, err := (&_ByteStringArray{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_ByteStringArray) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__byteStringArray ByteStringArray, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ByteStringArray"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ByteStringArray")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	arrayLength, err := ReadSimpleField(ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}
	m.ArrayLength = arrayLength

	value, err := ReadCountArrayField[uint8](ctx, "value", ReadUnsignedByte(readBuffer, uint8(8)), uint64(arrayLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("ByteStringArray"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ByteStringArray")
	}

	return m, nil
}

func (m *_ByteStringArray) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ByteStringArray) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ByteStringArray"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ByteStringArray")
	}

	if err := WriteSimpleField[int32](ctx, "arrayLength", m.GetArrayLength(), WriteSignedInt(writeBuffer, 32)); err != nil {
		return errors.Wrap(err, "Error serializing 'arrayLength' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "value", m.GetValue(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("ByteStringArray"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ByteStringArray")
	}
	return nil
}

func (m *_ByteStringArray) IsByteStringArray() {}

func (m *_ByteStringArray) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ByteStringArray) deepCopy() *_ByteStringArray {
	if m == nil {
		return nil
	}
	_ByteStringArrayCopy := &_ByteStringArray{
		m.ArrayLength,
		utils.DeepCopySlice[uint8, uint8](m.Value),
	}
	return _ByteStringArrayCopy
}

func (m *_ByteStringArray) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
