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

// CALDataStatusExtended is the corresponding interface of CALDataStatusExtended
type CALDataStatusExtended interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CALData
	// GetCoding returns Coding (property field)
	GetCoding() StatusCoding
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetBlockStart returns BlockStart (property field)
	GetBlockStart() uint8
	// GetStatusBytes returns StatusBytes (property field)
	GetStatusBytes() []StatusByte
	// GetLevelInformation returns LevelInformation (property field)
	GetLevelInformation() []LevelInformation
	// GetNumberOfStatusBytes returns NumberOfStatusBytes (virtual field)
	GetNumberOfStatusBytes() uint8
	// GetNumberOfLevelInformation returns NumberOfLevelInformation (virtual field)
	GetNumberOfLevelInformation() uint8
}

// CALDataStatusExtendedExactly can be used when we want exactly this type and not a type which fulfills CALDataStatusExtended.
// This is useful for switch cases.
type CALDataStatusExtendedExactly interface {
	CALDataStatusExtended
	isCALDataStatusExtended() bool
}

// _CALDataStatusExtended is the data-structure of this message
type _CALDataStatusExtended struct {
	*_CALData
	Coding           StatusCoding
	Application      ApplicationIdContainer
	BlockStart       uint8
	StatusBytes      []StatusByte
	LevelInformation []LevelInformation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataStatusExtended) InitializeParent(parent CALData, commandTypeContainer CALCommandTypeContainer, additionalData CALData) {
	m.CommandTypeContainer = commandTypeContainer
	m.AdditionalData = additionalData
}

func (m *_CALDataStatusExtended) GetParent() CALData {
	return m._CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataStatusExtended) GetCoding() StatusCoding {
	return m.Coding
}

func (m *_CALDataStatusExtended) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_CALDataStatusExtended) GetBlockStart() uint8 {
	return m.BlockStart
}

func (m *_CALDataStatusExtended) GetStatusBytes() []StatusByte {
	return m.StatusBytes
}

func (m *_CALDataStatusExtended) GetLevelInformation() []LevelInformation {
	return m.LevelInformation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_CALDataStatusExtended) GetNumberOfStatusBytes() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(utils.InlineIf((bool(bool((m.GetCoding()) == (StatusCoding_BINARY_BY_THIS_SERIAL_INTERFACE))) || bool(bool((m.GetCoding()) == (StatusCoding_BINARY_BY_ELSEWHERE)))), func() any { return uint8((uint8(m.GetCommandTypeContainer().NumBytes()) - uint8(uint8(3)))) }, func() any { return uint8((uint8(0))) }).(uint8))
}

func (m *_CALDataStatusExtended) GetNumberOfLevelInformation() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(utils.InlineIf((bool(bool((m.GetCoding()) == (StatusCoding_LEVEL_BY_THIS_SERIAL_INTERFACE))) || bool(bool((m.GetCoding()) == (StatusCoding_LEVEL_BY_ELSEWHERE)))), func() any {
		return uint8((uint8((uint8(m.GetCommandTypeContainer().NumBytes()) - uint8(uint8(3)))) / uint8(uint8(2))))
	}, func() any { return uint8((uint8(0))) }).(uint8))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataStatusExtended factory function for _CALDataStatusExtended
func NewCALDataStatusExtended(coding StatusCoding, application ApplicationIdContainer, blockStart uint8, statusBytes []StatusByte, levelInformation []LevelInformation, commandTypeContainer CALCommandTypeContainer, additionalData CALData, requestContext RequestContext) *_CALDataStatusExtended {
	_result := &_CALDataStatusExtended{
		Coding:           coding,
		Application:      application,
		BlockStart:       blockStart,
		StatusBytes:      statusBytes,
		LevelInformation: levelInformation,
		_CALData:         NewCALData(commandTypeContainer, additionalData, requestContext),
	}
	_result._CALData._CALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataStatusExtended(structType any) CALDataStatusExtended {
	if casted, ok := structType.(CALDataStatusExtended); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataStatusExtended); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataStatusExtended) GetTypeName() string {
	return "CALDataStatusExtended"
}

