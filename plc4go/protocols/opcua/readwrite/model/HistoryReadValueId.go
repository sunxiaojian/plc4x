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

// HistoryReadValueId is the corresponding interface of HistoryReadValueId
type HistoryReadValueId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() NodeId
	// GetIndexRange returns IndexRange (property field)
	GetIndexRange() PascalString
	// GetDataEncoding returns DataEncoding (property field)
	GetDataEncoding() QualifiedName
	// GetContinuationPoint returns ContinuationPoint (property field)
	GetContinuationPoint() PascalByteString
	// IsHistoryReadValueId is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHistoryReadValueId()
	// CreateBuilder creates a HistoryReadValueIdBuilder
	CreateHistoryReadValueIdBuilder() HistoryReadValueIdBuilder
}

// _HistoryReadValueId is the data-structure of this message
type _HistoryReadValueId struct {
	ExtensionObjectDefinitionContract
	NodeId            NodeId
	IndexRange        PascalString
	DataEncoding      QualifiedName
	ContinuationPoint PascalByteString
}

var _ HistoryReadValueId = (*_HistoryReadValueId)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_HistoryReadValueId)(nil)

// NewHistoryReadValueId factory function for _HistoryReadValueId
func NewHistoryReadValueId(nodeId NodeId, indexRange PascalString, dataEncoding QualifiedName, continuationPoint PascalByteString) *_HistoryReadValueId {
	if nodeId == nil {
		panic("nodeId of type NodeId for HistoryReadValueId must not be nil")
	}
	if indexRange == nil {
		panic("indexRange of type PascalString for HistoryReadValueId must not be nil")
	}
	if dataEncoding == nil {
		panic("dataEncoding of type QualifiedName for HistoryReadValueId must not be nil")
	}
	if continuationPoint == nil {
		panic("continuationPoint of type PascalByteString for HistoryReadValueId must not be nil")
	}
	_result := &_HistoryReadValueId{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NodeId:                            nodeId,
		IndexRange:                        indexRange,
		DataEncoding:                      dataEncoding,
		ContinuationPoint:                 continuationPoint,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// HistoryReadValueIdBuilder is a builder for HistoryReadValueId
type HistoryReadValueIdBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(nodeId NodeId, indexRange PascalString, dataEncoding QualifiedName, continuationPoint PascalByteString) HistoryReadValueIdBuilder
	// WithNodeId adds NodeId (property field)
	WithNodeId(NodeId) HistoryReadValueIdBuilder
	// WithNodeIdBuilder adds NodeId (property field) which is build by the builder
	WithNodeIdBuilder(func(NodeIdBuilder) NodeIdBuilder) HistoryReadValueIdBuilder
	// WithIndexRange adds IndexRange (property field)
	WithIndexRange(PascalString) HistoryReadValueIdBuilder
	// WithIndexRangeBuilder adds IndexRange (property field) which is build by the builder
	WithIndexRangeBuilder(func(PascalStringBuilder) PascalStringBuilder) HistoryReadValueIdBuilder
	// WithDataEncoding adds DataEncoding (property field)
	WithDataEncoding(QualifiedName) HistoryReadValueIdBuilder
	// WithDataEncodingBuilder adds DataEncoding (property field) which is build by the builder
	WithDataEncodingBuilder(func(QualifiedNameBuilder) QualifiedNameBuilder) HistoryReadValueIdBuilder
	// WithContinuationPoint adds ContinuationPoint (property field)
	WithContinuationPoint(PascalByteString) HistoryReadValueIdBuilder
	// WithContinuationPointBuilder adds ContinuationPoint (property field) which is build by the builder
	WithContinuationPointBuilder(func(PascalByteStringBuilder) PascalByteStringBuilder) HistoryReadValueIdBuilder
	// Build builds the HistoryReadValueId or returns an error if something is wrong
	Build() (HistoryReadValueId, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HistoryReadValueId
}

// NewHistoryReadValueIdBuilder() creates a HistoryReadValueIdBuilder
func NewHistoryReadValueIdBuilder() HistoryReadValueIdBuilder {
	return &_HistoryReadValueIdBuilder{_HistoryReadValueId: new(_HistoryReadValueId)}
}

type _HistoryReadValueIdBuilder struct {
	*_HistoryReadValueId

	err *utils.MultiError
}

var _ (HistoryReadValueIdBuilder) = (*_HistoryReadValueIdBuilder)(nil)

func (m *_HistoryReadValueIdBuilder) WithMandatoryFields(nodeId NodeId, indexRange PascalString, dataEncoding QualifiedName, continuationPoint PascalByteString) HistoryReadValueIdBuilder {
	return m.WithNodeId(nodeId).WithIndexRange(indexRange).WithDataEncoding(dataEncoding).WithContinuationPoint(continuationPoint)
}

func (m *_HistoryReadValueIdBuilder) WithNodeId(nodeId NodeId) HistoryReadValueIdBuilder {
	m.NodeId = nodeId
	return m
}

func (m *_HistoryReadValueIdBuilder) WithNodeIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) HistoryReadValueIdBuilder {
	builder := builderSupplier(m.NodeId.CreateNodeIdBuilder())
	var err error
	m.NodeId, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return m
}

func (m *_HistoryReadValueIdBuilder) WithIndexRange(indexRange PascalString) HistoryReadValueIdBuilder {
	m.IndexRange = indexRange
	return m
}

func (m *_HistoryReadValueIdBuilder) WithIndexRangeBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) HistoryReadValueIdBuilder {
	builder := builderSupplier(m.IndexRange.CreatePascalStringBuilder())
	var err error
	m.IndexRange, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return m
}

