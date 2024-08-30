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

// OpcuaNodeIdServicesVariableTopics is an enum
type OpcuaNodeIdServicesVariableTopics int32

type IOpcuaNodeIdServicesVariableTopics interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableTopics_Topics_FindAlias_InputArguments  OpcuaNodeIdServicesVariableTopics = 23495
	OpcuaNodeIdServicesVariableTopics_Topics_FindAlias_OutputArguments OpcuaNodeIdServicesVariableTopics = 23496
	OpcuaNodeIdServicesVariableTopics_Topics_LastChange                OpcuaNodeIdServicesVariableTopics = 32856
)

var OpcuaNodeIdServicesVariableTopicsValues []OpcuaNodeIdServicesVariableTopics

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableTopicsValues = []OpcuaNodeIdServicesVariableTopics{
		OpcuaNodeIdServicesVariableTopics_Topics_FindAlias_InputArguments,
		OpcuaNodeIdServicesVariableTopics_Topics_FindAlias_OutputArguments,
		OpcuaNodeIdServicesVariableTopics_Topics_LastChange,
	}
}

func OpcuaNodeIdServicesVariableTopicsByValue(value int32) (enum OpcuaNodeIdServicesVariableTopics, ok bool) {
	switch value {
	case 23495:
		return OpcuaNodeIdServicesVariableTopics_Topics_FindAlias_InputArguments, true
	case 23496:
		return OpcuaNodeIdServicesVariableTopics_Topics_FindAlias_OutputArguments, true
	case 32856:
		return OpcuaNodeIdServicesVariableTopics_Topics_LastChange, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTopicsByName(value string) (enum OpcuaNodeIdServicesVariableTopics, ok bool) {
	switch value {
	case "Topics_FindAlias_InputArguments":
		return OpcuaNodeIdServicesVariableTopics_Topics_FindAlias_InputArguments, true
	case "Topics_FindAlias_OutputArguments":
		return OpcuaNodeIdServicesVariableTopics_Topics_FindAlias_OutputArguments, true
	case "Topics_LastChange":
		return OpcuaNodeIdServicesVariableTopics_Topics_LastChange, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableTopicsKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableTopicsValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableTopics(structType any) OpcuaNodeIdServicesVariableTopics {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableTopics {
		if sOpcuaNodeIdServicesVariableTopics, ok := typ.(OpcuaNodeIdServicesVariableTopics); ok {
			return sOpcuaNodeIdServicesVariableTopics
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableTopics) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableTopics) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableTopicsParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableTopics, error) {
	return OpcuaNodeIdServicesVariableTopicsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableTopicsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableTopics, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableTopics", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableTopics")
	}
	if enum, ok := OpcuaNodeIdServicesVariableTopicsByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableTopics")
		return OpcuaNodeIdServicesVariableTopics(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableTopics) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableTopics) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableTopics", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableTopics) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableTopics_Topics_FindAlias_InputArguments:
		return "Topics_FindAlias_InputArguments"
	case OpcuaNodeIdServicesVariableTopics_Topics_FindAlias_OutputArguments:
		return "Topics_FindAlias_OutputArguments"
	case OpcuaNodeIdServicesVariableTopics_Topics_LastChange:
		return "Topics_LastChange"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableTopics) String() string {
	return e.PLC4XEnumName()
}
