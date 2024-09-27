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

// BACnetConstructedDataDescriptionOfHalt is the corresponding interface of BACnetConstructedDataDescriptionOfHalt
type BACnetConstructedDataDescriptionOfHalt interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDescriptionForHalt returns DescriptionForHalt (property field)
	GetDescriptionForHalt() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataDescriptionOfHalt is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDescriptionOfHalt()
	// CreateBuilder creates a BACnetConstructedDataDescriptionOfHaltBuilder
	CreateBACnetConstructedDataDescriptionOfHaltBuilder() BACnetConstructedDataDescriptionOfHaltBuilder
}

// _BACnetConstructedDataDescriptionOfHalt is the data-structure of this message
type _BACnetConstructedDataDescriptionOfHalt struct {
	BACnetConstructedDataContract
	DescriptionForHalt BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataDescriptionOfHalt = (*_BACnetConstructedDataDescriptionOfHalt)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDescriptionOfHalt)(nil)

// NewBACnetConstructedDataDescriptionOfHalt factory function for _BACnetConstructedDataDescriptionOfHalt
func NewBACnetConstructedDataDescriptionOfHalt(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, descriptionForHalt BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDescriptionOfHalt {
	if descriptionForHalt == nil {
		panic("descriptionForHalt of type BACnetApplicationTagCharacterString for BACnetConstructedDataDescriptionOfHalt must not be nil")
	}
	_result := &_BACnetConstructedDataDescriptionOfHalt{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DescriptionForHalt:            descriptionForHalt,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDescriptionOfHaltBuilder is a builder for BACnetConstructedDataDescriptionOfHalt
type BACnetConstructedDataDescriptionOfHaltBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(descriptionForHalt BACnetApplicationTagCharacterString) BACnetConstructedDataDescriptionOfHaltBuilder
	// WithDescriptionForHalt adds DescriptionForHalt (property field)
	WithDescriptionForHalt(BACnetApplicationTagCharacterString) BACnetConstructedDataDescriptionOfHaltBuilder
	// WithDescriptionForHaltBuilder adds DescriptionForHalt (property field) which is build by the builder
	WithDescriptionForHaltBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataDescriptionOfHaltBuilder
	// Build builds the BACnetConstructedDataDescriptionOfHalt or returns an error if something is wrong
	Build() (BACnetConstructedDataDescriptionOfHalt, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDescriptionOfHalt
}

// NewBACnetConstructedDataDescriptionOfHaltBuilder() creates a BACnetConstructedDataDescriptionOfHaltBuilder
func NewBACnetConstructedDataDescriptionOfHaltBuilder() BACnetConstructedDataDescriptionOfHaltBuilder {
	return &_BACnetConstructedDataDescriptionOfHaltBuilder{_BACnetConstructedDataDescriptionOfHalt: new(_BACnetConstructedDataDescriptionOfHalt)}
}

type _BACnetConstructedDataDescriptionOfHaltBuilder struct {
	*_BACnetConstructedDataDescriptionOfHalt

	err *utils.MultiError
}

var _ (BACnetConstructedDataDescriptionOfHaltBuilder) = (*_BACnetConstructedDataDescriptionOfHaltBuilder)(nil)

func (m *_BACnetConstructedDataDescriptionOfHaltBuilder) WithMandatoryFields(descriptionForHalt BACnetApplicationTagCharacterString) BACnetConstructedDataDescriptionOfHaltBuilder {
	return m.WithDescriptionForHalt(descriptionForHalt)
}

func (m *_BACnetConstructedDataDescriptionOfHaltBuilder) WithDescriptionForHalt(descriptionForHalt BACnetApplicationTagCharacterString) BACnetConstructedDataDescriptionOfHaltBuilder {
	m.DescriptionForHalt = descriptionForHalt
	return m
}

func (m *_BACnetConstructedDataDescriptionOfHaltBuilder) WithDescriptionForHaltBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataDescriptionOfHaltBuilder {
	builder := builderSupplier(m.DescriptionForHalt.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	m.DescriptionForHalt, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataDescriptionOfHaltBuilder) Build() (BACnetConstructedDataDescriptionOfHalt, error) {
	if m.DescriptionForHalt == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'descriptionForHalt' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataDescriptionOfHalt.deepCopy(), nil
}

func (m *_BACnetConstructedDataDescriptionOfHaltBuilder) MustBuild() BACnetConstructedDataDescriptionOfHalt {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataDescriptionOfHaltBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataDescriptionOfHaltBuilder()
}

// CreateBACnetConstructedDataDescriptionOfHaltBuilder creates a BACnetConstructedDataDescriptionOfHaltBuilder
func (m *_BACnetConstructedDataDescriptionOfHalt) CreateBACnetConstructedDataDescriptionOfHaltBuilder() BACnetConstructedDataDescriptionOfHaltBuilder {
	if m == nil {
		return NewBACnetConstructedDataDescriptionOfHaltBuilder()
	}
	return &_BACnetConstructedDataDescriptionOfHaltBuilder{_BACnetConstructedDataDescriptionOfHalt: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDescriptionOfHalt) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDescriptionOfHalt) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DESCRIPTION_OF_HALT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDescriptionOfHalt) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDescriptionOfHalt) GetDescriptionForHalt() BACnetApplicationTagCharacterString {
	return m.DescriptionForHalt
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDescriptionOfHalt) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetDescriptionForHalt())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDescriptionOfHalt(structType any) BACnetConstructedDataDescriptionOfHalt {
	if casted, ok := structType.(BACnetConstructedDataDescriptionOfHalt); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDescriptionOfHalt); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDescriptionOfHalt) GetTypeName() string {
	return "BACnetConstructedDataDescriptionOfHalt"
}

func (m *_BACnetConstructedDataDescriptionOfHalt) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (descriptionForHalt)
	lengthInBits += m.DescriptionForHalt.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDescriptionOfHalt) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDescriptionOfHalt) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDescriptionOfHalt BACnetConstructedDataDescriptionOfHalt, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDescriptionOfHalt"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDescriptionOfHalt")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	descriptionForHalt, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "descriptionForHalt", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'descriptionForHalt' field"))
	}
	m.DescriptionForHalt = descriptionForHalt

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), descriptionForHalt)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDescriptionOfHalt"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDescriptionOfHalt")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDescriptionOfHalt) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDescriptionOfHalt) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDescriptionOfHalt"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDescriptionOfHalt")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "descriptionForHalt", m.GetDescriptionForHalt(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'descriptionForHalt' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDescriptionOfHalt"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDescriptionOfHalt")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDescriptionOfHalt) IsBACnetConstructedDataDescriptionOfHalt() {}

func (m *_BACnetConstructedDataDescriptionOfHalt) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDescriptionOfHalt) deepCopy() *_BACnetConstructedDataDescriptionOfHalt {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDescriptionOfHaltCopy := &_BACnetConstructedDataDescriptionOfHalt{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.DescriptionForHalt.DeepCopy().(BACnetApplicationTagCharacterString),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDescriptionOfHaltCopy
}

func (m *_BACnetConstructedDataDescriptionOfHalt) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
