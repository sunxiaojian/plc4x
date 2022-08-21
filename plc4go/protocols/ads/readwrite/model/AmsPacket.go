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

// AmsPacket is the corresponding interface of AmsPacket
type AmsPacket interface {
	utils.LengthAware
	utils.Serializable
	// GetTargetAmsNetId returns TargetAmsNetId (property field)
	GetTargetAmsNetId() AmsNetId
	// GetTargetAmsPort returns TargetAmsPort (property field)
	GetTargetAmsPort() uint16
	// GetSourceAmsNetId returns SourceAmsNetId (property field)
	GetSourceAmsNetId() AmsNetId
	// GetSourceAmsPort returns SourceAmsPort (property field)
	GetSourceAmsPort() uint16
	// GetCommandId returns CommandId (property field)
	GetCommandId() CommandId
	// GetState returns State (property field)
	GetState() State
	// GetErrorCode returns ErrorCode (property field)
	GetErrorCode() uint32
	// GetInvokeId returns InvokeId (property field)
	GetInvokeId() uint32
	// GetData returns Data (property field)
	GetData() AdsData
}

// AmsPacketExactly can be used when we want exactly this type and not a type which fulfills AmsPacket.
// This is useful for switch cases.
type AmsPacketExactly interface {
	AmsPacket
	isAmsPacket() bool
}

// _AmsPacket is the data-structure of this message
type _AmsPacket struct {
	TargetAmsNetId AmsNetId
	TargetAmsPort  uint16
	SourceAmsNetId AmsNetId
	SourceAmsPort  uint16
	CommandId      CommandId
	State          State
	ErrorCode      uint32
	InvokeId       uint32
	Data           AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AmsPacket) GetTargetAmsNetId() AmsNetId {
	return m.TargetAmsNetId
}

func (m *_AmsPacket) GetTargetAmsPort() uint16 {
	return m.TargetAmsPort
}

func (m *_AmsPacket) GetSourceAmsNetId() AmsNetId {
	return m.SourceAmsNetId
}

func (m *_AmsPacket) GetSourceAmsPort() uint16 {
	return m.SourceAmsPort
}

func (m *_AmsPacket) GetCommandId() CommandId {
	return m.CommandId
}

func (m *_AmsPacket) GetState() State {
	return m.State
}

func (m *_AmsPacket) GetErrorCode() uint32 {
	return m.ErrorCode
}

func (m *_AmsPacket) GetInvokeId() uint32 {
	return m.InvokeId
}

func (m *_AmsPacket) GetData() AdsData {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAmsPacket factory function for _AmsPacket
func NewAmsPacket(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, commandId CommandId, state State, errorCode uint32, invokeId uint32, data AdsData) *_AmsPacket {
	return &_AmsPacket{TargetAmsNetId: targetAmsNetId, TargetAmsPort: targetAmsPort, SourceAmsNetId: sourceAmsNetId, SourceAmsPort: sourceAmsPort, CommandId: commandId, State: state, ErrorCode: errorCode, InvokeId: invokeId, Data: data}
}

// Deprecated: use the interface for direct cast
func CastAmsPacket(structType interface{}) AmsPacket {
	if casted, ok := structType.(AmsPacket); ok {
		return casted
	}
	if casted, ok := structType.(*AmsPacket); ok {
		return *casted
	}
	return nil
}

func (m *_AmsPacket) GetTypeName() string {
	return "AmsPacket"
}

func (m *_AmsPacket) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AmsPacket) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (targetAmsNetId)
	lengthInBits += m.TargetAmsNetId.GetLengthInBits()

	// Simple field (targetAmsPort)
	lengthInBits += 16

	// Simple field (sourceAmsNetId)
	lengthInBits += m.SourceAmsNetId.GetLengthInBits()

	// Simple field (sourceAmsPort)
	lengthInBits += 16

	// Simple field (commandId)
	lengthInBits += 16

	// Simple field (state)
	lengthInBits += m.State.GetLengthInBits()

	// Implicit Field (length)
	lengthInBits += 32

	// Simple field (errorCode)
	lengthInBits += 32

	// Simple field (invokeId)
	lengthInBits += 32

	// Simple field (data)
	lengthInBits += m.Data.GetLengthInBits()

	return lengthInBits
}

