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

// WriteRequest is the corresponding interface of WriteRequest
type WriteRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfNodesToWrite returns NoOfNodesToWrite (property field)
	GetNoOfNodesToWrite() int32
	// GetNodesToWrite returns NodesToWrite (property field)
	GetNodesToWrite() []ExtensionObjectDefinition
	// IsWriteRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsWriteRequest()
	// CreateBuilder creates a WriteRequestBuilder
	CreateWriteRequestBuilder() WriteRequestBuilder
}

// _WriteRequest is the data-structure of this message
type _WriteRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader    ExtensionObjectDefinition
	NoOfNodesToWrite int32
	NodesToWrite     []ExtensionObjectDefinition
}

var _ WriteRequest = (*_WriteRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_WriteRequest)(nil)

// NewWriteRequest factory function for _WriteRequest
func NewWriteRequest(requestHeader ExtensionObjectDefinition, noOfNodesToWrite int32, nodesToWrite []ExtensionObjectDefinition) *_WriteRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for WriteRequest must not be nil")
	}
	_result := &_WriteRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		NoOfNodesToWrite:                  noOfNodesToWrite,
		NodesToWrite:                      nodesToWrite,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// WriteRequestBuilder is a builder for WriteRequest
type WriteRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader ExtensionObjectDefinition, noOfNodesToWrite int32, nodesToWrite []ExtensionObjectDefinition) WriteRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(ExtensionObjectDefinition) WriteRequestBuilder
	// WithNoOfNodesToWrite adds NoOfNodesToWrite (property field)
	WithNoOfNodesToWrite(int32) WriteRequestBuilder
	// WithNodesToWrite adds NodesToWrite (property field)
	WithNodesToWrite(...ExtensionObjectDefinition) WriteRequestBuilder
	// Build builds the WriteRequest or returns an error if something is wrong
	Build() (WriteRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() WriteRequest
}

// NewWriteRequestBuilder() creates a WriteRequestBuilder
func NewWriteRequestBuilder() WriteRequestBuilder {
	return &_WriteRequestBuilder{_WriteRequest: new(_WriteRequest)}
}

type _WriteRequestBuilder struct {
	*_WriteRequest

	err *utils.MultiError
}

var _ (WriteRequestBuilder) = (*_WriteRequestBuilder)(nil)

func (m *_WriteRequestBuilder) WithMandatoryFields(requestHeader ExtensionObjectDefinition, noOfNodesToWrite int32, nodesToWrite []ExtensionObjectDefinition) WriteRequestBuilder {
	return m.WithRequestHeader(requestHeader).WithNoOfNodesToWrite(noOfNodesToWrite).WithNodesToWrite(nodesToWrite...)
}

func (m *_WriteRequestBuilder) WithRequestHeader(requestHeader ExtensionObjectDefinition) WriteRequestBuilder {
	m.RequestHeader = requestHeader
	return m
}

func (m *_WriteRequestBuilder) WithNoOfNodesToWrite(noOfNodesToWrite int32) WriteRequestBuilder {
	m.NoOfNodesToWrite = noOfNodesToWrite
	return m
}

func (m *_WriteRequestBuilder) WithNodesToWrite(nodesToWrite ...ExtensionObjectDefinition) WriteRequestBuilder {
	m.NodesToWrite = nodesToWrite
	return m
}

func (m *_WriteRequestBuilder) Build() (WriteRequest, error) {
	if m.RequestHeader == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._WriteRequest.deepCopy(), nil
}

func (m *_WriteRequestBuilder) MustBuild() WriteRequest {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_WriteRequestBuilder) DeepCopy() any {
	return m.CreateWriteRequestBuilder()
}

// CreateWriteRequestBuilder creates a WriteRequestBuilder
func (m *_WriteRequest) CreateWriteRequestBuilder() WriteRequestBuilder {
	if m == nil {
		return NewWriteRequestBuilder()
	}
	return &_WriteRequestBuilder{_WriteRequest: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_WriteRequest) GetIdentifier() string {
	return "673"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_WriteRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_WriteRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_WriteRequest) GetNoOfNodesToWrite() int32 {
	return m.NoOfNodesToWrite
}

func (m *_WriteRequest) GetNodesToWrite() []ExtensionObjectDefinition {
	return m.NodesToWrite
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastWriteRequest(structType any) WriteRequest {
	if casted, ok := structType.(WriteRequest); ok {
		return casted
	}
	if casted, ok := structType.(*WriteRequest); ok {
		return *casted
	}
	return nil
}

func (m *_WriteRequest) GetTypeName() string {
	return "WriteRequest"
}

func (m *_WriteRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfNodesToWrite)
	lengthInBits += 32

	// Array field
	if len(m.NodesToWrite) > 0 {
		for _curItem, element := range m.NodesToWrite {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodesToWrite), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_WriteRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_WriteRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__writeRequest WriteRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("WriteRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for WriteRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	noOfNodesToWrite, err := ReadSimpleField(ctx, "noOfNodesToWrite", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNodesToWrite' field"))
	}
	m.NoOfNodesToWrite = noOfNodesToWrite

	nodesToWrite, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "nodesToWrite", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("670")), readBuffer), uint64(noOfNodesToWrite))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodesToWrite' field"))
	}
	m.NodesToWrite = nodesToWrite

	if closeErr := readBuffer.CloseContext("WriteRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for WriteRequest")
	}

	return m, nil
}

func (m *_WriteRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_WriteRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("WriteRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for WriteRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfNodesToWrite", m.GetNoOfNodesToWrite(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfNodesToWrite' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "nodesToWrite", m.GetNodesToWrite(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'nodesToWrite' field")
		}

		if popErr := writeBuffer.PopContext("WriteRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for WriteRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_WriteRequest) IsWriteRequest() {}

func (m *_WriteRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_WriteRequest) deepCopy() *_WriteRequest {
	if m == nil {
		return nil
	}
	_WriteRequestCopy := &_WriteRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.NoOfNodesToWrite,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.NodesToWrite),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _WriteRequestCopy
}

func (m *_WriteRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
