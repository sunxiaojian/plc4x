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

// GroupObjectDescriptorRealisationType1 is the corresponding interface of GroupObjectDescriptorRealisationType1
type GroupObjectDescriptorRealisationType1 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetDataPointer returns DataPointer (property field)
	GetDataPointer() uint8
	// GetTransmitEnable returns TransmitEnable (property field)
	GetTransmitEnable() bool
	// GetSegmentSelectorEnable returns SegmentSelectorEnable (property field)
	GetSegmentSelectorEnable() bool
	// GetWriteEnable returns WriteEnable (property field)
	GetWriteEnable() bool
	// GetReadEnable returns ReadEnable (property field)
	GetReadEnable() bool
	// GetCommunicationEnable returns CommunicationEnable (property field)
	GetCommunicationEnable() bool
	// GetPriority returns Priority (property field)
	GetPriority() CEMIPriority
	// GetValueType returns ValueType (property field)
	GetValueType() ComObjectValueType
	// IsGroupObjectDescriptorRealisationType1 is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsGroupObjectDescriptorRealisationType1()
	// CreateBuilder creates a GroupObjectDescriptorRealisationType1Builder
	CreateGroupObjectDescriptorRealisationType1Builder() GroupObjectDescriptorRealisationType1Builder
}

// _GroupObjectDescriptorRealisationType1 is the data-structure of this message
type _GroupObjectDescriptorRealisationType1 struct {
	DataPointer           uint8
	TransmitEnable        bool
	SegmentSelectorEnable bool
	WriteEnable           bool
	ReadEnable            bool
	CommunicationEnable   bool
	Priority              CEMIPriority
	ValueType             ComObjectValueType
	// Reserved Fields
	reservedField0 *uint8
}

var _ GroupObjectDescriptorRealisationType1 = (*_GroupObjectDescriptorRealisationType1)(nil)

// NewGroupObjectDescriptorRealisationType1 factory function for _GroupObjectDescriptorRealisationType1
func NewGroupObjectDescriptorRealisationType1(dataPointer uint8, transmitEnable bool, segmentSelectorEnable bool, writeEnable bool, readEnable bool, communicationEnable bool, priority CEMIPriority, valueType ComObjectValueType) *_GroupObjectDescriptorRealisationType1 {
	return &_GroupObjectDescriptorRealisationType1{DataPointer: dataPointer, TransmitEnable: transmitEnable, SegmentSelectorEnable: segmentSelectorEnable, WriteEnable: writeEnable, ReadEnable: readEnable, CommunicationEnable: communicationEnable, Priority: priority, ValueType: valueType}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// GroupObjectDescriptorRealisationType1Builder is a builder for GroupObjectDescriptorRealisationType1
type GroupObjectDescriptorRealisationType1Builder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dataPointer uint8, transmitEnable bool, segmentSelectorEnable bool, writeEnable bool, readEnable bool, communicationEnable bool, priority CEMIPriority, valueType ComObjectValueType) GroupObjectDescriptorRealisationType1Builder
	// WithDataPointer adds DataPointer (property field)
	WithDataPointer(uint8) GroupObjectDescriptorRealisationType1Builder
	// WithTransmitEnable adds TransmitEnable (property field)
	WithTransmitEnable(bool) GroupObjectDescriptorRealisationType1Builder
	// WithSegmentSelectorEnable adds SegmentSelectorEnable (property field)
	WithSegmentSelectorEnable(bool) GroupObjectDescriptorRealisationType1Builder
	// WithWriteEnable adds WriteEnable (property field)
	WithWriteEnable(bool) GroupObjectDescriptorRealisationType1Builder
	// WithReadEnable adds ReadEnable (property field)
	WithReadEnable(bool) GroupObjectDescriptorRealisationType1Builder
	// WithCommunicationEnable adds CommunicationEnable (property field)
	WithCommunicationEnable(bool) GroupObjectDescriptorRealisationType1Builder
	// WithPriority adds Priority (property field)
	WithPriority(CEMIPriority) GroupObjectDescriptorRealisationType1Builder
	// WithValueType adds ValueType (property field)
	WithValueType(ComObjectValueType) GroupObjectDescriptorRealisationType1Builder
	// Build builds the GroupObjectDescriptorRealisationType1 or returns an error if something is wrong
	Build() (GroupObjectDescriptorRealisationType1, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() GroupObjectDescriptorRealisationType1
}

