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

// ApduDataExtAuthorizeRequest is the corresponding interface of ApduDataExtAuthorizeRequest
type ApduDataExtAuthorizeRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// GetLevel returns Level (property field)
	GetLevel() uint8
	// GetData returns Data (property field)
	GetData() []byte
	// IsApduDataExtAuthorizeRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtAuthorizeRequest()
	// CreateBuilder creates a ApduDataExtAuthorizeRequestBuilder
	CreateApduDataExtAuthorizeRequestBuilder() ApduDataExtAuthorizeRequestBuilder
}

// _ApduDataExtAuthorizeRequest is the data-structure of this message
type _ApduDataExtAuthorizeRequest struct {
	ApduDataExtContract
	Level uint8
	Data  []byte
}

var _ ApduDataExtAuthorizeRequest = (*_ApduDataExtAuthorizeRequest)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtAuthorizeRequest)(nil)

// NewApduDataExtAuthorizeRequest factory function for _ApduDataExtAuthorizeRequest
func NewApduDataExtAuthorizeRequest(level uint8, data []byte, length uint8) *_ApduDataExtAuthorizeRequest {
	_result := &_ApduDataExtAuthorizeRequest{
		ApduDataExtContract: NewApduDataExt(length),
		Level:               level,
		Data:                data,
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataExtAuthorizeRequestBuilder is a builder for ApduDataExtAuthorizeRequest
type ApduDataExtAuthorizeRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(level uint8, data []byte) ApduDataExtAuthorizeRequestBuilder
	// WithLevel adds Level (property field)
	WithLevel(uint8) ApduDataExtAuthorizeRequestBuilder
	// WithData adds Data (property field)
	WithData(...byte) ApduDataExtAuthorizeRequestBuilder
	// Build builds the ApduDataExtAuthorizeRequest or returns an error if something is wrong
	Build() (ApduDataExtAuthorizeRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataExtAuthorizeRequest
}

// NewApduDataExtAuthorizeRequestBuilder() creates a ApduDataExtAuthorizeRequestBuilder
func NewApduDataExtAuthorizeRequestBuilder() ApduDataExtAuthorizeRequestBuilder {
	return &_ApduDataExtAuthorizeRequestBuilder{_ApduDataExtAuthorizeRequest: new(_ApduDataExtAuthorizeRequest)}
}

type _ApduDataExtAuthorizeRequestBuilder struct {
	*_ApduDataExtAuthorizeRequest

	err *utils.MultiError
}

var _ (ApduDataExtAuthorizeRequestBuilder) = (*_ApduDataExtAuthorizeRequestBuilder)(nil)

func (m *_ApduDataExtAuthorizeRequestBuilder) WithMandatoryFields(level uint8, data []byte) ApduDataExtAuthorizeRequestBuilder {
	return m.WithLevel(level).WithData(data...)
}

func (m *_ApduDataExtAuthorizeRequestBuilder) WithLevel(level uint8) ApduDataExtAuthorizeRequestBuilder {
	m.Level = level
	return m
}

func (m *_ApduDataExtAuthorizeRequestBuilder) WithData(data ...byte) ApduDataExtAuthorizeRequestBuilder {
	m.Data = data
	return m
}

func (m *_ApduDataExtAuthorizeRequestBuilder) Build() (ApduDataExtAuthorizeRequest, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ApduDataExtAuthorizeRequest.deepCopy(), nil
}

func (m *_ApduDataExtAuthorizeRequestBuilder) MustBuild() ApduDataExtAuthorizeRequest {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ApduDataExtAuthorizeRequestBuilder) DeepCopy() any {
	return m.CreateApduDataExtAuthorizeRequestBuilder()
}

// CreateApduDataExtAuthorizeRequestBuilder creates a ApduDataExtAuthorizeRequestBuilder
func (m *_ApduDataExtAuthorizeRequest) CreateApduDataExtAuthorizeRequestBuilder() ApduDataExtAuthorizeRequestBuilder {
	if m == nil {
		return NewApduDataExtAuthorizeRequestBuilder()
	}
	return &_ApduDataExtAuthorizeRequestBuilder{_ApduDataExtAuthorizeRequest: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtAuthorizeRequest) GetExtApciType() uint8 {
	return 0x11
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtAuthorizeRequest) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataExtAuthorizeRequest) GetLevel() uint8 {
	return m.Level
}

func (m *_ApduDataExtAuthorizeRequest) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastApduDataExtAuthorizeRequest(structType any) ApduDataExtAuthorizeRequest {
	if casted, ok := structType.(ApduDataExtAuthorizeRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtAuthorizeRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtAuthorizeRequest) GetTypeName() string {
	return "ApduDataExtAuthorizeRequest"
}

func (m *_ApduDataExtAuthorizeRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	// Simple field (level)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ApduDataExtAuthorizeRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtAuthorizeRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtAuthorizeRequest ApduDataExtAuthorizeRequest, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtAuthorizeRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtAuthorizeRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	level, err := ReadSimpleField(ctx, "level", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'level' field"))
	}
	m.Level = level

	data, err := readBuffer.ReadByteArray("data", int(int32(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("ApduDataExtAuthorizeRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtAuthorizeRequest")
	}

	return m, nil
}

func (m *_ApduDataExtAuthorizeRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtAuthorizeRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtAuthorizeRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtAuthorizeRequest")
		}

		if err := WriteSimpleField[uint8](ctx, "level", m.GetLevel(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'level' field")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtAuthorizeRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtAuthorizeRequest")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtAuthorizeRequest) IsApduDataExtAuthorizeRequest() {}

func (m *_ApduDataExtAuthorizeRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtAuthorizeRequest) deepCopy() *_ApduDataExtAuthorizeRequest {
	if m == nil {
		return nil
	}
	_ApduDataExtAuthorizeRequestCopy := &_ApduDataExtAuthorizeRequest{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
		m.Level,
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	m.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtAuthorizeRequestCopy
}

func (m *_ApduDataExtAuthorizeRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
