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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// SecurityDataRaiseAlarm is the corresponding interface of SecurityDataRaiseAlarm
type SecurityDataRaiseAlarm interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataRaiseAlarm is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataRaiseAlarm()
	// CreateBuilder creates a SecurityDataRaiseAlarmBuilder
	CreateSecurityDataRaiseAlarmBuilder() SecurityDataRaiseAlarmBuilder
}

// _SecurityDataRaiseAlarm is the data-structure of this message
type _SecurityDataRaiseAlarm struct {
	SecurityDataContract
}

var _ SecurityDataRaiseAlarm = (*_SecurityDataRaiseAlarm)(nil)
var _ SecurityDataRequirements = (*_SecurityDataRaiseAlarm)(nil)

// NewSecurityDataRaiseAlarm factory function for _SecurityDataRaiseAlarm
func NewSecurityDataRaiseAlarm(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataRaiseAlarm {
	_result := &_SecurityDataRaiseAlarm{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataRaiseAlarmBuilder is a builder for SecurityDataRaiseAlarm
type SecurityDataRaiseAlarmBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SecurityDataRaiseAlarmBuilder
	// Build builds the SecurityDataRaiseAlarm or returns an error if something is wrong
	Build() (SecurityDataRaiseAlarm, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataRaiseAlarm
}

// NewSecurityDataRaiseAlarmBuilder() creates a SecurityDataRaiseAlarmBuilder
func NewSecurityDataRaiseAlarmBuilder() SecurityDataRaiseAlarmBuilder {
	return &_SecurityDataRaiseAlarmBuilder{_SecurityDataRaiseAlarm: new(_SecurityDataRaiseAlarm)}
}

type _SecurityDataRaiseAlarmBuilder struct {
	*_SecurityDataRaiseAlarm

	err *utils.MultiError
}

var _ (SecurityDataRaiseAlarmBuilder) = (*_SecurityDataRaiseAlarmBuilder)(nil)

func (m *_SecurityDataRaiseAlarmBuilder) WithMandatoryFields() SecurityDataRaiseAlarmBuilder {
	return m
}

func (m *_SecurityDataRaiseAlarmBuilder) Build() (SecurityDataRaiseAlarm, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SecurityDataRaiseAlarm.deepCopy(), nil
}

func (m *_SecurityDataRaiseAlarmBuilder) MustBuild() SecurityDataRaiseAlarm {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SecurityDataRaiseAlarmBuilder) DeepCopy() any {
	return m.CreateSecurityDataRaiseAlarmBuilder()
}

// CreateSecurityDataRaiseAlarmBuilder creates a SecurityDataRaiseAlarmBuilder
func (m *_SecurityDataRaiseAlarm) CreateSecurityDataRaiseAlarmBuilder() SecurityDataRaiseAlarmBuilder {
	if m == nil {
		return NewSecurityDataRaiseAlarmBuilder()
	}
	return &_SecurityDataRaiseAlarmBuilder{_SecurityDataRaiseAlarm: m.deepCopy()}
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

func (m *_SecurityDataRaiseAlarm) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataRaiseAlarm(structType any) SecurityDataRaiseAlarm {
	if casted, ok := structType.(SecurityDataRaiseAlarm); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataRaiseAlarm); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataRaiseAlarm) GetTypeName() string {
	return "SecurityDataRaiseAlarm"
}

func (m *_SecurityDataRaiseAlarm) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataRaiseAlarm) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataRaiseAlarm) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataRaiseAlarm SecurityDataRaiseAlarm, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataRaiseAlarm"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataRaiseAlarm")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataRaiseAlarm"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataRaiseAlarm")
	}

	return m, nil
}

func (m *_SecurityDataRaiseAlarm) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataRaiseAlarm) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataRaiseAlarm"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataRaiseAlarm")
		}

		if popErr := writeBuffer.PopContext("SecurityDataRaiseAlarm"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataRaiseAlarm")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataRaiseAlarm) IsSecurityDataRaiseAlarm() {}

func (m *_SecurityDataRaiseAlarm) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataRaiseAlarm) deepCopy() *_SecurityDataRaiseAlarm {
	if m == nil {
		return nil
	}
	_SecurityDataRaiseAlarmCopy := &_SecurityDataRaiseAlarm{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataRaiseAlarmCopy
}

func (m *_SecurityDataRaiseAlarm) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
