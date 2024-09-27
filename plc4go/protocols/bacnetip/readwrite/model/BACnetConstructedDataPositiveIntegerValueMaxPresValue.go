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

// BACnetConstructedDataPositiveIntegerValueMaxPresValue is the corresponding interface of BACnetConstructedDataPositiveIntegerValueMaxPresValue
type BACnetConstructedDataPositiveIntegerValueMaxPresValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMaxPresValue returns MaxPresValue (property field)
	GetMaxPresValue() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataPositiveIntegerValueMaxPresValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataPositiveIntegerValueMaxPresValue()
	// CreateBuilder creates a BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder
	CreateBACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder() BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder
}

// _BACnetConstructedDataPositiveIntegerValueMaxPresValue is the data-structure of this message
type _BACnetConstructedDataPositiveIntegerValueMaxPresValue struct {
	BACnetConstructedDataContract
	MaxPresValue BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataPositiveIntegerValueMaxPresValue = (*_BACnetConstructedDataPositiveIntegerValueMaxPresValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataPositiveIntegerValueMaxPresValue)(nil)

// NewBACnetConstructedDataPositiveIntegerValueMaxPresValue factory function for _BACnetConstructedDataPositiveIntegerValueMaxPresValue
func NewBACnetConstructedDataPositiveIntegerValueMaxPresValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, maxPresValue BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPositiveIntegerValueMaxPresValue {
	if maxPresValue == nil {
		panic("maxPresValue of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataPositiveIntegerValueMaxPresValue must not be nil")
	}
	_result := &_BACnetConstructedDataPositiveIntegerValueMaxPresValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaxPresValue:                  maxPresValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder is a builder for BACnetConstructedDataPositiveIntegerValueMaxPresValue
type BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maxPresValue BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder
	// WithMaxPresValue adds MaxPresValue (property field)
	WithMaxPresValue(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder
	// WithMaxPresValueBuilder adds MaxPresValue (property field) which is build by the builder
	WithMaxPresValueBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder
	// Build builds the BACnetConstructedDataPositiveIntegerValueMaxPresValue or returns an error if something is wrong
	Build() (BACnetConstructedDataPositiveIntegerValueMaxPresValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataPositiveIntegerValueMaxPresValue
}

// NewBACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder() creates a BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder
func NewBACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder() BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder {
	return &_BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder{_BACnetConstructedDataPositiveIntegerValueMaxPresValue: new(_BACnetConstructedDataPositiveIntegerValueMaxPresValue)}
}

type _BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder struct {
	*_BACnetConstructedDataPositiveIntegerValueMaxPresValue

	err *utils.MultiError
}

var _ (BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder) = (*_BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder)(nil)

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder) WithMandatoryFields(maxPresValue BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder {
	return m.WithMaxPresValue(maxPresValue)
}

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder) WithMaxPresValue(maxPresValue BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder {
	m.MaxPresValue = maxPresValue
	return m
}

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder) WithMaxPresValueBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder {
	builder := builderSupplier(m.MaxPresValue.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.MaxPresValue, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder) Build() (BACnetConstructedDataPositiveIntegerValueMaxPresValue, error) {
	if m.MaxPresValue == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'maxPresValue' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataPositiveIntegerValueMaxPresValue.deepCopy(), nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder) MustBuild() BACnetConstructedDataPositiveIntegerValueMaxPresValue {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder()
}

// CreateBACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder creates a BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder
func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValue) CreateBACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder() BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder {
	if m == nil {
		return NewBACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder()
	}
	return &_BACnetConstructedDataPositiveIntegerValueMaxPresValueBuilder{_BACnetConstructedDataPositiveIntegerValueMaxPresValue: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_POSITIVE_INTEGER_VALUE
}

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValue) GetMaxPresValue() BACnetApplicationTagUnsignedInteger {
	return m.MaxPresValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValue) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetMaxPresValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPositiveIntegerValueMaxPresValue(structType any) BACnetConstructedDataPositiveIntegerValueMaxPresValue {
	if casted, ok := structType.(BACnetConstructedDataPositiveIntegerValueMaxPresValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPositiveIntegerValueMaxPresValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValue) GetTypeName() string {
	return "BACnetConstructedDataPositiveIntegerValueMaxPresValue"
}

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maxPresValue)
	lengthInBits += m.MaxPresValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataPositiveIntegerValueMaxPresValue BACnetConstructedDataPositiveIntegerValueMaxPresValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPositiveIntegerValueMaxPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPositiveIntegerValueMaxPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maxPresValue, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxPresValue", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxPresValue' field"))
	}
	m.MaxPresValue = maxPresValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), maxPresValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPositiveIntegerValueMaxPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPositiveIntegerValueMaxPresValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPositiveIntegerValueMaxPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPositiveIntegerValueMaxPresValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxPresValue", m.GetMaxPresValue(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxPresValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPositiveIntegerValueMaxPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPositiveIntegerValueMaxPresValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValue) IsBACnetConstructedDataPositiveIntegerValueMaxPresValue() {
}

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValue) deepCopy() *_BACnetConstructedDataPositiveIntegerValueMaxPresValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataPositiveIntegerValueMaxPresValueCopy := &_BACnetConstructedDataPositiveIntegerValueMaxPresValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.MaxPresValue.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataPositiveIntegerValueMaxPresValueCopy
}

func (m *_BACnetConstructedDataPositiveIntegerValueMaxPresValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
