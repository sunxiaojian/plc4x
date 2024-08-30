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

// RegisterNodesRequest is the corresponding interface of RegisterNodesRequest
type RegisterNodesRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfNodesToRegister returns NoOfNodesToRegister (property field)
	GetNoOfNodesToRegister() int32
	// GetNodesToRegister returns NodesToRegister (property field)
	GetNodesToRegister() []NodeId
}

// RegisterNodesRequestExactly can be used when we want exactly this type and not a type which fulfills RegisterNodesRequest.
// This is useful for switch cases.
type RegisterNodesRequestExactly interface {
	RegisterNodesRequest
	isRegisterNodesRequest() bool
}

// _RegisterNodesRequest is the data-structure of this message
type _RegisterNodesRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader       ExtensionObjectDefinition
	NoOfNodesToRegister int32
	NodesToRegister     []NodeId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RegisterNodesRequest) GetIdentifier() string {
	return "560"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RegisterNodesRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_RegisterNodesRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RegisterNodesRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_RegisterNodesRequest) GetNoOfNodesToRegister() int32 {
	return m.NoOfNodesToRegister
}

func (m *_RegisterNodesRequest) GetNodesToRegister() []NodeId {
	return m.NodesToRegister
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRegisterNodesRequest factory function for _RegisterNodesRequest
func NewRegisterNodesRequest(requestHeader ExtensionObjectDefinition, noOfNodesToRegister int32, nodesToRegister []NodeId) *_RegisterNodesRequest {
	_result := &_RegisterNodesRequest{
		RequestHeader:              requestHeader,
		NoOfNodesToRegister:        noOfNodesToRegister,
		NodesToRegister:            nodesToRegister,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRegisterNodesRequest(structType any) RegisterNodesRequest {
	if casted, ok := structType.(RegisterNodesRequest); ok {
		return casted
	}
	if casted, ok := structType.(*RegisterNodesRequest); ok {
		return *casted
	}
	return nil
}

func (m *_RegisterNodesRequest) GetTypeName() string {
	return "RegisterNodesRequest"
}

func (m *_RegisterNodesRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfNodesToRegister)
	lengthInBits += 32

	// Array field
	if len(m.NodesToRegister) > 0 {
		for _curItem, element := range m.NodesToRegister {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodesToRegister), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_RegisterNodesRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RegisterNodesRequestParse(ctx context.Context, theBytes []byte, identifier string) (RegisterNodesRequest, error) {
	return RegisterNodesRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func RegisterNodesRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (RegisterNodesRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("RegisterNodesRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RegisterNodesRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of RegisterNodesRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (noOfNodesToRegister)
	_noOfNodesToRegister, _noOfNodesToRegisterErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfNodesToRegister", 32)
	if _noOfNodesToRegisterErr != nil {
		return nil, errors.Wrap(_noOfNodesToRegisterErr, "Error parsing 'noOfNodesToRegister' field of RegisterNodesRequest")
	}
	noOfNodesToRegister := _noOfNodesToRegister

	// Array field (nodesToRegister)
	if pullErr := readBuffer.PullContext("nodesToRegister", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodesToRegister")
	}
	// Count array
	nodesToRegister := make([]NodeId, max(noOfNodesToRegister, 0))
	// This happens when the size is set conditional to 0
	if len(nodesToRegister) == 0 {
		nodesToRegister = nil
	}
	{
		_numItems := uint16(max(noOfNodesToRegister, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := NodeIdParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'nodesToRegister' field of RegisterNodesRequest")
			}
			nodesToRegister[_curItem] = _item.(NodeId)
		}
	}
	if closeErr := readBuffer.CloseContext("nodesToRegister", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodesToRegister")
	}

	if closeErr := readBuffer.CloseContext("RegisterNodesRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RegisterNodesRequest")
	}

	// Create a partially initialized instance
	_child := &_RegisterNodesRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		NoOfNodesToRegister:        noOfNodesToRegister,
		NodesToRegister:            nodesToRegister,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_RegisterNodesRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RegisterNodesRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RegisterNodesRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RegisterNodesRequest")
		}

		// Simple Field (requestHeader)
		if pushErr := writeBuffer.PushContext("requestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestHeader")
		}
		_requestHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetRequestHeader())
		if popErr := writeBuffer.PopContext("requestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestHeader")
		}
		if _requestHeaderErr != nil {
			return errors.Wrap(_requestHeaderErr, "Error serializing 'requestHeader' field")
		}

		// Simple Field (noOfNodesToRegister)
		noOfNodesToRegister := int32(m.GetNoOfNodesToRegister())
		_noOfNodesToRegisterErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfNodesToRegister", 32, int32((noOfNodesToRegister)))
		if _noOfNodesToRegisterErr != nil {
			return errors.Wrap(_noOfNodesToRegisterErr, "Error serializing 'noOfNodesToRegister' field")
		}

		// Array Field (nodesToRegister)
		if pushErr := writeBuffer.PushContext("nodesToRegister", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nodesToRegister")
		}
		for _curItem, _element := range m.GetNodesToRegister() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetNodesToRegister()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'nodesToRegister' field")
			}
		}
		if popErr := writeBuffer.PopContext("nodesToRegister", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nodesToRegister")
		}

		if popErr := writeBuffer.PopContext("RegisterNodesRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RegisterNodesRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RegisterNodesRequest) isRegisterNodesRequest() bool {
	return true
}

func (m *_RegisterNodesRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
