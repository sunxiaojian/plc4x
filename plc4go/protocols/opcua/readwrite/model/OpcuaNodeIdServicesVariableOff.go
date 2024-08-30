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

// OpcuaNodeIdServicesVariableOff is an enum
type OpcuaNodeIdServicesVariableOff int32

type IOpcuaNodeIdServicesVariableOff interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableOff_OffNormalAlarmType_NormalState OpcuaNodeIdServicesVariableOff = 11158
)

var OpcuaNodeIdServicesVariableOffValues []OpcuaNodeIdServicesVariableOff

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableOffValues = []OpcuaNodeIdServicesVariableOff{
		OpcuaNodeIdServicesVariableOff_OffNormalAlarmType_NormalState,
	}
}

func OpcuaNodeIdServicesVariableOffByValue(value int32) (enum OpcuaNodeIdServicesVariableOff, ok bool) {
	switch value {
	case 11158:
		return OpcuaNodeIdServicesVariableOff_OffNormalAlarmType_NormalState, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableOffByName(value string) (enum OpcuaNodeIdServicesVariableOff, ok bool) {
	switch value {
	case "OffNormalAlarmType_NormalState":
		return OpcuaNodeIdServicesVariableOff_OffNormalAlarmType_NormalState, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableOffKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableOffValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableOff(structType any) OpcuaNodeIdServicesVariableOff {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableOff {
		if sOpcuaNodeIdServicesVariableOff, ok := typ.(OpcuaNodeIdServicesVariableOff); ok {
			return sOpcuaNodeIdServicesVariableOff
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableOff) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableOff) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableOffParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableOff, error) {
	return OpcuaNodeIdServicesVariableOffParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableOffParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableOff, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableOff", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableOff")
	}
	if enum, ok := OpcuaNodeIdServicesVariableOffByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableOff")
		return OpcuaNodeIdServicesVariableOff(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableOff) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableOff) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableOff", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableOff) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableOff_OffNormalAlarmType_NormalState:
		return "OffNormalAlarmType_NormalState"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableOff) String() string {
	return e.PLC4XEnumName()
}
