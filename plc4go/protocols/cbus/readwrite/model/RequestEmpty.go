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

// RequestEmpty is the corresponding interface of RequestEmpty
type RequestEmpty interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Request
	// IsRequestEmpty is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRequestEmpty()
	// CreateBuilder creates a RequestEmptyBuilder
	CreateRequestEmptyBuilder() RequestEmptyBuilder
}

// _RequestEmpty is the data-structure of this message
type _RequestEmpty struct {
	RequestContract
}

var _ RequestEmpty = (*_RequestEmpty)(nil)
var _ RequestRequirements = (*_RequestEmpty)(nil)

// NewRequestEmpty factory function for _RequestEmpty
func NewRequestEmpty(peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination, cBusOptions CBusOptions) *_RequestEmpty {
	_result := &_RequestEmpty{
		RequestContract: NewRequest(peekedByte, startingCR, resetMode, secondPeek, termination, cBusOptions),
	}
	_result.RequestContract.(*_Request)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// RequestEmptyBuilder is a builder for RequestEmpty
type RequestEmptyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() RequestEmptyBuilder
	// Build builds the RequestEmpty or returns an error if something is wrong
	Build() (RequestEmpty, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() RequestEmpty
}

// NewRequestEmptyBuilder() creates a RequestEmptyBuilder
func NewRequestEmptyBuilder() RequestEmptyBuilder {
	return &_RequestEmptyBuilder{_RequestEmpty: new(_RequestEmpty)}
}

type _RequestEmptyBuilder struct {
	*_RequestEmpty

	err *utils.MultiError
}

var _ (RequestEmptyBuilder) = (*_RequestEmptyBuilder)(nil)

func (m *_RequestEmptyBuilder) WithMandatoryFields() RequestEmptyBuilder {
	return m
}

func (m *_RequestEmptyBuilder) Build() (RequestEmpty, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._RequestEmpty.deepCopy(), nil
}

func (m *_RequestEmptyBuilder) MustBuild() RequestEmpty {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_RequestEmptyBuilder) DeepCopy() any {
	return m.CreateRequestEmptyBuilder()
}

// CreateRequestEmptyBuilder creates a RequestEmptyBuilder
func (m *_RequestEmpty) CreateRequestEmptyBuilder() RequestEmptyBuilder {
	if m == nil {
		return NewRequestEmptyBuilder()
	}
	return &_RequestEmptyBuilder{_RequestEmpty: m.deepCopy()}
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

func (m *_RequestEmpty) GetParent() RequestContract {
	return m.RequestContract
}

// Deprecated: use the interface for direct cast
func CastRequestEmpty(structType any) RequestEmpty {
	if casted, ok := structType.(RequestEmpty); ok {
		return casted
	}
	if casted, ok := structType.(*RequestEmpty); ok {
		return *casted
	}
	return nil
}

func (m *_RequestEmpty) GetTypeName() string {
	return "RequestEmpty"
}

func (m *_RequestEmpty) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.RequestContract.(*_Request).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_RequestEmpty) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RequestEmpty) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Request, cBusOptions CBusOptions) (__requestEmpty RequestEmpty, err error) {
	m.RequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestEmpty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestEmpty")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("RequestEmpty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestEmpty")
	}

	return m, nil
}

func (m *_RequestEmpty) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RequestEmpty) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RequestEmpty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RequestEmpty")
		}

		if popErr := writeBuffer.PopContext("RequestEmpty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RequestEmpty")
		}
		return nil
	}
	return m.RequestContract.(*_Request).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RequestEmpty) IsRequestEmpty() {}

func (m *_RequestEmpty) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RequestEmpty) deepCopy() *_RequestEmpty {
	if m == nil {
		return nil
	}
	_RequestEmptyCopy := &_RequestEmpty{
		m.RequestContract.(*_Request).deepCopy(),
	}
	m.RequestContract.(*_Request)._SubType = m
	return _RequestEmptyCopy
}

func (m *_RequestEmpty) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
