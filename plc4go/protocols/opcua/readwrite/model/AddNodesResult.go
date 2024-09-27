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

// AddNodesResult is the corresponding interface of AddNodesResult
type AddNodesResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetAddedNodeId returns AddedNodeId (property field)
	GetAddedNodeId() NodeId
	// IsAddNodesResult is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAddNodesResult()
	// CreateBuilder creates a AddNodesResultBuilder
	CreateAddNodesResultBuilder() AddNodesResultBuilder
}

// _AddNodesResult is the data-structure of this message
type _AddNodesResult struct {
	ExtensionObjectDefinitionContract
	StatusCode  StatusCode
	AddedNodeId NodeId
}

var _ AddNodesResult = (*_AddNodesResult)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_AddNodesResult)(nil)

// NewAddNodesResult factory function for _AddNodesResult
func NewAddNodesResult(statusCode StatusCode, addedNodeId NodeId) *_AddNodesResult {
	if statusCode == nil {
		panic("statusCode of type StatusCode for AddNodesResult must not be nil")
	}
	if addedNodeId == nil {
		panic("addedNodeId of type NodeId for AddNodesResult must not be nil")
	}
	_result := &_AddNodesResult{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		StatusCode:                        statusCode,
		AddedNodeId:                       addedNodeId,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AddNodesResultBuilder is a builder for AddNodesResult
type AddNodesResultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(statusCode StatusCode, addedNodeId NodeId) AddNodesResultBuilder
	// WithStatusCode adds StatusCode (property field)
	WithStatusCode(StatusCode) AddNodesResultBuilder
	// WithStatusCodeBuilder adds StatusCode (property field) which is build by the builder
	WithStatusCodeBuilder(func(StatusCodeBuilder) StatusCodeBuilder) AddNodesResultBuilder
	// WithAddedNodeId adds AddedNodeId (property field)
	WithAddedNodeId(NodeId) AddNodesResultBuilder
	// WithAddedNodeIdBuilder adds AddedNodeId (property field) which is build by the builder
	WithAddedNodeIdBuilder(func(NodeIdBuilder) NodeIdBuilder) AddNodesResultBuilder
	// Build builds the AddNodesResult or returns an error if something is wrong
	Build() (AddNodesResult, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AddNodesResult
}

// NewAddNodesResultBuilder() creates a AddNodesResultBuilder
func NewAddNodesResultBuilder() AddNodesResultBuilder {
	return &_AddNodesResultBuilder{_AddNodesResult: new(_AddNodesResult)}
}

type _AddNodesResultBuilder struct {
	*_AddNodesResult

	err *utils.MultiError
}

var _ (AddNodesResultBuilder) = (*_AddNodesResultBuilder)(nil)

func (m *_AddNodesResultBuilder) WithMandatoryFields(statusCode StatusCode, addedNodeId NodeId) AddNodesResultBuilder {
	return m.WithStatusCode(statusCode).WithAddedNodeId(addedNodeId)
}

func (m *_AddNodesResultBuilder) WithStatusCode(statusCode StatusCode) AddNodesResultBuilder {
	m.StatusCode = statusCode
	return m
}

func (m *_AddNodesResultBuilder) WithStatusCodeBuilder(builderSupplier func(StatusCodeBuilder) StatusCodeBuilder) AddNodesResultBuilder {
	builder := builderSupplier(m.StatusCode.CreateStatusCodeBuilder())
	var err error
	m.StatusCode, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "StatusCodeBuilder failed"))
	}
	return m
}

func (m *_AddNodesResultBuilder) WithAddedNodeId(addedNodeId NodeId) AddNodesResultBuilder {
	m.AddedNodeId = addedNodeId
	return m
}

func (m *_AddNodesResultBuilder) WithAddedNodeIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) AddNodesResultBuilder {
	builder := builderSupplier(m.AddedNodeId.CreateNodeIdBuilder())
	var err error
	m.AddedNodeId, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return m
}

func (m *_AddNodesResultBuilder) Build() (AddNodesResult, error) {
	if m.StatusCode == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'statusCode' not set"))
	}
	if m.AddedNodeId == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'addedNodeId' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._AddNodesResult.deepCopy(), nil
}

func (m *_AddNodesResultBuilder) MustBuild() AddNodesResult {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_AddNodesResultBuilder) DeepCopy() any {
	return m.CreateAddNodesResultBuilder()
}

// CreateAddNodesResultBuilder creates a AddNodesResultBuilder
func (m *_AddNodesResult) CreateAddNodesResultBuilder() AddNodesResultBuilder {
	if m == nil {
		return NewAddNodesResultBuilder()
	}
	return &_AddNodesResultBuilder{_AddNodesResult: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AddNodesResult) GetIdentifier() string {
	return "485"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AddNodesResult) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AddNodesResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_AddNodesResult) GetAddedNodeId() NodeId {
	return m.AddedNodeId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAddNodesResult(structType any) AddNodesResult {
	if casted, ok := structType.(AddNodesResult); ok {
		return casted
	}
	if casted, ok := structType.(*AddNodesResult); ok {
		return *casted
	}
	return nil
}

func (m *_AddNodesResult) GetTypeName() string {
	return "AddNodesResult"
}

func (m *_AddNodesResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (addedNodeId)
	lengthInBits += m.AddedNodeId.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AddNodesResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AddNodesResult) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__addNodesResult AddNodesResult, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AddNodesResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AddNodesResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusCode, err := ReadSimpleField[StatusCode](ctx, "statusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusCode' field"))
	}
	m.StatusCode = statusCode

	addedNodeId, err := ReadSimpleField[NodeId](ctx, "addedNodeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'addedNodeId' field"))
	}
	m.AddedNodeId = addedNodeId

	if closeErr := readBuffer.CloseContext("AddNodesResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AddNodesResult")
	}

	return m, nil
}

func (m *_AddNodesResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AddNodesResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AddNodesResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AddNodesResult")
		}

		if err := WriteSimpleField[StatusCode](ctx, "statusCode", m.GetStatusCode(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusCode' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "addedNodeId", m.GetAddedNodeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'addedNodeId' field")
		}

		if popErr := writeBuffer.PopContext("AddNodesResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AddNodesResult")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AddNodesResult) IsAddNodesResult() {}

func (m *_AddNodesResult) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AddNodesResult) deepCopy() *_AddNodesResult {
	if m == nil {
		return nil
	}
	_AddNodesResultCopy := &_AddNodesResult{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.StatusCode.DeepCopy().(StatusCode),
		m.AddedNodeId.DeepCopy().(NodeId),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _AddNodesResultCopy
}

func (m *_AddNodesResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
