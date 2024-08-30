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

// AccessControlData is the corresponding interface of AccessControlData
type AccessControlData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() AccessControlCommandTypeContainer
	// GetNetworkId returns NetworkId (property field)
	GetNetworkId() byte
	// GetAccessPointId returns AccessPointId (property field)
	GetAccessPointId() byte
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() AccessControlCommandType
}

// AccessControlDataExactly can be used when we want exactly this type and not a type which fulfills AccessControlData.
// This is useful for switch cases.
type AccessControlDataExactly interface {
	AccessControlData
	isAccessControlData() bool
}

// _AccessControlData is the data-structure of this message
type _AccessControlData struct {
	_AccessControlDataChildRequirements
	CommandTypeContainer AccessControlCommandTypeContainer
	NetworkId            byte
	AccessPointId        byte
}

type _AccessControlDataChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetCommandType() AccessControlCommandType
}

type AccessControlDataParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child AccessControlData, serializeChildFunction func() error) error
	GetTypeName() string
}

type AccessControlDataChild interface {
	utils.Serializable
	InitializeParent(parent AccessControlData, commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte)
	GetParent() *AccessControlData

	GetTypeName() string
	AccessControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AccessControlData) GetCommandTypeContainer() AccessControlCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_AccessControlData) GetNetworkId() byte {
	return m.NetworkId
}

func (m *_AccessControlData) GetAccessPointId() byte {
	return m.AccessPointId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_AccessControlData) GetCommandType() AccessControlCommandType {
	ctx := context.Background()
	_ = ctx
	return CastAccessControlCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAccessControlData factory function for _AccessControlData
func NewAccessControlData(commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) *_AccessControlData {
	return &_AccessControlData{CommandTypeContainer: commandTypeContainer, NetworkId: networkId, AccessPointId: accessPointId}
}

// Deprecated: use the interface for direct cast
func CastAccessControlData(structType any) AccessControlData {
	if casted, ok := structType.(AccessControlData); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlData); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlData) GetTypeName() string {
	return "AccessControlData"
}

func (m *_AccessControlData) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (networkId)
	lengthInBits += 8

	// Simple field (accessPointId)
	lengthInBits += 8

	return lengthInBits
}

