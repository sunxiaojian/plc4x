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

// AccessControlDataAccessPointForcedOpen is the corresponding interface of AccessControlDataAccessPointForcedOpen
type AccessControlDataAccessPointForcedOpen interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AccessControlData
	// IsAccessControlDataAccessPointForcedOpen is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAccessControlDataAccessPointForcedOpen()
	// CreateBuilder creates a AccessControlDataAccessPointForcedOpenBuilder
	CreateAccessControlDataAccessPointForcedOpenBuilder() AccessControlDataAccessPointForcedOpenBuilder
}

// _AccessControlDataAccessPointForcedOpen is the data-structure of this message
type _AccessControlDataAccessPointForcedOpen struct {
	AccessControlDataContract
}

var _ AccessControlDataAccessPointForcedOpen = (*_AccessControlDataAccessPointForcedOpen)(nil)
var _ AccessControlDataRequirements = (*_AccessControlDataAccessPointForcedOpen)(nil)

// NewAccessControlDataAccessPointForcedOpen factory function for _AccessControlDataAccessPointForcedOpen
func NewAccessControlDataAccessPointForcedOpen(commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) *_AccessControlDataAccessPointForcedOpen {
	_result := &_AccessControlDataAccessPointForcedOpen{
		AccessControlDataContract: NewAccessControlData(commandTypeContainer, networkId, accessPointId),
	}
	_result.AccessControlDataContract.(*_AccessControlData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AccessControlDataAccessPointForcedOpenBuilder is a builder for AccessControlDataAccessPointForcedOpen
type AccessControlDataAccessPointForcedOpenBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() AccessControlDataAccessPointForcedOpenBuilder
	// Build builds the AccessControlDataAccessPointForcedOpen or returns an error if something is wrong
	Build() (AccessControlDataAccessPointForcedOpen, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AccessControlDataAccessPointForcedOpen
}

// NewAccessControlDataAccessPointForcedOpenBuilder() creates a AccessControlDataAccessPointForcedOpenBuilder
func NewAccessControlDataAccessPointForcedOpenBuilder() AccessControlDataAccessPointForcedOpenBuilder {
	return &_AccessControlDataAccessPointForcedOpenBuilder{_AccessControlDataAccessPointForcedOpen: new(_AccessControlDataAccessPointForcedOpen)}
}

type _AccessControlDataAccessPointForcedOpenBuilder struct {
	*_AccessControlDataAccessPointForcedOpen

	err *utils.MultiError
}

var _ (AccessControlDataAccessPointForcedOpenBuilder) = (*_AccessControlDataAccessPointForcedOpenBuilder)(nil)

func (m *_AccessControlDataAccessPointForcedOpenBuilder) WithMandatoryFields() AccessControlDataAccessPointForcedOpenBuilder {
	return m
}

func (m *_AccessControlDataAccessPointForcedOpenBuilder) Build() (AccessControlDataAccessPointForcedOpen, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._AccessControlDataAccessPointForcedOpen.deepCopy(), nil
}

func (m *_AccessControlDataAccessPointForcedOpenBuilder) MustBuild() AccessControlDataAccessPointForcedOpen {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_AccessControlDataAccessPointForcedOpenBuilder) DeepCopy() any {
	return m.CreateAccessControlDataAccessPointForcedOpenBuilder()
}

// CreateAccessControlDataAccessPointForcedOpenBuilder creates a AccessControlDataAccessPointForcedOpenBuilder
func (m *_AccessControlDataAccessPointForcedOpen) CreateAccessControlDataAccessPointForcedOpenBuilder() AccessControlDataAccessPointForcedOpenBuilder {
	if m == nil {
		return NewAccessControlDataAccessPointForcedOpenBuilder()
	}
	return &_AccessControlDataAccessPointForcedOpenBuilder{_AccessControlDataAccessPointForcedOpen: m.deepCopy()}
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

func (m *_AccessControlDataAccessPointForcedOpen) GetParent() AccessControlDataContract {
	return m.AccessControlDataContract
}

// Deprecated: use the interface for direct cast
func CastAccessControlDataAccessPointForcedOpen(structType any) AccessControlDataAccessPointForcedOpen {
	if casted, ok := structType.(AccessControlDataAccessPointForcedOpen); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlDataAccessPointForcedOpen); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlDataAccessPointForcedOpen) GetTypeName() string {
	return "AccessControlDataAccessPointForcedOpen"
}

func (m *_AccessControlDataAccessPointForcedOpen) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AccessControlDataContract.(*_AccessControlData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_AccessControlDataAccessPointForcedOpen) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AccessControlDataAccessPointForcedOpen) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AccessControlData) (__accessControlDataAccessPointForcedOpen AccessControlDataAccessPointForcedOpen, err error) {
	m.AccessControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AccessControlDataAccessPointForcedOpen"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlDataAccessPointForcedOpen")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AccessControlDataAccessPointForcedOpen"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlDataAccessPointForcedOpen")
	}

	return m, nil
}

func (m *_AccessControlDataAccessPointForcedOpen) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AccessControlDataAccessPointForcedOpen) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AccessControlDataAccessPointForcedOpen"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AccessControlDataAccessPointForcedOpen")
		}

		if popErr := writeBuffer.PopContext("AccessControlDataAccessPointForcedOpen"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AccessControlDataAccessPointForcedOpen")
		}
		return nil
	}
	return m.AccessControlDataContract.(*_AccessControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AccessControlDataAccessPointForcedOpen) IsAccessControlDataAccessPointForcedOpen() {}

func (m *_AccessControlDataAccessPointForcedOpen) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AccessControlDataAccessPointForcedOpen) deepCopy() *_AccessControlDataAccessPointForcedOpen {
	if m == nil {
		return nil
	}
	_AccessControlDataAccessPointForcedOpenCopy := &_AccessControlDataAccessPointForcedOpen{
		m.AccessControlDataContract.(*_AccessControlData).deepCopy(),
	}
	m.AccessControlDataContract.(*_AccessControlData)._SubType = m
	return _AccessControlDataAccessPointForcedOpenCopy
}

func (m *_AccessControlDataAccessPointForcedOpen) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
