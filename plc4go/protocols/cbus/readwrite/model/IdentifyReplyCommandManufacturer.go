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

// IdentifyReplyCommandManufacturer is the corresponding interface of IdentifyReplyCommandManufacturer
type IdentifyReplyCommandManufacturer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetManufacturerName returns ManufacturerName (property field)
	GetManufacturerName() string
}

// IdentifyReplyCommandManufacturerExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandManufacturer.
// This is useful for switch cases.
type IdentifyReplyCommandManufacturerExactly interface {
	IdentifyReplyCommandManufacturer
	isIdentifyReplyCommandManufacturer() bool
}

// _IdentifyReplyCommandManufacturer is the data-structure of this message
type _IdentifyReplyCommandManufacturer struct {
	*_IdentifyReplyCommand
	ManufacturerName string
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandManufacturer) GetAttribute() Attribute {
	return Attribute_Manufacturer
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandManufacturer) InitializeParent(parent IdentifyReplyCommand) {}

func (m *_IdentifyReplyCommandManufacturer) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandManufacturer) GetManufacturerName() string {
	return m.ManufacturerName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandManufacturer factory function for _IdentifyReplyCommandManufacturer
func NewIdentifyReplyCommandManufacturer(manufacturerName string, numBytes uint8) *_IdentifyReplyCommandManufacturer {
	_result := &_IdentifyReplyCommandManufacturer{
		ManufacturerName:      manufacturerName,
		_IdentifyReplyCommand: NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandManufacturer(structType any) IdentifyReplyCommandManufacturer {
	if casted, ok := structType.(IdentifyReplyCommandManufacturer); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandManufacturer); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandManufacturer) GetTypeName() string {
	return "IdentifyReplyCommandManufacturer"
}

func (m *_IdentifyReplyCommandManufacturer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (manufacturerName)
	lengthInBits += 64

	return lengthInBits
}

func (m *_IdentifyReplyCommandManufacturer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IdentifyReplyCommandManufacturerParse(ctx context.Context, theBytes []byte, attribute Attribute, numBytes uint8) (IdentifyReplyCommandManufacturer, error) {
	return IdentifyReplyCommandManufacturerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), attribute, numBytes)
}

func IdentifyReplyCommandManufacturerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandManufacturer, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandManufacturer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandManufacturer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (manufacturerName)
	_manufacturerName, _manufacturerNameErr := /*TODO: migrate me*/ readBuffer.ReadString("manufacturerName", uint32(64), utils.WithEncoding("UTF-8"))
	if _manufacturerNameErr != nil {
		return nil, errors.Wrap(_manufacturerNameErr, "Error parsing 'manufacturerName' field of IdentifyReplyCommandManufacturer")
	}
	manufacturerName := _manufacturerName

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandManufacturer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandManufacturer")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandManufacturer{
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
		ManufacturerName: manufacturerName,
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandManufacturer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandManufacturer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandManufacturer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandManufacturer")
		}

		// Simple Field (manufacturerName)
		manufacturerName := string(m.GetManufacturerName())
		_manufacturerNameErr := /*TODO: migrate me*/ writeBuffer.WriteString("manufacturerName", uint32(64), (manufacturerName), utils.WithEncoding("UTF-8)"))
		if _manufacturerNameErr != nil {
			return errors.Wrap(_manufacturerNameErr, "Error serializing 'manufacturerName' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandManufacturer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandManufacturer")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandManufacturer) isIdentifyReplyCommandManufacturer() bool {
	return true
}

func (m *_IdentifyReplyCommandManufacturer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
