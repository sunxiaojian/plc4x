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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusPDUWriteMultipleCoilsRequest is the corresponding interface of ModbusPDUWriteMultipleCoilsRequest
type ModbusPDUWriteMultipleCoilsRequest interface {
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetStartingAddress returns StartingAddress (property field)
	GetStartingAddress() uint16
	// GetQuantity returns Quantity (property field)
	GetQuantity() uint16
	// GetValue returns Value (property field)
	GetValue() []byte
}

// ModbusPDUWriteMultipleCoilsRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUWriteMultipleCoilsRequest.
// This is useful for switch cases.
type ModbusPDUWriteMultipleCoilsRequestExactly interface {
	ModbusPDUWriteMultipleCoilsRequest
	isModbusPDUWriteMultipleCoilsRequest() bool
}

// _ModbusPDUWriteMultipleCoilsRequest is the data-structure of this message
type _ModbusPDUWriteMultipleCoilsRequest struct {
	*_ModbusPDU
	StartingAddress uint16
	Quantity        uint16
	Value           []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUWriteMultipleCoilsRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUWriteMultipleCoilsRequest) GetFunctionFlag() uint8 {
	return 0x0F
}

func (m *_ModbusPDUWriteMultipleCoilsRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUWriteMultipleCoilsRequest) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUWriteMultipleCoilsRequest) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteMultipleCoilsRequest) GetStartingAddress() uint16 {
	return m.StartingAddress
}

func (m *_ModbusPDUWriteMultipleCoilsRequest) GetQuantity() uint16 {
	return m.Quantity
}

func (m *_ModbusPDUWriteMultipleCoilsRequest) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUWriteMultipleCoilsRequest factory function for _ModbusPDUWriteMultipleCoilsRequest
func NewModbusPDUWriteMultipleCoilsRequest(startingAddress uint16, quantity uint16, value []byte) *_ModbusPDUWriteMultipleCoilsRequest {
	_result := &_ModbusPDUWriteMultipleCoilsRequest{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		Value:           value,
		_ModbusPDU:      NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteMultipleCoilsRequest(structType interface{}) ModbusPDUWriteMultipleCoilsRequest {
	if casted, ok := structType.(ModbusPDUWriteMultipleCoilsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteMultipleCoilsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteMultipleCoilsRequest) GetTypeName() string {
	return "ModbusPDUWriteMultipleCoilsRequest"
}

func (m *_ModbusPDUWriteMultipleCoilsRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ModbusPDUWriteMultipleCoilsRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_ModbusPDUWriteMultipleCoilsRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUWriteMultipleCoilsRequestParse(readBuffer utils.ReadBuffer, response bool) (ModbusPDUWriteMultipleCoilsRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteMultipleCoilsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteMultipleCoilsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (startingAddress)
	_startingAddress, _startingAddressErr := readBuffer.ReadUint16("startingAddress", 16)
	if _startingAddressErr != nil {
		return nil, errors.Wrap(_startingAddressErr, "Error parsing 'startingAddress' field of ModbusPDUWriteMultipleCoilsRequest")
	}
	startingAddress := _startingAddress

	// Simple Field (quantity)
	_quantity, _quantityErr := readBuffer.ReadUint16("quantity", 16)
	if _quantityErr != nil {
		return nil, errors.Wrap(_quantityErr, "Error parsing 'quantity' field of ModbusPDUWriteMultipleCoilsRequest")
	}
	quantity := _quantity

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field of ModbusPDUWriteMultipleCoilsRequest")
	}
	// Byte Array field (value)
	numberOfBytesvalue := int(byteCount)
	value, _readArrayErr := readBuffer.ReadByteArray("value", numberOfBytesvalue)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'value' field of ModbusPDUWriteMultipleCoilsRequest")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteMultipleCoilsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteMultipleCoilsRequest")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUWriteMultipleCoilsRequest{
		_ModbusPDU:      &_ModbusPDU{},
		StartingAddress: startingAddress,
		Quantity:        quantity,
		Value:           value,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUWriteMultipleCoilsRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteMultipleCoilsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteMultipleCoilsRequest")
		}

		// Simple Field (startingAddress)
		startingAddress := uint16(m.GetStartingAddress())
		_startingAddressErr := writeBuffer.WriteUint16("startingAddress", 16, (startingAddress))
		if _startingAddressErr != nil {
			return errors.Wrap(_startingAddressErr, "Error serializing 'startingAddress' field")
		}

		// Simple Field (quantity)
		quantity := uint16(m.GetQuantity())
		_quantityErr := writeBuffer.WriteUint16("quantity", 16, (quantity))
		if _quantityErr != nil {
			return errors.Wrap(_quantityErr, "Error serializing 'quantity' field")
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(len(m.GetValue())))
		_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Array Field (value)
		// Byte Array field (value)
		if err := writeBuffer.WriteByteArray("value", m.GetValue()); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteMultipleCoilsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUWriteMultipleCoilsRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ModbusPDUWriteMultipleCoilsRequest) isModbusPDUWriteMultipleCoilsRequest() bool {
	return true
}

func (m *_ModbusPDUWriteMultipleCoilsRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
