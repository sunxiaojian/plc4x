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

// BACnetConstructedDataChangeOfStateTime is the corresponding interface of BACnetConstructedDataChangeOfStateTime
type BACnetConstructedDataChangeOfStateTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetChangeOfStateTime returns ChangeOfStateTime (property field)
	GetChangeOfStateTime() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
	// IsBACnetConstructedDataChangeOfStateTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataChangeOfStateTime()
	// CreateBuilder creates a BACnetConstructedDataChangeOfStateTimeBuilder
	CreateBACnetConstructedDataChangeOfStateTimeBuilder() BACnetConstructedDataChangeOfStateTimeBuilder
}

// _BACnetConstructedDataChangeOfStateTime is the data-structure of this message
type _BACnetConstructedDataChangeOfStateTime struct {
	BACnetConstructedDataContract
	ChangeOfStateTime BACnetDateTime
}

var _ BACnetConstructedDataChangeOfStateTime = (*_BACnetConstructedDataChangeOfStateTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataChangeOfStateTime)(nil)

// NewBACnetConstructedDataChangeOfStateTime factory function for _BACnetConstructedDataChangeOfStateTime
func NewBACnetConstructedDataChangeOfStateTime(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, changeOfStateTime BACnetDateTime, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataChangeOfStateTime {
	if changeOfStateTime == nil {
		panic("changeOfStateTime of type BACnetDateTime for BACnetConstructedDataChangeOfStateTime must not be nil")
	}
	_result := &_BACnetConstructedDataChangeOfStateTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ChangeOfStateTime:             changeOfStateTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataChangeOfStateTimeBuilder is a builder for BACnetConstructedDataChangeOfStateTime
type BACnetConstructedDataChangeOfStateTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(changeOfStateTime BACnetDateTime) BACnetConstructedDataChangeOfStateTimeBuilder
	// WithChangeOfStateTime adds ChangeOfStateTime (property field)
	WithChangeOfStateTime(BACnetDateTime) BACnetConstructedDataChangeOfStateTimeBuilder
	// WithChangeOfStateTimeBuilder adds ChangeOfStateTime (property field) which is build by the builder
	WithChangeOfStateTimeBuilder(func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataChangeOfStateTimeBuilder
	// Build builds the BACnetConstructedDataChangeOfStateTime or returns an error if something is wrong
	Build() (BACnetConstructedDataChangeOfStateTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataChangeOfStateTime
}

// NewBACnetConstructedDataChangeOfStateTimeBuilder() creates a BACnetConstructedDataChangeOfStateTimeBuilder
func NewBACnetConstructedDataChangeOfStateTimeBuilder() BACnetConstructedDataChangeOfStateTimeBuilder {
	return &_BACnetConstructedDataChangeOfStateTimeBuilder{_BACnetConstructedDataChangeOfStateTime: new(_BACnetConstructedDataChangeOfStateTime)}
}

type _BACnetConstructedDataChangeOfStateTimeBuilder struct {
	*_BACnetConstructedDataChangeOfStateTime

	err *utils.MultiError
}

var _ (BACnetConstructedDataChangeOfStateTimeBuilder) = (*_BACnetConstructedDataChangeOfStateTimeBuilder)(nil)

func (m *_BACnetConstructedDataChangeOfStateTimeBuilder) WithMandatoryFields(changeOfStateTime BACnetDateTime) BACnetConstructedDataChangeOfStateTimeBuilder {
	return m.WithChangeOfStateTime(changeOfStateTime)
}

func (m *_BACnetConstructedDataChangeOfStateTimeBuilder) WithChangeOfStateTime(changeOfStateTime BACnetDateTime) BACnetConstructedDataChangeOfStateTimeBuilder {
	m.ChangeOfStateTime = changeOfStateTime
	return m
}

func (m *_BACnetConstructedDataChangeOfStateTimeBuilder) WithChangeOfStateTimeBuilder(builderSupplier func(BACnetDateTimeBuilder) BACnetDateTimeBuilder) BACnetConstructedDataChangeOfStateTimeBuilder {
	builder := builderSupplier(m.ChangeOfStateTime.CreateBACnetDateTimeBuilder())
	var err error
	m.ChangeOfStateTime, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetDateTimeBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataChangeOfStateTimeBuilder) Build() (BACnetConstructedDataChangeOfStateTime, error) {
	if m.ChangeOfStateTime == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'changeOfStateTime' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataChangeOfStateTime.deepCopy(), nil
}

func (m *_BACnetConstructedDataChangeOfStateTimeBuilder) MustBuild() BACnetConstructedDataChangeOfStateTime {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataChangeOfStateTimeBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataChangeOfStateTimeBuilder()
}

// CreateBACnetConstructedDataChangeOfStateTimeBuilder creates a BACnetConstructedDataChangeOfStateTimeBuilder
func (m *_BACnetConstructedDataChangeOfStateTime) CreateBACnetConstructedDataChangeOfStateTimeBuilder() BACnetConstructedDataChangeOfStateTimeBuilder {
	if m == nil {
		return NewBACnetConstructedDataChangeOfStateTimeBuilder()
	}
	return &_BACnetConstructedDataChangeOfStateTimeBuilder{_BACnetConstructedDataChangeOfStateTime: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataChangeOfStateTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataChangeOfStateTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CHANGE_OF_STATE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataChangeOfStateTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataChangeOfStateTime) GetChangeOfStateTime() BACnetDateTime {
	return m.ChangeOfStateTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataChangeOfStateTime) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetChangeOfStateTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataChangeOfStateTime(structType any) BACnetConstructedDataChangeOfStateTime {
	if casted, ok := structType.(BACnetConstructedDataChangeOfStateTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataChangeOfStateTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataChangeOfStateTime) GetTypeName() string {
	return "BACnetConstructedDataChangeOfStateTime"
}

func (m *_BACnetConstructedDataChangeOfStateTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (changeOfStateTime)
	lengthInBits += m.ChangeOfStateTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataChangeOfStateTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataChangeOfStateTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataChangeOfStateTime BACnetConstructedDataChangeOfStateTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataChangeOfStateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataChangeOfStateTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	changeOfStateTime, err := ReadSimpleField[BACnetDateTime](ctx, "changeOfStateTime", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'changeOfStateTime' field"))
	}
	m.ChangeOfStateTime = changeOfStateTime

	actualValue, err := ReadVirtualField[BACnetDateTime](ctx, "actualValue", (*BACnetDateTime)(nil), changeOfStateTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataChangeOfStateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataChangeOfStateTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataChangeOfStateTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataChangeOfStateTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataChangeOfStateTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataChangeOfStateTime")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "changeOfStateTime", m.GetChangeOfStateTime(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'changeOfStateTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataChangeOfStateTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataChangeOfStateTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataChangeOfStateTime) IsBACnetConstructedDataChangeOfStateTime() {}

func (m *_BACnetConstructedDataChangeOfStateTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataChangeOfStateTime) deepCopy() *_BACnetConstructedDataChangeOfStateTime {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataChangeOfStateTimeCopy := &_BACnetConstructedDataChangeOfStateTime{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.ChangeOfStateTime.DeepCopy().(BACnetDateTime),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataChangeOfStateTimeCopy
}

func (m *_BACnetConstructedDataChangeOfStateTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
