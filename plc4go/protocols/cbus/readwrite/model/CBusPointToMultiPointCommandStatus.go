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

// CBusPointToMultiPointCommandStatus is the corresponding interface of CBusPointToMultiPointCommandStatus
type CBusPointToMultiPointCommandStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CBusPointToMultiPointCommand
	// GetStatusRequest returns StatusRequest (property field)
	GetStatusRequest() StatusRequest
}

// CBusPointToMultiPointCommandStatusExactly can be used when we want exactly this type and not a type which fulfills CBusPointToMultiPointCommandStatus.
// This is useful for switch cases.
type CBusPointToMultiPointCommandStatusExactly interface {
	CBusPointToMultiPointCommandStatus
	isCBusPointToMultiPointCommandStatus() bool
}

// _CBusPointToMultiPointCommandStatus is the data-structure of this message
type _CBusPointToMultiPointCommandStatus struct {
	*_CBusPointToMultiPointCommand
	StatusRequest StatusRequest
	// Reserved Fields
	reservedField0 *byte
	reservedField1 *byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CBusPointToMultiPointCommandStatus) InitializeParent(parent CBusPointToMultiPointCommand, peekedApplication byte) {
	m.PeekedApplication = peekedApplication
}

func (m *_CBusPointToMultiPointCommandStatus) GetParent() CBusPointToMultiPointCommand {
	return m._CBusPointToMultiPointCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToMultiPointCommandStatus) GetStatusRequest() StatusRequest {
	return m.StatusRequest
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusPointToMultiPointCommandStatus factory function for _CBusPointToMultiPointCommandStatus
func NewCBusPointToMultiPointCommandStatus(statusRequest StatusRequest, peekedApplication byte, cBusOptions CBusOptions) *_CBusPointToMultiPointCommandStatus {
	_result := &_CBusPointToMultiPointCommandStatus{
		StatusRequest:                 statusRequest,
		_CBusPointToMultiPointCommand: NewCBusPointToMultiPointCommand(peekedApplication, cBusOptions),
	}
	_result._CBusPointToMultiPointCommand._CBusPointToMultiPointCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCBusPointToMultiPointCommandStatus(structType any) CBusPointToMultiPointCommandStatus {
	if casted, ok := structType.(CBusPointToMultiPointCommandStatus); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToMultiPointCommandStatus); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToMultiPointCommandStatus) GetTypeName() string {
	return "CBusPointToMultiPointCommandStatus"
}

func (m *_CBusPointToMultiPointCommandStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (statusRequest)
	lengthInBits += m.StatusRequest.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CBusPointToMultiPointCommandStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CBusPointToMultiPointCommandStatusParse(ctx context.Context, theBytes []byte, cBusOptions CBusOptions) (CBusPointToMultiPointCommandStatus, error) {
	return CBusPointToMultiPointCommandStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func CBusPointToMultiPointCommandStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (CBusPointToMultiPointCommandStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CBusPointToMultiPointCommandStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToMultiPointCommandStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *byte
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CBusPointToMultiPointCommandStatus")
		}
		if reserved != byte(0xFF) {
			log.Info().Fields(map[string]any{
				"expected value": byte(0xFF),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	var reservedField1 *byte
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CBusPointToMultiPointCommandStatus")
		}
		if reserved != byte(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": byte(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	// Simple Field (statusRequest)
	if pullErr := readBuffer.PullContext("statusRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusRequest")
	}
	_statusRequest, _statusRequestErr := StatusRequestParseWithBuffer(ctx, readBuffer)
	if _statusRequestErr != nil {
		return nil, errors.Wrap(_statusRequestErr, "Error parsing 'statusRequest' field of CBusPointToMultiPointCommandStatus")
	}
	statusRequest := _statusRequest.(StatusRequest)
	if closeErr := readBuffer.CloseContext("statusRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusRequest")
	}

	if closeErr := readBuffer.CloseContext("CBusPointToMultiPointCommandStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToMultiPointCommandStatus")
	}

	// Create a partially initialized instance
	_child := &_CBusPointToMultiPointCommandStatus{
		_CBusPointToMultiPointCommand: &_CBusPointToMultiPointCommand{
			CBusOptions: cBusOptions,
		},
		StatusRequest:  statusRequest,
		reservedField0: reservedField0,
		reservedField1: reservedField1,
	}
	_child._CBusPointToMultiPointCommand._CBusPointToMultiPointCommandChildRequirements = _child
	return _child, nil
}

func (m *_CBusPointToMultiPointCommandStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusPointToMultiPointCommandStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusPointToMultiPointCommandStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusPointToMultiPointCommandStatus")
		}

		// Reserved Field (reserved)
		{
			var reserved byte = byte(0xFF)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": byte(0xFF),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := /*TODO: migrate me*/ writeBuffer.WriteByte("reserved", reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			var reserved byte = byte(0x00)
			if m.reservedField1 != nil {
				log.Info().Fields(map[string]any{
					"expected value": byte(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField1
			}
			_err := /*TODO: migrate me*/ writeBuffer.WriteByte("reserved", reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (statusRequest)
		if pushErr := writeBuffer.PushContext("statusRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusRequest")
		}
		_statusRequestErr := writeBuffer.WriteSerializable(ctx, m.GetStatusRequest())
		if popErr := writeBuffer.PopContext("statusRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusRequest")
		}
		if _statusRequestErr != nil {
			return errors.Wrap(_statusRequestErr, "Error serializing 'statusRequest' field")
		}

		if popErr := writeBuffer.PopContext("CBusPointToMultiPointCommandStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusPointToMultiPointCommandStatus")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CBusPointToMultiPointCommandStatus) isCBusPointToMultiPointCommandStatus() bool {
	return true
}

func (m *_CBusPointToMultiPointCommandStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
