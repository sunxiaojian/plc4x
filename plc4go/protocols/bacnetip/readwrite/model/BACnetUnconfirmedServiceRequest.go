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

// BACnetUnconfirmedServiceRequest is the corresponding interface of BACnetUnconfirmedServiceRequest
type BACnetUnconfirmedServiceRequest interface {
	BACnetUnconfirmedServiceRequestContract
	BACnetUnconfirmedServiceRequestRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsBACnetUnconfirmedServiceRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetUnconfirmedServiceRequest()
	// CreateBuilder creates a BACnetUnconfirmedServiceRequestBuilder
	CreateBACnetUnconfirmedServiceRequestBuilder() BACnetUnconfirmedServiceRequestBuilder
}

// BACnetUnconfirmedServiceRequestContract provides a set of functions which can be overwritten by a sub struct
type BACnetUnconfirmedServiceRequestContract interface {
	// GetServiceRequestLength() returns a parser argument
	GetServiceRequestLength() uint16
	// IsBACnetUnconfirmedServiceRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetUnconfirmedServiceRequest()
	// CreateBuilder creates a BACnetUnconfirmedServiceRequestBuilder
	CreateBACnetUnconfirmedServiceRequestBuilder() BACnetUnconfirmedServiceRequestBuilder
}

// BACnetUnconfirmedServiceRequestRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetUnconfirmedServiceRequestRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetServiceChoice returns ServiceChoice (discriminator field)
	GetServiceChoice() BACnetUnconfirmedServiceChoice
}

// _BACnetUnconfirmedServiceRequest is the data-structure of this message
type _BACnetUnconfirmedServiceRequest struct {
	_SubType BACnetUnconfirmedServiceRequest

	// Arguments.
	ServiceRequestLength uint16
}

var _ BACnetUnconfirmedServiceRequestContract = (*_BACnetUnconfirmedServiceRequest)(nil)

