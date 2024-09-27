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

// BACnetConstructedDataSegmentationSupported is the corresponding interface of BACnetConstructedDataSegmentationSupported
type BACnetConstructedDataSegmentationSupported interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetSegmentationSupported returns SegmentationSupported (property field)
	GetSegmentationSupported() BACnetSegmentationTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetSegmentationTagged
	// IsBACnetConstructedDataSegmentationSupported is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataSegmentationSupported()
	// CreateBuilder creates a BACnetConstructedDataSegmentationSupportedBuilder
	CreateBACnetConstructedDataSegmentationSupportedBuilder() BACnetConstructedDataSegmentationSupportedBuilder
}

// _BACnetConstructedDataSegmentationSupported is the data-structure of this message
type _BACnetConstructedDataSegmentationSupported struct {
	BACnetConstructedDataContract
	SegmentationSupported BACnetSegmentationTagged
}

var _ BACnetConstructedDataSegmentationSupported = (*_BACnetConstructedDataSegmentationSupported)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataSegmentationSupported)(nil)

// NewBACnetConstructedDataSegmentationSupported factory function for _BACnetConstructedDataSegmentationSupported
func NewBACnetConstructedDataSegmentationSupported(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, segmentationSupported BACnetSegmentationTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSegmentationSupported {
	if segmentationSupported == nil {
		panic("segmentationSupported of type BACnetSegmentationTagged for BACnetConstructedDataSegmentationSupported must not be nil")
	}
	_result := &_BACnetConstructedDataSegmentationSupported{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		SegmentationSupported:         segmentationSupported,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataSegmentationSupportedBuilder is a builder for BACnetConstructedDataSegmentationSupported
type BACnetConstructedDataSegmentationSupportedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(segmentationSupported BACnetSegmentationTagged) BACnetConstructedDataSegmentationSupportedBuilder
	// WithSegmentationSupported adds SegmentationSupported (property field)
	WithSegmentationSupported(BACnetSegmentationTagged) BACnetConstructedDataSegmentationSupportedBuilder
	// WithSegmentationSupportedBuilder adds SegmentationSupported (property field) which is build by the builder
	WithSegmentationSupportedBuilder(func(BACnetSegmentationTaggedBuilder) BACnetSegmentationTaggedBuilder) BACnetConstructedDataSegmentationSupportedBuilder
	// Build builds the BACnetConstructedDataSegmentationSupported or returns an error if something is wrong
	Build() (BACnetConstructedDataSegmentationSupported, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataSegmentationSupported
}

// NewBACnetConstructedDataSegmentationSupportedBuilder() creates a BACnetConstructedDataSegmentationSupportedBuilder
func NewBACnetConstructedDataSegmentationSupportedBuilder() BACnetConstructedDataSegmentationSupportedBuilder {
	return &_BACnetConstructedDataSegmentationSupportedBuilder{_BACnetConstructedDataSegmentationSupported: new(_BACnetConstructedDataSegmentationSupported)}
}

type _BACnetConstructedDataSegmentationSupportedBuilder struct {
	*_BACnetConstructedDataSegmentationSupported

	err *utils.MultiError
}

var _ (BACnetConstructedDataSegmentationSupportedBuilder) = (*_BACnetConstructedDataSegmentationSupportedBuilder)(nil)

func (m *_BACnetConstructedDataSegmentationSupportedBuilder) WithMandatoryFields(segmentationSupported BACnetSegmentationTagged) BACnetConstructedDataSegmentationSupportedBuilder {
	return m.WithSegmentationSupported(segmentationSupported)
}

func (m *_BACnetConstructedDataSegmentationSupportedBuilder) WithSegmentationSupported(segmentationSupported BACnetSegmentationTagged) BACnetConstructedDataSegmentationSupportedBuilder {
	m.SegmentationSupported = segmentationSupported
	return m
}

func (m *_BACnetConstructedDataSegmentationSupportedBuilder) WithSegmentationSupportedBuilder(builderSupplier func(BACnetSegmentationTaggedBuilder) BACnetSegmentationTaggedBuilder) BACnetConstructedDataSegmentationSupportedBuilder {
	builder := builderSupplier(m.SegmentationSupported.CreateBACnetSegmentationTaggedBuilder())
	var err error
	m.SegmentationSupported, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetSegmentationTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataSegmentationSupportedBuilder) Build() (BACnetConstructedDataSegmentationSupported, error) {
	if m.SegmentationSupported == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'segmentationSupported' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataSegmentationSupported.deepCopy(), nil
}

func (m *_BACnetConstructedDataSegmentationSupportedBuilder) MustBuild() BACnetConstructedDataSegmentationSupported {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataSegmentationSupportedBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataSegmentationSupportedBuilder()
}

// CreateBACnetConstructedDataSegmentationSupportedBuilder creates a BACnetConstructedDataSegmentationSupportedBuilder
func (m *_BACnetConstructedDataSegmentationSupported) CreateBACnetConstructedDataSegmentationSupportedBuilder() BACnetConstructedDataSegmentationSupportedBuilder {
	if m == nil {
		return NewBACnetConstructedDataSegmentationSupportedBuilder()
	}
	return &_BACnetConstructedDataSegmentationSupportedBuilder{_BACnetConstructedDataSegmentationSupported: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSegmentationSupported) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSegmentationSupported) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SEGMENTATION_SUPPORTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSegmentationSupported) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSegmentationSupported) GetSegmentationSupported() BACnetSegmentationTagged {
	return m.SegmentationSupported
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSegmentationSupported) GetActualValue() BACnetSegmentationTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetSegmentationTagged(m.GetSegmentationSupported())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSegmentationSupported(structType any) BACnetConstructedDataSegmentationSupported {
	if casted, ok := structType.(BACnetConstructedDataSegmentationSupported); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSegmentationSupported); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSegmentationSupported) GetTypeName() string {
	return "BACnetConstructedDataSegmentationSupported"
}

func (m *_BACnetConstructedDataSegmentationSupported) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (segmentationSupported)
	lengthInBits += m.SegmentationSupported.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSegmentationSupported) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataSegmentationSupported) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataSegmentationSupported BACnetConstructedDataSegmentationSupported, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSegmentationSupported"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSegmentationSupported")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	segmentationSupported, err := ReadSimpleField[BACnetSegmentationTagged](ctx, "segmentationSupported", ReadComplex[BACnetSegmentationTagged](BACnetSegmentationTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentationSupported' field"))
	}
	m.SegmentationSupported = segmentationSupported

	actualValue, err := ReadVirtualField[BACnetSegmentationTagged](ctx, "actualValue", (*BACnetSegmentationTagged)(nil), segmentationSupported)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSegmentationSupported"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSegmentationSupported")
	}

	return m, nil
}

func (m *_BACnetConstructedDataSegmentationSupported) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSegmentationSupported) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSegmentationSupported"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSegmentationSupported")
		}

		if err := WriteSimpleField[BACnetSegmentationTagged](ctx, "segmentationSupported", m.GetSegmentationSupported(), WriteComplex[BACnetSegmentationTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'segmentationSupported' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSegmentationSupported"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSegmentationSupported")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSegmentationSupported) IsBACnetConstructedDataSegmentationSupported() {
}

func (m *_BACnetConstructedDataSegmentationSupported) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataSegmentationSupported) deepCopy() *_BACnetConstructedDataSegmentationSupported {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataSegmentationSupportedCopy := &_BACnetConstructedDataSegmentationSupported{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.SegmentationSupported.DeepCopy().(BACnetSegmentationTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataSegmentationSupportedCopy
}

func (m *_BACnetConstructedDataSegmentationSupported) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
