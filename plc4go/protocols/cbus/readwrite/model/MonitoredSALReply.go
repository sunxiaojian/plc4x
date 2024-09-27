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

// MonitoredSALReply is the corresponding interface of MonitoredSALReply
type MonitoredSALReply interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	EncodedReply
	// GetMonitoredSAL returns MonitoredSAL (property field)
	GetMonitoredSAL() MonitoredSAL
	// IsMonitoredSALReply is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMonitoredSALReply()
	// CreateBuilder creates a MonitoredSALReplyBuilder
	CreateMonitoredSALReplyBuilder() MonitoredSALReplyBuilder
}

// _MonitoredSALReply is the data-structure of this message
type _MonitoredSALReply struct {
	EncodedReplyContract
	MonitoredSAL MonitoredSAL
}

var _ MonitoredSALReply = (*_MonitoredSALReply)(nil)
var _ EncodedReplyRequirements = (*_MonitoredSALReply)(nil)

// NewMonitoredSALReply factory function for _MonitoredSALReply
func NewMonitoredSALReply(peekedByte byte, monitoredSAL MonitoredSAL, cBusOptions CBusOptions, requestContext RequestContext) *_MonitoredSALReply {
	if monitoredSAL == nil {
		panic("monitoredSAL of type MonitoredSAL for MonitoredSALReply must not be nil")
	}
	_result := &_MonitoredSALReply{
		EncodedReplyContract: NewEncodedReply(peekedByte, cBusOptions, requestContext),
		MonitoredSAL:         monitoredSAL,
	}
	_result.EncodedReplyContract.(*_EncodedReply)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MonitoredSALReplyBuilder is a builder for MonitoredSALReply
type MonitoredSALReplyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(monitoredSAL MonitoredSAL) MonitoredSALReplyBuilder
	// WithMonitoredSAL adds MonitoredSAL (property field)
	WithMonitoredSAL(MonitoredSAL) MonitoredSALReplyBuilder
	// Build builds the MonitoredSALReply or returns an error if something is wrong
	Build() (MonitoredSALReply, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MonitoredSALReply
}

// NewMonitoredSALReplyBuilder() creates a MonitoredSALReplyBuilder
func NewMonitoredSALReplyBuilder() MonitoredSALReplyBuilder {
	return &_MonitoredSALReplyBuilder{_MonitoredSALReply: new(_MonitoredSALReply)}
}

type _MonitoredSALReplyBuilder struct {
	*_MonitoredSALReply

	err *utils.MultiError
}

var _ (MonitoredSALReplyBuilder) = (*_MonitoredSALReplyBuilder)(nil)

func (m *_MonitoredSALReplyBuilder) WithMandatoryFields(monitoredSAL MonitoredSAL) MonitoredSALReplyBuilder {
	return m.WithMonitoredSAL(monitoredSAL)
}

func (m *_MonitoredSALReplyBuilder) WithMonitoredSAL(monitoredSAL MonitoredSAL) MonitoredSALReplyBuilder {
	m.MonitoredSAL = monitoredSAL
	return m
}

func (m *_MonitoredSALReplyBuilder) Build() (MonitoredSALReply, error) {
	if m.MonitoredSAL == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'monitoredSAL' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._MonitoredSALReply.deepCopy(), nil
}

func (m *_MonitoredSALReplyBuilder) MustBuild() MonitoredSALReply {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_MonitoredSALReplyBuilder) DeepCopy() any {
	return m.CreateMonitoredSALReplyBuilder()
}

// CreateMonitoredSALReplyBuilder creates a MonitoredSALReplyBuilder
func (m *_MonitoredSALReply) CreateMonitoredSALReplyBuilder() MonitoredSALReplyBuilder {
	if m == nil {
		return NewMonitoredSALReplyBuilder()
	}
	return &_MonitoredSALReplyBuilder{_MonitoredSALReply: m.deepCopy()}
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

func (m *_MonitoredSALReply) GetParent() EncodedReplyContract {
	return m.EncodedReplyContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredSALReply) GetMonitoredSAL() MonitoredSAL {
	return m.MonitoredSAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMonitoredSALReply(structType any) MonitoredSALReply {
	if casted, ok := structType.(MonitoredSALReply); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredSALReply); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredSALReply) GetTypeName() string {
	return "MonitoredSALReply"
}

func (m *_MonitoredSALReply) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.EncodedReplyContract.(*_EncodedReply).getLengthInBits(ctx))

	// Simple field (monitoredSAL)
	lengthInBits += m.MonitoredSAL.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_MonitoredSALReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MonitoredSALReply) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_EncodedReply, cBusOptions CBusOptions, requestContext RequestContext) (__monitoredSALReply MonitoredSALReply, err error) {
	m.EncodedReplyContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MonitoredSALReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredSALReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	monitoredSAL, err := ReadSimpleField[MonitoredSAL](ctx, "monitoredSAL", ReadComplex[MonitoredSAL](MonitoredSALParseWithBufferProducer[MonitoredSAL]((CBusOptions)(cBusOptions)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredSAL' field"))
	}
	m.MonitoredSAL = monitoredSAL

	if closeErr := readBuffer.CloseContext("MonitoredSALReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredSALReply")
	}

	return m, nil
}

func (m *_MonitoredSALReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredSALReply) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredSALReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredSALReply")
		}

		if err := WriteSimpleField[MonitoredSAL](ctx, "monitoredSAL", m.GetMonitoredSAL(), WriteComplex[MonitoredSAL](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'monitoredSAL' field")
		}

		if popErr := writeBuffer.PopContext("MonitoredSALReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredSALReply")
		}
		return nil
	}
	return m.EncodedReplyContract.(*_EncodedReply).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredSALReply) IsMonitoredSALReply() {}

func (m *_MonitoredSALReply) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MonitoredSALReply) deepCopy() *_MonitoredSALReply {
	if m == nil {
		return nil
	}
	_MonitoredSALReplyCopy := &_MonitoredSALReply{
		m.EncodedReplyContract.(*_EncodedReply).deepCopy(),
		m.MonitoredSAL.DeepCopy().(MonitoredSAL),
	}
	m.EncodedReplyContract.(*_EncodedReply)._SubType = m
	return _MonitoredSALReplyCopy
}

func (m *_MonitoredSALReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
