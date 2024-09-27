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

package readwrite

import (
	"context"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	. "github.com/apache/plc4x/plc4go/protocols/cbus/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

type CbusXmlParserHelper struct {
}

// Temporary imports to silent compiler warnings (TODO: migrate from static to emission based imports)
func init() {
	_ = strconv.ParseUint
	_ = strconv.ParseInt
	_ = strings.Join
	_ = utils.Dump
}

func (m CbusXmlParserHelper) Parse(typeName string, xmlString string, parserArguments ...string) (any, error) {
	switch typeName {
	case "HVACStatusFlags":
		return HVACStatusFlagsParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ParameterValue":
		parameterType, _ := ParameterTypeByName(parserArguments[0])
		parsedUint1, err := strconv.ParseUint(parserArguments[1], 10, 8)
		if err != nil {
			return nil, err
		}
		numBytes := uint8(parsedUint1)
		return ParameterValueParseWithBuffer[ParameterValue](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), parameterType, numBytes)
	case "ReplyOrConfirmation":
		// TODO: find a way to parse the sub types
		var cBusOptions CBusOptions
		// TODO: find a way to parse the sub types
		var requestContext RequestContext
		return ReplyOrConfirmationParseWithBuffer[ReplyOrConfirmation](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions, requestContext)
	case "CBusOptions":
		return CBusOptionsParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "TemperatureBroadcastData":
		return TemperatureBroadcastDataParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "PanicStatus":
		return PanicStatusParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "IdentifyReplyCommandUnitSummary":
		return IdentifyReplyCommandUnitSummaryParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "InterfaceOptions1PowerUpSettings":
		return InterfaceOptions1PowerUpSettingsParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "MonitoredSAL":
		// TODO: find a way to parse the sub types
		var cBusOptions CBusOptions
		return MonitoredSALParseWithBuffer[MonitoredSAL](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions)
	case "ReplyNetwork":
		return ReplyNetworkParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SerialNumber":
		return SerialNumberParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CBusPointToMultiPointCommand":
		// TODO: find a way to parse the sub types
		var cBusOptions CBusOptions
		return CBusPointToMultiPointCommandParseWithBuffer[CBusPointToMultiPointCommand](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions)
	case "StatusRequest":
		return StatusRequestParseWithBuffer[StatusRequest](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "InterfaceOptions3":
		return InterfaceOptions3ParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "InterfaceOptions1":
		return InterfaceOptions1ParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "InterfaceOptions2":
		return InterfaceOptions2ParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "HVACModeAndFlags":
		return HVACModeAndFlagsParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "LightingData":
		return LightingDataParseWithBuffer[LightingData](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SALData":
		applicationId, _ := ApplicationIdByName(parserArguments[0])
		return SALDataParseWithBuffer[SALData](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), applicationId)
	case "CBusCommand":
		// TODO: find a way to parse the sub types
		var cBusOptions CBusOptions
		return CBusCommandParseWithBuffer[CBusCommand](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions)
	case "HVACHumidity":
		return HVACHumidityParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "HVACHumidityModeAndFlags":
		return HVACHumidityModeAndFlagsParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CBusConstants":
		return CBusConstantsParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SerialInterfaceAddress":
		return SerialInterfaceAddressParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "MeasurementData":
		return MeasurementDataParseWithBuffer[MeasurementData](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "HVACZoneList":
		return HVACZoneListParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "MediaTransportControlData":
		return MediaTransportControlDataParseWithBuffer[MediaTransportControlData](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "StatusByte":
		return StatusByteParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "TriggerControlLabelOptions":
		return TriggerControlLabelOptionsParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "HVACAuxiliaryLevel":
		return HVACAuxiliaryLevelParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ErrorReportingData":
		return ErrorReportingDataParseWithBuffer[ErrorReportingData](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "UnitAddress":
		return UnitAddressParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SecurityArmCode":
		return SecurityArmCodeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "MeteringData":
		return MeteringDataParseWithBuffer[MeteringData](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "EnableControlData":
		return EnableControlDataParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ApplicationAddress2":
		return ApplicationAddress2ParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ApplicationAddress1":
		return ApplicationAddress1ParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "RequestContext":
		return RequestContextParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "TriggerControlData":
		return TriggerControlDataParseWithBuffer[TriggerControlData](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "HVACStartTime":
		return HVACStartTimeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "HVACTemperature":
		return HVACTemperatureParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "RequestTermination":
		return RequestTerminationParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CBusMessage":
		isResponse := parserArguments[0] == "true"
		// TODO: find a way to parse the sub types
		var requestContext RequestContext
		// TODO: find a way to parse the sub types
		var cBusOptions CBusOptions
		return CBusMessageParseWithBuffer[CBusMessage](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), isResponse, requestContext, cBusOptions)
	case "ErrorReportingSystemCategory":
		return ErrorReportingSystemCategoryParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "PowerUp":
		return PowerUpParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "Reply":
		// TODO: find a way to parse the sub types
		var cBusOptions CBusOptions
		// TODO: find a way to parse the sub types
		var requestContext RequestContext
		return ReplyParseWithBuffer[Reply](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions, requestContext)
	case "TelephonyData":
		return TelephonyDataParseWithBuffer[TelephonyData](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "HVACHumidityStatusFlags":
		return HVACHumidityStatusFlagsParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ParameterChange":
		return ParameterChangeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ErrorReportingSystemCategoryType":
		errorReportingSystemCategoryClass, _ := ErrorReportingSystemCategoryClassByName(parserArguments[0])
		return ErrorReportingSystemCategoryTypeParseWithBuffer[ErrorReportingSystemCategoryType](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), errorReportingSystemCategoryClass)
	case "Confirmation":
		return ConfirmationParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SecurityData":
		return SecurityDataParseWithBuffer[SecurityData](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NetworkProtocolControlInformation":
		return NetworkProtocolControlInformationParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CBusHeader":
		return CBusHeaderParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "Request":
		// TODO: find a way to parse the sub types
		var cBusOptions CBusOptions
		return RequestParseWithBuffer[Request](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions)
	case "Alpha":
		return AlphaParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CALData":
		// TODO: find a way to parse the sub types
		var requestContext RequestContext
		return CALDataParseWithBuffer[CALData](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), requestContext)
	case "Checksum":
		return ChecksumParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CALReply":
		// TODO: find a way to parse the sub types
		var cBusOptions CBusOptions
		// TODO: find a way to parse the sub types
		var requestContext RequestContext
		return CALReplyParseWithBuffer[CALReply](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions, requestContext)
	case "CustomManufacturer":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		numBytes := uint8(parsedUint0)
		return CustomManufacturerParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), numBytes)
	case "AccessControlData":
		return AccessControlDataParseWithBuffer[AccessControlData](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ClockAndTimekeepingData":
		return ClockAndTimekeepingDataParseWithBuffer[ClockAndTimekeepingData](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NetworkRoute":
		return NetworkRouteParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ResponseTermination":
		return ResponseTerminationParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "LevelInformation":
		return LevelInformationParseWithBuffer[LevelInformation](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "TamperStatus":
		return TamperStatusParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "IdentifyReplyCommand":
		attribute, _ := AttributeByName(parserArguments[0])
		parsedUint1, err := strconv.ParseUint(parserArguments[1], 10, 5)
		if err != nil {
			return nil, err
		}
		numBytes := uint8(parsedUint1)
		return IdentifyReplyCommandParseWithBuffer[IdentifyReplyCommand](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), attribute, numBytes)
	case "HVACRawLevels":
		return HVACRawLevelsParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ZoneStatus":
		return ZoneStatusParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BridgeAddress":
		return BridgeAddressParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "LightingLabelOptions":
		return LightingLabelOptionsParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "CustomTypes":
		parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 8)
		if err != nil {
			return nil, err
		}
		numBytes := uint8(parsedUint0)
		return CustomTypesParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), numBytes)
	case "EncodedReply":
		// TODO: find a way to parse the sub types
		var cBusOptions CBusOptions
		// TODO: find a way to parse the sub types
		var requestContext RequestContext
		return EncodedReplyParseWithBuffer[EncodedReply](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions, requestContext)
	case "CBusPointToPointToMultiPointCommand":
		// TODO: find a way to parse the sub types
		var cBusOptions CBusOptions
		return CBusPointToPointToMultiPointCommandParseWithBuffer[CBusPointToPointToMultiPointCommand](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions)
	case "CBusPointToPointCommand":
		// TODO: find a way to parse the sub types
		var cBusOptions CBusOptions
		return CBusPointToPointCommandParseWithBuffer[CBusPointToPointCommand](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), cBusOptions)
	case "AirConditioningData":
		return AirConditioningDataParseWithBuffer[AirConditioningData](context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "LogicAssignment":
		return LogicAssignmentParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
