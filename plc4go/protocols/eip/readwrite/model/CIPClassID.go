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

// CIPClassID is an enum
type CIPClassID uint16

type ICIPClassID interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	CIPClassID_Identity                     CIPClassID = 0x0001
	CIPClassID_MessageRouter                CIPClassID = 0x0002
	CIPClassID_Assembly                     CIPClassID = 0x0004
	CIPClassID_Connection                   CIPClassID = 0x0005
	CIPClassID_ConnectionManager            CIPClassID = 0x0006
	CIPClassID_Register                     CIPClassID = 0x0007
	CIPClassID_Parameter                    CIPClassID = 0x000F
	CIPClassID_ParameterGroup               CIPClassID = 0x0010
	CIPClassID_AcknowledgeHandler           CIPClassID = 0x002B
	CIPClassID_Selection                    CIPClassID = 0x002E
	CIPClassID_File                         CIPClassID = 0x0037
	CIPClassID_OriginatorConnectionList     CIPClassID = 0x0045
	CIPClassID_ConnectionConfiguration      CIPClassID = 0x00F3
	CIPClassID_Port                         CIPClassID = 0x00F4
	CIPClassID_DiscreteInputPoint           CIPClassID = 0x0008
	CIPClassID_DiscreteOutputPoint          CIPClassID = 0x0009
	CIPClassID_AnalogInputPoint             CIPClassID = 0x000A
	CIPClassID_AnalogOutputPoint            CIPClassID = 0x000B
	CIPClassID_PresenceSensing              CIPClassID = 0x000E
	CIPClassID_Group                        CIPClassID = 0x0012
	CIPClassID_DiscreteInputGroup           CIPClassID = 0x001D
	CIPClassID_DiscreteOutputGroup          CIPClassID = 0x001E
	CIPClassID_DiscreteGroup                CIPClassID = 0x001F
	CIPClassID_AnalogInputGroup             CIPClassID = 0x0020
	CIPClassID_AnalogOutputGroup            CIPClassID = 0x0021
	CIPClassID_AnalogGroup                  CIPClassID = 0x0022
	CIPClassID_PositionSensor               CIPClassID = 0x0023
	CIPClassID_PositionControllerSupervisor CIPClassID = 0x0024
	CIPClassID_PositionController           CIPClassID = 0x0025
	CIPClassID_BlockSequencer               CIPClassID = 0x0026
	CIPClassID_CommandBlock                 CIPClassID = 0x0027
	CIPClassID_MotorData                    CIPClassID = 0x0028
	CIPClassID_ControlSupervisor            CIPClassID = 0x0029
	CIPClassID_AcDcDrive                    CIPClassID = 0x002A
	CIPClassID_Overload                     CIPClassID = 0x002C
	CIPClassID_SDeviceSupervisor            CIPClassID = 0x0030
	CIPClassID_SAnalogSensor                CIPClassID = 0x0031
	CIPClassID_SAnalogActuator              CIPClassID = 0x0032
	CIPClassID_SSingleStageController       CIPClassID = 0x0033
	CIPClassID_SGasCalibration              CIPClassID = 0x0034
	CIPClassID_SPartialPressure             CIPClassID = 0x0038
	CIPClassID_SSensorCalibration           CIPClassID = 0x0040
	CIPClassID_EventLog                     CIPClassID = 0x0041
	CIPClassID_MotionDeviceAxis             CIPClassID = 0x0042
	CIPClassID_SafetyAnalogInputGroup       CIPClassID = 0x004A
	CIPClassID_BaseEnergy                   CIPClassID = 0x004E
	CIPClassID_ElectricalEnergy             CIPClassID = 0x004F
	CIPClassID_NonElectricalEnergy          CIPClassID = 0x0050
	CIPClassID_PowerManagementObject        CIPClassID = 0x0053
	CIPClassID_PowerCurtailmentObject       CIPClassID = 0x005C
	CIPClassID_DeviceNet                    CIPClassID = 0x0003
	CIPClassID_Modbus                       CIPClassID = 0x0044
	CIPClassID_ModbusSerialLink             CIPClassID = 0x0046
	CIPClassID_DeviceLevelRing              CIPClassID = 0x0047
	CIPClassID_QOS                          CIPClassID = 0x0048
	CIPClassID_Sercos3Link                  CIPClassID = 0x004C
	CIPClassID_BaseSwitch                   CIPClassID = 0x0051
	CIPClassID_Snmp                         CIPClassID = 0x0052
	CIPClassID_PowerManagement              CIPClassID = 0x0053
	CIPClassID_RstpBridge                   CIPClassID = 0x0054
	CIPClassID_RstpPort                     CIPClassID = 0x0055
	CIPClassID_ParallelRedundancyProtocol   CIPClassID = 0x0056
	CIPClassID_PrpNodesTable                CIPClassID = 0x0057
	CIPClassID_ControlNet                   CIPClassID = 0x00F0
	CIPClassID_ControlNetKeeper             CIPClassID = 0x00F1
	CIPClassID_ControlNetScheduling         CIPClassID = 0x00F2
	CIPClassID_TcpIpInterface               CIPClassID = 0x00F5
	CIPClassID_EthernetLink                 CIPClassID = 0x00F6
	CIPClassID_CompoNetLink                 CIPClassID = 0x00F7
	CIPClassID_CompoNetRepeater             CIPClassID = 0x00F8
)

