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

// BACnetAuthorizationExemption is an enum
type BACnetAuthorizationExemption uint8

type IBACnetAuthorizationExemption interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetAuthorizationExemption_PASSBACK                 BACnetAuthorizationExemption = 0
	BACnetAuthorizationExemption_OCCUPANCY_CHECK          BACnetAuthorizationExemption = 1
	BACnetAuthorizationExemption_ACCESS_RIGHTS            BACnetAuthorizationExemption = 2
	BACnetAuthorizationExemption_LOCKOUT                  BACnetAuthorizationExemption = 3
	BACnetAuthorizationExemption_DENY                     BACnetAuthorizationExemption = 4
	BACnetAuthorizationExemption_VERIFICATION             BACnetAuthorizationExemption = 5
	BACnetAuthorizationExemption_AUTHORIZATION_DELAY      BACnetAuthorizationExemption = 6
	BACnetAuthorizationExemption_VENDOR_PROPRIETARY_VALUE BACnetAuthorizationExemption = 0xFF
)

var BACnetAuthorizationExemptionValues []BACnetAuthorizationExemption

func init() {
	_ = errors.New
	BACnetAuthorizationExemptionValues = []BACnetAuthorizationExemption{
		BACnetAuthorizationExemption_PASSBACK,
		BACnetAuthorizationExemption_OCCUPANCY_CHECK,
		BACnetAuthorizationExemption_ACCESS_RIGHTS,
		BACnetAuthorizationExemption_LOCKOUT,
		BACnetAuthorizationExemption_DENY,
		BACnetAuthorizationExemption_VERIFICATION,
		BACnetAuthorizationExemption_AUTHORIZATION_DELAY,
		BACnetAuthorizationExemption_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetAuthorizationExemptionByValue(value uint8) (enum BACnetAuthorizationExemption, ok bool) {
	switch value {
	case 0:
		return BACnetAuthorizationExemption_PASSBACK, true
	case 0xFF:
		return BACnetAuthorizationExemption_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetAuthorizationExemption_OCCUPANCY_CHECK, true
	case 2:
		return BACnetAuthorizationExemption_ACCESS_RIGHTS, true
	case 3:
		return BACnetAuthorizationExemption_LOCKOUT, true
	case 4:
		return BACnetAuthorizationExemption_DENY, true
	case 5:
		return BACnetAuthorizationExemption_VERIFICATION, true
	case 6:
		return BACnetAuthorizationExemption_AUTHORIZATION_DELAY, true
	}
	return 0, false
}

func BACnetAuthorizationExemptionByName(value string) (enum BACnetAuthorizationExemption, ok bool) {
	switch value {
	case "PASSBACK":
		return BACnetAuthorizationExemption_PASSBACK, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetAuthorizationExemption_VENDOR_PROPRIETARY_VALUE, true
	case "OCCUPANCY_CHECK":
		return BACnetAuthorizationExemption_OCCUPANCY_CHECK, true
	case "ACCESS_RIGHTS":
		return BACnetAuthorizationExemption_ACCESS_RIGHTS, true
	case "LOCKOUT":
		return BACnetAuthorizationExemption_LOCKOUT, true
	case "DENY":
		return BACnetAuthorizationExemption_DENY, true
	case "VERIFICATION":
		return BACnetAuthorizationExemption_VERIFICATION, true
	case "AUTHORIZATION_DELAY":
		return BACnetAuthorizationExemption_AUTHORIZATION_DELAY, true
	}
	return 0, false
}

func BACnetAuthorizationExemptionKnows(value uint8) bool {
	for _, typeValue := range BACnetAuthorizationExemptionValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetAuthorizationExemption(structType any) BACnetAuthorizationExemption {
	castFunc := func(typ any) BACnetAuthorizationExemption {
		if sBACnetAuthorizationExemption, ok := typ.(BACnetAuthorizationExemption); ok {
			return sBACnetAuthorizationExemption
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAuthorizationExemption) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetAuthorizationExemption) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAuthorizationExemptionParse(ctx context.Context, theBytes []byte) (BACnetAuthorizationExemption, error) {
	return BACnetAuthorizationExemptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAuthorizationExemptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAuthorizationExemption, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("BACnetAuthorizationExemption", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetAuthorizationExemption")
	}
	if enum, ok := BACnetAuthorizationExemptionByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetAuthorizationExemption")
		return BACnetAuthorizationExemption(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetAuthorizationExemption) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetAuthorizationExemption) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("BACnetAuthorizationExemption", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAuthorizationExemption) PLC4XEnumName() string {
	switch e {
	case BACnetAuthorizationExemption_PASSBACK:
		return "PASSBACK"
	case BACnetAuthorizationExemption_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetAuthorizationExemption_OCCUPANCY_CHECK:
		return "OCCUPANCY_CHECK"
	case BACnetAuthorizationExemption_ACCESS_RIGHTS:
		return "ACCESS_RIGHTS"
	case BACnetAuthorizationExemption_LOCKOUT:
		return "LOCKOUT"
	case BACnetAuthorizationExemption_DENY:
		return "DENY"
	case BACnetAuthorizationExemption_VERIFICATION:
		return "VERIFICATION"
	case BACnetAuthorizationExemption_AUTHORIZATION_DELAY:
		return "AUTHORIZATION_DELAY"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetAuthorizationExemption) String() string {
	return e.PLC4XEnumName()
}
