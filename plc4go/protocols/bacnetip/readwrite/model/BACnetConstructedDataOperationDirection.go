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

// BACnetConstructedDataOperationDirection is the corresponding interface of BACnetConstructedDataOperationDirection
type BACnetConstructedDataOperationDirection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetOperationDirection returns OperationDirection (property field)
	GetOperationDirection() BACnetEscalatorOperationDirectionTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEscalatorOperationDirectionTagged
	// IsBACnetConstructedDataOperationDirection is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataOperationDirection()
	// CreateBuilder creates a BACnetConstructedDataOperationDirectionBuilder
	CreateBACnetConstructedDataOperationDirectionBuilder() BACnetConstructedDataOperationDirectionBuilder
}

// _BACnetConstructedDataOperationDirection is the data-structure of this message
type _BACnetConstructedDataOperationDirection struct {
	BACnetConstructedDataContract
	OperationDirection BACnetEscalatorOperationDirectionTagged
}

var _ BACnetConstructedDataOperationDirection = (*_BACnetConstructedDataOperationDirection)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataOperationDirection)(nil)

// NewBACnetConstructedDataOperationDirection factory function for _BACnetConstructedDataOperationDirection
func NewBACnetConstructedDataOperationDirection(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, operationDirection BACnetEscalatorOperationDirectionTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOperationDirection {
	if operationDirection == nil {
		panic("operationDirection of type BACnetEscalatorOperationDirectionTagged for BACnetConstructedDataOperationDirection must not be nil")
	}
	_result := &_BACnetConstructedDataOperationDirection{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		OperationDirection:            operationDirection,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataOperationDirectionBuilder is a builder for BACnetConstructedDataOperationDirection
type BACnetConstructedDataOperationDirectionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(operationDirection BACnetEscalatorOperationDirectionTagged) BACnetConstructedDataOperationDirectionBuilder
	// WithOperationDirection adds OperationDirection (property field)
	WithOperationDirection(BACnetEscalatorOperationDirectionTagged) BACnetConstructedDataOperationDirectionBuilder
	// WithOperationDirectionBuilder adds OperationDirection (property field) which is build by the builder
	WithOperationDirectionBuilder(func(BACnetEscalatorOperationDirectionTaggedBuilder) BACnetEscalatorOperationDirectionTaggedBuilder) BACnetConstructedDataOperationDirectionBuilder
	// Build builds the BACnetConstructedDataOperationDirection or returns an error if something is wrong
	Build() (BACnetConstructedDataOperationDirection, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataOperationDirection
}

// NewBACnetConstructedDataOperationDirectionBuilder() creates a BACnetConstructedDataOperationDirectionBuilder
func NewBACnetConstructedDataOperationDirectionBuilder() BACnetConstructedDataOperationDirectionBuilder {
	return &_BACnetConstructedDataOperationDirectionBuilder{_BACnetConstructedDataOperationDirection: new(_BACnetConstructedDataOperationDirection)}
}

type _BACnetConstructedDataOperationDirectionBuilder struct {
	*_BACnetConstructedDataOperationDirection

	err *utils.MultiError
}

var _ (BACnetConstructedDataOperationDirectionBuilder) = (*_BACnetConstructedDataOperationDirectionBuilder)(nil)

func (m *_BACnetConstructedDataOperationDirectionBuilder) WithMandatoryFields(operationDirection BACnetEscalatorOperationDirectionTagged) BACnetConstructedDataOperationDirectionBuilder {
	return m.WithOperationDirection(operationDirection)
}

func (m *_BACnetConstructedDataOperationDirectionBuilder) WithOperationDirection(operationDirection BACnetEscalatorOperationDirectionTagged) BACnetConstructedDataOperationDirectionBuilder {
	m.OperationDirection = operationDirection
	return m
}

func (m *_BACnetConstructedDataOperationDirectionBuilder) WithOperationDirectionBuilder(builderSupplier func(BACnetEscalatorOperationDirectionTaggedBuilder) BACnetEscalatorOperationDirectionTaggedBuilder) BACnetConstructedDataOperationDirectionBuilder {
	builder := builderSupplier(m.OperationDirection.CreateBACnetEscalatorOperationDirectionTaggedBuilder())
	var err error
	m.OperationDirection, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetEscalatorOperationDirectionTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataOperationDirectionBuilder) Build() (BACnetConstructedDataOperationDirection, error) {
	if m.OperationDirection == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'operationDirection' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataOperationDirection.deepCopy(), nil
}

func (m *_BACnetConstructedDataOperationDirectionBuilder) MustBuild() BACnetConstructedDataOperationDirection {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataOperationDirectionBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataOperationDirectionBuilder()
}

// CreateBACnetConstructedDataOperationDirectionBuilder creates a BACnetConstructedDataOperationDirectionBuilder
func (m *_BACnetConstructedDataOperationDirection) CreateBACnetConstructedDataOperationDirectionBuilder() BACnetConstructedDataOperationDirectionBuilder {
	if m == nil {
		return NewBACnetConstructedDataOperationDirectionBuilder()
	}
	return &_BACnetConstructedDataOperationDirectionBuilder{_BACnetConstructedDataOperationDirection: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOperationDirection) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOperationDirection) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OPERATION_DIRECTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOperationDirection) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOperationDirection) GetOperationDirection() BACnetEscalatorOperationDirectionTagged {
	return m.OperationDirection
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOperationDirection) GetActualValue() BACnetEscalatorOperationDirectionTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetEscalatorOperationDirectionTagged(m.GetOperationDirection())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOperationDirection(structType any) BACnetConstructedDataOperationDirection {
	if casted, ok := structType.(BACnetConstructedDataOperationDirection); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOperationDirection); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOperationDirection) GetTypeName() string {
	return "BACnetConstructedDataOperationDirection"
}

func (m *_BACnetConstructedDataOperationDirection) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (operationDirection)
	lengthInBits += m.OperationDirection.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOperationDirection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataOperationDirection) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataOperationDirection BACnetConstructedDataOperationDirection, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOperationDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOperationDirection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	operationDirection, err := ReadSimpleField[BACnetEscalatorOperationDirectionTagged](ctx, "operationDirection", ReadComplex[BACnetEscalatorOperationDirectionTagged](BACnetEscalatorOperationDirectionTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operationDirection' field"))
	}
	m.OperationDirection = operationDirection

	actualValue, err := ReadVirtualField[BACnetEscalatorOperationDirectionTagged](ctx, "actualValue", (*BACnetEscalatorOperationDirectionTagged)(nil), operationDirection)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOperationDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOperationDirection")
	}

	return m, nil
}

func (m *_BACnetConstructedDataOperationDirection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOperationDirection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOperationDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOperationDirection")
		}

		if err := WriteSimpleField[BACnetEscalatorOperationDirectionTagged](ctx, "operationDirection", m.GetOperationDirection(), WriteComplex[BACnetEscalatorOperationDirectionTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'operationDirection' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOperationDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOperationDirection")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOperationDirection) IsBACnetConstructedDataOperationDirection() {}

func (m *_BACnetConstructedDataOperationDirection) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataOperationDirection) deepCopy() *_BACnetConstructedDataOperationDirection {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataOperationDirectionCopy := &_BACnetConstructedDataOperationDirection{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.OperationDirection.DeepCopy().(BACnetEscalatorOperationDirectionTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataOperationDirectionCopy
}

func (m *_BACnetConstructedDataOperationDirection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
