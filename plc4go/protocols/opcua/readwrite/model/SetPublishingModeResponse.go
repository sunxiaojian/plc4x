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

// SetPublishingModeResponse is the corresponding interface of SetPublishingModeResponse
type SetPublishingModeResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfResults returns NoOfResults (property field)
	GetNoOfResults() int32
	// GetResults returns Results (property field)
	GetResults() []StatusCode
	// GetNoOfDiagnosticInfos returns NoOfDiagnosticInfos (property field)
	GetNoOfDiagnosticInfos() int32
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
}

// SetPublishingModeResponseExactly can be used when we want exactly this type and not a type which fulfills SetPublishingModeResponse.
// This is useful for switch cases.
type SetPublishingModeResponseExactly interface {
	SetPublishingModeResponse
	isSetPublishingModeResponse() bool
}

// _SetPublishingModeResponse is the data-structure of this message
type _SetPublishingModeResponse struct {
	*_ExtensionObjectDefinition
	ResponseHeader      ExtensionObjectDefinition
	NoOfResults         int32
	Results             []StatusCode
	NoOfDiagnosticInfos int32
	DiagnosticInfos     []DiagnosticInfo
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SetPublishingModeResponse) GetIdentifier() string {
	return "802"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SetPublishingModeResponse) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_SetPublishingModeResponse) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SetPublishingModeResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_SetPublishingModeResponse) GetNoOfResults() int32 {
	return m.NoOfResults
}

func (m *_SetPublishingModeResponse) GetResults() []StatusCode {
	return m.Results
}

func (m *_SetPublishingModeResponse) GetNoOfDiagnosticInfos() int32 {
	return m.NoOfDiagnosticInfos
}

