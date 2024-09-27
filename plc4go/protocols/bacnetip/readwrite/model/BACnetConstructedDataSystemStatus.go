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

// BACnetConstructedDataSystemStatus is the corresponding interface of BACnetConstructedDataSystemStatus
type BACnetConstructedDataSystemStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetSystemStatus returns SystemStatus (property field)
	GetSystemStatus() BACnetDeviceStatusTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceStatusTagged
	// IsBACnetConstructedDataSystemStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataSystemStatus()
	// CreateBuilder creates a BACnetConstructedDataSystemStatusBuilder
	CreateBACnetConstructedDataSystemStatusBuilder() BACnetConstructedDataSystemStatusBuilder
}

// _BACnetConstructedDataSystemStatus is the data-structure of this message
type _BACnetConstructedDataSystemStatus struct {
	BACnetConstructedDataContract
	SystemStatus BACnetDeviceStatusTagged
}

var _ BACnetConstructedDataSystemStatus = (*_BACnetConstructedDataSystemStatus)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataSystemStatus)(nil)

// NewBACnetConstructedDataSystemStatus factory function for _BACnetConstructedDataSystemStatus
func NewBACnetConstructedDataSystemStatus(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, systemStatus BACnetDeviceStatusTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSystemStatus {
	if systemStatus == nil {
		panic("systemStatus of type BACnetDeviceStatusTagged for BACnetConstructedDataSystemStatus must not be nil")
	}
	_result := &_BACnetConstructedDataSystemStatus{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		SystemStatus:                  systemStatus,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataSystemStatusBuilder is a builder for BACnetConstructedDataSystemStatus
type BACnetConstructedDataSystemStatusBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(systemStatus BACnetDeviceStatusTagged) BACnetConstructedDataSystemStatusBuilder
	// WithSystemStatus adds SystemStatus (property field)
	WithSystemStatus(BACnetDeviceStatusTagged) BACnetConstructedDataSystemStatusBuilder
	// WithSystemStatusBuilder adds SystemStatus (property field) which is build by the builder
	WithSystemStatusBuilder(func(BACnetDeviceStatusTaggedBuilder) BACnetDeviceStatusTaggedBuilder) BACnetConstructedDataSystemStatusBuilder
	// Build builds the BACnetConstructedDataSystemStatus or returns an error if something is wrong
	Build() (BACnetConstructedDataSystemStatus, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataSystemStatus
}

// NewBACnetConstructedDataSystemStatusBuilder() creates a BACnetConstructedDataSystemStatusBuilder
func NewBACnetConstructedDataSystemStatusBuilder() BACnetConstructedDataSystemStatusBuilder {
	return &_BACnetConstructedDataSystemStatusBuilder{_BACnetConstructedDataSystemStatus: new(_BACnetConstructedDataSystemStatus)}
}

type _BACnetConstructedDataSystemStatusBuilder struct {
	*_BACnetConstructedDataSystemStatus

	err *utils.MultiError
}

var _ (BACnetConstructedDataSystemStatusBuilder) = (*_BACnetConstructedDataSystemStatusBuilder)(nil)

func (m *_BACnetConstructedDataSystemStatusBuilder) WithMandatoryFields(systemStatus BACnetDeviceStatusTagged) BACnetConstructedDataSystemStatusBuilder {
	return m.WithSystemStatus(systemStatus)
}

func (m *_BACnetConstructedDataSystemStatusBuilder) WithSystemStatus(systemStatus BACnetDeviceStatusTagged) BACnetConstructedDataSystemStatusBuilder {
	m.SystemStatus = systemStatus
	return m
}

func (m *_BACnetConstructedDataSystemStatusBuilder) WithSystemStatusBuilder(builderSupplier func(BACnetDeviceStatusTaggedBuilder) BACnetDeviceStatusTaggedBuilder) BACnetConstructedDataSystemStatusBuilder {
	builder := builderSupplier(m.SystemStatus.CreateBACnetDeviceStatusTaggedBuilder())
	var err error
	m.SystemStatus, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetDeviceStatusTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataSystemStatusBuilder) Build() (BACnetConstructedDataSystemStatus, error) {
	if m.SystemStatus == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'systemStatus' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataSystemStatus.deepCopy(), nil
}

func (m *_BACnetConstructedDataSystemStatusBuilder) MustBuild() BACnetConstructedDataSystemStatus {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataSystemStatusBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataSystemStatusBuilder()
}

// CreateBACnetConstructedDataSystemStatusBuilder creates a BACnetConstructedDataSystemStatusBuilder
func (m *_BACnetConstructedDataSystemStatus) CreateBACnetConstructedDataSystemStatusBuilder() BACnetConstructedDataSystemStatusBuilder {
	if m == nil {
		return NewBACnetConstructedDataSystemStatusBuilder()
	}
	return &_BACnetConstructedDataSystemStatusBuilder{_BACnetConstructedDataSystemStatus: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSystemStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSystemStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SYSTEM_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSystemStatus) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSystemStatus) GetSystemStatus() BACnetDeviceStatusTagged {
	return m.SystemStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSystemStatus) GetActualValue() BACnetDeviceStatusTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceStatusTagged(m.GetSystemStatus())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSystemStatus(structType any) BACnetConstructedDataSystemStatus {
	if casted, ok := structType.(BACnetConstructedDataSystemStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSystemStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSystemStatus) GetTypeName() string {
	return "BACnetConstructedDataSystemStatus"
}

func (m *_BACnetConstructedDataSystemStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (systemStatus)
	lengthInBits += m.SystemStatus.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSystemStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataSystemStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataSystemStatus BACnetConstructedDataSystemStatus, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSystemStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSystemStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	systemStatus, err := ReadSimpleField[BACnetDeviceStatusTagged](ctx, "systemStatus", ReadComplex[BACnetDeviceStatusTagged](BACnetDeviceStatusTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'systemStatus' field"))
	}
	m.SystemStatus = systemStatus

	actualValue, err := ReadVirtualField[BACnetDeviceStatusTagged](ctx, "actualValue", (*BACnetDeviceStatusTagged)(nil), systemStatus)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSystemStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSystemStatus")
	}

	return m, nil
}

func (m *_BACnetConstructedDataSystemStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSystemStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSystemStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSystemStatus")
		}

		if err := WriteSimpleField[BACnetDeviceStatusTagged](ctx, "systemStatus", m.GetSystemStatus(), WriteComplex[BACnetDeviceStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'systemStatus' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSystemStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSystemStatus")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSystemStatus) IsBACnetConstructedDataSystemStatus() {}

func (m *_BACnetConstructedDataSystemStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataSystemStatus) deepCopy() *_BACnetConstructedDataSystemStatus {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataSystemStatusCopy := &_BACnetConstructedDataSystemStatus{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.SystemStatus.DeepCopy().(BACnetDeviceStatusTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataSystemStatusCopy
}

func (m *_BACnetConstructedDataSystemStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