// NewGroupObjectDescriptorRealisationType1Builder() creates a GroupObjectDescriptorRealisationType1Builder
func NewGroupObjectDescriptorRealisationType1Builder() GroupObjectDescriptorRealisationType1Builder {
	return &_GroupObjectDescriptorRealisationType1Builder{_GroupObjectDescriptorRealisationType1: new(_GroupObjectDescriptorRealisationType1)}
}

type _GroupObjectDescriptorRealisationType1Builder struct {
	*_GroupObjectDescriptorRealisationType1

	err *utils.MultiError
}

var _ (GroupObjectDescriptorRealisationType1Builder) = (*_GroupObjectDescriptorRealisationType1Builder)(nil)

func (m *_GroupObjectDescriptorRealisationType1Builder) WithMandatoryFields(dataPointer uint8, transmitEnable bool, segmentSelectorEnable bool, writeEnable bool, readEnable bool, communicationEnable bool, priority CEMIPriority, valueType ComObjectValueType) GroupObjectDescriptorRealisationType1Builder {
	return m.WithDataPointer(dataPointer).WithTransmitEnable(transmitEnable).WithSegmentSelectorEnable(segmentSelectorEnable).WithWriteEnable(writeEnable).WithReadEnable(readEnable).WithCommunicationEnable(communicationEnable).WithPriority(priority).WithValueType(valueType)
}

func (m *_GroupObjectDescriptorRealisationType1Builder) WithDataPointer(dataPointer uint8) GroupObjectDescriptorRealisationType1Builder {
	m.DataPointer = dataPointer
	return m
}

func (m *_GroupObjectDescriptorRealisationType1Builder) WithTransmitEnable(transmitEnable bool) GroupObjectDescriptorRealisationType1Builder {
	m.TransmitEnable = transmitEnable
	return m
}

func (m *_GroupObjectDescriptorRealisationType1Builder) WithSegmentSelectorEnable(segmentSelectorEnable bool) GroupObjectDescriptorRealisationType1Builder {
	m.SegmentSelectorEnable = segmentSelectorEnable
	return m
}

func (m *_GroupObjectDescriptorRealisationType1Builder) WithWriteEnable(writeEnable bool) GroupObjectDescriptorRealisationType1Builder {
	m.WriteEnable = writeEnable
	return m
}

func (m *_GroupObjectDescriptorRealisationType1Builder) WithReadEnable(readEnable bool) GroupObjectDescriptorRealisationType1Builder {
	m.ReadEnable = readEnable
	return m
}

func (m *_GroupObjectDescriptorRealisationType1Builder) WithCommunicationEnable(communicationEnable bool) GroupObjectDescriptorRealisationType1Builder {
	m.CommunicationEnable = communicationEnable
	return m
}

func (m *_GroupObjectDescriptorRealisationType1Builder) WithPriority(priority CEMIPriority) GroupObjectDescriptorRealisationType1Builder {
	m.Priority = priority
	return m
}

func (m *_GroupObjectDescriptorRealisationType1Builder) WithValueType(valueType ComObjectValueType) GroupObjectDescriptorRealisationType1Builder {
	m.ValueType = valueType
	return m
}

func (m *_GroupObjectDescriptorRealisationType1Builder) Build() (GroupObjectDescriptorRealisationType1, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._GroupObjectDescriptorRealisationType1.deepCopy(), nil
}

