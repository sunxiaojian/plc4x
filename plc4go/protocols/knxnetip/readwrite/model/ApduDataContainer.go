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

// ApduDataContainer is the corresponding interface of ApduDataContainer
type ApduDataContainer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Apdu
	// GetDataApdu returns DataApdu (property field)
	GetDataApdu() ApduData
	// IsApduDataContainer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataContainer()
	// CreateBuilder creates a ApduDataContainerBuilder
	CreateApduDataContainerBuilder() ApduDataContainerBuilder
}

// _ApduDataContainer is the data-structure of this message
type _ApduDataContainer struct {
	ApduContract
	DataApdu ApduData
}

var _ ApduDataContainer = (*_ApduDataContainer)(nil)
var _ ApduRequirements = (*_ApduDataContainer)(nil)

// NewApduDataContainer factory function for _ApduDataContainer
func NewApduDataContainer(numbered bool, counter uint8, dataApdu ApduData, dataLength uint8) *_ApduDataContainer {
	if dataApdu == nil {
		panic("dataApdu of type ApduData for ApduDataContainer must not be nil")
	}
	_result := &_ApduDataContainer{
		ApduContract: NewApdu(numbered, counter, dataLength),
		DataApdu:     dataApdu,
	}
	_result.ApduContract.(*_Apdu)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataContainerBuilder is a builder for ApduDataContainer
type ApduDataContainerBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dataApdu ApduData) ApduDataContainerBuilder
	// WithDataApdu adds DataApdu (property field)
	WithDataApdu(ApduData) ApduDataContainerBuilder
	// Build builds the ApduDataContainer or returns an error if something is wrong
	Build() (ApduDataContainer, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataContainer
}

// NewApduDataContainerBuilder() creates a ApduDataContainerBuilder
func NewApduDataContainerBuilder() ApduDataContainerBuilder {
	return &_ApduDataContainerBuilder{_ApduDataContainer: new(_ApduDataContainer)}
}

type _ApduDataContainerBuilder struct {
	*_ApduDataContainer

	err *utils.MultiError
}

var _ (ApduDataContainerBuilder) = (*_ApduDataContainerBuilder)(nil)

func (m *_ApduDataContainerBuilder) WithMandatoryFields(dataApdu ApduData) ApduDataContainerBuilder {
	return m.WithDataApdu(dataApdu)
}

func (m *_ApduDataContainerBuilder) WithDataApdu(dataApdu ApduData) ApduDataContainerBuilder {
	m.DataApdu = dataApdu
	return m
}

func (m *_ApduDataContainerBuilder) Build() (ApduDataContainer, error) {
	if m.DataApdu == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'dataApdu' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._ApduDataContainer.deepCopy(), nil
}

func (m *_ApduDataContainerBuilder) MustBuild() ApduDataContainer {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_ApduDataContainerBuilder) DeepCopy() any {
	return m.CreateApduDataContainerBuilder()
}

// CreateApduDataContainerBuilder creates a ApduDataContainerBuilder
func (m *_ApduDataContainer) CreateApduDataContainerBuilder() ApduDataContainerBuilder {
	if m == nil {
		return NewApduDataContainerBuilder()
	}
	return &_ApduDataContainerBuilder{_ApduDataContainer: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataContainer) GetControl() uint8 {
	return uint8(0)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataContainer) GetParent() ApduContract {
	return m.ApduContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataContainer) GetDataApdu() ApduData {
	return m.DataApdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastApduDataContainer(structType any) ApduDataContainer {
	if casted, ok := structType.(ApduDataContainer); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataContainer); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataContainer) GetTypeName() string {
	return "ApduDataContainer"
}

func (m *_ApduDataContainer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduContract.(*_Apdu).getLengthInBits(ctx))

	// Simple field (dataApdu)
	lengthInBits += m.DataApdu.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ApduDataContainer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataContainer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Apdu, dataLength uint8) (__apduDataContainer ApduDataContainer, err error) {
	m.ApduContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataContainer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataContainer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataApdu, err := ReadSimpleField[ApduData](ctx, "dataApdu", ReadComplex[ApduData](ApduDataParseWithBufferProducer[ApduData]((uint8)(dataLength)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataApdu' field"))
	}
	m.DataApdu = dataApdu

	if closeErr := readBuffer.CloseContext("ApduDataContainer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataContainer")
	}

	return m, nil
}

func (m *_ApduDataContainer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataContainer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataContainer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataContainer")
		}

		if err := WriteSimpleField[ApduData](ctx, "dataApdu", m.GetDataApdu(), WriteComplex[ApduData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataApdu' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataContainer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataContainer")
		}
		return nil
	}
	return m.ApduContract.(*_Apdu).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataContainer) IsApduDataContainer() {}

func (m *_ApduDataContainer) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataContainer) deepCopy() *_ApduDataContainer {
	if m == nil {
		return nil
	}
	_ApduDataContainerCopy := &_ApduDataContainer{
		m.ApduContract.(*_Apdu).deepCopy(),
		m.DataApdu.DeepCopy().(ApduData),
	}
	m.ApduContract.(*_Apdu)._SubType = m
	return _ApduDataContainerCopy
}

func (m *_ApduDataContainer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
