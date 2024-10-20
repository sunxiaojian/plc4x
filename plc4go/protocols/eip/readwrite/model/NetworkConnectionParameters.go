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

// NetworkConnectionParameters is the corresponding interface of NetworkConnectionParameters
type NetworkConnectionParameters interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetConnectionSize returns ConnectionSize (property field)
	GetConnectionSize() uint16
	// GetOwner returns Owner (property field)
	GetOwner() bool
	// GetConnectionType returns ConnectionType (property field)
	GetConnectionType() uint8
	// GetPriority returns Priority (property field)
	GetPriority() uint8
	// GetConnectionSizeType returns ConnectionSizeType (property field)
	GetConnectionSizeType() bool
	// IsNetworkConnectionParameters is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNetworkConnectionParameters()
	// CreateBuilder creates a NetworkConnectionParametersBuilder
	CreateNetworkConnectionParametersBuilder() NetworkConnectionParametersBuilder
}

// _NetworkConnectionParameters is the data-structure of this message
type _NetworkConnectionParameters struct {
	ConnectionSize     uint16
	Owner              bool
	ConnectionType     uint8
	Priority           uint8
	ConnectionSizeType bool
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *bool
	reservedField2 *bool
}

var _ NetworkConnectionParameters = (*_NetworkConnectionParameters)(nil)

