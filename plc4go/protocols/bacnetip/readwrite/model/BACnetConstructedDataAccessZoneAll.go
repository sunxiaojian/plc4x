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

// BACnetConstructedDataAccessZoneAll is the corresponding interface of BACnetConstructedDataAccessZoneAll
type BACnetConstructedDataAccessZoneAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataAccessZoneAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccessZoneAll()
	// CreateBuilder creates a BACnetConstructedDataAccessZoneAllBuilder
	CreateBACnetConstructedDataAccessZoneAllBuilder() BACnetConstructedDataAccessZoneAllBuilder
}

// _BACnetConstructedDataAccessZoneAll is the data-structure of this message
type _BACnetConstructedDataAccessZoneAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataAccessZoneAll = (*_BACnetConstructedDataAccessZoneAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessZoneAll)(nil)

// NewBACnetConstructedDataAccessZoneAll factory function for _BACnetConstructedDataAccessZoneAll
func NewBACnetConstructedDataAccessZoneAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessZoneAll {
	_result := &_BACnetConstructedDataAccessZoneAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAccessZoneAllBuilder is a builder for BACnetConstructedDataAccessZoneAll
type BACnetConstructedDataAccessZoneAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataAccessZoneAllBuilder
	// Build builds the BACnetConstructedDataAccessZoneAll or returns an error if something is wrong
	Build() (BACnetConstructedDataAccessZoneAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAccessZoneAll
}

// NewBACnetConstructedDataAccessZoneAllBuilder() creates a BACnetConstructedDataAccessZoneAllBuilder
func NewBACnetConstructedDataAccessZoneAllBuilder() BACnetConstructedDataAccessZoneAllBuilder {
	return &_BACnetConstructedDataAccessZoneAllBuilder{_BACnetConstructedDataAccessZoneAll: new(_BACnetConstructedDataAccessZoneAll)}
}

type _BACnetConstructedDataAccessZoneAllBuilder struct {
	*_BACnetConstructedDataAccessZoneAll

	err *utils.MultiError
}

var _ (BACnetConstructedDataAccessZoneAllBuilder) = (*_BACnetConstructedDataAccessZoneAllBuilder)(nil)

func (m *_BACnetConstructedDataAccessZoneAllBuilder) WithMandatoryFields() BACnetConstructedDataAccessZoneAllBuilder {
	return m
}

func (m *_BACnetConstructedDataAccessZoneAllBuilder) Build() (BACnetConstructedDataAccessZoneAll, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataAccessZoneAll.deepCopy(), nil
}

func (m *_BACnetConstructedDataAccessZoneAllBuilder) MustBuild() BACnetConstructedDataAccessZoneAll {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataAccessZoneAllBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataAccessZoneAllBuilder()
}

// CreateBACnetConstructedDataAccessZoneAllBuilder creates a BACnetConstructedDataAccessZoneAllBuilder
func (m *_BACnetConstructedDataAccessZoneAll) CreateBACnetConstructedDataAccessZoneAllBuilder() BACnetConstructedDataAccessZoneAllBuilder {
	if m == nil {
		return NewBACnetConstructedDataAccessZoneAllBuilder()
	}
	return &_BACnetConstructedDataAccessZoneAllBuilder{_BACnetConstructedDataAccessZoneAll: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessZoneAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCESS_ZONE
}

func (m *_BACnetConstructedDataAccessZoneAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessZoneAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessZoneAll(structType any) BACnetConstructedDataAccessZoneAll {
	if casted, ok := structType.(BACnetConstructedDataAccessZoneAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessZoneAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessZoneAll) GetTypeName() string {
	return "BACnetConstructedDataAccessZoneAll"
}

func (m *_BACnetConstructedDataAccessZoneAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessZoneAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessZoneAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessZoneAll BACnetConstructedDataAccessZoneAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessZoneAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessZoneAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessZoneAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessZoneAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessZoneAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessZoneAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessZoneAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessZoneAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessZoneAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessZoneAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessZoneAll) IsBACnetConstructedDataAccessZoneAll() {}

func (m *_BACnetConstructedDataAccessZoneAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccessZoneAll) deepCopy() *_BACnetConstructedDataAccessZoneAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccessZoneAllCopy := &_BACnetConstructedDataAccessZoneAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccessZoneAllCopy
}

func (m *_BACnetConstructedDataAccessZoneAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
