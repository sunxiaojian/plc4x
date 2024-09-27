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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// CycServiceItemAnyType is the corresponding interface of CycServiceItemAnyType
type CycServiceItemAnyType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CycServiceItemType
	// GetTransportSize returns TransportSize (property field)
	GetTransportSize() TransportSize
	// GetLength returns Length (property field)
	GetLength() uint16
	// GetDbNumber returns DbNumber (property field)
	GetDbNumber() uint16
	// GetMemoryArea returns MemoryArea (property field)
	GetMemoryArea() MemoryArea
	// GetAddress returns Address (property field)
	GetAddress() uint32
	// IsCycServiceItemAnyType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCycServiceItemAnyType()
	// CreateBuilder creates a CycServiceItemAnyTypeBuilder
	CreateCycServiceItemAnyTypeBuilder() CycServiceItemAnyTypeBuilder
}

// _CycServiceItemAnyType is the data-structure of this message
type _CycServiceItemAnyType struct {
	CycServiceItemTypeContract
	TransportSize TransportSize
	Length        uint16
	DbNumber      uint16
	MemoryArea    MemoryArea
	Address       uint32
}

var _ CycServiceItemAnyType = (*_CycServiceItemAnyType)(nil)
var _ CycServiceItemTypeRequirements = (*_CycServiceItemAnyType)(nil)

