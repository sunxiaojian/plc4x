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

// S7MessageObjectResponse is the corresponding interface of S7MessageObjectResponse
type S7MessageObjectResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7DataAlarmMessage
	// GetReturnCode returns ReturnCode (property field)
	GetReturnCode() DataTransportErrorCode
	// GetTransportSize returns TransportSize (property field)
	GetTransportSize() DataTransportSize
}

// S7MessageObjectResponseExactly can be used when we want exactly this type and not a type which fulfills S7MessageObjectResponse.
// This is useful for switch cases.
type S7MessageObjectResponseExactly interface {
	S7MessageObjectResponse
	isS7MessageObjectResponse() bool
}

// _S7MessageObjectResponse is the data-structure of this message
type _S7MessageObjectResponse struct {
	*_S7DataAlarmMessage
	ReturnCode    DataTransportErrorCode
	TransportSize DataTransportSize
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7MessageObjectResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7MessageObjectResponse) InitializeParent(parent S7DataAlarmMessage) {}

func (m *_S7MessageObjectResponse) GetParent() S7DataAlarmMessage {
	return m._S7DataAlarmMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7MessageObjectResponse) GetReturnCode() DataTransportErrorCode {
	return m.ReturnCode
}

func (m *_S7MessageObjectResponse) GetTransportSize() DataTransportSize {
	return m.TransportSize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7MessageObjectResponse factory function for _S7MessageObjectResponse
func NewS7MessageObjectResponse(returnCode DataTransportErrorCode, transportSize DataTransportSize) *_S7MessageObjectResponse {
	_result := &_S7MessageObjectResponse{
		ReturnCode:          returnCode,
		TransportSize:       transportSize,
		_S7DataAlarmMessage: NewS7DataAlarmMessage(),
	}
	_result._S7DataAlarmMessage._S7DataAlarmMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7MessageObjectResponse(structType any) S7MessageObjectResponse {
	if casted, ok := structType.(S7MessageObjectResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7MessageObjectResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7MessageObjectResponse) GetTypeName() string {
	return "S7MessageObjectResponse"
}

func (m *_S7MessageObjectResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (returnCode)
	lengthInBits += 8

	// Simple field (transportSize)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7MessageObjectResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7MessageObjectResponseParse(ctx context.Context, theBytes []byte, cpuFunctionType uint8) (S7MessageObjectResponse, error) {
	return S7MessageObjectResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cpuFunctionType)
}

func S7MessageObjectResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionType uint8) (S7MessageObjectResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("S7MessageObjectResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7MessageObjectResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (returnCode)
	if pullErr := readBuffer.PullContext("returnCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for returnCode")
	}
	_returnCode, _returnCodeErr := DataTransportErrorCodeParseWithBuffer(ctx, readBuffer)
	if _returnCodeErr != nil {
		return nil, errors.Wrap(_returnCodeErr, "Error parsing 'returnCode' field of S7MessageObjectResponse")
	}
	returnCode := _returnCode
	if closeErr := readBuffer.CloseContext("returnCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for returnCode")
	}

	// Simple Field (transportSize)
	if pullErr := readBuffer.PullContext("transportSize"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transportSize")
	}
	_transportSize, _transportSizeErr := DataTransportSizeParseWithBuffer(ctx, readBuffer)
	if _transportSizeErr != nil {
		return nil, errors.Wrap(_transportSizeErr, "Error parsing 'transportSize' field of S7MessageObjectResponse")
	}
	transportSize := _transportSize
	if closeErr := readBuffer.CloseContext("transportSize"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transportSize")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of S7MessageObjectResponse")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	if closeErr := readBuffer.CloseContext("S7MessageObjectResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7MessageObjectResponse")
	}

	// Create a partially initialized instance
	_child := &_S7MessageObjectResponse{
		_S7DataAlarmMessage: &_S7DataAlarmMessage{},
		ReturnCode:          returnCode,
		TransportSize:       transportSize,
		reservedField0:      reservedField0,
	}
	_child._S7DataAlarmMessage._S7DataAlarmMessageChildRequirements = _child
	return _child, nil
}

func (m *_S7MessageObjectResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7MessageObjectResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7MessageObjectResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7MessageObjectResponse")
		}

		// Simple Field (returnCode)
		if pushErr := writeBuffer.PushContext("returnCode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for returnCode")
		}
		_returnCodeErr := writeBuffer.WriteSerializable(ctx, m.GetReturnCode())
		if popErr := writeBuffer.PopContext("returnCode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for returnCode")
		}
		if _returnCodeErr != nil {
			return errors.Wrap(_returnCodeErr, "Error serializing 'returnCode' field")
		}

		// Simple Field (transportSize)
		if pushErr := writeBuffer.PushContext("transportSize"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for transportSize")
		}
		_transportSizeErr := writeBuffer.WriteSerializable(ctx, m.GetTransportSize())
		if popErr := writeBuffer.PopContext("transportSize"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for transportSize")
		}
		if _transportSizeErr != nil {
			return errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := /*TODO: migrate me*/ writeBuffer.WriteUint8("reserved", 8, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("S7MessageObjectResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7MessageObjectResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7MessageObjectResponse) isS7MessageObjectResponse() bool {
	return true
}

func (m *_S7MessageObjectResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
