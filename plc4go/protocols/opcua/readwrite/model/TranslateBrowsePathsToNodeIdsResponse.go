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

// TranslateBrowsePathsToNodeIdsResponse is the corresponding interface of TranslateBrowsePathsToNodeIdsResponse
type TranslateBrowsePathsToNodeIdsResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfResults returns NoOfResults (property field)
	GetNoOfResults() int32
	// GetResults returns Results (property field)
	GetResults() []ExtensionObjectDefinition
	// GetNoOfDiagnosticInfos returns NoOfDiagnosticInfos (property field)
	GetNoOfDiagnosticInfos() int32
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
	// IsTranslateBrowsePathsToNodeIdsResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTranslateBrowsePathsToNodeIdsResponse()
	// CreateBuilder creates a TranslateBrowsePathsToNodeIdsResponseBuilder
	CreateTranslateBrowsePathsToNodeIdsResponseBuilder() TranslateBrowsePathsToNodeIdsResponseBuilder
}

// _TranslateBrowsePathsToNodeIdsResponse is the data-structure of this message
type _TranslateBrowsePathsToNodeIdsResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader      ExtensionObjectDefinition
	NoOfResults         int32
	Results             []ExtensionObjectDefinition
	NoOfDiagnosticInfos int32
	DiagnosticInfos     []DiagnosticInfo
}

var _ TranslateBrowsePathsToNodeIdsResponse = (*_TranslateBrowsePathsToNodeIdsResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_TranslateBrowsePathsToNodeIdsResponse)(nil)