func (m *_AmsPacket) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AmsPacketParse(readBuffer utils.ReadBuffer) (AmsPacket, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AmsPacket"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AmsPacket")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (targetAmsNetId)
	if pullErr := readBuffer.PullContext("targetAmsNetId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for targetAmsNetId")
	}
	_targetAmsNetId, _targetAmsNetIdErr := AmsNetIdParse(readBuffer)
	if _targetAmsNetIdErr != nil {
		return nil, errors.Wrap(_targetAmsNetIdErr, "Error parsing 'targetAmsNetId' field of AmsPacket")
	}
	targetAmsNetId := _targetAmsNetId.(AmsNetId)
	if closeErr := readBuffer.CloseContext("targetAmsNetId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for targetAmsNetId")
	}

	// Simple Field (targetAmsPort)
	_targetAmsPort, _targetAmsPortErr := readBuffer.ReadUint16("targetAmsPort", 16)
	if _targetAmsPortErr != nil {
		return nil, errors.Wrap(_targetAmsPortErr, "Error parsing 'targetAmsPort' field of AmsPacket")
	}
	targetAmsPort := _targetAmsPort

	// Simple Field (sourceAmsNetId)
	if pullErr := readBuffer.PullContext("sourceAmsNetId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for sourceAmsNetId")
	}
	_sourceAmsNetId, _sourceAmsNetIdErr := AmsNetIdParse(readBuffer)
	if _sourceAmsNetIdErr != nil {
		return nil, errors.Wrap(_sourceAmsNetIdErr, "Error parsing 'sourceAmsNetId' field of AmsPacket")
	}
	sourceAmsNetId := _sourceAmsNetId.(AmsNetId)
	if closeErr := readBuffer.CloseContext("sourceAmsNetId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for sourceAmsNetId")
	}

	// Simple Field (sourceAmsPort)
	_sourceAmsPort, _sourceAmsPortErr := readBuffer.ReadUint16("sourceAmsPort", 16)
	if _sourceAmsPortErr != nil {
		return nil, errors.Wrap(_sourceAmsPortErr, "Error parsing 'sourceAmsPort' field of AmsPacket")
	}
	sourceAmsPort := _sourceAmsPort

	// Simple Field (commandId)
	if pullErr := readBuffer.PullContext("commandId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for commandId")
	}
	_commandId, _commandIdErr := CommandIdParse(readBuffer)
	if _commandIdErr != nil {
		return nil, errors.Wrap(_commandIdErr, "Error parsing 'commandId' field of AmsPacket")
	}
	commandId := _commandId
	if closeErr := readBuffer.CloseContext("commandId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for commandId")
	}

	// Simple Field (state)
	if pullErr := readBuffer.PullContext("state"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for state")
	}
	_state, _stateErr := StateParse(readBuffer)
	if _stateErr != nil {
		return nil, errors.Wrap(_stateErr, "Error parsing 'state' field of AmsPacket")
	}
	state := _state.(State)
	if closeErr := readBuffer.CloseContext("state"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for state")
	}

	// Implicit Field (length) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := readBuffer.ReadUint32("length", 32)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field of AmsPacket")
	}

	// Simple Field (errorCode)
	_errorCode, _errorCodeErr := readBuffer.ReadUint32("errorCode", 32)
	if _errorCodeErr != nil {
		return nil, errors.Wrap(_errorCodeErr, "Error parsing 'errorCode' field of AmsPacket")
	}
	errorCode := _errorCode

	// Simple Field (invokeId)
	_invokeId, _invokeIdErr := readBuffer.ReadUint32("invokeId", 32)
	if _invokeIdErr != nil {
		return nil, errors.Wrap(_invokeIdErr, "Error parsing 'invokeId' field of AmsPacket")
	}
	invokeId := _invokeId

	// Simple Field (data)
	if pullErr := readBuffer.PullContext("data"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for data")
	}
	_data, _dataErr := AdsDataParse(readBuffer, CommandId(commandId), bool(state.GetResponse()))
	if _dataErr != nil {
		return nil, errors.Wrap(_dataErr, "Error parsing 'data' field of AmsPacket")
	}
	data := _data.(AdsData)
	if closeErr := readBuffer.CloseContext("data"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for data")
	}

	if closeErr := readBuffer.CloseContext("AmsPacket"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AmsPacket")
	}

	// Create the instance
	return &_AmsPacket{
		TargetAmsNetId: targetAmsNetId,
		TargetAmsPort:  targetAmsPort,
		SourceAmsNetId: sourceAmsNetId,
		SourceAmsPort:  sourceAmsPort,
		CommandId:      commandId,
		State:          state,
		ErrorCode:      errorCode,
		InvokeId:       invokeId,
		Data:           data,
	}, nil
}

