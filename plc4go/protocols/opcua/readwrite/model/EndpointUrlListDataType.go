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

// EndpointUrlListDataType is the corresponding interface of EndpointUrlListDataType
type EndpointUrlListDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNoOfEndpointUrlList returns NoOfEndpointUrlList (property field)
	GetNoOfEndpointUrlList() int32
	// GetEndpointUrlList returns EndpointUrlList (property field)
	GetEndpointUrlList() []PascalString
}

// EndpointUrlListDataTypeExactly can be used when we want exactly this type and not a type which fulfills EndpointUrlListDataType.
// This is useful for switch cases.
type EndpointUrlListDataTypeExactly interface {
	EndpointUrlListDataType
	isEndpointUrlListDataType() bool
}

// _EndpointUrlListDataType is the data-structure of this message
type _EndpointUrlListDataType struct {
	*_ExtensionObjectDefinition
	NoOfEndpointUrlList int32
	EndpointUrlList     []PascalString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EndpointUrlListDataType) GetIdentifier() string {
	return "11945"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EndpointUrlListDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_EndpointUrlListDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EndpointUrlListDataType) GetNoOfEndpointUrlList() int32 {
	return m.NoOfEndpointUrlList
}

func (m *_EndpointUrlListDataType) GetEndpointUrlList() []PascalString {
	return m.EndpointUrlList
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEndpointUrlListDataType factory function for _EndpointUrlListDataType
func NewEndpointUrlListDataType(noOfEndpointUrlList int32, endpointUrlList []PascalString) *_EndpointUrlListDataType {
	_result := &_EndpointUrlListDataType{
		NoOfEndpointUrlList:        noOfEndpointUrlList,
		EndpointUrlList:            endpointUrlList,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastEndpointUrlListDataType(structType any) EndpointUrlListDataType {
	if casted, ok := structType.(EndpointUrlListDataType); ok {
		return casted
	}
	if casted, ok := structType.(*EndpointUrlListDataType); ok {
		return *casted
	}
	return nil
}

func (m *_EndpointUrlListDataType) GetTypeName() string {
	return "EndpointUrlListDataType"
}

func (m *_EndpointUrlListDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (noOfEndpointUrlList)
	lengthInBits += 32

	// Array field
	if len(m.EndpointUrlList) > 0 {
		for _curItem, element := range m.EndpointUrlList {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.EndpointUrlList), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_EndpointUrlListDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func EndpointUrlListDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (EndpointUrlListDataType, error) {
	return EndpointUrlListDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func EndpointUrlListDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (EndpointUrlListDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("EndpointUrlListDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EndpointUrlListDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (noOfEndpointUrlList)
	_noOfEndpointUrlList, _noOfEndpointUrlListErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfEndpointUrlList", 32)
	if _noOfEndpointUrlListErr != nil {
		return nil, errors.Wrap(_noOfEndpointUrlListErr, "Error parsing 'noOfEndpointUrlList' field of EndpointUrlListDataType")
	}
	noOfEndpointUrlList := _noOfEndpointUrlList

	// Array field (endpointUrlList)
	if pullErr := readBuffer.PullContext("endpointUrlList", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for endpointUrlList")
	}
	// Count array
	endpointUrlList := make([]PascalString, max(noOfEndpointUrlList, 0))
	// This happens when the size is set conditional to 0
	if len(endpointUrlList) == 0 {
		endpointUrlList = nil
	}
	{
		_numItems := uint16(max(noOfEndpointUrlList, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := PascalStringParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'endpointUrlList' field of EndpointUrlListDataType")
			}
			endpointUrlList[_curItem] = _item.(PascalString)
		}
	}
	if closeErr := readBuffer.CloseContext("endpointUrlList", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for endpointUrlList")
	}

	if closeErr := readBuffer.CloseContext("EndpointUrlListDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EndpointUrlListDataType")
	}

	// Create a partially initialized instance
	_child := &_EndpointUrlListDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NoOfEndpointUrlList:        noOfEndpointUrlList,
		EndpointUrlList:            endpointUrlList,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_EndpointUrlListDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EndpointUrlListDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EndpointUrlListDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EndpointUrlListDataType")
		}

		// Simple Field (noOfEndpointUrlList)
		noOfEndpointUrlList := int32(m.GetNoOfEndpointUrlList())
		_noOfEndpointUrlListErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfEndpointUrlList", 32, int32((noOfEndpointUrlList)))
		if _noOfEndpointUrlListErr != nil {
			return errors.Wrap(_noOfEndpointUrlListErr, "Error serializing 'noOfEndpointUrlList' field")
		}

		// Array Field (endpointUrlList)
		if pushErr := writeBuffer.PushContext("endpointUrlList", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for endpointUrlList")
		}
		for _curItem, _element := range m.GetEndpointUrlList() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetEndpointUrlList()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'endpointUrlList' field")
			}
		}
		if popErr := writeBuffer.PopContext("endpointUrlList", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for endpointUrlList")
		}

		if popErr := writeBuffer.PopContext("EndpointUrlListDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EndpointUrlListDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EndpointUrlListDataType) isEndpointUrlListDataType() bool {
	return true
}

func (m *_EndpointUrlListDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
