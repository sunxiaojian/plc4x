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

// ActivateSessionRequest is the corresponding interface of ActivateSessionRequest
type ActivateSessionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetClientSignature returns ClientSignature (property field)
	GetClientSignature() ExtensionObjectDefinition
	// GetNoOfClientSoftwareCertificates returns NoOfClientSoftwareCertificates (property field)
	GetNoOfClientSoftwareCertificates() int32
	// GetClientSoftwareCertificates returns ClientSoftwareCertificates (property field)
	GetClientSoftwareCertificates() []ExtensionObjectDefinition
	// GetNoOfLocaleIds returns NoOfLocaleIds (property field)
	GetNoOfLocaleIds() int32
	// GetLocaleIds returns LocaleIds (property field)
	GetLocaleIds() []PascalString
	// GetUserIdentityToken returns UserIdentityToken (property field)
	GetUserIdentityToken() ExtensionObject
	// GetUserTokenSignature returns UserTokenSignature (property field)
	GetUserTokenSignature() ExtensionObjectDefinition
}

// ActivateSessionRequestExactly can be used when we want exactly this type and not a type which fulfills ActivateSessionRequest.
// This is useful for switch cases.
type ActivateSessionRequestExactly interface {
	ActivateSessionRequest
	isActivateSessionRequest() bool
}

// _ActivateSessionRequest is the data-structure of this message
type _ActivateSessionRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader                  ExtensionObjectDefinition
	ClientSignature                ExtensionObjectDefinition
	NoOfClientSoftwareCertificates int32
	ClientSoftwareCertificates     []ExtensionObjectDefinition
	NoOfLocaleIds                  int32
	LocaleIds                      []PascalString
	UserIdentityToken              ExtensionObject
	UserTokenSignature             ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ActivateSessionRequest) GetIdentifier() string {
	return "467"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ActivateSessionRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ActivateSessionRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ActivateSessionRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_ActivateSessionRequest) GetClientSignature() ExtensionObjectDefinition {
	return m.ClientSignature
}

func (m *_ActivateSessionRequest) GetNoOfClientSoftwareCertificates() int32 {
	return m.NoOfClientSoftwareCertificates
}

func (m *_ActivateSessionRequest) GetClientSoftwareCertificates() []ExtensionObjectDefinition {
	return m.ClientSoftwareCertificates
}

func (m *_ActivateSessionRequest) GetNoOfLocaleIds() int32 {
	return m.NoOfLocaleIds
}

func (m *_ActivateSessionRequest) GetLocaleIds() []PascalString {
	return m.LocaleIds
}

func (m *_ActivateSessionRequest) GetUserIdentityToken() ExtensionObject {
	return m.UserIdentityToken
}

