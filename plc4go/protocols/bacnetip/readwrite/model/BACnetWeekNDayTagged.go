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

// BACnetWeekNDayTagged is the corresponding interface of BACnetWeekNDayTagged
type BACnetWeekNDayTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetMonth returns Month (property field)
	GetMonth() uint8
	// GetWeekOfMonth returns WeekOfMonth (property field)
	GetWeekOfMonth() uint8
	// GetDayOfWeek returns DayOfWeek (property field)
	GetDayOfWeek() uint8
	// GetOddMonths returns OddMonths (virtual field)
	GetOddMonths() bool
	// GetEvenMonths returns EvenMonths (virtual field)
	GetEvenMonths() bool
	// GetAnyMonth returns AnyMonth (virtual field)
	GetAnyMonth() bool
	// GetDays1to7 returns Days1to7 (virtual field)
	GetDays1to7() bool
	// GetDays8to14 returns Days8to14 (virtual field)
	GetDays8to14() bool
	// GetDays15to21 returns Days15to21 (virtual field)
	GetDays15to21() bool
	// GetDays22to28 returns Days22to28 (virtual field)
	GetDays22to28() bool
	// GetDays29to31 returns Days29to31 (virtual field)
	GetDays29to31() bool
	// GetLast7DaysOfThisMonth returns Last7DaysOfThisMonth (virtual field)
	GetLast7DaysOfThisMonth() bool
	// GetAny7DaysPriorToLast7DaysOfThisMonth returns Any7DaysPriorToLast7DaysOfThisMonth (virtual field)
	GetAny7DaysPriorToLast7DaysOfThisMonth() bool
	// GetAny7DaysPriorToLast14DaysOfThisMonth returns Any7DaysPriorToLast14DaysOfThisMonth (virtual field)
	GetAny7DaysPriorToLast14DaysOfThisMonth() bool
	// GetAny7DaysPriorToLast21DaysOfThisMonth returns Any7DaysPriorToLast21DaysOfThisMonth (virtual field)
	GetAny7DaysPriorToLast21DaysOfThisMonth() bool
	// GetAnyWeekOfthisMonth returns AnyWeekOfthisMonth (virtual field)
	GetAnyWeekOfthisMonth() bool
	// GetAnyDayOfWeek returns AnyDayOfWeek (virtual field)
	GetAnyDayOfWeek() bool
	// IsBACnetWeekNDayTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetWeekNDayTagged()
	// CreateBuilder creates a BACnetWeekNDayTaggedBuilder
	CreateBACnetWeekNDayTaggedBuilder() BACnetWeekNDayTaggedBuilder
}

// _BACnetWeekNDayTagged is the data-structure of this message
type _BACnetWeekNDayTagged struct {
	Header      BACnetTagHeader
	Month       uint8
	WeekOfMonth uint8
	DayOfWeek   uint8

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetWeekNDayTagged = (*_BACnetWeekNDayTagged)(nil)

// NewBACnetWeekNDayTagged factory function for _BACnetWeekNDayTagged
func NewBACnetWeekNDayTagged(header BACnetTagHeader, month uint8, weekOfMonth uint8, dayOfWeek uint8, tagNumber uint8, tagClass TagClass) *_BACnetWeekNDayTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetWeekNDayTagged must not be nil")
	}
	return &_BACnetWeekNDayTagged{Header: header, Month: month, WeekOfMonth: weekOfMonth, DayOfWeek: dayOfWeek, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetWeekNDayTaggedBuilder is a builder for BACnetWeekNDayTagged
type BACnetWeekNDayTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, month uint8, weekOfMonth uint8, dayOfWeek uint8) BACnetWeekNDayTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetWeekNDayTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetWeekNDayTaggedBuilder
	// WithMonth adds Month (property field)
	WithMonth(uint8) BACnetWeekNDayTaggedBuilder
	// WithWeekOfMonth adds WeekOfMonth (property field)
	WithWeekOfMonth(uint8) BACnetWeekNDayTaggedBuilder
	// WithDayOfWeek adds DayOfWeek (property field)
	WithDayOfWeek(uint8) BACnetWeekNDayTaggedBuilder
	// Build builds the BACnetWeekNDayTagged or returns an error if something is wrong
	Build() (BACnetWeekNDayTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetWeekNDayTagged
}

