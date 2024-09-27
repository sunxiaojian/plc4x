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

// BACnetConstructedDataAlarmValues is the corresponding interface of BACnetConstructedDataAlarmValues
type BACnetConstructedDataAlarmValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetAlarmValues returns AlarmValues (property field)
	GetAlarmValues() []BACnetLifeSafetyStateTagged
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataAlarmValues is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAlarmValues()
	// CreateBuilder creates a BACnetConstructedDataAlarmValuesBuilder
	CreateBACnetConstructedDataAlarmValuesBuilder() BACnetConstructedDataAlarmValuesBuilder
}

// _BACnetConstructedDataAlarmValues is the data-structure of this message
type _BACnetConstructedDataAlarmValues struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	AlarmValues          []BACnetLifeSafetyStateTagged
}

var _ BACnetConstructedDataAlarmValues = (*_BACnetConstructedDataAlarmValues)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAlarmValues)(nil)

// NewBACnetConstructedDataAlarmValues factory function for _BACnetConstructedDataAlarmValues
func NewBACnetConstructedDataAlarmValues(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, alarmValues []BACnetLifeSafetyStateTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAlarmValues {
	_result := &_BACnetConstructedDataAlarmValues{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		AlarmValues:                   alarmValues,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAlarmValuesBuilder is a builder for BACnetConstructedDataAlarmValues
type BACnetConstructedDataAlarmValuesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(alarmValues []BACnetLifeSafetyStateTagged) BACnetConstructedDataAlarmValuesBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAlarmValuesBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAlarmValuesBuilder
	// WithAlarmValues adds AlarmValues (property field)
	WithAlarmValues(...BACnetLifeSafetyStateTagged) BACnetConstructedDataAlarmValuesBuilder
	// Build builds the BACnetConstructedDataAlarmValues or returns an error if something is wrong
	Build() (BACnetConstructedDataAlarmValues, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAlarmValues
}

// NewBACnetConstructedDataAlarmValuesBuilder() creates a BACnetConstructedDataAlarmValuesBuilder
func NewBACnetConstructedDataAlarmValuesBuilder() BACnetConstructedDataAlarmValuesBuilder {
	return &_BACnetConstructedDataAlarmValuesBuilder{_BACnetConstructedDataAlarmValues: new(_BACnetConstructedDataAlarmValues)}
}

type _BACnetConstructedDataAlarmValuesBuilder struct {
	*_BACnetConstructedDataAlarmValues

	err *utils.MultiError
}

var _ (BACnetConstructedDataAlarmValuesBuilder) = (*_BACnetConstructedDataAlarmValuesBuilder)(nil)

func (m *_BACnetConstructedDataAlarmValuesBuilder) WithMandatoryFields(alarmValues []BACnetLifeSafetyStateTagged) BACnetConstructedDataAlarmValuesBuilder {
	return m.WithAlarmValues(alarmValues...)
}

func (m *_BACnetConstructedDataAlarmValuesBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataAlarmValuesBuilder {
	m.NumberOfDataElements = numberOfDataElements
	return m
}

func (m *_BACnetConstructedDataAlarmValuesBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataAlarmValuesBuilder {
	builder := builderSupplier(m.NumberOfDataElements.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.NumberOfDataElements, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataAlarmValuesBuilder) WithAlarmValues(alarmValues ...BACnetLifeSafetyStateTagged) BACnetConstructedDataAlarmValuesBuilder {
	m.AlarmValues = alarmValues
	return m
}

func (m *_BACnetConstructedDataAlarmValuesBuilder) Build() (BACnetConstructedDataAlarmValues, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataAlarmValues.deepCopy(), nil
}

func (m *_BACnetConstructedDataAlarmValuesBuilder) MustBuild() BACnetConstructedDataAlarmValues {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataAlarmValuesBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataAlarmValuesBuilder()
}

// CreateBACnetConstructedDataAlarmValuesBuilder creates a BACnetConstructedDataAlarmValuesBuilder
func (m *_BACnetConstructedDataAlarmValues) CreateBACnetConstructedDataAlarmValuesBuilder() BACnetConstructedDataAlarmValuesBuilder {
	if m == nil {
		return NewBACnetConstructedDataAlarmValuesBuilder()
	}
	return &_BACnetConstructedDataAlarmValuesBuilder{_BACnetConstructedDataAlarmValues: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAlarmValues) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAlarmValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALARM_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAlarmValues) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAlarmValues) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataAlarmValues) GetAlarmValues() []BACnetLifeSafetyStateTagged {
	return m.AlarmValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAlarmValues) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAlarmValues(structType any) BACnetConstructedDataAlarmValues {
	if casted, ok := structType.(BACnetConstructedDataAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAlarmValues) GetTypeName() string {
	return "BACnetConstructedDataAlarmValues"
}

func (m *_BACnetConstructedDataAlarmValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.AlarmValues) > 0 {
		for _, element := range m.AlarmValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAlarmValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAlarmValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAlarmValues BACnetConstructedDataAlarmValues, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	alarmValues, err := ReadTerminatedArrayField[BACnetLifeSafetyStateTagged](ctx, "alarmValues", ReadComplex[BACnetLifeSafetyStateTagged](BACnetLifeSafetyStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alarmValues' field"))
	}
	m.AlarmValues = alarmValues

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAlarmValues")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAlarmValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAlarmValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAlarmValues")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "alarmValues", m.GetAlarmValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'alarmValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAlarmValues")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAlarmValues) IsBACnetConstructedDataAlarmValues() {}

func (m *_BACnetConstructedDataAlarmValues) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAlarmValues) deepCopy() *_BACnetConstructedDataAlarmValues {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAlarmValuesCopy := &_BACnetConstructedDataAlarmValues{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetLifeSafetyStateTagged, BACnetLifeSafetyStateTagged](m.AlarmValues),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAlarmValuesCopy
}

func (m *_BACnetConstructedDataAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
