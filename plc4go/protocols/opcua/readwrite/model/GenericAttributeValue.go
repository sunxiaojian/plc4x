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

// GenericAttributeValue is the corresponding interface of GenericAttributeValue
type GenericAttributeValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetAttributeId returns AttributeId (property field)
	GetAttributeId() uint32
	// GetValue returns Value (property field)
	GetValue() Variant
}

// GenericAttributeValueExactly can be used when we want exactly this type and not a type which fulfills GenericAttributeValue.
// This is useful for switch cases.
type GenericAttributeValueExactly interface {
	GenericAttributeValue
	isGenericAttributeValue() bool
}

// _GenericAttributeValue is the data-structure of this message
type _GenericAttributeValue struct {
	*_ExtensionObjectDefinition
	AttributeId uint32
	Value       Variant
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GenericAttributeValue) GetIdentifier() string {
	return "17608"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GenericAttributeValue) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_GenericAttributeValue) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GenericAttributeValue) GetAttributeId() uint32 {
	return m.AttributeId
}

func (m *_GenericAttributeValue) GetValue() Variant {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewGenericAttributeValue factory function for _GenericAttributeValue
func NewGenericAttributeValue(attributeId uint32, value Variant) *_GenericAttributeValue {
	_result := &_GenericAttributeValue{
		AttributeId:                attributeId,
		Value:                      value,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastGenericAttributeValue(structType any) GenericAttributeValue {
	if casted, ok := structType.(GenericAttributeValue); ok {
		return casted
	}
	if casted, ok := structType.(*GenericAttributeValue); ok {
		return *casted
	}
	return nil
}

func (m *_GenericAttributeValue) GetTypeName() string {
	return "GenericAttributeValue"
}

func (m *_GenericAttributeValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (attributeId)
	lengthInBits += 32

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_GenericAttributeValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func GenericAttributeValueParse(ctx context.Context, theBytes []byte, identifier string) (GenericAttributeValue, error) {
	return GenericAttributeValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func GenericAttributeValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (GenericAttributeValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("GenericAttributeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GenericAttributeValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (attributeId)
	_attributeId, _attributeIdErr := /*TODO: migrate me*/ readBuffer.ReadUint32("attributeId", 32)
	if _attributeIdErr != nil {
		return nil, errors.Wrap(_attributeIdErr, "Error parsing 'attributeId' field of GenericAttributeValue")
	}
	attributeId := _attributeId

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
	_value, _valueErr := VariantParseWithBuffer(ctx, readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of GenericAttributeValue")
	}
	value := _value.(Variant)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}

	if closeErr := readBuffer.CloseContext("GenericAttributeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GenericAttributeValue")
	}

	// Create a partially initialized instance
	_child := &_GenericAttributeValue{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		AttributeId:                attributeId,
		Value:                      value,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_GenericAttributeValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GenericAttributeValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GenericAttributeValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GenericAttributeValue")
		}

		// Simple Field (attributeId)
		attributeId := uint32(m.GetAttributeId())
		_attributeIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("attributeId", 32, uint32((attributeId)))
		if _attributeIdErr != nil {
			return errors.Wrap(_attributeIdErr, "Error serializing 'attributeId' field")
		}

		// Simple Field (value)
		if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for value")
		}
		_valueErr := writeBuffer.WriteSerializable(ctx, m.GetValue())
		if popErr := writeBuffer.PopContext("value"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for value")
		}
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("GenericAttributeValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GenericAttributeValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GenericAttributeValue) isGenericAttributeValue() bool {
	return true
}

func (m *_GenericAttributeValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
