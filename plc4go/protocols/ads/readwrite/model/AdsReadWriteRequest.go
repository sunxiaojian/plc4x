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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// AdsReadWriteRequest is the corresponding interface of AdsReadWriteRequest
type AdsReadWriteRequest interface {
	utils.LengthAware
	utils.Serializable
	AdsData
	// GetIndexGroup returns IndexGroup (property field)
	GetIndexGroup() uint32
	// GetIndexOffset returns IndexOffset (property field)
	GetIndexOffset() uint32
	// GetReadLength returns ReadLength (property field)
	GetReadLength() uint32
	// GetItems returns Items (property field)
	GetItems() []AdsMultiRequestItem
	// GetData returns Data (property field)
	GetData() []byte
}

// AdsReadWriteRequestExactly can be used when we want exactly this type and not a type which fulfills AdsReadWriteRequest.
// This is useful for switch cases.
type AdsReadWriteRequestExactly interface {
	AdsReadWriteRequest
	isAdsReadWriteRequest() bool
}

// _AdsReadWriteRequest is the data-structure of this message
type _AdsReadWriteRequest struct {
	*_AdsData
	IndexGroup  uint32
	IndexOffset uint32
	ReadLength  uint32
	Items       []AdsMultiRequestItem
	Data        []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsReadWriteRequest) GetCommandId() CommandId {
	return CommandId_ADS_READ_WRITE
}

func (m *_AdsReadWriteRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsReadWriteRequest) InitializeParent(parent AdsData) {}

func (m *_AdsReadWriteRequest) GetParent() AdsData {
	return m._AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsReadWriteRequest) GetIndexGroup() uint32 {
	return m.IndexGroup
}

func (m *_AdsReadWriteRequest) GetIndexOffset() uint32 {
	return m.IndexOffset
}

func (m *_AdsReadWriteRequest) GetReadLength() uint32 {
	return m.ReadLength
}

func (m *_AdsReadWriteRequest) GetItems() []AdsMultiRequestItem {
	return m.Items
}

