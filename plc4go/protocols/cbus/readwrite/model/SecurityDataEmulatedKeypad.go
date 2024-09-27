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

// SecurityDataEmulatedKeypad is the corresponding interface of SecurityDataEmulatedKeypad
type SecurityDataEmulatedKeypad interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// GetKey returns Key (property field)
	GetKey() byte
	// GetIsAscii returns IsAscii (virtual field)
	GetIsAscii() bool
	// GetIsCustom returns IsCustom (virtual field)
	GetIsCustom() bool
	// GetIsEnter returns IsEnter (virtual field)
	GetIsEnter() bool
	// GetIsShift returns IsShift (virtual field)
	GetIsShift() bool
	// GetIsPanic returns IsPanic (virtual field)
	GetIsPanic() bool
	// GetIsFire returns IsFire (virtual field)
	GetIsFire() bool
	// GetIsARM returns IsARM (virtual field)
	GetIsARM() bool
	// GetIsAway returns IsAway (virtual field)
	GetIsAway() bool
	// GetIsNight returns IsNight (virtual field)
	GetIsNight() bool
	// GetIsDay returns IsDay (virtual field)
	GetIsDay() bool
	// GetIsVacation returns IsVacation (virtual field)
	GetIsVacation() bool
	// IsSecurityDataEmulatedKeypad is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataEmulatedKeypad()
	// CreateBuilder creates a SecurityDataEmulatedKeypadBuilder
	CreateSecurityDataEmulatedKeypadBuilder() SecurityDataEmulatedKeypadBuilder
}

// _SecurityDataEmulatedKeypad is the data-structure of this message
type _SecurityDataEmulatedKeypad struct {
	SecurityDataContract
	Key byte
}

var _ SecurityDataEmulatedKeypad = (*_SecurityDataEmulatedKeypad)(nil)
var _ SecurityDataRequirements = (*_SecurityDataEmulatedKeypad)(nil)

// NewSecurityDataEmulatedKeypad factory function for _SecurityDataEmulatedKeypad
func NewSecurityDataEmulatedKeypad(commandTypeContainer SecurityCommandTypeContainer, argument byte, key byte) *_SecurityDataEmulatedKeypad {
	_result := &_SecurityDataEmulatedKeypad{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		Key:                  key,
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataEmulatedKeypadBuilder is a builder for SecurityDataEmulatedKeypad
type SecurityDataEmulatedKeypadBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(key byte) SecurityDataEmulatedKeypadBuilder
	// WithKey adds Key (property field)
	WithKey(byte) SecurityDataEmulatedKeypadBuilder
	// Build builds the SecurityDataEmulatedKeypad or returns an error if something is wrong
	Build() (SecurityDataEmulatedKeypad, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataEmulatedKeypad
}

// NewSecurityDataEmulatedKeypadBuilder() creates a SecurityDataEmulatedKeypadBuilder
func NewSecurityDataEmulatedKeypadBuilder() SecurityDataEmulatedKeypadBuilder {
	return &_SecurityDataEmulatedKeypadBuilder{_SecurityDataEmulatedKeypad: new(_SecurityDataEmulatedKeypad)}
}

type _SecurityDataEmulatedKeypadBuilder struct {
	*_SecurityDataEmulatedKeypad

	err *utils.MultiError
}

var _ (SecurityDataEmulatedKeypadBuilder) = (*_SecurityDataEmulatedKeypadBuilder)(nil)

func (m *_SecurityDataEmulatedKeypadBuilder) WithMandatoryFields(key byte) SecurityDataEmulatedKeypadBuilder {
	return m.WithKey(key)
}

func (m *_SecurityDataEmulatedKeypadBuilder) WithKey(key byte) SecurityDataEmulatedKeypadBuilder {
	m.Key = key
	return m
}

func (m *_SecurityDataEmulatedKeypadBuilder) Build() (SecurityDataEmulatedKeypad, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SecurityDataEmulatedKeypad.deepCopy(), nil
}

func (m *_SecurityDataEmulatedKeypadBuilder) MustBuild() SecurityDataEmulatedKeypad {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SecurityDataEmulatedKeypadBuilder) DeepCopy() any {
	return m.CreateSecurityDataEmulatedKeypadBuilder()
}

// CreateSecurityDataEmulatedKeypadBuilder creates a SecurityDataEmulatedKeypadBuilder
func (m *_SecurityDataEmulatedKeypad) CreateSecurityDataEmulatedKeypadBuilder() SecurityDataEmulatedKeypadBuilder {
	if m == nil {
		return NewSecurityDataEmulatedKeypadBuilder()
	}
	return &_SecurityDataEmulatedKeypadBuilder{_SecurityDataEmulatedKeypad: m.deepCopy()}
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

func (m *_SecurityDataEmulatedKeypad) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataEmulatedKeypad) GetKey() byte {
	return m.Key
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_SecurityDataEmulatedKeypad) GetIsAscii() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(bool((m.GetKey()) >= (0x00))) && bool(bool((m.GetKey()) <= (0x7F))))
}

