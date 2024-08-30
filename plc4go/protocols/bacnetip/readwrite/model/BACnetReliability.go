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

// BACnetReliability is an enum
type BACnetReliability uint16

type IBACnetReliability interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetReliability_NO_FAULT_DETECTED                BACnetReliability = 0
	BACnetReliability_NO_SENSOR                        BACnetReliability = 1
	BACnetReliability_OVER_RANGE                       BACnetReliability = 2
	BACnetReliability_UNDER_RANGE                      BACnetReliability = 3
	BACnetReliability_OPEN_LOOP                        BACnetReliability = 4
	BACnetReliability_SHORTED_LOOP                     BACnetReliability = 5
	BACnetReliability_NO_OUTPUT                        BACnetReliability = 6
	BACnetReliability_UNRELIABLE_OTHER                 BACnetReliability = 7
	BACnetReliability_PROCESS_ERROR                    BACnetReliability = 8
	BACnetReliability_MULTI_STATE_FAULT                BACnetReliability = 9
	BACnetReliability_CONFIGURATION_ERROR              BACnetReliability = 10
	BACnetReliability_COMMUNICATION_FAILURE            BACnetReliability = 12
	BACnetReliability_MEMBER_FAULT                     BACnetReliability = 13
	BACnetReliability_MONITORED_OBJECT_FAULT           BACnetReliability = 14
	BACnetReliability_TRIPPED                          BACnetReliability = 15
	BACnetReliability_LAMP_FAILURE                     BACnetReliability = 16
	BACnetReliability_ACTIVATION_FAILURE               BACnetReliability = 17
	BACnetReliability_RENEW_DHCP_FAILURE               BACnetReliability = 18
	BACnetReliability_RENEW_FD_REGISTRATION_FAILURE    BACnetReliability = 19
	BACnetReliability_RESTART_AUTO_NEGOTIATION_FAILURE BACnetReliability = 20
	BACnetReliability_RESTART_FAILURE                  BACnetReliability = 21
	BACnetReliability_PROPRIETARY_COMMAND_FAILURE      BACnetReliability = 22
	BACnetReliability_FAULTS_LISTED                    BACnetReliability = 23
	BACnetReliability_REFERENCED_OBJECT_FAULT          BACnetReliability = 24
	BACnetReliability_VENDOR_PROPRIETARY_VALUE         BACnetReliability = 0xFFFF
)

var BACnetReliabilityValues []BACnetReliability

