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

// BACnetObjectTypesSupportedTagged is the corresponding interface of BACnetObjectTypesSupportedTagged
type BACnetObjectTypesSupportedTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBitString
	// GetTimeValue returns TimeValue (virtual field)
	GetTimeValue() bool
	// GetNotificationForwarder returns NotificationForwarder (virtual field)
	GetNotificationForwarder() bool
	// GetAlertEnrollment returns AlertEnrollment (virtual field)
	GetAlertEnrollment() bool
	// GetChannel returns Channel (virtual field)
	GetChannel() bool
	// GetLightingOutput returns LightingOutput (virtual field)
	GetLightingOutput() bool
	// GetBinaryLightingOutput returns BinaryLightingOutput (virtual field)
	GetBinaryLightingOutput() bool
	// GetNetworkPort returns NetworkPort (virtual field)
	GetNetworkPort() bool
	// GetElevatorGroup returns ElevatorGroup (virtual field)
	GetElevatorGroup() bool
	// GetEscalator returns Escalator (virtual field)
	GetEscalator() bool
	// GetLift returns Lift (virtual field)
	GetLift() bool
	// IsBACnetObjectTypesSupportedTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetObjectTypesSupportedTagged()
	// CreateBuilder creates a BACnetObjectTypesSupportedTaggedBuilder
	CreateBACnetObjectTypesSupportedTaggedBuilder() BACnetObjectTypesSupportedTaggedBuilder
}

// _BACnetObjectTypesSupportedTagged is the data-structure of this message
type _BACnetObjectTypesSupportedTagged struct {
	Header  BACnetTagHeader
	Payload BACnetTagPayloadBitString

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetObjectTypesSupportedTagged = (*_BACnetObjectTypesSupportedTagged)(nil)

// NewBACnetObjectTypesSupportedTagged factory function for _BACnetObjectTypesSupportedTagged
func NewBACnetObjectTypesSupportedTagged(header BACnetTagHeader, payload BACnetTagPayloadBitString, tagNumber uint8, tagClass TagClass) *_BACnetObjectTypesSupportedTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetObjectTypesSupportedTagged must not be nil")
	}
	if payload == nil {
		panic("payload of type BACnetTagPayloadBitString for BACnetObjectTypesSupportedTagged must not be nil")
	}
	return &_BACnetObjectTypesSupportedTagged{Header: header, Payload: payload, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetObjectTypesSupportedTaggedBuilder is a builder for BACnetObjectTypesSupportedTagged
type BACnetObjectTypesSupportedTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, payload BACnetTagPayloadBitString) BACnetObjectTypesSupportedTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetObjectTypesSupportedTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetObjectTypesSupportedTaggedBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadBitString) BACnetObjectTypesSupportedTaggedBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadBitStringBuilder) BACnetTagPayloadBitStringBuilder) BACnetObjectTypesSupportedTaggedBuilder
	// Build builds the BACnetObjectTypesSupportedTagged or returns an error if something is wrong
	Build() (BACnetObjectTypesSupportedTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetObjectTypesSupportedTagged
}

// NewBACnetObjectTypesSupportedTaggedBuilder() creates a BACnetObjectTypesSupportedTaggedBuilder
func NewBACnetObjectTypesSupportedTaggedBuilder() BACnetObjectTypesSupportedTaggedBuilder {
	return &_BACnetObjectTypesSupportedTaggedBuilder{_BACnetObjectTypesSupportedTagged: new(_BACnetObjectTypesSupportedTagged)}
}

type _BACnetObjectTypesSupportedTaggedBuilder struct {
	*_BACnetObjectTypesSupportedTagged

	err *utils.MultiError
}

var _ (BACnetObjectTypesSupportedTaggedBuilder) = (*_BACnetObjectTypesSupportedTaggedBuilder)(nil)

func (m *_BACnetObjectTypesSupportedTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, payload BACnetTagPayloadBitString) BACnetObjectTypesSupportedTaggedBuilder {
	return m.WithHeader(header).WithPayload(payload)
}

func (m *_BACnetObjectTypesSupportedTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetObjectTypesSupportedTaggedBuilder {
	m.Header = header
	return m
}

func (m *_BACnetObjectTypesSupportedTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetObjectTypesSupportedTaggedBuilder {
	builder := builderSupplier(m.Header.CreateBACnetTagHeaderBuilder())
	var err error
	m.Header, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return m
}

func (m *_BACnetObjectTypesSupportedTaggedBuilder) WithPayload(payload BACnetTagPayloadBitString) BACnetObjectTypesSupportedTaggedBuilder {
	m.Payload = payload
	return m
}

func (m *_BACnetObjectTypesSupportedTaggedBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadBitStringBuilder) BACnetTagPayloadBitStringBuilder) BACnetObjectTypesSupportedTaggedBuilder {
	builder := builderSupplier(m.Payload.CreateBACnetTagPayloadBitStringBuilder())
	var err error
	m.Payload, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetTagPayloadBitStringBuilder failed"))
	}
	return m
}