// NewNetworkConnectionParameters factory function for _NetworkConnectionParameters
func NewNetworkConnectionParameters(connectionSize uint16, owner bool, connectionType uint8, priority uint8, connectionSizeType bool) *_NetworkConnectionParameters {
	return &_NetworkConnectionParameters{ConnectionSize: connectionSize, Owner: owner, ConnectionType: connectionType, Priority: priority, ConnectionSizeType: connectionSizeType}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NetworkConnectionParametersBuilder is a builder for NetworkConnectionParameters
type NetworkConnectionParametersBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(connectionSize uint16, owner bool, connectionType uint8, priority uint8, connectionSizeType bool) NetworkConnectionParametersBuilder
	// WithConnectionSize adds ConnectionSize (property field)
	WithConnectionSize(uint16) NetworkConnectionParametersBuilder
	// WithOwner adds Owner (property field)
	WithOwner(bool) NetworkConnectionParametersBuilder
	// WithConnectionType adds ConnectionType (property field)
	WithConnectionType(uint8) NetworkConnectionParametersBuilder
	// WithPriority adds Priority (property field)
	WithPriority(uint8) NetworkConnectionParametersBuilder
	// WithConnectionSizeType adds ConnectionSizeType (property field)
	WithConnectionSizeType(bool) NetworkConnectionParametersBuilder
	// Build builds the NetworkConnectionParameters or returns an error if something is wrong
	Build() (NetworkConnectionParameters, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NetworkConnectionParameters
}

// NewNetworkConnectionParametersBuilder() creates a NetworkConnectionParametersBuilder
func NewNetworkConnectionParametersBuilder() NetworkConnectionParametersBuilder {
	return &_NetworkConnectionParametersBuilder{_NetworkConnectionParameters: new(_NetworkConnectionParameters)}
}

type _NetworkConnectionParametersBuilder struct {
	*_NetworkConnectionParameters

	err *utils.MultiError
}

var _ (NetworkConnectionParametersBuilder) = (*_NetworkConnectionParametersBuilder)(nil)

func (b *_NetworkConnectionParametersBuilder) WithMandatoryFields(connectionSize uint16, owner bool, connectionType uint8, priority uint8, connectionSizeType bool) NetworkConnectionParametersBuilder {
	return b.WithConnectionSize(connectionSize).WithOwner(owner).WithConnectionType(connectionType).WithPriority(priority).WithConnectionSizeType(connectionSizeType)
}

func (b *_NetworkConnectionParametersBuilder) WithConnectionSize(connectionSize uint16) NetworkConnectionParametersBuilder {
	b.ConnectionSize = connectionSize
	return b
}

func (b *_NetworkConnectionParametersBuilder) WithOwner(owner bool) NetworkConnectionParametersBuilder {
	b.Owner = owner
	return b
}

func (b *_NetworkConnectionParametersBuilder) WithConnectionType(connectionType uint8) NetworkConnectionParametersBuilder {
	b.ConnectionType = connectionType
	return b
}

func (b *_NetworkConnectionParametersBuilder) WithPriority(priority uint8) NetworkConnectionParametersBuilder {
	b.Priority = priority
	return b
}

func (b *_NetworkConnectionParametersBuilder) WithConnectionSizeType(connectionSizeType bool) NetworkConnectionParametersBuilder {
	b.ConnectionSizeType = connectionSizeType
	return b
}

func (b *_NetworkConnectionParametersBuilder) Build() (NetworkConnectionParameters, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NetworkConnectionParameters.deepCopy(), nil
}

func (b *_NetworkConnectionParametersBuilder) MustBuild() NetworkConnectionParameters {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_NetworkConnectionParametersBuilder) DeepCopy() any {
	_copy := b.CreateNetworkConnectionParametersBuilder().(*_NetworkConnectionParametersBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNetworkConnectionParametersBuilder creates a NetworkConnectionParametersBuilder
func (b *_NetworkConnectionParameters) CreateNetworkConnectionParametersBuilder() NetworkConnectionParametersBuilder {
	if b == nil {
		return NewNetworkConnectionParametersBuilder()
	}
	return &_NetworkConnectionParametersBuilder{_NetworkConnectionParameters: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NetworkConnectionParameters) GetConnectionSize() uint16 {
	return m.ConnectionSize
}

func (m *_NetworkConnectionParameters) GetOwner() bool {
	return m.Owner
}

func (m *_NetworkConnectionParameters) GetConnectionType() uint8 {
	return m.ConnectionType
}

func (m *_NetworkConnectionParameters) GetPriority() uint8 {
	return m.Priority
}

func (m *_NetworkConnectionParameters) GetConnectionSizeType() bool {
	return m.ConnectionSizeType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNetworkConnectionParameters(structType any) NetworkConnectionParameters {
	if casted, ok := structType.(NetworkConnectionParameters); ok {
		return casted
	}
	if casted, ok := structType.(*NetworkConnectionParameters); ok {
		return *casted
	}
	return nil
}

func (m *_NetworkConnectionParameters) GetTypeName() string {
	return "NetworkConnectionParameters"
}

func (m *_NetworkConnectionParameters) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (connectionSize)
	lengthInBits += 16

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (owner)
	lengthInBits += 1

	// Simple field (connectionType)
	lengthInBits += 2

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (priority)
	lengthInBits += 2

	// Simple field (connectionSizeType)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	return lengthInBits
}

func (m *_NetworkConnectionParameters) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NetworkConnectionParametersParse(ctx context.Context, theBytes []byte) (NetworkConnectionParameters, error) {
	return NetworkConnectionParametersParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NetworkConnectionParametersParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (NetworkConnectionParameters, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (NetworkConnectionParameters, error) {
		return NetworkConnectionParametersParseWithBuffer(ctx, readBuffer)
	}
}

func NetworkConnectionParametersParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NetworkConnectionParameters, error) {
	v, err := (&_NetworkConnectionParameters{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_NetworkConnectionParameters) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__networkConnectionParameters NetworkConnectionParameters, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NetworkConnectionParameters"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NetworkConnectionParameters")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	connectionSize, err := ReadSimpleField(ctx, "connectionSize", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionSize' field"))
	}
	m.ConnectionSize = connectionSize

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	owner, err := ReadSimpleField(ctx, "owner", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'owner' field"))
	}
	m.Owner = owner

	connectionType, err := ReadSimpleField(ctx, "connectionType", ReadUnsignedByte(readBuffer, uint8(2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionType' field"))
	}
	m.ConnectionType = connectionType

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	priority, err := ReadSimpleField(ctx, "priority", ReadUnsignedByte(readBuffer, uint8(2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	m.Priority = priority

	connectionSizeType, err := ReadSimpleField(ctx, "connectionSizeType", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionSizeType' field"))
	}
	m.ConnectionSizeType = connectionSizeType

	reservedField2, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField2 = reservedField2

	if closeErr := readBuffer.CloseContext("NetworkConnectionParameters"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NetworkConnectionParameters")
	}

	return m, nil
}

func (m *_NetworkConnectionParameters) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NetworkConnectionParameters) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NetworkConnectionParameters"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NetworkConnectionParameters")
	}

	if err := WriteSimpleField[uint16](ctx, "connectionSize", m.GetConnectionSize(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'connectionSize' field")
	}

	if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteSimpleField[bool](ctx, "owner", m.GetOwner(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'owner' field")
	}

	if err := WriteSimpleField[uint8](ctx, "connectionType", m.GetConnectionType(), WriteUnsignedByte(writeBuffer, 2)); err != nil {
		return errors.Wrap(err, "Error serializing 'connectionType' field")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 2")
	}

	if err := WriteSimpleField[uint8](ctx, "priority", m.GetPriority(), WriteUnsignedByte(writeBuffer, 2)); err != nil {
		return errors.Wrap(err, "Error serializing 'priority' field")
	}

	if err := WriteSimpleField[bool](ctx, "connectionSizeType", m.GetConnectionSizeType(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'connectionSizeType' field")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 3")
	}

	if popErr := writeBuffer.PopContext("NetworkConnectionParameters"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NetworkConnectionParameters")
	}
	return nil
}

func (m *_NetworkConnectionParameters) IsNetworkConnectionParameters() {}

func (m *_NetworkConnectionParameters) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NetworkConnectionParameters) deepCopy() *_NetworkConnectionParameters {
	if m == nil {
		return nil
	}
	_NetworkConnectionParametersCopy := &_NetworkConnectionParameters{
		m.ConnectionSize,
		m.Owner,
		m.ConnectionType,
		m.Priority,
		m.ConnectionSizeType,
		m.reservedField0,
		m.reservedField1,
		m.reservedField2,
	}
	return _NetworkConnectionParametersCopy
}

func (m *_NetworkConnectionParameters) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
