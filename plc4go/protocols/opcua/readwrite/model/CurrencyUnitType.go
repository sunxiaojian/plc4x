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

// CurrencyUnitType is the corresponding interface of CurrencyUnitType
type CurrencyUnitType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNumericCode returns NumericCode (property field)
	GetNumericCode() int16
	// GetExponent returns Exponent (property field)
	GetExponent() int8
	// GetAlphabeticCode returns AlphabeticCode (property field)
	GetAlphabeticCode() PascalString
	// GetCurrency returns Currency (property field)
	GetCurrency() LocalizedText
}

// CurrencyUnitTypeExactly can be used when we want exactly this type and not a type which fulfills CurrencyUnitType.
// This is useful for switch cases.
type CurrencyUnitTypeExactly interface {
	CurrencyUnitType
	isCurrencyUnitType() bool
}

// _CurrencyUnitType is the data-structure of this message
type _CurrencyUnitType struct {
	*_ExtensionObjectDefinition
	NumericCode    int16
	Exponent       int8
	AlphabeticCode PascalString
	Currency       LocalizedText
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CurrencyUnitType) GetIdentifier() string {
	return "23500"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CurrencyUnitType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_CurrencyUnitType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CurrencyUnitType) GetNumericCode() int16 {
	return m.NumericCode
}

func (m *_CurrencyUnitType) GetExponent() int8 {
	return m.Exponent
}

func (m *_CurrencyUnitType) GetAlphabeticCode() PascalString {
	return m.AlphabeticCode
}

func (m *_CurrencyUnitType) GetCurrency() LocalizedText {
	return m.Currency
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCurrencyUnitType factory function for _CurrencyUnitType
func NewCurrencyUnitType(numericCode int16, exponent int8, alphabeticCode PascalString, currency LocalizedText) *_CurrencyUnitType {
	_result := &_CurrencyUnitType{
		NumericCode:                numericCode,
		Exponent:                   exponent,
		AlphabeticCode:             alphabeticCode,
		Currency:                   currency,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCurrencyUnitType(structType any) CurrencyUnitType {
	if casted, ok := structType.(CurrencyUnitType); ok {
		return casted
	}
	if casted, ok := structType.(*CurrencyUnitType); ok {
		return *casted
	}
	return nil
}

func (m *_CurrencyUnitType) GetTypeName() string {
	return "CurrencyUnitType"
}

func (m *_CurrencyUnitType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (numericCode)
	lengthInBits += 16

	// Simple field (exponent)
	lengthInBits += 8

	// Simple field (alphabeticCode)
	lengthInBits += m.AlphabeticCode.GetLengthInBits(ctx)

	// Simple field (currency)
	lengthInBits += m.Currency.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CurrencyUnitType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CurrencyUnitTypeParse(ctx context.Context, theBytes []byte, identifier string) (CurrencyUnitType, error) {
	return CurrencyUnitTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func CurrencyUnitTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (CurrencyUnitType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CurrencyUnitType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CurrencyUnitType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (numericCode)
	_numericCode, _numericCodeErr := /*TODO: migrate me*/ readBuffer.ReadInt16("numericCode", 16)
	if _numericCodeErr != nil {
		return nil, errors.Wrap(_numericCodeErr, "Error parsing 'numericCode' field of CurrencyUnitType")
	}
	numericCode := _numericCode

	// Simple Field (exponent)
	_exponent, _exponentErr := /*TODO: migrate me*/ readBuffer.ReadInt8("exponent", 8)
	if _exponentErr != nil {
		return nil, errors.Wrap(_exponentErr, "Error parsing 'exponent' field of CurrencyUnitType")
	}
	exponent := _exponent

	// Simple Field (alphabeticCode)
	if pullErr := readBuffer.PullContext("alphabeticCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for alphabeticCode")
	}
	_alphabeticCode, _alphabeticCodeErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _alphabeticCodeErr != nil {
		return nil, errors.Wrap(_alphabeticCodeErr, "Error parsing 'alphabeticCode' field of CurrencyUnitType")
	}
	alphabeticCode := _alphabeticCode.(PascalString)
	if closeErr := readBuffer.CloseContext("alphabeticCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for alphabeticCode")
	}

	// Simple Field (currency)
	if pullErr := readBuffer.PullContext("currency"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for currency")
	}
	_currency, _currencyErr := LocalizedTextParseWithBuffer(ctx, readBuffer)
	if _currencyErr != nil {
		return nil, errors.Wrap(_currencyErr, "Error parsing 'currency' field of CurrencyUnitType")
	}
	currency := _currency.(LocalizedText)
	if closeErr := readBuffer.CloseContext("currency"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for currency")
	}

	if closeErr := readBuffer.CloseContext("CurrencyUnitType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CurrencyUnitType")
	}

	// Create a partially initialized instance
	_child := &_CurrencyUnitType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NumericCode:                numericCode,
		Exponent:                   exponent,
		AlphabeticCode:             alphabeticCode,
		Currency:                   currency,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_CurrencyUnitType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CurrencyUnitType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CurrencyUnitType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CurrencyUnitType")
		}

		// Simple Field (numericCode)
		numericCode := int16(m.GetNumericCode())
		_numericCodeErr := /*TODO: migrate me*/ writeBuffer.WriteInt16("numericCode", 16, int16((numericCode)))
		if _numericCodeErr != nil {
			return errors.Wrap(_numericCodeErr, "Error serializing 'numericCode' field")
		}

		// Simple Field (exponent)
		exponent := int8(m.GetExponent())
		_exponentErr := /*TODO: migrate me*/ writeBuffer.WriteInt8("exponent", 8, int8((exponent)))
		if _exponentErr != nil {
			return errors.Wrap(_exponentErr, "Error serializing 'exponent' field")
		}

		// Simple Field (alphabeticCode)
		if pushErr := writeBuffer.PushContext("alphabeticCode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for alphabeticCode")
		}
		_alphabeticCodeErr := writeBuffer.WriteSerializable(ctx, m.GetAlphabeticCode())
		if popErr := writeBuffer.PopContext("alphabeticCode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for alphabeticCode")
		}
		if _alphabeticCodeErr != nil {
			return errors.Wrap(_alphabeticCodeErr, "Error serializing 'alphabeticCode' field")
		}

		// Simple Field (currency)
		if pushErr := writeBuffer.PushContext("currency"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for currency")
		}
		_currencyErr := writeBuffer.WriteSerializable(ctx, m.GetCurrency())
		if popErr := writeBuffer.PopContext("currency"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for currency")
		}
		if _currencyErr != nil {
			return errors.Wrap(_currencyErr, "Error serializing 'currency' field")
		}

		if popErr := writeBuffer.PopContext("CurrencyUnitType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CurrencyUnitType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CurrencyUnitType) isCurrencyUnitType() bool {
	return true
}

func (m *_CurrencyUnitType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
