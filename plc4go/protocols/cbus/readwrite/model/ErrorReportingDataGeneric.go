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

// ErrorReportingDataGeneric is the corresponding interface of ErrorReportingDataGeneric
type ErrorReportingDataGeneric interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ErrorReportingData
	// GetSystemCategory returns SystemCategory (property field)
	GetSystemCategory() ErrorReportingSystemCategory
	// GetMostRecent returns MostRecent (property field)
	GetMostRecent() bool
	// GetAcknowledge returns Acknowledge (property field)
	GetAcknowledge() bool
	// GetMostSevere returns MostSevere (property field)
	GetMostSevere() bool
	// GetSeverity returns Severity (property field)
	GetSeverity() ErrorReportingSeverity
	// GetDeviceId returns DeviceId (property field)
	GetDeviceId() uint8
	// GetErrorData1 returns ErrorData1 (property field)
	GetErrorData1() uint8
	// GetErrorData2 returns ErrorData2 (property field)
	GetErrorData2() uint8
	// GetIsMostSevereError returns IsMostSevereError (virtual field)
	GetIsMostSevereError() bool
	// GetIsMostRecentError returns IsMostRecentError (virtual field)
	GetIsMostRecentError() bool
	// GetIsMostRecentAndMostSevere returns IsMostRecentAndMostSevere (virtual field)
	GetIsMostRecentAndMostSevere() bool
}

// ErrorReportingDataGenericExactly can be used when we want exactly this type and not a type which fulfills ErrorReportingDataGeneric.
// This is useful for switch cases.
type ErrorReportingDataGenericExactly interface {
	ErrorReportingDataGeneric
	isErrorReportingDataGeneric() bool
}

// _ErrorReportingDataGeneric is the data-structure of this message
type _ErrorReportingDataGeneric struct {
	*_ErrorReportingData
	SystemCategory ErrorReportingSystemCategory
	MostRecent     bool
	Acknowledge    bool
	MostSevere     bool
	Severity       ErrorReportingSeverity
	DeviceId       uint8
	ErrorData1     uint8
	ErrorData2     uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ErrorReportingDataGeneric) InitializeParent(parent ErrorReportingData, commandTypeContainer ErrorReportingCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_ErrorReportingDataGeneric) GetParent() ErrorReportingData {
	return m._ErrorReportingData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorReportingDataGeneric) GetSystemCategory() ErrorReportingSystemCategory {
	return m.SystemCategory
}

func (m *_ErrorReportingDataGeneric) GetMostRecent() bool {
	return m.MostRecent
}

func (m *_ErrorReportingDataGeneric) GetAcknowledge() bool {
	return m.Acknowledge
}

func (m *_ErrorReportingDataGeneric) GetMostSevere() bool {
	return m.MostSevere
}

func (m *_ErrorReportingDataGeneric) GetSeverity() ErrorReportingSeverity {
	return m.Severity
}

func (m *_ErrorReportingDataGeneric) GetDeviceId() uint8 {
	return m.DeviceId
}

func (m *_ErrorReportingDataGeneric) GetErrorData1() uint8 {
	return m.ErrorData1
}

func (m *_ErrorReportingDataGeneric) GetErrorData2() uint8 {
	return m.ErrorData2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_ErrorReportingDataGeneric) GetIsMostSevereError() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetMostSevere())
}

func (m *_ErrorReportingDataGeneric) GetIsMostRecentError() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetMostRecent())
}

func (m *_ErrorReportingDataGeneric) GetIsMostRecentAndMostSevere() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(m.GetIsMostRecentError()) && bool(m.GetIsMostSevereError()))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewErrorReportingDataGeneric factory function for _ErrorReportingDataGeneric
func NewErrorReportingDataGeneric(systemCategory ErrorReportingSystemCategory, mostRecent bool, acknowledge bool, mostSevere bool, severity ErrorReportingSeverity, deviceId uint8, errorData1 uint8, errorData2 uint8, commandTypeContainer ErrorReportingCommandTypeContainer) *_ErrorReportingDataGeneric {
	_result := &_ErrorReportingDataGeneric{
		SystemCategory:      systemCategory,
		MostRecent:          mostRecent,
		Acknowledge:         acknowledge,
		MostSevere:          mostSevere,
		Severity:            severity,
		DeviceId:            deviceId,
		ErrorData1:          errorData1,
		ErrorData2:          errorData2,
		_ErrorReportingData: NewErrorReportingData(commandTypeContainer),
	}
	_result._ErrorReportingData._ErrorReportingDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastErrorReportingDataGeneric(structType any) ErrorReportingDataGeneric {
	if casted, ok := structType.(ErrorReportingDataGeneric); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorReportingDataGeneric); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorReportingDataGeneric) GetTypeName() string {
	return "ErrorReportingDataGeneric"
}

