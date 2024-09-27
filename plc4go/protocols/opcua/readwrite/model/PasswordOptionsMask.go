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

// PasswordOptionsMask is an enum
type PasswordOptionsMask uint32

type IPasswordOptionsMask interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	PasswordOptionsMask_passwordOptionsMaskNone                         PasswordOptionsMask = 0
	PasswordOptionsMask_passwordOptionsMaskSupportInitialPasswordChange PasswordOptionsMask = 1
	PasswordOptionsMask_passwordOptionsMaskSupportDisableUser           PasswordOptionsMask = 2
	PasswordOptionsMask_passwordOptionsMaskSupportDisableDeleteForUser  PasswordOptionsMask = 4
	PasswordOptionsMask_passwordOptionsMaskSupportNoChangeForUser       PasswordOptionsMask = 8
	PasswordOptionsMask_passwordOptionsMaskSupportDescriptionForUser    PasswordOptionsMask = 16
	PasswordOptionsMask_passwordOptionsMaskRequiresUpperCaseCharacters  PasswordOptionsMask = 32
	PasswordOptionsMask_passwordOptionsMaskRequiresLowerCaseCharacters  PasswordOptionsMask = 64
	PasswordOptionsMask_passwordOptionsMaskRequiresDigitCharacters      PasswordOptionsMask = 128
	PasswordOptionsMask_passwordOptionsMaskRequiresSpecialCharacters    PasswordOptionsMask = 256
)

var PasswordOptionsMaskValues []PasswordOptionsMask

func init() {
	_ = errors.New
	PasswordOptionsMaskValues = []PasswordOptionsMask{
		PasswordOptionsMask_passwordOptionsMaskNone,
		PasswordOptionsMask_passwordOptionsMaskSupportInitialPasswordChange,
		PasswordOptionsMask_passwordOptionsMaskSupportDisableUser,
		PasswordOptionsMask_passwordOptionsMaskSupportDisableDeleteForUser,
		PasswordOptionsMask_passwordOptionsMaskSupportNoChangeForUser,
		PasswordOptionsMask_passwordOptionsMaskSupportDescriptionForUser,
		PasswordOptionsMask_passwordOptionsMaskRequiresUpperCaseCharacters,
		PasswordOptionsMask_passwordOptionsMaskRequiresLowerCaseCharacters,
		PasswordOptionsMask_passwordOptionsMaskRequiresDigitCharacters,
		PasswordOptionsMask_passwordOptionsMaskRequiresSpecialCharacters,
	}
}

func PasswordOptionsMaskByValue(value uint32) (enum PasswordOptionsMask, ok bool) {
	switch value {
	case 0:
		return PasswordOptionsMask_passwordOptionsMaskNone, true
	case 1:
		return PasswordOptionsMask_passwordOptionsMaskSupportInitialPasswordChange, true
	case 128:
		return PasswordOptionsMask_passwordOptionsMaskRequiresDigitCharacters, true
	case 16:
		return PasswordOptionsMask_passwordOptionsMaskSupportDescriptionForUser, true
	case 2:
		return PasswordOptionsMask_passwordOptionsMaskSupportDisableUser, true
	case 256:
		return PasswordOptionsMask_passwordOptionsMaskRequiresSpecialCharacters, true
	case 32:
		return PasswordOptionsMask_passwordOptionsMaskRequiresUpperCaseCharacters, true
	case 4:
		return PasswordOptionsMask_passwordOptionsMaskSupportDisableDeleteForUser, true
	case 64:
		return PasswordOptionsMask_passwordOptionsMaskRequiresLowerCaseCharacters, true
	case 8:
		return PasswordOptionsMask_passwordOptionsMaskSupportNoChangeForUser, true
	}
	return 0, false
}