func (m *_AmsPacket) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AmsPacket"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AmsPacket")
	}

	// Simple Field (targetAmsNetId)
	if pushErr := writeBuffer.PushContext("targetAmsNetId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for targetAmsNetId")
	}
	_targetAmsNetIdErr := writeBuffer.WriteSerializable(m.GetTargetAmsNetId())
	if popErr := writeBuffer.PopContext("targetAmsNetId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for targetAmsNetId")
	}
	if _targetAmsNetIdErr != nil {
		return errors.Wrap(_targetAmsNetIdErr, "Error serializing 'targetAmsNetId' field")
	}

	// Simple Field (targetAmsPort)
	targetAmsPort := uint16(m.GetTargetAmsPort())
	_targetAmsPortErr := writeBuffer.WriteUint16("targetAmsPort", 16, (targetAmsPort))
	if _targetAmsPortErr != nil {
		return errors.Wrap(_targetAmsPortErr, "Error serializing 'targetAmsPort' field")
	}

	// Simple Field (sourceAmsNetId)
	if pushErr := writeBuffer.PushContext("sourceAmsNetId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for sourceAmsNetId")
	}
	_sourceAmsNetIdErr := writeBuffer.WriteSerializable(m.GetSourceAmsNetId())
	if popErr := writeBuffer.PopContext("sourceAmsNetId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for sourceAmsNetId")
	}
	if _sourceAmsNetIdErr != nil {
		return errors.Wrap(_sourceAmsNetIdErr, "Error serializing 'sourceAmsNetId' field")
	}

	// Simple Field (sourceAmsPort)
	sourceAmsPort := uint16(m.GetSourceAmsPort())
	_sourceAmsPortErr := writeBuffer.WriteUint16("sourceAmsPort", 16, (sourceAmsPort))
	if _sourceAmsPortErr != nil {
		return errors.Wrap(_sourceAmsPortErr, "Error serializing 'sourceAmsPort' field")
	}

	// Simple Field (commandId)
	if pushErr := writeBuffer.PushContext("commandId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for commandId")
	}
	_commandIdErr := writeBuffer.WriteSerializable(m.GetCommandId())
	if popErr := writeBuffer.PopContext("commandId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for commandId")
	}
	if _commandIdErr != nil {
		return errors.Wrap(_commandIdErr, "Error serializing 'commandId' field")
	}

	// Simple Field (state)
	if pushErr := writeBuffer.PushContext("state"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for state")
	}
	_stateErr := writeBuffer.WriteSerializable(m.GetState())
	if popErr := writeBuffer.PopContext("state"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for state")
	}
	if _stateErr != nil {
		return errors.Wrap(_stateErr, "Error serializing 'state' field")
	}

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	length := uint32(m.GetData().GetLengthInBytes())
	_lengthErr := writeBuffer.WriteUint32("length", 32, (length))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Simple Field (errorCode)
	errorCode := uint32(m.GetErrorCode())
	_errorCodeErr := writeBuffer.WriteUint32("errorCode", 32, (errorCode))
	if _errorCodeErr != nil {
		return errors.Wrap(_errorCodeErr, "Error serializing 'errorCode' field")
	}

	// Simple Field (invokeId)
	invokeId := uint32(m.GetInvokeId())
	_invokeIdErr := writeBuffer.WriteUint32("invokeId", 32, (invokeId))
	if _invokeIdErr != nil {
		return errors.Wrap(_invokeIdErr, "Error serializing 'invokeId' field")
	}

	// Simple Field (data)
	if pushErr := writeBuffer.PushContext("data"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for data")
	}
	_dataErr := writeBuffer.WriteSerializable(m.GetData())
	if popErr := writeBuffer.PopContext("data"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for data")
	}
	if _dataErr != nil {
		return errors.Wrap(_dataErr, "Error serializing 'data' field")
	}

	if popErr := writeBuffer.PopContext("AmsPacket"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AmsPacket")
	}
	return nil
}

func (m *_AmsPacket) isAmsPacket() bool {
	return true
}

func (m *_AmsPacket) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
