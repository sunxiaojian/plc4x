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

// BACnetEventParameterExtended is the corresponding interface of BACnetEventParameterExtended
type BACnetEventParameterExtended interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetExtendedEventType returns ExtendedEventType (property field)
	GetExtendedEventType() BACnetContextTagUnsignedInteger
	// GetParameters returns Parameters (property field)
	GetParameters() BACnetEventParameterExtendedParameters
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterExtended is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterExtended()
	// CreateBuilder creates a BACnetEventParameterExtendedBuilder
	CreateBACnetEventParameterExtendedBuilder() BACnetEventParameterExtendedBuilder
}

// _BACnetEventParameterExtended is the data-structure of this message
type _BACnetEventParameterExtended struct {
	BACnetEventParameterContract
	OpeningTag        BACnetOpeningTag
	VendorId          BACnetVendorIdTagged
	ExtendedEventType BACnetContextTagUnsignedInteger
	Parameters        BACnetEventParameterExtendedParameters
	ClosingTag        BACnetClosingTag
}

var _ BACnetEventParameterExtended = (*_BACnetEventParameterExtended)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterExtended)(nil)

// NewBACnetEventParameterExtended factory function for _BACnetEventParameterExtended
func NewBACnetEventParameterExtended(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, vendorId BACnetVendorIdTagged, extendedEventType BACnetContextTagUnsignedInteger, parameters BACnetEventParameterExtendedParameters, closingTag BACnetClosingTag) *_BACnetEventParameterExtended {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterExtended must not be nil")
	}
	if vendorId == nil {
		panic("vendorId of type BACnetVendorIdTagged for BACnetEventParameterExtended must not be nil")
	}
	if extendedEventType == nil {
		panic("extendedEventType of type BACnetContextTagUnsignedInteger for BACnetEventParameterExtended must not be nil")
	}
	if parameters == nil {
		panic("parameters of type BACnetEventParameterExtendedParameters for BACnetEventParameterExtended must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterExtended must not be nil")
	}
	_result := &_BACnetEventParameterExtended{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		VendorId:                     vendorId,
		ExtendedEventType:            extendedEventType,
		Parameters:                   parameters,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterExtendedBuilder is a builder for BACnetEventParameterExtended
type BACnetEventParameterExtendedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, vendorId BACnetVendorIdTagged, extendedEventType BACnetContextTagUnsignedInteger, parameters BACnetEventParameterExtendedParameters, closingTag BACnetClosingTag) BACnetEventParameterExtendedBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterExtendedBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterExtendedBuilder
	// WithVendorId adds VendorId (property field)
	WithVendorId(BACnetVendorIdTagged) BACnetEventParameterExtendedBuilder
	// WithVendorIdBuilder adds VendorId (property field) which is build by the builder
	WithVendorIdBuilder(func(BACnetVendorIdTaggedBuilder) BACnetVendorIdTaggedBuilder) BACnetEventParameterExtendedBuilder
	// WithExtendedEventType adds ExtendedEventType (property field)
	WithExtendedEventType(BACnetContextTagUnsignedInteger) BACnetEventParameterExtendedBuilder
	// WithExtendedEventTypeBuilder adds ExtendedEventType (property field) which is build by the builder
	WithExtendedEventTypeBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterExtendedBuilder
	// WithParameters adds Parameters (property field)
	WithParameters(BACnetEventParameterExtendedParameters) BACnetEventParameterExtendedBuilder
	// WithParametersBuilder adds Parameters (property field) which is build by the builder
	WithParametersBuilder(func(BACnetEventParameterExtendedParametersBuilder) BACnetEventParameterExtendedParametersBuilder) BACnetEventParameterExtendedBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterExtendedBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterExtendedBuilder
	// Build builds the BACnetEventParameterExtended or returns an error if something is wrong
	Build() (BACnetEventParameterExtended, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterExtended
}

// NewBACnetEventParameterExtendedBuilder() creates a BACnetEventParameterExtendedBuilder
func NewBACnetEventParameterExtendedBuilder() BACnetEventParameterExtendedBuilder {
	return &_BACnetEventParameterExtendedBuilder{_BACnetEventParameterExtended: new(_BACnetEventParameterExtended)}
}

type _BACnetEventParameterExtendedBuilder struct {
	*_BACnetEventParameterExtended

	err *utils.MultiError
}

var _ (BACnetEventParameterExtendedBuilder) = (*_BACnetEventParameterExtendedBuilder)(nil)

func (m *_BACnetEventParameterExtendedBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, vendorId BACnetVendorIdTagged, extendedEventType BACnetContextTagUnsignedInteger, parameters BACnetEventParameterExtendedParameters, closingTag BACnetClosingTag) BACnetEventParameterExtendedBuilder {
	return m.WithOpeningTag(openingTag).WithVendorId(vendorId).WithExtendedEventType(extendedEventType).WithParameters(parameters).WithClosingTag(closingTag)
}

func (m *_BACnetEventParameterExtendedBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterExtendedBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetEventParameterExtendedBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterExtendedBuilder {
	builder := builderSupplier(m.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	m.OpeningTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return m
}

func (m *_BACnetEventParameterExtendedBuilder) WithVendorId(vendorId BACnetVendorIdTagged) BACnetEventParameterExtendedBuilder {
	m.VendorId = vendorId
	return m
}

func (m *_BACnetEventParameterExtendedBuilder) WithVendorIdBuilder(builderSupplier func(BACnetVendorIdTaggedBuilder) BACnetVendorIdTaggedBuilder) BACnetEventParameterExtendedBuilder {
	builder := builderSupplier(m.VendorId.CreateBACnetVendorIdTaggedBuilder())
	var err error
	m.VendorId, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetVendorIdTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetEventParameterExtendedBuilder) WithExtendedEventType(extendedEventType BACnetContextTagUnsignedInteger) BACnetEventParameterExtendedBuilder {
	m.ExtendedEventType = extendedEventType
	return m
}

func (m *_BACnetEventParameterExtendedBuilder) WithExtendedEventTypeBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterExtendedBuilder {
	builder := builderSupplier(m.ExtendedEventType.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.ExtendedEventType, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetEventParameterExtendedBuilder) WithParameters(parameters BACnetEventParameterExtendedParameters) BACnetEventParameterExtendedBuilder {
	m.Parameters = parameters
	return m
}

func (m *_BACnetEventParameterExtendedBuilder) WithParametersBuilder(builderSupplier func(BACnetEventParameterExtendedParametersBuilder) BACnetEventParameterExtendedParametersBuilder) BACnetEventParameterExtendedBuilder {
	builder := builderSupplier(m.Parameters.CreateBACnetEventParameterExtendedParametersBuilder())
	var err error
	m.Parameters, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetEventParameterExtendedParametersBuilder failed"))
	}
	return m
}

func (m *_BACnetEventParameterExtendedBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterExtendedBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetEventParameterExtendedBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterExtendedBuilder {
	builder := builderSupplier(m.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	m.ClosingTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return m
}

func (m *_BACnetEventParameterExtendedBuilder) Build() (BACnetEventParameterExtended, error) {
	if m.OpeningTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if m.VendorId == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'vendorId' not set"))
	}
	if m.ExtendedEventType == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'extendedEventType' not set"))
	}
	if m.Parameters == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'parameters' not set"))
	}
	if m.ClosingTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetEventParameterExtended.deepCopy(), nil
}

func (m *_BACnetEventParameterExtendedBuilder) MustBuild() BACnetEventParameterExtended {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetEventParameterExtendedBuilder) DeepCopy() any {
	return m.CreateBACnetEventParameterExtendedBuilder()
}

// CreateBACnetEventParameterExtendedBuilder creates a BACnetEventParameterExtendedBuilder
func (m *_BACnetEventParameterExtended) CreateBACnetEventParameterExtendedBuilder() BACnetEventParameterExtendedBuilder {
	if m == nil {
		return NewBACnetEventParameterExtendedBuilder()
	}
	return &_BACnetEventParameterExtendedBuilder{_BACnetEventParameterExtended: m.deepCopy()}
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

func (m *_BACnetEventParameterExtended) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterExtended) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterExtended) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_BACnetEventParameterExtended) GetExtendedEventType() BACnetContextTagUnsignedInteger {
	return m.ExtendedEventType
}

func (m *_BACnetEventParameterExtended) GetParameters() BACnetEventParameterExtendedParameters {
	return m.Parameters
}

func (m *_BACnetEventParameterExtended) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterExtended(structType any) BACnetEventParameterExtended {
	if casted, ok := structType.(BACnetEventParameterExtended); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterExtended); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterExtended) GetTypeName() string {
	return "BACnetEventParameterExtended"
}

func (m *_BACnetEventParameterExtended) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits(ctx)

	// Simple field (extendedEventType)
	lengthInBits += m.ExtendedEventType.GetLengthInBits(ctx)

	// Simple field (parameters)
	lengthInBits += m.Parameters.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterExtended) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterExtended) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterExtended BACnetEventParameterExtended, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterExtended"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterExtended")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(9))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	vendorId, err := ReadSimpleField[BACnetVendorIdTagged](ctx, "vendorId", ReadComplex[BACnetVendorIdTagged](BACnetVendorIdTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorId' field"))
	}
	m.VendorId = vendorId

	extendedEventType, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "extendedEventType", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extendedEventType' field"))
	}
	m.ExtendedEventType = extendedEventType

	parameters, err := ReadSimpleField[BACnetEventParameterExtendedParameters](ctx, "parameters", ReadComplex[BACnetEventParameterExtendedParameters](BACnetEventParameterExtendedParametersParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameters' field"))
	}
	m.Parameters = parameters

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(9))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterExtended"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterExtended")
	}

	return m, nil
}

func (m *_BACnetEventParameterExtended) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterExtended) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterExtended"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterExtended")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetVendorIdTagged](ctx, "vendorId", m.GetVendorId(), WriteComplex[BACnetVendorIdTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vendorId' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "extendedEventType", m.GetExtendedEventType(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'extendedEventType' field")
		}

		if err := WriteSimpleField[BACnetEventParameterExtendedParameters](ctx, "parameters", m.GetParameters(), WriteComplex[BACnetEventParameterExtendedParameters](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'parameters' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterExtended"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterExtended")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterExtended) IsBACnetEventParameterExtended() {}

func (m *_BACnetEventParameterExtended) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterExtended) deepCopy() *_BACnetEventParameterExtended {
	if m == nil {
		return nil
	}
	_BACnetEventParameterExtendedCopy := &_BACnetEventParameterExtended{
		m.BACnetEventParameterContract.(*_BACnetEventParameter).deepCopy(),
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.VendorId.DeepCopy().(BACnetVendorIdTagged),
		m.ExtendedEventType.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.Parameters.DeepCopy().(BACnetEventParameterExtendedParameters),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = m
	return _BACnetEventParameterExtendedCopy
}

func (m *_BACnetEventParameterExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
