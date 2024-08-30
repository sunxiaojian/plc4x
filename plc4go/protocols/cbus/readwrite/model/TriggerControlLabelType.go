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

// TriggerControlLabelType is an enum
type TriggerControlLabelType uint8

type ITriggerControlLabelType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	TriggerControlLabelType_TEXT_LABEL             TriggerControlLabelType = 0
	TriggerControlLabelType_PREDEFINED_ICON        TriggerControlLabelType = 1
	TriggerControlLabelType_LOAD_DYNAMIC_ICON      TriggerControlLabelType = 2
	TriggerControlLabelType_SET_PREFERRED_LANGUAGE TriggerControlLabelType = 3
)

var TriggerControlLabelTypeValues []TriggerControlLabelType

func init() {
	_ = errors.New
	TriggerControlLabelTypeValues = []TriggerControlLabelType{
		TriggerControlLabelType_TEXT_LABEL,
		TriggerControlLabelType_PREDEFINED_ICON,
		TriggerControlLabelType_LOAD_DYNAMIC_ICON,
		TriggerControlLabelType_SET_PREFERRED_LANGUAGE,
	}
}

func TriggerControlLabelTypeByValue(value uint8) (enum TriggerControlLabelType, ok bool) {
	switch value {
	case 0:
		return TriggerControlLabelType_TEXT_LABEL, true
	case 1:
		return TriggerControlLabelType_PREDEFINED_ICON, true
	case 2:
		return TriggerControlLabelType_LOAD_DYNAMIC_ICON, true
	case 3:
		return TriggerControlLabelType_SET_PREFERRED_LANGUAGE, true
	}
	return 0, false
}

func TriggerControlLabelTypeByName(value string) (enum TriggerControlLabelType, ok bool) {
	switch value {
	case "TEXT_LABEL":
		return TriggerControlLabelType_TEXT_LABEL, true
	case "PREDEFINED_ICON":
		return TriggerControlLabelType_PREDEFINED_ICON, true
	case "LOAD_DYNAMIC_ICON":
		return TriggerControlLabelType_LOAD_DYNAMIC_ICON, true
	case "SET_PREFERRED_LANGUAGE":
		return TriggerControlLabelType_SET_PREFERRED_LANGUAGE, true
	}
	return 0, false
}

func TriggerControlLabelTypeKnows(value uint8) bool {
	for _, typeValue := range TriggerControlLabelTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastTriggerControlLabelType(structType any) TriggerControlLabelType {
	castFunc := func(typ any) TriggerControlLabelType {
		if sTriggerControlLabelType, ok := typ.(TriggerControlLabelType); ok {
			return sTriggerControlLabelType
		}
		return 0
	}
	return castFunc(structType)
}

func (m TriggerControlLabelType) GetLengthInBits(ctx context.Context) uint16 {
	return 2
}

func (m TriggerControlLabelType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TriggerControlLabelTypeParse(ctx context.Context, theBytes []byte) (TriggerControlLabelType, error) {
	return TriggerControlLabelTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TriggerControlLabelTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TriggerControlLabelType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("TriggerControlLabelType", 2)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TriggerControlLabelType")
	}
	if enum, ok := TriggerControlLabelTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for TriggerControlLabelType")
		return TriggerControlLabelType(val), nil
	} else {
		return enum, nil
	}
}

func (e TriggerControlLabelType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e TriggerControlLabelType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("TriggerControlLabelType", 2, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TriggerControlLabelType) PLC4XEnumName() string {
	switch e {
	case TriggerControlLabelType_TEXT_LABEL:
		return "TEXT_LABEL"
	case TriggerControlLabelType_PREDEFINED_ICON:
		return "PREDEFINED_ICON"
	case TriggerControlLabelType_LOAD_DYNAMIC_ICON:
		return "LOAD_DYNAMIC_ICON"
	case TriggerControlLabelType_SET_PREFERRED_LANGUAGE:
		return "SET_PREFERRED_LANGUAGE"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e TriggerControlLabelType) String() string {
	return e.PLC4XEnumName()
}
