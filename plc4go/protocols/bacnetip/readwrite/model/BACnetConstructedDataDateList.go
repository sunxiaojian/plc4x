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

// BACnetConstructedDataDateList is the corresponding interface of BACnetConstructedDataDateList
type BACnetConstructedDataDateList interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDateList returns DateList (property field)
	GetDateList() []BACnetCalendarEntry
}

// BACnetConstructedDataDateListExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDateList.
// This is useful for switch cases.
type BACnetConstructedDataDateListExactly interface {
	BACnetConstructedDataDateList
	isBACnetConstructedDataDateList() bool
}

// _BACnetConstructedDataDateList is the data-structure of this message
type _BACnetConstructedDataDateList struct {
	*_BACnetConstructedData
	DateList []BACnetCalendarEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDateList) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDateList) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DATE_LIST
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDateList) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDateList) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDateList) GetDateList() []BACnetCalendarEntry {
	return m.DateList
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDateList factory function for _BACnetConstructedDataDateList
func NewBACnetConstructedDataDateList(dateList []BACnetCalendarEntry, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDateList {
	_result := &_BACnetConstructedDataDateList{
		DateList:               dateList,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDateList(structType interface{}) BACnetConstructedDataDateList {
	if casted, ok := structType.(BACnetConstructedDataDateList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDateList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDateList) GetTypeName() string {
	return "BACnetConstructedDataDateList"
}

func (m *_BACnetConstructedDataDateList) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDateList) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.DateList) > 0 {
		for _, element := range m.DateList {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataDateList) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDateListParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDateList, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDateList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDateList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (dateList)
	if pullErr := readBuffer.PullContext("dateList", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateList")
	}
	// Terminated array
	var dateList []BACnetCalendarEntry
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetCalendarEntryParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'dateList' field of BACnetConstructedDataDateList")
			}
			dateList = append(dateList, _item.(BACnetCalendarEntry))

		}
	}
	if closeErr := readBuffer.CloseContext("dateList", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateList")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDateList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDateList")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDateList{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DateList: dateList,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDateList) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDateList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDateList")
		}

		// Array Field (dateList)
		if pushErr := writeBuffer.PushContext("dateList", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dateList")
		}
		for _, _element := range m.GetDateList() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'dateList' field")
			}
		}
		if popErr := writeBuffer.PopContext("dateList", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dateList")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDateList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDateList")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDateList) isBACnetConstructedDataDateList() bool {
	return true
}

func (m *_BACnetConstructedDataDateList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