// NewTranslateBrowsePathsToNodeIdsResponse factory function for _TranslateBrowsePathsToNodeIdsResponse
func NewTranslateBrowsePathsToNodeIdsResponse(responseHeader ExtensionObjectDefinition, noOfResults int32, results []ExtensionObjectDefinition, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) *_TranslateBrowsePathsToNodeIdsResponse {
	if responseHeader == nil {
		panic("responseHeader of type ExtensionObjectDefinition for TranslateBrowsePathsToNodeIdsResponse must not be nil")
	}
	_result := &_TranslateBrowsePathsToNodeIdsResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		NoOfResults:                       noOfResults,
		Results:                           results,
		NoOfDiagnosticInfos:               noOfDiagnosticInfos,
		DiagnosticInfos:                   diagnosticInfos,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TranslateBrowsePathsToNodeIdsResponseBuilder is a builder for TranslateBrowsePathsToNodeIdsResponse
type TranslateBrowsePathsToNodeIdsResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(responseHeader ExtensionObjectDefinition, noOfResults int32, results []ExtensionObjectDefinition, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) TranslateBrowsePathsToNodeIdsResponseBuilder
	// WithResponseHeader adds ResponseHeader (property field)
	WithResponseHeader(ExtensionObjectDefinition) TranslateBrowsePathsToNodeIdsResponseBuilder
	// WithNoOfResults adds NoOfResults (property field)
	WithNoOfResults(int32) TranslateBrowsePathsToNodeIdsResponseBuilder
	// WithResults adds Results (property field)
	WithResults(...ExtensionObjectDefinition) TranslateBrowsePathsToNodeIdsResponseBuilder
	// WithNoOfDiagnosticInfos adds NoOfDiagnosticInfos (property field)
	WithNoOfDiagnosticInfos(int32) TranslateBrowsePathsToNodeIdsResponseBuilder
	// WithDiagnosticInfos adds DiagnosticInfos (property field)
	WithDiagnosticInfos(...DiagnosticInfo) TranslateBrowsePathsToNodeIdsResponseBuilder
	// Build builds the TranslateBrowsePathsToNodeIdsResponse or returns an error if something is wrong
	Build() (TranslateBrowsePathsToNodeIdsResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TranslateBrowsePathsToNodeIdsResponse
}

// NewTranslateBrowsePathsToNodeIdsResponseBuilder() creates a TranslateBrowsePathsToNodeIdsResponseBuilder
func NewTranslateBrowsePathsToNodeIdsResponseBuilder() TranslateBrowsePathsToNodeIdsResponseBuilder {
	return &_TranslateBrowsePathsToNodeIdsResponseBuilder{_TranslateBrowsePathsToNodeIdsResponse: new(_TranslateBrowsePathsToNodeIdsResponse)}
}

type _TranslateBrowsePathsToNodeIdsResponseBuilder struct {
	*_TranslateBrowsePathsToNodeIdsResponse

	err *utils.MultiError
}

var _ (TranslateBrowsePathsToNodeIdsResponseBuilder) = (*_TranslateBrowsePathsToNodeIdsResponseBuilder)(nil)

func (m *_TranslateBrowsePathsToNodeIdsResponseBuilder) WithMandatoryFields(responseHeader ExtensionObjectDefinition, noOfResults int32, results []ExtensionObjectDefinition, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) TranslateBrowsePathsToNodeIdsResponseBuilder {
	return m.WithResponseHeader(responseHeader).WithNoOfResults(noOfResults).WithResults(results...).WithNoOfDiagnosticInfos(noOfDiagnosticInfos).WithDiagnosticInfos(diagnosticInfos...)
}

func (m *_TranslateBrowsePathsToNodeIdsResponseBuilder) WithResponseHeader(responseHeader ExtensionObjectDefinition) TranslateBrowsePathsToNodeIdsResponseBuilder {
	m.ResponseHeader = responseHeader
	return m
}

func (m *_TranslateBrowsePathsToNodeIdsResponseBuilder) WithNoOfResults(noOfResults int32) TranslateBrowsePathsToNodeIdsResponseBuilder {
	m.NoOfResults = noOfResults
	return m
}

func (m *_TranslateBrowsePathsToNodeIdsResponseBuilder) WithResults(results ...ExtensionObjectDefinition) TranslateBrowsePathsToNodeIdsResponseBuilder {
	m.Results = results
	return m
}

func (m *_TranslateBrowsePathsToNodeIdsResponseBuilder) WithNoOfDiagnosticInfos(noOfDiagnosticInfos int32) TranslateBrowsePathsToNodeIdsResponseBuilder {
	m.NoOfDiagnosticInfos = noOfDiagnosticInfos
	return m
}

func (m *_TranslateBrowsePathsToNodeIdsResponseBuilder) WithDiagnosticInfos(diagnosticInfos ...DiagnosticInfo) TranslateBrowsePathsToNodeIdsResponseBuilder {
	m.DiagnosticInfos = diagnosticInfos
	return m
}

func (m *_TranslateBrowsePathsToNodeIdsResponseBuilder) Build() (TranslateBrowsePathsToNodeIdsResponse, error) {
	if m.ResponseHeader == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'responseHeader' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._TranslateBrowsePathsToNodeIdsResponse.deepCopy(), nil
}

func (m *_TranslateBrowsePathsToNodeIdsResponseBuilder) MustBuild() TranslateBrowsePathsToNodeIdsResponse {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_TranslateBrowsePathsToNodeIdsResponseBuilder) DeepCopy() any {
	return m.CreateTranslateBrowsePathsToNodeIdsResponseBuilder()
}

// CreateTranslateBrowsePathsToNodeIdsResponseBuilder creates a TranslateBrowsePathsToNodeIdsResponseBuilder
func (m *_TranslateBrowsePathsToNodeIdsResponse) CreateTranslateBrowsePathsToNodeIdsResponseBuilder() TranslateBrowsePathsToNodeIdsResponseBuilder {
	if m == nil {
		return NewTranslateBrowsePathsToNodeIdsResponseBuilder()
	}
	return &_TranslateBrowsePathsToNodeIdsResponseBuilder{_TranslateBrowsePathsToNodeIdsResponse: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetIdentifier() string {
	return "557"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetNoOfResults() int32 {
	return m.NoOfResults
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetResults() []ExtensionObjectDefinition {
	return m.Results
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetNoOfDiagnosticInfos() int32 {
	return m.NoOfDiagnosticInfos
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastTranslateBrowsePathsToNodeIdsResponse(structType any) TranslateBrowsePathsToNodeIdsResponse {
	if casted, ok := structType.(TranslateBrowsePathsToNodeIdsResponse); ok {
		return casted
	}
	if casted, ok := structType.(*TranslateBrowsePathsToNodeIdsResponse); ok {
		return *casted
	}
	return nil
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetTypeName() string {
	return "TranslateBrowsePathsToNodeIdsResponse"
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (noOfResults)
	lengthInBits += 32

	// Array field
	if len(m.Results) > 0 {
		for _curItem, element := range m.Results {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Results), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.DiagnosticInfos) > 0 {
		for _curItem, element := range m.DiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.DiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__translateBrowsePathsToNodeIdsResponse TranslateBrowsePathsToNodeIdsResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TranslateBrowsePathsToNodeIdsResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TranslateBrowsePathsToNodeIdsResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	noOfResults, err := ReadSimpleField(ctx, "noOfResults", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfResults' field"))
	}
	m.NoOfResults = noOfResults

	results, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "results", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("551")), readBuffer), uint64(noOfResults))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'results' field"))
	}
	m.Results = results

	noOfDiagnosticInfos, err := ReadSimpleField(ctx, "noOfDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfDiagnosticInfos' field"))
	}
	m.NoOfDiagnosticInfos = noOfDiagnosticInfos

	diagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "diagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'diagnosticInfos' field"))
	}
	m.DiagnosticInfos = diagnosticInfos

	if closeErr := readBuffer.CloseContext("TranslateBrowsePathsToNodeIdsResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TranslateBrowsePathsToNodeIdsResponse")
	}

	return m, nil
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TranslateBrowsePathsToNodeIdsResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TranslateBrowsePathsToNodeIdsResponse")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfResults", m.GetNoOfResults(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfResults' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "results", m.GetResults(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'results' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfDiagnosticInfos", m.GetNoOfDiagnosticInfos(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "diagnosticInfos", m.GetDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'diagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("TranslateBrowsePathsToNodeIdsResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TranslateBrowsePathsToNodeIdsResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) IsTranslateBrowsePathsToNodeIdsResponse() {}

func (m *_TranslateBrowsePathsToNodeIdsResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) deepCopy() *_TranslateBrowsePathsToNodeIdsResponse {
	if m == nil {
		return nil
	}
	_TranslateBrowsePathsToNodeIdsResponseCopy := &_TranslateBrowsePathsToNodeIdsResponse{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ResponseHeader.DeepCopy().(ExtensionObjectDefinition),
		m.NoOfResults,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.Results),
		m.NoOfDiagnosticInfos,
		utils.DeepCopySlice[DiagnosticInfo, DiagnosticInfo](m.DiagnosticInfos),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _TranslateBrowsePathsToNodeIdsResponseCopy
}

func (m *_TranslateBrowsePathsToNodeIdsResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