func (m *_HistoryReadValueIdBuilder) WithDataEncoding(dataEncoding QualifiedName) HistoryReadValueIdBuilder {
	m.DataEncoding = dataEncoding
	return m
}

func (m *_HistoryReadValueIdBuilder) WithDataEncodingBuilder(builderSupplier func(QualifiedNameBuilder) QualifiedNameBuilder) HistoryReadValueIdBuilder {
	builder := builderSupplier(m.DataEncoding.CreateQualifiedNameBuilder())
	var err error
	m.DataEncoding, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "QualifiedNameBuilder failed"))
	}
	return m
}

func (m *_HistoryReadValueIdBuilder) WithContinuationPoint(continuationPoint PascalByteString) HistoryReadValueIdBuilder {
	m.ContinuationPoint = continuationPoint
	return m
}

func (m *_HistoryReadValueIdBuilder) WithContinuationPointBuilder(builderSupplier func(PascalByteStringBuilder) PascalByteStringBuilder) HistoryReadValueIdBuilder {
	builder := builderSupplier(m.ContinuationPoint.CreatePascalByteStringBuilder())
	var err error
	m.ContinuationPoint, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "PascalByteStringBuilder failed"))
	}
	return m
}

func (m *_HistoryReadValueIdBuilder) Build() (HistoryReadValueId, error) {
	if m.NodeId == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'nodeId' not set"))
	}
	if m.IndexRange == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'indexRange' not set"))
	}
	if m.DataEncoding == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'dataEncoding' not set"))
	}
	if m.ContinuationPoint == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'continuationPoint' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._HistoryReadValueId.deepCopy(), nil
}

func (m *_HistoryReadValueIdBuilder) MustBuild() HistoryReadValueId {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_HistoryReadValueIdBuilder) DeepCopy() any {
	return m.CreateHistoryReadValueIdBuilder()
}

// CreateHistoryReadValueIdBuilder creates a HistoryReadValueIdBuilder
func (m *_HistoryReadValueId) CreateHistoryReadValueIdBuilder() HistoryReadValueIdBuilder {
	if m == nil {
		return NewHistoryReadValueIdBuilder()
	}
	return &_HistoryReadValueIdBuilder{_HistoryReadValueId: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryReadValueId) GetIdentifier() string {
	return "637"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryReadValueId) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryReadValueId) GetNodeId() NodeId {
	return m.NodeId
}

func (m *_HistoryReadValueId) GetIndexRange() PascalString {
	return m.IndexRange
}

func (m *_HistoryReadValueId) GetDataEncoding() QualifiedName {
	return m.DataEncoding
}

func (m *_HistoryReadValueId) GetContinuationPoint() PascalByteString {
	return m.ContinuationPoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHistoryReadValueId(structType any) HistoryReadValueId {
	if casted, ok := structType.(HistoryReadValueId); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryReadValueId); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryReadValueId) GetTypeName() string {
	return "HistoryReadValueId"
}

func (m *_HistoryReadValueId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Simple field (indexRange)
	lengthInBits += m.IndexRange.GetLengthInBits(ctx)

	// Simple field (dataEncoding)
	lengthInBits += m.DataEncoding.GetLengthInBits(ctx)

	// Simple field (continuationPoint)
	lengthInBits += m.ContinuationPoint.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_HistoryReadValueId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_HistoryReadValueId) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__historyReadValueId HistoryReadValueId, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HistoryReadValueId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryReadValueId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nodeId, err := ReadSimpleField[NodeId](ctx, "nodeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeId' field"))
	}
	m.NodeId = nodeId

	indexRange, err := ReadSimpleField[PascalString](ctx, "indexRange", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'indexRange' field"))
	}
	m.IndexRange = indexRange

	dataEncoding, err := ReadSimpleField[QualifiedName](ctx, "dataEncoding", ReadComplex[QualifiedName](QualifiedNameParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataEncoding' field"))
	}
	m.DataEncoding = dataEncoding

	continuationPoint, err := ReadSimpleField[PascalByteString](ctx, "continuationPoint", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'continuationPoint' field"))
	}
	m.ContinuationPoint = continuationPoint

	if closeErr := readBuffer.CloseContext("HistoryReadValueId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryReadValueId")
	}

	return m, nil
}

func (m *_HistoryReadValueId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryReadValueId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryReadValueId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryReadValueId")
		}

		if err := WriteSimpleField[NodeId](ctx, "nodeId", m.GetNodeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeId' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "indexRange", m.GetIndexRange(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'indexRange' field")
		}

		if err := WriteSimpleField[QualifiedName](ctx, "dataEncoding", m.GetDataEncoding(), WriteComplex[QualifiedName](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataEncoding' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "continuationPoint", m.GetContinuationPoint(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'continuationPoint' field")
		}

		if popErr := writeBuffer.PopContext("HistoryReadValueId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryReadValueId")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryReadValueId) IsHistoryReadValueId() {}

func (m *_HistoryReadValueId) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HistoryReadValueId) deepCopy() *_HistoryReadValueId {
	if m == nil {
		return nil
	}
	_HistoryReadValueIdCopy := &_HistoryReadValueId{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.NodeId.DeepCopy().(NodeId),
		m.IndexRange.DeepCopy().(PascalString),
		m.DataEncoding.DeepCopy().(QualifiedName),
		m.ContinuationPoint.DeepCopy().(PascalByteString),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _HistoryReadValueIdCopy
}

func (m *_HistoryReadValueId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
