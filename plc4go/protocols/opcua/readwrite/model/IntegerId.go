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

// IntegerId is the corresponding interface of IntegerId
type IntegerId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsIntegerId is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIntegerId()
	// CreateBuilder creates a IntegerIdBuilder
	CreateIntegerIdBuilder() IntegerIdBuilder
}

// _IntegerId is the data-structure of this message
type _IntegerId struct {
}

var _ IntegerId = (*_IntegerId)(nil)

// NewIntegerId factory function for _IntegerId
func NewIntegerId() *_IntegerId {
	return &_IntegerId{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// IntegerIdBuilder is a builder for IntegerId
type IntegerIdBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() IntegerIdBuilder
	// Build builds the IntegerId or returns an error if something is wrong
	Build() (IntegerId, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() IntegerId
}

// NewIntegerIdBuilder() creates a IntegerIdBuilder
func NewIntegerIdBuilder() IntegerIdBuilder {
	return &_IntegerIdBuilder{_IntegerId: new(_IntegerId)}
}

type _IntegerIdBuilder struct {
	*_IntegerId

	err *utils.MultiError
}

var _ (IntegerIdBuilder) = (*_IntegerIdBuilder)(nil)

func (m *_IntegerIdBuilder) WithMandatoryFields() IntegerIdBuilder {
	return m
}

func (m *_IntegerIdBuilder) Build() (IntegerId, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._IntegerId.deepCopy(), nil
}

func (m *_IntegerIdBuilder) MustBuild() IntegerId {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_IntegerIdBuilder) DeepCopy() any {
	return m.CreateIntegerIdBuilder()
}

// CreateIntegerIdBuilder creates a IntegerIdBuilder
func (m *_IntegerId) CreateIntegerIdBuilder() IntegerIdBuilder {
	if m == nil {
		return NewIntegerIdBuilder()
	}
	return &_IntegerIdBuilder{_IntegerId: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastIntegerId(structType any) IntegerId {
	if casted, ok := structType.(IntegerId); ok {
		return casted
	}
	if casted, ok := structType.(*IntegerId); ok {
		return *casted
	}
	return nil
}

func (m *_IntegerId) GetTypeName() string {
	return "IntegerId"
}

func (m *_IntegerId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_IntegerId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IntegerIdParse(ctx context.Context, theBytes []byte) (IntegerId, error) {
	return IntegerIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func IntegerIdParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (IntegerId, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (IntegerId, error) {
		return IntegerIdParseWithBuffer(ctx, readBuffer)
	}
}

func IntegerIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (IntegerId, error) {
	v, err := (&_IntegerId{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_IntegerId) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__integerId IntegerId, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IntegerId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IntegerId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("IntegerId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IntegerId")
	}

	return m, nil
}

func (m *_IntegerId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IntegerId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("IntegerId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for IntegerId")
	}

	if popErr := writeBuffer.PopContext("IntegerId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for IntegerId")
	}
	return nil
}

func (m *_IntegerId) IsIntegerId() {}

func (m *_IntegerId) DeepCopy() any {
	return m.deepCopy()
}

func (m *_IntegerId) deepCopy() *_IntegerId {
	if m == nil {
		return nil
	}
	_IntegerIdCopy := &_IntegerId{}
	return _IntegerIdCopy
}

func (m *_IntegerId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
