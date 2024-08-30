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

// OpcuaNodeIdServicesVariableCartesian is an enum
type OpcuaNodeIdServicesVariableCartesian int32

type IOpcuaNodeIdServicesVariableCartesian interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableCartesian_CartesianCoordinatesType_LengthUnit OpcuaNodeIdServicesVariableCartesian = 18773
)

var OpcuaNodeIdServicesVariableCartesianValues []OpcuaNodeIdServicesVariableCartesian

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableCartesianValues = []OpcuaNodeIdServicesVariableCartesian{
		OpcuaNodeIdServicesVariableCartesian_CartesianCoordinatesType_LengthUnit,
	}
}

func OpcuaNodeIdServicesVariableCartesianByValue(value int32) (enum OpcuaNodeIdServicesVariableCartesian, ok bool) {
	switch value {
	case 18773:
		return OpcuaNodeIdServicesVariableCartesian_CartesianCoordinatesType_LengthUnit, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableCartesianByName(value string) (enum OpcuaNodeIdServicesVariableCartesian, ok bool) {
	switch value {
	case "CartesianCoordinatesType_LengthUnit":
		return OpcuaNodeIdServicesVariableCartesian_CartesianCoordinatesType_LengthUnit, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableCartesianKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableCartesianValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableCartesian(structType any) OpcuaNodeIdServicesVariableCartesian {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableCartesian {
		if sOpcuaNodeIdServicesVariableCartesian, ok := typ.(OpcuaNodeIdServicesVariableCartesian); ok {
			return sOpcuaNodeIdServicesVariableCartesian
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableCartesian) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableCartesian) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableCartesianParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableCartesian, error) {
	return OpcuaNodeIdServicesVariableCartesianParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableCartesianParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableCartesian, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableCartesian", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableCartesian")
	}
	if enum, ok := OpcuaNodeIdServicesVariableCartesianByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableCartesian")
		return OpcuaNodeIdServicesVariableCartesian(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableCartesian) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableCartesian) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableCartesian", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableCartesian) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableCartesian_CartesianCoordinatesType_LengthUnit:
		return "CartesianCoordinatesType_LengthUnit"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableCartesian) String() string {
	return e.PLC4XEnumName()
}
