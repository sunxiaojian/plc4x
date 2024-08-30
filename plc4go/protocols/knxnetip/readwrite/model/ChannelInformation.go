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

// ChannelInformation is the corresponding interface of ChannelInformation
type ChannelInformation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNumChannels returns NumChannels (property field)
	GetNumChannels() uint8
	// GetChannelCode returns ChannelCode (property field)
	GetChannelCode() uint16
}

// ChannelInformationExactly can be used when we want exactly this type and not a type which fulfills ChannelInformation.
// This is useful for switch cases.
type ChannelInformationExactly interface {
	ChannelInformation
	isChannelInformation() bool
}

// _ChannelInformation is the data-structure of this message
type _ChannelInformation struct {
	NumChannels uint8
	ChannelCode uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ChannelInformation) GetNumChannels() uint8 {
	return m.NumChannels
}

func (m *_ChannelInformation) GetChannelCode() uint16 {
	return m.ChannelCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewChannelInformation factory function for _ChannelInformation
func NewChannelInformation(numChannels uint8, channelCode uint16) *_ChannelInformation {
	return &_ChannelInformation{NumChannels: numChannels, ChannelCode: channelCode}
}

// Deprecated: use the interface for direct cast
func CastChannelInformation(structType any) ChannelInformation {
	if casted, ok := structType.(ChannelInformation); ok {
		return casted
	}
	if casted, ok := structType.(*ChannelInformation); ok {
		return *casted
	}
	return nil
}

func (m *_ChannelInformation) GetTypeName() string {
	return "ChannelInformation"
}

func (m *_ChannelInformation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (numChannels)
	lengthInBits += 3

	// Simple field (channelCode)
	lengthInBits += 13

	return lengthInBits
}

func (m *_ChannelInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ChannelInformationParse(ctx context.Context, theBytes []byte) (ChannelInformation, error) {
	return ChannelInformationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ChannelInformationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ChannelInformation, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ChannelInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ChannelInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (numChannels)
	_numChannels, _numChannelsErr := /*TODO: migrate me*/ readBuffer.ReadUint8("numChannels", 3)
	if _numChannelsErr != nil {
		return nil, errors.Wrap(_numChannelsErr, "Error parsing 'numChannels' field of ChannelInformation")
	}
	numChannels := _numChannels

	// Simple Field (channelCode)
	_channelCode, _channelCodeErr := /*TODO: migrate me*/ readBuffer.ReadUint16("channelCode", 13)
	if _channelCodeErr != nil {
		return nil, errors.Wrap(_channelCodeErr, "Error parsing 'channelCode' field of ChannelInformation")
	}
	channelCode := _channelCode

	if closeErr := readBuffer.CloseContext("ChannelInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ChannelInformation")
	}

	// Create the instance
	return &_ChannelInformation{
		NumChannels: numChannels,
		ChannelCode: channelCode,
	}, nil
}

func (m *_ChannelInformation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ChannelInformation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ChannelInformation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ChannelInformation")
	}

	// Simple Field (numChannels)
	numChannels := uint8(m.GetNumChannels())
	_numChannelsErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("numChannels", 3, uint8((numChannels)))
	if _numChannelsErr != nil {
		return errors.Wrap(_numChannelsErr, "Error serializing 'numChannels' field")
	}

	// Simple Field (channelCode)
	channelCode := uint16(m.GetChannelCode())
	_channelCodeErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("channelCode", 13, uint16((channelCode)))
	if _channelCodeErr != nil {
		return errors.Wrap(_channelCodeErr, "Error serializing 'channelCode' field")
	}

	if popErr := writeBuffer.PopContext("ChannelInformation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ChannelInformation")
	}
	return nil
}

func (m *_ChannelInformation) isChannelInformation() bool {
	return true
}

func (m *_ChannelInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
