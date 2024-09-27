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

// BACnetNotificationParametersUnsignedOutOfRange is the corresponding interface of BACnetNotificationParametersUnsignedOutOfRange
type BACnetNotificationParametersUnsignedOutOfRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetExceedingValue returns ExceedingValue (property field)
	GetExceedingValue() BACnetContextTagUnsignedInteger
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetDeadband returns Deadband (property field)
	GetDeadband() BACnetContextTagUnsignedInteger
	// GetExceededLimit returns ExceededLimit (property field)
	GetExceededLimit() BACnetContextTagUnsignedInteger
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
	// IsBACnetNotificationParametersUnsignedOutOfRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersUnsignedOutOfRange()
	// CreateBuilder creates a BACnetNotificationParametersUnsignedOutOfRangeBuilder
	CreateBACnetNotificationParametersUnsignedOutOfRangeBuilder() BACnetNotificationParametersUnsignedOutOfRangeBuilder
}

// _BACnetNotificationParametersUnsignedOutOfRange is the data-structure of this message
type _BACnetNotificationParametersUnsignedOutOfRange struct {
	BACnetNotificationParametersContract
	InnerOpeningTag BACnetOpeningTag
	ExceedingValue  BACnetContextTagUnsignedInteger
	StatusFlags     BACnetStatusFlagsTagged
	Deadband        BACnetContextTagUnsignedInteger
	ExceededLimit   BACnetContextTagUnsignedInteger
	InnerClosingTag BACnetClosingTag
}

var _ BACnetNotificationParametersUnsignedOutOfRange = (*_BACnetNotificationParametersUnsignedOutOfRange)(nil)
var _ BACnetNotificationParametersRequirements = (*_BACnetNotificationParametersUnsignedOutOfRange)(nil)

