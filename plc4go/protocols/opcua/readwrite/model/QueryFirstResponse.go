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

// QueryFirstResponse is the corresponding interface of QueryFirstResponse
type QueryFirstResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfQueryDataSets returns NoOfQueryDataSets (property field)
	GetNoOfQueryDataSets() int32
	// GetQueryDataSets returns QueryDataSets (property field)
	GetQueryDataSets() []ExtensionObjectDefinition
	// GetContinuationPoint returns ContinuationPoint (property field)
	GetContinuationPoint() PascalByteString
	// GetNoOfParsingResults returns NoOfParsingResults (property field)
	GetNoOfParsingResults() int32
	// GetParsingResults returns ParsingResults (property field)
	GetParsingResults() []ExtensionObjectDefinition
	// GetNoOfDiagnosticInfos returns NoOfDiagnosticInfos (property field)
	GetNoOfDiagnosticInfos() int32
	// GetDiagnosticInfos returns DiagnosticInfos (property field)
	GetDiagnosticInfos() []DiagnosticInfo
	// GetFilterResult returns FilterResult (property field)
	GetFilterResult() ExtensionObjectDefinition
}

// QueryFirstResponseExactly can be used when we want exactly this type and not a type which fulfills QueryFirstResponse.
// This is useful for switch cases.
type QueryFirstResponseExactly interface {
	QueryFirstResponse
	isQueryFirstResponse() bool
}

// _QueryFirstResponse is the data-structure of this message
type _QueryFirstResponse struct {
	*_ExtensionObjectDefinition
	ResponseHeader      ExtensionObjectDefinition
	NoOfQueryDataSets   int32
	QueryDataSets       []ExtensionObjectDefinition
	ContinuationPoint   PascalByteString
	NoOfParsingResults  int32
	ParsingResults      []ExtensionObjectDefinition
	NoOfDiagnosticInfos int32
	DiagnosticInfos     []DiagnosticInfo
	FilterResult        ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_QueryFirstResponse) GetIdentifier() string {
	return "618"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_QueryFirstResponse) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_QueryFirstResponse) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_QueryFirstResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_QueryFirstResponse) GetNoOfQueryDataSets() int32 {
	return m.NoOfQueryDataSets
}

func (m *_QueryFirstResponse) GetQueryDataSets() []ExtensionObjectDefinition {
	return m.QueryDataSets
}

func (m *_QueryFirstResponse) GetContinuationPoint() PascalByteString {
	return m.ContinuationPoint
}

func (m *_QueryFirstResponse) GetNoOfParsingResults() int32 {
	return m.NoOfParsingResults
}

func (m *_QueryFirstResponse) GetParsingResults() []ExtensionObjectDefinition {
	return m.ParsingResults
}

func (m *_QueryFirstResponse) GetNoOfDiagnosticInfos() int32 {
	return m.NoOfDiagnosticInfos
}

func (m *_QueryFirstResponse) GetDiagnosticInfos() []DiagnosticInfo {
	return m.DiagnosticInfos
}

