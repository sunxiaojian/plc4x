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

// BrowsePathTarget is the corresponding interface of BrowsePathTarget
type BrowsePathTarget interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetTargetId returns TargetId (property field)
	GetTargetId() ExpandedNodeId
	// GetRemainingPathIndex returns RemainingPathIndex (property field)
	GetRemainingPathIndex() uint32
	// IsBrowsePathTarget is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBrowsePathTarget()
	// CreateBuilder creates a BrowsePathTargetBuilder
	CreateBrowsePathTargetBuilder() BrowsePathTargetBuilder
}

// _BrowsePathTarget is the data-structure of this message
type _BrowsePathTarget struct {
	ExtensionObjectDefinitionContract
	TargetId           ExpandedNodeId
	RemainingPathIndex uint32
}

var _ BrowsePathTarget = (*_BrowsePathTarget)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_BrowsePathTarget)(nil)

// NewBrowsePathTarget factory function for _BrowsePathTarget
func NewBrowsePathTarget(targetId ExpandedNodeId, remainingPathIndex uint32) *_BrowsePathTarget {
	if targetId == nil {
		panic("targetId of type ExpandedNodeId for BrowsePathTarget must not be nil")
	}
	_result := &_BrowsePathTarget{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		TargetId:                          targetId,
		RemainingPathIndex:                remainingPathIndex,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BrowsePathTargetBuilder is a builder for BrowsePathTarget
type BrowsePathTargetBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(targetId ExpandedNodeId, remainingPathIndex uint32) BrowsePathTargetBuilder
	// WithTargetId adds TargetId (property field)
	WithTargetId(ExpandedNodeId) BrowsePathTargetBuilder
	// WithTargetIdBuilder adds TargetId (property field) which is build by the builder
	WithTargetIdBuilder(func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) BrowsePathTargetBuilder
	// WithRemainingPathIndex adds RemainingPathIndex (property field)
	WithRemainingPathIndex(uint32) BrowsePathTargetBuilder
	// Build builds the BrowsePathTarget or returns an error if something is wrong
	Build() (BrowsePathTarget, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BrowsePathTarget
}

// NewBrowsePathTargetBuilder() creates a BrowsePathTargetBuilder
func NewBrowsePathTargetBuilder() BrowsePathTargetBuilder {
	return &_BrowsePathTargetBuilder{_BrowsePathTarget: new(_BrowsePathTarget)}
}

type _BrowsePathTargetBuilder struct {
	*_BrowsePathTarget

	err *utils.MultiError
}

var _ (BrowsePathTargetBuilder) = (*_BrowsePathTargetBuilder)(nil)

func (m *_BrowsePathTargetBuilder) WithMandatoryFields(targetId ExpandedNodeId, remainingPathIndex uint32) BrowsePathTargetBuilder {
	return m.WithTargetId(targetId).WithRemainingPathIndex(remainingPathIndex)
}

func (m *_BrowsePathTargetBuilder) WithTargetId(targetId ExpandedNodeId) BrowsePathTargetBuilder {
	m.TargetId = targetId
	return m
}

func (m *_BrowsePathTargetBuilder) WithTargetIdBuilder(builderSupplier func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) BrowsePathTargetBuilder {
	builder := builderSupplier(m.TargetId.CreateExpandedNodeIdBuilder())
	var err error
	m.TargetId, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "ExpandedNodeIdBuilder failed"))
	}
	return m
}

func (m *_BrowsePathTargetBuilder) WithRemainingPathIndex(remainingPathIndex uint32) BrowsePathTargetBuilder {
	m.RemainingPathIndex = remainingPathIndex
	return m
}

func (m *_BrowsePathTargetBuilder) Build() (BrowsePathTarget, error) {
	if m.TargetId == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'targetId' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BrowsePathTarget.deepCopy(), nil
}

func (m *_BrowsePathTargetBuilder) MustBuild() BrowsePathTarget {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BrowsePathTargetBuilder) DeepCopy() any {
	return m.CreateBrowsePathTargetBuilder()
}

// CreateBrowsePathTargetBuilder creates a BrowsePathTargetBuilder
func (m *_BrowsePathTarget) CreateBrowsePathTargetBuilder() BrowsePathTargetBuilder {
	if m == nil {
		return NewBrowsePathTargetBuilder()
	}
	return &_BrowsePathTargetBuilder{_BrowsePathTarget: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrowsePathTarget) GetIdentifier() string {
	return "548"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrowsePathTarget) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrowsePathTarget) GetTargetId() ExpandedNodeId {
	return m.TargetId
}

func (m *_BrowsePathTarget) GetRemainingPathIndex() uint32 {
	return m.RemainingPathIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBrowsePathTarget(structType any) BrowsePathTarget {
	if casted, ok := structType.(BrowsePathTarget); ok {
		return casted
	}
	if casted, ok := structType.(*BrowsePathTarget); ok {
		return *casted
	}
	return nil
}

func (m *_BrowsePathTarget) GetTypeName() string {
	return "BrowsePathTarget"
}

func (m *_BrowsePathTarget) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (targetId)
	lengthInBits += m.TargetId.GetLengthInBits(ctx)

	// Simple field (remainingPathIndex)
	lengthInBits += 32

	return lengthInBits
}

func (m *_BrowsePathTarget) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BrowsePathTarget) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__browsePathTarget BrowsePathTarget, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BrowsePathTarget"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrowsePathTarget")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	targetId, err := ReadSimpleField[ExpandedNodeId](ctx, "targetId", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetId' field"))
	}
	m.TargetId = targetId

	remainingPathIndex, err := ReadSimpleField(ctx, "remainingPathIndex", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'remainingPathIndex' field"))
	}
	m.RemainingPathIndex = remainingPathIndex

	if closeErr := readBuffer.CloseContext("BrowsePathTarget"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrowsePathTarget")
	}

	return m, nil
}

func (m *_BrowsePathTarget) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrowsePathTarget) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrowsePathTarget"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrowsePathTarget")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "targetId", m.GetTargetId(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'targetId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "remainingPathIndex", m.GetRemainingPathIndex(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'remainingPathIndex' field")
		}

		if popErr := writeBuffer.PopContext("BrowsePathTarget"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrowsePathTarget")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrowsePathTarget) IsBrowsePathTarget() {}

func (m *_BrowsePathTarget) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BrowsePathTarget) deepCopy() *_BrowsePathTarget {
	if m == nil {
		return nil
	}
	_BrowsePathTargetCopy := &_BrowsePathTarget{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.TargetId.DeepCopy().(ExpandedNodeId),
		m.RemainingPathIndex,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _BrowsePathTargetCopy
}

func (m *_BrowsePathTarget) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
