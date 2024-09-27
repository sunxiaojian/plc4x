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

// BACnetObjectType is an enum
type BACnetObjectType uint16

type IBACnetObjectType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetObjectType_ACCESS_CREDENTIAL        BACnetObjectType = 32
	BACnetObjectType_ACCESS_DOOR              BACnetObjectType = 30
	BACnetObjectType_ACCESS_POINT             BACnetObjectType = 33
	BACnetObjectType_ACCESS_RIGHTS            BACnetObjectType = 34
	BACnetObjectType_ACCESS_USER              BACnetObjectType = 35
	BACnetObjectType_ACCESS_ZONE              BACnetObjectType = 36
	BACnetObjectType_ACCUMULATOR              BACnetObjectType = 23
	BACnetObjectType_ALERT_ENROLLMENT         BACnetObjectType = 52
	BACnetObjectType_ANALOG_INPUT             BACnetObjectType = 0
	BACnetObjectType_ANALOG_OUTPUT            BACnetObjectType = 1
	BACnetObjectType_ANALOG_VALUE             BACnetObjectType = 2
	BACnetObjectType_AVERAGING                BACnetObjectType = 18
	BACnetObjectType_BINARY_INPUT             BACnetObjectType = 3
	BACnetObjectType_BINARY_LIGHTING_OUTPUT   BACnetObjectType = 55
	BACnetObjectType_BINARY_OUTPUT            BACnetObjectType = 4
	BACnetObjectType_BINARY_VALUE             BACnetObjectType = 5
	BACnetObjectType_BITSTRING_VALUE          BACnetObjectType = 39
	BACnetObjectType_CALENDAR                 BACnetObjectType = 6
	BACnetObjectType_CHANNEL                  BACnetObjectType = 53
	BACnetObjectType_CHARACTERSTRING_VALUE    BACnetObjectType = 40
	BACnetObjectType_COMMAND                  BACnetObjectType = 7
	BACnetObjectType_CREDENTIAL_DATA_INPUT    BACnetObjectType = 37
	BACnetObjectType_DATEPATTERN_VALUE        BACnetObjectType = 41
	BACnetObjectType_DATE_VALUE               BACnetObjectType = 42
	BACnetObjectType_DATETIMEPATTERN_VALUE    BACnetObjectType = 43
	BACnetObjectType_DATETIME_VALUE           BACnetObjectType = 44
	BACnetObjectType_DEVICE                   BACnetObjectType = 8
	BACnetObjectType_ELEVATOR_GROUP           BACnetObjectType = 57
	BACnetObjectType_ESCALATOR                BACnetObjectType = 58
	BACnetObjectType_EVENT_ENROLLMENT         BACnetObjectType = 9
	BACnetObjectType_EVENT_LOG                BACnetObjectType = 25
	BACnetObjectType_FILE                     BACnetObjectType = 10
	BACnetObjectType_GLOBAL_GROUP             BACnetObjectType = 26
	BACnetObjectType_GROUP                    BACnetObjectType = 11
	BACnetObjectType_INTEGER_VALUE            BACnetObjectType = 45
	BACnetObjectType_LARGE_ANALOG_VALUE       BACnetObjectType = 46
	BACnetObjectType_LIFE_SAFETY_POINT        BACnetObjectType = 21
	BACnetObjectType_LIFE_SAFETY_ZONE         BACnetObjectType = 22
	BACnetObjectType_LIFT                     BACnetObjectType = 59
	BACnetObjectType_LIGHTING_OUTPUT          BACnetObjectType = 54
	BACnetObjectType_LOAD_CONTROL             BACnetObjectType = 28
	BACnetObjectType_LOOP                     BACnetObjectType = 12
	BACnetObjectType_MULTI_STATE_INPUT        BACnetObjectType = 13
	BACnetObjectType_MULTI_STATE_OUTPUT       BACnetObjectType = 14
	BACnetObjectType_MULTI_STATE_VALUE        BACnetObjectType = 19
	BACnetObjectType_NETWORK_PORT             BACnetObjectType = 56
	BACnetObjectType_NETWORK_SECURITY         BACnetObjectType = 38
	BACnetObjectType_NOTIFICATION_CLASS       BACnetObjectType = 15
	BACnetObjectType_NOTIFICATION_FORWARDER   BACnetObjectType = 51
	BACnetObjectType_OCTETSTRING_VALUE        BACnetObjectType = 47
	BACnetObjectType_POSITIVE_INTEGER_VALUE   BACnetObjectType = 48
	BACnetObjectType_PROGRAM                  BACnetObjectType = 16
	BACnetObjectType_PULSE_CONVERTER          BACnetObjectType = 24
	BACnetObjectType_SCHEDULE                 BACnetObjectType = 17
	BACnetObjectType_STRUCTURED_VIEW          BACnetObjectType = 29
	BACnetObjectType_TIMEPATTERN_VALUE        BACnetObjectType = 49
	BACnetObjectType_TIME_VALUE               BACnetObjectType = 50
	BACnetObjectType_TIMER                    BACnetObjectType = 31
	BACnetObjectType_TREND_LOG                BACnetObjectType = 20
	BACnetObjectType_TREND_LOG_MULTIPLE       BACnetObjectType = 27
	BACnetObjectType_VENDOR_PROPRIETARY_VALUE BACnetObjectType = 0x3FF
)

