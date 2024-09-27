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

// BACnetConstructedDataSubordinateTags is the corresponding interface of BACnetConstructedDataSubordinateTags
type BACnetConstructedDataSubordinateTags interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetSubordinateList returns SubordinateList (property field)
	GetSubordinateList() []BACnetNameValueCollection
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataSubordinateTags is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataSubordinateTags()
	// CreateBuilder creates a BACnetConstructedDataSubordinateTagsBuilder
	CreateBACnetConstructedDataSubordinateTagsBuilder() BACnetConstructedDataSubordinateTagsBuilder
}

// _BACnetConstructedDataSubordinateTags is the data-structure of this message
type _BACnetConstructedDataSubordinateTags struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	SubordinateList      []BACnetNameValueCollection
}

var _ BACnetConstructedDataSubordinateTags = (*_BACnetConstructedDataSubordinateTags)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataSubordinateTags)(nil)

// NewBACnetConstructedDataSubordinateTags factory function for _BACnetConstructedDataSubordinateTags
func NewBACnetConstructedDataSubordinateTags(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, subordinateList []BACnetNameValueCollection, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSubordinateTags {
	_result := &_BACnetConstructedDataSubordinateTags{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		SubordinateList:               subordinateList,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataSubordinateTagsBuilder is a builder for BACnetConstructedDataSubordinateTags
type BACnetConstructedDataSubordinateTagsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(subordinateList []BACnetNameValueCollection) BACnetConstructedDataSubordinateTagsBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSubordinateTagsBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataSubordinateTagsBuilder
	// WithSubordinateList adds SubordinateList (property field)
	WithSubordinateList(...BACnetNameValueCollection) BACnetConstructedDataSubordinateTagsBuilder
	// Build builds the BACnetConstructedDataSubordinateTags or returns an error if something is wrong
	Build() (BACnetConstructedDataSubordinateTags, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataSubordinateTags
}

// NewBACnetConstructedDataSubordinateTagsBuilder() creates a BACnetConstructedDataSubordinateTagsBuilder
func NewBACnetConstructedDataSubordinateTagsBuilder() BACnetConstructedDataSubordinateTagsBuilder {
	return &_BACnetConstructedDataSubordinateTagsBuilder{_BACnetConstructedDataSubordinateTags: new(_BACnetConstructedDataSubordinateTags)}
}

type _BACnetConstructedDataSubordinateTagsBuilder struct {
	*_BACnetConstructedDataSubordinateTags

	err *utils.MultiError
}

var _ (BACnetConstructedDataSubordinateTagsBuilder) = (*_BACnetConstructedDataSubordinateTagsBuilder)(nil)

func (m *_BACnetConstructedDataSubordinateTagsBuilder) WithMandatoryFields(subordinateList []BACnetNameValueCollection) BACnetConstructedDataSubordinateTagsBuilder {
	return m.WithSubordinateList(subordinateList...)
}

func (m *_BACnetConstructedDataSubordinateTagsBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataSubordinateTagsBuilder {
	m.NumberOfDataElements = numberOfDataElements
	return m
}

func (m *_BACnetConstructedDataSubordinateTagsBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataSubordinateTagsBuilder {
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

func (m *_BACnetConstructedDataSubordinateTagsBuilder) WithSubordinateList(subordinateList ...BACnetNameValueCollection) BACnetConstructedDataSubordinateTagsBuilder {
	m.SubordinateList = subordinateList
	return m
}

func (m *_BACnetConstructedDataSubordinateTagsBuilder) Build() (BACnetConstructedDataSubordinateTags, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataSubordinateTags.deepCopy(), nil
}

func (m *_BACnetConstructedDataSubordinateTagsBuilder) MustBuild() BACnetConstructedDataSubordinateTags {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataSubordinateTagsBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataSubordinateTagsBuilder()
}

// CreateBACnetConstructedDataSubordinateTagsBuilder creates a BACnetConstructedDataSubordinateTagsBuilder
func (m *_BACnetConstructedDataSubordinateTags) CreateBACnetConstructedDataSubordinateTagsBuilder() BACnetConstructedDataSubordinateTagsBuilder {
	if m == nil {
		return NewBACnetConstructedDataSubordinateTagsBuilder()
	}
	return &_BACnetConstructedDataSubordinateTagsBuilder{_BACnetConstructedDataSubordinateTags: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSubordinateTags) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSubordinateTags) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUBORDINATE_TAGS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSubordinateTags) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSubordinateTags) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataSubordinateTags) GetSubordinateList() []BACnetNameValueCollection {
	return m.SubordinateList
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSubordinateTags) GetZero() uint64 {
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
func CastBACnetConstructedDataSubordinateTags(structType any) BACnetConstructedDataSubordinateTags {
	if casted, ok := structType.(BACnetConstructedDataSubordinateTags); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSubordinateTags); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSubordinateTags) GetTypeName() string {
	return "BACnetConstructedDataSubordinateTags"
}

func (m *_BACnetConstructedDataSubordinateTags) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.SubordinateList) > 0 {
		for _, element := range m.SubordinateList {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataSubordinateTags) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataSubordinateTags) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataSubordinateTags BACnetConstructedDataSubordinateTags, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSubordinateTags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSubordinateTags")
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

	subordinateList, err := ReadTerminatedArrayField[BACnetNameValueCollection](ctx, "subordinateList", ReadComplex[BACnetNameValueCollection](BACnetNameValueCollectionParseWithBufferProducer((uint8)(uint8(0))), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subordinateList' field"))
	}
	m.SubordinateList = subordinateList

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSubordinateTags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSubordinateTags")
	}

	return m, nil
}

func (m *_BACnetConstructedDataSubordinateTags) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSubordinateTags) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSubordinateTags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSubordinateTags")
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

		if err := WriteComplexTypeArrayField(ctx, "subordinateList", m.GetSubordinateList(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'subordinateList' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSubordinateTags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSubordinateTags")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSubordinateTags) IsBACnetConstructedDataSubordinateTags() {}

func (m *_BACnetConstructedDataSubordinateTags) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataSubordinateTags) deepCopy() *_BACnetConstructedDataSubordinateTags {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataSubordinateTagsCopy := &_BACnetConstructedDataSubordinateTags{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.NumberOfDataElements.DeepCopy().(BACnetApplicationTagUnsignedInteger),
		utils.DeepCopySlice[BACnetNameValueCollection, BACnetNameValueCollection](m.SubordinateList),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataSubordinateTagsCopy
}

func (m *_BACnetConstructedDataSubordinateTags) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
