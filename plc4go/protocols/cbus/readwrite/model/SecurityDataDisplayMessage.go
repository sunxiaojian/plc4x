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

// SecurityDataDisplayMessage is the corresponding interface of SecurityDataDisplayMessage
type SecurityDataDisplayMessage interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// GetMessage returns Message (property field)
	GetMessage() string
	// IsSecurityDataDisplayMessage is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataDisplayMessage()
	// CreateBuilder creates a SecurityDataDisplayMessageBuilder
	CreateSecurityDataDisplayMessageBuilder() SecurityDataDisplayMessageBuilder
}

// _SecurityDataDisplayMessage is the data-structure of this message
type _SecurityDataDisplayMessage struct {
	SecurityDataContract
	Message string
}

var _ SecurityDataDisplayMessage = (*_SecurityDataDisplayMessage)(nil)
var _ SecurityDataRequirements = (*_SecurityDataDisplayMessage)(nil)

// NewSecurityDataDisplayMessage factory function for _SecurityDataDisplayMessage
func NewSecurityDataDisplayMessage(commandTypeContainer SecurityCommandTypeContainer, argument byte, message string) *_SecurityDataDisplayMessage {
	_result := &_SecurityDataDisplayMessage{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		Message:              message,
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataDisplayMessageBuilder is a builder for SecurityDataDisplayMessage
type SecurityDataDisplayMessageBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(message string) SecurityDataDisplayMessageBuilder
	// WithMessage adds Message (property field)
	WithMessage(string) SecurityDataDisplayMessageBuilder
	// Build builds the SecurityDataDisplayMessage or returns an error if something is wrong
	Build() (SecurityDataDisplayMessage, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataDisplayMessage
}

// NewSecurityDataDisplayMessageBuilder() creates a SecurityDataDisplayMessageBuilder
func NewSecurityDataDisplayMessageBuilder() SecurityDataDisplayMessageBuilder {
	return &_SecurityDataDisplayMessageBuilder{_SecurityDataDisplayMessage: new(_SecurityDataDisplayMessage)}
}

type _SecurityDataDisplayMessageBuilder struct {
	*_SecurityDataDisplayMessage

	err *utils.MultiError
}

var _ (SecurityDataDisplayMessageBuilder) = (*_SecurityDataDisplayMessageBuilder)(nil)

func (m *_SecurityDataDisplayMessageBuilder) WithMandatoryFields(message string) SecurityDataDisplayMessageBuilder {
	return m.WithMessage(message)
}

func (m *_SecurityDataDisplayMessageBuilder) WithMessage(message string) SecurityDataDisplayMessageBuilder {
	m.Message = message
	return m
}

func (m *_SecurityDataDisplayMessageBuilder) Build() (SecurityDataDisplayMessage, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SecurityDataDisplayMessage.deepCopy(), nil
}

func (m *_SecurityDataDisplayMessageBuilder) MustBuild() SecurityDataDisplayMessage {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SecurityDataDisplayMessageBuilder) DeepCopy() any {
	return m.CreateSecurityDataDisplayMessageBuilder()
}

// CreateSecurityDataDisplayMessageBuilder creates a SecurityDataDisplayMessageBuilder
func (m *_SecurityDataDisplayMessage) CreateSecurityDataDisplayMessageBuilder() SecurityDataDisplayMessageBuilder {
	if m == nil {
		return NewSecurityDataDisplayMessageBuilder()
	}
	return &_SecurityDataDisplayMessageBuilder{_SecurityDataDisplayMessage: m.deepCopy()}
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

func (m *_SecurityDataDisplayMessage) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataDisplayMessage) GetMessage() string {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSecurityDataDisplayMessage(structType any) SecurityDataDisplayMessage {
	if casted, ok := structType.(SecurityDataDisplayMessage); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataDisplayMessage); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataDisplayMessage) GetTypeName() string {
	return "SecurityDataDisplayMessage"
}

func (m *_SecurityDataDisplayMessage) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Simple field (message)
	lengthInBits += uint16(int32((int32(m.GetCommandTypeContainer().NumBytes()) - int32(int32(1)))) * int32(int32(8)))

	return lengthInBits
}

func (m *_SecurityDataDisplayMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataDisplayMessage) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData, commandTypeContainer SecurityCommandTypeContainer) (__securityDataDisplayMessage SecurityDataDisplayMessage, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataDisplayMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataDisplayMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	message, err := ReadSimpleField(ctx, "message", ReadString(readBuffer, uint32(int32((int32(commandTypeContainer.NumBytes())-int32(int32(1))))*int32(int32(8)))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'message' field"))
	}
	m.Message = message

	if closeErr := readBuffer.CloseContext("SecurityDataDisplayMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataDisplayMessage")
	}

	return m, nil
}

func (m *_SecurityDataDisplayMessage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataDisplayMessage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataDisplayMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataDisplayMessage")
		}

		if err := WriteSimpleField[string](ctx, "message", m.GetMessage(), WriteString(writeBuffer, int32(int32((int32(m.GetCommandTypeContainer().NumBytes())-int32(int32(1))))*int32(int32(8))))); err != nil {
			return errors.Wrap(err, "Error serializing 'message' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataDisplayMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataDisplayMessage")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataDisplayMessage) IsSecurityDataDisplayMessage() {}

func (m *_SecurityDataDisplayMessage) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataDisplayMessage) deepCopy() *_SecurityDataDisplayMessage {
	if m == nil {
		return nil
	}
	_SecurityDataDisplayMessageCopy := &_SecurityDataDisplayMessage{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
		m.Message,
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataDisplayMessageCopy
}

func (m *_SecurityDataDisplayMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
