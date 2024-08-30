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

// AirConditioningDataSetZoneGroupOff is the corresponding interface of AirConditioningDataSetZoneGroupOff
type AirConditioningDataSetZoneGroupOff interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
}

// AirConditioningDataSetZoneGroupOffExactly can be used when we want exactly this type and not a type which fulfills AirConditioningDataSetZoneGroupOff.
// This is useful for switch cases.
type AirConditioningDataSetZoneGroupOffExactly interface {
	AirConditioningDataSetZoneGroupOff
	isAirConditioningDataSetZoneGroupOff() bool
}

// _AirConditioningDataSetZoneGroupOff is the data-structure of this message
type _AirConditioningDataSetZoneGroupOff struct {
	*_AirConditioningData
	ZoneGroup byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataSetZoneGroupOff) InitializeParent(parent AirConditioningData, commandTypeContainer AirConditioningCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_AirConditioningDataSetZoneGroupOff) GetParent() AirConditioningData {
	return m._AirConditioningData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetZoneGroupOff) GetZoneGroup() byte {
	return m.ZoneGroup
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataSetZoneGroupOff factory function for _AirConditioningDataSetZoneGroupOff
func NewAirConditioningDataSetZoneGroupOff(zoneGroup byte, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataSetZoneGroupOff {
	_result := &_AirConditioningDataSetZoneGroupOff{
		ZoneGroup:            zoneGroup,
		_AirConditioningData: NewAirConditioningData(commandTypeContainer),
	}
	_result._AirConditioningData._AirConditioningDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetZoneGroupOff(structType any) AirConditioningDataSetZoneGroupOff {
	if casted, ok := structType.(AirConditioningDataSetZoneGroupOff); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetZoneGroupOff); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetZoneGroupOff) GetTypeName() string {
	return "AirConditioningDataSetZoneGroupOff"
}

func (m *_AirConditioningDataSetZoneGroupOff) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	return lengthInBits
}

func (m *_AirConditioningDataSetZoneGroupOff) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AirConditioningDataSetZoneGroupOffParse(ctx context.Context, theBytes []byte) (AirConditioningDataSetZoneGroupOff, error) {
	return AirConditioningDataSetZoneGroupOffParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AirConditioningDataSetZoneGroupOffParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AirConditioningDataSetZoneGroupOff, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AirConditioningDataSetZoneGroupOff"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetZoneGroupOff")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneGroup)
	_zoneGroup, _zoneGroupErr := /*TODO: migrate me*/ readBuffer.ReadByte("zoneGroup")
	if _zoneGroupErr != nil {
		return nil, errors.Wrap(_zoneGroupErr, "Error parsing 'zoneGroup' field of AirConditioningDataSetZoneGroupOff")
	}
	zoneGroup := _zoneGroup

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetZoneGroupOff"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetZoneGroupOff")
	}

	// Create a partially initialized instance
	_child := &_AirConditioningDataSetZoneGroupOff{
		_AirConditioningData: &_AirConditioningData{},
		ZoneGroup:            zoneGroup,
	}
	_child._AirConditioningData._AirConditioningDataChildRequirements = _child
	return _child, nil
}

func (m *_AirConditioningDataSetZoneGroupOff) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataSetZoneGroupOff) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetZoneGroupOff"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetZoneGroupOff")
		}

		// Simple Field (zoneGroup)
		zoneGroup := byte(m.GetZoneGroup())
		_zoneGroupErr := /*TODO: migrate me*/ writeBuffer.WriteByte("zoneGroup", (zoneGroup))
		if _zoneGroupErr != nil {
			return errors.Wrap(_zoneGroupErr, "Error serializing 'zoneGroup' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetZoneGroupOff"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetZoneGroupOff")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetZoneGroupOff) isAirConditioningDataSetZoneGroupOff() bool {
	return true
}

func (m *_AirConditioningDataSetZoneGroupOff) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