var CIPClassIDValues []CIPClassID

func init() {
	_ = errors.New
	CIPClassIDValues = []CIPClassID{
		CIPClassID_Identity,
		CIPClassID_MessageRouter,
		CIPClassID_Assembly,
		CIPClassID_Connection,
		CIPClassID_ConnectionManager,
		CIPClassID_Register,
		CIPClassID_Parameter,
		CIPClassID_ParameterGroup,
		CIPClassID_AcknowledgeHandler,
		CIPClassID_Selection,
		CIPClassID_File,
		CIPClassID_OriginatorConnectionList,
		CIPClassID_ConnectionConfiguration,
		CIPClassID_Port,
		CIPClassID_DiscreteInputPoint,
		CIPClassID_DiscreteOutputPoint,
		CIPClassID_AnalogInputPoint,
		CIPClassID_AnalogOutputPoint,
		CIPClassID_PresenceSensing,
		CIPClassID_Group,
		CIPClassID_DiscreteInputGroup,
		CIPClassID_DiscreteOutputGroup,
		CIPClassID_DiscreteGroup,
		CIPClassID_AnalogInputGroup,
		CIPClassID_AnalogOutputGroup,
		CIPClassID_AnalogGroup,
		CIPClassID_PositionSensor,
		CIPClassID_PositionControllerSupervisor,
		CIPClassID_PositionController,
		CIPClassID_BlockSequencer,
		CIPClassID_CommandBlock,
		CIPClassID_MotorData,
		CIPClassID_ControlSupervisor,
		CIPClassID_AcDcDrive,
		CIPClassID_Overload,
		CIPClassID_SDeviceSupervisor,
		CIPClassID_SAnalogSensor,
		CIPClassID_SAnalogActuator,
		CIPClassID_SSingleStageController,
		CIPClassID_SGasCalibration,
		CIPClassID_SPartialPressure,
		CIPClassID_SSensorCalibration,
		CIPClassID_EventLog,
		CIPClassID_MotionDeviceAxis,
		CIPClassID_SafetyAnalogInputGroup,
		CIPClassID_BaseEnergy,
		CIPClassID_ElectricalEnergy,
		CIPClassID_NonElectricalEnergy,
		CIPClassID_PowerManagementObject,
		CIPClassID_PowerCurtailmentObject,
		CIPClassID_DeviceNet,
		CIPClassID_Modbus,
		CIPClassID_ModbusSerialLink,
		CIPClassID_DeviceLevelRing,
		CIPClassID_QOS,
		CIPClassID_Sercos3Link,
		CIPClassID_BaseSwitch,
		CIPClassID_Snmp,
		CIPClassID_PowerManagement,
		CIPClassID_RstpBridge,
		CIPClassID_RstpPort,
		CIPClassID_ParallelRedundancyProtocol,
		CIPClassID_PrpNodesTable,
		CIPClassID_ControlNet,
		CIPClassID_ControlNetKeeper,
		CIPClassID_ControlNetScheduling,
		CIPClassID_TcpIpInterface,
		CIPClassID_EthernetLink,
		CIPClassID_CompoNetLink,
		CIPClassID_CompoNetRepeater,
	}
}

