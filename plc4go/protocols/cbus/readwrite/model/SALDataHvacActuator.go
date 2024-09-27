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

// SALDataHvacActuator is the corresponding interface of SALDataHvacActuator
type SALDataHvacActuator interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SALData
	// GetHvacActuatorData returns HvacActuatorData (property field)
	GetHvacActuatorData() LightingData
	// IsSALDataHvacActuator is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALDataHvacActuator()
	// CreateBuilder creates a SALDataHvacActuatorBuilder
	CreateSALDataHvacActuatorBuilder() SALDataHvacActuatorBuilder
}

// _SALDataHvacActuator is the data-structure of this message
type _SALDataHvacActuator struct {
	SALDataContract
	HvacActuatorData LightingData
}

var _ SALDataHvacActuator = (*_SALDataHvacActuator)(nil)
var _ SALDataRequirements = (*_SALDataHvacActuator)(nil)

// NewSALDataHvacActuator factory function for _SALDataHvacActuator
func NewSALDataHvacActuator(salData SALData, hvacActuatorData LightingData) *_SALDataHvacActuator {
	if hvacActuatorData == nil {
		panic("hvacActuatorData of type LightingData for SALDataHvacActuator must not be nil")
	}
	_result := &_SALDataHvacActuator{
		SALDataContract:  NewSALData(salData),
		HvacActuatorData: hvacActuatorData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SALDataHvacActuatorBuilder is a builder for SALDataHvacActuator
type SALDataHvacActuatorBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(hvacActuatorData LightingData) SALDataHvacActuatorBuilder
	// WithHvacActuatorData adds HvacActuatorData (property field)
	WithHvacActuatorData(LightingData) SALDataHvacActuatorBuilder
	// Build builds the SALDataHvacActuator or returns an error if something is wrong
	Build() (SALDataHvacActuator, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SALDataHvacActuator
}

// NewSALDataHvacActuatorBuilder() creates a SALDataHvacActuatorBuilder
func NewSALDataHvacActuatorBuilder() SALDataHvacActuatorBuilder {
	return &_SALDataHvacActuatorBuilder{_SALDataHvacActuator: new(_SALDataHvacActuator)}
}

type _SALDataHvacActuatorBuilder struct {
	*_SALDataHvacActuator

	err *utils.MultiError
}

var _ (SALDataHvacActuatorBuilder) = (*_SALDataHvacActuatorBuilder)(nil)

func (m *_SALDataHvacActuatorBuilder) WithMandatoryFields(hvacActuatorData LightingData) SALDataHvacActuatorBuilder {
	return m.WithHvacActuatorData(hvacActuatorData)
}

func (m *_SALDataHvacActuatorBuilder) WithHvacActuatorData(hvacActuatorData LightingData) SALDataHvacActuatorBuilder {
	m.HvacActuatorData = hvacActuatorData
	return m
}

func (m *_SALDataHvacActuatorBuilder) Build() (SALDataHvacActuator, error) {
	if m.HvacActuatorData == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'hvacActuatorData' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SALDataHvacActuator.deepCopy(), nil
}

func (m *_SALDataHvacActuatorBuilder) MustBuild() SALDataHvacActuator {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SALDataHvacActuatorBuilder) DeepCopy() any {
	return m.CreateSALDataHvacActuatorBuilder()
}

// CreateSALDataHvacActuatorBuilder creates a SALDataHvacActuatorBuilder
func (m *_SALDataHvacActuator) CreateSALDataHvacActuatorBuilder() SALDataHvacActuatorBuilder {
	if m == nil {
		return NewSALDataHvacActuatorBuilder()
	}
	return &_SALDataHvacActuatorBuilder{_SALDataHvacActuator: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataHvacActuator) GetApplicationId() ApplicationId {
	return ApplicationId_HVAC_ACTUATOR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataHvacActuator) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataHvacActuator) GetHvacActuatorData() LightingData {
	return m.HvacActuatorData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSALDataHvacActuator(structType any) SALDataHvacActuator {
	if casted, ok := structType.(SALDataHvacActuator); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataHvacActuator); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataHvacActuator) GetTypeName() string {
	return "SALDataHvacActuator"
}

func (m *_SALDataHvacActuator) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (hvacActuatorData)
	lengthInBits += m.HvacActuatorData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataHvacActuator) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataHvacActuator) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataHvacActuator SALDataHvacActuator, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataHvacActuator"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataHvacActuator")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	hvacActuatorData, err := ReadSimpleField[LightingData](ctx, "hvacActuatorData", ReadComplex[LightingData](LightingDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacActuatorData' field"))
	}
	m.HvacActuatorData = hvacActuatorData

	if closeErr := readBuffer.CloseContext("SALDataHvacActuator"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataHvacActuator")
	}

	return m, nil
}

func (m *_SALDataHvacActuator) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataHvacActuator) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataHvacActuator"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataHvacActuator")
		}

		if err := WriteSimpleField[LightingData](ctx, "hvacActuatorData", m.GetHvacActuatorData(), WriteComplex[LightingData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacActuatorData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataHvacActuator"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataHvacActuator")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataHvacActuator) IsSALDataHvacActuator() {}

func (m *_SALDataHvacActuator) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SALDataHvacActuator) deepCopy() *_SALDataHvacActuator {
	if m == nil {
		return nil
	}
	_SALDataHvacActuatorCopy := &_SALDataHvacActuator{
		m.SALDataContract.(*_SALData).deepCopy(),
		m.HvacActuatorData.DeepCopy().(LightingData),
	}
	m.SALDataContract.(*_SALData)._SubType = m
	return _SALDataHvacActuatorCopy
}

func (m *_SALDataHvacActuator) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
