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

// S7ParameterWriteVarRequest is the corresponding interface of S7ParameterWriteVarRequest
type S7ParameterWriteVarRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7Parameter
	// GetItems returns Items (property field)
	GetItems() []S7VarRequestParameterItem
	// IsS7ParameterWriteVarRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7ParameterWriteVarRequest()
	// CreateBuilder creates a S7ParameterWriteVarRequestBuilder
	CreateS7ParameterWriteVarRequestBuilder() S7ParameterWriteVarRequestBuilder
}

// _S7ParameterWriteVarRequest is the data-structure of this message
type _S7ParameterWriteVarRequest struct {
	S7ParameterContract
	Items []S7VarRequestParameterItem
}

var _ S7ParameterWriteVarRequest = (*_S7ParameterWriteVarRequest)(nil)
var _ S7ParameterRequirements = (*_S7ParameterWriteVarRequest)(nil)

// NewS7ParameterWriteVarRequest factory function for _S7ParameterWriteVarRequest
func NewS7ParameterWriteVarRequest(items []S7VarRequestParameterItem) *_S7ParameterWriteVarRequest {
	_result := &_S7ParameterWriteVarRequest{
		S7ParameterContract: NewS7Parameter(),
		Items:               items,
	}
	_result.S7ParameterContract.(*_S7Parameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7ParameterWriteVarRequestBuilder is a builder for S7ParameterWriteVarRequest
type S7ParameterWriteVarRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(items []S7VarRequestParameterItem) S7ParameterWriteVarRequestBuilder
	// WithItems adds Items (property field)
	WithItems(...S7VarRequestParameterItem) S7ParameterWriteVarRequestBuilder
	// Build builds the S7ParameterWriteVarRequest or returns an error if something is wrong
	Build() (S7ParameterWriteVarRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7ParameterWriteVarRequest
}

// NewS7ParameterWriteVarRequestBuilder() creates a S7ParameterWriteVarRequestBuilder
func NewS7ParameterWriteVarRequestBuilder() S7ParameterWriteVarRequestBuilder {
	return &_S7ParameterWriteVarRequestBuilder{_S7ParameterWriteVarRequest: new(_S7ParameterWriteVarRequest)}
}

type _S7ParameterWriteVarRequestBuilder struct {
	*_S7ParameterWriteVarRequest

	err *utils.MultiError
}

var _ (S7ParameterWriteVarRequestBuilder) = (*_S7ParameterWriteVarRequestBuilder)(nil)

func (m *_S7ParameterWriteVarRequestBuilder) WithMandatoryFields(items []S7VarRequestParameterItem) S7ParameterWriteVarRequestBuilder {
	return m.WithItems(items...)
}

func (m *_S7ParameterWriteVarRequestBuilder) WithItems(items ...S7VarRequestParameterItem) S7ParameterWriteVarRequestBuilder {
	m.Items = items
	return m
}

func (m *_S7ParameterWriteVarRequestBuilder) Build() (S7ParameterWriteVarRequest, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._S7ParameterWriteVarRequest.deepCopy(), nil
}

func (m *_S7ParameterWriteVarRequestBuilder) MustBuild() S7ParameterWriteVarRequest {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_S7ParameterWriteVarRequestBuilder) DeepCopy() any {
	return m.CreateS7ParameterWriteVarRequestBuilder()
}

// CreateS7ParameterWriteVarRequestBuilder creates a S7ParameterWriteVarRequestBuilder
func (m *_S7ParameterWriteVarRequest) CreateS7ParameterWriteVarRequestBuilder() S7ParameterWriteVarRequestBuilder {
	if m == nil {
		return NewS7ParameterWriteVarRequestBuilder()
	}
	return &_S7ParameterWriteVarRequestBuilder{_S7ParameterWriteVarRequest: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7ParameterWriteVarRequest) GetParameterType() uint8 {
	return 0x05
}

func (m *_S7ParameterWriteVarRequest) GetMessageType() uint8 {
	return 0x01
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7ParameterWriteVarRequest) GetParent() S7ParameterContract {
	return m.S7ParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7ParameterWriteVarRequest) GetItems() []S7VarRequestParameterItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7ParameterWriteVarRequest(structType any) S7ParameterWriteVarRequest {
	if casted, ok := structType.(S7ParameterWriteVarRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7ParameterWriteVarRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7ParameterWriteVarRequest) GetTypeName() string {
	return "S7ParameterWriteVarRequest"
}

func (m *_S7ParameterWriteVarRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7ParameterContract.(*_S7Parameter).getLengthInBits(ctx))

	// Implicit Field (numItems)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_S7ParameterWriteVarRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7ParameterWriteVarRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7Parameter, messageType uint8) (__s7ParameterWriteVarRequest S7ParameterWriteVarRequest, err error) {
	m.S7ParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterWriteVarRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7ParameterWriteVarRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numItems, err := ReadImplicitField[uint8](ctx, "numItems", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numItems' field"))
	}
	_ = numItems

	items, err := ReadCountArrayField[S7VarRequestParameterItem](ctx, "items", ReadComplex[S7VarRequestParameterItem](S7VarRequestParameterItemParseWithBuffer, readBuffer), uint64(numItems))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("S7ParameterWriteVarRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7ParameterWriteVarRequest")
	}

	return m, nil
}

func (m *_S7ParameterWriteVarRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7ParameterWriteVarRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterWriteVarRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7ParameterWriteVarRequest")
		}
		numItems := uint8(uint8(len(m.GetItems())))
		if err := WriteImplicitField(ctx, "numItems", numItems, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'numItems' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "items", m.GetItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("S7ParameterWriteVarRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7ParameterWriteVarRequest")
		}
		return nil
	}
	return m.S7ParameterContract.(*_S7Parameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7ParameterWriteVarRequest) IsS7ParameterWriteVarRequest() {}

func (m *_S7ParameterWriteVarRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7ParameterWriteVarRequest) deepCopy() *_S7ParameterWriteVarRequest {
	if m == nil {
		return nil
	}
	_S7ParameterWriteVarRequestCopy := &_S7ParameterWriteVarRequest{
		m.S7ParameterContract.(*_S7Parameter).deepCopy(),
		utils.DeepCopySlice[S7VarRequestParameterItem, S7VarRequestParameterItem](m.Items),
	}
	m.S7ParameterContract.(*_S7Parameter)._SubType = m
	return _S7ParameterWriteVarRequestCopy
}

func (m *_S7ParameterWriteVarRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