// NewBACnetWeekNDayTaggedBuilder() creates a BACnetWeekNDayTaggedBuilder
func NewBACnetWeekNDayTaggedBuilder() BACnetWeekNDayTaggedBuilder {
	return &_BACnetWeekNDayTaggedBuilder{_BACnetWeekNDayTagged: new(_BACnetWeekNDayTagged)}
}

type _BACnetWeekNDayTaggedBuilder struct {
	*_BACnetWeekNDayTagged

	err *utils.MultiError
}

var _ (BACnetWeekNDayTaggedBuilder) = (*_BACnetWeekNDayTaggedBuilder)(nil)

func (m *_BACnetWeekNDayTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, month uint8, weekOfMonth uint8, dayOfWeek uint8) BACnetWeekNDayTaggedBuilder {
	return m.WithHeader(header).WithMonth(month).WithWeekOfMonth(weekOfMonth).WithDayOfWeek(dayOfWeek)
}

func (m *_BACnetWeekNDayTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetWeekNDayTaggedBuilder {
	m.Header = header
	return m
}

func (m *_BACnetWeekNDayTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetWeekNDayTaggedBuilder {
	builder := builderSupplier(m.Header.CreateBACnetTagHeaderBuilder())
	var err error
	m.Header, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return m
}

func (m *_BACnetWeekNDayTaggedBuilder) WithMonth(month uint8) BACnetWeekNDayTaggedBuilder {
	m.Month = month
	return m
}

func (m *_BACnetWeekNDayTaggedBuilder) WithWeekOfMonth(weekOfMonth uint8) BACnetWeekNDayTaggedBuilder {
	m.WeekOfMonth = weekOfMonth
	return m
}

func (m *_BACnetWeekNDayTaggedBuilder) WithDayOfWeek(dayOfWeek uint8) BACnetWeekNDayTaggedBuilder {
	m.DayOfWeek = dayOfWeek
	return m
}

func (m *_BACnetWeekNDayTaggedBuilder) Build() (BACnetWeekNDayTagged, error) {
	if m.Header == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetWeekNDayTagged.deepCopy(), nil
}

func (m *_BACnetWeekNDayTaggedBuilder) MustBuild() BACnetWeekNDayTagged {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetWeekNDayTaggedBuilder) DeepCopy() any {
	return m.CreateBACnetWeekNDayTaggedBuilder()
}

// CreateBACnetWeekNDayTaggedBuilder creates a BACnetWeekNDayTaggedBuilder
func (m *_BACnetWeekNDayTagged) CreateBACnetWeekNDayTaggedBuilder() BACnetWeekNDayTaggedBuilder {
	if m == nil {
		return NewBACnetWeekNDayTaggedBuilder()
	}
	return &_BACnetWeekNDayTaggedBuilder{_BACnetWeekNDayTagged: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetWeekNDayTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetWeekNDayTagged) GetMonth() uint8 {
	return m.Month
}

func (m *_BACnetWeekNDayTagged) GetWeekOfMonth() uint8 {
	return m.WeekOfMonth
}

func (m *_BACnetWeekNDayTagged) GetDayOfWeek() uint8 {
	return m.DayOfWeek
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetWeekNDayTagged) GetOddMonths() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetMonth()) == (13)))
}

func (m *_BACnetWeekNDayTagged) GetEvenMonths() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetMonth()) == (14)))
}

func (m *_BACnetWeekNDayTagged) GetAnyMonth() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetMonth()) == (0xFF)))
}

func (m *_BACnetWeekNDayTagged) GetDays1to7() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetWeekOfMonth()) == (1)))
}

func (m *_BACnetWeekNDayTagged) GetDays8to14() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetWeekOfMonth()) == (2)))
}

func (m *_BACnetWeekNDayTagged) GetDays15to21() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetWeekOfMonth()) == (3)))
}

func (m *_BACnetWeekNDayTagged) GetDays22to28() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetWeekOfMonth()) == (4)))
}

func (m *_BACnetWeekNDayTagged) GetDays29to31() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetWeekOfMonth()) == (5)))
}

func (m *_BACnetWeekNDayTagged) GetLast7DaysOfThisMonth() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetWeekOfMonth()) == (6)))
}

func (m *_BACnetWeekNDayTagged) GetAny7DaysPriorToLast7DaysOfThisMonth() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetWeekOfMonth()) == (7)))
}

