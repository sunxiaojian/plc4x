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

// ApduDataIndividualAddressResponse is the corresponding interface of ApduDataIndividualAddressResponse
type ApduDataIndividualAddressResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduData
	// IsApduDataIndividualAddressResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataIndividualAddressResponse()
	// CreateBuilder creates a ApduDataIndividualAddressResponseBuilder
	CreateApduDataIndividualAddressResponseBuilder() ApduDataIndividualAddressResponseBuilder
}

// _ApduDataIndividualAddressResponse is the data-structure of this message
type _ApduDataIndividualAddressResponse struct {
	ApduDataContract
}

var _ ApduDataIndividualAddressResponse = (*_ApduDataIndividualAddressResponse)(nil)
var _ ApduDataRequirements = (*_ApduDataIndividualAddressResponse)(nil)

// NewApduDataIndividualAddressResponse factory function for _ApduDataIndividualAddressResponse
func NewApduDataIndividualAddressResponse(dataLength uint8) *_ApduDataIndividualAddressResponse {
	_result := &_ApduDataIndividualAddressResponse{
		ApduDataContract: NewApduData(dataLength),
	}
	_result.ApduDataContract.(*_ApduData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataIndividualAddressResponseBuilder is a builder for ApduDataIndividualAddressResponse
type ApduDataIndividualAddressResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataIndividualAddressResponseBuilder
	// Build builds the ApduDataIndividualAddressResponse or returns an error if something is wrong
	Build() (ApduDataIndividualAddressResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataIndividualAddressResponse
}

// NewApduDataIndividualAddressResponseBuilder() creates a ApduDataIndividualAddressResponseBuilder
func NewApduDataIndividualAddressResponseBuilder() ApduDataIndividualAddressResponseBuilder {
	return &_ApduDataIndividualAddressResponseBuilder{_ApduDataIndividualAddressResponse: new(_ApduDataIndividualAddressResponse)}
}

type _ApduDataIndividualAddressResponseBuilder struct {
	*_ApduDataIndividualAddressResponse

	err *utils.MultiError
}

var _ (ApduDataIndividualAddressResponseBuilder) = (*_ApduDataIndividualAddressResponseBuilder)(nil)

func (m *_ApduDataIndividualAddressResponseBuilder) WithMandatoryFields() ApduDataIndividualAddressResponseBuilder {
	return m
}

func (m *_ApduDataIndividualAddressResponseBuilder) Build() (ApduDataIndividualAddressResponse, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ApduDataIndividualAddressResponse.deepCopy(), nil
}

func (m *_ApduDataIndividualAddressResponseBuilder) MustBuild() ApduDataIndividualAddressResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ApduDataIndividualAddressResponseBuilder) DeepCopy() any {
	return m.CreateApduDataIndividualAddressResponseBuilder()
}

// CreateApduDataIndividualAddressResponseBuilder creates a ApduDataIndividualAddressResponseBuilder
func (m *_ApduDataIndividualAddressResponse) CreateApduDataIndividualAddressResponseBuilder() ApduDataIndividualAddressResponseBuilder {
	if m == nil {
		return NewApduDataIndividualAddressResponseBuilder()
	}
	return &_ApduDataIndividualAddressResponseBuilder{_ApduDataIndividualAddressResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataIndividualAddressResponse) GetApciType() uint8 {
	return 0x5
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataIndividualAddressResponse) GetParent() ApduDataContract {
	return m.ApduDataContract
}

// Deprecated: use the interface for direct cast
func CastApduDataIndividualAddressResponse(structType any) ApduDataIndividualAddressResponse {
	if casted, ok := structType.(ApduDataIndividualAddressResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataIndividualAddressResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataIndividualAddressResponse) GetTypeName() string {
	return "ApduDataIndividualAddressResponse"
}

func (m *_ApduDataIndividualAddressResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataContract.(*_ApduData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataIndividualAddressResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataIndividualAddressResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduData, dataLength uint8) (__apduDataIndividualAddressResponse ApduDataIndividualAddressResponse, err error) {
	m.ApduDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataIndividualAddressResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataIndividualAddressResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataIndividualAddressResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataIndividualAddressResponse")
	}

	return m, nil
}

func (m *_ApduDataIndividualAddressResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataIndividualAddressResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataIndividualAddressResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataIndividualAddressResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataIndividualAddressResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataIndividualAddressResponse")
		}
		return nil
	}
	return m.ApduDataContract.(*_ApduData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataIndividualAddressResponse) IsApduDataIndividualAddressResponse() {}

func (m *_ApduDataIndividualAddressResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataIndividualAddressResponse) deepCopy() *_ApduDataIndividualAddressResponse {
	if m == nil {
		return nil
	}
	_ApduDataIndividualAddressResponseCopy := &_ApduDataIndividualAddressResponse{
		m.ApduDataContract.(*_ApduData).deepCopy(),
	}
	m.ApduDataContract.(*_ApduData)._SubType = m
	return _ApduDataIndividualAddressResponseCopy
}

func (m *_ApduDataIndividualAddressResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
