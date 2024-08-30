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

// Constant values.
const RequestTermination_CR byte = 0x0D

// RequestTermination is the corresponding interface of RequestTermination
type RequestTermination interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// RequestTerminationExactly can be used when we want exactly this type and not a type which fulfills RequestTermination.
// This is useful for switch cases.
type RequestTerminationExactly interface {
	RequestTermination
	isRequestTermination() bool
}

// _RequestTermination is the data-structure of this message
type _RequestTermination struct {
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_RequestTermination) GetCr() byte {
	return RequestTermination_CR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRequestTermination factory function for _RequestTermination
func NewRequestTermination() *_RequestTermination {
	return &_RequestTermination{}
}

// Deprecated: use the interface for direct cast
func CastRequestTermination(structType any) RequestTermination {
	if casted, ok := structType.(RequestTermination); ok {
		return casted
	}
	if casted, ok := structType.(*RequestTermination); ok {
		return *casted
	}
	return nil
}

func (m *_RequestTermination) GetTypeName() string {
	return "RequestTermination"
}

func (m *_RequestTermination) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (cr)
	lengthInBits += 8

	return lengthInBits
}

func (m *_RequestTermination) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RequestTerminationParse(ctx context.Context, theBytes []byte) (RequestTermination, error) {
	return RequestTerminationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func RequestTerminationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (RequestTermination, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("RequestTermination"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestTermination")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (cr)
	cr, _crErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadByte("cr")
	if _crErr != nil {
		return nil, errors.Wrap(_crErr, "Error parsing 'cr' field of RequestTermination")
	}
	if cr != RequestTermination_CR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", RequestTermination_CR) + " but got " + fmt.Sprintf("%d", cr))
	}

	if closeErr := readBuffer.CloseContext("RequestTermination"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestTermination")
	}

	// Create the instance
	return &_RequestTermination{}, nil
}

func (m *_RequestTermination) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RequestTermination) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("RequestTermination"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for RequestTermination")
	}

	// Const Field (cr)
	_crErr := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteByte("cr", 0x0D)
	if _crErr != nil {
		return errors.Wrap(_crErr, "Error serializing 'cr' field")
	}

	if popErr := writeBuffer.PopContext("RequestTermination"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for RequestTermination")
	}
	return nil
}

func (m *_RequestTermination) isRequestTermination() bool {
	return true
}

func (m *_RequestTermination) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
