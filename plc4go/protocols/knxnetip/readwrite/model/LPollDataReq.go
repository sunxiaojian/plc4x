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

// LPollDataReq is the corresponding interface of LPollDataReq
type LPollDataReq interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// IsLPollDataReq is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLPollDataReq()
	// CreateBuilder creates a LPollDataReqBuilder
	CreateLPollDataReqBuilder() LPollDataReqBuilder
}

// _LPollDataReq is the data-structure of this message
type _LPollDataReq struct {
	CEMIContract
}

var _ LPollDataReq = (*_LPollDataReq)(nil)
var _ CEMIRequirements = (*_LPollDataReq)(nil)

// NewLPollDataReq factory function for _LPollDataReq
func NewLPollDataReq(size uint16) *_LPollDataReq {
	_result := &_LPollDataReq{
		CEMIContract: NewCEMI(size),
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// LPollDataReqBuilder is a builder for LPollDataReq
type LPollDataReqBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() LPollDataReqBuilder
	// Build builds the LPollDataReq or returns an error if something is wrong
	Build() (LPollDataReq, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() LPollDataReq
}

// NewLPollDataReqBuilder() creates a LPollDataReqBuilder
func NewLPollDataReqBuilder() LPollDataReqBuilder {
	return &_LPollDataReqBuilder{_LPollDataReq: new(_LPollDataReq)}
}

type _LPollDataReqBuilder struct {
	*_LPollDataReq

	err *utils.MultiError
}

var _ (LPollDataReqBuilder) = (*_LPollDataReqBuilder)(nil)

func (m *_LPollDataReqBuilder) WithMandatoryFields() LPollDataReqBuilder {
	return m
}

func (m *_LPollDataReqBuilder) Build() (LPollDataReq, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._LPollDataReq.deepCopy(), nil
}

func (m *_LPollDataReqBuilder) MustBuild() LPollDataReq {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_LPollDataReqBuilder) DeepCopy() any {
	return m.CreateLPollDataReqBuilder()
}

// CreateLPollDataReqBuilder creates a LPollDataReqBuilder
func (m *_LPollDataReq) CreateLPollDataReqBuilder() LPollDataReqBuilder {
	if m == nil {
		return NewLPollDataReqBuilder()
	}
	return &_LPollDataReqBuilder{_LPollDataReq: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LPollDataReq) GetMessageCode() uint8 {
	return 0x13
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LPollDataReq) GetParent() CEMIContract {
	return m.CEMIContract
}

// Deprecated: use the interface for direct cast
func CastLPollDataReq(structType any) LPollDataReq {
	if casted, ok := structType.(LPollDataReq); ok {
		return casted
	}
	if casted, ok := structType.(*LPollDataReq); ok {
		return *casted
	}
	return nil
}

func (m *_LPollDataReq) GetTypeName() string {
	return "LPollDataReq"
}

func (m *_LPollDataReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_LPollDataReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LPollDataReq) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__lPollDataReq LPollDataReq, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LPollDataReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LPollDataReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("LPollDataReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LPollDataReq")
	}

	return m, nil
}

func (m *_LPollDataReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LPollDataReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LPollDataReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LPollDataReq")
		}

		if popErr := writeBuffer.PopContext("LPollDataReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LPollDataReq")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LPollDataReq) IsLPollDataReq() {}

func (m *_LPollDataReq) DeepCopy() any {
	return m.deepCopy()
}

func (m *_LPollDataReq) deepCopy() *_LPollDataReq {
	if m == nil {
		return nil
	}
	_LPollDataReqCopy := &_LPollDataReq{
		m.CEMIContract.(*_CEMI).deepCopy(),
	}
	m.CEMIContract.(*_CEMI)._SubType = m
	return _LPollDataReqCopy
}

func (m *_LPollDataReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
