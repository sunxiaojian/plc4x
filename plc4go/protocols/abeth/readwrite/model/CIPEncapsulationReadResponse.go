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

// CIPEncapsulationReadResponse is the corresponding interface of CIPEncapsulationReadResponse
type CIPEncapsulationReadResponse interface {
	utils.LengthAware
	utils.Serializable
	CIPEncapsulationPacket
	// GetResponse returns Response (property field)
	GetResponse() DF1ResponseMessage
}

// CIPEncapsulationReadResponseExactly can be used when we want exactly this type and not a type which fulfills CIPEncapsulationReadResponse.
// This is useful for switch cases.
type CIPEncapsulationReadResponseExactly interface {
	CIPEncapsulationReadResponse
	isCIPEncapsulationReadResponse() bool
}

// _CIPEncapsulationReadResponse is the data-structure of this message
type _CIPEncapsulationReadResponse struct {
	*_CIPEncapsulationPacket
	Response DF1ResponseMessage

	// Arguments.
	PacketLen uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CIPEncapsulationReadResponse) GetCommandType() uint16 {
	return 0x0207
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CIPEncapsulationReadResponse) InitializeParent(parent CIPEncapsulationPacket, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) {
	m.SessionHandle = sessionHandle
	m.Status = status
	m.SenderContext = senderContext
	m.Options = options
}

func (m *_CIPEncapsulationReadResponse) GetParent() CIPEncapsulationPacket {
	return m._CIPEncapsulationPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CIPEncapsulationReadResponse) GetResponse() DF1ResponseMessage {
	return m.Response
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCIPEncapsulationReadResponse factory function for _CIPEncapsulationReadResponse
func NewCIPEncapsulationReadResponse(response DF1ResponseMessage, sessionHandle uint32, status uint32, senderContext []uint8, options uint32, packetLen uint16) *_CIPEncapsulationReadResponse {
	_result := &_CIPEncapsulationReadResponse{
		Response:                response,
		_CIPEncapsulationPacket: NewCIPEncapsulationPacket(sessionHandle, status, senderContext, options),
	}
	_result._CIPEncapsulationPacket._CIPEncapsulationPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCIPEncapsulationReadResponse(structType interface{}) CIPEncapsulationReadResponse {
	if casted, ok := structType.(CIPEncapsulationReadResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CIPEncapsulationReadResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CIPEncapsulationReadResponse) GetTypeName() string {
	return "CIPEncapsulationReadResponse"
}

func (m *_CIPEncapsulationReadResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CIPEncapsulationReadResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (response)
	lengthInBits += m.Response.GetLengthInBits()

	return lengthInBits
}

func (m *_CIPEncapsulationReadResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CIPEncapsulationReadResponseParse(readBuffer utils.ReadBuffer, packetLen uint16) (CIPEncapsulationReadResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CIPEncapsulationReadResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CIPEncapsulationReadResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (response)
	if pullErr := readBuffer.PullContext("response"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for response")
	}
	_response, _responseErr := DF1ResponseMessageParse(readBuffer, uint16(packetLen))
	if _responseErr != nil {
		return nil, errors.Wrap(_responseErr, "Error parsing 'response' field of CIPEncapsulationReadResponse")
	}
	response := _response.(DF1ResponseMessage)
	if closeErr := readBuffer.CloseContext("response"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for response")
	}

	if closeErr := readBuffer.CloseContext("CIPEncapsulationReadResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CIPEncapsulationReadResponse")
	}

	// Create a partially initialized instance
	_child := &_CIPEncapsulationReadResponse{
		_CIPEncapsulationPacket: &_CIPEncapsulationPacket{},
		Response:                response,
	}
	_child._CIPEncapsulationPacket._CIPEncapsulationPacketChildRequirements = _child
	return _child, nil
}

func (m *_CIPEncapsulationReadResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CIPEncapsulationReadResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CIPEncapsulationReadResponse")
		}

		// Simple Field (response)
		if pushErr := writeBuffer.PushContext("response"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for response")
		}
		_responseErr := writeBuffer.WriteSerializable(m.GetResponse())
		if popErr := writeBuffer.PopContext("response"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for response")
		}
		if _responseErr != nil {
			return errors.Wrap(_responseErr, "Error serializing 'response' field")
		}

		if popErr := writeBuffer.PopContext("CIPEncapsulationReadResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CIPEncapsulationReadResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_CIPEncapsulationReadResponse) GetPacketLen() uint16 {
	return m.PacketLen
}

//
////

func (m *_CIPEncapsulationReadResponse) isCIPEncapsulationReadResponse() bool {
	return true
}

func (m *_CIPEncapsulationReadResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
