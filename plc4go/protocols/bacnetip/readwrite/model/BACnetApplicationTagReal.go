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

// BACnetApplicationTagReal is the corresponding interface of BACnetApplicationTagReal
type BACnetApplicationTagReal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() float32
	// IsBACnetApplicationTagReal is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetApplicationTagReal()
	// CreateBuilder creates a BACnetApplicationTagRealBuilder
	CreateBACnetApplicationTagRealBuilder() BACnetApplicationTagRealBuilder
}

// _BACnetApplicationTagReal is the data-structure of this message
type _BACnetApplicationTagReal struct {
	BACnetApplicationTagContract
	Payload BACnetTagPayloadReal
}

var _ BACnetApplicationTagReal = (*_BACnetApplicationTagReal)(nil)
var _ BACnetApplicationTagRequirements = (*_BACnetApplicationTagReal)(nil)

// NewBACnetApplicationTagReal factory function for _BACnetApplicationTagReal
func NewBACnetApplicationTagReal(header BACnetTagHeader, payload BACnetTagPayloadReal) *_BACnetApplicationTagReal {
	if payload == nil {
		panic("payload of type BACnetTagPayloadReal for BACnetApplicationTagReal must not be nil")
	}
	_result := &_BACnetApplicationTagReal{
		BACnetApplicationTagContract: NewBACnetApplicationTag(header),
		Payload:                      payload,
	}
	_result.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetApplicationTagRealBuilder is a builder for BACnetApplicationTagReal
type BACnetApplicationTagRealBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload BACnetTagPayloadReal) BACnetApplicationTagRealBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadReal) BACnetApplicationTagRealBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadRealBuilder) BACnetTagPayloadRealBuilder) BACnetApplicationTagRealBuilder
	// Build builds the BACnetApplicationTagReal or returns an error if something is wrong
	Build() (BACnetApplicationTagReal, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetApplicationTagReal
}

// NewBACnetApplicationTagRealBuilder() creates a BACnetApplicationTagRealBuilder
func NewBACnetApplicationTagRealBuilder() BACnetApplicationTagRealBuilder {
	return &_BACnetApplicationTagRealBuilder{_BACnetApplicationTagReal: new(_BACnetApplicationTagReal)}
}

type _BACnetApplicationTagRealBuilder struct {
	*_BACnetApplicationTagReal

	err *utils.MultiError
}

var _ (BACnetApplicationTagRealBuilder) = (*_BACnetApplicationTagRealBuilder)(nil)

func (m *_BACnetApplicationTagRealBuilder) WithMandatoryFields(payload BACnetTagPayloadReal) BACnetApplicationTagRealBuilder {
	return m.WithPayload(payload)
}

func (m *_BACnetApplicationTagRealBuilder) WithPayload(payload BACnetTagPayloadReal) BACnetApplicationTagRealBuilder {
	m.Payload = payload
	return m
}

func (m *_BACnetApplicationTagRealBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadRealBuilder) BACnetTagPayloadRealBuilder) BACnetApplicationTagRealBuilder {
	builder := builderSupplier(m.Payload.CreateBACnetTagPayloadRealBuilder())
	var err error
	m.Payload, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetTagPayloadRealBuilder failed"))
	}
	return m
}

func (m *_BACnetApplicationTagRealBuilder) Build() (BACnetApplicationTagReal, error) {
	if m.Payload == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetApplicationTagReal.deepCopy(), nil
}

func (m *_BACnetApplicationTagRealBuilder) MustBuild() BACnetApplicationTagReal {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetApplicationTagRealBuilder) DeepCopy() any {
	return m.CreateBACnetApplicationTagRealBuilder()
}

// CreateBACnetApplicationTagRealBuilder creates a BACnetApplicationTagRealBuilder
func (m *_BACnetApplicationTagReal) CreateBACnetApplicationTagRealBuilder() BACnetApplicationTagRealBuilder {
	if m == nil {
		return NewBACnetApplicationTagRealBuilder()
	}
	return &_BACnetApplicationTagRealBuilder{_BACnetApplicationTagReal: m.deepCopy()}
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

func (m *_BACnetApplicationTagReal) GetParent() BACnetApplicationTagContract {
	return m.BACnetApplicationTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagReal) GetPayload() BACnetTagPayloadReal {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetApplicationTagReal) GetActualValue() float32 {
	ctx := context.Background()
	_ = ctx
	return float32(m.GetPayload().GetValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagReal(structType any) BACnetApplicationTagReal {
	if casted, ok := structType.(BACnetApplicationTagReal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagReal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagReal) GetTypeName() string {
	return "BACnetApplicationTagReal"
}

func (m *_BACnetApplicationTagReal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetApplicationTagContract.(*_BACnetApplicationTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetApplicationTagReal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetApplicationTagReal) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetApplicationTag) (__bACnetApplicationTagReal BACnetApplicationTagReal, err error) {
	m.BACnetApplicationTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagReal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagReal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadReal](ctx, "payload", ReadComplex[BACnetTagPayloadReal](BACnetTagPayloadRealParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	actualValue, err := ReadVirtualField[float32](ctx, "actualValue", (*float32)(nil), payload.GetValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagReal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagReal")
	}

	return m, nil
}

func (m *_BACnetApplicationTagReal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagReal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagReal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagReal")
		}

		if err := WriteSimpleField[BACnetTagPayloadReal](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagReal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagReal")
		}
		return nil
	}
	return m.BACnetApplicationTagContract.(*_BACnetApplicationTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagReal) IsBACnetApplicationTagReal() {}

func (m *_BACnetApplicationTagReal) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetApplicationTagReal) deepCopy() *_BACnetApplicationTagReal {
	if m == nil {
		return nil
	}
	_BACnetApplicationTagRealCopy := &_BACnetApplicationTagReal{
		m.BACnetApplicationTagContract.(*_BACnetApplicationTag).deepCopy(),
		m.Payload.DeepCopy().(BACnetTagPayloadReal),
	}
	m.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = m
	return _BACnetApplicationTagRealCopy
}

func (m *_BACnetApplicationTagReal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