func (m *_AccessControlData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AccessControlDataParse(ctx context.Context, theBytes []byte) (AccessControlData, error) {
	return AccessControlDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AccessControlDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AccessControlData, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AccessControlData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsAccessControlCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "no command type could be found"})
	}

	// Simple Field (commandTypeContainer)
	if pullErr := readBuffer.PullContext("commandTypeContainer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for commandTypeContainer")
	}
	_commandTypeContainer, _commandTypeContainerErr := AccessControlCommandTypeContainerParseWithBuffer(ctx, readBuffer)
	if _commandTypeContainerErr != nil {
		return nil, errors.Wrap(_commandTypeContainerErr, "Error parsing 'commandTypeContainer' field of AccessControlData")
	}
	commandTypeContainer := _commandTypeContainer
	if closeErr := readBuffer.CloseContext("commandTypeContainer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for commandTypeContainer")
	}

	// Virtual field
	_commandType := commandTypeContainer.CommandType()
	commandType := AccessControlCommandType(_commandType)
	_ = commandType

	// Simple Field (networkId)
	_networkId, _networkIdErr := /*TODO: migrate me*/ readBuffer.ReadByte("networkId")
	if _networkIdErr != nil {
		return nil, errors.Wrap(_networkIdErr, "Error parsing 'networkId' field of AccessControlData")
	}
	networkId := _networkId

	// Simple Field (accessPointId)
	_accessPointId, _accessPointIdErr := /*TODO: migrate me*/ readBuffer.ReadByte("accessPointId")
	if _accessPointIdErr != nil {
		return nil, errors.Wrap(_accessPointIdErr, "Error parsing 'accessPointId' field of AccessControlData")
	}
	accessPointId := _accessPointId

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type AccessControlDataChildSerializeRequirement interface {
		AccessControlData
		InitializeParent(AccessControlData, AccessControlCommandTypeContainer, byte, byte)
		GetParent() AccessControlData
	}
	var _childTemp any
	var _child AccessControlDataChildSerializeRequirement
	var typeSwitchError error
	switch {
	case commandType == AccessControlCommandType_VALID_ACCESS: // AccessControlDataValidAccessRequest
		_childTemp, typeSwitchError = AccessControlDataValidAccessRequestParseWithBuffer(ctx, readBuffer, commandTypeContainer)
	case commandType == AccessControlCommandType_INVALID_ACCESS: // AccessControlDataInvalidAccessRequest
		_childTemp, typeSwitchError = AccessControlDataInvalidAccessRequestParseWithBuffer(ctx, readBuffer, commandTypeContainer)
	case commandType == AccessControlCommandType_ACCESS_POINT_LEFT_OPEN: // AccessControlDataAccessPointLeftOpen
		_childTemp, typeSwitchError = AccessControlDataAccessPointLeftOpenParseWithBuffer(ctx, readBuffer)
	case commandType == AccessControlCommandType_ACCESS_POINT_FORCED_OPEN: // AccessControlDataAccessPointForcedOpen
		_childTemp, typeSwitchError = AccessControlDataAccessPointForcedOpenParseWithBuffer(ctx, readBuffer)
	case commandType == AccessControlCommandType_ACCESS_POINT_CLOSED: // AccessControlDataAccessPointClosed
		_childTemp, typeSwitchError = AccessControlDataAccessPointClosedParseWithBuffer(ctx, readBuffer)
	case commandType == AccessControlCommandType_REQUEST_TO_EXIT: // AccessControlDataRequestToExit
		_childTemp, typeSwitchError = AccessControlDataRequestToExitParseWithBuffer(ctx, readBuffer)
	case commandType == AccessControlCommandType_CLOSE_ACCESS_POINT: // AccessControlDataCloseAccessPoint
		_childTemp, typeSwitchError = AccessControlDataCloseAccessPointParseWithBuffer(ctx, readBuffer)
	case commandType == AccessControlCommandType_LOCK_ACCESS_POINT: // AccessControlDataLockAccessPoint
		_childTemp, typeSwitchError = AccessControlDataLockAccessPointParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of AccessControlData")
	}
	_child = _childTemp.(AccessControlDataChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("AccessControlData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlData")
	}

	// Finish initializing
	_child.InitializeParent(_child, commandTypeContainer, networkId, accessPointId)
	return _child, nil
}

func (pm *_AccessControlData) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child AccessControlData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AccessControlData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AccessControlData")
	}

	// Simple Field (commandTypeContainer)
	if pushErr := writeBuffer.PushContext("commandTypeContainer"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for commandTypeContainer")
	}
	_commandTypeContainerErr := writeBuffer.WriteSerializable(ctx, m.GetCommandTypeContainer())
	if popErr := writeBuffer.PopContext("commandTypeContainer"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for commandTypeContainer")
	}
	if _commandTypeContainerErr != nil {
		return errors.Wrap(_commandTypeContainerErr, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Simple Field (networkId)
	networkId := byte(m.GetNetworkId())
	_networkIdErr := /*TODO: migrate me*/ writeBuffer.WriteByte("networkId", (networkId))
	if _networkIdErr != nil {
		return errors.Wrap(_networkIdErr, "Error serializing 'networkId' field")
	}

	// Simple Field (accessPointId)
	accessPointId := byte(m.GetAccessPointId())
	_accessPointIdErr := /*TODO: migrate me*/ writeBuffer.WriteByte("accessPointId", (accessPointId))
	if _accessPointIdErr != nil {
		return errors.Wrap(_accessPointIdErr, "Error serializing 'accessPointId' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("AccessControlData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AccessControlData")
	}
	return nil
}

func (m *_AccessControlData) isAccessControlData() bool {
	return true
}

func (m *_AccessControlData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