func (m *_SetPublishingModeResponse) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSetPublishingModeResponse factory function for _SetPublishingModeResponse
func NewSetPublishingModeResponse(responseHeader ExtensionObjectDefinition, noOfResults int32, results []StatusCode, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo) *_SetPublishingModeResponse {
	_result := &_SetPublishingModeResponse{
		ResponseHeader:             responseHeader,
		NoOfResults:                noOfResults,
		Results:                    results,
		NoOfDiagnosticInfos:        noOfDiagnosticInfos,
		DiagnosticInfos:            diagnosticInfos,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSetPublishingModeResponse(structType any) SetPublishingModeResponse {
	if casted, ok := structType.(SetPublishingModeResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SetPublishingModeResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SetPublishingModeResponse) GetTypeName() string {
	return "SetPublishingModeResponse"
}

func (m *_SetPublishingModeResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

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

func (m *_SetPublishingModeResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SetPublishingModeResponseParse(ctx context.Context, theBytes []byte, identifier string) (SetPublishingModeResponse, error) {
	return SetPublishingModeResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func SetPublishingModeResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (SetPublishingModeResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SetPublishingModeResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SetPublishingModeResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (responseHeader)
	if pullErr := readBuffer.PullContext("responseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for responseHeader")
	}
	_responseHeader, _responseHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("394"))
	if _responseHeaderErr != nil {
		return nil, errors.Wrap(_responseHeaderErr, "Error parsing 'responseHeader' field of SetPublishingModeResponse")
	}
	responseHeader := _responseHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("responseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for responseHeader")
	}

	// Simple Field (noOfResults)
	_noOfResults, _noOfResultsErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfResults", 32)
	if _noOfResultsErr != nil {
		return nil, errors.Wrap(_noOfResultsErr, "Error parsing 'noOfResults' field of SetPublishingModeResponse")
	}
	noOfResults := _noOfResults

	// Array field (results)
	if pullErr := readBuffer.PullContext("results", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for results")
	}
	// Count array
	results := make([]StatusCode, max(noOfResults, 0))
	// This happens when the size is set conditional to 0
	if len(results) == 0 {
		results = nil
	}
	{
		_numItems := uint16(max(noOfResults, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := StatusCodeParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'results' field of SetPublishingModeResponse")
			}
			results[_curItem] = _item.(StatusCode)
		}
	}
	if closeErr := readBuffer.CloseContext("results", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for results")
	}

	// Simple Field (noOfDiagnosticInfos)
	_noOfDiagnosticInfos, _noOfDiagnosticInfosErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfDiagnosticInfos", 32)
	if _noOfDiagnosticInfosErr != nil {
		return nil, errors.Wrap(_noOfDiagnosticInfosErr, "Error parsing 'noOfDiagnosticInfos' field of SetPublishingModeResponse")
	}
	noOfDiagnosticInfos := _noOfDiagnosticInfos

	// Array field (diagnosticInfos)
	if pullErr := readBuffer.PullContext("diagnosticInfos", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for diagnosticInfos")
	}
	// Count array
	diagnosticInfos := make([]DiagnosticInfo, max(noOfDiagnosticInfos, 0))
	// This happens when the size is set conditional to 0
	if len(diagnosticInfos) == 0 {
		diagnosticInfos = nil
	}
	{
		_numItems := uint16(max(noOfDiagnosticInfos, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := DiagnosticInfoParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'diagnosticInfos' field of SetPublishingModeResponse")
			}
			diagnosticInfos[_curItem] = _item.(DiagnosticInfo)
		}
	}
	if closeErr := readBuffer.CloseContext("diagnosticInfos", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for diagnosticInfos")
	}

	if closeErr := readBuffer.CloseContext("SetPublishingModeResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SetPublishingModeResponse")
	}

	// Create a partially initialized instance
	_child := &_SetPublishingModeResponse{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ResponseHeader:             responseHeader,
		NoOfResults:                noOfResults,
		Results:                    results,
		NoOfDiagnosticInfos:        noOfDiagnosticInfos,
		DiagnosticInfos:            diagnosticInfos,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_SetPublishingModeResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SetPublishingModeResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SetPublishingModeResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SetPublishingModeResponse")
		}

		// Simple Field (responseHeader)
		if pushErr := writeBuffer.PushContext("responseHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for responseHeader")
		}
		_responseHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetResponseHeader())
		if popErr := writeBuffer.PopContext("responseHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for responseHeader")
		}
		if _responseHeaderErr != nil {
			return errors.Wrap(_responseHeaderErr, "Error serializing 'responseHeader' field")
		}

		// Simple Field (noOfResults)
		noOfResults := int32(m.GetNoOfResults())
		_noOfResultsErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfResults", 32, int32((noOfResults)))
		if _noOfResultsErr != nil {
			return errors.Wrap(_noOfResultsErr, "Error serializing 'noOfResults' field")
		}

		// Array Field (results)
		if pushErr := writeBuffer.PushContext("results", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for results")
		}
		for _curItem, _element := range m.GetResults() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetResults()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'results' field")
			}
		}
		if popErr := writeBuffer.PopContext("results", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for results")
		}

		// Simple Field (noOfDiagnosticInfos)
		noOfDiagnosticInfos := int32(m.GetNoOfDiagnosticInfos())
		_noOfDiagnosticInfosErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfDiagnosticInfos", 32, int32((noOfDiagnosticInfos)))
		if _noOfDiagnosticInfosErr != nil {
			return errors.Wrap(_noOfDiagnosticInfosErr, "Error serializing 'noOfDiagnosticInfos' field")
		}

		// Array Field (diagnosticInfos)
		if pushErr := writeBuffer.PushContext("diagnosticInfos", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for diagnosticInfos")
		}
		for _curItem, _element := range m.GetDiagnosticInfos() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetDiagnosticInfos()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'diagnosticInfos' field")
			}
		}
		if popErr := writeBuffer.PopContext("diagnosticInfos", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for diagnosticInfos")
		}

		if popErr := writeBuffer.PopContext("SetPublishingModeResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SetPublishingModeResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SetPublishingModeResponse) isSetPublishingModeResponse() bool {
	return true
}

func (m *_SetPublishingModeResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