func (m *_GroupObjectDescriptorRealisationType1Builder) MustBuild() GroupObjectDescriptorRealisationType1 {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_GroupObjectDescriptorRealisationType1Builder) DeepCopy() any {
	return m.CreateGroupObjectDescriptorRealisationType1Builder()
}

// CreateGroupObjectDescriptorRealisationType1Builder creates a GroupObjectDescriptorRealisationType1Builder
func (m *_GroupObjectDescriptorRealisationType1) CreateGroupObjectDescriptorRealisationType1Builder() GroupObjectDescriptorRealisationType1Builder {
	if m == nil {
		return NewGroupObjectDescriptorRealisationType1Builder()
	}
	return &_GroupObjectDescriptorRealisationType1Builder{_GroupObjectDescriptorRealisationType1: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GroupObjectDescriptorRealisationType1) GetDataPointer() uint8 {
	return m.DataPointer
}

func (m *_GroupObjectDescriptorRealisationType1) GetTransmitEnable() bool {
	return m.TransmitEnable
}

func (m *_GroupObjectDescriptorRealisationType1) GetSegmentSelectorEnable() bool {
	return m.SegmentSelectorEnable
}

func (m *_GroupObjectDescriptorRealisationType1) GetWriteEnable() bool {
	return m.WriteEnable
}

func (m *_GroupObjectDescriptorRealisationType1) GetReadEnable() bool {
	return m.ReadEnable
}

func (m *_GroupObjectDescriptorRealisationType1) GetCommunicationEnable() bool {
	return m.CommunicationEnable
}

func (m *_GroupObjectDescriptorRealisationType1) GetPriority() CEMIPriority {
	return m.Priority
}

func (m *_GroupObjectDescriptorRealisationType1) GetValueType() ComObjectValueType {
	return m.ValueType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastGroupObjectDescriptorRealisationType1(structType any) GroupObjectDescriptorRealisationType1 {
	if casted, ok := structType.(GroupObjectDescriptorRealisationType1); ok {
		return casted
	}
	if casted, ok := structType.(*GroupObjectDescriptorRealisationType1); ok {
		return *casted
	}
	return nil
}

func (m *_GroupObjectDescriptorRealisationType1) GetTypeName() string {
	return "GroupObjectDescriptorRealisationType1"
}

func (m *_GroupObjectDescriptorRealisationType1) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (dataPointer)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (transmitEnable)
	lengthInBits += 1

	// Simple field (segmentSelectorEnable)
	lengthInBits += 1

	// Simple field (writeEnable)
	lengthInBits += 1

	// Simple field (readEnable)
	lengthInBits += 1

	// Simple field (communicationEnable)
	lengthInBits += 1

	// Simple field (priority)
	lengthInBits += 2

	// Simple field (valueType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_GroupObjectDescriptorRealisationType1) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func GroupObjectDescriptorRealisationType1Parse(ctx context.Context, theBytes []byte) (GroupObjectDescriptorRealisationType1, error) {
	return GroupObjectDescriptorRealisationType1ParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func GroupObjectDescriptorRealisationType1ParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationType1, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationType1, error) {
		return GroupObjectDescriptorRealisationType1ParseWithBuffer(ctx, readBuffer)
	}
}