func (m *_ActivateSessionRequest) GetUserTokenSignature() ExtensionObjectDefinition {
	return m.UserTokenSignature
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewActivateSessionRequest factory function for _ActivateSessionRequest
func NewActivateSessionRequest(requestHeader ExtensionObjectDefinition, clientSignature ExtensionObjectDefinition, noOfClientSoftwareCertificates int32, clientSoftwareCertificates []ExtensionObjectDefinition, noOfLocaleIds int32, localeIds []PascalString, userIdentityToken ExtensionObject, userTokenSignature ExtensionObjectDefinition) *_ActivateSessionRequest {
	_result := &_ActivateSessionRequest{
		RequestHeader:                  requestHeader,
		ClientSignature:                clientSignature,
		NoOfClientSoftwareCertificates: noOfClientSoftwareCertificates,
		ClientSoftwareCertificates:     clientSoftwareCertificates,
		NoOfLocaleIds:                  noOfLocaleIds,
		LocaleIds:                      localeIds,
		UserIdentityToken:              userIdentityToken,
		UserTokenSignature:             userTokenSignature,
		_ExtensionObjectDefinition:     NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastActivateSessionRequest(structType any) ActivateSessionRequest {
	if casted, ok := structType.(ActivateSessionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ActivateSessionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ActivateSessionRequest) GetTypeName() string {
	return "ActivateSessionRequest"
}

func (m *_ActivateSessionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (clientSignature)
	lengthInBits += m.ClientSignature.GetLengthInBits(ctx)

	// Simple field (noOfClientSoftwareCertificates)
	lengthInBits += 32

	// Array field
	if len(m.ClientSoftwareCertificates) > 0 {
		for _curItem, element := range m.ClientSoftwareCertificates {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ClientSoftwareCertificates), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfLocaleIds)
	lengthInBits += 32

	// Array field
	if len(m.LocaleIds) > 0 {
		for _curItem, element := range m.LocaleIds {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LocaleIds), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (userIdentityToken)
	lengthInBits += m.UserIdentityToken.GetLengthInBits(ctx)

	// Simple field (userTokenSignature)
	lengthInBits += m.UserTokenSignature.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ActivateSessionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ActivateSessionRequestParse(ctx context.Context, theBytes []byte, identifier string) (ActivateSessionRequest, error) {
	return ActivateSessionRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ActivateSessionRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ActivateSessionRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ActivateSessionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ActivateSessionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of ActivateSessionRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (clientSignature)
	if pullErr := readBuffer.PullContext("clientSignature"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for clientSignature")
	}
	_clientSignature, _clientSignatureErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("458"))
	if _clientSignatureErr != nil {
		return nil, errors.Wrap(_clientSignatureErr, "Error parsing 'clientSignature' field of ActivateSessionRequest")
	}
	clientSignature := _clientSignature.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("clientSignature"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for clientSignature")
	}

	// Simple Field (noOfClientSoftwareCertificates)
	_noOfClientSoftwareCertificates, _noOfClientSoftwareCertificatesErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfClientSoftwareCertificates", 32)
	if _noOfClientSoftwareCertificatesErr != nil {
		return nil, errors.Wrap(_noOfClientSoftwareCertificatesErr, "Error parsing 'noOfClientSoftwareCertificates' field of ActivateSessionRequest")
	}
	noOfClientSoftwareCertificates := _noOfClientSoftwareCertificates

	// Array field (clientSoftwareCertificates)
	if pullErr := readBuffer.PullContext("clientSoftwareCertificates", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for clientSoftwareCertificates")
	}
	// Count array
	clientSoftwareCertificates := make([]ExtensionObjectDefinition, max(noOfClientSoftwareCertificates, 0))
	// This happens when the size is set conditional to 0
	if len(clientSoftwareCertificates) == 0 {
		clientSoftwareCertificates = nil
	}
	{
		_numItems := uint16(max(noOfClientSoftwareCertificates, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "346")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'clientSoftwareCertificates' field of ActivateSessionRequest")
			}
			clientSoftwareCertificates[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("clientSoftwareCertificates", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for clientSoftwareCertificates")
	}

	// Simple Field (noOfLocaleIds)
	_noOfLocaleIds, _noOfLocaleIdsErr := /*TODO: migrate me*/ readBuffer.ReadInt32("noOfLocaleIds", 32)
	if _noOfLocaleIdsErr != nil {
		return nil, errors.Wrap(_noOfLocaleIdsErr, "Error parsing 'noOfLocaleIds' field of ActivateSessionRequest")
	}
	noOfLocaleIds := _noOfLocaleIds

	// Array field (localeIds)
	if pullErr := readBuffer.PullContext("localeIds", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for localeIds")
	}
	// Count array
	localeIds := make([]PascalString, max(noOfLocaleIds, 0))
	// This happens when the size is set conditional to 0
	if len(localeIds) == 0 {
		localeIds = nil
	}
	{
		_numItems := uint16(max(noOfLocaleIds, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := PascalStringParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'localeIds' field of ActivateSessionRequest")
			}
			localeIds[_curItem] = _item.(PascalString)
		}
	}
	if closeErr := readBuffer.CloseContext("localeIds", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for localeIds")
	}

	// Simple Field (userIdentityToken)
	if pullErr := readBuffer.PullContext("userIdentityToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for userIdentityToken")
	}
	_userIdentityToken, _userIdentityTokenErr := ExtensionObjectParseWithBuffer(ctx, readBuffer, bool(bool(true)))
	if _userIdentityTokenErr != nil {
		return nil, errors.Wrap(_userIdentityTokenErr, "Error parsing 'userIdentityToken' field of ActivateSessionRequest")
	}
	userIdentityToken := _userIdentityToken.(ExtensionObject)
	if closeErr := readBuffer.CloseContext("userIdentityToken"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for userIdentityToken")
	}

	// Simple Field (userTokenSignature)
	if pullErr := readBuffer.PullContext("userTokenSignature"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for userTokenSignature")
	}
	_userTokenSignature, _userTokenSignatureErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("458"))
	if _userTokenSignatureErr != nil {
		return nil, errors.Wrap(_userTokenSignatureErr, "Error parsing 'userTokenSignature' field of ActivateSessionRequest")
	}
	userTokenSignature := _userTokenSignature.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("userTokenSignature"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for userTokenSignature")
	}

	if closeErr := readBuffer.CloseContext("ActivateSessionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ActivateSessionRequest")
	}

	// Create a partially initialized instance
	_child := &_ActivateSessionRequest{
		_ExtensionObjectDefinition:     &_ExtensionObjectDefinition{},
		RequestHeader:                  requestHeader,
		ClientSignature:                clientSignature,
		NoOfClientSoftwareCertificates: noOfClientSoftwareCertificates,
		ClientSoftwareCertificates:     clientSoftwareCertificates,
		NoOfLocaleIds:                  noOfLocaleIds,
		LocaleIds:                      localeIds,
		UserIdentityToken:              userIdentityToken,
		UserTokenSignature:             userTokenSignature,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ActivateSessionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ActivateSessionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ActivateSessionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ActivateSessionRequest")
		}

		// Simple Field (requestHeader)
		if pushErr := writeBuffer.PushContext("requestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestHeader")
		}
		_requestHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetRequestHeader())
		if popErr := writeBuffer.PopContext("requestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestHeader")
		}
		if _requestHeaderErr != nil {
			return errors.Wrap(_requestHeaderErr, "Error serializing 'requestHeader' field")
		}

		// Simple Field (clientSignature)
		if pushErr := writeBuffer.PushContext("clientSignature"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for clientSignature")
		}
		_clientSignatureErr := writeBuffer.WriteSerializable(ctx, m.GetClientSignature())
		if popErr := writeBuffer.PopContext("clientSignature"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for clientSignature")
		}
		if _clientSignatureErr != nil {
			return errors.Wrap(_clientSignatureErr, "Error serializing 'clientSignature' field")
		}

		// Simple Field (noOfClientSoftwareCertificates)
		noOfClientSoftwareCertificates := int32(m.GetNoOfClientSoftwareCertificates())
		_noOfClientSoftwareCertificatesErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfClientSoftwareCertificates", 32, int32((noOfClientSoftwareCertificates)))
		if _noOfClientSoftwareCertificatesErr != nil {
			return errors.Wrap(_noOfClientSoftwareCertificatesErr, "Error serializing 'noOfClientSoftwareCertificates' field")
		}

		// Array Field (clientSoftwareCertificates)
		if pushErr := writeBuffer.PushContext("clientSoftwareCertificates", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for clientSoftwareCertificates")
		}
		for _curItem, _element := range m.GetClientSoftwareCertificates() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetClientSoftwareCertificates()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'clientSoftwareCertificates' field")
			}
		}
		if popErr := writeBuffer.PopContext("clientSoftwareCertificates", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for clientSoftwareCertificates")
		}

		// Simple Field (noOfLocaleIds)
		noOfLocaleIds := int32(m.GetNoOfLocaleIds())
		_noOfLocaleIdsErr := /*TODO: migrate me*/ writeBuffer.WriteInt32("noOfLocaleIds", 32, int32((noOfLocaleIds)))
		if _noOfLocaleIdsErr != nil {
			return errors.Wrap(_noOfLocaleIdsErr, "Error serializing 'noOfLocaleIds' field")
		}

		// Array Field (localeIds)
		if pushErr := writeBuffer.PushContext("localeIds", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for localeIds")
		}
		for _curItem, _element := range m.GetLocaleIds() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetLocaleIds()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'localeIds' field")
			}
		}
		if popErr := writeBuffer.PopContext("localeIds", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for localeIds")
		}

		// Simple Field (userIdentityToken)
		if pushErr := writeBuffer.PushContext("userIdentityToken"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for userIdentityToken")
		}
		_userIdentityTokenErr := writeBuffer.WriteSerializable(ctx, m.GetUserIdentityToken())
		if popErr := writeBuffer.PopContext("userIdentityToken"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for userIdentityToken")
		}
		if _userIdentityTokenErr != nil {
			return errors.Wrap(_userIdentityTokenErr, "Error serializing 'userIdentityToken' field")
		}

		// Simple Field (userTokenSignature)
		if pushErr := writeBuffer.PushContext("userTokenSignature"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for userTokenSignature")
		}
		_userTokenSignatureErr := writeBuffer.WriteSerializable(ctx, m.GetUserTokenSignature())
		if popErr := writeBuffer.PopContext("userTokenSignature"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for userTokenSignature")
		}
		if _userTokenSignatureErr != nil {
			return errors.Wrap(_userTokenSignatureErr, "Error serializing 'userTokenSignature' field")
		}

		if popErr := writeBuffer.PopContext("ActivateSessionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ActivateSessionRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ActivateSessionRequest) isActivateSessionRequest() bool {
	return true
}

func (m *_ActivateSessionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
