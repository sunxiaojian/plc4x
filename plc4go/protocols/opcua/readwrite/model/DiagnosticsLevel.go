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

// DiagnosticsLevel is an enum
type DiagnosticsLevel uint32

type IDiagnosticsLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	DiagnosticsLevel_diagnosticsLevelBasic    DiagnosticsLevel = 0
	DiagnosticsLevel_diagnosticsLevelAdvanced DiagnosticsLevel = 1
	DiagnosticsLevel_diagnosticsLevelInfo     DiagnosticsLevel = 2
	DiagnosticsLevel_diagnosticsLevelLog      DiagnosticsLevel = 3
	DiagnosticsLevel_diagnosticsLevelDebug    DiagnosticsLevel = 4
)

var DiagnosticsLevelValues []DiagnosticsLevel

func init() {
	_ = errors.New
	DiagnosticsLevelValues = []DiagnosticsLevel{
		DiagnosticsLevel_diagnosticsLevelBasic,
		DiagnosticsLevel_diagnosticsLevelAdvanced,
		DiagnosticsLevel_diagnosticsLevelInfo,
		DiagnosticsLevel_diagnosticsLevelLog,
		DiagnosticsLevel_diagnosticsLevelDebug,
	}
}

func DiagnosticsLevelByValue(value uint32) (enum DiagnosticsLevel, ok bool) {
	switch value {
	case 0:
		return DiagnosticsLevel_diagnosticsLevelBasic, true
	case 1:
		return DiagnosticsLevel_diagnosticsLevelAdvanced, true
	case 2:
		return DiagnosticsLevel_diagnosticsLevelInfo, true
	case 3:
		return DiagnosticsLevel_diagnosticsLevelLog, true
	case 4:
		return DiagnosticsLevel_diagnosticsLevelDebug, true
	}
	return 0, false
}

func DiagnosticsLevelByName(value string) (enum DiagnosticsLevel, ok bool) {
	switch value {
	case "diagnosticsLevelBasic":
		return DiagnosticsLevel_diagnosticsLevelBasic, true
	case "diagnosticsLevelAdvanced":
		return DiagnosticsLevel_diagnosticsLevelAdvanced, true
	case "diagnosticsLevelInfo":
		return DiagnosticsLevel_diagnosticsLevelInfo, true
	case "diagnosticsLevelLog":
		return DiagnosticsLevel_diagnosticsLevelLog, true
	case "diagnosticsLevelDebug":
		return DiagnosticsLevel_diagnosticsLevelDebug, true
	}
	return 0, false
}

func DiagnosticsLevelKnows(value uint32) bool {
	for _, typeValue := range DiagnosticsLevelValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastDiagnosticsLevel(structType any) DiagnosticsLevel {
	castFunc := func(typ any) DiagnosticsLevel {
		if sDiagnosticsLevel, ok := typ.(DiagnosticsLevel); ok {
			return sDiagnosticsLevel
		}
		return 0
	}
	return castFunc(structType)
}

func (m DiagnosticsLevel) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m DiagnosticsLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DiagnosticsLevelParse(ctx context.Context, theBytes []byte) (DiagnosticsLevel, error) {
	return DiagnosticsLevelParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DiagnosticsLevelParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DiagnosticsLevel, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("DiagnosticsLevel", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading DiagnosticsLevel")
	}
	if enum, ok := DiagnosticsLevelByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for DiagnosticsLevel")
		return DiagnosticsLevel(val), nil
	} else {
		return enum, nil
	}
}

func (e DiagnosticsLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e DiagnosticsLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("DiagnosticsLevel", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e DiagnosticsLevel) PLC4XEnumName() string {
	switch e {
	case DiagnosticsLevel_diagnosticsLevelBasic:
		return "diagnosticsLevelBasic"
	case DiagnosticsLevel_diagnosticsLevelAdvanced:
		return "diagnosticsLevelAdvanced"
	case DiagnosticsLevel_diagnosticsLevelInfo:
		return "diagnosticsLevelInfo"
	case DiagnosticsLevel_diagnosticsLevelLog:
		return "diagnosticsLevelLog"
	case DiagnosticsLevel_diagnosticsLevelDebug:
		return "diagnosticsLevelDebug"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e DiagnosticsLevel) String() string {
	return e.PLC4XEnumName()
}