// NewBACnetNotificationParametersUnsignedOutOfRange factory function for _BACnetNotificationParametersUnsignedOutOfRange
func NewBACnetNotificationParametersUnsignedOutOfRange(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, innerOpeningTag BACnetOpeningTag, exceedingValue BACnetContextTagUnsignedInteger, statusFlags BACnetStatusFlagsTagged, deadband BACnetContextTagUnsignedInteger, exceededLimit BACnetContextTagUnsignedInteger, innerClosingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersUnsignedOutOfRange {
	if innerOpeningTag == nil {
		panic("innerOpeningTag of type BACnetOpeningTag for BACnetNotificationParametersUnsignedOutOfRange must not be nil")
	}
	if exceedingValue == nil {
		panic("exceedingValue of type BACnetContextTagUnsignedInteger for BACnetNotificationParametersUnsignedOutOfRange must not be nil")
	}
	if statusFlags == nil {
		panic("statusFlags of type BACnetStatusFlagsTagged for BACnetNotificationParametersUnsignedOutOfRange must not be nil")
	}
	if deadband == nil {
		panic("deadband of type BACnetContextTagUnsignedInteger for BACnetNotificationParametersUnsignedOutOfRange must not be nil")
	}
	if exceededLimit == nil {
		panic("exceededLimit of type BACnetContextTagUnsignedInteger for BACnetNotificationParametersUnsignedOutOfRange must not be nil")
	}
	if innerClosingTag == nil {
		panic("innerClosingTag of type BACnetClosingTag for BACnetNotificationParametersUnsignedOutOfRange must not be nil")
	}
	_result := &_BACnetNotificationParametersUnsignedOutOfRange{
		BACnetNotificationParametersContract: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
		InnerOpeningTag:                      innerOpeningTag,
		ExceedingValue:                       exceedingValue,
		StatusFlags:                          statusFlags,
		Deadband:                             deadband,
		ExceededLimit:                        exceededLimit,
		InnerClosingTag:                      innerClosingTag,
	}
	_result.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersUnsignedOutOfRangeBuilder is a builder for BACnetNotificationParametersUnsignedOutOfRange
type BACnetNotificationParametersUnsignedOutOfRangeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(innerOpeningTag BACnetOpeningTag, exceedingValue BACnetContextTagUnsignedInteger, statusFlags BACnetStatusFlagsTagged, deadband BACnetContextTagUnsignedInteger, exceededLimit BACnetContextTagUnsignedInteger, innerClosingTag BACnetClosingTag) BACnetNotificationParametersUnsignedOutOfRangeBuilder
	// WithInnerOpeningTag adds InnerOpeningTag (property field)
	WithInnerOpeningTag(BACnetOpeningTag) BACnetNotificationParametersUnsignedOutOfRangeBuilder
	// WithInnerOpeningTagBuilder adds InnerOpeningTag (property field) which is build by the builder
	WithInnerOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersUnsignedOutOfRangeBuilder
	// WithExceedingValue adds ExceedingValue (property field)
	WithExceedingValue(BACnetContextTagUnsignedInteger) BACnetNotificationParametersUnsignedOutOfRangeBuilder
	// WithExceedingValueBuilder adds ExceedingValue (property field) which is build by the builder
	WithExceedingValueBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersUnsignedOutOfRangeBuilder
	// WithStatusFlags adds StatusFlags (property field)
	WithStatusFlags(BACnetStatusFlagsTagged) BACnetNotificationParametersUnsignedOutOfRangeBuilder
	// WithStatusFlagsBuilder adds StatusFlags (property field) which is build by the builder
	WithStatusFlagsBuilder(func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetNotificationParametersUnsignedOutOfRangeBuilder
	// WithDeadband adds Deadband (property field)
	WithDeadband(BACnetContextTagUnsignedInteger) BACnetNotificationParametersUnsignedOutOfRangeBuilder
	// WithDeadbandBuilder adds Deadband (property field) which is build by the builder
	WithDeadbandBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersUnsignedOutOfRangeBuilder
	// WithExceededLimit adds ExceededLimit (property field)
	WithExceededLimit(BACnetContextTagUnsignedInteger) BACnetNotificationParametersUnsignedOutOfRangeBuilder
	// WithExceededLimitBuilder adds ExceededLimit (property field) which is build by the builder
	WithExceededLimitBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersUnsignedOutOfRangeBuilder
	// WithInnerClosingTag adds InnerClosingTag (property field)
	WithInnerClosingTag(BACnetClosingTag) BACnetNotificationParametersUnsignedOutOfRangeBuilder
	// WithInnerClosingTagBuilder adds InnerClosingTag (property field) which is build by the builder
	WithInnerClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersUnsignedOutOfRangeBuilder
	// Build builds the BACnetNotificationParametersUnsignedOutOfRange or returns an error if something is wrong
	Build() (BACnetNotificationParametersUnsignedOutOfRange, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersUnsignedOutOfRange
}

// NewBACnetNotificationParametersUnsignedOutOfRangeBuilder() creates a BACnetNotificationParametersUnsignedOutOfRangeBuilder
func NewBACnetNotificationParametersUnsignedOutOfRangeBuilder() BACnetNotificationParametersUnsignedOutOfRangeBuilder {
	return &_BACnetNotificationParametersUnsignedOutOfRangeBuilder{_BACnetNotificationParametersUnsignedOutOfRange: new(_BACnetNotificationParametersUnsignedOutOfRange)}
}

type _BACnetNotificationParametersUnsignedOutOfRangeBuilder struct {
	*_BACnetNotificationParametersUnsignedOutOfRange

	err *utils.MultiError
}

var _ (BACnetNotificationParametersUnsignedOutOfRangeBuilder) = (*_BACnetNotificationParametersUnsignedOutOfRangeBuilder)(nil)

func (m *_BACnetNotificationParametersUnsignedOutOfRangeBuilder) WithMandatoryFields(innerOpeningTag BACnetOpeningTag, exceedingValue BACnetContextTagUnsignedInteger, statusFlags BACnetStatusFlagsTagged, deadband BACnetContextTagUnsignedInteger, exceededLimit BACnetContextTagUnsignedInteger, innerClosingTag BACnetClosingTag) BACnetNotificationParametersUnsignedOutOfRangeBuilder {
	return m.WithInnerOpeningTag(innerOpeningTag).WithExceedingValue(exceedingValue).WithStatusFlags(statusFlags).WithDeadband(deadband).WithExceededLimit(exceededLimit).WithInnerClosingTag(innerClosingTag)
}

func (m *_BACnetNotificationParametersUnsignedOutOfRangeBuilder) WithInnerOpeningTag(innerOpeningTag BACnetOpeningTag) BACnetNotificationParametersUnsignedOutOfRangeBuilder {
	m.InnerOpeningTag = innerOpeningTag
	return m
}

func (m *_BACnetNotificationParametersUnsignedOutOfRangeBuilder) WithInnerOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersUnsignedOutOfRangeBuilder {
	builder := builderSupplier(m.InnerOpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	m.InnerOpeningTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return m
}

func (m *_BACnetNotificationParametersUnsignedOutOfRangeBuilder) WithExceedingValue(exceedingValue BACnetContextTagUnsignedInteger) BACnetNotificationParametersUnsignedOutOfRangeBuilder {
	m.ExceedingValue = exceedingValue
	return m
}

func (m *_BACnetNotificationParametersUnsignedOutOfRangeBuilder) WithExceedingValueBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersUnsignedOutOfRangeBuilder {
	builder := builderSupplier(m.ExceedingValue.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.ExceedingValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetNotificationParametersUnsignedOutOfRangeBuilder) WithStatusFlags(statusFlags BACnetStatusFlagsTagged) BACnetNotificationParametersUnsignedOutOfRangeBuilder {
	m.StatusFlags = statusFlags
	return m
}

func (m *_BACnetNotificationParametersUnsignedOutOfRangeBuilder) WithStatusFlagsBuilder(builderSupplier func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetNotificationParametersUnsignedOutOfRangeBuilder {
	builder := builderSupplier(m.StatusFlags.CreateBACnetStatusFlagsTaggedBuilder())
	var err error
	m.StatusFlags, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetStatusFlagsTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetNotificationParametersUnsignedOutOfRangeBuilder) WithDeadband(deadband BACnetContextTagUnsignedInteger) BACnetNotificationParametersUnsignedOutOfRangeBuilder {
	m.Deadband = deadband
	return m
}

func (m *_BACnetNotificationParametersUnsignedOutOfRangeBuilder) WithDeadbandBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersUnsignedOutOfRangeBuilder {
	builder := builderSupplier(m.Deadband.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.Deadband, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetNotificationParametersUnsignedOutOfRangeBuilder) WithExceededLimit(exceededLimit BACnetContextTagUnsignedInteger) BACnetNotificationParametersUnsignedOutOfRangeBuilder {
	m.ExceededLimit = exceededLimit
	return m
}

func (m *_BACnetNotificationParametersUnsignedOutOfRangeBuilder) WithExceededLimitBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersUnsignedOutOfRangeBuilder {
	builder := builderSupplier(m.ExceededLimit.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.ExceededLimit, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetNotificationParametersUnsignedOutOfRangeBuilder) WithInnerClosingTag(innerClosingTag BACnetClosingTag) BACnetNotificationParametersUnsignedOutOfRangeBuilder {
	m.InnerClosingTag = innerClosingTag
	return m
}

func (m *_BACnetNotificationParametersUnsignedOutOfRangeBuilder) WithInnerClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersUnsignedOutOfRangeBuilder {
	builder := builderSupplier(m.InnerClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	m.InnerClosingTag, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return m
}

func (m *_BACnetNotificationParametersUnsignedOutOfRangeBuilder) Build() (BACnetNotificationParametersUnsignedOutOfRange, error) {
	if m.InnerOpeningTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'innerOpeningTag' not set"))
	}
	if m.ExceedingValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'exceedingValue' not set"))
	}
	if m.StatusFlags == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'statusFlags' not set"))
	}
	if m.Deadband == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'deadband' not set"))
	}
	if m.ExceededLimit == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'exceededLimit' not set"))
	}
	if m.InnerClosingTag == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'innerClosingTag' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetNotificationParametersUnsignedOutOfRange.deepCopy(), nil
}

func (m *_BACnetNotificationParametersUnsignedOutOfRangeBuilder) MustBuild() BACnetNotificationParametersUnsignedOutOfRange {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetNotificationParametersUnsignedOutOfRangeBuilder) DeepCopy() any {
	return m.CreateBACnetNotificationParametersUnsignedOutOfRangeBuilder()
}

// CreateBACnetNotificationParametersUnsignedOutOfRangeBuilder creates a BACnetNotificationParametersUnsignedOutOfRangeBuilder
func (m *_BACnetNotificationParametersUnsignedOutOfRange) CreateBACnetNotificationParametersUnsignedOutOfRangeBuilder() BACnetNotificationParametersUnsignedOutOfRangeBuilder {
	if m == nil {
		return NewBACnetNotificationParametersUnsignedOutOfRangeBuilder()
	}
	return &_BACnetNotificationParametersUnsignedOutOfRangeBuilder{_BACnetNotificationParametersUnsignedOutOfRange: m.deepCopy()}
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

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetParent() BACnetNotificationParametersContract {
	return m.BACnetNotificationParametersContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetExceedingValue() BACnetContextTagUnsignedInteger {
	return m.ExceedingValue
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetDeadband() BACnetContextTagUnsignedInteger {
	return m.Deadband
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetExceededLimit() BACnetContextTagUnsignedInteger {
	return m.ExceededLimit
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersUnsignedOutOfRange(structType any) BACnetNotificationParametersUnsignedOutOfRange {
	if casted, ok := structType.(BACnetNotificationParametersUnsignedOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersUnsignedOutOfRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetTypeName() string {
	return "BACnetNotificationParametersUnsignedOutOfRange"
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).getLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (exceedingValue)
	lengthInBits += m.ExceedingValue.GetLengthInBits(ctx)

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits(ctx)

	// Simple field (exceededLimit)
	lengthInBits += m.ExceededLimit.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParameters, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (__bACnetNotificationParametersUnsignedOutOfRange BACnetNotificationParametersUnsignedOutOfRange, err error) {
	m.BACnetNotificationParametersContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersUnsignedOutOfRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersUnsignedOutOfRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	exceedingValue, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "exceedingValue", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exceedingValue' field"))
	}
	m.ExceedingValue = exceedingValue

	statusFlags, err := ReadSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusFlags' field"))
	}
	m.StatusFlags = statusFlags

	deadband, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "deadband", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deadband' field"))
	}
	m.Deadband = deadband

	exceededLimit, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "exceededLimit", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exceededLimit' field"))
	}
	m.ExceededLimit = exceededLimit

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersUnsignedOutOfRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersUnsignedOutOfRange")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersUnsignedOutOfRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersUnsignedOutOfRange")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "exceedingValue", m.GetExceedingValue(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'exceedingValue' field")
		}

		if err := WriteSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", m.GetStatusFlags(), WriteComplex[BACnetStatusFlagsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusFlags' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "deadband", m.GetDeadband(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deadband' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "exceededLimit", m.GetExceededLimit(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'exceededLimit' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersUnsignedOutOfRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersUnsignedOutOfRange")
		}
		return nil
	}
	return m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) IsBACnetNotificationParametersUnsignedOutOfRange() {
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) deepCopy() *_BACnetNotificationParametersUnsignedOutOfRange {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersUnsignedOutOfRangeCopy := &_BACnetNotificationParametersUnsignedOutOfRange{
		m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).deepCopy(),
		m.InnerOpeningTag.DeepCopy().(BACnetOpeningTag),
		m.ExceedingValue.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.StatusFlags.DeepCopy().(BACnetStatusFlagsTagged),
		m.Deadband.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.ExceededLimit.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.InnerClosingTag.DeepCopy().(BACnetClosingTag),
	}
	m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = m
	return _BACnetNotificationParametersUnsignedOutOfRangeCopy
}

func (m *_BACnetNotificationParametersUnsignedOutOfRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
