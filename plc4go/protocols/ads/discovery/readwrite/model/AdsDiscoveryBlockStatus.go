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
const AdsDiscoveryBlockStatus_STATUSLENGTH uint16 = 0x0004

// AdsDiscoveryBlockStatus is the corresponding interface of AdsDiscoveryBlockStatus
type AdsDiscoveryBlockStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AdsDiscoveryBlock
	// GetStatus returns Status (property field)
	GetStatus() Status
}

// AdsDiscoveryBlockStatusExactly can be used when we want exactly this type and not a type which fulfills AdsDiscoveryBlockStatus.
// This is useful for switch cases.
type AdsDiscoveryBlockStatusExactly interface {
	AdsDiscoveryBlockStatus
	isAdsDiscoveryBlockStatus() bool
}

// _AdsDiscoveryBlockStatus is the data-structure of this message
type _AdsDiscoveryBlockStatus struct {
	*_AdsDiscoveryBlock
	Status Status
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsDiscoveryBlockStatus) GetBlockType() AdsDiscoveryBlockType {
	return AdsDiscoveryBlockType_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsDiscoveryBlockStatus) InitializeParent(parent AdsDiscoveryBlock) {}

func (m *_AdsDiscoveryBlockStatus) GetParent() AdsDiscoveryBlock {
	return m._AdsDiscoveryBlock
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsDiscoveryBlockStatus) GetStatus() Status {
	return m.Status
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_AdsDiscoveryBlockStatus) GetStatusLength() uint16 {
	return AdsDiscoveryBlockStatus_STATUSLENGTH
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsDiscoveryBlockStatus factory function for _AdsDiscoveryBlockStatus
func NewAdsDiscoveryBlockStatus(status Status) *_AdsDiscoveryBlockStatus {
	_result := &_AdsDiscoveryBlockStatus{
		Status:             status,
		_AdsDiscoveryBlock: NewAdsDiscoveryBlock(),
	}
	_result._AdsDiscoveryBlock._AdsDiscoveryBlockChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsDiscoveryBlockStatus(structType any) AdsDiscoveryBlockStatus {
	if casted, ok := structType.(AdsDiscoveryBlockStatus); ok {
		return casted
	}
	if casted, ok := structType.(*AdsDiscoveryBlockStatus); ok {
		return *casted
	}
	return nil
}

func (m *_AdsDiscoveryBlockStatus) GetTypeName() string {
	return "AdsDiscoveryBlockStatus"
}

func (m *_AdsDiscoveryBlockStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Const Field (statusLength)
	lengthInBits += 16

	// Simple field (status)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsDiscoveryBlockStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AdsDiscoveryBlockStatusParse(ctx context.Context, theBytes []byte) (AdsDiscoveryBlockStatus, error) {
	return AdsDiscoveryBlockStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AdsDiscoveryBlockStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AdsDiscoveryBlockStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AdsDiscoveryBlockStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsDiscoveryBlockStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (statusLength)
	statusLength, _statusLengthErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint16("statusLength", 16)
	if _statusLengthErr != nil {
		return nil, errors.Wrap(_statusLengthErr, "Error parsing 'statusLength' field of AdsDiscoveryBlockStatus")
	}
	if statusLength != AdsDiscoveryBlockStatus_STATUSLENGTH {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", AdsDiscoveryBlockStatus_STATUSLENGTH) + " but got " + fmt.Sprintf("%d", statusLength))
	}

	// Simple Field (status)
	if pullErr := readBuffer.PullContext("status"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for status")
	}
	_status, _statusErr := StatusParseWithBuffer(ctx, readBuffer)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field of AdsDiscoveryBlockStatus")
	}
	status := _status
	if closeErr := readBuffer.CloseContext("status"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for status")
	}

	if closeErr := readBuffer.CloseContext("AdsDiscoveryBlockStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsDiscoveryBlockStatus")
	}

	// Create a partially initialized instance
	_child := &_AdsDiscoveryBlockStatus{
		_AdsDiscoveryBlock: &_AdsDiscoveryBlock{},
		Status:             status,
	}
	_child._AdsDiscoveryBlock._AdsDiscoveryBlockChildRequirements = _child
	return _child, nil
}

func (m *_AdsDiscoveryBlockStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsDiscoveryBlockStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsDiscoveryBlockStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsDiscoveryBlockStatus")
		}

		// Const Field (statusLength)
		_statusLengthErr := /*TODO: migrate me*/ /*TODO: migrate me*/ writeBuffer.WriteUint16("statusLength", 16, uint16(0x0004))
		if _statusLengthErr != nil {
			return errors.Wrap(_statusLengthErr, "Error serializing 'statusLength' field")
		}

		// Simple Field (status)
		if pushErr := writeBuffer.PushContext("status"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for status")
		}
		_statusErr := writeBuffer.WriteSerializable(ctx, m.GetStatus())
		if popErr := writeBuffer.PopContext("status"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for status")
		}
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		if popErr := writeBuffer.PopContext("AdsDiscoveryBlockStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsDiscoveryBlockStatus")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsDiscoveryBlockStatus) isAdsDiscoveryBlockStatus() bool {
	return true
}

func (m *_AdsDiscoveryBlockStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
