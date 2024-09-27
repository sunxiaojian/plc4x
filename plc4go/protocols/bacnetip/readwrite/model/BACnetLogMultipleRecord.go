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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetLogMultipleRecord is the corresponding interface of BACnetLogMultipleRecord
type BACnetLogMultipleRecord interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetDateTimeEnclosed
	// GetLogData returns LogData (property field)
	GetLogData() BACnetLogData
	// IsBACnetLogMultipleRecord is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetLogMultipleRecord()
	// CreateBuilder creates a BACnetLogMultipleRecordBuilder
	CreateBACnetLogMultipleRecordBuilder() BACnetLogMultipleRecordBuilder
}

// _BACnetLogMultipleRecord is the data-structure of this message
type _BACnetLogMultipleRecord struct {
	Timestamp BACnetDateTimeEnclosed
	LogData   BACnetLogData
}

var _ BACnetLogMultipleRecord = (*_BACnetLogMultipleRecord)(nil)

// NewBACnetLogMultipleRecord factory function for _BACnetLogMultipleRecord
func NewBACnetLogMultipleRecord(timestamp BACnetDateTimeEnclosed, logData BACnetLogData) *_BACnetLogMultipleRecord {
	if timestamp == nil {
		panic("timestamp of type BACnetDateTimeEnclosed for BACnetLogMultipleRecord must not be nil")
	}
	if logData == nil {
		panic("logData of type BACnetLogData for BACnetLogMultipleRecord must not be nil")
	}
	return &_BACnetLogMultipleRecord{Timestamp: timestamp, LogData: logData}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetLogMultipleRecordBuilder is a builder for BACnetLogMultipleRecord
type BACnetLogMultipleRecordBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timestamp BACnetDateTimeEnclosed, logData BACnetLogData) BACnetLogMultipleRecordBuilder
	// WithTimestamp adds Timestamp (property field)
	WithTimestamp(BACnetDateTimeEnclosed) BACnetLogMultipleRecordBuilder
	// WithTimestampBuilder adds Timestamp (property field) which is build by the builder
	WithTimestampBuilder(func(BACnetDateTimeEnclosedBuilder) BACnetDateTimeEnclosedBuilder) BACnetLogMultipleRecordBuilder
	// WithLogData adds LogData (property field)
	WithLogData(BACnetLogData) BACnetLogMultipleRecordBuilder
	// Build builds the BACnetLogMultipleRecord or returns an error if something is wrong
	Build() (BACnetLogMultipleRecord, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetLogMultipleRecord
}

// NewBACnetLogMultipleRecordBuilder() creates a BACnetLogMultipleRecordBuilder
func NewBACnetLogMultipleRecordBuilder() BACnetLogMultipleRecordBuilder {
	return &_BACnetLogMultipleRecordBuilder{_BACnetLogMultipleRecord: new(_BACnetLogMultipleRecord)}
}

type _BACnetLogMultipleRecordBuilder struct {
	*_BACnetLogMultipleRecord

	err *utils.MultiError
}

var _ (BACnetLogMultipleRecordBuilder) = (*_BACnetLogMultipleRecordBuilder)(nil)

func (m *_BACnetLogMultipleRecordBuilder) WithMandatoryFields(timestamp BACnetDateTimeEnclosed, logData BACnetLogData) BACnetLogMultipleRecordBuilder {
	return m.WithTimestamp(timestamp).WithLogData(logData)
}

func (m *_BACnetLogMultipleRecordBuilder) WithTimestamp(timestamp BACnetDateTimeEnclosed) BACnetLogMultipleRecordBuilder {
	m.Timestamp = timestamp
	return m
}

func (m *_BACnetLogMultipleRecordBuilder) WithTimestampBuilder(builderSupplier func(BACnetDateTimeEnclosedBuilder) BACnetDateTimeEnclosedBuilder) BACnetLogMultipleRecordBuilder {
	builder := builderSupplier(m.Timestamp.CreateBACnetDateTimeEnclosedBuilder())
	var err error
	m.Timestamp, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetDateTimeEnclosedBuilder failed"))
	}
	return m
}

