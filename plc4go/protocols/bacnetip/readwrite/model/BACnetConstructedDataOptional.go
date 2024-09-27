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

// BACnetConstructedDataOptional is the corresponding interface of BACnetConstructedDataOptional
type BACnetConstructedDataOptional interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataOptional is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataOptional()
	// CreateBuilder creates a BACnetConstructedDataOptionalBuilder
	CreateBACnetConstructedDataOptionalBuilder() BACnetConstructedDataOptionalBuilder
}

// _BACnetConstructedDataOptional is the data-structure of this message
type _BACnetConstructedDataOptional struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataOptional = (*_BACnetConstructedDataOptional)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataOptional)(nil)

// NewBACnetConstructedDataOptional factory function for _BACnetConstructedDataOptional
func NewBACnetConstructedDataOptional(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOptional {
	_result := &_BACnetConstructedDataOptional{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataOptionalBuilder is a builder for BACnetConstructedDataOptional
type BACnetConstructedDataOptionalBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataOptionalBuilder
	// Build builds the BACnetConstructedDataOptional or returns an error if something is wrong
	Build() (BACnetConstructedDataOptional, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataOptional
}

// NewBACnetConstructedDataOptionalBuilder() creates a BACnetConstructedDataOptionalBuilder
func NewBACnetConstructedDataOptionalBuilder() BACnetConstructedDataOptionalBuilder {
	return &_BACnetConstructedDataOptionalBuilder{_BACnetConstructedDataOptional: new(_BACnetConstructedDataOptional)}
}

type _BACnetConstructedDataOptionalBuilder struct {
	*_BACnetConstructedDataOptional

	err *utils.MultiError
}

var _ (BACnetConstructedDataOptionalBuilder) = (*_BACnetConstructedDataOptionalBuilder)(nil)

func (m *_BACnetConstructedDataOptionalBuilder) WithMandatoryFields() BACnetConstructedDataOptionalBuilder {
	return m
}

func (m *_BACnetConstructedDataOptionalBuilder) Build() (BACnetConstructedDataOptional, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataOptional.deepCopy(), nil
}

func (m *_BACnetConstructedDataOptionalBuilder) MustBuild() BACnetConstructedDataOptional {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataOptionalBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataOptionalBuilder()
}

// CreateBACnetConstructedDataOptionalBuilder creates a BACnetConstructedDataOptionalBuilder
func (m *_BACnetConstructedDataOptional) CreateBACnetConstructedDataOptionalBuilder() BACnetConstructedDataOptionalBuilder {
	if m == nil {
		return NewBACnetConstructedDataOptionalBuilder()
	}
	return &_BACnetConstructedDataOptionalBuilder{_BACnetConstructedDataOptional: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOptional) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOptional) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OPTIONAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOptional) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOptional(structType any) BACnetConstructedDataOptional {
	if casted, ok := structType.(BACnetConstructedDataOptional); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOptional); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOptional) GetTypeName() string {
	return "BACnetConstructedDataOptional"
}

func (m *_BACnetConstructedDataOptional) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataOptional) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataOptional) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataOptional BACnetConstructedDataOptional, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOptional"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOptional")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "An property identified by OPTIONAL should never occur in the wild"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOptional"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOptional")
	}

	return m, nil
}

func (m *_BACnetConstructedDataOptional) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOptional) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOptional"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOptional")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOptional"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOptional")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOptional) IsBACnetConstructedDataOptional() {}

func (m *_BACnetConstructedDataOptional) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataOptional) deepCopy() *_BACnetConstructedDataOptional {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataOptionalCopy := &_BACnetConstructedDataOptional{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataOptionalCopy
}

func (m *_BACnetConstructedDataOptional) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
