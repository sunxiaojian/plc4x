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

// BACnetPropertyStatesBinaryLightningValue is the corresponding interface of BACnetPropertyStatesBinaryLightningValue
type BACnetPropertyStatesBinaryLightningValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetBinaryLightningValue returns BinaryLightningValue (property field)
	GetBinaryLightningValue() BACnetBinaryLightingPVTagged
	// IsBACnetPropertyStatesBinaryLightningValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesBinaryLightningValue()
	// CreateBuilder creates a BACnetPropertyStatesBinaryLightningValueBuilder
	CreateBACnetPropertyStatesBinaryLightningValueBuilder() BACnetPropertyStatesBinaryLightningValueBuilder
}

// _BACnetPropertyStatesBinaryLightningValue is the data-structure of this message
type _BACnetPropertyStatesBinaryLightningValue struct {
	BACnetPropertyStatesContract
	BinaryLightningValue BACnetBinaryLightingPVTagged
}

var _ BACnetPropertyStatesBinaryLightningValue = (*_BACnetPropertyStatesBinaryLightningValue)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesBinaryLightningValue)(nil)

// NewBACnetPropertyStatesBinaryLightningValue factory function for _BACnetPropertyStatesBinaryLightningValue
func NewBACnetPropertyStatesBinaryLightningValue(peekedTagHeader BACnetTagHeader, binaryLightningValue BACnetBinaryLightingPVTagged) *_BACnetPropertyStatesBinaryLightningValue {
	if binaryLightningValue == nil {
		panic("binaryLightningValue of type BACnetBinaryLightingPVTagged for BACnetPropertyStatesBinaryLightningValue must not be nil")
	}
	_result := &_BACnetPropertyStatesBinaryLightningValue{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		BinaryLightningValue:         binaryLightningValue,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesBinaryLightningValueBuilder is a builder for BACnetPropertyStatesBinaryLightningValue
type BACnetPropertyStatesBinaryLightningValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(binaryLightningValue BACnetBinaryLightingPVTagged) BACnetPropertyStatesBinaryLightningValueBuilder
	// WithBinaryLightningValue adds BinaryLightningValue (property field)
	WithBinaryLightningValue(BACnetBinaryLightingPVTagged) BACnetPropertyStatesBinaryLightningValueBuilder
	// WithBinaryLightningValueBuilder adds BinaryLightningValue (property field) which is build by the builder
	WithBinaryLightningValueBuilder(func(BACnetBinaryLightingPVTaggedBuilder) BACnetBinaryLightingPVTaggedBuilder) BACnetPropertyStatesBinaryLightningValueBuilder
	// Build builds the BACnetPropertyStatesBinaryLightningValue or returns an error if something is wrong
	Build() (BACnetPropertyStatesBinaryLightningValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesBinaryLightningValue
}

// NewBACnetPropertyStatesBinaryLightningValueBuilder() creates a BACnetPropertyStatesBinaryLightningValueBuilder
func NewBACnetPropertyStatesBinaryLightningValueBuilder() BACnetPropertyStatesBinaryLightningValueBuilder {
	return &_BACnetPropertyStatesBinaryLightningValueBuilder{_BACnetPropertyStatesBinaryLightningValue: new(_BACnetPropertyStatesBinaryLightningValue)}
}

type _BACnetPropertyStatesBinaryLightningValueBuilder struct {
	*_BACnetPropertyStatesBinaryLightningValue

	err *utils.MultiError
}

var _ (BACnetPropertyStatesBinaryLightningValueBuilder) = (*_BACnetPropertyStatesBinaryLightningValueBuilder)(nil)

func (m *_BACnetPropertyStatesBinaryLightningValueBuilder) WithMandatoryFields(binaryLightningValue BACnetBinaryLightingPVTagged) BACnetPropertyStatesBinaryLightningValueBuilder {
	return m.WithBinaryLightningValue(binaryLightningValue)
}

func (m *_BACnetPropertyStatesBinaryLightningValueBuilder) WithBinaryLightningValue(binaryLightningValue BACnetBinaryLightingPVTagged) BACnetPropertyStatesBinaryLightningValueBuilder {
	m.BinaryLightningValue = binaryLightningValue
	return m
}

func (m *_BACnetPropertyStatesBinaryLightningValueBuilder) WithBinaryLightningValueBuilder(builderSupplier func(BACnetBinaryLightingPVTaggedBuilder) BACnetBinaryLightingPVTaggedBuilder) BACnetPropertyStatesBinaryLightningValueBuilder {
	builder := builderSupplier(m.BinaryLightningValue.CreateBACnetBinaryLightingPVTaggedBuilder())
	var err error
	m.BinaryLightningValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetBinaryLightingPVTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetPropertyStatesBinaryLightningValueBuilder) Build() (BACnetPropertyStatesBinaryLightningValue, error) {
	if m.BinaryLightningValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'binaryLightningValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetPropertyStatesBinaryLightningValue.deepCopy(), nil
}

func (m *_BACnetPropertyStatesBinaryLightningValueBuilder) MustBuild() BACnetPropertyStatesBinaryLightningValue {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetPropertyStatesBinaryLightningValueBuilder) DeepCopy() any {
	return m.CreateBACnetPropertyStatesBinaryLightningValueBuilder()
}

// CreateBACnetPropertyStatesBinaryLightningValueBuilder creates a BACnetPropertyStatesBinaryLightningValueBuilder
func (m *_BACnetPropertyStatesBinaryLightningValue) CreateBACnetPropertyStatesBinaryLightningValueBuilder() BACnetPropertyStatesBinaryLightningValueBuilder {
	if m == nil {
		return NewBACnetPropertyStatesBinaryLightningValueBuilder()
	}
	return &_BACnetPropertyStatesBinaryLightningValueBuilder{_BACnetPropertyStatesBinaryLightningValue: m.deepCopy()}
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

func (m *_BACnetPropertyStatesBinaryLightningValue) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesBinaryLightningValue) GetBinaryLightningValue() BACnetBinaryLightingPVTagged {
	return m.BinaryLightningValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesBinaryLightningValue(structType any) BACnetPropertyStatesBinaryLightningValue {
	if casted, ok := structType.(BACnetPropertyStatesBinaryLightningValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesBinaryLightningValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesBinaryLightningValue) GetTypeName() string {
	return "BACnetPropertyStatesBinaryLightningValue"
}

func (m *_BACnetPropertyStatesBinaryLightningValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (binaryLightningValue)
	lengthInBits += m.BinaryLightningValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesBinaryLightningValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesBinaryLightningValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesBinaryLightningValue BACnetPropertyStatesBinaryLightningValue, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesBinaryLightningValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesBinaryLightningValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	binaryLightningValue, err := ReadSimpleField[BACnetBinaryLightingPVTagged](ctx, "binaryLightningValue", ReadComplex[BACnetBinaryLightingPVTagged](BACnetBinaryLightingPVTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'binaryLightningValue' field"))
	}
	m.BinaryLightningValue = binaryLightningValue

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesBinaryLightningValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesBinaryLightningValue")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesBinaryLightningValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesBinaryLightningValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesBinaryLightningValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesBinaryLightningValue")
		}

		if err := WriteSimpleField[BACnetBinaryLightingPVTagged](ctx, "binaryLightningValue", m.GetBinaryLightningValue(), WriteComplex[BACnetBinaryLightingPVTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'binaryLightningValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesBinaryLightningValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesBinaryLightningValue")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesBinaryLightningValue) IsBACnetPropertyStatesBinaryLightningValue() {}

func (m *_BACnetPropertyStatesBinaryLightningValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesBinaryLightningValue) deepCopy() *_BACnetPropertyStatesBinaryLightningValue {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesBinaryLightningValueCopy := &_BACnetPropertyStatesBinaryLightningValue{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.BinaryLightningValue.DeepCopy().(BACnetBinaryLightingPVTagged),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesBinaryLightningValueCopy
}

func (m *_BACnetPropertyStatesBinaryLightningValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