// NewBACnetUnconfirmedServiceRequest factory function for _BACnetUnconfirmedServiceRequest
func NewBACnetUnconfirmedServiceRequest(serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequest {
	return &_BACnetUnconfirmedServiceRequest{ServiceRequestLength: serviceRequestLength}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetUnconfirmedServiceRequestBuilder is a builder for BACnetUnconfirmedServiceRequest
type BACnetUnconfirmedServiceRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetUnconfirmedServiceRequestBuilder
	// Build builds the BACnetUnconfirmedServiceRequest or returns an error if something is wrong
	Build() (BACnetUnconfirmedServiceRequestContract, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetUnconfirmedServiceRequestContract
}

// NewBACnetUnconfirmedServiceRequestBuilder() creates a BACnetUnconfirmedServiceRequestBuilder
func NewBACnetUnconfirmedServiceRequestBuilder() BACnetUnconfirmedServiceRequestBuilder {
	return &_BACnetUnconfirmedServiceRequestBuilder{_BACnetUnconfirmedServiceRequest: new(_BACnetUnconfirmedServiceRequest)}
}

type _BACnetUnconfirmedServiceRequestBuilder struct {
	*_BACnetUnconfirmedServiceRequest

	err *utils.MultiError
}

var _ (BACnetUnconfirmedServiceRequestBuilder) = (*_BACnetUnconfirmedServiceRequestBuilder)(nil)

func (m *_BACnetUnconfirmedServiceRequestBuilder) WithMandatoryFields() BACnetUnconfirmedServiceRequestBuilder {
	return m
}

func (m *_BACnetUnconfirmedServiceRequestBuilder) Build() (BACnetUnconfirmedServiceRequestContract, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetUnconfirmedServiceRequest.deepCopy(), nil
}

func (m *_BACnetUnconfirmedServiceRequestBuilder) MustBuild() BACnetUnconfirmedServiceRequestContract {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetUnconfirmedServiceRequestBuilder) DeepCopy() any {
	return m.CreateBACnetUnconfirmedServiceRequestBuilder()
}

// CreateBACnetUnconfirmedServiceRequestBuilder creates a BACnetUnconfirmedServiceRequestBuilder
func (m *_BACnetUnconfirmedServiceRequest) CreateBACnetUnconfirmedServiceRequestBuilder() BACnetUnconfirmedServiceRequestBuilder {
	if m == nil {
		return NewBACnetUnconfirmedServiceRequestBuilder()
	}
	return &_BACnetUnconfirmedServiceRequestBuilder{_BACnetUnconfirmedServiceRequest: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequest(structType any) BACnetUnconfirmedServiceRequest {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequest); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequest); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequest) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequest"
}

func (m *_BACnetUnconfirmedServiceRequest) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetUnconfirmedServiceRequestParse[T BACnetUnconfirmedServiceRequest](ctx context.Context, theBytes []byte, serviceRequestLength uint16) (T, error) {
	return BACnetUnconfirmedServiceRequestParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetUnconfirmedServiceRequestParseWithBufferProducer[T BACnetUnconfirmedServiceRequest](serviceRequestLength uint16) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetUnconfirmedServiceRequestParseWithBuffer[T](ctx, readBuffer, serviceRequestLength)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetUnconfirmedServiceRequestParseWithBuffer[T BACnetUnconfirmedServiceRequest](ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint16) (T, error) {
	v, err := (&_BACnetUnconfirmedServiceRequest{ServiceRequestLength: serviceRequestLength}).parse(ctx, readBuffer, serviceRequestLength)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_BACnetUnconfirmedServiceRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint16) (__bACnetUnconfirmedServiceRequest BACnetUnconfirmedServiceRequest, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	serviceChoice, err := ReadDiscriminatorEnumField[BACnetUnconfirmedServiceChoice](ctx, "serviceChoice", "BACnetUnconfirmedServiceChoice", ReadEnum(BACnetUnconfirmedServiceChoiceByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceChoice' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetUnconfirmedServiceRequest
	switch {
	case serviceChoice == BACnetUnconfirmedServiceChoice_I_AM: // BACnetUnconfirmedServiceRequestIAm
		if _child, err = new(_BACnetUnconfirmedServiceRequestIAm).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetUnconfirmedServiceRequestIAm for type-switch of BACnetUnconfirmedServiceRequest")
		}
	case serviceChoice == BACnetUnconfirmedServiceChoice_I_HAVE: // BACnetUnconfirmedServiceRequestIHave
		if _child, err = new(_BACnetUnconfirmedServiceRequestIHave).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetUnconfirmedServiceRequestIHave for type-switch of BACnetUnconfirmedServiceRequest")
		}
	case serviceChoice == BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION: // BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification
		if _child, err = new(_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification for type-switch of BACnetUnconfirmedServiceRequest")
		}
	case serviceChoice == BACnetUnconfirmedServiceChoice_UNCONFIRMED_EVENT_NOTIFICATION: // BACnetUnconfirmedServiceRequestUnconfirmedEventNotification
		if _child, err = new(_BACnetUnconfirmedServiceRequestUnconfirmedEventNotification).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetUnconfirmedServiceRequestUnconfirmedEventNotification for type-switch of BACnetUnconfirmedServiceRequest")
		}
	case serviceChoice == BACnetUnconfirmedServiceChoice_UNCONFIRMED_PRIVATE_TRANSFER: // BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
		if _child, err = new(_BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer for type-switch of BACnetUnconfirmedServiceRequest")
		}
	case serviceChoice == BACnetUnconfirmedServiceChoice_UNCONFIRMED_TEXT_MESSAGE: // BACnetUnconfirmedServiceRequestUnconfirmedTextMessage
		if _child, err = new(_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetUnconfirmedServiceRequestUnconfirmedTextMessage for type-switch of BACnetUnconfirmedServiceRequest")
		}
	case serviceChoice == BACnetUnconfirmedServiceChoice_TIME_SYNCHRONIZATION: // BACnetUnconfirmedServiceRequestTimeSynchronization
		if _child, err = new(_BACnetUnconfirmedServiceRequestTimeSynchronization).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetUnconfirmedServiceRequestTimeSynchronization for type-switch of BACnetUnconfirmedServiceRequest")
		}
	case serviceChoice == BACnetUnconfirmedServiceChoice_WHO_HAS: // BACnetUnconfirmedServiceRequestWhoHas
		if _child, err = new(_BACnetUnconfirmedServiceRequestWhoHas).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetUnconfirmedServiceRequestWhoHas for type-switch of BACnetUnconfirmedServiceRequest")
		}
	case serviceChoice == BACnetUnconfirmedServiceChoice_WHO_IS: // BACnetUnconfirmedServiceRequestWhoIs
		if _child, err = new(_BACnetUnconfirmedServiceRequestWhoIs).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetUnconfirmedServiceRequestWhoIs for type-switch of BACnetUnconfirmedServiceRequest")
		}
	case serviceChoice == BACnetUnconfirmedServiceChoice_UTC_TIME_SYNCHRONIZATION: // BACnetUnconfirmedServiceRequestUTCTimeSynchronization
		if _child, err = new(_BACnetUnconfirmedServiceRequestUTCTimeSynchronization).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetUnconfirmedServiceRequestUTCTimeSynchronization for type-switch of BACnetUnconfirmedServiceRequest")
		}
	case serviceChoice == BACnetUnconfirmedServiceChoice_WRITE_GROUP: // BACnetUnconfirmedServiceRequestWriteGroup
		if _child, err = new(_BACnetUnconfirmedServiceRequestWriteGroup).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetUnconfirmedServiceRequestWriteGroup for type-switch of BACnetUnconfirmedServiceRequest")
		}
	case serviceChoice == BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION_MULTIPLE: // BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
		if _child, err = new(_BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple for type-switch of BACnetUnconfirmedServiceRequest")
		}
	case 0 == 0: // BACnetUnconfirmedServiceRequestUnknown
		if _child, err = new(_BACnetUnconfirmedServiceRequestUnknown).parse(ctx, readBuffer, m, serviceRequestLength); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetUnconfirmedServiceRequestUnknown for type-switch of BACnetUnconfirmedServiceRequest")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [serviceChoice=%v]", serviceChoice)
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequest")
	}

	return _child, nil
}

func (pm *_BACnetUnconfirmedServiceRequest) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetUnconfirmedServiceRequest, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequest"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequest")
	}

	if err := WriteDiscriminatorEnumField(ctx, "serviceChoice", "BACnetUnconfirmedServiceChoice", m.GetServiceChoice(), WriteEnum[BACnetUnconfirmedServiceChoice, uint8](BACnetUnconfirmedServiceChoice.GetValue, BACnetUnconfirmedServiceChoice.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'serviceChoice' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequest"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequest")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetUnconfirmedServiceRequest) GetServiceRequestLength() uint16 {
	return m.ServiceRequestLength
}

//
////

func (m *_BACnetUnconfirmedServiceRequest) IsBACnetUnconfirmedServiceRequest() {}

func (m *_BACnetUnconfirmedServiceRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetUnconfirmedServiceRequest) deepCopy() *_BACnetUnconfirmedServiceRequest {
	if m == nil {
		return nil
	}
	_BACnetUnconfirmedServiceRequestCopy := &_BACnetUnconfirmedServiceRequest{
		nil, // will be set by child
		m.ServiceRequestLength,
	}
	return _BACnetUnconfirmedServiceRequestCopy
}