func (m *_ErrorReportingDataGeneric) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (systemCategory)
	lengthInBits += m.SystemCategory.GetLengthInBits(ctx)

	// Simple field (mostRecent)
	lengthInBits += 1

	// Simple field (acknowledge)
	lengthInBits += 1

	// Simple field (mostSevere)
	lengthInBits += 1

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Simple field (severity)
	lengthInBits += 3

	// Simple field (deviceId)
	lengthInBits += 8

	// Simple field (errorData1)
	lengthInBits += 8

	// Simple field (errorData2)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ErrorReportingDataGeneric) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorReportingDataGenericParse(ctx context.Context, theBytes []byte) (ErrorReportingDataGeneric, error) {
	return ErrorReportingDataGenericParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorReportingDataGenericParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorReportingDataGeneric, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ErrorReportingDataGeneric"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingDataGeneric")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (systemCategory)
	if pullErr := readBuffer.PullContext("systemCategory"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for systemCategory")
	}
	_systemCategory, _systemCategoryErr := ErrorReportingSystemCategoryParseWithBuffer(ctx, readBuffer)
	if _systemCategoryErr != nil {
		return nil, errors.Wrap(_systemCategoryErr, "Error parsing 'systemCategory' field of ErrorReportingDataGeneric")
	}
	systemCategory := _systemCategory.(ErrorReportingSystemCategory)
	if closeErr := readBuffer.CloseContext("systemCategory"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for systemCategory")
	}

	// Simple Field (mostRecent)
	_mostRecent, _mostRecentErr := /*TODO: migrate me*/ readBuffer.ReadBit("mostRecent")
	if _mostRecentErr != nil {
		return nil, errors.Wrap(_mostRecentErr, "Error parsing 'mostRecent' field of ErrorReportingDataGeneric")
	}
	mostRecent := _mostRecent

	// Simple Field (acknowledge)
	_acknowledge, _acknowledgeErr := /*TODO: migrate me*/ readBuffer.ReadBit("acknowledge")
	if _acknowledgeErr != nil {
		return nil, errors.Wrap(_acknowledgeErr, "Error parsing 'acknowledge' field of ErrorReportingDataGeneric")
	}
	acknowledge := _acknowledge

	// Simple Field (mostSevere)
	_mostSevere, _mostSevereErr := /*TODO: migrate me*/ readBuffer.ReadBit("mostSevere")
	if _mostSevereErr != nil {
		return nil, errors.Wrap(_mostSevereErr, "Error parsing 'mostSevere' field of ErrorReportingDataGeneric")
	}
	mostSevere := _mostSevere

	// Validation
	if !(bool(mostRecent) || bool(mostSevere)) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "Invalid Error condition"})
	}

	// Virtual field
	_isMostSevereError := mostSevere
	isMostSevereError := bool(_isMostSevereError)
	_ = isMostSevereError

	// Virtual field
	_isMostRecentError := mostRecent
	isMostRecentError := bool(_isMostRecentError)
	_ = isMostRecentError

	// Virtual field
	_isMostRecentAndMostSevere := bool(isMostRecentError) && bool(isMostSevereError)
	isMostRecentAndMostSevere := bool(_isMostRecentAndMostSevere)
	_ = isMostRecentAndMostSevere

	// Simple Field (severity)
	if pullErr := readBuffer.PullContext("severity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for severity")
	}
	_severity, _severityErr := ErrorReportingSeverityParseWithBuffer(ctx, readBuffer)
	if _severityErr != nil {
		return nil, errors.Wrap(_severityErr, "Error parsing 'severity' field of ErrorReportingDataGeneric")
	}
	severity := _severity
	if closeErr := readBuffer.CloseContext("severity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for severity")
	}

	// Simple Field (deviceId)
	_deviceId, _deviceIdErr := /*TODO: migrate me*/ readBuffer.ReadUint8("deviceId", 8)
	if _deviceIdErr != nil {
		return nil, errors.Wrap(_deviceIdErr, "Error parsing 'deviceId' field of ErrorReportingDataGeneric")
	}
	deviceId := _deviceId

	// Simple Field (errorData1)
	_errorData1, _errorData1Err := /*TODO: migrate me*/ readBuffer.ReadUint8("errorData1", 8)
	if _errorData1Err != nil {
		return nil, errors.Wrap(_errorData1Err, "Error parsing 'errorData1' field of ErrorReportingDataGeneric")
	}
	errorData1 := _errorData1

	// Simple Field (errorData2)
	_errorData2, _errorData2Err := /*TODO: migrate me*/ readBuffer.ReadUint8("errorData2", 8)
	if _errorData2Err != nil {
		return nil, errors.Wrap(_errorData2Err, "Error parsing 'errorData2' field of ErrorReportingDataGeneric")
	}
	errorData2 := _errorData2

	if closeErr := readBuffer.CloseContext("ErrorReportingDataGeneric"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingDataGeneric")
	}

	// Create a partially initialized instance
	_child := &_ErrorReportingDataGeneric{
		_ErrorReportingData: &_ErrorReportingData{},
		SystemCategory:      systemCategory,
		MostRecent:          mostRecent,
		Acknowledge:         acknowledge,
		MostSevere:          mostSevere,
		Severity:            severity,
		DeviceId:            deviceId,
		ErrorData1:          errorData1,
		ErrorData2:          errorData2,
	}
	_child._ErrorReportingData._ErrorReportingDataChildRequirements = _child
	return _child, nil
}