func (m *_BACnetObjectTypesSupportedTaggedBuilder) Build() (BACnetObjectTypesSupportedTagged, error) {
	if m.Header == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if m.Payload == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetObjectTypesSupportedTagged.deepCopy(), nil
}

func (m *_BACnetObjectTypesSupportedTaggedBuilder) MustBuild() BACnetObjectTypesSupportedTagged {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetObjectTypesSupportedTaggedBuilder) DeepCopy() any {
	return m.CreateBACnetObjectTypesSupportedTaggedBuilder()
}

// CreateBACnetObjectTypesSupportedTaggedBuilder creates a BACnetObjectTypesSupportedTaggedBuilder
func (m *_BACnetObjectTypesSupportedTagged) CreateBACnetObjectTypesSupportedTaggedBuilder() BACnetObjectTypesSupportedTaggedBuilder {
	if m == nil {
		return NewBACnetObjectTypesSupportedTaggedBuilder()
	}
	return &_BACnetObjectTypesSupportedTaggedBuilder{_BACnetObjectTypesSupportedTagged: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetObjectTypesSupportedTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetObjectTypesSupportedTagged) GetPayload() BACnetTagPayloadBitString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetObjectTypesSupportedTagged) GetTimeValue() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (0))), func() any { return bool(m.GetPayload().GetData()[0]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetNotificationForwarder() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (1))), func() any { return bool(m.GetPayload().GetData()[1]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetAlertEnrollment() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (2))), func() any { return bool(m.GetPayload().GetData()[2]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetChannel() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (3))), func() any { return bool(m.GetPayload().GetData()[3]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetLightingOutput() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (4))), func() any { return bool(m.GetPayload().GetData()[4]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetBinaryLightingOutput() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (5))), func() any { return bool(m.GetPayload().GetData()[5]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetNetworkPort() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (6))), func() any { return bool(m.GetPayload().GetData()[6]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetElevatorGroup() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (7))), func() any { return bool(m.GetPayload().GetData()[7]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetEscalator() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (8))), func() any { return bool(m.GetPayload().GetData()[8]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetObjectTypesSupportedTagged) GetLift() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (9))), func() any { return bool(m.GetPayload().GetData()[9]) }, func() any { return bool(bool(false)) }).(bool))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetObjectTypesSupportedTagged(structType any) BACnetObjectTypesSupportedTagged {
	if casted, ok := structType.(BACnetObjectTypesSupportedTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetObjectTypesSupportedTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetObjectTypesSupportedTagged) GetTypeName() string {
	return "BACnetObjectTypesSupportedTagged"
}

func (m *_BACnetObjectTypesSupportedTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetObjectTypesSupportedTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetObjectTypesSupportedTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetObjectTypesSupportedTagged, error) {
	return BACnetObjectTypesSupportedTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetObjectTypesSupportedTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetObjectTypesSupportedTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetObjectTypesSupportedTagged, error) {
		return BACnetObjectTypesSupportedTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetObjectTypesSupportedTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetObjectTypesSupportedTagged, error) {
	v, err := (&_BACnetObjectTypesSupportedTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetObjectTypesSupportedTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetObjectTypesSupportedTagged BACnetObjectTypesSupportedTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetObjectTypesSupportedTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetObjectTypesSupportedTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	payload, err := ReadSimpleField[BACnetTagPayloadBitString](ctx, "payload", ReadComplex[BACnetTagPayloadBitString](BACnetTagPayloadBitStringParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	timeValue, err := ReadVirtualField[bool](ctx, "timeValue", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (0))), func() any { return bool(payload.GetData()[0]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeValue' field"))
	}
	_ = timeValue

	notificationForwarder, err := ReadVirtualField[bool](ctx, "notificationForwarder", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (1))), func() any { return bool(payload.GetData()[1]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationForwarder' field"))
	}
	_ = notificationForwarder

	alertEnrollment, err := ReadVirtualField[bool](ctx, "alertEnrollment", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (2))), func() any { return bool(payload.GetData()[2]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alertEnrollment' field"))
	}
	_ = alertEnrollment

	channel, err := ReadVirtualField[bool](ctx, "channel", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (3))), func() any { return bool(payload.GetData()[3]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channel' field"))
	}
	_ = channel

	lightingOutput, err := ReadVirtualField[bool](ctx, "lightingOutput", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (4))), func() any { return bool(payload.GetData()[4]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lightingOutput' field"))
	}
	_ = lightingOutput

	binaryLightingOutput, err := ReadVirtualField[bool](ctx, "binaryLightingOutput", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (5))), func() any { return bool(payload.GetData()[5]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'binaryLightingOutput' field"))
	}
	_ = binaryLightingOutput

	networkPort, err := ReadVirtualField[bool](ctx, "networkPort", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (6))), func() any { return bool(payload.GetData()[6]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkPort' field"))
	}
	_ = networkPort

	elevatorGroup, err := ReadVirtualField[bool](ctx, "elevatorGroup", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (7))), func() any { return bool(payload.GetData()[7]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'elevatorGroup' field"))
	}
	_ = elevatorGroup

	escalator, err := ReadVirtualField[bool](ctx, "escalator", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (8))), func() any { return bool(payload.GetData()[8]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'escalator' field"))
	}
	_ = escalator

	lift, err := ReadVirtualField[bool](ctx, "lift", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (9))), func() any { return bool(payload.GetData()[9]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lift' field"))
	}
	_ = lift

	if closeErr := readBuffer.CloseContext("BACnetObjectTypesSupportedTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetObjectTypesSupportedTagged")
	}

	return m, nil
}

func (m *_BACnetObjectTypesSupportedTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetObjectTypesSupportedTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetObjectTypesSupportedTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetObjectTypesSupportedTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteSimpleField[BACnetTagPayloadBitString](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadBitString](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'payload' field")
	}
	// Virtual field
	timeValue := m.GetTimeValue()
	_ = timeValue
	if _timeValueErr := writeBuffer.WriteVirtual(ctx, "timeValue", m.GetTimeValue()); _timeValueErr != nil {
		return errors.Wrap(_timeValueErr, "Error serializing 'timeValue' field")
	}
	// Virtual field
	notificationForwarder := m.GetNotificationForwarder()
	_ = notificationForwarder
	if _notificationForwarderErr := writeBuffer.WriteVirtual(ctx, "notificationForwarder", m.GetNotificationForwarder()); _notificationForwarderErr != nil {
		return errors.Wrap(_notificationForwarderErr, "Error serializing 'notificationForwarder' field")
	}
	// Virtual field
	alertEnrollment := m.GetAlertEnrollment()
	_ = alertEnrollment
	if _alertEnrollmentErr := writeBuffer.WriteVirtual(ctx, "alertEnrollment", m.GetAlertEnrollment()); _alertEnrollmentErr != nil {
		return errors.Wrap(_alertEnrollmentErr, "Error serializing 'alertEnrollment' field")
	}
	// Virtual field
	channel := m.GetChannel()
	_ = channel
	if _channelErr := writeBuffer.WriteVirtual(ctx, "channel", m.GetChannel()); _channelErr != nil {
		return errors.Wrap(_channelErr, "Error serializing 'channel' field")
	}
	// Virtual field
	lightingOutput := m.GetLightingOutput()
	_ = lightingOutput
	if _lightingOutputErr := writeBuffer.WriteVirtual(ctx, "lightingOutput", m.GetLightingOutput()); _lightingOutputErr != nil {
		return errors.Wrap(_lightingOutputErr, "Error serializing 'lightingOutput' field")
	}
	// Virtual field
	binaryLightingOutput := m.GetBinaryLightingOutput()
	_ = binaryLightingOutput
	if _binaryLightingOutputErr := writeBuffer.WriteVirtual(ctx, "binaryLightingOutput", m.GetBinaryLightingOutput()); _binaryLightingOutputErr != nil {
		return errors.Wrap(_binaryLightingOutputErr, "Error serializing 'binaryLightingOutput' field")
	}
	// Virtual field
	networkPort := m.GetNetworkPort()
	_ = networkPort
	if _networkPortErr := writeBuffer.WriteVirtual(ctx, "networkPort", m.GetNetworkPort()); _networkPortErr != nil {
		return errors.Wrap(_networkPortErr, "Error serializing 'networkPort' field")
	}
	// Virtual field
	elevatorGroup := m.GetElevatorGroup()
	_ = elevatorGroup
	if _elevatorGroupErr := writeBuffer.WriteVirtual(ctx, "elevatorGroup", m.GetElevatorGroup()); _elevatorGroupErr != nil {
		return errors.Wrap(_elevatorGroupErr, "Error serializing 'elevatorGroup' field")
	}
	// Virtual field
	escalator := m.GetEscalator()
	_ = escalator
	if _escalatorErr := writeBuffer.WriteVirtual(ctx, "escalator", m.GetEscalator()); _escalatorErr != nil {
		return errors.Wrap(_escalatorErr, "Error serializing 'escalator' field")
	}
	// Virtual field
	lift := m.GetLift()
	_ = lift
	if _liftErr := writeBuffer.WriteVirtual(ctx, "lift", m.GetLift()); _liftErr != nil {
		return errors.Wrap(_liftErr, "Error serializing 'lift' field")
	}

	if popErr := writeBuffer.PopContext("BACnetObjectTypesSupportedTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetObjectTypesSupportedTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetObjectTypesSupportedTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetObjectTypesSupportedTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetObjectTypesSupportedTagged) IsBACnetObjectTypesSupportedTagged() {}

func (m *_BACnetObjectTypesSupportedTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetObjectTypesSupportedTagged) deepCopy() *_BACnetObjectTypesSupportedTagged {
	if m == nil {
		return nil
	}
	_BACnetObjectTypesSupportedTaggedCopy := &_BACnetObjectTypesSupportedTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Payload.DeepCopy().(BACnetTagPayloadBitString),
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetObjectTypesSupportedTaggedCopy
}

func (m *_BACnetObjectTypesSupportedTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
