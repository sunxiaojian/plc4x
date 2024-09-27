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

// BACnetPropertyStatesEscalatorFault is the corresponding interface of BACnetPropertyStatesEscalatorFault
type BACnetPropertyStatesEscalatorFault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetEscalatorFault returns EscalatorFault (property field)
	GetEscalatorFault() BACnetEscalatorFaultTagged
	// IsBACnetPropertyStatesEscalatorFault is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesEscalatorFault()
	// CreateBuilder creates a BACnetPropertyStatesEscalatorFaultBuilder
	CreateBACnetPropertyStatesEscalatorFaultBuilder() BACnetPropertyStatesEscalatorFaultBuilder
}

// _BACnetPropertyStatesEscalatorFault is the data-structure of this message
type _BACnetPropertyStatesEscalatorFault struct {
	BACnetPropertyStatesContract
	EscalatorFault BACnetEscalatorFaultTagged
}

var _ BACnetPropertyStatesEscalatorFault = (*_BACnetPropertyStatesEscalatorFault)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesEscalatorFault)(nil)

// NewBACnetPropertyStatesEscalatorFault factory function for _BACnetPropertyStatesEscalatorFault
func NewBACnetPropertyStatesEscalatorFault(peekedTagHeader BACnetTagHeader, escalatorFault BACnetEscalatorFaultTagged) *_BACnetPropertyStatesEscalatorFault {
	if escalatorFault == nil {
		panic("escalatorFault of type BACnetEscalatorFaultTagged for BACnetPropertyStatesEscalatorFault must not be nil")
	}
	_result := &_BACnetPropertyStatesEscalatorFault{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		EscalatorFault:               escalatorFault,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesEscalatorFaultBuilder is a builder for BACnetPropertyStatesEscalatorFault
type BACnetPropertyStatesEscalatorFaultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(escalatorFault BACnetEscalatorFaultTagged) BACnetPropertyStatesEscalatorFaultBuilder
	// WithEscalatorFault adds EscalatorFault (property field)
	WithEscalatorFault(BACnetEscalatorFaultTagged) BACnetPropertyStatesEscalatorFaultBuilder
	// WithEscalatorFaultBuilder adds EscalatorFault (property field) which is build by the builder
	WithEscalatorFaultBuilder(func(BACnetEscalatorFaultTaggedBuilder) BACnetEscalatorFaultTaggedBuilder) BACnetPropertyStatesEscalatorFaultBuilder
	// Build builds the BACnetPropertyStatesEscalatorFault or returns an error if something is wrong
	Build() (BACnetPropertyStatesEscalatorFault, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesEscalatorFault
}

// NewBACnetPropertyStatesEscalatorFaultBuilder() creates a BACnetPropertyStatesEscalatorFaultBuilder
func NewBACnetPropertyStatesEscalatorFaultBuilder() BACnetPropertyStatesEscalatorFaultBuilder {
	return &_BACnetPropertyStatesEscalatorFaultBuilder{_BACnetPropertyStatesEscalatorFault: new(_BACnetPropertyStatesEscalatorFault)}
}

type _BACnetPropertyStatesEscalatorFaultBuilder struct {
	*_BACnetPropertyStatesEscalatorFault

	err *utils.MultiError
}

var _ (BACnetPropertyStatesEscalatorFaultBuilder) = (*_BACnetPropertyStatesEscalatorFaultBuilder)(nil)

func (m *_BACnetPropertyStatesEscalatorFaultBuilder) WithMandatoryFields(escalatorFault BACnetEscalatorFaultTagged) BACnetPropertyStatesEscalatorFaultBuilder {
	return m.WithEscalatorFault(escalatorFault)
}

func (m *_BACnetPropertyStatesEscalatorFaultBuilder) WithEscalatorFault(escalatorFault BACnetEscalatorFaultTagged) BACnetPropertyStatesEscalatorFaultBuilder {
	m.EscalatorFault = escalatorFault
	return m
}

func (m *_BACnetPropertyStatesEscalatorFaultBuilder) WithEscalatorFaultBuilder(builderSupplier func(BACnetEscalatorFaultTaggedBuilder) BACnetEscalatorFaultTaggedBuilder) BACnetPropertyStatesEscalatorFaultBuilder {
	builder := builderSupplier(m.EscalatorFault.CreateBACnetEscalatorFaultTaggedBuilder())
	var err error
	m.EscalatorFault, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetEscalatorFaultTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetPropertyStatesEscalatorFaultBuilder) Build() (BACnetPropertyStatesEscalatorFault, error) {
	if m.EscalatorFault == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'escalatorFault' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetPropertyStatesEscalatorFault.deepCopy(), nil
}

func (m *_BACnetPropertyStatesEscalatorFaultBuilder) MustBuild() BACnetPropertyStatesEscalatorFault {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetPropertyStatesEscalatorFaultBuilder) DeepCopy() any {
	return m.CreateBACnetPropertyStatesEscalatorFaultBuilder()
}

// CreateBACnetPropertyStatesEscalatorFaultBuilder creates a BACnetPropertyStatesEscalatorFaultBuilder
func (m *_BACnetPropertyStatesEscalatorFault) CreateBACnetPropertyStatesEscalatorFaultBuilder() BACnetPropertyStatesEscalatorFaultBuilder {
	if m == nil {
		return NewBACnetPropertyStatesEscalatorFaultBuilder()
	}
	return &_BACnetPropertyStatesEscalatorFaultBuilder{_BACnetPropertyStatesEscalatorFault: m.deepCopy()}
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

func (m *_BACnetPropertyStatesEscalatorFault) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesEscalatorFault) GetEscalatorFault() BACnetEscalatorFaultTagged {
	return m.EscalatorFault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesEscalatorFault(structType any) BACnetPropertyStatesEscalatorFault {
	if casted, ok := structType.(BACnetPropertyStatesEscalatorFault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesEscalatorFault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesEscalatorFault) GetTypeName() string {
	return "BACnetPropertyStatesEscalatorFault"
}

func (m *_BACnetPropertyStatesEscalatorFault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (escalatorFault)
	lengthInBits += m.EscalatorFault.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesEscalatorFault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesEscalatorFault) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesEscalatorFault BACnetPropertyStatesEscalatorFault, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesEscalatorFault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesEscalatorFault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	escalatorFault, err := ReadSimpleField[BACnetEscalatorFaultTagged](ctx, "escalatorFault", ReadComplex[BACnetEscalatorFaultTagged](BACnetEscalatorFaultTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'escalatorFault' field"))
	}
	m.EscalatorFault = escalatorFault

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesEscalatorFault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesEscalatorFault")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesEscalatorFault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesEscalatorFault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesEscalatorFault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesEscalatorFault")
		}

		if err := WriteSimpleField[BACnetEscalatorFaultTagged](ctx, "escalatorFault", m.GetEscalatorFault(), WriteComplex[BACnetEscalatorFaultTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'escalatorFault' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesEscalatorFault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesEscalatorFault")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesEscalatorFault) IsBACnetPropertyStatesEscalatorFault() {}

func (m *_BACnetPropertyStatesEscalatorFault) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesEscalatorFault) deepCopy() *_BACnetPropertyStatesEscalatorFault {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesEscalatorFaultCopy := &_BACnetPropertyStatesEscalatorFault{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.EscalatorFault.DeepCopy().(BACnetEscalatorFaultTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesEscalatorFaultCopy
}

func (m *_BACnetPropertyStatesEscalatorFault) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
