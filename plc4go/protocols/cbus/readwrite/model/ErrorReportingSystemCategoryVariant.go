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

// ErrorReportingSystemCategoryVariant is an enum
type ErrorReportingSystemCategoryVariant uint8

type IErrorReportingSystemCategoryVariant interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	ErrorReportingSystemCategoryVariant_RESERVED_0 ErrorReportingSystemCategoryVariant = 0x0
	ErrorReportingSystemCategoryVariant_RESERVED_1 ErrorReportingSystemCategoryVariant = 0x1
	ErrorReportingSystemCategoryVariant_RESERVED_2 ErrorReportingSystemCategoryVariant = 0x2
	ErrorReportingSystemCategoryVariant_RESERVED_3 ErrorReportingSystemCategoryVariant = 0x3
)

var ErrorReportingSystemCategoryVariantValues []ErrorReportingSystemCategoryVariant

func init() {
	_ = errors.New
	ErrorReportingSystemCategoryVariantValues = []ErrorReportingSystemCategoryVariant{
		ErrorReportingSystemCategoryVariant_RESERVED_0,
		ErrorReportingSystemCategoryVariant_RESERVED_1,
		ErrorReportingSystemCategoryVariant_RESERVED_2,
		ErrorReportingSystemCategoryVariant_RESERVED_3,
	}
}

func ErrorReportingSystemCategoryVariantByValue(value uint8) (enum ErrorReportingSystemCategoryVariant, ok bool) {
	switch value {
	case 0x0:
		return ErrorReportingSystemCategoryVariant_RESERVED_0, true
	case 0x1:
		return ErrorReportingSystemCategoryVariant_RESERVED_1, true
	case 0x2:
		return ErrorReportingSystemCategoryVariant_RESERVED_2, true
	case 0x3:
		return ErrorReportingSystemCategoryVariant_RESERVED_3, true
	}
	return 0, false
}

func ErrorReportingSystemCategoryVariantByName(value string) (enum ErrorReportingSystemCategoryVariant, ok bool) {
	switch value {
	case "RESERVED_0":
		return ErrorReportingSystemCategoryVariant_RESERVED_0, true
	case "RESERVED_1":
		return ErrorReportingSystemCategoryVariant_RESERVED_1, true
	case "RESERVED_2":
		return ErrorReportingSystemCategoryVariant_RESERVED_2, true
	case "RESERVED_3":
		return ErrorReportingSystemCategoryVariant_RESERVED_3, true
	}
	return 0, false
}

func ErrorReportingSystemCategoryVariantKnows(value uint8) bool {
	for _, typeValue := range ErrorReportingSystemCategoryVariantValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastErrorReportingSystemCategoryVariant(structType any) ErrorReportingSystemCategoryVariant {
	castFunc := func(typ any) ErrorReportingSystemCategoryVariant {
		if sErrorReportingSystemCategoryVariant, ok := typ.(ErrorReportingSystemCategoryVariant); ok {
			return sErrorReportingSystemCategoryVariant
		}
		return 0
	}
	return castFunc(structType)
}

func (m ErrorReportingSystemCategoryVariant) GetLengthInBits(ctx context.Context) uint16 {
	return 2
}

func (m ErrorReportingSystemCategoryVariant) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorReportingSystemCategoryVariantParse(ctx context.Context, theBytes []byte) (ErrorReportingSystemCategoryVariant, error) {
	return ErrorReportingSystemCategoryVariantParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ErrorReportingSystemCategoryVariantParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ErrorReportingSystemCategoryVariant, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("ErrorReportingSystemCategoryVariant", 2)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ErrorReportingSystemCategoryVariant")
	}
	if enum, ok := ErrorReportingSystemCategoryVariantByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for ErrorReportingSystemCategoryVariant")
		return ErrorReportingSystemCategoryVariant(val), nil
	} else {
		return enum, nil
	}
}

func (e ErrorReportingSystemCategoryVariant) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ErrorReportingSystemCategoryVariant) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("ErrorReportingSystemCategoryVariant", 2, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ErrorReportingSystemCategoryVariant) PLC4XEnumName() string {
	switch e {
	case ErrorReportingSystemCategoryVariant_RESERVED_0:
		return "RESERVED_0"
	case ErrorReportingSystemCategoryVariant_RESERVED_1:
		return "RESERVED_1"
	case ErrorReportingSystemCategoryVariant_RESERVED_2:
		return "RESERVED_2"
	case ErrorReportingSystemCategoryVariant_RESERVED_3:
		return "RESERVED_3"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e ErrorReportingSystemCategoryVariant) String() string {
	return e.PLC4XEnumName()
}
