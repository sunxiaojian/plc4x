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

// BACnetEventParameterCommandFailure is the corresponding interface of BACnetEventParameterCommandFailure
type BACnetEventParameterCommandFailure interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetFeedbackPropertyReference returns FeedbackPropertyReference (property field)
	GetFeedbackPropertyReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterCommandFailure is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterCommandFailure()
	// CreateBuilder creates a BACnetEventParameterCommandFailureBuilder
	CreateBACnetEventParameterCommandFailureBuilder() BACnetEventParameterCommandFailureBuilder
}

// _BACnetEventParameterCommandFailure is the data-structure of this message
type _BACnetEventParameterCommandFailure struct {
	BACnetEventParameterContract
	OpeningTag                BACnetOpeningTag
	TimeDelay                 BACnetContextTagUnsignedInteger
	FeedbackPropertyReference BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag                BACnetClosingTag
}

var _ BACnetEventParameterCommandFailure = (*_BACnetEventParameterCommandFailure)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterCommandFailure)(nil)

// NewBACnetEventParameterCommandFailure factory function for _BACnetEventParameterCommandFailure
func NewBACnetEventParameterCommandFailure(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, feedbackPropertyReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) *_BACnetEventParameterCommandFailure {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterCommandFailure must not be nil")
	}
	if timeDelay == nil {
		panic("timeDelay of type BACnetContextTagUnsignedInteger for BACnetEventParameterCommandFailure must not be nil")
	}
	if feedbackPropertyReference == nil {
		panic("feedbackPropertyReference of type BACnetDeviceObjectPropertyReferenceEnclosed for BACnetEventParameterCommandFailure must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterCommandFailure must not be nil")
	}
	_result := &_BACnetEventParameterCommandFailure{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		TimeDelay:                    timeDelay,
		FeedbackPropertyReference:    feedbackPropertyReference,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterCommandFailureBuilder is a builder for BACnetEventParameterCommandFailure
type BACnetEventParameterCommandFailureBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, feedbackPropertyReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) BACnetEventParameterCommandFailureBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterCommandFailureBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterCommandFailureBuilder
	// WithTimeDelay adds TimeDelay (property field)
	WithTimeDelay(BACnetContextTagUnsignedInteger) BACnetEventParameterCommandFailureBuilder
	// WithTimeDelayBuilder adds TimeDelay (property field) which is build by the builder
	WithTimeDelayBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterCommandFailureBuilder
	// WithFeedbackPropertyReference adds FeedbackPropertyReference (property field)
	WithFeedbackPropertyReference(BACnetDeviceObjectPropertyReferenceEnclosed) BACnetEventParameterCommandFailureBuilder
	// WithFeedbackPropertyReferenceBuilder adds FeedbackPropertyReference (property field) which is build by the builder
	WithFeedbackPropertyReferenceBuilder(func(BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetEventParameterCommandFailureBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterCommandFailureBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterCommandFailureBuilder
	// Build builds the BACnetEventParameterCommandFailure or returns an error if something is wrong
	Build() (BACnetEventParameterCommandFailure, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterCommandFailure
}

// NewBACnetEventParameterCommandFailureBuilder() creates a BACnetEventParameterCommandFailureBuilder
func NewBACnetEventParameterCommandFailureBuilder() BACnetEventParameterCommandFailureBuilder {
	return &_BACnetEventParameterCommandFailureBuilder{_BACnetEventParameterCommandFailure: new(_BACnetEventParameterCommandFailure)}
}

type _BACnetEventParameterCommandFailureBuilder struct {
	*_BACnetEventParameterCommandFailure

	err *utils.MultiError
}

var _ (BACnetEventParameterCommandFailureBuilder) = (*_BACnetEventParameterCommandFailureBuilder)(nil)

func (m *_BACnetEventParameterCommandFailureBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, feedbackPropertyReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) BACnetEventParameterCommandFailureBuilder {
	return m.WithOpeningTag(openingTag).WithTimeDelay(timeDelay).WithFeedbackPropertyReference(feedbackPropertyReference).WithClosingTag(closingTag)
}

func (m *_BACnetEventParameterCommandFailureBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterCommandFailureBuilder {
	m.OpeningTag = openingTag
	return m
}

func (m *_BACnetEventParameterCommandFailureBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterCommandFailureBuilder {
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

func (m *_BACnetEventParameterCommandFailureBuilder) WithTimeDelay(timeDelay BACnetContextTagUnsignedInteger) BACnetEventParameterCommandFailureBuilder {
	m.TimeDelay = timeDelay
	return m
}

func (m *_BACnetEventParameterCommandFailureBuilder) WithTimeDelayBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterCommandFailureBuilder {
	builder := builderSupplier(m.TimeDelay.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.TimeDelay, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetEventParameterCommandFailureBuilder) WithFeedbackPropertyReference(feedbackPropertyReference BACnetDeviceObjectPropertyReferenceEnclosed) BACnetEventParameterCommandFailureBuilder {
	m.FeedbackPropertyReference = feedbackPropertyReference
	return m
}

func (m *_BACnetEventParameterCommandFailureBuilder) WithFeedbackPropertyReferenceBuilder(builderSupplier func(BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetEventParameterCommandFailureBuilder {
	builder := builderSupplier(m.FeedbackPropertyReference.CreateBACnetDeviceObjectPropertyReferenceEnclosedBuilder())
	var err error
	m.FeedbackPropertyReference, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetDeviceObjectPropertyReferenceEnclosedBuilder failed"))
	}
	return m
}

func (m *_BACnetEventParameterCommandFailureBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterCommandFailureBuilder {
	m.ClosingTag = closingTag
	return m
}

func (m *_BACnetEventParameterCommandFailureBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterCommandFailureBuilder {
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

func (m *_BACnetEventParameterCommandFailureBuilder) Build() (BACnetEventParameterCommandFailure, error) {
	if m.OpeningTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if m.TimeDelay == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'timeDelay' not set"))
	}
	if m.FeedbackPropertyReference == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'feedbackPropertyReference' not set"))
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
	return m._BACnetEventParameterCommandFailure.deepCopy(), nil
}

func (m *_BACnetEventParameterCommandFailureBuilder) MustBuild() BACnetEventParameterCommandFailure {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetEventParameterCommandFailureBuilder) DeepCopy() any {
	return m.CreateBACnetEventParameterCommandFailureBuilder()
}

// CreateBACnetEventParameterCommandFailureBuilder creates a BACnetEventParameterCommandFailureBuilder
func (m *_BACnetEventParameterCommandFailure) CreateBACnetEventParameterCommandFailureBuilder() BACnetEventParameterCommandFailureBuilder {
	if m == nil {
		return NewBACnetEventParameterCommandFailureBuilder()
	}
	return &_BACnetEventParameterCommandFailureBuilder{_BACnetEventParameterCommandFailure: m.deepCopy()}
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

func (m *_BACnetEventParameterCommandFailure) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterCommandFailure) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterCommandFailure) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterCommandFailure) GetFeedbackPropertyReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.FeedbackPropertyReference
}

func (m *_BACnetEventParameterCommandFailure) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterCommandFailure(structType any) BACnetEventParameterCommandFailure {
	if casted, ok := structType.(BACnetEventParameterCommandFailure); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterCommandFailure); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterCommandFailure) GetTypeName() string {
	return "BACnetEventParameterCommandFailure"
}

func (m *_BACnetEventParameterCommandFailure) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (feedbackPropertyReference)
	lengthInBits += m.FeedbackPropertyReference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterCommandFailure) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterCommandFailure) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterCommandFailure BACnetEventParameterCommandFailure, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterCommandFailure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterCommandFailure")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(3))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	timeDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}
	m.TimeDelay = timeDelay

	feedbackPropertyReference, err := ReadSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "feedbackPropertyReference", ReadComplex[BACnetDeviceObjectPropertyReferenceEnclosed](BACnetDeviceObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'feedbackPropertyReference' field"))
	}
	m.FeedbackPropertyReference = feedbackPropertyReference

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(3))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterCommandFailure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterCommandFailure")
	}

	return m, nil
}

func (m *_BACnetEventParameterCommandFailure) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterCommandFailure) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterCommandFailure"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterCommandFailure")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", m.GetTimeDelay(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelay' field")
		}

		if err := WriteSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "feedbackPropertyReference", m.GetFeedbackPropertyReference(), WriteComplex[BACnetDeviceObjectPropertyReferenceEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'feedbackPropertyReference' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterCommandFailure"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterCommandFailure")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterCommandFailure) IsBACnetEventParameterCommandFailure() {}

func (m *_BACnetEventParameterCommandFailure) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterCommandFailure) deepCopy() *_BACnetEventParameterCommandFailure {
	if m == nil {
		return nil
	}
	_BACnetEventParameterCommandFailureCopy := &_BACnetEventParameterCommandFailure{
		m.BACnetEventParameterContract.(*_BACnetEventParameter).deepCopy(),
		m.OpeningTag.DeepCopy().(BACnetOpeningTag),
		m.TimeDelay.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.FeedbackPropertyReference.DeepCopy().(BACnetDeviceObjectPropertyReferenceEnclosed),
		m.ClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = m
	return _BACnetEventParameterCommandFailureCopy
}

func (m *_BACnetEventParameterCommandFailure) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