func (m *_BACnetWeekNDayTagged) GetAny7DaysPriorToLast14DaysOfThisMonth() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetWeekOfMonth()) == (8)))
}

func (m *_BACnetWeekNDayTagged) GetAny7DaysPriorToLast21DaysOfThisMonth() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetWeekOfMonth()) == (9)))
}

func (m *_BACnetWeekNDayTagged) GetAnyWeekOfthisMonth() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetWeekOfMonth()) == (0xFF)))
}

func (m *_BACnetWeekNDayTagged) GetAnyDayOfWeek() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetDayOfWeek()) == (0xFF)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetWeekNDayTagged(structType any) BACnetWeekNDayTagged {
	if casted, ok := structType.(BACnetWeekNDayTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetWeekNDayTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetWeekNDayTagged) GetTypeName() string {
	return "BACnetWeekNDayTagged"
}

func (m *_BACnetWeekNDayTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Simple field (month)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (weekOfMonth)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (dayOfWeek)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetWeekNDayTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetWeekNDayTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetWeekNDayTagged, error) {
	return BACnetWeekNDayTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetWeekNDayTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetWeekNDayTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetWeekNDayTagged, error) {
		return BACnetWeekNDayTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetWeekNDayTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetWeekNDayTagged, error) {
	v, err := (&_BACnetWeekNDayTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetWeekNDayTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetWeekNDayTagged BACnetWeekNDayTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetWeekNDayTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetWeekNDayTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	// Validation
	if !(bool((header.GetActualLength()) == (3))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "We should have at least 3 octets"})
	}

	month, err := ReadSimpleField(ctx, "month", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'month' field"))
	}
	m.Month = month

	oddMonths, err := ReadVirtualField[bool](ctx, "oddMonths", (*bool)(nil), bool((month) == (13)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'oddMonths' field"))
	}
	_ = oddMonths

	evenMonths, err := ReadVirtualField[bool](ctx, "evenMonths", (*bool)(nil), bool((month) == (14)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'evenMonths' field"))
	}
	_ = evenMonths

	anyMonth, err := ReadVirtualField[bool](ctx, "anyMonth", (*bool)(nil), bool((month) == (0xFF)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'anyMonth' field"))
	}
	_ = anyMonth

	weekOfMonth, err := ReadSimpleField(ctx, "weekOfMonth", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'weekOfMonth' field"))
	}
	m.WeekOfMonth = weekOfMonth

	days1to7, err := ReadVirtualField[bool](ctx, "days1to7", (*bool)(nil), bool((weekOfMonth) == (1)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'days1to7' field"))
	}
	_ = days1to7

	days8to14, err := ReadVirtualField[bool](ctx, "days8to14", (*bool)(nil), bool((weekOfMonth) == (2)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'days8to14' field"))
	}
	_ = days8to14

	days15to21, err := ReadVirtualField[bool](ctx, "days15to21", (*bool)(nil), bool((weekOfMonth) == (3)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'days15to21' field"))
	}
	_ = days15to21

	days22to28, err := ReadVirtualField[bool](ctx, "days22to28", (*bool)(nil), bool((weekOfMonth) == (4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'days22to28' field"))
	}
	_ = days22to28

	days29to31, err := ReadVirtualField[bool](ctx, "days29to31", (*bool)(nil), bool((weekOfMonth) == (5)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'days29to31' field"))
	}
	_ = days29to31

	last7DaysOfThisMonth, err := ReadVirtualField[bool](ctx, "last7DaysOfThisMonth", (*bool)(nil), bool((weekOfMonth) == (6)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'last7DaysOfThisMonth' field"))
	}
	_ = last7DaysOfThisMonth

	any7DaysPriorToLast7DaysOfThisMonth, err := ReadVirtualField[bool](ctx, "any7DaysPriorToLast7DaysOfThisMonth", (*bool)(nil), bool((weekOfMonth) == (7)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'any7DaysPriorToLast7DaysOfThisMonth' field"))
	}
	_ = any7DaysPriorToLast7DaysOfThisMonth

	any7DaysPriorToLast14DaysOfThisMonth, err := ReadVirtualField[bool](ctx, "any7DaysPriorToLast14DaysOfThisMonth", (*bool)(nil), bool((weekOfMonth) == (8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'any7DaysPriorToLast14DaysOfThisMonth' field"))
	}
	_ = any7DaysPriorToLast14DaysOfThisMonth

	any7DaysPriorToLast21DaysOfThisMonth, err := ReadVirtualField[bool](ctx, "any7DaysPriorToLast21DaysOfThisMonth", (*bool)(nil), bool((weekOfMonth) == (9)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'any7DaysPriorToLast21DaysOfThisMonth' field"))
	}
	_ = any7DaysPriorToLast21DaysOfThisMonth

	anyWeekOfthisMonth, err := ReadVirtualField[bool](ctx, "anyWeekOfthisMonth", (*bool)(nil), bool((weekOfMonth) == (0xFF)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'anyWeekOfthisMonth' field"))
	}
	_ = anyWeekOfthisMonth

	dayOfWeek, err := ReadSimpleField(ctx, "dayOfWeek", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dayOfWeek' field"))
	}
	m.DayOfWeek = dayOfWeek

	anyDayOfWeek, err := ReadVirtualField[bool](ctx, "anyDayOfWeek", (*bool)(nil), bool((dayOfWeek) == (0xFF)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'anyDayOfWeek' field"))
	}
	_ = anyDayOfWeek

	if closeErr := readBuffer.CloseContext("BACnetWeekNDayTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetWeekNDayTagged")
	}

	return m, nil
}

func (m *_BACnetWeekNDayTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetWeekNDayTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetWeekNDayTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetWeekNDayTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteSimpleField[uint8](ctx, "month", m.GetMonth(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'month' field")
	}
	// Virtual field
	oddMonths := m.GetOddMonths()
	_ = oddMonths
	if _oddMonthsErr := writeBuffer.WriteVirtual(ctx, "oddMonths", m.GetOddMonths()); _oddMonthsErr != nil {
		return errors.Wrap(_oddMonthsErr, "Error serializing 'oddMonths' field")
	}
	// Virtual field
	evenMonths := m.GetEvenMonths()
	_ = evenMonths
	if _evenMonthsErr := writeBuffer.WriteVirtual(ctx, "evenMonths", m.GetEvenMonths()); _evenMonthsErr != nil {
		return errors.Wrap(_evenMonthsErr, "Error serializing 'evenMonths' field")
	}
	// Virtual field
	anyMonth := m.GetAnyMonth()
	_ = anyMonth
	if _anyMonthErr := writeBuffer.WriteVirtual(ctx, "anyMonth", m.GetAnyMonth()); _anyMonthErr != nil {
		return errors.Wrap(_anyMonthErr, "Error serializing 'anyMonth' field")
	}

	if err := WriteSimpleField[uint8](ctx, "weekOfMonth", m.GetWeekOfMonth(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'weekOfMonth' field")
	}
	// Virtual field
	days1to7 := m.GetDays1to7()
	_ = days1to7
	if _days1to7Err := writeBuffer.WriteVirtual(ctx, "days1to7", m.GetDays1to7()); _days1to7Err != nil {
		return errors.Wrap(_days1to7Err, "Error serializing 'days1to7' field")
	}
	// Virtual field
	days8to14 := m.GetDays8to14()
	_ = days8to14
	if _days8to14Err := writeBuffer.WriteVirtual(ctx, "days8to14", m.GetDays8to14()); _days8to14Err != nil {
		return errors.Wrap(_days8to14Err, "Error serializing 'days8to14' field")
	}
	// Virtual field
	days15to21 := m.GetDays15to21()
	_ = days15to21
	if _days15to21Err := writeBuffer.WriteVirtual(ctx, "days15to21", m.GetDays15to21()); _days15to21Err != nil {
		return errors.Wrap(_days15to21Err, "Error serializing 'days15to21' field")
	}
	// Virtual field
	days22to28 := m.GetDays22to28()
	_ = days22to28
	if _days22to28Err := writeBuffer.WriteVirtual(ctx, "days22to28", m.GetDays22to28()); _days22to28Err != nil {
		return errors.Wrap(_days22to28Err, "Error serializing 'days22to28' field")
	}
	// Virtual field
	days29to31 := m.GetDays29to31()
	_ = days29to31
	if _days29to31Err := writeBuffer.WriteVirtual(ctx, "days29to31", m.GetDays29to31()); _days29to31Err != nil {
		return errors.Wrap(_days29to31Err, "Error serializing 'days29to31' field")
	}
	// Virtual field
	last7DaysOfThisMonth := m.GetLast7DaysOfThisMonth()
	_ = last7DaysOfThisMonth
	if _last7DaysOfThisMonthErr := writeBuffer.WriteVirtual(ctx, "last7DaysOfThisMonth", m.GetLast7DaysOfThisMonth()); _last7DaysOfThisMonthErr != nil {
		return errors.Wrap(_last7DaysOfThisMonthErr, "Error serializing 'last7DaysOfThisMonth' field")
	}
	// Virtual field
	any7DaysPriorToLast7DaysOfThisMonth := m.GetAny7DaysPriorToLast7DaysOfThisMonth()
	_ = any7DaysPriorToLast7DaysOfThisMonth
	if _any7DaysPriorToLast7DaysOfThisMonthErr := writeBuffer.WriteVirtual(ctx, "any7DaysPriorToLast7DaysOfThisMonth", m.GetAny7DaysPriorToLast7DaysOfThisMonth()); _any7DaysPriorToLast7DaysOfThisMonthErr != nil {
		return errors.Wrap(_any7DaysPriorToLast7DaysOfThisMonthErr, "Error serializing 'any7DaysPriorToLast7DaysOfThisMonth' field")
	}
	// Virtual field
	any7DaysPriorToLast14DaysOfThisMonth := m.GetAny7DaysPriorToLast14DaysOfThisMonth()
	_ = any7DaysPriorToLast14DaysOfThisMonth
	if _any7DaysPriorToLast14DaysOfThisMonthErr := writeBuffer.WriteVirtual(ctx, "any7DaysPriorToLast14DaysOfThisMonth", m.GetAny7DaysPriorToLast14DaysOfThisMonth()); _any7DaysPriorToLast14DaysOfThisMonthErr != nil {
		return errors.Wrap(_any7DaysPriorToLast14DaysOfThisMonthErr, "Error serializing 'any7DaysPriorToLast14DaysOfThisMonth' field")
	}
	// Virtual field
	any7DaysPriorToLast21DaysOfThisMonth := m.GetAny7DaysPriorToLast21DaysOfThisMonth()
	_ = any7DaysPriorToLast21DaysOfThisMonth
	if _any7DaysPriorToLast21DaysOfThisMonthErr := writeBuffer.WriteVirtual(ctx, "any7DaysPriorToLast21DaysOfThisMonth", m.GetAny7DaysPriorToLast21DaysOfThisMonth()); _any7DaysPriorToLast21DaysOfThisMonthErr != nil {
		return errors.Wrap(_any7DaysPriorToLast21DaysOfThisMonthErr, "Error serializing 'any7DaysPriorToLast21DaysOfThisMonth' field")
	}
	// Virtual field
	anyWeekOfthisMonth := m.GetAnyWeekOfthisMonth()
	_ = anyWeekOfthisMonth
	if _anyWeekOfthisMonthErr := writeBuffer.WriteVirtual(ctx, "anyWeekOfthisMonth", m.GetAnyWeekOfthisMonth()); _anyWeekOfthisMonthErr != nil {
		return errors.Wrap(_anyWeekOfthisMonthErr, "Error serializing 'anyWeekOfthisMonth' field")
	}

	if err := WriteSimpleField[uint8](ctx, "dayOfWeek", m.GetDayOfWeek(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'dayOfWeek' field")
	}
	// Virtual field
	anyDayOfWeek := m.GetAnyDayOfWeek()
	_ = anyDayOfWeek
	if _anyDayOfWeekErr := writeBuffer.WriteVirtual(ctx, "anyDayOfWeek", m.GetAnyDayOfWeek()); _anyDayOfWeekErr != nil {
		return errors.Wrap(_anyDayOfWeekErr, "Error serializing 'anyDayOfWeek' field")
	}

	if popErr := writeBuffer.PopContext("BACnetWeekNDayTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetWeekNDayTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetWeekNDayTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetWeekNDayTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetWeekNDayTagged) IsBACnetWeekNDayTagged() {}

func (m *_BACnetWeekNDayTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetWeekNDayTagged) deepCopy() *_BACnetWeekNDayTagged {
	if m == nil {
		return nil
	}
	_BACnetWeekNDayTaggedCopy := &_BACnetWeekNDayTagged{
		m.Header.DeepCopy().(BACnetTagHeader),
		m.Month,
		m.WeekOfMonth,
		m.DayOfWeek,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetWeekNDayTaggedCopy
}

func (m *_BACnetWeekNDayTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
