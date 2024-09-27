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

// BACnetConstructedDataMultiStateOutputAll is the corresponding interface of BACnetConstructedDataMultiStateOutputAll
type BACnetConstructedDataMultiStateOutputAll interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// IsBACnetConstructedDataMultiStateOutputAll is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMultiStateOutputAll()
	// CreateBuilder creates a BACnetConstructedDataMultiStateOutputAllBuilder
	CreateBACnetConstructedDataMultiStateOutputAllBuilder() BACnetConstructedDataMultiStateOutputAllBuilder
}

// _BACnetConstructedDataMultiStateOutputAll is the data-structure of this message
type _BACnetConstructedDataMultiStateOutputAll struct {
	BACnetConstructedDataContract
}

var _ BACnetConstructedDataMultiStateOutputAll = (*_BACnetConstructedDataMultiStateOutputAll)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMultiStateOutputAll)(nil)

// NewBACnetConstructedDataMultiStateOutputAll factory function for _BACnetConstructedDataMultiStateOutputAll
func NewBACnetConstructedDataMultiStateOutputAll(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMultiStateOutputAll {
	_result := &_BACnetConstructedDataMultiStateOutputAll{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMultiStateOutputAllBuilder is a builder for BACnetConstructedDataMultiStateOutputAll
type BACnetConstructedDataMultiStateOutputAllBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() BACnetConstructedDataMultiStateOutputAllBuilder
	// Build builds the BACnetConstructedDataMultiStateOutputAll or returns an error if something is wrong
	Build() (BACnetConstructedDataMultiStateOutputAll, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMultiStateOutputAll
}

// NewBACnetConstructedDataMultiStateOutputAllBuilder() creates a BACnetConstructedDataMultiStateOutputAllBuilder
func NewBACnetConstructedDataMultiStateOutputAllBuilder() BACnetConstructedDataMultiStateOutputAllBuilder {
	return &_BACnetConstructedDataMultiStateOutputAllBuilder{_BACnetConstructedDataMultiStateOutputAll: new(_BACnetConstructedDataMultiStateOutputAll)}
}

type _BACnetConstructedDataMultiStateOutputAllBuilder struct {
	*_BACnetConstructedDataMultiStateOutputAll

	err *utils.MultiError
}

var _ (BACnetConstructedDataMultiStateOutputAllBuilder) = (*_BACnetConstructedDataMultiStateOutputAllBuilder)(nil)

func (m *_BACnetConstructedDataMultiStateOutputAllBuilder) WithMandatoryFields() BACnetConstructedDataMultiStateOutputAllBuilder {
	return m
}

func (m *_BACnetConstructedDataMultiStateOutputAllBuilder) Build() (BACnetConstructedDataMultiStateOutputAll, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataMultiStateOutputAll.deepCopy(), nil
}

func (m *_BACnetConstructedDataMultiStateOutputAllBuilder) MustBuild() BACnetConstructedDataMultiStateOutputAll {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataMultiStateOutputAllBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataMultiStateOutputAllBuilder()
}

// CreateBACnetConstructedDataMultiStateOutputAllBuilder creates a BACnetConstructedDataMultiStateOutputAllBuilder
func (m *_BACnetConstructedDataMultiStateOutputAll) CreateBACnetConstructedDataMultiStateOutputAllBuilder() BACnetConstructedDataMultiStateOutputAllBuilder {
	if m == nil {
		return NewBACnetConstructedDataMultiStateOutputAllBuilder()
	}
	return &_BACnetConstructedDataMultiStateOutputAllBuilder{_BACnetConstructedDataMultiStateOutputAll: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMultiStateOutputAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_MULTI_STATE_OUTPUT
}

func (m *_BACnetConstructedDataMultiStateOutputAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMultiStateOutputAll) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMultiStateOutputAll(structType any) BACnetConstructedDataMultiStateOutputAll {
	if casted, ok := structType.(BACnetConstructedDataMultiStateOutputAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMultiStateOutputAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMultiStateOutputAll) GetTypeName() string {
	return "BACnetConstructedDataMultiStateOutputAll"
}

func (m *_BACnetConstructedDataMultiStateOutputAll) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_BACnetConstructedDataMultiStateOutputAll) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMultiStateOutputAll) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMultiStateOutputAll BACnetConstructedDataMultiStateOutputAll, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMultiStateOutputAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMultiStateOutputAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMultiStateOutputAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMultiStateOutputAll")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMultiStateOutputAll) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMultiStateOutputAll) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMultiStateOutputAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMultiStateOutputAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMultiStateOutputAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMultiStateOutputAll")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMultiStateOutputAll) IsBACnetConstructedDataMultiStateOutputAll() {}

func (m *_BACnetConstructedDataMultiStateOutputAll) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMultiStateOutputAll) deepCopy() *_BACnetConstructedDataMultiStateOutputAll {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMultiStateOutputAllCopy := &_BACnetConstructedDataMultiStateOutputAll{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMultiStateOutputAllCopy
}

func (m *_BACnetConstructedDataMultiStateOutputAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
