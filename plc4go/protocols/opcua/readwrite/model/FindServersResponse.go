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

// FindServersResponse is the corresponding interface of FindServersResponse
type FindServersResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfServers returns NoOfServers (property field)
	GetNoOfServers() int32
	// GetServers returns Servers (property field)
	GetServers() []ExtensionObjectDefinition
}

// FindServersResponseExactly can be used when we want exactly this type and not a type which fulfills FindServersResponse.
// This is useful for switch cases.
type FindServersResponseExactly interface {
	FindServersResponse
	isFindServersResponse() bool
}

// _FindServersResponse is the data-structure of this message
type _FindServersResponse struct {
	*_ExtensionObjectDefinition
	ResponseHeader ExtensionObjectDefinition
	NoOfServers    int32
	Servers        []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FindServersResponse) GetIdentifier() string {
	return "425"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FindServersResponse) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_FindServersResponse) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FindServersResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_FindServersResponse) GetNoOfServers() int32 {
	return m.NoOfServers
}

func (m *_FindServersResponse) GetServers() []ExtensionObjectDefinition {
	return m.Servers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFindServersResponse factory function for _FindServersResponse
func NewFindServersResponse(responseHeader ExtensionObjectDefinition, noOfServers int32, servers []ExtensionObjectDefinition) *_FindServersResponse {
	_result := &_FindServersResponse{
		ResponseHeader:             responseHeader,
		NoOfServers:                noOfServers,
		Servers:                    servers,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFindServersResponse(structType any) FindServersResponse {
	if casted, ok := structType.(FindServersResponse); ok {
		return casted
	}
	if casted, ok := structType.(*FindServersResponse); ok {
		return *casted
	}
	return nil
}

func (m *_FindServersResponse) GetTypeName() string {
	return "FindServersResponse"
}

func (m *_FindServersResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (noOfServers)
	lengthInBits += 32

	// Array field
	if len(m.Servers) > 0 {
		for _curItem, element := range m.Servers {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Servers), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_FindServersResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func FindServersResponseParse(ctx context.Context, theBytes []byte, identifier string) (FindServersResponse, error) {
	return FindServersResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func FindServersResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (FindServersResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("FindServersResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FindServersResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (responseHeader)
	if pullErr := readBuffer.PullContext("responseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for responseHeader")
	}
	_responseHeader, _responseHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("394"))
	if _responseHeaderErr != nil {
		return nil, errors.Wrap(_responseHeaderErr, "Error parsing 'responseHeader' field of FindServersResponse")
	}
	responseHeader := _responseHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("responseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for responseHeader")
	}

	// Simple Field (noOfServers)
	_noOfServers, _noOfServersErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfServers", 32)
	if _noOfServersErr != nil {
		return nil, errors.Wrap(_noOfServersErr, "Error parsing 'noOfServers' field of FindServersResponse")
	}
	noOfServers := _noOfServers

	// Array field (servers)
	if pullErr := readBuffer.PullContext("servers", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for servers")
	}
	// Count array
	servers := make([]ExtensionObjectDefinition, max(noOfServers, 0))
	// This happens when the size is set conditional to 0
	if len(servers) == 0 {
		servers = nil
	}
	{
		_numItems := uint16(max(noOfServers, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "310")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'servers' field of FindServersResponse")
			}
			servers[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("servers", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for servers")
	}

	if closeErr := readBuffer.CloseContext("FindServersResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FindServersResponse")
	}

	// Create a partially initialized instance
	_child := &_FindServersResponse{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ResponseHeader:             responseHeader,
		NoOfServers:                noOfServers,
		Servers:                    servers,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_FindServersResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FindServersResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FindServersResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FindServersResponse")
		}

		// Simple Field (responseHeader)
		if pushErr := writeBuffer.PushContext("responseHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for responseHeader")
		}
		_responseHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetResponseHeader())
		if popErr := writeBuffer.PopContext("responseHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for responseHeader")
		}
		if _responseHeaderErr != nil {
			return errors.Wrap(_responseHeaderErr, "Error serializing 'responseHeader' field")
		}

		// Simple Field (noOfServers)
		noOfServers := int32(m.GetNoOfServers())
		_noOfServersErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfServers", 32, int32((noOfServers)))
		if _noOfServersErr != nil {
			return errors.Wrap(_noOfServersErr, "Error serializing 'noOfServers' field")
		}

		// Array Field (servers)
		if pushErr := writeBuffer.PushContext("servers", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for servers")
		}
		for _curItem, _element := range m.GetServers() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetServers()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'servers' field")
			}
		}
		if popErr := writeBuffer.PopContext("servers", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for servers")
		}

		if popErr := writeBuffer.PopContext("FindServersResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FindServersResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FindServersResponse) isFindServersResponse() bool {
	return true
}

func (m *_FindServersResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