func (m *_QueryFirstResponse) GetFilterResult() ExtensionObjectDefinition {
	return m.FilterResult
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewQueryFirstResponse factory function for _QueryFirstResponse
func NewQueryFirstResponse(responseHeader ExtensionObjectDefinition, noOfQueryDataSets int32, queryDataSets []ExtensionObjectDefinition, continuationPoint PascalByteString, noOfParsingResults int32, parsingResults []ExtensionObjectDefinition, noOfDiagnosticInfos int32, diagnosticInfos []DiagnosticInfo, filterResult ExtensionObjectDefinition) *_QueryFirstResponse {
	_result := &_QueryFirstResponse{
		ResponseHeader:             responseHeader,
		NoOfQueryDataSets:          noOfQueryDataSets,
		QueryDataSets:              queryDataSets,
		ContinuationPoint:          continuationPoint,
		NoOfParsingResults:         noOfParsingResults,
		ParsingResults:             parsingResults,
		NoOfDiagnosticInfos:        noOfDiagnosticInfos,
		DiagnosticInfos:            diagnosticInfos,
		FilterResult:               filterResult,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastQueryFirstResponse(structType any) QueryFirstResponse {
	if casted, ok := structType.(QueryFirstResponse); ok {
		return casted
	}
	if casted, ok := structType.(*QueryFirstResponse); ok {
		return *casted
	}
	return nil
}

func (m *_QueryFirstResponse) GetTypeName() string {
	return "QueryFirstResponse"
}

func (m *_QueryFirstResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (noOfQueryDataSets)
	lengthInBits += 32

	// Array field
	if len(m.QueryDataSets) > 0 {
		for _curItem, element := range m.QueryDataSets {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.QueryDataSets), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (continuationPoint)
	lengthInBits += m.ContinuationPoint.GetLengthInBits(ctx)

	// Simple field (noOfParsingResults)
	lengthInBits += 32

	// Array field
	if len(m.ParsingResults) > 0 {
		for _curItem, element := range m.ParsingResults {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ParsingResults), _curItem)
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

	// Simple field (filterResult)
	lengthInBits += m.FilterResult.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_QueryFirstResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func QueryFirstResponseParse(ctx context.Context, theBytes []byte, identifier string) (QueryFirstResponse, error) {
	return QueryFirstResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func QueryFirstResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (QueryFirstResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("QueryFirstResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for QueryFirstResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (responseHeader)
	if pullErr := readBuffer.PullContext("responseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for responseHeader")
	}
	_responseHeader, _responseHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("394"))
	if _responseHeaderErr != nil {
		return nil, errors.Wrap(_responseHeaderErr, "Error parsing 'responseHeader' field of QueryFirstResponse")
	}
	responseHeader := _responseHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("responseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for responseHeader")
	}

	// Simple Field (noOfQueryDataSets)
	_noOfQueryDataSets, _noOfQueryDataSetsErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfQueryDataSets", 32)
	if _noOfQueryDataSetsErr != nil {
		return nil, errors.Wrap(_noOfQueryDataSetsErr, "Error parsing 'noOfQueryDataSets' field of QueryFirstResponse")
	}
	noOfQueryDataSets := _noOfQueryDataSets

	// Array field (queryDataSets)
	if pullErr := readBuffer.PullContext("queryDataSets", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for queryDataSets")
	}
	// Count array
	queryDataSets := make([]ExtensionObjectDefinition, max(noOfQueryDataSets, 0))
	// This happens when the size is set conditional to 0
	if len(queryDataSets) == 0 {
		queryDataSets = nil
	}
	{
		_numItems := uint16(max(noOfQueryDataSets, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "579")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'queryDataSets' field of QueryFirstResponse")
			}
			queryDataSets[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("queryDataSets", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for queryDataSets")
	}

	// Simple Field (continuationPoint)
	if pullErr := readBuffer.PullContext("continuationPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for continuationPoint")
	}
	_continuationPoint, _continuationPointErr := PascalByteStringParseWithBuffer(ctx, readBuffer)
	if _continuationPointErr != nil {
		return nil, errors.Wrap(_continuationPointErr, "Error parsing 'continuationPoint' field of QueryFirstResponse")
	}
	continuationPoint := _continuationPoint.(PascalByteString)
	if closeErr := readBuffer.CloseContext("continuationPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for continuationPoint")
	}

	// Simple Field (noOfParsingResults)
	_noOfParsingResults, _noOfParsingResultsErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfParsingResults", 32)
	if _noOfParsingResultsErr != nil {
		return nil, errors.Wrap(_noOfParsingResultsErr, "Error parsing 'noOfParsingResults' field of QueryFirstResponse")
	}
	noOfParsingResults := _noOfParsingResults

	// Array field (parsingResults)
	if pullErr := readBuffer.PullContext("parsingResults", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for parsingResults")
	}
	// Count array
	parsingResults := make([]ExtensionObjectDefinition, max(noOfParsingResults, 0))
	// This happens when the size is set conditional to 0
	if len(parsingResults) == 0 {
		parsingResults = nil
	}
	{
		_numItems := uint16(max(noOfParsingResults, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "612")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'parsingResults' field of QueryFirstResponse")
			}
			parsingResults[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("parsingResults", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for parsingResults")
	}

	// Simple Field (noOfDiagnosticInfos)
	_noOfDiagnosticInfos, _noOfDiagnosticInfosErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfDiagnosticInfos", 32)
	if _noOfDiagnosticInfosErr != nil {
		return nil, errors.Wrap(_noOfDiagnosticInfosErr, "Error parsing 'noOfDiagnosticInfos' field of QueryFirstResponse")
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
				return nil, errors.Wrap(_err, "Error parsing 'diagnosticInfos' field of QueryFirstResponse")
			}
			diagnosticInfos[_curItem] = _item.(DiagnosticInfo)
		}
	}
	if closeErr := readBuffer.CloseContext("diagnosticInfos", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for diagnosticInfos")
	}

	// Simple Field (filterResult)
	if pullErr := readBuffer.PullContext("filterResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for filterResult")
	}
	_filterResult, _filterResultErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("609"))
	if _filterResultErr != nil {
		return nil, errors.Wrap(_filterResultErr, "Error parsing 'filterResult' field of QueryFirstResponse")
	}
	filterResult := _filterResult.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("filterResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for filterResult")
	}

	if closeErr := readBuffer.CloseContext("QueryFirstResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for QueryFirstResponse")
	}

	// Create a partially initialized instance
	_child := &_QueryFirstResponse{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ResponseHeader:             responseHeader,
		NoOfQueryDataSets:          noOfQueryDataSets,
		QueryDataSets:              queryDataSets,
		ContinuationPoint:          continuationPoint,
		NoOfParsingResults:         noOfParsingResults,
		ParsingResults:             parsingResults,
		NoOfDiagnosticInfos:        noOfDiagnosticInfos,
		DiagnosticInfos:            diagnosticInfos,
		FilterResult:               filterResult,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_QueryFirstResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_QueryFirstResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("QueryFirstResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for QueryFirstResponse")
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

		// Simple Field (noOfQueryDataSets)
		noOfQueryDataSets := int32(m.GetNoOfQueryDataSets())
		_noOfQueryDataSetsErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfQueryDataSets", 32, int32((noOfQueryDataSets)))
		if _noOfQueryDataSetsErr != nil {
			return errors.Wrap(_noOfQueryDataSetsErr, "Error serializing 'noOfQueryDataSets' field")
		}

		// Array Field (queryDataSets)
		if pushErr := writeBuffer.PushContext("queryDataSets", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for queryDataSets")
		}
		for _curItem, _element := range m.GetQueryDataSets() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetQueryDataSets()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'queryDataSets' field")
			}
		}
		if popErr := writeBuffer.PopContext("queryDataSets", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for queryDataSets")
		}

		// Simple Field (continuationPoint)
		if pushErr := writeBuffer.PushContext("continuationPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for continuationPoint")
		}
		_continuationPointErr := writeBuffer.WriteSerializable(ctx, m.GetContinuationPoint())
		if popErr := writeBuffer.PopContext("continuationPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for continuationPoint")
		}
		if _continuationPointErr != nil {
			return errors.Wrap(_continuationPointErr, "Error serializing 'continuationPoint' field")
		}

		// Simple Field (noOfParsingResults)
		noOfParsingResults := int32(m.GetNoOfParsingResults())
		_noOfParsingResultsErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfParsingResults", 32, int32((noOfParsingResults)))
		if _noOfParsingResultsErr != nil {
			return errors.Wrap(_noOfParsingResultsErr, "Error serializing 'noOfParsingResults' field")
		}

		// Array Field (parsingResults)
		if pushErr := writeBuffer.PushContext("parsingResults", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for parsingResults")
		}
		for _curItem, _element := range m.GetParsingResults() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetParsingResults()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'parsingResults' field")
			}
		}
		if popErr := writeBuffer.PopContext("parsingResults", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for parsingResults")
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

		// Simple Field (filterResult)
		if pushErr := writeBuffer.PushContext("filterResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for filterResult")
		}
		_filterResultErr := writeBuffer.WriteSerializable(ctx, m.GetFilterResult())
		if popErr := writeBuffer.PopContext("filterResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for filterResult")
		}
		if _filterResultErr != nil {
			return errors.Wrap(_filterResultErr, "Error serializing 'filterResult' field")
		}

		if popErr := writeBuffer.PopContext("QueryFirstResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for QueryFirstResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_QueryFirstResponse) isQueryFirstResponse() bool {
	return true
}

func (m *_QueryFirstResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
