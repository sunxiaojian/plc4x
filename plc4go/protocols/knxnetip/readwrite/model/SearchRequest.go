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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// SearchRequest is the corresponding interface of SearchRequest
type SearchRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	KnxNetIpMessage
	// GetHpaiIDiscoveryEndpoint returns HpaiIDiscoveryEndpoint (property field)
	GetHpaiIDiscoveryEndpoint() HPAIDiscoveryEndpoint
	// IsSearchRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSearchRequest()
	// CreateBuilder creates a SearchRequestBuilder
	CreateSearchRequestBuilder() SearchRequestBuilder
}

// _SearchRequest is the data-structure of this message
type _SearchRequest struct {
	KnxNetIpMessageContract
	HpaiIDiscoveryEndpoint HPAIDiscoveryEndpoint
}

var _ SearchRequest = (*_SearchRequest)(nil)
var _ KnxNetIpMessageRequirements = (*_SearchRequest)(nil)

// NewSearchRequest factory function for _SearchRequest
func NewSearchRequest(hpaiIDiscoveryEndpoint HPAIDiscoveryEndpoint) *_SearchRequest {
	if hpaiIDiscoveryEndpoint == nil {
		panic("hpaiIDiscoveryEndpoint of type HPAIDiscoveryEndpoint for SearchRequest must not be nil")
	}
	_result := &_SearchRequest{
		KnxNetIpMessageContract: NewKnxNetIpMessage(),
		HpaiIDiscoveryEndpoint:  hpaiIDiscoveryEndpoint,
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SearchRequestBuilder is a builder for SearchRequest
type SearchRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(hpaiIDiscoveryEndpoint HPAIDiscoveryEndpoint) SearchRequestBuilder
	// WithHpaiIDiscoveryEndpoint adds HpaiIDiscoveryEndpoint (property field)
	WithHpaiIDiscoveryEndpoint(HPAIDiscoveryEndpoint) SearchRequestBuilder
	// WithHpaiIDiscoveryEndpointBuilder adds HpaiIDiscoveryEndpoint (property field) which is build by the builder
	WithHpaiIDiscoveryEndpointBuilder(func(HPAIDiscoveryEndpointBuilder) HPAIDiscoveryEndpointBuilder) SearchRequestBuilder
	// Build builds the SearchRequest or returns an error if something is wrong
	Build() (SearchRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SearchRequest
}

// NewSearchRequestBuilder() creates a SearchRequestBuilder
func NewSearchRequestBuilder() SearchRequestBuilder {
	return &_SearchRequestBuilder{_SearchRequest: new(_SearchRequest)}
}

type _SearchRequestBuilder struct {
	*_SearchRequest

	err *utils.MultiError
}

var _ (SearchRequestBuilder) = (*_SearchRequestBuilder)(nil)

func (m *_SearchRequestBuilder) WithMandatoryFields(hpaiIDiscoveryEndpoint HPAIDiscoveryEndpoint) SearchRequestBuilder {
	return m.WithHpaiIDiscoveryEndpoint(hpaiIDiscoveryEndpoint)
}

func (m *_SearchRequestBuilder) WithHpaiIDiscoveryEndpoint(hpaiIDiscoveryEndpoint HPAIDiscoveryEndpoint) SearchRequestBuilder {
	m.HpaiIDiscoveryEndpoint = hpaiIDiscoveryEndpoint
	return m
}

func (m *_SearchRequestBuilder) WithHpaiIDiscoveryEndpointBuilder(builderSupplier func(HPAIDiscoveryEndpointBuilder) HPAIDiscoveryEndpointBuilder) SearchRequestBuilder {
	builder := builderSupplier(m.HpaiIDiscoveryEndpoint.CreateHPAIDiscoveryEndpointBuilder())
	var err error
	m.HpaiIDiscoveryEndpoint, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "HPAIDiscoveryEndpointBuilder failed"))
	}
	return m
}

func (m *_SearchRequestBuilder) Build() (SearchRequest, error) {
	if m.HpaiIDiscoveryEndpoint == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'hpaiIDiscoveryEndpoint' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._SearchRequest.deepCopy(), nil
}

func (m *_SearchRequestBuilder) MustBuild() SearchRequest {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_SearchRequestBuilder) DeepCopy() any {
	return m.CreateSearchRequestBuilder()
}

// CreateSearchRequestBuilder creates a SearchRequestBuilder
func (m *_SearchRequest) CreateSearchRequestBuilder() SearchRequestBuilder {
	if m == nil {
		return NewSearchRequestBuilder()
	}
	return &_SearchRequestBuilder{_SearchRequest: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SearchRequest) GetMsgType() uint16 {
	return 0x0201
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SearchRequest) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SearchRequest) GetHpaiIDiscoveryEndpoint() HPAIDiscoveryEndpoint {
	return m.HpaiIDiscoveryEndpoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSearchRequest(structType any) SearchRequest {
	if casted, ok := structType.(SearchRequest); ok {
		return casted
	}
	if casted, ok := structType.(*SearchRequest); ok {
		return *casted
	}
	return nil
}

func (m *_SearchRequest) GetTypeName() string {
	return "SearchRequest"
}

func (m *_SearchRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (hpaiIDiscoveryEndpoint)
	lengthInBits += m.HpaiIDiscoveryEndpoint.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SearchRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SearchRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage) (__searchRequest SearchRequest, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SearchRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SearchRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	hpaiIDiscoveryEndpoint, err := ReadSimpleField[HPAIDiscoveryEndpoint](ctx, "hpaiIDiscoveryEndpoint", ReadComplex[HPAIDiscoveryEndpoint](HPAIDiscoveryEndpointParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hpaiIDiscoveryEndpoint' field"))
	}
	m.HpaiIDiscoveryEndpoint = hpaiIDiscoveryEndpoint

	if closeErr := readBuffer.CloseContext("SearchRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SearchRequest")
	}

	return m, nil
}

func (m *_SearchRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SearchRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SearchRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SearchRequest")
		}

		if err := WriteSimpleField[HPAIDiscoveryEndpoint](ctx, "hpaiIDiscoveryEndpoint", m.GetHpaiIDiscoveryEndpoint(), WriteComplex[HPAIDiscoveryEndpoint](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'hpaiIDiscoveryEndpoint' field")
		}

		if popErr := writeBuffer.PopContext("SearchRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SearchRequest")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SearchRequest) IsSearchRequest() {}

func (m *_SearchRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SearchRequest) deepCopy() *_SearchRequest {
	if m == nil {
		return nil
	}
	_SearchRequestCopy := &_SearchRequest{
		m.KnxNetIpMessageContract.(*_KnxNetIpMessage).deepCopy(),
		m.HpaiIDiscoveryEndpoint.DeepCopy().(HPAIDiscoveryEndpoint),
	}
	m.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = m
	return _SearchRequestCopy
}

func (m *_SearchRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
