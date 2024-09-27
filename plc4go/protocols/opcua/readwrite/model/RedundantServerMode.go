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

// RedundantServerMode is an enum
type RedundantServerMode uint32

type IRedundantServerMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	RedundantServerMode_redundantServerModePrimaryWithBackup RedundantServerMode = 0
	RedundantServerMode_redundantServerModePrimaryOnly       RedundantServerMode = 1
	RedundantServerMode_redundantServerModeBackupReady       RedundantServerMode = 2
	RedundantServerMode_redundantServerModeBackupNotReady    RedundantServerMode = 3
)

var RedundantServerModeValues []RedundantServerMode

func init() {
	_ = errors.New
	RedundantServerModeValues = []RedundantServerMode{
		RedundantServerMode_redundantServerModePrimaryWithBackup,
		RedundantServerMode_redundantServerModePrimaryOnly,
		RedundantServerMode_redundantServerModeBackupReady,
		RedundantServerMode_redundantServerModeBackupNotReady,
	}
}

func RedundantServerModeByValue(value uint32) (enum RedundantServerMode, ok bool) {
	switch value {
	case 0:
		return RedundantServerMode_redundantServerModePrimaryWithBackup, true
	case 1:
		return RedundantServerMode_redundantServerModePrimaryOnly, true
	case 2:
		return RedundantServerMode_redundantServerModeBackupReady, true
	case 3:
		return RedundantServerMode_redundantServerModeBackupNotReady, true
	}
	return 0, false
}

func RedundantServerModeByName(value string) (enum RedundantServerMode, ok bool) {
	switch value {
	case "redundantServerModePrimaryWithBackup":
		return RedundantServerMode_redundantServerModePrimaryWithBackup, true
	case "redundantServerModePrimaryOnly":
		return RedundantServerMode_redundantServerModePrimaryOnly, true
	case "redundantServerModeBackupReady":
		return RedundantServerMode_redundantServerModeBackupReady, true
	case "redundantServerModeBackupNotReady":
		return RedundantServerMode_redundantServerModeBackupNotReady, true
	}
	return 0, false
}

func RedundantServerModeKnows(value uint32) bool {
	for _, typeValue := range RedundantServerModeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastRedundantServerMode(structType any) RedundantServerMode {
	castFunc := func(typ any) RedundantServerMode {
		if sRedundantServerMode, ok := typ.(RedundantServerMode); ok {
			return sRedundantServerMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m RedundantServerMode) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m RedundantServerMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RedundantServerModeParse(ctx context.Context, theBytes []byte) (RedundantServerMode, error) {
	return RedundantServerModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func RedundantServerModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (RedundantServerMode, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("RedundantServerMode", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading RedundantServerMode")
	}
	if enum, ok := RedundantServerModeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for RedundantServerMode")
		return RedundantServerMode(val), nil
	} else {
		return enum, nil
	}
}

func (e RedundantServerMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e RedundantServerMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("RedundantServerMode", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e RedundantServerMode) GetValue() uint32 {
	return uint32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e RedundantServerMode) PLC4XEnumName() string {
	switch e {
	case RedundantServerMode_redundantServerModePrimaryWithBackup:
		return "redundantServerModePrimaryWithBackup"
	case RedundantServerMode_redundantServerModePrimaryOnly:
		return "redundantServerModePrimaryOnly"
	case RedundantServerMode_redundantServerModeBackupReady:
		return "redundantServerModeBackupReady"
	case RedundantServerMode_redundantServerModeBackupNotReady:
		return "redundantServerModeBackupNotReady"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e RedundantServerMode) String() string {
	return e.PLC4XEnumName()
}