func init() {
	_ = errors.New
	BACnetReliabilityValues = []BACnetReliability{
		BACnetReliability_NO_FAULT_DETECTED,
		BACnetReliability_NO_SENSOR,
		BACnetReliability_OVER_RANGE,
		BACnetReliability_UNDER_RANGE,
		BACnetReliability_OPEN_LOOP,
		BACnetReliability_SHORTED_LOOP,
		BACnetReliability_NO_OUTPUT,
		BACnetReliability_UNRELIABLE_OTHER,
		BACnetReliability_PROCESS_ERROR,
		BACnetReliability_MULTI_STATE_FAULT,
		BACnetReliability_CONFIGURATION_ERROR,
		BACnetReliability_COMMUNICATION_FAILURE,
		BACnetReliability_MEMBER_FAULT,
		BACnetReliability_MONITORED_OBJECT_FAULT,
		BACnetReliability_TRIPPED,
		BACnetReliability_LAMP_FAILURE,
		BACnetReliability_ACTIVATION_FAILURE,
		BACnetReliability_RENEW_DHCP_FAILURE,
		BACnetReliability_RENEW_FD_REGISTRATION_FAILURE,
		BACnetReliability_RESTART_AUTO_NEGOTIATION_FAILURE,
		BACnetReliability_RESTART_FAILURE,
		BACnetReliability_PROPRIETARY_COMMAND_FAILURE,
		BACnetReliability_FAULTS_LISTED,
		BACnetReliability_REFERENCED_OBJECT_FAULT,
		BACnetReliability_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetReliabilityByValue(value uint16) (enum BACnetReliability, ok bool) {
	switch value {
	case 0:
		return BACnetReliability_NO_FAULT_DETECTED, true
	case 0xFFFF:
		return BACnetReliability_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetReliability_NO_SENSOR, true
	case 10:
		return BACnetReliability_CONFIGURATION_ERROR, true
	case 12:
		return BACnetReliability_COMMUNICATION_FAILURE, true
	case 13:
		return BACnetReliability_MEMBER_FAULT, true
	case 14:
		return BACnetReliability_MONITORED_OBJECT_FAULT, true
	case 15:
		return BACnetReliability_TRIPPED, true
	case 16:
		return BACnetReliability_LAMP_FAILURE, true
	case 17:
		return BACnetReliability_ACTIVATION_FAILURE, true
	case 18:
		return BACnetReliability_RENEW_DHCP_FAILURE, true
	case 19:
		return BACnetReliability_RENEW_FD_REGISTRATION_FAILURE, true
	case 2:
		return BACnetReliability_OVER_RANGE, true
	case 20:
		return BACnetReliability_RESTART_AUTO_NEGOTIATION_FAILURE, true
	case 21:
		return BACnetReliability_RESTART_FAILURE, true
	case 22:
		return BACnetReliability_PROPRIETARY_COMMAND_FAILURE, true
	case 23:
		return BACnetReliability_FAULTS_LISTED, true
	case 24:
		return BACnetReliability_REFERENCED_OBJECT_FAULT, true
	case 3:
		return BACnetReliability_UNDER_RANGE, true
	case 4:
		return BACnetReliability_OPEN_LOOP, true
	case 5:
		return BACnetReliability_SHORTED_LOOP, true
	case 6:
		return BACnetReliability_NO_OUTPUT, true
	case 7:
		return BACnetReliability_UNRELIABLE_OTHER, true
	case 8:
		return BACnetReliability_PROCESS_ERROR, true
	case 9:
		return BACnetReliability_MULTI_STATE_FAULT, true
	}
	return 0, false
}

func BACnetReliabilityByName(value string) (enum BACnetReliability, ok bool) {
	switch value {
	case "NO_FAULT_DETECTED":
		return BACnetReliability_NO_FAULT_DETECTED, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetReliability_VENDOR_PROPRIETARY_VALUE, true
	case "NO_SENSOR":
		return BACnetReliability_NO_SENSOR, true
	case "CONFIGURATION_ERROR":
		return BACnetReliability_CONFIGURATION_ERROR, true
	case "COMMUNICATION_FAILURE":
		return BACnetReliability_COMMUNICATION_FAILURE, true
	case "MEMBER_FAULT":
		return BACnetReliability_MEMBER_FAULT, true
	case "MONITORED_OBJECT_FAULT":
		return BACnetReliability_MONITORED_OBJECT_FAULT, true
	case "TRIPPED":
		return BACnetReliability_TRIPPED, true
	case "LAMP_FAILURE":
		return BACnetReliability_LAMP_FAILURE, true
	case "ACTIVATION_FAILURE":
		return BACnetReliability_ACTIVATION_FAILURE, true
	case "RENEW_DHCP_FAILURE":
		return BACnetReliability_RENEW_DHCP_FAILURE, true
	case "RENEW_FD_REGISTRATION_FAILURE":
		return BACnetReliability_RENEW_FD_REGISTRATION_FAILURE, true
	case "OVER_RANGE":
		return BACnetReliability_OVER_RANGE, true
	case "RESTART_AUTO_NEGOTIATION_FAILURE":
		return BACnetReliability_RESTART_AUTO_NEGOTIATION_FAILURE, true
	case "RESTART_FAILURE":
		return BACnetReliability_RESTART_FAILURE, true
	case "PROPRIETARY_COMMAND_FAILURE":
		return BACnetReliability_PROPRIETARY_COMMAND_FAILURE, true
	case "FAULTS_LISTED":
		return BACnetReliability_FAULTS_LISTED, true
	case "REFERENCED_OBJECT_FAULT":
		return BACnetReliability_REFERENCED_OBJECT_FAULT, true
	case "UNDER_RANGE":
		return BACnetReliability_UNDER_RANGE, true
	case "OPEN_LOOP":
		return BACnetReliability_OPEN_LOOP, true
	case "SHORTED_LOOP":
		return BACnetReliability_SHORTED_LOOP, true
	case "NO_OUTPUT":
		return BACnetReliability_NO_OUTPUT, true
	case "UNRELIABLE_OTHER":
		return BACnetReliability_UNRELIABLE_OTHER, true
	case "PROCESS_ERROR":
		return BACnetReliability_PROCESS_ERROR, true
	case "MULTI_STATE_FAULT":
		return BACnetReliability_MULTI_STATE_FAULT, true
	}
	return 0, false
}

func BACnetReliabilityKnows(value uint16) bool {
	for _, typeValue := range BACnetReliabilityValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetReliability(structType any) BACnetReliability {
	castFunc := func(typ any) BACnetReliability {
		if sBACnetReliability, ok := typ.(BACnetReliability); ok {
			return sBACnetReliability
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetReliability) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m BACnetReliability) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetReliabilityParse(ctx context.Context, theBytes []byte) (BACnetReliability, error) {
	return BACnetReliabilityParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetReliabilityParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetReliability, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint16("BACnetReliability", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetReliability")
	}
	if enum, ok := BACnetReliabilityByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetReliability")
		return BACnetReliability(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetReliability) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetReliability) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint16("BACnetReliability", 16, uint16(uint16(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetReliability) PLC4XEnumName() string {
	switch e {
	case BACnetReliability_NO_FAULT_DETECTED:
		return "NO_FAULT_DETECTED"
	case BACnetReliability_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetReliability_NO_SENSOR:
		return "NO_SENSOR"
	case BACnetReliability_CONFIGURATION_ERROR:
		return "CONFIGURATION_ERROR"
	case BACnetReliability_COMMUNICATION_FAILURE:
		return "COMMUNICATION_FAILURE"
	case BACnetReliability_MEMBER_FAULT:
		return "MEMBER_FAULT"
	case BACnetReliability_MONITORED_OBJECT_FAULT:
		return "MONITORED_OBJECT_FAULT"
	case BACnetReliability_TRIPPED:
		return "TRIPPED"
	case BACnetReliability_LAMP_FAILURE:
		return "LAMP_FAILURE"
	case BACnetReliability_ACTIVATION_FAILURE:
		return "ACTIVATION_FAILURE"
	case BACnetReliability_RENEW_DHCP_FAILURE:
		return "RENEW_DHCP_FAILURE"
	case BACnetReliability_RENEW_FD_REGISTRATION_FAILURE:
		return "RENEW_FD_REGISTRATION_FAILURE"
	case BACnetReliability_OVER_RANGE:
		return "OVER_RANGE"
	case BACnetReliability_RESTART_AUTO_NEGOTIATION_FAILURE:
		return "RESTART_AUTO_NEGOTIATION_FAILURE"
	case BACnetReliability_RESTART_FAILURE:
		return "RESTART_FAILURE"
	case BACnetReliability_PROPRIETARY_COMMAND_FAILURE:
		return "PROPRIETARY_COMMAND_FAILURE"
	case BACnetReliability_FAULTS_LISTED:
		return "FAULTS_LISTED"
	case BACnetReliability_REFERENCED_OBJECT_FAULT:
		return "REFERENCED_OBJECT_FAULT"
	case BACnetReliability_UNDER_RANGE:
		return "UNDER_RANGE"
	case BACnetReliability_OPEN_LOOP:
		return "OPEN_LOOP"
	case BACnetReliability_SHORTED_LOOP:
		return "SHORTED_LOOP"
	case BACnetReliability_NO_OUTPUT:
		return "NO_OUTPUT"
	case BACnetReliability_UNRELIABLE_OTHER:
		return "UNRELIABLE_OTHER"
	case BACnetReliability_PROCESS_ERROR:
		return "PROCESS_ERROR"
	case BACnetReliability_MULTI_STATE_FAULT:
		return "MULTI_STATE_FAULT"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e BACnetReliability) String() string {
	return e.PLC4XEnumName()
}
