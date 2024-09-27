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

// BACnetTimerStateChangeValueDateTime is the corresponding interface of BACnetTimerStateChangeValueDateTime
type BACnetTimerStateChangeValueDateTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetTimerStateChangeValue
	// GetDateTimeValue returns DateTimeValue (property field)
	GetDateTimeValue() BACnetDateTimeEnclosed
	// IsBACnetTimerStateChangeValueDateTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimerStateChangeValueDateTime()
	// CreateBuilder creates a BACnetTimerStateChangeValueDateTimeBuilder
	CreateBACnetTimerStateChangeValueDateTimeBuilder() BACnetTimerStateChangeValueDateTimeBuilder
}

// _BACnetTimerStateChangeValueDateTime is the data-structure of this message
type _BACnetTimerStateChangeValueDateTime struct {
	BACnetTimerStateChangeValueContract
	DateTimeValue BACnetDateTimeEnclosed
}

var _ BACnetTimerStateChangeValueDateTime = (*_BACnetTimerStateChangeValueDateTime)(nil)
var _ BACnetTimerStateChangeValueRequirements = (*_BACnetTimerStateChangeValueDateTime)(nil)

// NewBACnetTimerStateChangeValueDateTime factory function for _BACnetTimerStateChangeValueDateTime
func NewBACnetTimerStateChangeValueDateTime(peekedTagHeader BACnetTagHeader, dateTimeValue BACnetDateTimeEnclosed, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueDateTime {
	if dateTimeValue == nil {
		panic("dateTimeValue of type BACnetDateTimeEnclosed for BACnetTimerStateChangeValueDateTime must not be nil")
	}
	_result := &_BACnetTimerStateChangeValueDateTime{
		BACnetTimerStateChangeValueContract: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
		DateTimeValue:                       dateTimeValue,
	}
	_result.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTimerStateChangeValueDateTimeBuilder is a builder for BACnetTimerStateChangeValueDateTime
type BACnetTimerStateChangeValueDateTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dateTimeValue BACnetDateTimeEnclosed) BACnetTimerStateChangeValueDateTimeBuilder
	// WithDateTimeValue adds DateTimeValue (property field)
	WithDateTimeValue(BACnetDateTimeEnclosed) BACnetTimerStateChangeValueDateTimeBuilder
	// WithDateTimeValueBuilder adds DateTimeValue (property field) which is build by the builder
	WithDateTimeValueBuilder(func(BACnetDateTimeEnclosedBuilder) BACnetDateTimeEnclosedBuilder) BACnetTimerStateChangeValueDateTimeBuilder
	// Build builds the BACnetTimerStateChangeValueDateTime or returns an error if something is wrong
	Build() (BACnetTimerStateChangeValueDateTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTimerStateChangeValueDateTime
}

// NewBACnetTimerStateChangeValueDateTimeBuilder() creates a BACnetTimerStateChangeValueDateTimeBuilder
func NewBACnetTimerStateChangeValueDateTimeBuilder() BACnetTimerStateChangeValueDateTimeBuilder {
	return &_BACnetTimerStateChangeValueDateTimeBuilder{_BACnetTimerStateChangeValueDateTime: new(_BACnetTimerStateChangeValueDateTime)}
}

type _BACnetTimerStateChangeValueDateTimeBuilder struct {
	*_BACnetTimerStateChangeValueDateTime

	err *utils.MultiError
}

var _ (BACnetTimerStateChangeValueDateTimeBuilder) = (*_BACnetTimerStateChangeValueDateTimeBuilder)(nil)

func (m *_BACnetTimerStateChangeValueDateTimeBuilder) WithMandatoryFields(dateTimeValue BACnetDateTimeEnclosed) BACnetTimerStateChangeValueDateTimeBuilder {
	return m.WithDateTimeValue(dateTimeValue)
}

func (m *_BACnetTimerStateChangeValueDateTimeBuilder) WithDateTimeValue(dateTimeValue BACnetDateTimeEnclosed) BACnetTimerStateChangeValueDateTimeBuilder {
	m.DateTimeValue = dateTimeValue
	return m
}

func (m *_BACnetTimerStateChangeValueDateTimeBuilder) WithDateTimeValueBuilder(builderSupplier func(BACnetDateTimeEnclosedBuilder) BACnetDateTimeEnclosedBuilder) BACnetTimerStateChangeValueDateTimeBuilder {
	builder := builderSupplier(m.DateTimeValue.CreateBACnetDateTimeEnclosedBuilder())
	var err error
	m.DateTimeValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetDateTimeEnclosedBuilder failed"))
	}
	return m
}

