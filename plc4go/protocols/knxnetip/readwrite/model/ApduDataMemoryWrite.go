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

// ApduDataMemoryWrite is the corresponding interface of ApduDataMemoryWrite
type ApduDataMemoryWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduData
	// IsApduDataMemoryWrite is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataMemoryWrite()
	// CreateBuilder creates a ApduDataMemoryWriteBuilder
	CreateApduDataMemoryWriteBuilder() ApduDataMemoryWriteBuilder
}

// _ApduDataMemoryWrite is the data-structure of this message
type _ApduDataMemoryWrite struct {
	ApduDataContract
}

var _ ApduDataMemoryWrite = (*_ApduDataMemoryWrite)(nil)
var _ ApduDataRequirements = (*_ApduDataMemoryWrite)(nil)

// NewApduDataMemoryWrite factory function for _ApduDataMemoryWrite
func NewApduDataMemoryWrite(dataLength uint8) *_ApduDataMemoryWrite {
	_result := &_ApduDataMemoryWrite{
		ApduDataContract: NewApduData(dataLength),
	}
	_result.ApduDataContract.(*_ApduData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataMemoryWriteBuilder is a builder for ApduDataMemoryWrite
type ApduDataMemoryWriteBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataMemoryWriteBuilder
	// Build builds the ApduDataMemoryWrite or returns an error if something is wrong
	Build() (ApduDataMemoryWrite, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataMemoryWrite
}

// NewApduDataMemoryWriteBuilder() creates a ApduDataMemoryWriteBuilder
func NewApduDataMemoryWriteBuilder() ApduDataMemoryWriteBuilder {
	return &_ApduDataMemoryWriteBuilder{_ApduDataMemoryWrite: new(_ApduDataMemoryWrite)}
}

type _ApduDataMemoryWriteBuilder struct {
	*_ApduDataMemoryWrite

	err *utils.MultiError
}

var _ (ApduDataMemoryWriteBuilder) = (*_ApduDataMemoryWriteBuilder)(nil)

func (m *_ApduDataMemoryWriteBuilder) WithMandatoryFields() ApduDataMemoryWriteBuilder {
	return m
}

func (m *_ApduDataMemoryWriteBuilder) Build() (ApduDataMemoryWrite, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ApduDataMemoryWrite.deepCopy(), nil
}

func (m *_ApduDataMemoryWriteBuilder) MustBuild() ApduDataMemoryWrite {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ApduDataMemoryWriteBuilder) DeepCopy() any {
	return m.CreateApduDataMemoryWriteBuilder()
}

// CreateApduDataMemoryWriteBuilder creates a ApduDataMemoryWriteBuilder
func (m *_ApduDataMemoryWrite) CreateApduDataMemoryWriteBuilder() ApduDataMemoryWriteBuilder {
	if m == nil {
		return NewApduDataMemoryWriteBuilder()
	}
	return &_ApduDataMemoryWriteBuilder{_ApduDataMemoryWrite: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataMemoryWrite) GetApciType() uint8 {
	return 0xA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataMemoryWrite) GetParent() ApduDataContract {
	return m.ApduDataContract
}

// Deprecated: use the interface for direct cast
func CastApduDataMemoryWrite(structType any) ApduDataMemoryWrite {
	if casted, ok := structType.(ApduDataMemoryWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataMemoryWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataMemoryWrite) GetTypeName() string {
	return "ApduDataMemoryWrite"
}

func (m *_ApduDataMemoryWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataContract.(*_ApduData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataMemoryWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataMemoryWrite) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduData, dataLength uint8) (__apduDataMemoryWrite ApduDataMemoryWrite, err error) {
	m.ApduDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataMemoryWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataMemoryWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataMemoryWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataMemoryWrite")
	}

	return m, nil
}

func (m *_ApduDataMemoryWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataMemoryWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataMemoryWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataMemoryWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataMemoryWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataMemoryWrite")
		}
		return nil
	}
	return m.ApduDataContract.(*_ApduData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataMemoryWrite) IsApduDataMemoryWrite() {}

func (m *_ApduDataMemoryWrite) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataMemoryWrite) deepCopy() *_ApduDataMemoryWrite {
	if m == nil {
		return nil
	}
	_ApduDataMemoryWriteCopy := &_ApduDataMemoryWrite{
		m.ApduDataContract.(*_ApduData).deepCopy(),
	}
	m.ApduDataContract.(*_ApduData)._SubType = m
	return _ApduDataMemoryWriteCopy
}

func (m *_ApduDataMemoryWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
