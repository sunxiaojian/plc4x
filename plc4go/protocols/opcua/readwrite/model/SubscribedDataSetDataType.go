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

// SubscribedDataSetDataType is the corresponding interface of SubscribedDataSetDataType
type SubscribedDataSetDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// IsSubscribedDataSetDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSubscribedDataSetDataType()
	// CreateBuilder creates a SubscribedDataSetDataTypeBuilder
	CreateSubscribedDataSetDataTypeBuilder() SubscribedDataSetDataTypeBuilder
}

// _SubscribedDataSetDataType is the data-structure of this message
type _SubscribedDataSetDataType struct {
	ExtensionObjectDefinitionContract
}

var _ SubscribedDataSetDataType = (*_SubscribedDataSetDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_SubscribedDataSetDataType)(nil)

// NewSubscribedDataSetDataType factory function for _SubscribedDataSetDataType
func NewSubscribedDataSetDataType() *_SubscribedDataSetDataType {
	_result := &_SubscribedDataSetDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SubscribedDataSetDataTypeBuilder is a builder for SubscribedDataSetDataType
type SubscribedDataSetDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SubscribedDataSetDataTypeBuilder
	// Build builds the SubscribedDataSetDataType or returns an error if something is wrong
	Build() (SubscribedDataSetDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SubscribedDataSetDataType
}

// NewSubscribedDataSetDataTypeBuilder() creates a SubscribedDataSetDataTypeBuilder
func NewSubscribedDataSetDataTypeBuilder() SubscribedDataSetDataTypeBuilder {
	return &_SubscribedDataSetDataTypeBuilder{_SubscribedDataSetDataType: new(_SubscribedDataSetDataType)}
}

type _SubscribedDataSetDataTypeBuilder struct {
	*_SubscribedDataSetDataType

	err *utils.MultiError
}

var _ (SubscribedDataSetDataTypeBuilder) = (*_SubscribedDataSetDataTypeBuilder)(nil)

func (m *_SubscribedDataSetDataTypeBuilder) WithMandatoryFields() SubscribedDataSetDataTypeBuilder {
	return m
}

func (m *_SubscribedDataSetDataTypeBuilder) Build() (SubscribedDataSetDataType, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SubscribedDataSetDataType.deepCopy(), nil
}

func (m *_SubscribedDataSetDataTypeBuilder) MustBuild() SubscribedDataSetDataType {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SubscribedDataSetDataTypeBuilder) DeepCopy() any {
	return m.CreateSubscribedDataSetDataTypeBuilder()
}

// CreateSubscribedDataSetDataTypeBuilder creates a SubscribedDataSetDataTypeBuilder
func (m *_SubscribedDataSetDataType) CreateSubscribedDataSetDataTypeBuilder() SubscribedDataSetDataTypeBuilder {
	if m == nil {
		return NewSubscribedDataSetDataTypeBuilder()
	}
	return &_SubscribedDataSetDataTypeBuilder{_SubscribedDataSetDataType: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SubscribedDataSetDataType) GetIdentifier() string {
	return "15632"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SubscribedDataSetDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastSubscribedDataSetDataType(structType any) SubscribedDataSetDataType {
	if casted, ok := structType.(SubscribedDataSetDataType); ok {
		return casted
	}
	if casted, ok := structType.(*SubscribedDataSetDataType); ok {
		return *casted
	}
	return nil
}

func (m *_SubscribedDataSetDataType) GetTypeName() string {
	return "SubscribedDataSetDataType"
}

func (m *_SubscribedDataSetDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SubscribedDataSetDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SubscribedDataSetDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__subscribedDataSetDataType SubscribedDataSetDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SubscribedDataSetDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SubscribedDataSetDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SubscribedDataSetDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SubscribedDataSetDataType")
	}

	return m, nil
}

func (m *_SubscribedDataSetDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SubscribedDataSetDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SubscribedDataSetDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SubscribedDataSetDataType")
		}

		if popErr := writeBuffer.PopContext("SubscribedDataSetDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SubscribedDataSetDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SubscribedDataSetDataType) IsSubscribedDataSetDataType() {}

func (m *_SubscribedDataSetDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SubscribedDataSetDataType) deepCopy() *_SubscribedDataSetDataType {
	if m == nil {
		return nil
	}
	_SubscribedDataSetDataTypeCopy := &_SubscribedDataSetDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _SubscribedDataSetDataTypeCopy
}

func (m *_SubscribedDataSetDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