func (m *_ErrorReportingDataGeneric) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorReportingDataGeneric) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ErrorReportingDataGeneric"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ErrorReportingDataGeneric")
		}

		// Simple Field (systemCategory)
		if pushErr := writeBuffer.PushContext("systemCategory"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for systemCategory")
		}
		_systemCategoryErr := writeBuffer.WriteSerializable(ctx, m.GetSystemCategory())
		if popErr := writeBuffer.PopContext("systemCategory"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for systemCategory")
		}
		if _systemCategoryErr != nil {
			return errors.Wrap(_systemCategoryErr, "Error serializing 'systemCategory' field")
		}

		// Simple Field (mostRecent)
		mostRecent := bool(m.GetMostRecent())
		_mostRecentErr := /*TODO: migrate me*/ writeBuffer.WriteBit("mostRecent", (mostRecent))
		if _mostRecentErr != nil {
			return errors.Wrap(_mostRecentErr, "Error serializing 'mostRecent' field")
		}

		// Simple Field (acknowledge)
		acknowledge := bool(m.GetAcknowledge())
		_acknowledgeErr := /*TODO: migrate me*/ writeBuffer.WriteBit("acknowledge", (acknowledge))
		if _acknowledgeErr != nil {
			return errors.Wrap(_acknowledgeErr, "Error serializing 'acknowledge' field")
		}

		// Simple Field (mostSevere)
		mostSevere := bool(m.GetMostSevere())
		_mostSevereErr := /*TODO: migrate me*/ writeBuffer.WriteBit("mostSevere", (mostSevere))
		if _mostSevereErr != nil {
			return errors.Wrap(_mostSevereErr, "Error serializing 'mostSevere' field")
		}
		// Virtual field
		isMostSevereError := m.GetIsMostSevereError()
		_ = isMostSevereError
		if _isMostSevereErrorErr := writeBuffer.WriteVirtual(ctx, "isMostSevereError", m.GetIsMostSevereError()); _isMostSevereErrorErr != nil {
			return errors.Wrap(_isMostSevereErrorErr, "Error serializing 'isMostSevereError' field")
		}
		// Virtual field
		isMostRecentError := m.GetIsMostRecentError()
		_ = isMostRecentError
		if _isMostRecentErrorErr := writeBuffer.WriteVirtual(ctx, "isMostRecentError", m.GetIsMostRecentError()); _isMostRecentErrorErr != nil {
			return errors.Wrap(_isMostRecentErrorErr, "Error serializing 'isMostRecentError' field")
		}
		// Virtual field
		isMostRecentAndMostSevere := m.GetIsMostRecentAndMostSevere()
		_ = isMostRecentAndMostSevere
		if _isMostRecentAndMostSevereErr := writeBuffer.WriteVirtual(ctx, "isMostRecentAndMostSevere", m.GetIsMostRecentAndMostSevere()); _isMostRecentAndMostSevereErr != nil {
			return errors.Wrap(_isMostRecentAndMostSevereErr, "Error serializing 'isMostRecentAndMostSevere' field")
		}

		// Simple Field (severity)
		if pushErr := writeBuffer.PushContext("severity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for severity")
		}
		_severityErr := writeBuffer.WriteSerializable(ctx, m.GetSeverity())
		if popErr := writeBuffer.PopContext("severity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for severity")
		}
		if _severityErr != nil {
			return errors.Wrap(_severityErr, "Error serializing 'severity' field")
		}

		// Simple Field (deviceId)
		deviceId := uint8(m.GetDeviceId())
		_deviceIdErr := /*TODO: migrate me*/ writeBuffer.WriteUint8("deviceId", 8, uint8((deviceId)))
		if _deviceIdErr != nil {
			return errors.Wrap(_deviceIdErr, "Error serializing 'deviceId' field")
		}

		// Simple Field (errorData1)
		errorData1 := uint8(m.GetErrorData1())
		_errorData1Err := /*TODO: migrate me*/ writeBuffer.WriteUint8("errorData1", 8, uint8((errorData1)))
		if _errorData1Err != nil {
			return errors.Wrap(_errorData1Err, "Error serializing 'errorData1' field")
		}

		// Simple Field (errorData2)
		errorData2 := uint8(m.GetErrorData2())
		_errorData2Err := /*TODO: migrate me*/ writeBuffer.WriteUint8("errorData2", 8, uint8((errorData2)))
		if _errorData2Err != nil {
			return errors.Wrap(_errorData2Err, "Error serializing 'errorData2' field")
		}

		if popErr := writeBuffer.PopContext("ErrorReportingDataGeneric"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ErrorReportingDataGeneric")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ErrorReportingDataGeneric) isErrorReportingDataGeneric() bool {
	return true
}

func (m *_ErrorReportingDataGeneric) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
