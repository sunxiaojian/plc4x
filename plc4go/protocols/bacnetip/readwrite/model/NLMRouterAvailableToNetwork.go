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

// NLMRouterAvailableToNetwork is the corresponding interface of NLMRouterAvailableToNetwork
type NLMRouterAvailableToNetwork interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetDestinationNetworkAddresses returns DestinationNetworkAddresses (property field)
	GetDestinationNetworkAddresses() []uint16
}

// NLMRouterAvailableToNetworkExactly can be used when we want exactly this type and not a type which fulfills NLMRouterAvailableToNetwork.
// This is useful for switch cases.
type NLMRouterAvailableToNetworkExactly interface {
	NLMRouterAvailableToNetwork
	isNLMRouterAvailableToNetwork() bool
}

// _NLMRouterAvailableToNetwork is the data-structure of this message
type _NLMRouterAvailableToNetwork struct {
	*_NLM
	DestinationNetworkAddresses []uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMRouterAvailableToNetwork) GetMessageType() uint8 {
	return 0x05
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMRouterAvailableToNetwork) InitializeParent(parent NLM) {}

func (m *_NLMRouterAvailableToNetwork) GetParent() NLM {
	return m._NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMRouterAvailableToNetwork) GetDestinationNetworkAddresses() []uint16 {
	return m.DestinationNetworkAddresses
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMRouterAvailableToNetwork factory function for _NLMRouterAvailableToNetwork
func NewNLMRouterAvailableToNetwork(destinationNetworkAddresses []uint16, apduLength uint16) *_NLMRouterAvailableToNetwork {
	_result := &_NLMRouterAvailableToNetwork{
		DestinationNetworkAddresses: destinationNetworkAddresses,
		_NLM:                        NewNLM(apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMRouterAvailableToNetwork(structType any) NLMRouterAvailableToNetwork {
	if casted, ok := structType.(NLMRouterAvailableToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*NLMRouterAvailableToNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_NLMRouterAvailableToNetwork) GetTypeName() string {
	return "NLMRouterAvailableToNetwork"
}

func (m *_NLMRouterAvailableToNetwork) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.DestinationNetworkAddresses) > 0 {
		lengthInBits += 16 * uint16(len(m.DestinationNetworkAddresses))
	}

	return lengthInBits
}

func (m *_NLMRouterAvailableToNetwork) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMRouterAvailableToNetworkParse(ctx context.Context, theBytes []byte, apduLength uint16) (NLMRouterAvailableToNetwork, error) {
	return NLMRouterAvailableToNetworkParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func NLMRouterAvailableToNetworkParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (NLMRouterAvailableToNetwork, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NLMRouterAvailableToNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMRouterAvailableToNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (destinationNetworkAddresses)
	if pullErr := readBuffer.PullContext("destinationNetworkAddresses", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for destinationNetworkAddresses")
	}
	// Length array
	var destinationNetworkAddresses []uint16
	{
		_destinationNetworkAddressesLength := uint16(apduLength) - uint16(uint16(1))
		_destinationNetworkAddressesEndPos := positionAware.GetPos() + uint16(_destinationNetworkAddressesLength)
		for positionAware.GetPos() < _destinationNetworkAddressesEndPos {
			_item, _err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint16("", 16)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'destinationNetworkAddresses' field of NLMRouterAvailableToNetwork")
			}
			destinationNetworkAddresses = append(destinationNetworkAddresses, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("destinationNetworkAddresses", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for destinationNetworkAddresses")
	}

	if closeErr := readBuffer.CloseContext("NLMRouterAvailableToNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMRouterAvailableToNetwork")
	}

	// Create a partially initialized instance
	_child := &_NLMRouterAvailableToNetwork{
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
		DestinationNetworkAddresses: destinationNetworkAddresses,
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMRouterAvailableToNetwork) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMRouterAvailableToNetwork) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMRouterAvailableToNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMRouterAvailableToNetwork")
		}

		// Array Field (destinationNetworkAddresses)
		if pushErr := writeBuffer.PushContext("destinationNetworkAddresses", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for destinationNetworkAddresses")
		}
		for _curItem, _element := range m.GetDestinationNetworkAddresses() {
			_ = _curItem
			_elementErr := /*TODO: migrate me*/ writeBuffer.WriteUint16("", 16, uint16(_element))
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'destinationNetworkAddresses' field")
			}
		}
		if popErr := writeBuffer.PopContext("destinationNetworkAddresses", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for destinationNetworkAddresses")
		}

		if popErr := writeBuffer.PopContext("NLMRouterAvailableToNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMRouterAvailableToNetwork")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMRouterAvailableToNetwork) isNLMRouterAvailableToNetwork() bool {
	return true
}

func (m *_NLMRouterAvailableToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
