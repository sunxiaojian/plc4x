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

// ModbusPDUWriteMultipleCoilsResponse is the corresponding interface of ModbusPDUWriteMultipleCoilsResponse
type ModbusPDUWriteMultipleCoilsResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ModbusPDU
	// GetStartingAddress returns StartingAddress (property field)
	GetStartingAddress() uint16
	// GetQuantity returns Quantity (property field)
	GetQuantity() uint16
	// IsModbusPDUWriteMultipleCoilsResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusPDUWriteMultipleCoilsResponse()
	// CreateBuilder creates a ModbusPDUWriteMultipleCoilsResponseBuilder
	CreateModbusPDUWriteMultipleCoilsResponseBuilder() ModbusPDUWriteMultipleCoilsResponseBuilder
}

// _ModbusPDUWriteMultipleCoilsResponse is the data-structure of this message
type _ModbusPDUWriteMultipleCoilsResponse struct {
	ModbusPDUContract
	StartingAddress uint16
	Quantity        uint16
}

var _ ModbusPDUWriteMultipleCoilsResponse = (*_ModbusPDUWriteMultipleCoilsResponse)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUWriteMultipleCoilsResponse)(nil)

// NewModbusPDUWriteMultipleCoilsResponse factory function for _ModbusPDUWriteMultipleCoilsResponse
func NewModbusPDUWriteMultipleCoilsResponse(startingAddress uint16, quantity uint16) *_ModbusPDUWriteMultipleCoilsResponse {
	_result := &_ModbusPDUWriteMultipleCoilsResponse{
		ModbusPDUContract: NewModbusPDU(),
		StartingAddress:   startingAddress,
		Quantity:          quantity,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ModbusPDUWriteMultipleCoilsResponseBuilder is a builder for ModbusPDUWriteMultipleCoilsResponse
type ModbusPDUWriteMultipleCoilsResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(startingAddress uint16, quantity uint16) ModbusPDUWriteMultipleCoilsResponseBuilder
	// WithStartingAddress adds StartingAddress (property field)
	WithStartingAddress(uint16) ModbusPDUWriteMultipleCoilsResponseBuilder
	// WithQuantity adds Quantity (property field)
	WithQuantity(uint16) ModbusPDUWriteMultipleCoilsResponseBuilder
	// Build builds the ModbusPDUWriteMultipleCoilsResponse or returns an error if something is wrong
	Build() (ModbusPDUWriteMultipleCoilsResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ModbusPDUWriteMultipleCoilsResponse
}

// NewModbusPDUWriteMultipleCoilsResponseBuilder() creates a ModbusPDUWriteMultipleCoilsResponseBuilder
func NewModbusPDUWriteMultipleCoilsResponseBuilder() ModbusPDUWriteMultipleCoilsResponseBuilder {
	return &_ModbusPDUWriteMultipleCoilsResponseBuilder{_ModbusPDUWriteMultipleCoilsResponse: new(_ModbusPDUWriteMultipleCoilsResponse)}
}

type _ModbusPDUWriteMultipleCoilsResponseBuilder struct {
	*_ModbusPDUWriteMultipleCoilsResponse

	err *utils.MultiError
}

var _ (ModbusPDUWriteMultipleCoilsResponseBuilder) = (*_ModbusPDUWriteMultipleCoilsResponseBuilder)(nil)

func (m *_ModbusPDUWriteMultipleCoilsResponseBuilder) WithMandatoryFields(startingAddress uint16, quantity uint16) ModbusPDUWriteMultipleCoilsResponseBuilder {
	return m.WithStartingAddress(startingAddress).WithQuantity(quantity)
}

func (m *_ModbusPDUWriteMultipleCoilsResponseBuilder) WithStartingAddress(startingAddress uint16) ModbusPDUWriteMultipleCoilsResponseBuilder {
	m.StartingAddress = startingAddress
	return m
}

func (m *_ModbusPDUWriteMultipleCoilsResponseBuilder) WithQuantity(quantity uint16) ModbusPDUWriteMultipleCoilsResponseBuilder {
	m.Quantity = quantity
	return m
}

func (m *_ModbusPDUWriteMultipleCoilsResponseBuilder) Build() (ModbusPDUWriteMultipleCoilsResponse, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ModbusPDUWriteMultipleCoilsResponse.deepCopy(), nil
}

func (m *_ModbusPDUWriteMultipleCoilsResponseBuilder) MustBuild() ModbusPDUWriteMultipleCoilsResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ModbusPDUWriteMultipleCoilsResponseBuilder) DeepCopy() any {
	return m.CreateModbusPDUWriteMultipleCoilsResponseBuilder()
}

// CreateModbusPDUWriteMultipleCoilsResponseBuilder creates a ModbusPDUWriteMultipleCoilsResponseBuilder
func (m *_ModbusPDUWriteMultipleCoilsResponse) CreateModbusPDUWriteMultipleCoilsResponseBuilder() ModbusPDUWriteMultipleCoilsResponseBuilder {
	if m == nil {
		return NewModbusPDUWriteMultipleCoilsResponseBuilder()
	}
	return &_ModbusPDUWriteMultipleCoilsResponseBuilder{_ModbusPDUWriteMultipleCoilsResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetFunctionFlag() uint8 {
	return 0x0F
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetStartingAddress() uint16 {
	return m.StartingAddress
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetQuantity() uint16 {
	return m.Quantity
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteMultipleCoilsResponse(structType any) ModbusPDUWriteMultipleCoilsResponse {
	if casted, ok := structType.(ModbusPDUWriteMultipleCoilsResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteMultipleCoilsResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetTypeName() string {
	return "ModbusPDUWriteMultipleCoilsResponse"
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUWriteMultipleCoilsResponse ModbusPDUWriteMultipleCoilsResponse, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteMultipleCoilsResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteMultipleCoilsResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	startingAddress, err := ReadSimpleField(ctx, "startingAddress", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startingAddress' field"))
	}
	m.StartingAddress = startingAddress

	quantity, err := ReadSimpleField(ctx, "quantity", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'quantity' field"))
	}
	m.Quantity = quantity

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteMultipleCoilsResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteMultipleCoilsResponse")
	}

	return m, nil
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteMultipleCoilsResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteMultipleCoilsResponse")
		}

		if err := WriteSimpleField[uint16](ctx, "startingAddress", m.GetStartingAddress(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'startingAddress' field")
		}

		if err := WriteSimpleField[uint16](ctx, "quantity", m.GetQuantity(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'quantity' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteMultipleCoilsResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUWriteMultipleCoilsResponse")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) IsModbusPDUWriteMultipleCoilsResponse() {}

func (m *_ModbusPDUWriteMultipleCoilsResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) deepCopy() *_ModbusPDUWriteMultipleCoilsResponse {
	if m == nil {
		return nil
	}
	_ModbusPDUWriteMultipleCoilsResponseCopy := &_ModbusPDUWriteMultipleCoilsResponse{
		m.ModbusPDUContract.(*_ModbusPDU).deepCopy(),
		m.StartingAddress,
		m.Quantity,
	}
	m.ModbusPDUContract.(*_ModbusPDU)._SubType = m
	return _ModbusPDUWriteMultipleCoilsResponseCopy
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