func CIPClassIDByValue(value uint16) (enum CIPClassID, ok bool) {
	switch value {
	case 0x0001:
		return CIPClassID_Identity, true
	case 0x0002:
		return CIPClassID_MessageRouter, true
	case 0x0003:
		return CIPClassID_DeviceNet, true
	case 0x0004:
		return CIPClassID_Assembly, true
	case 0x0005:
		return CIPClassID_Connection, true
	case 0x0006:
		return CIPClassID_ConnectionManager, true
	case 0x0007:
		return CIPClassID_Register, true
	case 0x0008:
		return CIPClassID_DiscreteInputPoint, true
	case 0x0009:
		return CIPClassID_DiscreteOutputPoint, true
	case 0x000A:
		return CIPClassID_AnalogInputPoint, true
	case 0x000B:
		return CIPClassID_AnalogOutputPoint, true
	case 0x000E:
		return CIPClassID_PresenceSensing, true
	case 0x000F:
		return CIPClassID_Parameter, true
	case 0x0010:
		return CIPClassID_ParameterGroup, true
	case 0x0012:
		return CIPClassID_Group, true
	case 0x001D:
		return CIPClassID_DiscreteInputGroup, true
	case 0x001E:
		return CIPClassID_DiscreteOutputGroup, true
	case 0x001F:
		return CIPClassID_DiscreteGroup, true
	case 0x0020:
		return CIPClassID_AnalogInputGroup, true
	case 0x0021:
		return CIPClassID_AnalogOutputGroup, true
	case 0x0022:
		return CIPClassID_AnalogGroup, true
	case 0x0023:
		return CIPClassID_PositionSensor, true
	case 0x0024:
		return CIPClassID_PositionControllerSupervisor, true
	case 0x0025:
		return CIPClassID_PositionController, true
	case 0x0026:
		return CIPClassID_BlockSequencer, true
	case 0x0027:
		return CIPClassID_CommandBlock, true
	case 0x0028:
		return CIPClassID_MotorData, true
	case 0x0029:
		return CIPClassID_ControlSupervisor, true
	case 0x002A:
		return CIPClassID_AcDcDrive, true
	case 0x002B:
		return CIPClassID_AcknowledgeHandler, true
	case 0x002C:
		return CIPClassID_Overload, true
	case 0x002E:
		return CIPClassID_Selection, true
	case 0x0030:
		return CIPClassID_SDeviceSupervisor, true
	case 0x0031:
		return CIPClassID_SAnalogSensor, true
	case 0x0032:
		return CIPClassID_SAnalogActuator, true
	case 0x0033:
		return CIPClassID_SSingleStageController, true
	case 0x0034:
		return CIPClassID_SGasCalibration, true
	case 0x0037:
		return CIPClassID_File, true
	case 0x0038:
		return CIPClassID_SPartialPressure, true
	case 0x0040:
		return CIPClassID_SSensorCalibration, true
	case 0x0041:
		return CIPClassID_EventLog, true
	case 0x0042:
		return CIPClassID_MotionDeviceAxis, true
	case 0x0044:
		return CIPClassID_Modbus, true
	case 0x0045:
		return CIPClassID_OriginatorConnectionList, true
	case 0x0046:
		return CIPClassID_ModbusSerialLink, true
	case 0x0047:
		return CIPClassID_DeviceLevelRing, true
	case 0x0048:
		return CIPClassID_QOS, true
	case 0x004A:
		return CIPClassID_SafetyAnalogInputGroup, true
	case 0x004C:
		return CIPClassID_Sercos3Link, true
	case 0x004E:
		return CIPClassID_BaseEnergy, true
	case 0x004F:
		return CIPClassID_ElectricalEnergy, true
	case 0x0050:
		return CIPClassID_NonElectricalEnergy, true
	case 0x0051:
		return CIPClassID_BaseSwitch, true
	case 0x0052:
		return CIPClassID_Snmp, true
	case 0x0053:
		return CIPClassID_PowerManagementObject, true
	case 0x0054:
		return CIPClassID_RstpBridge, true
	case 0x0055:
		return CIPClassID_RstpPort, true
	case 0x0056:
		return CIPClassID_ParallelRedundancyProtocol, true
	case 0x0057:
		return CIPClassID_PrpNodesTable, true
	case 0x005C:
		return CIPClassID_PowerCurtailmentObject, true
	case 0x00F0:
		return CIPClassID_ControlNet, true
	case 0x00F1:
		return CIPClassID_ControlNetKeeper, true
	case 0x00F2:
		return CIPClassID_ControlNetScheduling, true
	case 0x00F3:
		return CIPClassID_ConnectionConfiguration, true
	case 0x00F4:
		return CIPClassID_Port, true
	case 0x00F5:
		return CIPClassID_TcpIpInterface, true
	case 0x00F6:
		return CIPClassID_EthernetLink, true
	case 0x00F7:
		return CIPClassID_CompoNetLink, true
	case 0x00F8:
		return CIPClassID_CompoNetRepeater, true
	}
	return 0, false
}

