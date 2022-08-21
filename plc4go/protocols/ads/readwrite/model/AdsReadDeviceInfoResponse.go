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

// AdsReadDeviceInfoResponse is the corresponding interface of AdsReadDeviceInfoResponse
type AdsReadDeviceInfoResponse interface {
	utils.LengthAware
	utils.Serializable
	AdsData
	// GetResult returns Result (property field)
	GetResult() ReturnCode
	// GetMajorVersion returns MajorVersion (property field)
	GetMajorVersion() uint8
	// GetMinorVersion returns MinorVersion (property field)
	GetMinorVersion() uint8
	// GetVersion returns Version (property field)
	GetVersion() uint16
	// GetDevice returns Device (property field)
	GetDevice() []byte
}

// AdsReadDeviceInfoResponseExactly can be used when we want exactly this type and not a type which fulfills AdsReadDeviceInfoResponse.
// This is useful for switch cases.
type AdsReadDeviceInfoResponseExactly interface {
	AdsReadDeviceInfoResponse
	isAdsReadDeviceInfoResponse() bool
}

// _AdsReadDeviceInfoResponse is the data-structure of this message
type _AdsReadDeviceInfoResponse struct {
	*_AdsData
	Result       ReturnCode
	MajorVersion uint8
	MinorVersion uint8
	Version      uint16
	Device       []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsReadDeviceInfoResponse) GetCommandId() CommandId {
	return CommandId_ADS_READ_DEVICE_INFO
}

func (m *_AdsReadDeviceInfoResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsReadDeviceInfoResponse) InitializeParent(parent AdsData) {}

func (m *_AdsReadDeviceInfoResponse) GetParent() AdsData {
	return m._AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsReadDeviceInfoResponse) GetResult() ReturnCode {
	return m.Result
}

func (m *_AdsReadDeviceInfoResponse) GetMajorVersion() uint8 {
	return m.MajorVersion
}

func (m *_AdsReadDeviceInfoResponse) GetMinorVersion() uint8 {
	return m.MinorVersion
}

func (m *_AdsReadDeviceInfoResponse) GetVersion() uint16 {
	return m.Version
}

func (m *_AdsReadDeviceInfoResponse) GetDevice() []byte {
	return m.Device
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsReadDeviceInfoResponse factory function for _AdsReadDeviceInfoResponse
func NewAdsReadDeviceInfoResponse(result ReturnCode, majorVersion uint8, minorVersion uint8, version uint16, device []byte) *_AdsReadDeviceInfoResponse {
	_result := &_AdsReadDeviceInfoResponse{
		Result:       result,
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		Version:      version,
		Device:       device,
		_AdsData:     NewAdsData(),
	}
	_result._AdsData._AdsDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsReadDeviceInfoResponse(structType interface{}) AdsReadDeviceInfoResponse {
	if casted, ok := structType.(AdsReadDeviceInfoResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsReadDeviceInfoResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsReadDeviceInfoResponse) GetTypeName() string {
	return "AdsReadDeviceInfoResponse"
}

func (m *_AdsReadDeviceInfoResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsReadDeviceInfoResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (result)
	lengthInBits += 32

	// Simple field (majorVersion)
	lengthInBits += 8

	// Simple field (minorVersion)
	lengthInBits += 8

	// Simple field (version)
	lengthInBits += 16

	// Array field
	if len(m.Device) > 0 {
		lengthInBits += 8 * uint16(len(m.Device))
	}

	return lengthInBits
}

func (m *_AdsReadDeviceInfoResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsReadDeviceInfoResponseParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (AdsReadDeviceInfoResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsReadDeviceInfoResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadDeviceInfoResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (result)
	if pullErr := readBuffer.PullContext("result"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for result")
	}
	_result, _resultErr := ReturnCodeParse(readBuffer)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field of AdsReadDeviceInfoResponse")
	}
	result := _result
	if closeErr := readBuffer.CloseContext("result"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for result")
	}

	// Simple Field (majorVersion)
	_majorVersion, _majorVersionErr := readBuffer.ReadUint8("majorVersion", 8)
	if _majorVersionErr != nil {
		return nil, errors.Wrap(_majorVersionErr, "Error parsing 'majorVersion' field of AdsReadDeviceInfoResponse")
	}
	majorVersion := _majorVersion

	// Simple Field (minorVersion)
	_minorVersion, _minorVersionErr := readBuffer.ReadUint8("minorVersion", 8)
	if _minorVersionErr != nil {
		return nil, errors.Wrap(_minorVersionErr, "Error parsing 'minorVersion' field of AdsReadDeviceInfoResponse")
	}
	minorVersion := _minorVersion

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadUint16("version", 16)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field of AdsReadDeviceInfoResponse")
	}
	version := _version
	// Byte Array field (device)
	numberOfBytesdevice := int(uint16(16))
	device, _readArrayErr := readBuffer.ReadByteArray("device", numberOfBytesdevice)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'device' field of AdsReadDeviceInfoResponse")
	}

	if closeErr := readBuffer.CloseContext("AdsReadDeviceInfoResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadDeviceInfoResponse")
	}

	// Create a partially initialized instance
	_child := &_AdsReadDeviceInfoResponse{
		_AdsData:     &_AdsData{},
		Result:       result,
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		Version:      version,
		Device:       device,
	}
	_child._AdsData._AdsDataChildRequirements = _child
	return _child, nil
}

func (m *_AdsReadDeviceInfoResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadDeviceInfoResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadDeviceInfoResponse")
		}

		// Simple Field (result)
		if pushErr := writeBuffer.PushContext("result"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for result")
		}
		_resultErr := writeBuffer.WriteSerializable(m.GetResult())
		if popErr := writeBuffer.PopContext("result"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for result")
		}
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		// Simple Field (majorVersion)
		majorVersion := uint8(m.GetMajorVersion())
		_majorVersionErr := writeBuffer.WriteUint8("majorVersion", 8, (majorVersion))
		if _majorVersionErr != nil {
			return errors.Wrap(_majorVersionErr, "Error serializing 'majorVersion' field")
		}

		// Simple Field (minorVersion)
		minorVersion := uint8(m.GetMinorVersion())
		_minorVersionErr := writeBuffer.WriteUint8("minorVersion", 8, (minorVersion))
		if _minorVersionErr != nil {
			return errors.Wrap(_minorVersionErr, "Error serializing 'minorVersion' field")
		}

		// Simple Field (version)
		version := uint16(m.GetVersion())
		_versionErr := writeBuffer.WriteUint16("version", 16, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		// Array Field (device)
		// Byte Array field (device)
		if err := writeBuffer.WriteByteArray("device", m.GetDevice()); err != nil {
			return errors.Wrap(err, "Error serializing 'device' field")
		}

		if popErr := writeBuffer.PopContext("AdsReadDeviceInfoResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadDeviceInfoResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AdsReadDeviceInfoResponse) isAdsReadDeviceInfoResponse() bool {
	return true
}

func (m *_AdsReadDeviceInfoResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
