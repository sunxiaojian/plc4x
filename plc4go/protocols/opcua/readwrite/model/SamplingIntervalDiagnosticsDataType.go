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

// SamplingIntervalDiagnosticsDataType is the corresponding interface of SamplingIntervalDiagnosticsDataType
type SamplingIntervalDiagnosticsDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetSamplingInterval returns SamplingInterval (property field)
	GetSamplingInterval() float64
	// GetMonitoredItemCount returns MonitoredItemCount (property field)
	GetMonitoredItemCount() uint32
	// GetMaxMonitoredItemCount returns MaxMonitoredItemCount (property field)
	GetMaxMonitoredItemCount() uint32
	// GetDisabledMonitoredItemCount returns DisabledMonitoredItemCount (property field)
	GetDisabledMonitoredItemCount() uint32
}

// SamplingIntervalDiagnosticsDataTypeExactly can be used when we want exactly this type and not a type which fulfills SamplingIntervalDiagnosticsDataType.
// This is useful for switch cases.
type SamplingIntervalDiagnosticsDataTypeExactly interface {
	SamplingIntervalDiagnosticsDataType
	isSamplingIntervalDiagnosticsDataType() bool
}

// _SamplingIntervalDiagnosticsDataType is the data-structure of this message
type _SamplingIntervalDiagnosticsDataType struct {
	*_ExtensionObjectDefinition
	SamplingInterval           float64
	MonitoredItemCount         uint32
	MaxMonitoredItemCount      uint32
	DisabledMonitoredItemCount uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SamplingIntervalDiagnosticsDataType) GetIdentifier() string {
	return "858"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SamplingIntervalDiagnosticsDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_SamplingIntervalDiagnosticsDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SamplingIntervalDiagnosticsDataType) GetSamplingInterval() float64 {
	return m.SamplingInterval
}

func (m *_SamplingIntervalDiagnosticsDataType) GetMonitoredItemCount() uint32 {
	return m.MonitoredItemCount
}

func (m *_SamplingIntervalDiagnosticsDataType) GetMaxMonitoredItemCount() uint32 {
	return m.MaxMonitoredItemCount
}