func (m *_BACnetTimerStateChangeValueDateTimeBuilder) Build() (BACnetTimerStateChangeValueDateTime, error) {
	if m.DateTimeValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'dateTimeValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetTimerStateChangeValueDateTime.deepCopy(), nil
}

func (m *_BACnetTimerStateChangeValueDateTimeBuilder) MustBuild() BACnetTimerStateChangeValueDateTime {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetTimerStateChangeValueDateTimeBuilder) DeepCopy() any {
	return m.CreateBACnetTimerStateChangeValueDateTimeBuilder()
}

// CreateBACnetTimerStateChangeValueDateTimeBuilder creates a BACnetTimerStateChangeValueDateTimeBuilder
func (m *_BACnetTimerStateChangeValueDateTime) CreateBACnetTimerStateChangeValueDateTimeBuilder() BACnetTimerStateChangeValueDateTimeBuilder {
	if m == nil {
		return NewBACnetTimerStateChangeValueDateTimeBuilder()
	}
	return &_BACnetTimerStateChangeValueDateTimeBuilder{_BACnetTimerStateChangeValueDateTime: m.deepCopy()}
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

func (m *_BACnetTimerStateChangeValueDateTime) GetParent() BACnetTimerStateChangeValueContract {
	return m.BACnetTimerStateChangeValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueDateTime) GetDateTimeValue() BACnetDateTimeEnclosed {
	return m.DateTimeValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueDateTime(structType any) BACnetTimerStateChangeValueDateTime {
	if casted, ok := structType.(BACnetTimerStateChangeValueDateTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueDateTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueDateTime) GetTypeName() string {
	return "BACnetTimerStateChangeValueDateTime"
}

func (m *_BACnetTimerStateChangeValueDateTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).getLengthInBits(ctx))

	// Simple field (dateTimeValue)
	lengthInBits += m.DateTimeValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueDateTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimerStateChangeValueDateTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimerStateChangeValue, objectTypeArgument BACnetObjectType) (__bACnetTimerStateChangeValueDateTime BACnetTimerStateChangeValueDateTime, err error) {
	m.BACnetTimerStateChangeValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueDateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueDateTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dateTimeValue, err := ReadSimpleField[BACnetDateTimeEnclosed](ctx, "dateTimeValue", ReadComplex[BACnetDateTimeEnclosed](BACnetDateTimeEnclosedParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dateTimeValue' field"))
	}
	m.DateTimeValue = dateTimeValue

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueDateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueDateTime")
	}

	return m, nil
}

func (m *_BACnetTimerStateChangeValueDateTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueDateTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueDateTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueDateTime")
		}

		if err := WriteSimpleField[BACnetDateTimeEnclosed](ctx, "dateTimeValue", m.GetDateTimeValue(), WriteComplex[BACnetDateTimeEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dateTimeValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueDateTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueDateTime")
		}
		return nil
	}
	return m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueDateTime) IsBACnetTimerStateChangeValueDateTime() {}

func (m *_BACnetTimerStateChangeValueDateTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTimerStateChangeValueDateTime) deepCopy() *_BACnetTimerStateChangeValueDateTime {
	if m == nil {
		return nil
	}
	_BACnetTimerStateChangeValueDateTimeCopy := &_BACnetTimerStateChangeValueDateTime{
		m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).deepCopy(),
		m.DateTimeValue.DeepCopy().(BACnetDateTimeEnclosed),
	}
	m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = m
	return _BACnetTimerStateChangeValueDateTimeCopy
}

func (m *_BACnetTimerStateChangeValueDateTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