func (m *_AdsReadWriteRequest) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsReadWriteRequest factory function for _AdsReadWriteRequest
func NewAdsReadWriteRequest(indexGroup uint32, indexOffset uint32, readLength uint32, items []AdsMultiRequestItem, data []byte) *_AdsReadWriteRequest {
	_result := &_AdsReadWriteRequest{
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		ReadLength:  readLength,
		Items:       items,
		Data:        data,
		_AdsData:    NewAdsData(),
	}
	_result._AdsData._AdsDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsReadWriteRequest(structType interface{}) AdsReadWriteRequest {
	if casted, ok := structType.(AdsReadWriteRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsReadWriteRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsReadWriteRequest) GetTypeName() string {
	return "AdsReadWriteRequest"
}

func (m *_AdsReadWriteRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsReadWriteRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (indexGroup)
	lengthInBits += 32

	// Simple field (indexOffset)
	lengthInBits += 32

	// Simple field (readLength)
	lengthInBits += 32

	// Implicit Field (writeLength)
	lengthInBits += 32

	// Array field
	if len(m.Items) > 0 {
		for i, element := range m.Items {
			last := i == len(m.Items)-1
			lengthInBits += element.(interface{ GetLengthInBitsConditional(bool) uint16 }).GetLengthInBitsConditional(last)
		}
	}

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_AdsReadWriteRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsReadWriteRequestParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (AdsReadWriteRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsReadWriteRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadWriteRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (indexGroup)
	_indexGroup, _indexGroupErr := readBuffer.ReadUint32("indexGroup", 32)
	if _indexGroupErr != nil {
		return nil, errors.Wrap(_indexGroupErr, "Error parsing 'indexGroup' field of AdsReadWriteRequest")
	}
	indexGroup := _indexGroup

	// Simple Field (indexOffset)
	_indexOffset, _indexOffsetErr := readBuffer.ReadUint32("indexOffset", 32)
	if _indexOffsetErr != nil {
		return nil, errors.Wrap(_indexOffsetErr, "Error parsing 'indexOffset' field of AdsReadWriteRequest")
	}
	indexOffset := _indexOffset

	// Simple Field (readLength)
	_readLength, _readLengthErr := readBuffer.ReadUint32("readLength", 32)
	if _readLengthErr != nil {
		return nil, errors.Wrap(_readLengthErr, "Error parsing 'readLength' field of AdsReadWriteRequest")
	}
	readLength := _readLength

	// Implicit Field (writeLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	writeLength, _writeLengthErr := readBuffer.ReadUint32("writeLength", 32)
	_ = writeLength
	if _writeLengthErr != nil {
		return nil, errors.Wrap(_writeLengthErr, "Error parsing 'writeLength' field of AdsReadWriteRequest")
	}

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for items")
	}
	// Count array
	items := make([]AdsMultiRequestItem, utils.InlineIf((bool(bool((bool((indexGroup) == (61568)))) || bool((bool((indexGroup) == (61569))))) || bool((bool((indexGroup) == (61570))))), func() interface{} { return uint16(indexOffset) }, func() interface{} { return uint16(uint16(0)) }).(uint16))
	// This happens when the size is set conditional to 0
	if len(items) == 0 {
		items = nil
	}
	{
		for curItem := uint16(0); curItem < uint16(utils.InlineIf((bool(bool((bool((indexGroup) == (61568)))) || bool((bool((indexGroup) == (61569))))) || bool((bool((indexGroup) == (61570))))), func() interface{} { return uint16(indexOffset) }, func() interface{} { return uint16(uint16(0)) }).(uint16)); curItem++ {
			_item, _err := AdsMultiRequestItemParse(readBuffer, indexGroup)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field of AdsReadWriteRequest")
			}
			items[curItem] = _item.(AdsMultiRequestItem)
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for items")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(uint16(writeLength) - uint16((uint16(uint16(len(items))) * uint16(uint16(12)))))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of AdsReadWriteRequest")
	}

	if closeErr := readBuffer.CloseContext("AdsReadWriteRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadWriteRequest")
	}

	// Create a partially initialized instance
	_child := &_AdsReadWriteRequest{
		_AdsData:    &_AdsData{},
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		ReadLength:  readLength,
		Items:       items,
		Data:        data,
	}
	_child._AdsData._AdsDataChildRequirements = _child
	return _child, nil
}

func (m *_AdsReadWriteRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadWriteRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadWriteRequest")
		}

		// Simple Field (indexGroup)
		indexGroup := uint32(m.GetIndexGroup())
		_indexGroupErr := writeBuffer.WriteUint32("indexGroup", 32, (indexGroup))
		if _indexGroupErr != nil {
			return errors.Wrap(_indexGroupErr, "Error serializing 'indexGroup' field")
		}

		// Simple Field (indexOffset)
		indexOffset := uint32(m.GetIndexOffset())
		_indexOffsetErr := writeBuffer.WriteUint32("indexOffset", 32, (indexOffset))
		if _indexOffsetErr != nil {
			return errors.Wrap(_indexOffsetErr, "Error serializing 'indexOffset' field")
		}

		// Simple Field (readLength)
		readLength := uint32(m.GetReadLength())
		_readLengthErr := writeBuffer.WriteUint32("readLength", 32, (readLength))
		if _readLengthErr != nil {
			return errors.Wrap(_readLengthErr, "Error serializing 'readLength' field")
		}

		// Implicit Field (writeLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		writeLength := uint32(uint32((uint32(uint32(len(m.GetItems()))) * uint32((utils.InlineIf((bool((m.GetIndexGroup()) == (61570))), func() interface{} { return uint32(uint32(16)) }, func() interface{} { return uint32(uint32(12)) }).(uint32))))) + uint32(uint32(len(m.GetData()))))
		_writeLengthErr := writeBuffer.WriteUint32("writeLength", 32, (writeLength))
		if _writeLengthErr != nil {
			return errors.Wrap(_writeLengthErr, "Error serializing 'writeLength' field")
		}

		// Array Field (items)
		if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for items")
		}
		for _, _element := range m.GetItems() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'items' field")
			}
		}
		if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for items")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("AdsReadWriteRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadWriteRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AdsReadWriteRequest) isAdsReadWriteRequest() bool {
	return true
}

func (m *_AdsReadWriteRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