func (m *_BACnetLogMultipleRecordBuilder) WithLogData(logData BACnetLogData) BACnetLogMultipleRecordBuilder {
	m.LogData = logData
	return m
}

func (m *_BACnetLogMultipleRecordBuilder) Build() (BACnetLogMultipleRecord, error) {
	if m.Timestamp == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'timestamp' not set"))
	}
	if m.LogData == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'logData' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetLogMultipleRecord.deepCopy(), nil
}

func (m *_BACnetLogMultipleRecordBuilder) MustBuild() BACnetLogMultipleRecord {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetLogMultipleRecordBuilder) DeepCopy() any {
	return m.CreateBACnetLogMultipleRecordBuilder()
}

// CreateBACnetLogMultipleRecordBuilder creates a BACnetLogMultipleRecordBuilder
func (m *_BACnetLogMultipleRecord) CreateBACnetLogMultipleRecordBuilder() BACnetLogMultipleRecordBuilder {
	if m == nil {
		return NewBACnetLogMultipleRecordBuilder()
	}
	return &_BACnetLogMultipleRecordBuilder{_BACnetLogMultipleRecord: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogMultipleRecord) GetTimestamp() BACnetDateTimeEnclosed {
	return m.Timestamp
}

func (m *_BACnetLogMultipleRecord) GetLogData() BACnetLogData {
	return m.LogData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetLogMultipleRecord(structType any) BACnetLogMultipleRecord {
	if casted, ok := structType.(BACnetLogMultipleRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogMultipleRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogMultipleRecord) GetTypeName() string {
	return "BACnetLogMultipleRecord"
}

func (m *_BACnetLogMultipleRecord) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits(ctx)

	// Simple field (logData)
	lengthInBits += m.LogData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogMultipleRecord) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLogMultipleRecordParse(ctx context.Context, theBytes []byte) (BACnetLogMultipleRecord, error) {
	return BACnetLogMultipleRecordParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLogMultipleRecordParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLogMultipleRecord, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLogMultipleRecord, error) {
		return BACnetLogMultipleRecordParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetLogMultipleRecordParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLogMultipleRecord, error) {
	v, err := (&_BACnetLogMultipleRecord{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetLogMultipleRecord) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetLogMultipleRecord BACnetLogMultipleRecord, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogMultipleRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogMultipleRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timestamp, err := ReadSimpleField[BACnetDateTimeEnclosed](ctx, "timestamp", ReadComplex[BACnetDateTimeEnclosed](BACnetDateTimeEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}
	m.Timestamp = timestamp

	logData, err := ReadSimpleField[BACnetLogData](ctx, "logData", ReadComplex[BACnetLogData](BACnetLogDataParseWithBufferProducer[BACnetLogData]((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logData' field"))
	}
	m.LogData = logData

	if closeErr := readBuffer.CloseContext("BACnetLogMultipleRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogMultipleRecord")
	}

	return m, nil
}

func (m *_BACnetLogMultipleRecord) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogMultipleRecord) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLogMultipleRecord"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLogMultipleRecord")
	}

	if err := WriteSimpleField[BACnetDateTimeEnclosed](ctx, "timestamp", m.GetTimestamp(), WriteComplex[BACnetDateTimeEnclosed](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'timestamp' field")
	}

	if err := WriteSimpleField[BACnetLogData](ctx, "logData", m.GetLogData(), WriteComplex[BACnetLogData](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'logData' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLogMultipleRecord"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLogMultipleRecord")
	}
	return nil
}

func (m *_BACnetLogMultipleRecord) IsBACnetLogMultipleRecord() {}

func (m *_BACnetLogMultipleRecord) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetLogMultipleRecord) deepCopy() *_BACnetLogMultipleRecord {
	if m == nil {
		return nil
	}
	_BACnetLogMultipleRecordCopy := &_BACnetLogMultipleRecord{
		m.Timestamp.DeepCopy().(BACnetDateTimeEnclosed),
		m.LogData.DeepCopy().(BACnetLogData),
	}
	return _BACnetLogMultipleRecordCopy
}

func (m *_BACnetLogMultipleRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
