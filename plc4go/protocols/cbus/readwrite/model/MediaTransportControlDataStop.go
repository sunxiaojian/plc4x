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

// MediaTransportControlDataStop is the corresponding interface of MediaTransportControlDataStop
type MediaTransportControlDataStop interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MediaTransportControlData
	// IsMediaTransportControlDataStop is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMediaTransportControlDataStop()
	// CreateBuilder creates a MediaTransportControlDataStopBuilder
	CreateMediaTransportControlDataStopBuilder() MediaTransportControlDataStopBuilder
}

// _MediaTransportControlDataStop is the data-structure of this message
type _MediaTransportControlDataStop struct {
	MediaTransportControlDataContract
}

var _ MediaTransportControlDataStop = (*_MediaTransportControlDataStop)(nil)
var _ MediaTransportControlDataRequirements = (*_MediaTransportControlDataStop)(nil)

// NewMediaTransportControlDataStop factory function for _MediaTransportControlDataStop
func NewMediaTransportControlDataStop(commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataStop {
	_result := &_MediaTransportControlDataStop{
		MediaTransportControlDataContract: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MediaTransportControlDataStopBuilder is a builder for MediaTransportControlDataStop
type MediaTransportControlDataStopBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() MediaTransportControlDataStopBuilder
	// Build builds the MediaTransportControlDataStop or returns an error if something is wrong
	Build() (MediaTransportControlDataStop, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MediaTransportControlDataStop
}

// NewMediaTransportControlDataStopBuilder() creates a MediaTransportControlDataStopBuilder
func NewMediaTransportControlDataStopBuilder() MediaTransportControlDataStopBuilder {
	return &_MediaTransportControlDataStopBuilder{_MediaTransportControlDataStop: new(_MediaTransportControlDataStop)}
}

type _MediaTransportControlDataStopBuilder struct {
	*_MediaTransportControlDataStop

	err *utils.MultiError
}

var _ (MediaTransportControlDataStopBuilder) = (*_MediaTransportControlDataStopBuilder)(nil)

func (m *_MediaTransportControlDataStopBuilder) WithMandatoryFields() MediaTransportControlDataStopBuilder {
	return m
}

func (m *_MediaTransportControlDataStopBuilder) Build() (MediaTransportControlDataStop, error) {
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._MediaTransportControlDataStop.deepCopy(), nil
}

func (m *_MediaTransportControlDataStopBuilder) MustBuild() MediaTransportControlDataStop {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_MediaTransportControlDataStopBuilder) DeepCopy() any {
	return m.CreateMediaTransportControlDataStopBuilder()
}

// CreateMediaTransportControlDataStopBuilder creates a MediaTransportControlDataStopBuilder
func (m *_MediaTransportControlDataStop) CreateMediaTransportControlDataStopBuilder() MediaTransportControlDataStopBuilder {
	if m == nil {
		return NewMediaTransportControlDataStopBuilder()
	}
	return &_MediaTransportControlDataStopBuilder{_MediaTransportControlDataStop: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataStop) GetParent() MediaTransportControlDataContract {
	return m.MediaTransportControlDataContract
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataStop(structType any) MediaTransportControlDataStop {
	if casted, ok := structType.(MediaTransportControlDataStop); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataStop); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataStop) GetTypeName() string {
	return "MediaTransportControlDataStop"
}

func (m *_MediaTransportControlDataStop) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MediaTransportControlDataContract.(*_MediaTransportControlData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_MediaTransportControlDataStop) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MediaTransportControlDataStop) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MediaTransportControlData) (__mediaTransportControlDataStop MediaTransportControlDataStop, err error) {
	m.MediaTransportControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataStop"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataStop")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataStop"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataStop")
	}

	return m, nil
}

func (m *_MediaTransportControlDataStop) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataStop) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataStop"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataStop")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataStop"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataStop")
		}
		return nil
	}
	return m.MediaTransportControlDataContract.(*_MediaTransportControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataStop) IsMediaTransportControlDataStop() {}

func (m *_MediaTransportControlDataStop) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MediaTransportControlDataStop) deepCopy() *_MediaTransportControlDataStop {
	if m == nil {
		return nil
	}
	_MediaTransportControlDataStopCopy := &_MediaTransportControlDataStop{
		m.MediaTransportControlDataContract.(*_MediaTransportControlData).deepCopy(),
	}
	m.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = m
	return _MediaTransportControlDataStopCopy
}

func (m *_MediaTransportControlDataStop) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