func PasswordOptionsMaskByName(value string) (enum PasswordOptionsMask, ok bool) {
	switch value {
	case "passwordOptionsMaskNone":
		return PasswordOptionsMask_passwordOptionsMaskNone, true
	case "passwordOptionsMaskSupportInitialPasswordChange":
		return PasswordOptionsMask_passwordOptionsMaskSupportInitialPasswordChange, true
	case "passwordOptionsMaskRequiresDigitCharacters":
		return PasswordOptionsMask_passwordOptionsMaskRequiresDigitCharacters, true
	case "passwordOptionsMaskSupportDescriptionForUser":
		return PasswordOptionsMask_passwordOptionsMaskSupportDescriptionForUser, true
	case "passwordOptionsMaskSupportDisableUser":
		return PasswordOptionsMask_passwordOptionsMaskSupportDisableUser, true
	case "passwordOptionsMaskRequiresSpecialCharacters":
		return PasswordOptionsMask_passwordOptionsMaskRequiresSpecialCharacters, true
	case "passwordOptionsMaskRequiresUpperCaseCharacters":
		return PasswordOptionsMask_passwordOptionsMaskRequiresUpperCaseCharacters, true
	case "passwordOptionsMaskSupportDisableDeleteForUser":
		return PasswordOptionsMask_passwordOptionsMaskSupportDisableDeleteForUser, true
	case "passwordOptionsMaskRequiresLowerCaseCharacters":
		return PasswordOptionsMask_passwordOptionsMaskRequiresLowerCaseCharacters, true
	case "passwordOptionsMaskSupportNoChangeForUser":
		return PasswordOptionsMask_passwordOptionsMaskSupportNoChangeForUser, true
	}
	return 0, false
}

func PasswordOptionsMaskKnows(value uint32) bool {
	for _, typeValue := range PasswordOptionsMaskValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastPasswordOptionsMask(structType any) PasswordOptionsMask {
	castFunc := func(typ any) PasswordOptionsMask {
		if sPasswordOptionsMask, ok := typ.(PasswordOptionsMask); ok {
			return sPasswordOptionsMask
		}
		return 0
	}
	return castFunc(structType)
}

func (m PasswordOptionsMask) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m PasswordOptionsMask) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PasswordOptionsMaskParse(ctx context.Context, theBytes []byte) (PasswordOptionsMask, error) {
	return PasswordOptionsMaskParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func PasswordOptionsMaskParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (PasswordOptionsMask, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("PasswordOptionsMask", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading PasswordOptionsMask")
	}
	if enum, ok := PasswordOptionsMaskByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for PasswordOptionsMask")
		return PasswordOptionsMask(val), nil
	} else {
		return enum, nil
	}
}

func (e PasswordOptionsMask) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e PasswordOptionsMask) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("PasswordOptionsMask", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e PasswordOptionsMask) GetValue() uint32 {
	return uint32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e PasswordOptionsMask) PLC4XEnumName() string {
	switch e {
	case PasswordOptionsMask_passwordOptionsMaskNone:
		return "passwordOptionsMaskNone"
	case PasswordOptionsMask_passwordOptionsMaskSupportInitialPasswordChange:
		return "passwordOptionsMaskSupportInitialPasswordChange"
	case PasswordOptionsMask_passwordOptionsMaskRequiresDigitCharacters:
		return "passwordOptionsMaskRequiresDigitCharacters"
	case PasswordOptionsMask_passwordOptionsMaskSupportDescriptionForUser:
		return "passwordOptionsMaskSupportDescriptionForUser"
	case PasswordOptionsMask_passwordOptionsMaskSupportDisableUser:
		return "passwordOptionsMaskSupportDisableUser"
	case PasswordOptionsMask_passwordOptionsMaskRequiresSpecialCharacters:
		return "passwordOptionsMaskRequiresSpecialCharacters"
	case PasswordOptionsMask_passwordOptionsMaskRequiresUpperCaseCharacters:
		return "passwordOptionsMaskRequiresUpperCaseCharacters"
	case PasswordOptionsMask_passwordOptionsMaskSupportDisableDeleteForUser:
		return "passwordOptionsMaskSupportDisableDeleteForUser"
	case PasswordOptionsMask_passwordOptionsMaskRequiresLowerCaseCharacters:
		return "passwordOptionsMaskRequiresLowerCaseCharacters"
	case PasswordOptionsMask_passwordOptionsMaskSupportNoChangeForUser:
		return "passwordOptionsMaskSupportNoChangeForUser"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e PasswordOptionsMask) String() string {
	return e.PLC4XEnumName()
}
