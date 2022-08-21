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

// BACnetPropertyStates is the corresponding interface of BACnetPropertyStates
type BACnetPropertyStates interface {
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetPropertyStatesExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStates.
// This is useful for switch cases.
type BACnetPropertyStatesExactly interface {
	BACnetPropertyStates
	isBACnetPropertyStates() bool
}

// _BACnetPropertyStates is the data-structure of this message
type _BACnetPropertyStates struct {
	_BACnetPropertyStatesChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetPropertyStatesChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}

type BACnetPropertyStatesParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child BACnetPropertyStates, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetPropertyStatesChild interface {
	utils.Serializable
	InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetPropertyStates

	GetTypeName() string
	BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStates) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetPropertyStates) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStates factory function for _BACnetPropertyStates
func NewBACnetPropertyStates(peekedTagHeader BACnetTagHeader) *_BACnetPropertyStates {
	return &_BACnetPropertyStates{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStates(structType interface{}) BACnetPropertyStates {
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStates) GetTypeName() string {
	return "BACnetPropertyStates"
}

func (m *_BACnetPropertyStates) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetPropertyStates) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesParse(readBuffer utils.ReadBuffer) (BACnetPropertyStates, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStates"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStates")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetPropertyStatesChildSerializeRequirement interface {
		BACnetPropertyStates
		InitializeParent(BACnetPropertyStates, BACnetTagHeader)
		GetParent() BACnetPropertyStates
	}
	var _childTemp interface{}
	var _child BACnetPropertyStatesChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetPropertyStatesBoolean
		_childTemp, typeSwitchError = BACnetPropertyStatesBooleanParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(1): // BACnetPropertyStatesBinaryValue
		_childTemp, typeSwitchError = BACnetPropertyStatesBinaryValueParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(2): // BACnetPropertyStatesEventType
		_childTemp, typeSwitchError = BACnetPropertyStatesEventTypeParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(3): // BACnetPropertyStatesPolarity
		_childTemp, typeSwitchError = BACnetPropertyStatesPolarityParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(4): // BACnetPropertyStatesProgramChange
		_childTemp, typeSwitchError = BACnetPropertyStatesProgramChangeParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(5): // BACnetPropertyStatesProgramChange
		_childTemp, typeSwitchError = BACnetPropertyStatesProgramChangeParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(6): // BACnetPropertyStatesReasonForHalt
		_childTemp, typeSwitchError = BACnetPropertyStatesReasonForHaltParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(7): // BACnetPropertyStatesReliability
		_childTemp, typeSwitchError = BACnetPropertyStatesReliabilityParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(8): // BACnetPropertyStatesState
		_childTemp, typeSwitchError = BACnetPropertyStatesStateParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(9): // BACnetPropertyStatesSystemStatus
		_childTemp, typeSwitchError = BACnetPropertyStatesSystemStatusParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(10): // BACnetPropertyStatesUnits
		_childTemp, typeSwitchError = BACnetPropertyStatesUnitsParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(11): // BACnetPropertyStatesExtendedValue
		_childTemp, typeSwitchError = BACnetPropertyStatesExtendedValueParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(12): // BACnetPropertyStatesLifeSafetyMode
		_childTemp, typeSwitchError = BACnetPropertyStatesLifeSafetyModeParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(13): // BACnetPropertyStatesLifeSafetyState
		_childTemp, typeSwitchError = BACnetPropertyStatesLifeSafetyStateParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(14): // BACnetPropertyStatesRestartReason
		_childTemp, typeSwitchError = BACnetPropertyStatesRestartReasonParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(15): // BACnetPropertyStatesDoorAlarmState
		_childTemp, typeSwitchError = BACnetPropertyStatesDoorAlarmStateParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(16): // BACnetPropertyStatesAction
		_childTemp, typeSwitchError = BACnetPropertyStatesActionParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(17): // BACnetPropertyStatesDoorSecuredStatus
		_childTemp, typeSwitchError = BACnetPropertyStatesDoorSecuredStatusParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(18): // BACnetPropertyStatesDoorStatus
		_childTemp, typeSwitchError = BACnetPropertyStatesDoorStatusParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(19): // BACnetPropertyStatesDoorValue
		_childTemp, typeSwitchError = BACnetPropertyStatesDoorValueParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(20): // BACnetPropertyStatesFileAccessMethod
		_childTemp, typeSwitchError = BACnetPropertyStatesFileAccessMethodParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(21): // BACnetPropertyStatesLockStatus
		_childTemp, typeSwitchError = BACnetPropertyStatesLockStatusParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(22): // BACnetPropertyStatesLifeSafetyOperations
		_childTemp, typeSwitchError = BACnetPropertyStatesLifeSafetyOperationsParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(23): // BACnetPropertyStatesMaintenance
		_childTemp, typeSwitchError = BACnetPropertyStatesMaintenanceParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(24): // BACnetPropertyStatesNodeType
		_childTemp, typeSwitchError = BACnetPropertyStatesNodeTypeParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(25): // BACnetPropertyStatesNotifyType
		_childTemp, typeSwitchError = BACnetPropertyStatesNotifyTypeParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(26): // BACnetPropertyStatesSecurityLevel
		_childTemp, typeSwitchError = BACnetPropertyStatesSecurityLevelParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(27): // BACnetPropertyStatesShedState
		_childTemp, typeSwitchError = BACnetPropertyStatesShedStateParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(28): // BACnetPropertyStatesSilencedState
		_childTemp, typeSwitchError = BACnetPropertyStatesSilencedStateParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(30): // BACnetPropertyStatesAccessEvent
		_childTemp, typeSwitchError = BACnetPropertyStatesAccessEventParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(31): // BACnetPropertyStatesZoneOccupanyState
		_childTemp, typeSwitchError = BACnetPropertyStatesZoneOccupanyStateParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(32): // BACnetPropertyStatesAccessCredentialDisableReason
		_childTemp, typeSwitchError = BACnetPropertyStatesAccessCredentialDisableReasonParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(33): // BACnetPropertyStatesAccessCredentialDisable
		_childTemp, typeSwitchError = BACnetPropertyStatesAccessCredentialDisableParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(34): // BACnetPropertyStatesAuthenticationStatus
		_childTemp, typeSwitchError = BACnetPropertyStatesAuthenticationStatusParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(36): // BACnetPropertyStatesBackupState
		_childTemp, typeSwitchError = BACnetPropertyStatesBackupStateParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(37): // BACnetPropertyStatesWriteStatus
		_childTemp, typeSwitchError = BACnetPropertyStatesWriteStatusParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(38): // BACnetPropertyStatesLightningInProgress
		_childTemp, typeSwitchError = BACnetPropertyStatesLightningInProgressParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(39): // BACnetPropertyStatesLightningOperation
		_childTemp, typeSwitchError = BACnetPropertyStatesLightningOperationParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(40): // BACnetPropertyStatesLightningTransition
		_childTemp, typeSwitchError = BACnetPropertyStatesLightningTransitionParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(41): // BACnetPropertyStatesIntegerValue
		_childTemp, typeSwitchError = BACnetPropertyStatesIntegerValueParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(42): // BACnetPropertyStatesBinaryLightningValue
		_childTemp, typeSwitchError = BACnetPropertyStatesBinaryLightningValueParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(43): // BACnetPropertyStatesTimerState
		_childTemp, typeSwitchError = BACnetPropertyStatesTimerStateParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(44): // BACnetPropertyStatesTimerTransition
		_childTemp, typeSwitchError = BACnetPropertyStatesTimerTransitionParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(45): // BACnetPropertyStatesBacnetIpMode
		_childTemp, typeSwitchError = BACnetPropertyStatesBacnetIpModeParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(46): // BACnetPropertyStatesNetworkPortCommand
		_childTemp, typeSwitchError = BACnetPropertyStatesNetworkPortCommandParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(47): // BACnetPropertyStatesNetworkType
		_childTemp, typeSwitchError = BACnetPropertyStatesNetworkTypeParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(48): // BACnetPropertyStatesNetworkNumberQuality
		_childTemp, typeSwitchError = BACnetPropertyStatesNetworkNumberQualityParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(49): // BACnetPropertyStatesEscalatorOperationDirection
		_childTemp, typeSwitchError = BACnetPropertyStatesEscalatorOperationDirectionParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(50): // BACnetPropertyStatesEscalatorFault
		_childTemp, typeSwitchError = BACnetPropertyStatesEscalatorFaultParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(51): // BACnetPropertyStatesEscalatorMode
		_childTemp, typeSwitchError = BACnetPropertyStatesEscalatorModeParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(52): // BACnetPropertyStatesLiftCarDirection
		_childTemp, typeSwitchError = BACnetPropertyStatesLiftCarDirectionParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(53): // BACnetPropertyStatesLiftCarDoorCommand
		_childTemp, typeSwitchError = BACnetPropertyStatesLiftCarDoorCommandParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(54): // BACnetPropertyStatesLiftCarDriveStatus
		_childTemp, typeSwitchError = BACnetPropertyStatesLiftCarDriveStatusParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(55): // BACnetPropertyStatesLiftCarMode
		_childTemp, typeSwitchError = BACnetPropertyStatesLiftCarModeParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(56): // BACnetPropertyStatesLiftGroupMode
		_childTemp, typeSwitchError = BACnetPropertyStatesLiftGroupModeParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(57): // BACnetPropertyStatesLiftFault
		_childTemp, typeSwitchError = BACnetPropertyStatesLiftFaultParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(58): // BACnetPropertyStatesProtocolLevel
		_childTemp, typeSwitchError = BACnetPropertyStatesProtocolLevelParse(readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(63): // BACnetPropertyStatesExtendedValue
		_childTemp, typeSwitchError = BACnetPropertyStatesExtendedValueParse(readBuffer, peekedTagNumber)
	case true: // BACnetPropertyStateActionUnknown
		_childTemp, typeSwitchError = BACnetPropertyStateActionUnknownParse(readBuffer, peekedTagNumber)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetPropertyStates")
	}
	_child = _childTemp.(BACnetPropertyStatesChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetPropertyStates"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStates")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetPropertyStates) SerializeParent(writeBuffer utils.WriteBuffer, child BACnetPropertyStates, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetPropertyStates"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStates")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyStates"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyStates")
	}
	return nil
}

func (m *_BACnetPropertyStates) isBACnetPropertyStates() bool {
	return true
}

func (m *_BACnetPropertyStates) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
