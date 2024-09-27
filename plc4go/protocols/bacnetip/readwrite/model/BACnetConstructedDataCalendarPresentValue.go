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

// BACnetConstructedDataCalendarPresentValue is the corresponding interface of BACnetConstructedDataCalendarPresentValue
type BACnetConstructedDataCalendarPresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataCalendarPresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCalendarPresentValue()
	// CreateBuilder creates a BACnetConstructedDataCalendarPresentValueBuilder
	CreateBACnetConstructedDataCalendarPresentValueBuilder() BACnetConstructedDataCalendarPresentValueBuilder
}

// _BACnetConstructedDataCalendarPresentValue is the data-structure of this message
type _BACnetConstructedDataCalendarPresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataCalendarPresentValue = (*_BACnetConstructedDataCalendarPresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCalendarPresentValue)(nil)

// NewBACnetConstructedDataCalendarPresentValue factory function for _BACnetConstructedDataCalendarPresentValue
func NewBACnetConstructedDataCalendarPresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCalendarPresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetApplicationTagBoolean for BACnetConstructedDataCalendarPresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataCalendarPresentValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PresentValue:                  presentValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataCalendarPresentValueBuilder is a builder for BACnetConstructedDataCalendarPresentValue
type BACnetConstructedDataCalendarPresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetApplicationTagBoolean) BACnetConstructedDataCalendarPresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetApplicationTagBoolean) BACnetConstructedDataCalendarPresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataCalendarPresentValueBuilder
	// Build builds the BACnetConstructedDataCalendarPresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataCalendarPresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataCalendarPresentValue
}

// NewBACnetConstructedDataCalendarPresentValueBuilder() creates a BACnetConstructedDataCalendarPresentValueBuilder
func NewBACnetConstructedDataCalendarPresentValueBuilder() BACnetConstructedDataCalendarPresentValueBuilder {
	return &_BACnetConstructedDataCalendarPresentValueBuilder{_BACnetConstructedDataCalendarPresentValue: new(_BACnetConstructedDataCalendarPresentValue)}
}

type _BACnetConstructedDataCalendarPresentValueBuilder struct {
	*_BACnetConstructedDataCalendarPresentValue

	err *utils.MultiError
}

var _ (BACnetConstructedDataCalendarPresentValueBuilder) = (*_BACnetConstructedDataCalendarPresentValueBuilder)(nil)

func (m *_BACnetConstructedDataCalendarPresentValueBuilder) WithMandatoryFields(presentValue BACnetApplicationTagBoolean) BACnetConstructedDataCalendarPresentValueBuilder {
	return m.WithPresentValue(presentValue)
}

func (m *_BACnetConstructedDataCalendarPresentValueBuilder) WithPresentValue(presentValue BACnetApplicationTagBoolean) BACnetConstructedDataCalendarPresentValueBuilder {
	m.PresentValue = presentValue
	return m
}

func (m *_BACnetConstructedDataCalendarPresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataCalendarPresentValueBuilder {
	builder := builderSupplier(m.PresentValue.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	m.PresentValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataCalendarPresentValueBuilder) Build() (BACnetConstructedDataCalendarPresentValue, error) {
	if m.PresentValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataCalendarPresentValue.deepCopy(), nil
}

func (m *_BACnetConstructedDataCalendarPresentValueBuilder) MustBuild() BACnetConstructedDataCalendarPresentValue {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataCalendarPresentValueBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataCalendarPresentValueBuilder()
}

// CreateBACnetConstructedDataCalendarPresentValueBuilder creates a BACnetConstructedDataCalendarPresentValueBuilder
func (m *_BACnetConstructedDataCalendarPresentValue) CreateBACnetConstructedDataCalendarPresentValueBuilder() BACnetConstructedDataCalendarPresentValueBuilder {
	if m == nil {
		return NewBACnetConstructedDataCalendarPresentValueBuilder()
	}
	return &_BACnetConstructedDataCalendarPresentValueBuilder{_BACnetConstructedDataCalendarPresentValue: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCalendarPresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_CALENDAR
}

func (m *_BACnetConstructedDataCalendarPresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCalendarPresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCalendarPresentValue) GetPresentValue() BACnetApplicationTagBoolean {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCalendarPresentValue) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCalendarPresentValue(structType any) BACnetConstructedDataCalendarPresentValue {
	if casted, ok := structType.(BACnetConstructedDataCalendarPresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCalendarPresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCalendarPresentValue) GetTypeName() string {
	return "BACnetConstructedDataCalendarPresentValue"
}

func (m *_BACnetConstructedDataCalendarPresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCalendarPresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCalendarPresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCalendarPresentValue BACnetConstructedDataCalendarPresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCalendarPresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCalendarPresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "presentValue", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCalendarPresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCalendarPresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCalendarPresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCalendarPresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCalendarPresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCalendarPresentValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCalendarPresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCalendarPresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCalendarPresentValue) IsBACnetConstructedDataCalendarPresentValue() {}

func (m *_BACnetConstructedDataCalendarPresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataCalendarPresentValue) deepCopy() *_BACnetConstructedDataCalendarPresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataCalendarPresentValueCopy := &_BACnetConstructedDataCalendarPresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.PresentValue.DeepCopy().(BACnetApplicationTagBoolean),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataCalendarPresentValueCopy
}

func (m *_BACnetConstructedDataCalendarPresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