func (m *_SecurityDataEmulatedKeypad) GetIsCustom() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) >= (0x80)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsEnter() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x0D)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsShift() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x80)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsPanic() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x81)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsFire() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x82)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsARM() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x83)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsAway() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x84)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsNight() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x85)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsDay() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x86)))
}

func (m *_SecurityDataEmulatedKeypad) GetIsVacation() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetKey()) == (0x87)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSecurityDataEmulatedKeypad(structType any) SecurityDataEmulatedKeypad {
	if casted, ok := structType.(SecurityDataEmulatedKeypad); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataEmulatedKeypad); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataEmulatedKeypad) GetTypeName() string {
	return "SecurityDataEmulatedKeypad"
}

func (m *_SecurityDataEmulatedKeypad) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Simple field (key)
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

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_SecurityDataEmulatedKeypad) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataEmulatedKeypad) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataEmulatedKeypad SecurityDataEmulatedKeypad, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataEmulatedKeypad"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataEmulatedKeypad")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	key, err := ReadSimpleField(ctx, "key", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'key' field"))
	}
	m.Key = key

	isAscii, err := ReadVirtualField[bool](ctx, "isAscii", (*bool)(nil), bool(bool((key) >= (0x00))) && bool(bool((key) <= (0x7F))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isAscii' field"))
	}
	_ = isAscii

	isCustom, err := ReadVirtualField[bool](ctx, "isCustom", (*bool)(nil), bool((key) >= (0x80)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isCustom' field"))
	}
	_ = isCustom

	isEnter, err := ReadVirtualField[bool](ctx, "isEnter", (*bool)(nil), bool((key) == (0x0D)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isEnter' field"))
	}
	_ = isEnter

	isShift, err := ReadVirtualField[bool](ctx, "isShift", (*bool)(nil), bool((key) == (0x80)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isShift' field"))
	}
	_ = isShift

	isPanic, err := ReadVirtualField[bool](ctx, "isPanic", (*bool)(nil), bool((key) == (0x81)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isPanic' field"))
	}
	_ = isPanic

	isFire, err := ReadVirtualField[bool](ctx, "isFire", (*bool)(nil), bool((key) == (0x82)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isFire' field"))
	}
	_ = isFire

	isARM, err := ReadVirtualField[bool](ctx, "isARM", (*bool)(nil), bool((key) == (0x83)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isARM' field"))
	}
	_ = isARM

	isAway, err := ReadVirtualField[bool](ctx, "isAway", (*bool)(nil), bool((key) == (0x84)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isAway' field"))
	}
	_ = isAway

	isNight, err := ReadVirtualField[bool](ctx, "isNight", (*bool)(nil), bool((key) == (0x85)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isNight' field"))
	}
	_ = isNight

	isDay, err := ReadVirtualField[bool](ctx, "isDay", (*bool)(nil), bool((key) == (0x86)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isDay' field"))
	}
	_ = isDay

	isVacation, err := ReadVirtualField[bool](ctx, "isVacation", (*bool)(nil), bool((key) == (0x87)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isVacation' field"))
	}
	_ = isVacation

	if closeErr := readBuffer.CloseContext("SecurityDataEmulatedKeypad"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataEmulatedKeypad")
	}

	return m, nil
}

func (m *_SecurityDataEmulatedKeypad) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataEmulatedKeypad) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataEmulatedKeypad"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataEmulatedKeypad")
		}

		if err := WriteSimpleField[byte](ctx, "key", m.GetKey(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'key' field")
		}
		// Virtual field
		isAscii := m.GetIsAscii()
		_ = isAscii
		if _isAsciiErr := writeBuffer.WriteVirtual(ctx, "isAscii", m.GetIsAscii()); _isAsciiErr != nil {
			return errors.Wrap(_isAsciiErr, "Error serializing 'isAscii' field")
		}
		// Virtual field
		isCustom := m.GetIsCustom()
		_ = isCustom
		if _isCustomErr := writeBuffer.WriteVirtual(ctx, "isCustom", m.GetIsCustom()); _isCustomErr != nil {
			return errors.Wrap(_isCustomErr, "Error serializing 'isCustom' field")
		}
		// Virtual field
		isEnter := m.GetIsEnter()
		_ = isEnter
		if _isEnterErr := writeBuffer.WriteVirtual(ctx, "isEnter", m.GetIsEnter()); _isEnterErr != nil {
			return errors.Wrap(_isEnterErr, "Error serializing 'isEnter' field")
		}
		// Virtual field
		isShift := m.GetIsShift()
		_ = isShift
		if _isShiftErr := writeBuffer.WriteVirtual(ctx, "isShift", m.GetIsShift()); _isShiftErr != nil {
			return errors.Wrap(_isShiftErr, "Error serializing 'isShift' field")
		}
		// Virtual field
		isPanic := m.GetIsPanic()
		_ = isPanic
		if _isPanicErr := writeBuffer.WriteVirtual(ctx, "isPanic", m.GetIsPanic()); _isPanicErr != nil {
			return errors.Wrap(_isPanicErr, "Error serializing 'isPanic' field")
		}
		// Virtual field
		isFire := m.GetIsFire()
		_ = isFire
		if _isFireErr := writeBuffer.WriteVirtual(ctx, "isFire", m.GetIsFire()); _isFireErr != nil {
			return errors.Wrap(_isFireErr, "Error serializing 'isFire' field")
		}
		// Virtual field
		isARM := m.GetIsARM()
		_ = isARM
		if _isARMErr := writeBuffer.WriteVirtual(ctx, "isARM", m.GetIsARM()); _isARMErr != nil {
			return errors.Wrap(_isARMErr, "Error serializing 'isARM' field")
		}
		// Virtual field
		isAway := m.GetIsAway()
		_ = isAway
		if _isAwayErr := writeBuffer.WriteVirtual(ctx, "isAway", m.GetIsAway()); _isAwayErr != nil {
			return errors.Wrap(_isAwayErr, "Error serializing 'isAway' field")
		}
		// Virtual field
		isNight := m.GetIsNight()
		_ = isNight
		if _isNightErr := writeBuffer.WriteVirtual(ctx, "isNight", m.GetIsNight()); _isNightErr != nil {
			return errors.Wrap(_isNightErr, "Error serializing 'isNight' field")
		}
		// Virtual field
		isDay := m.GetIsDay()
		_ = isDay
		if _isDayErr := writeBuffer.WriteVirtual(ctx, "isDay", m.GetIsDay()); _isDayErr != nil {
			return errors.Wrap(_isDayErr, "Error serializing 'isDay' field")
		}
		// Virtual field
		isVacation := m.GetIsVacation()
		_ = isVacation
		if _isVacationErr := writeBuffer.WriteVirtual(ctx, "isVacation", m.GetIsVacation()); _isVacationErr != nil {
			return errors.Wrap(_isVacationErr, "Error serializing 'isVacation' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataEmulatedKeypad"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataEmulatedKeypad")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataEmulatedKeypad) IsSecurityDataEmulatedKeypad() {}

func (m *_SecurityDataEmulatedKeypad) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataEmulatedKeypad) deepCopy() *_SecurityDataEmulatedKeypad {
	if m == nil {
		return nil
	}
	_SecurityDataEmulatedKeypadCopy := &_SecurityDataEmulatedKeypad{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
		m.Key,
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataEmulatedKeypadCopy
}

func (m *_SecurityDataEmulatedKeypad) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