func GroupObjectDescriptorRealisationType1ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (GroupObjectDescriptorRealisationType1, error) {
	v, err := (&_GroupObjectDescriptorRealisationType1{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_GroupObjectDescriptorRealisationType1) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__groupObjectDescriptorRealisationType1 GroupObjectDescriptorRealisationType1, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GroupObjectDescriptorRealisationType1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GroupObjectDescriptorRealisationType1")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataPointer, err := ReadSimpleField(ctx, "dataPointer", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataPointer' field"))
	}
	m.DataPointer = dataPointer

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(1)), uint8(0x1))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	transmitEnable, err := ReadSimpleField(ctx, "transmitEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transmitEnable' field"))
	}
	m.TransmitEnable = transmitEnable

	segmentSelectorEnable, err := ReadSimpleField(ctx, "segmentSelectorEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentSelectorEnable' field"))
	}
	m.SegmentSelectorEnable = segmentSelectorEnable

	writeEnable, err := ReadSimpleField(ctx, "writeEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writeEnable' field"))
	}
	m.WriteEnable = writeEnable

	readEnable, err := ReadSimpleField(ctx, "readEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readEnable' field"))
	}
	m.ReadEnable = readEnable

	communicationEnable, err := ReadSimpleField(ctx, "communicationEnable", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'communicationEnable' field"))
	}
	m.CommunicationEnable = communicationEnable

	priority, err := ReadEnumField[CEMIPriority](ctx, "priority", "CEMIPriority", ReadEnum(CEMIPriorityByValue, ReadUnsignedByte(readBuffer, uint8(2))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	m.Priority = priority

	valueType, err := ReadEnumField[ComObjectValueType](ctx, "valueType", "ComObjectValueType", ReadEnum(ComObjectValueTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'valueType' field"))
	}
	m.ValueType = valueType

	if closeErr := readBuffer.CloseContext("GroupObjectDescriptorRealisationType1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GroupObjectDescriptorRealisationType1")
	}

	return m, nil
}

func (m *_GroupObjectDescriptorRealisationType1) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GroupObjectDescriptorRealisationType1) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("GroupObjectDescriptorRealisationType1"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for GroupObjectDescriptorRealisationType1")
	}

	if err := WriteSimpleField[uint8](ctx, "dataPointer", m.GetDataPointer(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'dataPointer' field")
	}

	if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x1), WriteUnsignedByte(writeBuffer, 1)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteSimpleField[bool](ctx, "transmitEnable", m.GetTransmitEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'transmitEnable' field")
	}

	if err := WriteSimpleField[bool](ctx, "segmentSelectorEnable", m.GetSegmentSelectorEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'segmentSelectorEnable' field")
	}

	if err := WriteSimpleField[bool](ctx, "writeEnable", m.GetWriteEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'writeEnable' field")
	}

	if err := WriteSimpleField[bool](ctx, "readEnable", m.GetReadEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'readEnable' field")
	}

	if err := WriteSimpleField[bool](ctx, "communicationEnable", m.GetCommunicationEnable(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'communicationEnable' field")
	}

	if err := WriteSimpleEnumField[CEMIPriority](ctx, "priority", "CEMIPriority", m.GetPriority(), WriteEnum[CEMIPriority, uint8](CEMIPriority.GetValue, CEMIPriority.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 2))); err != nil {
		return errors.Wrap(err, "Error serializing 'priority' field")
	}

	if err := WriteSimpleEnumField[ComObjectValueType](ctx, "valueType", "ComObjectValueType", m.GetValueType(), WriteEnum[ComObjectValueType, uint8](ComObjectValueType.GetValue, ComObjectValueType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'valueType' field")
	}

	if popErr := writeBuffer.PopContext("GroupObjectDescriptorRealisationType1"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for GroupObjectDescriptorRealisationType1")
	}
	return nil
}

func (m *_GroupObjectDescriptorRealisationType1) IsGroupObjectDescriptorRealisationType1() {}

func (m *_GroupObjectDescriptorRealisationType1) DeepCopy() any {
	return m.deepCopy()
}

func (m *_GroupObjectDescriptorRealisationType1) deepCopy() *_GroupObjectDescriptorRealisationType1 {
	if m == nil {
		return nil
	}
	_GroupObjectDescriptorRealisationType1Copy := &_GroupObjectDescriptorRealisationType1{
		m.DataPointer,
		m.TransmitEnable,
		m.SegmentSelectorEnable,
		m.WriteEnable,
		m.ReadEnable,
		m.CommunicationEnable,
		m.Priority,
		m.ValueType,
		m.reservedField0,
	}
	return _GroupObjectDescriptorRealisationType1Copy
}

func (m *_GroupObjectDescriptorRealisationType1) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