// NewCycServiceItemAnyType factory function for _CycServiceItemAnyType
func NewCycServiceItemAnyType(byteLength uint8, syntaxId uint8, transportSize TransportSize, length uint16, dbNumber uint16, memoryArea MemoryArea, address uint32) *_CycServiceItemAnyType {
	_result := &_CycServiceItemAnyType{
		CycServiceItemTypeContract: NewCycServiceItemType(byteLength, syntaxId),
		TransportSize:              transportSize,
		Length:                     length,
		DbNumber:                   dbNumber,
		MemoryArea:                 memoryArea,
		Address:                    address,
	}
	_result.CycServiceItemTypeContract.(*_CycServiceItemType)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CycServiceItemAnyTypeBuilder is a builder for CycServiceItemAnyType
type CycServiceItemAnyTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(transportSize TransportSize, length uint16, dbNumber uint16, memoryArea MemoryArea, address uint32) CycServiceItemAnyTypeBuilder
	// WithTransportSize adds TransportSize (property field)
	WithTransportSize(TransportSize) CycServiceItemAnyTypeBuilder
	// WithLength adds Length (property field)
	WithLength(uint16) CycServiceItemAnyTypeBuilder
	// WithDbNumber adds DbNumber (property field)
	WithDbNumber(uint16) CycServiceItemAnyTypeBuilder
	// WithMemoryArea adds MemoryArea (property field)
	WithMemoryArea(MemoryArea) CycServiceItemAnyTypeBuilder
	// WithAddress adds Address (property field)
	WithAddress(uint32) CycServiceItemAnyTypeBuilder
	// Build builds the CycServiceItemAnyType or returns an error if something is wrong
	Build() (CycServiceItemAnyType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CycServiceItemAnyType
}

// NewCycServiceItemAnyTypeBuilder() creates a CycServiceItemAnyTypeBuilder
func NewCycServiceItemAnyTypeBuilder() CycServiceItemAnyTypeBuilder {
	return &_CycServiceItemAnyTypeBuilder{_CycServiceItemAnyType: new(_CycServiceItemAnyType)}
}

type _CycServiceItemAnyTypeBuilder struct {
	*_CycServiceItemAnyType

	err *utils.MultiError
}

var _ (CycServiceItemAnyTypeBuilder) = (*_CycServiceItemAnyTypeBuilder)(nil)

func (m *_CycServiceItemAnyTypeBuilder) WithMandatoryFields(transportSize TransportSize, length uint16, dbNumber uint16, memoryArea MemoryArea, address uint32) CycServiceItemAnyTypeBuilder {
	return m.WithTransportSize(transportSize).WithLength(length).WithDbNumber(dbNumber).WithMemoryArea(memoryArea).WithAddress(address)
}

func (m *_CycServiceItemAnyTypeBuilder) WithTransportSize(transportSize TransportSize) CycServiceItemAnyTypeBuilder {
	m.TransportSize = transportSize
	return m
}

func (m *_CycServiceItemAnyTypeBuilder) WithLength(length uint16) CycServiceItemAnyTypeBuilder {
	m.Length = length
	return m
}

func (m *_CycServiceItemAnyTypeBuilder) WithDbNumber(dbNumber uint16) CycServiceItemAnyTypeBuilder {
	m.DbNumber = dbNumber
	return m
}

func (m *_CycServiceItemAnyTypeBuilder) WithMemoryArea(memoryArea MemoryArea) CycServiceItemAnyTypeBuilder {
	m.MemoryArea = memoryArea
	return m
}

func (m *_CycServiceItemAnyTypeBuilder) WithAddress(address uint32) CycServiceItemAnyTypeBuilder {
	m.Address = address
	return m
}

func (m *_CycServiceItemAnyTypeBuilder) Build() (CycServiceItemAnyType, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._CycServiceItemAnyType.deepCopy(), nil
}

func (m *_CycServiceItemAnyTypeBuilder) MustBuild() CycServiceItemAnyType {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_CycServiceItemAnyTypeBuilder) DeepCopy() any {
	return m.CreateCycServiceItemAnyTypeBuilder()
}

// CreateCycServiceItemAnyTypeBuilder creates a CycServiceItemAnyTypeBuilder
func (m *_CycServiceItemAnyType) CreateCycServiceItemAnyTypeBuilder() CycServiceItemAnyTypeBuilder {
	if m == nil {
		return NewCycServiceItemAnyTypeBuilder()
	}
	return &_CycServiceItemAnyTypeBuilder{_CycServiceItemAnyType: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CycServiceItemAnyType) GetParent() CycServiceItemTypeContract {
	return m.CycServiceItemTypeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CycServiceItemAnyType) GetTransportSize() TransportSize {
	return m.TransportSize
}

func (m *_CycServiceItemAnyType) GetLength() uint16 {
	return m.Length
}

func (m *_CycServiceItemAnyType) GetDbNumber() uint16 {
	return m.DbNumber
}

func (m *_CycServiceItemAnyType) GetMemoryArea() MemoryArea {
	return m.MemoryArea
}

func (m *_CycServiceItemAnyType) GetAddress() uint32 {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCycServiceItemAnyType(structType any) CycServiceItemAnyType {
	if casted, ok := structType.(CycServiceItemAnyType); ok {
		return casted
	}
	if casted, ok := structType.(*CycServiceItemAnyType); ok {
		return *casted
	}
	return nil
}

func (m *_CycServiceItemAnyType) GetTypeName() string {
	return "CycServiceItemAnyType"
}

func (m *_CycServiceItemAnyType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CycServiceItemTypeContract.(*_CycServiceItemType).getLengthInBits(ctx))

	// Enum Field (transportSize)
	lengthInBits += 8

	// Simple field (length)
	lengthInBits += 16

	// Simple field (dbNumber)
	lengthInBits += 16

	// Simple field (memoryArea)
	lengthInBits += 8

	// Simple field (address)
	lengthInBits += 24

	return lengthInBits
}

func (m *_CycServiceItemAnyType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CycServiceItemAnyType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CycServiceItemType) (__cycServiceItemAnyType CycServiceItemAnyType, err error) {
	m.CycServiceItemTypeContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CycServiceItemAnyType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CycServiceItemAnyType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	transportSize, err := ReadEnumField[TransportSize](ctx, "transportSize", "TransportSize", ReadEnum[TransportSize, uint8](TransportSizeFirstEnumForFieldCode, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transportSize' field"))
	}
	m.TransportSize = transportSize

	length, err := ReadSimpleField(ctx, "length", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'length' field"))
	}
	m.Length = length

	dbNumber, err := ReadSimpleField(ctx, "dbNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dbNumber' field"))
	}
	m.DbNumber = dbNumber

	memoryArea, err := ReadEnumField[MemoryArea](ctx, "memoryArea", "MemoryArea", ReadEnum(MemoryAreaByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'memoryArea' field"))
	}
	m.MemoryArea = memoryArea

	address, err := ReadSimpleField(ctx, "address", ReadUnsignedInt(readBuffer, uint8(24)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	if closeErr := readBuffer.CloseContext("CycServiceItemAnyType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CycServiceItemAnyType")
	}

	return m, nil
}

func (m *_CycServiceItemAnyType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CycServiceItemAnyType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CycServiceItemAnyType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CycServiceItemAnyType")
		}

		if err := WriteEnumField(ctx, "transportSize", "TransportSize", m.GetTransportSize(), WriteEnum[TransportSize, uint8](TransportSize.GetCode, TransportSize.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'transportSize' field")
		}

		if err := WriteSimpleField[uint16](ctx, "length", m.GetLength(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'length' field")
		}

		if err := WriteSimpleField[uint16](ctx, "dbNumber", m.GetDbNumber(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'dbNumber' field")
		}

		if err := WriteSimpleEnumField[MemoryArea](ctx, "memoryArea", "MemoryArea", m.GetMemoryArea(), WriteEnum[MemoryArea, uint8](MemoryArea.GetValue, MemoryArea.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'memoryArea' field")
		}

		if err := WriteSimpleField[uint32](ctx, "address", m.GetAddress(), WriteUnsignedInt(writeBuffer, 24)); err != nil {
			return errors.Wrap(err, "Error serializing 'address' field")
		}

		if popErr := writeBuffer.PopContext("CycServiceItemAnyType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CycServiceItemAnyType")
		}
		return nil
	}
	return m.CycServiceItemTypeContract.(*_CycServiceItemType).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CycServiceItemAnyType) IsCycServiceItemAnyType() {}

func (m *_CycServiceItemAnyType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CycServiceItemAnyType) deepCopy() *_CycServiceItemAnyType {
	if m == nil {
		return nil
	}
	_CycServiceItemAnyTypeCopy := &_CycServiceItemAnyType{
		m.CycServiceItemTypeContract.(*_CycServiceItemType).deepCopy(),
		m.TransportSize,
		m.Length,
		m.DbNumber,
		m.MemoryArea,
		m.Address,
	}
	m.CycServiceItemTypeContract.(*_CycServiceItemType)._SubType = m
	return _CycServiceItemAnyTypeCopy
}

func (m *_CycServiceItemAnyType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
