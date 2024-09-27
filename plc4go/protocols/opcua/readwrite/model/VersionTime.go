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

// VersionTime is the corresponding interface of VersionTime
type VersionTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsVersionTime is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVersionTime()
	// CreateBuilder creates a VersionTimeBuilder
	CreateVersionTimeBuilder() VersionTimeBuilder
}

// _VersionTime is the data-structure of this message
type _VersionTime struct {
}

var _ VersionTime = (*_VersionTime)(nil)

// NewVersionTime factory function for _VersionTime
func NewVersionTime() *_VersionTime {
	return &_VersionTime{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// VersionTimeBuilder is a builder for VersionTime
type VersionTimeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() VersionTimeBuilder
	// Build builds the VersionTime or returns an error if something is wrong
	Build() (VersionTime, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() VersionTime
}

// NewVersionTimeBuilder() creates a VersionTimeBuilder
func NewVersionTimeBuilder() VersionTimeBuilder {
	return &_VersionTimeBuilder{_VersionTime: new(_VersionTime)}
}

type _VersionTimeBuilder struct {
	*_VersionTime

	err *utils.MultiError
}

var _ (VersionTimeBuilder) = (*_VersionTimeBuilder)(nil)

func (m *_VersionTimeBuilder) WithMandatoryFields() VersionTimeBuilder {
	return m
}

func (m *_VersionTimeBuilder) Build() (VersionTime, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._VersionTime.deepCopy(), nil
}

func (m *_VersionTimeBuilder) MustBuild() VersionTime {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_VersionTimeBuilder) DeepCopy() any {
	return m.CreateVersionTimeBuilder()
}

// CreateVersionTimeBuilder creates a VersionTimeBuilder
func (m *_VersionTime) CreateVersionTimeBuilder() VersionTimeBuilder {
	if m == nil {
		return NewVersionTimeBuilder()
	}
	return &_VersionTimeBuilder{_VersionTime: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastVersionTime(structType any) VersionTime {
	if casted, ok := structType.(VersionTime); ok {
		return casted
	}
	if casted, ok := structType.(*VersionTime); ok {
		return *casted
	}
	return nil
}

func (m *_VersionTime) GetTypeName() string {
	return "VersionTime"
}

func (m *_VersionTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_VersionTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func VersionTimeParse(ctx context.Context, theBytes []byte) (VersionTime, error) {
	return VersionTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func VersionTimeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (VersionTime, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (VersionTime, error) {
		return VersionTimeParseWithBuffer(ctx, readBuffer)
	}
}

func VersionTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (VersionTime, error) {
	v, err := (&_VersionTime{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_VersionTime) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__versionTime VersionTime, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VersionTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VersionTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("VersionTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VersionTime")
	}

	return m, nil
}

func (m *_VersionTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VersionTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("VersionTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for VersionTime")
	}

	if popErr := writeBuffer.PopContext("VersionTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for VersionTime")
	}
	return nil
}

func (m *_VersionTime) IsVersionTime() {}

func (m *_VersionTime) DeepCopy() any {
	return m.deepCopy()
}

func (m *_VersionTime) deepCopy() *_VersionTime {
	if m == nil {
		return nil
	}
	_VersionTimeCopy := &_VersionTime{}
	return _VersionTimeCopy
}

func (m *_VersionTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
