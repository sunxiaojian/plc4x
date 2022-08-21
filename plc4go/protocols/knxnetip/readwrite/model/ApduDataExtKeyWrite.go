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

// ApduDataExtKeyWrite is the corresponding interface of ApduDataExtKeyWrite
type ApduDataExtKeyWrite interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtKeyWriteExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtKeyWrite.
// This is useful for switch cases.
type ApduDataExtKeyWriteExactly interface {
	ApduDataExtKeyWrite
	isApduDataExtKeyWrite() bool
}

// _ApduDataExtKeyWrite is the data-structure of this message
type _ApduDataExtKeyWrite struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtKeyWrite) GetExtApciType() uint8 {
	return 0x13
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtKeyWrite) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtKeyWrite) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtKeyWrite factory function for _ApduDataExtKeyWrite
func NewApduDataExtKeyWrite(length uint8) *_ApduDataExtKeyWrite {
	_result := &_ApduDataExtKeyWrite{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtKeyWrite(structType interface{}) ApduDataExtKeyWrite {
	if casted, ok := structType.(ApduDataExtKeyWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtKeyWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtKeyWrite) GetTypeName() string {
	return "ApduDataExtKeyWrite"
}

func (m *_ApduDataExtKeyWrite) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtKeyWrite) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataExtKeyWrite) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtKeyWriteParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtKeyWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtKeyWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtKeyWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtKeyWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtKeyWrite")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtKeyWrite{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtKeyWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtKeyWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtKeyWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtKeyWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtKeyWrite")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtKeyWrite) isApduDataExtKeyWrite() bool {
	return true
}

func (m *_ApduDataExtKeyWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