func (m *_SamplingIntervalDiagnosticsDataType) GetDisabledMonitoredItemCount() uint32 {
	return m.DisabledMonitoredItemCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSamplingIntervalDiagnosticsDataType factory function for _SamplingIntervalDiagnosticsDataType
func NewSamplingIntervalDiagnosticsDataType(samplingInterval float64, monitoredItemCount uint32, maxMonitoredItemCount uint32, disabledMonitoredItemCount uint32) *_SamplingIntervalDiagnosticsDataType {
	_result := &_SamplingIntervalDiagnosticsDataType{
		SamplingInterval:           samplingInterval,
		MonitoredItemCount:         monitoredItemCount,
		MaxMonitoredItemCount:      maxMonitoredItemCount,
		DisabledMonitoredItemCount: disabledMonitoredItemCount,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSamplingIntervalDiagnosticsDataType(structType any) SamplingIntervalDiagnosticsDataType {
	if casted, ok := structType.(SamplingIntervalDiagnosticsDataType); ok {
		return casted
	}
	if casted, ok := structType.(*SamplingIntervalDiagnosticsDataType); ok {
		return *casted
	}
	return nil
}

func (m *_SamplingIntervalDiagnosticsDataType) GetTypeName() string {
	return "SamplingIntervalDiagnosticsDataType"
}

func (m *_SamplingIntervalDiagnosticsDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (samplingInterval)
	lengthInBits += 64

	// Simple field (monitoredItemCount)
	lengthInBits += 32

	// Simple field (maxMonitoredItemCount)
	lengthInBits += 32

	// Simple field (disabledMonitoredItemCount)
	lengthInBits += 32

	return lengthInBits
}

func (m *_SamplingIntervalDiagnosticsDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SamplingIntervalDiagnosticsDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (SamplingIntervalDiagnosticsDataType, error) {
	return SamplingIntervalDiagnosticsDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func SamplingIntervalDiagnosticsDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (SamplingIntervalDiagnosticsDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SamplingIntervalDiagnosticsDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SamplingIntervalDiagnosticsDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (samplingInterval)
	_samplingInterval, _samplingIntervalErr := /*TODO: migrate me*/ readBuffer.ReadFloat64("samplingInterval", 64)
	if _samplingIntervalErr != nil {
		return nil, errors.Wrap(_samplingIntervalErr, "Error parsing 'samplingInterval' field of SamplingIntervalDiagnosticsDataType")
	}
	samplingInterval := _samplingInterval

	// Simple Field (monitoredItemCount)
	_monitoredItemCount, _monitoredItemCountErr := /*TODO: migrate me*/ readBuffer.ReadUint32("monitoredItemCount", 32)
	if _monitoredItemCountErr != nil {
		return nil, errors.Wrap(_monitoredItemCountErr, "Error parsing 'monitoredItemCount' field of SamplingIntervalDiagnosticsDataType")
	}
	monitoredItemCount := _monitoredItemCount

	// Simple Field (maxMonitoredItemCount)
	_maxMonitoredItemCount, _maxMonitoredItemCountErr := /*TODO: migrate me*/ readBuffer.ReadUint32("maxMonitoredItemCount", 32)
	if _maxMonitoredItemCountErr != nil {
		return nil, errors.Wrap(_maxMonitoredItemCountErr, "Error parsing 'maxMonitoredItemCount' field of SamplingIntervalDiagnosticsDataType")
	}
	maxMonitoredItemCount := _maxMonitoredItemCount

	// Simple Field (disabledMonitoredItemCount)
	_disabledMonitoredItemCount, _disabledMonitoredItemCountErr := /*TODO: migrate me*/ readBuffer.ReadUint32("disabledMonitoredItemCount", 32)
	if _disabledMonitoredItemCountErr != nil {
		return nil, errors.Wrap(_disabledMonitoredItemCountErr, "Error parsing 'disabledMonitoredItemCount' field of SamplingIntervalDiagnosticsDataType")
	}
	disabledMonitoredItemCount := _disabledMonitoredItemCount

	if closeErr := readBuffer.CloseContext("SamplingIntervalDiagnosticsDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SamplingIntervalDiagnosticsDataType")
	}

	// Create a partially initialized instance
	_child := &_SamplingIntervalDiagnosticsDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		SamplingInterval:           samplingInterval,
		MonitoredItemCount:         monitoredItemCount,
		MaxMonitoredItemCount:      maxMonitoredItemCount,
		DisabledMonitoredItemCount: disabledMonitoredItemCount,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_SamplingIntervalDiagnosticsDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SamplingIntervalDiagnosticsDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SamplingIntervalDiagnosticsDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SamplingIntervalDiagnosticsDataType")
		}

		// Simple Field (samplingInterval)
		samplingInterval := float64(m.GetSamplingInterval())
		_samplingIntervalErr := /*TODO: migrate me*/ writeBuffer.WriteFloat64("samplingInterval", 64, (samplingInterval))
		if _samplingIntervalErr != nil {
			return errors.Wrap(_samplingIntervalErr, "Error serializing 'samplingInterval' field")
		}

		// Simple Field (monitoredItemCount)
		monitoredItemCount := uint32(m.GetMonitoredItemCount())
		_monitoredItemCountErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("monitoredItemCount", 32, uint32((monitoredItemCount)))
		if _monitoredItemCountErr != nil {
			return errors.Wrap(_monitoredItemCountErr, "Error serializing 'monitoredItemCount' field")
		}

		// Simple Field (maxMonitoredItemCount)
		maxMonitoredItemCount := uint32(m.GetMaxMonitoredItemCount())
		_maxMonitoredItemCountErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("maxMonitoredItemCount", 32, uint32((maxMonitoredItemCount)))
		if _maxMonitoredItemCountErr != nil {
			return errors.Wrap(_maxMonitoredItemCountErr, "Error serializing 'maxMonitoredItemCount' field")
		}

		// Simple Field (disabledMonitoredItemCount)
		disabledMonitoredItemCount := uint32(m.GetDisabledMonitoredItemCount())
		_disabledMonitoredItemCountErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("disabledMonitoredItemCount", 32, uint32((disabledMonitoredItemCount)))
		if _disabledMonitoredItemCountErr != nil {
			return errors.Wrap(_disabledMonitoredItemCountErr, "Error serializing 'disabledMonitoredItemCount' field")
		}

		if popErr := writeBuffer.PopContext("SamplingIntervalDiagnosticsDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SamplingIntervalDiagnosticsDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SamplingIntervalDiagnosticsDataType) isSamplingIntervalDiagnosticsDataType() bool {
	return true
}

func (m *_SamplingIntervalDiagnosticsDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