func CIPClassIDByName(value string) (enum CIPClassID, ok bool) {
	switch value {
	case "Identity":
		return CIPClassID_Identity, true
	case "MessageRouter":
		return CIPClassID_MessageRouter, true
	case "DeviceNet":
		return CIPClassID_DeviceNet, true
	case "Assembly":
		return CIPClassID_Assembly, true
	case "Connection":
		return CIPClassID_Connection, true
	case "ConnectionManager":
		return CIPClassID_ConnectionManager, true
	case "Register":
		return CIPClassID_Register, true
	case "DiscreteInputPoint":
		return CIPClassID_DiscreteInputPoint, true
	case "DiscreteOutputPoint":
		return CIPClassID_DiscreteOutputPoint, true
	case "AnalogInputPoint":
		return CIPClassID_AnalogInputPoint, true
	case "AnalogOutputPoint":
		return CIPClassID_AnalogOutputPoint, true
	case "PresenceSensing":
		return CIPClassID_PresenceSensing, true
	case "Parameter":
		return CIPClassID_Parameter, true
	case "ParameterGroup":
		return CIPClassID_ParameterGroup, true
	case "Group":
		return CIPClassID_Group, true
	case "DiscreteInputGroup":
		return CIPClassID_DiscreteInputGroup, true
	case "DiscreteOutputGroup":
		return CIPClassID_DiscreteOutputGroup, true
	case "DiscreteGroup":
		return CIPClassID_DiscreteGroup, true
	case "AnalogInputGroup":
		return CIPClassID_AnalogInputGroup, true
	case "AnalogOutputGroup":
		return CIPClassID_AnalogOutputGroup, true
	case "AnalogGroup":
		return CIPClassID_AnalogGroup, true
	case "PositionSensor":
		return CIPClassID_PositionSensor, true
	case "PositionControllerSupervisor":
		return CIPClassID_PositionControllerSupervisor, true
	case "PositionController":
		return CIPClassID_PositionController, true
	case "BlockSequencer":
		return CIPClassID_BlockSequencer, true
	case "CommandBlock":
		return CIPClassID_CommandBlock, true
	case "MotorData":
		return CIPClassID_MotorData, true
	case "ControlSupervisor":
		return CIPClassID_ControlSupervisor, true
	case "AcDcDrive":
		return CIPClassID_AcDcDrive, true
	case "AcknowledgeHandler":
		return CIPClassID_AcknowledgeHandler, true
	case "Overload":
		return CIPClassID_Overload, true
	case "Selection":
		return CIPClassID_Selection, true
	case "SDeviceSupervisor":
		return CIPClassID_SDeviceSupervisor, true
	case "SAnalogSensor":
		return CIPClassID_SAnalogSensor, true
	case "SAnalogActuator":
		return CIPClassID_SAnalogActuator, true
	case "SSingleStageController":
		return CIPClassID_SSingleStageController, true
	case "SGasCalibration":
		return CIPClassID_SGasCalibration, true
	case "File":
		return CIPClassID_File, true
	case "SPartialPressure":
		return CIPClassID_SPartialPressure, true
	case "SSensorCalibration":
		return CIPClassID_SSensorCalibration, true
	case "EventLog":
		return CIPClassID_EventLog, true
	case "MotionDeviceAxis":
		return CIPClassID_MotionDeviceAxis, true
	case "Modbus":
		return CIPClassID_Modbus, true
	case "OriginatorConnectionList":
		return CIPClassID_OriginatorConnectionList, true
	case "ModbusSerialLink":
		return CIPClassID_ModbusSerialLink, true
	case "DeviceLevelRing":
		return CIPClassID_DeviceLevelRing, true
	case "QOS":
		return CIPClassID_QOS, true
	case "SafetyAnalogInputGroup":
		return CIPClassID_SafetyAnalogInputGroup, true
	case "Sercos3Link":
		return CIPClassID_Sercos3Link, true
	case "BaseEnergy":
		return CIPClassID_BaseEnergy, true
	case "ElectricalEnergy":
		return CIPClassID_ElectricalEnergy, true
	case "NonElectricalEnergy":
		return CIPClassID_NonElectricalEnergy, true
	case "BaseSwitch":
		return CIPClassID_BaseSwitch, true
	case "Snmp":
		return CIPClassID_Snmp, true
	case "PowerManagementObject":
		return CIPClassID_PowerManagementObject, true
	case "RstpBridge":
		return CIPClassID_RstpBridge, true
	case "RstpPort":
		return CIPClassID_RstpPort, true
	case "ParallelRedundancyProtocol":
		return CIPClassID_ParallelRedundancyProtocol, true
	case "PrpNodesTable":
		return CIPClassID_PrpNodesTable, true
	case "PowerCurtailmentObject":
		return CIPClassID_PowerCurtailmentObject, true
	case "ControlNet":
		return CIPClassID_ControlNet, true
	case "ControlNetKeeper":
		return CIPClassID_ControlNetKeeper, true
	case "ControlNetScheduling":
		return CIPClassID_ControlNetScheduling, true
	case "ConnectionConfiguration":
		return CIPClassID_ConnectionConfiguration, true
	case "Port":
		return CIPClassID_Port, true
	case "TcpIpInterface":
		return CIPClassID_TcpIpInterface, true
	case "EthernetLink":
		return CIPClassID_EthernetLink, true
	case "CompoNetLink":
		return CIPClassID_CompoNetLink, true
	case "CompoNetRepeater":
		return CIPClassID_CompoNetRepeater, true
	}
	return 0, false
}

