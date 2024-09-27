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

// BACnetConstructedDataShedLevels is the corresponding interface of BACnetConstructedDataShedLevels
type BACnetConstructedDataShedLevels interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetShedLevels returns ShedLevels (property field)
	GetShedLevels() []BACnetApplicationTagUnsignedInteger
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataShedLevels is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataShedLevels()
	// CreateBuilder creates a BACnetConstructedDataShedLevelsBuilder
	CreateBACnetConstructedDataShedLevelsBuilder() BACnetConstructedDataShedLevelsBuilder
}

// _BACnetConstructedDataShedLevels is the data-structure of this message
type _BACnetConstructedDataShedLevels struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	ShedLevels           []BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataShedLevels = (*_BACnetConstructedDataShedLevels)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataShedLevels)(nil)

// NewBACnetConstructedDataShedLevels factory function for _BACnetConstructedDataShedLevels
func NewBACnetConstructedDataShedLevels(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, shedLevels []BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataShedLevels {
	_result := &_BACnetConstructedDataShedLevels{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		ShedLevels:                    shedLevels,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataShedLevelsBuilder is a builder for BACnetConstructedDataShedLevels
type BACnetConstructedDataShedLevelsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(shedLevels []BACnetApplicationTagUnsignedInteger) BACnetConstructedDataShedLevelsBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataShedLevelsBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataShedLevelsBuilder
	// WithShedLevels adds ShedLevels (property field)
	WithShedLevels(...BACnetApplicationTagUnsignedInteger) BACnetConstructedDataShedLevelsBuilder
	// Build builds the BACnetConstructedDataShedLevels or returns an error if something is wrong
	Build() (BACnetConstructedDataShedLevels, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataShedLevels
}

// NewBACnetConstructedDataShedLevelsBuilder() creates a BACnetConstructedDataShedLevelsBuilder
func NewBACnetConstructedDataShedLevelsBuilder() BACnetConstructedDataShedLevelsBuilder {
	return &_BACnetConstructedDataShedLevelsBuilder{_BACnetConstructedDataShedLevels: new(_BACnetConstructedDataShedLevels)}
}

type _BACnetConstructedDataShedLevelsBuilder struct {
	*_BACnetConstructedDataShedLevels

	err *utils.MultiError
}

var _ (BACnetConstructedDataShedLevelsBuilder) = (*_BACnetConstructedDataShedLevelsBuilder)(nil)

func (m *_BACnetConstructedDataShedLevelsBuilder) WithMandatoryFields(shedLevels []BACnetApplicationTagUnsignedInteger) BACnetConstructedDataShedLevelsBuilder {
	return m.WithShedLevels(shedLevels...)
}

func (m *_BACnetConstructedDataShedLevelsBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataShedLevelsBuilder {
	m.NumberOfDataElements = numberOfDataElements
	return m
}

func (m *_BACnetConstructedDataShedLevelsBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataShedLevelsBuilder {
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

func (m *_BACnetConstructedDataShedLevelsBuilder) WithShedLevels(shedLevels ...BACnetApplicationTagUnsignedInteger) BACnetConstructedDataShedLevelsBuilder {
	m.ShedLevels = shedLevels
	return m
}

func (m *_BACnetConstructedDataShedLevelsBuilder) Build() (BACnetConstructedDataShedLevels, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataShedLevels.deepCopy(), nil
}

func (m *_BACnetConstructedDataShedLevelsBuilder) MustBuild() BACnetConstructedDataShedLevels {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataShedLevelsBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataShedLevelsBuilder()
}

// CreateBACnetConstructedDataShedLevelsBuilder creates a BACnetConstructedDataShedLevelsBuilder
func (m *_BACnetConstructedDataShedLevels) CreateBACnetConstructedDataShedLevelsBuilder() BACnetConstructedDataShedLevelsBuilder {
	if m == nil {
		return NewBACnetConstructedDataShedLevelsBuilder()
	}
	return &_BACnetConstructedDataShedLevelsBuilder{_BACnetConstructedDataShedLevels: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataShedLevels) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataShedLevels) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SHED_LEVELS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataShedLevels) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataShedLevels) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataShedLevels) GetShedLevels() []BACnetApplicationTagUnsignedInteger {
	return m.ShedLevels
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataShedLevels) GetZero() uint64 {
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
func CastBACnetConstructedDataShedLevels(structType any) BACnetConstructedDataShedLevels {
	if casted, ok := structType.(BACnetConstructedDataShedLevels); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataShedLevels); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataShedLevels) GetTypeName() string {
	return "BACnetConstructedDataShedLevels"
}

func (m *_BACnetConstructedDataShedLevels) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.ShedLevels) > 0 {
		for _, element := range m.ShedLevels {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataShedLevels) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataShedLevels) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataShedLevels BACnetConstructedDataShedLevels, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataShedLevels"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataShedLevels")
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

	shedLevels, err := ReadTerminatedArrayField[BACnetApplicationTagUnsignedInteger](ctx, "shedLevels", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'shedLevels' field"))
	}
	m.ShedLevels = shedLevels

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataShedLevels"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataShedLevels")
	}

	return m, nil
}

func (m *_BACnetConstructedDataShedLevels) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataShedLevels) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataShedLevels"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataShedLevels")
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

		if err := WriteComplexTypeArrayField(ctx, "shedLevels", m.GetShedLevels(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'shedLevels' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataShedLevels"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataShedLevels")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataShedLevels) IsBACnetConstructedDataShedLevels() {}

func (m *_BACnetConstructedDataShedLevels) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataShedLevels) deepCopy() *_BACnetConstructedDataShedLevels {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataShedLevelsCopy := &_BACnetConstructedDataShedLevels{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetApplicationTagUnsignedInteger, BACnetApplicationTagUnsignedInteger](m.ShedLevels),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataShedLevelsCopy
}

func (m *_BACnetConstructedDataShedLevels) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