var BACnetObjectTypeValues []BACnetObjectType

func init() {
	_ = errors.New
	BACnetObjectTypeValues = []BACnetObjectType{
		BACnetObjectType_ACCESS_CREDENTIAL,
		BACnetObjectType_ACCESS_DOOR,
		BACnetObjectType_ACCESS_POINT,
		BACnetObjectType_ACCESS_RIGHTS,
		BACnetObjectType_ACCESS_USER,
		BACnetObjectType_ACCESS_ZONE,
		BACnetObjectType_ACCUMULATOR,
		BACnetObjectType_ALERT_ENROLLMENT,
		BACnetObjectType_ANALOG_INPUT,
		BACnetObjectType_ANALOG_OUTPUT,
		BACnetObjectType_ANALOG_VALUE,
		BACnetObjectType_AVERAGING,
		BACnetObjectType_BINARY_INPUT,
		BACnetObjectType_BINARY_LIGHTING_OUTPUT,
		BACnetObjectType_BINARY_OUTPUT,
		BACnetObjectType_BINARY_VALUE,
		BACnetObjectType_BITSTRING_VALUE,
		BACnetObjectType_CALENDAR,
		BACnetObjectType_CHANNEL,
		BACnetObjectType_CHARACTERSTRING_VALUE,
		BACnetObjectType_COMMAND,
		BACnetObjectType_CREDENTIAL_DATA_INPUT,
		BACnetObjectType_DATEPATTERN_VALUE,
		BACnetObjectType_DATE_VALUE,
		BACnetObjectType_DATETIMEPATTERN_VALUE,
		BACnetObjectType_DATETIME_VALUE,
		BACnetObjectType_DEVICE,
		BACnetObjectType_ELEVATOR_GROUP,
		BACnetObjectType_ESCALATOR,
		BACnetObjectType_EVENT_ENROLLMENT,
		BACnetObjectType_EVENT_LOG,
		BACnetObjectType_FILE,
		BACnetObjectType_GLOBAL_GROUP,
		BACnetObjectType_GROUP,
		BACnetObjectType_INTEGER_VALUE,
		BACnetObjectType_LARGE_ANALOG_VALUE,
		BACnetObjectType_LIFE_SAFETY_POINT,
		BACnetObjectType_LIFE_SAFETY_ZONE,
		BACnetObjectType_LIFT,
		BACnetObjectType_LIGHTING_OUTPUT,
		BACnetObjectType_LOAD_CONTROL,
		BACnetObjectType_LOOP,
		BACnetObjectType_MULTI_STATE_INPUT,
		BACnetObjectType_MULTI_STATE_OUTPUT,
		BACnetObjectType_MULTI_STATE_VALUE,
		BACnetObjectType_NETWORK_PORT,
		BACnetObjectType_NETWORK_SECURITY,
		BACnetObjectType_NOTIFICATION_CLASS,
		BACnetObjectType_NOTIFICATION_FORWARDER,
		BACnetObjectType_OCTETSTRING_VALUE,
		BACnetObjectType_POSITIVE_INTEGER_VALUE,
		BACnetObjectType_PROGRAM,
		BACnetObjectType_PULSE_CONVERTER,
		BACnetObjectType_SCHEDULE,
		BACnetObjectType_STRUCTURED_VIEW,
		BACnetObjectType_TIMEPATTERN_VALUE,
		BACnetObjectType_TIME_VALUE,
		BACnetObjectType_TIMER,
		BACnetObjectType_TREND_LOG,
		BACnetObjectType_TREND_LOG_MULTIPLE,
		BACnetObjectType_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetObjectTypeByValue(value uint16) (enum BACnetObjectType, ok bool) {
	switch value {
	case 0:
		return BACnetObjectType_ANALOG_INPUT, true
	case 0x3FF:
		return BACnetObjectType_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetObjectType_ANALOG_OUTPUT, true
	case 10:
		return BACnetObjectType_FILE, true
	case 11:
		return BACnetObjectType_GROUP, true
	case 12:
		return BACnetObjectType_LOOP, true
	case 13:
		return BACnetObjectType_MULTI_STATE_INPUT, true
	case 14:
		return BACnetObjectType_MULTI_STATE_OUTPUT, true
	case 15:
		return BACnetObjectType_NOTIFICATION_CLASS, true
	case 16:
		return BACnetObjectType_PROGRAM, true
	case 17:
		return BACnetObjectType_SCHEDULE, true
	case 18:
		return BACnetObjectType_AVERAGING, true
	case 19:
		return BACnetObjectType_MULTI_STATE_VALUE, true
	case 2:
		return BACnetObjectType_ANALOG_VALUE, true
	case 20:
		return BACnetObjectType_TREND_LOG, true
	case 21:
		return BACnetObjectType_LIFE_SAFETY_POINT, true
	case 22:
		return BACnetObjectType_LIFE_SAFETY_ZONE, true
	case 23:
		return BACnetObjectType_ACCUMULATOR, true
	case 24:
		return BACnetObjectType_PULSE_CONVERTER, true
	case 25:
		return BACnetObjectType_EVENT_LOG, true
	case 26:
		return BACnetObjectType_GLOBAL_GROUP, true
	case 27:
		return BACnetObjectType_TREND_LOG_MULTIPLE, true
	case 28:
		return BACnetObjectType_LOAD_CONTROL, true
	case 29:
		return BACnetObjectType_STRUCTURED_VIEW, true
	case 3:
		return BACnetObjectType_BINARY_INPUT, true
	case 30:
		return BACnetObjectType_ACCESS_DOOR, true
	case 31:
		return BACnetObjectType_TIMER, true
	case 32:
		return BACnetObjectType_ACCESS_CREDENTIAL, true
	case 33:
		return BACnetObjectType_ACCESS_POINT, true
	case 34:
		return BACnetObjectType_ACCESS_RIGHTS, true
	case 35:
		return BACnetObjectType_ACCESS_USER, true
	case 36:
		return BACnetObjectType_ACCESS_ZONE, true
	case 37:
		return BACnetObjectType_CREDENTIAL_DATA_INPUT, true
	case 38:
		return BACnetObjectType_NETWORK_SECURITY, true
	case 39:
		return BACnetObjectType_BITSTRING_VALUE, true
	case 4:
		return BACnetObjectType_BINARY_OUTPUT, true
	case 40:
		return BACnetObjectType_CHARACTERSTRING_VALUE, true
	case 41:
		return BACnetObjectType_DATEPATTERN_VALUE, true
	case 42:
		return BACnetObjectType_DATE_VALUE, true
	case 43:
		return BACnetObjectType_DATETIMEPATTERN_VALUE, true
	case 44:
		return BACnetObjectType_DATETIME_VALUE, true
	case 45:
		return BACnetObjectType_INTEGER_VALUE, true
	case 46:
		return BACnetObjectType_LARGE_ANALOG_VALUE, true
	case 47:
		return BACnetObjectType_OCTETSTRING_VALUE, true
	case 48:
		return BACnetObjectType_POSITIVE_INTEGER_VALUE, true
	case 49:
		return BACnetObjectType_TIMEPATTERN_VALUE, true
	case 5:
		return BACnetObjectType_BINARY_VALUE, true
	case 50:
		return BACnetObjectType_TIME_VALUE, true
	case 51:
		return BACnetObjectType_NOTIFICATION_FORWARDER, true
	case 52:
		return BACnetObjectType_ALERT_ENROLLMENT, true
	case 53:
		return BACnetObjectType_CHANNEL, true
	case 54:
		return BACnetObjectType_LIGHTING_OUTPUT, true
	case 55:
		return BACnetObjectType_BINARY_LIGHTING_OUTPUT, true
	case 56:
		return BACnetObjectType_NETWORK_PORT, true
	case 57:
		return BACnetObjectType_ELEVATOR_GROUP, true
	case 58:
		return BACnetObjectType_ESCALATOR, true
	case 59:
		return BACnetObjectType_LIFT, true
	case 6:
		return BACnetObjectType_CALENDAR, true
	case 7:
		return BACnetObjectType_COMMAND, true
	case 8:
		return BACnetObjectType_DEVICE, true
	case 9:
		return BACnetObjectType_EVENT_ENROLLMENT, true
	}
	return 0, false
}

func BACnetObjectTypeByName(value string) (enum BACnetObjectType, ok bool) {
	switch value {
	case "ANALOG_INPUT":
		return BACnetObjectType_ANALOG_INPUT, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetObjectType_VENDOR_PROPRIETARY_VALUE, true
	case "ANALOG_OUTPUT":
		return BACnetObjectType_ANALOG_OUTPUT, true
	case "FILE":
		return BACnetObjectType_FILE, true
	case "GROUP":
		return BACnetObjectType_GROUP, true
	case "LOOP":
		return BACnetObjectType_LOOP, true
	case "MULTI_STATE_INPUT":
		return BACnetObjectType_MULTI_STATE_INPUT, true
	case "MULTI_STATE_OUTPUT":
		return BACnetObjectType_MULTI_STATE_OUTPUT, true
	case "NOTIFICATION_CLASS":
		return BACnetObjectType_NOTIFICATION_CLASS, true
	case "PROGRAM":
		return BACnetObjectType_PROGRAM, true
	case "SCHEDULE":
		return BACnetObjectType_SCHEDULE, true
	case "AVERAGING":
		return BACnetObjectType_AVERAGING, true
	case "MULTI_STATE_VALUE":
		return BACnetObjectType_MULTI_STATE_VALUE, true
	case "ANALOG_VALUE":
		return BACnetObjectType_ANALOG_VALUE, true
	case "TREND_LOG":
		return BACnetObjectType_TREND_LOG, true
	case "LIFE_SAFETY_POINT":
		return BACnetObjectType_LIFE_SAFETY_POINT, true
	case "LIFE_SAFETY_ZONE":
		return BACnetObjectType_LIFE_SAFETY_ZONE, true
	case "ACCUMULATOR":
		return BACnetObjectType_ACCUMULATOR, true
	case "PULSE_CONVERTER":
		return BACnetObjectType_PULSE_CONVERTER, true
	case "EVENT_LOG":
		return BACnetObjectType_EVENT_LOG, true
	case "GLOBAL_GROUP":
		return BACnetObjectType_GLOBAL_GROUP, true
	case "TREND_LOG_MULTIPLE":
		return BACnetObjectType_TREND_LOG_MULTIPLE, true
	case "LOAD_CONTROL":
		return BACnetObjectType_LOAD_CONTROL, true
	case "STRUCTURED_VIEW":
		return BACnetObjectType_STRUCTURED_VIEW, true
	case "BINARY_INPUT":
		return BACnetObjectType_BINARY_INPUT, true
	case "ACCESS_DOOR":
		return BACnetObjectType_ACCESS_DOOR, true
	case "TIMER":
		return BACnetObjectType_TIMER, true
	case "ACCESS_CREDENTIAL":
		return BACnetObjectType_ACCESS_CREDENTIAL, true
	case "ACCESS_POINT":
		return BACnetObjectType_ACCESS_POINT, true
	case "ACCESS_RIGHTS":
		return BACnetObjectType_ACCESS_RIGHTS, true
	case "ACCESS_USER":
		return BACnetObjectType_ACCESS_USER, true
	case "ACCESS_ZONE":
		return BACnetObjectType_ACCESS_ZONE, true
	case "CREDENTIAL_DATA_INPUT":
		return BACnetObjectType_CREDENTIAL_DATA_INPUT, true
	case "NETWORK_SECURITY":
		return BACnetObjectType_NETWORK_SECURITY, true
	case "BITSTRING_VALUE":
		return BACnetObjectType_BITSTRING_VALUE, true
	case "BINARY_OUTPUT":
		return BACnetObjectType_BINARY_OUTPUT, true
	case "CHARACTERSTRING_VALUE":
		return BACnetObjectType_CHARACTERSTRING_VALUE, true
	case "DATEPATTERN_VALUE":
		return BACnetObjectType_DATEPATTERN_VALUE, true
	case "DATE_VALUE":
		return BACnetObjectType_DATE_VALUE, true
	case "DATETIMEPATTERN_VALUE":
		return BACnetObjectType_DATETIMEPATTERN_VALUE, true
	case "DATETIME_VALUE":
		return BACnetObjectType_DATETIME_VALUE, true
	case "INTEGER_VALUE":
		return BACnetObjectType_INTEGER_VALUE, true
	case "LARGE_ANALOG_VALUE":
		return BACnetObjectType_LARGE_ANALOG_VALUE, true
	case "OCTETSTRING_VALUE":
		return BACnetObjectType_OCTETSTRING_VALUE, true
	case "POSITIVE_INTEGER_VALUE":
		return BACnetObjectType_POSITIVE_INTEGER_VALUE, true
	case "TIMEPATTERN_VALUE":
		return BACnetObjectType_TIMEPATTERN_VALUE, true
	case "BINARY_VALUE":
		return BACnetObjectType_BINARY_VALUE, true
	case "TIME_VALUE":
		return BACnetObjectType_TIME_VALUE, true
	case "NOTIFICATION_FORWARDER":
		return BACnetObjectType_NOTIFICATION_FORWARDER, true
	case "ALERT_ENROLLMENT":
		return BACnetObjectType_ALERT_ENROLLMENT, true
	case "CHANNEL":
		return BACnetObjectType_CHANNEL, true
	case "LIGHTING_OUTPUT":
		return BACnetObjectType_LIGHTING_OUTPUT, true
	case "BINARY_LIGHTING_OUTPUT":
		return BACnetObjectType_BINARY_LIGHTING_OUTPUT, true
	case "NETWORK_PORT":
		return BACnetObjectType_NETWORK_PORT, true
	case "ELEVATOR_GROUP":
		return BACnetObjectType_ELEVATOR_GROUP, true
	case "ESCALATOR":
		return BACnetObjectType_ESCALATOR, true
	case "LIFT":
		return BACnetObjectType_LIFT, true
	case "CALENDAR":
		return BACnetObjectType_CALENDAR, true
	case "COMMAND":
		return BACnetObjectType_COMMAND, true
	case "DEVICE":
		return BACnetObjectType_DEVICE, true
	case "EVENT_ENROLLMENT":
		return BACnetObjectType_EVENT_ENROLLMENT, true
	}
	return 0, false
}

func BACnetObjectTypeKnows(value uint16) bool {
	for _, typeValue := range BACnetObjectTypeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetObjectType(structType any) BACnetObjectType {
	castFunc := func(typ any) BACnetObjectType {
		if sBACnetObjectType, ok := typ.(BACnetObjectType); ok {
			return sBACnetObjectType
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetObjectType) GetLengthInBits(ctx context.Context) uint16 {
	return 10
}

func (m BACnetObjectType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetObjectTypeParse(ctx context.Context, theBytes []byte) (BACnetObjectType, error) {
	return BACnetObjectTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetObjectTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetObjectType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint16("BACnetObjectType", 10)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetObjectType")
	}
	if enum, ok := BACnetObjectTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetObjectType")
		return BACnetObjectType(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetObjectType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetObjectType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint16("BACnetObjectType", 10, uint16(uint16(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e BACnetObjectType) GetValue() uint16 {
	return uint16(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetObjectType) PLC4XEnumName() string {
	switch e {
	case BACnetObjectType_ANALOG_INPUT:
		return "ANALOG_INPUT"
	case BACnetObjectType_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetObjectType_ANALOG_OUTPUT:
		return "ANALOG_OUTPUT"
	case BACnetObjectType_FILE:
		return "FILE"
	case BACnetObjectType_GROUP:
		return "GROUP"
	case BACnetObjectType_LOOP:
		return "LOOP"
	case BACnetObjectType_MULTI_STATE_INPUT:
		return "MULTI_STATE_INPUT"
	case BACnetObjectType_MULTI_STATE_OUTPUT:
		return "MULTI_STATE_OUTPUT"
	case BACnetObjectType_NOTIFICATION_CLASS:
		return "NOTIFICATION_CLASS"
	case BACnetObjectType_PROGRAM:
		return "PROGRAM"
	case BACnetObjectType_SCHEDULE:
		return "SCHEDULE"
	case BACnetObjectType_AVERAGING:
		return "AVERAGING"
	case BACnetObjectType_MULTI_STATE_VALUE:
		return "MULTI_STATE_VALUE"
	case BACnetObjectType_ANALOG_VALUE:
		return "ANALOG_VALUE"
	case BACnetObjectType_TREND_LOG:
		return "TREND_LOG"
	case BACnetObjectType_LIFE_SAFETY_POINT:
		return "LIFE_SAFETY_POINT"
	case BACnetObjectType_LIFE_SAFETY_ZONE:
		return "LIFE_SAFETY_ZONE"
	case BACnetObjectType_ACCUMULATOR:
		return "ACCUMULATOR"
	case BACnetObjectType_PULSE_CONVERTER:
		return "PULSE_CONVERTER"
	case BACnetObjectType_EVENT_LOG:
		return "EVENT_LOG"
	case BACnetObjectType_GLOBAL_GROUP:
		return "GLOBAL_GROUP"
	case BACnetObjectType_TREND_LOG_MULTIPLE:
		return "TREND_LOG_MULTIPLE"
	case BACnetObjectType_LOAD_CONTROL:
		return "LOAD_CONTROL"
	case BACnetObjectType_STRUCTURED_VIEW:
		return "STRUCTURED_VIEW"
	case BACnetObjectType_BINARY_INPUT:
		return "BINARY_INPUT"
	case BACnetObjectType_ACCESS_DOOR:
		return "ACCESS_DOOR"
	case BACnetObjectType_TIMER:
		return "TIMER"
	case BACnetObjectType_ACCESS_CREDENTIAL:
		return "ACCESS_CREDENTIAL"
	case BACnetObjectType_ACCESS_POINT:
		return "ACCESS_POINT"
	case BACnetObjectType_ACCESS_RIGHTS:
		return "ACCESS_RIGHTS"
	case BACnetObjectType_ACCESS_USER:
		return "ACCESS_USER"
	case BACnetObjectType_ACCESS_ZONE:
		return "ACCESS_ZONE"
	case BACnetObjectType_CREDENTIAL_DATA_INPUT:
		return "CREDENTIAL_DATA_INPUT"
	case BACnetObjectType_NETWORK_SECURITY:
		return "NETWORK_SECURITY"
	case BACnetObjectType_BITSTRING_VALUE:
		return "BITSTRING_VALUE"
	case BACnetObjectType_BINARY_OUTPUT:
		return "BINARY_OUTPUT"
	case BACnetObjectType_CHARACTERSTRING_VALUE:
		return "CHARACTERSTRING_VALUE"
	case BACnetObjectType_DATEPATTERN_VALUE:
		return "DATEPATTERN_VALUE"
	case BACnetObjectType_DATE_VALUE:
		return "DATE_VALUE"
	case BACnetObjectType_DATETIMEPATTERN_VALUE:
		return "DATETIMEPATTERN_VALUE"
	case BACnetObjectType_DATETIME_VALUE:
		return "DATETIME_VALUE"
	case BACnetObjectType_INTEGER_VALUE:
		return "INTEGER_VALUE"
	case BACnetObjectType_LARGE_ANALOG_VALUE:
		return "LARGE_ANALOG_VALUE"
	case BACnetObjectType_OCTETSTRING_VALUE:
		return "OCTETSTRING_VALUE"
	case BACnetObjectType_POSITIVE_INTEGER_VALUE:
		return "POSITIVE_INTEGER_VALUE"
	case BACnetObjectType_TIMEPATTERN_VALUE:
		return "TIMEPATTERN_VALUE"
	case BACnetObjectType_BINARY_VALUE:
		return "BINARY_VALUE"
	case BACnetObjectType_TIME_VALUE:
		return "TIME_VALUE"
	case BACnetObjectType_NOTIFICATION_FORWARDER:
		return "NOTIFICATION_FORWARDER"
	case BACnetObjectType_ALERT_ENROLLMENT:
		return "ALERT_ENROLLMENT"
	case BACnetObjectType_CHANNEL:
		return "CHANNEL"
	case BACnetObjectType_LIGHTING_OUTPUT:
		return "LIGHTING_OUTPUT"
	case BACnetObjectType_BINARY_LIGHTING_OUTPUT:
		return "BINARY_LIGHTING_OUTPUT"
	case BACnetObjectType_NETWORK_PORT:
		return "NETWORK_PORT"
	case BACnetObjectType_ELEVATOR_GROUP:
		return "ELEVATOR_GROUP"
	case BACnetObjectType_ESCALATOR:
		return "ESCALATOR"
	case BACnetObjectType_LIFT:
		return "LIFT"
	case BACnetObjectType_CALENDAR:
		return "CALENDAR"
	case BACnetObjectType_COMMAND:
		return "COMMAND"
	case BACnetObjectType_DEVICE:
		return "DEVICE"
	case BACnetObjectType_EVENT_ENROLLMENT:
		return "EVENT_ENROLLMENT"
	}
	return fmt.Sprintf("Unknown(%v)", uint16(e))
}

func (e BACnetObjectType) String() string {
	return e.PLC4XEnumName()
}