func (m *_CALDataStatusExtended) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (coding)
	lengthInBits += 8

	// Simple field (application)
	lengthInBits += 8

	// Simple field (blockStart)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Array field
	if len(m.StatusBytes) > 0 {
		for _curItem, element := range m.StatusBytes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.StatusBytes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Array field
	if len(m.LevelInformation) > 0 {
		for _curItem, element := range m.LevelInformation {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LevelInformation), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_CALDataStatusExtended) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CALDataStatusExtendedParse(ctx context.Context, theBytes []byte, commandTypeContainer CALCommandTypeContainer, requestContext RequestContext) (CALDataStatusExtended, error) {
	return CALDataStatusExtendedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), commandTypeContainer, requestContext)
}

func CALDataStatusExtendedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, commandTypeContainer CALCommandTypeContainer, requestContext RequestContext) (CALDataStatusExtended, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CALDataStatusExtended"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataStatusExtended")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (coding)
	if pullErr := readBuffer.PullContext("coding"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for coding")
	}
	_coding, _codingErr := StatusCodingParseWithBuffer(ctx, readBuffer)
	if _codingErr != nil {
		return nil, errors.Wrap(_codingErr, "Error parsing 'coding' field of CALDataStatusExtended")
	}
	coding := _coding
	if closeErr := readBuffer.CloseContext("coding"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for coding")
	}

	// Simple Field (application)
	if pullErr := readBuffer.PullContext("application"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for application")
	}
	_application, _applicationErr := ApplicationIdContainerParseWithBuffer(ctx, readBuffer)
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field of CALDataStatusExtended")
	}
	application := _application
	if closeErr := readBuffer.CloseContext("application"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for application")
	}

	// Simple Field (blockStart)
	_blockStart, _blockStartErr := /*TODO: migrate me*/ readBuffer.ReadUint8("blockStart", 8)
	if _blockStartErr != nil {
		return nil, errors.Wrap(_blockStartErr, "Error parsing 'blockStart' field of CALDataStatusExtended")
	}
	blockStart := _blockStart

	// Virtual field
	_numberOfStatusBytes := utils.InlineIf((bool(bool((coding) == (StatusCoding_BINARY_BY_THIS_SERIAL_INTERFACE))) || bool(bool((coding) == (StatusCoding_BINARY_BY_ELSEWHERE)))), func() any { return uint8((uint8(commandTypeContainer.NumBytes()) - uint8(uint8(3)))) }, func() any { return uint8((uint8(0))) }).(uint8)
	numberOfStatusBytes := uint8(_numberOfStatusBytes)
	_ = numberOfStatusBytes

	// Virtual field
	_numberOfLevelInformation := utils.InlineIf((bool(bool((coding) == (StatusCoding_LEVEL_BY_THIS_SERIAL_INTERFACE))) || bool(bool((coding) == (StatusCoding_LEVEL_BY_ELSEWHERE)))), func() any {
		return uint8((uint8((uint8(commandTypeContainer.NumBytes()) - uint8(uint8(3)))) / uint8(uint8(2))))
	}, func() any { return uint8((uint8(0))) }).(uint8)
	numberOfLevelInformation := uint8(_numberOfLevelInformation)
	_ = numberOfLevelInformation

	// Array field (statusBytes)
	if pullErr := readBuffer.PullContext("statusBytes", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusBytes")
	}
	// Count array
	statusBytes := make([]StatusByte, max(numberOfStatusBytes, 0))
	// This happens when the size is set conditional to 0
	if len(statusBytes) == 0 {
		statusBytes = nil
	}
	{
		_numItems := uint16(max(numberOfStatusBytes, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := StatusByteParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'statusBytes' field of CALDataStatusExtended")
			}
			statusBytes[_curItem] = _item.(StatusByte)
		}
	}
	if closeErr := readBuffer.CloseContext("statusBytes", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusBytes")
	}

	// Array field (levelInformation)
	if pullErr := readBuffer.PullContext("levelInformation", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for levelInformation")
	}
	// Count array
	levelInformation := make([]LevelInformation, max(numberOfLevelInformation, 0))
	// This happens when the size is set conditional to 0
	if len(levelInformation) == 0 {
		levelInformation = nil
	}
	{
		_numItems := uint16(max(numberOfLevelInformation, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := LevelInformationParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'levelInformation' field of CALDataStatusExtended")
			}
			levelInformation[_curItem] = _item.(LevelInformation)
		}
	}
	if closeErr := readBuffer.CloseContext("levelInformation", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for levelInformation")
	}

	if closeErr := readBuffer.CloseContext("CALDataStatusExtended"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataStatusExtended")
	}

	// Create a partially initialized instance
	_child := &_CALDataStatusExtended{
		_CALData: &_CALData{
			RequestContext: requestContext,
		},
		Coding:           coding,
		Application:      application,
		BlockStart:       blockStart,
		StatusBytes:      statusBytes,
		LevelInformation: levelInformation,
	}
	_child._CALData._CALDataChildRequirements = _child
	return _child, nil
}

