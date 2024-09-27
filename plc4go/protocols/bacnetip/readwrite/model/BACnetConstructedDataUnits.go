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

// BACnetConstructedDataUnits is the corresponding interface of BACnetConstructedDataUnits
type BACnetConstructedDataUnits interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetUnits returns Units (property field)
	GetUnits() BACnetEngineeringUnitsTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEngineeringUnitsTagged
	// IsBACnetConstructedDataUnits is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataUnits()
	// CreateBuilder creates a BACnetConstructedDataUnitsBuilder
	CreateBACnetConstructedDataUnitsBuilder() BACnetConstructedDataUnitsBuilder
}

// _BACnetConstructedDataUnits is the data-structure of this message
type _BACnetConstructedDataUnits struct {
	BACnetConstructedDataContract
	Units BACnetEngineeringUnitsTagged
}

var _ BACnetConstructedDataUnits = (*_BACnetConstructedDataUnits)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataUnits)(nil)

// NewBACnetConstructedDataUnits factory function for _BACnetConstructedDataUnits
func NewBACnetConstructedDataUnits(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, units BACnetEngineeringUnitsTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUnits {
	if units == nil {
		panic("units of type BACnetEngineeringUnitsTagged for BACnetConstructedDataUnits must not be nil")
	}
	_result := &_BACnetConstructedDataUnits{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Units:                         units,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataUnitsBuilder is a builder for BACnetConstructedDataUnits
type BACnetConstructedDataUnitsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(units BACnetEngineeringUnitsTagged) BACnetConstructedDataUnitsBuilder
	// WithUnits adds Units (property field)
	WithUnits(BACnetEngineeringUnitsTagged) BACnetConstructedDataUnitsBuilder
	// WithUnitsBuilder adds Units (property field) which is build by the builder
	WithUnitsBuilder(func(BACnetEngineeringUnitsTaggedBuilder) BACnetEngineeringUnitsTaggedBuilder) BACnetConstructedDataUnitsBuilder
	// Build builds the BACnetConstructedDataUnits or returns an error if something is wrong
	Build() (BACnetConstructedDataUnits, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataUnits
}

// NewBACnetConstructedDataUnitsBuilder() creates a BACnetConstructedDataUnitsBuilder
func NewBACnetConstructedDataUnitsBuilder() BACnetConstructedDataUnitsBuilder {
	return &_BACnetConstructedDataUnitsBuilder{_BACnetConstructedDataUnits: new(_BACnetConstructedDataUnits)}
}

type _BACnetConstructedDataUnitsBuilder struct {
	*_BACnetConstructedDataUnits

	err *utils.MultiError
}

var _ (BACnetConstructedDataUnitsBuilder) = (*_BACnetConstructedDataUnitsBuilder)(nil)

func (m *_BACnetConstructedDataUnitsBuilder) WithMandatoryFields(units BACnetEngineeringUnitsTagged) BACnetConstructedDataUnitsBuilder {
	return m.WithUnits(units)
}

func (m *_BACnetConstructedDataUnitsBuilder) WithUnits(units BACnetEngineeringUnitsTagged) BACnetConstructedDataUnitsBuilder {
	m.Units = units
	return m
}

func (m *_BACnetConstructedDataUnitsBuilder) WithUnitsBuilder(builderSupplier func(BACnetEngineeringUnitsTaggedBuilder) BACnetEngineeringUnitsTaggedBuilder) BACnetConstructedDataUnitsBuilder {
	builder := builderSupplier(m.Units.CreateBACnetEngineeringUnitsTaggedBuilder())
	var err error
	m.Units, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetEngineeringUnitsTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataUnitsBuilder) Build() (BACnetConstructedDataUnits, error) {
	if m.Units == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'units' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataUnits.deepCopy(), nil
}

func (m *_BACnetConstructedDataUnitsBuilder) MustBuild() BACnetConstructedDataUnits {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataUnitsBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataUnitsBuilder()
}

// CreateBACnetConstructedDataUnitsBuilder creates a BACnetConstructedDataUnitsBuilder
func (m *_BACnetConstructedDataUnits) CreateBACnetConstructedDataUnitsBuilder() BACnetConstructedDataUnitsBuilder {
	if m == nil {
		return NewBACnetConstructedDataUnitsBuilder()
	}
	return &_BACnetConstructedDataUnitsBuilder{_BACnetConstructedDataUnits: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUnits) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUnits) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_UNITS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUnits) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUnits) GetUnits() BACnetEngineeringUnitsTagged {
	return m.Units
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataUnits) GetActualValue() BACnetEngineeringUnitsTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetEngineeringUnitsTagged(m.GetUnits())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUnits(structType any) BACnetConstructedDataUnits {
	if casted, ok := structType.(BACnetConstructedDataUnits); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUnits); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUnits) GetTypeName() string {
	return "BACnetConstructedDataUnits"
}

func (m *_BACnetConstructedDataUnits) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (units)
	lengthInBits += m.Units.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataUnits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataUnits) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataUnits BACnetConstructedDataUnits, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUnits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUnits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	units, err := ReadSimpleField[BACnetEngineeringUnitsTagged](ctx, "units", ReadComplex[BACnetEngineeringUnitsTagged](BACnetEngineeringUnitsTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'units' field"))
	}
	m.Units = units

	actualValue, err := ReadVirtualField[BACnetEngineeringUnitsTagged](ctx, "actualValue", (*BACnetEngineeringUnitsTagged)(nil), units)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUnits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUnits")
	}

	return m, nil
}

func (m *_BACnetConstructedDataUnits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataUnits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUnits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUnits")
		}

		if err := WriteSimpleField[BACnetEngineeringUnitsTagged](ctx, "units", m.GetUnits(), WriteComplex[BACnetEngineeringUnitsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'units' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUnits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUnits")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUnits) IsBACnetConstructedDataUnits() {}

func (m *_BACnetConstructedDataUnits) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataUnits) deepCopy() *_BACnetConstructedDataUnits {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataUnitsCopy := &_BACnetConstructedDataUnits{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.Units.DeepCopy().(BACnetEngineeringUnitsTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataUnitsCopy
}

func (m *_BACnetConstructedDataUnits) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
