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

// BACnetConstructedDataActivationTime is the corresponding interface of BACnetConstructedDataActivationTime
type BACnetConstructedDataActivationTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetActivationTime returns ActivationTime (property field)
	GetActivationTime() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
	// IsBACnetConstructedDataActivationTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataActivationTime()
	// CreateBuilder creates a BACnetConstructedDataActivationTimeBuilder
	CreateBACnetConstructedDataActivationTimeBuilder() BACnetConstructedDataActivationTimeBuilder
}

// _BACnetConstructedDataActivationTime is the data-structure of this message
type _BACnetConstructedDataActivationTime struct {
	BACnetConstructedDataContract
	ActivationTime BACnetDateTime
}

var _ BACnetConstructedDataActivationTime = (*_BACnetConstructedDataActivationTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataActivationTime)(nil)

// NewBACnetConstructedDataActivationTime factory function for _BACnetConstructedDataActivationTime
func NewBACnetConstructedDataActivationTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, activationTime BACnetDateTime, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataActivationTime {
	if activationTime == nil {
		panic("activationTime of type BACnetDateTime for BACnetConstructedDataActivationTime must not be nil")
	}
	_result := &_BACnetConstructedDataActivationTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ActivationTime:                activationTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataActivationTimeBuilder is a builder for BACnetConstructedDataActivationTime
type BACnetConstructedDataActivationTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(activationTime BACnetDateTime) BACnetConstructedDataActivationTimeBuilder
	// WithActivationTime adds ActivationTime (property field)
	WithActivationTime(BACnetDateTime) BACnetConstructedDataActivationTimeBuilder
	// WithActivationTimeBuilder adds ActivationTime (property field) which is build by the builder
	WithActivationTimeBuilder(func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataActivationTimeBuilder
	// Build builds the BACnetConstructedDataActivationTime or returns an error if something is wrong
	Build() (BACnetConstructedDataActivationTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataActivationTime
}

// NewBACnetConstructedDataActivationTimeBuilder() creates a BACnetConstructedDataActivationTimeBuilder
func NewBACnetConstructedDataActivationTimeBuilder() BACnetConstructedDataActivationTimeBuilder {
	return &_BACnetConstructedDataActivationTimeBuilder{_BACnetConstructedDataActivationTime: new(_BACnetConstructedDataActivationTime)}
}

type _BACnetConstructedDataActivationTimeBuilder struct {
	*_BACnetConstructedDataActivationTime

	err *utils.MultiError
}

var _ (BACnetConstructedDataActivationTimeBuilder) = (*_BACnetConstructedDataActivationTimeBuilder)(nil)

func (m *_BACnetConstructedDataActivationTimeBuilder) WithMandatoryFields(activationTime BACnetDateTime) BACnetConstructedDataActivationTimeBuilder {
	return m.WithActivationTime(activationTime)
}

func (m *_BACnetConstructedDataActivationTimeBuilder) WithActivationTime(activationTime BACnetDateTime) BACnetConstructedDataActivationTimeBuilder {
	m.ActivationTime = activationTime
	return m
}

func (m *_BACnetConstructedDataActivationTimeBuilder) WithActivationTimeBuilder(builderSupplier func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataActivationTimeBuilder {
	builder := builderSupplier(m.ActivationTime.CreateBACnetDateTimeBuilder())
	var err error
	m.ActivationTime, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetDateTimeBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataActivationTimeBuilder) Build() (BACnetConstructedDataActivationTime, error) {
	if m.ActivationTime == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'activationTime' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataActivationTime.deepCopy(), nil
}

func (m *_BACnetConstructedDataActivationTimeBuilder) MustBuild() BACnetConstructedDataActivationTime {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataActivationTimeBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataActivationTimeBuilder()
}

// CreateBACnetConstructedDataActivationTimeBuilder creates a BACnetConstructedDataActivationTimeBuilder
func (m *_BACnetConstructedDataActivationTime) CreateBACnetConstructedDataActivationTimeBuilder() BACnetConstructedDataActivationTimeBuilder {
	if m == nil {
		return NewBACnetConstructedDataActivationTimeBuilder()
	}
	return &_BACnetConstructedDataActivationTimeBuilder{_BACnetConstructedDataActivationTime: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataActivationTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataActivationTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACTIVATION_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataActivationTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataActivationTime) GetActivationTime() BACnetDateTime {
	return m.ActivationTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataActivationTime) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetActivationTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataActivationTime(structType any) BACnetConstructedDataActivationTime {
	if casted, ok := structType.(BACnetConstructedDataActivationTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataActivationTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataActivationTime) GetTypeName() string {
	return "BACnetConstructedDataActivationTime"
}

func (m *_BACnetConstructedDataActivationTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (activationTime)
	lengthInBits += m.ActivationTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataActivationTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataActivationTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataActivationTime BACnetConstructedDataActivationTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataActivationTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataActivationTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	activationTime, err := ReadSimpleField[BACnetDateTime](ctx, "activationTime", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'activationTime' field"))
	}
	m.ActivationTime = activationTime

	actualValue, err := ReadVirtualField[BACnetDateTime](ctx, "actualValue", (*BACnetDateTime)(nil), activationTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataActivationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataActivationTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataActivationTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataActivationTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataActivationTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataActivationTime")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "activationTime", m.GetActivationTime(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'activationTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataActivationTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataActivationTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataActivationTime) IsBACnetConstructedDataActivationTime() {}

func (m *_BACnetConstructedDataActivationTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataActivationTime) deepCopy() *_BACnetConstructedDataActivationTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataActivationTimeCopy := &_BACnetConstructedDataActivationTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.ActivationTime.DeepCopy().(BACnetDateTime),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataActivationTimeCopy
}

func (m *_BACnetConstructedDataActivationTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