func (m *_CALDataStatusExtended) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataStatusExtended) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataStatusExtended"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataStatusExtended")
		}

		// Simple Field (coding)
		if pushErr := writeBuffer.PushContext("coding"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for coding")
		}
		_codingErr := writeBuffer.WriteSerializable(ctx, m.GetCoding())
		if popErr := writeBuffer.PopContext("coding"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for coding")
		}
		if _codingErr != nil {
			return errors.Wrap(_codingErr, "Error serializing 'coding' field")
		}

		// Simple Field (application)
		if pushErr := writeBuffer.PushContext("application"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for application")
		}
		_applicationErr := writeBuffer.WriteSerializable(ctx, m.GetApplication())
		if popErr := writeBuffer.PopContext("application"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for application")
		}
		if _applicationErr != nil {
			return errors.Wrap(_applicationErr, "Error serializing 'application' field")
		}

		// Simple Field (blockStart)
		blockStart := uint8(m.GetBlockStart())
		_blockStartErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("blockStart", 8, uint8((blockStart)))
		if _blockStartErr != nil {
			return errors.Wrap(_blockStartErr, "Error serializing 'blockStart' field")
		}
		// Virtual field
		numberOfStatusBytes := m.GetNumberOfStatusBytes()
		_ = numberOfStatusBytes
		if _numberOfStatusBytesErr := writeBuffer.WriteVirtual(ctx, "numberOfStatusBytes", m.GetNumberOfStatusBytes()); _numberOfStatusBytesErr != nil {
			return errors.Wrap(_numberOfStatusBytesErr, "Error serializing 'numberOfStatusBytes' field")
		}
		// Virtual field
		numberOfLevelInformation := m.GetNumberOfLevelInformation()
		_ = numberOfLevelInformation
		if _numberOfLevelInformationErr := writeBuffer.WriteVirtual(ctx, "numberOfLevelInformation", m.GetNumberOfLevelInformation()); _numberOfLevelInformationErr != nil {
			return errors.Wrap(_numberOfLevelInformationErr, "Error serializing 'numberOfLevelInformation' field")
		}

		// Array Field (statusBytes)
		if pushErr := writeBuffer.PushContext("statusBytes", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusBytes")
		}
		for _curItem, _element := range m.GetStatusBytes() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetStatusBytes()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'statusBytes' field")
			}
		}
		if popErr := writeBuffer.PopContext("statusBytes", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusBytes")
		}

		// Array Field (levelInformation)
		if pushErr := writeBuffer.PushContext("levelInformation", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for levelInformation")
		}
		for _curItem, _element := range m.GetLevelInformation() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetLevelInformation()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'levelInformation' field")
			}
		}
		if popErr := writeBuffer.PopContext("levelInformation", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for levelInformation")
		}

		if popErr := writeBuffer.PopContext("CALDataStatusExtended"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataStatusExtended")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALDataStatusExtended) isCALDataStatusExtended() bool {
	return true
}

func (m *_CALDataStatusExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