func CIPClassIDKnows(value uint16) bool {
	for _, typeValue := range CIPClassIDValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastCIPClassID(structType any) CIPClassID {
	castFunc := func(typ any) CIPClassID {
		if sCIPClassID, ok := typ.(CIPClassID); ok {
			return sCIPClassID
		}
		return 0
	}
	return castFunc(structType)
}

func (m CIPClassID) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m CIPClassID) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CIPClassIDParse(ctx context.Context, theBytes []byte) (CIPClassID, error) {
	return CIPClassIDParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func CIPClassIDParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CIPClassID, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint16("CIPClassID", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading CIPClassID")
	}
	if enum, ok := CIPClassIDByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for CIPClassID")
		return CIPClassID(val), nil
	} else {
		return enum, nil
	}
}

func (e CIPClassID) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e CIPClassID) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint16("CIPClassID", 16, uint16(uint16(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e CIPClassID) PLC4XEnumName() string {
	switch e {
	case CIPClassID_Identity:
		return "Identity"
	case CIPClassID_MessageRouter:
		return "MessageRouter"
	case CIPClassID_DeviceNet:
		return "DeviceNet"
	case CIPClassID_Assembly:
		return "Assembly"
	case CIPClassID_Connection:
		return "Connection"
	case CIPClassID_ConnectionManager:
		return "ConnectionManager"
	case CIPClassID_Register:
		return "Register"
	case CIPClassID_DiscreteInputPoint:
		return "DiscreteInputPoint"
	case CIPClassID_DiscreteOutputPoint:
		return "DiscreteOutputPoint"
	case CIPClassID_AnalogInputPoint:
		return "AnalogInputPoint"
	case CIPClassID_AnalogOutputPoint:
		return "AnalogOutputPoint"
	case CIPClassID_PresenceSensing:
		return "PresenceSensing"
	case CIPClassID_Parameter:
		return "Parameter"
	case CIPClassID_ParameterGroup:
		return "ParameterGroup"
	case CIPClassID_Group:
		return "Group"
	case CIPClassID_DiscreteInputGroup:
		return "DiscreteInputGroup"
	case CIPClassID_DiscreteOutputGroup:
		return "DiscreteOutputGroup"
	case CIPClassID_DiscreteGroup:
		return "DiscreteGroup"
	case CIPClassID_AnalogInputGroup:
		return "AnalogInputGroup"
	case CIPClassID_AnalogOutputGroup:
		return "AnalogOutputGroup"
	case CIPClassID_AnalogGroup:
		return "AnalogGroup"
	case CIPClassID_PositionSensor:
		return "PositionSensor"
	case CIPClassID_PositionControllerSupervisor:
		return "PositionControllerSupervisor"
	case CIPClassID_PositionController:
		return "PositionController"
	case CIPClassID_BlockSequencer:
		return "BlockSequencer"
	case CIPClassID_CommandBlock:
		return "CommandBlock"
	case CIPClassID_MotorData:
		return "MotorData"
	case CIPClassID_ControlSupervisor:
		return "ControlSupervisor"
	case CIPClassID_AcDcDrive:
		return "AcDcDrive"
	case CIPClassID_AcknowledgeHandler:
		return "AcknowledgeHandler"
	case CIPClassID_Overload:
		return "Overload"
	case CIPClassID_Selection:
		return "Selection"
	case CIPClassID_SDeviceSupervisor:
		return "SDeviceSupervisor"
	case CIPClassID_SAnalogSensor:
		return "SAnalogSensor"
	case CIPClassID_SAnalogActuator:
		return "SAnalogActuator"
	case CIPClassID_SSingleStageController:
		return "SSingleStageController"
	case CIPClassID_SGasCalibration:
		return "SGasCalibration"
	case CIPClassID_File:
		return "File"
	case CIPClassID_SPartialPressure:
		return "SPartialPressure"
	case CIPClassID_SSensorCalibration:
		return "SSensorCalibration"
	case CIPClassID_EventLog:
		return "EventLog"
	case CIPClassID_MotionDeviceAxis:
		return "MotionDeviceAxis"
	case CIPClassID_Modbus:
		return "Modbus"
	case CIPClassID_OriginatorConnectionList:
		return "OriginatorConnectionList"
	case CIPClassID_ModbusSerialLink:
		return "ModbusSerialLink"
	case CIPClassID_DeviceLevelRing:
		return "DeviceLevelRing"
	case CIPClassID_QOS:
		return "QOS"
	case CIPClassID_SafetyAnalogInputGroup:
		return "SafetyAnalogInputGroup"
	case CIPClassID_Sercos3Link:
		return "Sercos3Link"
	case CIPClassID_BaseEnergy:
		return "BaseEnergy"
	case CIPClassID_ElectricalEnergy:
		return "ElectricalEnergy"
	case CIPClassID_NonElectricalEnergy:
		return "NonElectricalEnergy"
	case CIPClassID_BaseSwitch:
		return "BaseSwitch"
	case CIPClassID_Snmp:
		return "Snmp"
	case CIPClassID_PowerManagementObject:
		return "PowerManagementObject"
	case CIPClassID_RstpBridge:
		return "RstpBridge"
	case CIPClassID_RstpPort:
		return "RstpPort"
	case CIPClassID_ParallelRedundancyProtocol:
		return "ParallelRedundancyProtocol"
	case CIPClassID_PrpNodesTable:
		return "PrpNodesTable"
	case CIPClassID_PowerCurtailmentObject:
		return "PowerCurtailmentObject"
	case CIPClassID_ControlNet:
		return "ControlNet"
	case CIPClassID_ControlNetKeeper:
		return "ControlNetKeeper"
	case CIPClassID_ControlNetScheduling:
		return "ControlNetScheduling"
	case CIPClassID_ConnectionConfiguration:
		return "ConnectionConfiguration"
	case CIPClassID_Port:
		return "Port"
	case CIPClassID_TcpIpInterface:
		return "TcpIpInterface"
	case CIPClassID_EthernetLink:
		return "EthernetLink"
	case CIPClassID_CompoNetLink:
		return "CompoNetLink"
	case CIPClassID_CompoNetRepeater:
		return "CompoNetRepeater"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e CIPClassID) String() string {
	return e.PLC4XEnumName()
}
