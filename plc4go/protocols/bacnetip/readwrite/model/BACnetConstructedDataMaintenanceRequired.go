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

// BACnetConstructedDataMaintenanceRequired is the corresponding interface of BACnetConstructedDataMaintenanceRequired
type BACnetConstructedDataMaintenanceRequired interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMaintenanceRequired returns MaintenanceRequired (property field)
	GetMaintenanceRequired() BACnetMaintenanceTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetMaintenanceTagged
	// IsBACnetConstructedDataMaintenanceRequired is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMaintenanceRequired()
	// CreateBuilder creates a BACnetConstructedDataMaintenanceRequiredBuilder
	CreateBACnetConstructedDataMaintenanceRequiredBuilder() BACnetConstructedDataMaintenanceRequiredBuilder
}

// _BACnetConstructedDataMaintenanceRequired is the data-structure of this message
type _BACnetConstructedDataMaintenanceRequired struct {
	BACnetConstructedDataContract
	MaintenanceRequired BACnetMaintenanceTagged
}

var _ BACnetConstructedDataMaintenanceRequired = (*_BACnetConstructedDataMaintenanceRequired)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMaintenanceRequired)(nil)

// NewBACnetConstructedDataMaintenanceRequired factory function for _BACnetConstructedDataMaintenanceRequired
func NewBACnetConstructedDataMaintenanceRequired(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, maintenanceRequired BACnetMaintenanceTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMaintenanceRequired {
	if maintenanceRequired == nil {
		panic("maintenanceRequired of type BACnetMaintenanceTagged for BACnetConstructedDataMaintenanceRequired must not be nil")
	}
	_result := &_BACnetConstructedDataMaintenanceRequired{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaintenanceRequired:           maintenanceRequired,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMaintenanceRequiredBuilder is a builder for BACnetConstructedDataMaintenanceRequired
type BACnetConstructedDataMaintenanceRequiredBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maintenanceRequired BACnetMaintenanceTagged) BACnetConstructedDataMaintenanceRequiredBuilder
	// WithMaintenanceRequired adds MaintenanceRequired (property field)
	WithMaintenanceRequired(BACnetMaintenanceTagged) BACnetConstructedDataMaintenanceRequiredBuilder
	// WithMaintenanceRequiredBuilder adds MaintenanceRequired (property field) which is build by the builder
	WithMaintenanceRequiredBuilder(func(BACnetMaintenanceTaggedBuilder) BACnetMaintenanceTaggedBuilder) BACnetConstructedDataMaintenanceRequiredBuilder
	// Build builds the BACnetConstructedDataMaintenanceRequired or returns an error if something is wrong
	Build() (BACnetConstructedDataMaintenanceRequired, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMaintenanceRequired
}

// NewBACnetConstructedDataMaintenanceRequiredBuilder() creates a BACnetConstructedDataMaintenanceRequiredBuilder
func NewBACnetConstructedDataMaintenanceRequiredBuilder() BACnetConstructedDataMaintenanceRequiredBuilder {
	return &_BACnetConstructedDataMaintenanceRequiredBuilder{_BACnetConstructedDataMaintenanceRequired: new(_BACnetConstructedDataMaintenanceRequired)}
}

type _BACnetConstructedDataMaintenanceRequiredBuilder struct {
	*_BACnetConstructedDataMaintenanceRequired

	err *utils.MultiError
}

var _ (BACnetConstructedDataMaintenanceRequiredBuilder) = (*_BACnetConstructedDataMaintenanceRequiredBuilder)(nil)

func (m *_BACnetConstructedDataMaintenanceRequiredBuilder) WithMandatoryFields(maintenanceRequired BACnetMaintenanceTagged) BACnetConstructedDataMaintenanceRequiredBuilder {
	return m.WithMaintenanceRequired(maintenanceRequired)
}

func (m *_BACnetConstructedDataMaintenanceRequiredBuilder) WithMaintenanceRequired(maintenanceRequired BACnetMaintenanceTagged) BACnetConstructedDataMaintenanceRequiredBuilder {
	m.MaintenanceRequired = maintenanceRequired
	return m
}

func (m *_BACnetConstructedDataMaintenanceRequiredBuilder) WithMaintenanceRequiredBuilder(builderSupplier func(BACnetMaintenanceTaggedBuilder) BACnetMaintenanceTaggedBuilder) BACnetConstructedDataMaintenanceRequiredBuilder {
	builder := builderSupplier(m.MaintenanceRequired.CreateBACnetMaintenanceTaggedBuilder())
	var err error
	m.MaintenanceRequired, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetMaintenanceTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataMaintenanceRequiredBuilder) Build() (BACnetConstructedDataMaintenanceRequired, error) {
	if m.MaintenanceRequired == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'maintenanceRequired' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataMaintenanceRequired.deepCopy(), nil
}

func (m *_BACnetConstructedDataMaintenanceRequiredBuilder) MustBuild() BACnetConstructedDataMaintenanceRequired {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataMaintenanceRequiredBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataMaintenanceRequiredBuilder()
}

// CreateBACnetConstructedDataMaintenanceRequiredBuilder creates a BACnetConstructedDataMaintenanceRequiredBuilder
func (m *_BACnetConstructedDataMaintenanceRequired) CreateBACnetConstructedDataMaintenanceRequiredBuilder() BACnetConstructedDataMaintenanceRequiredBuilder {
	if m == nil {
		return NewBACnetConstructedDataMaintenanceRequiredBuilder()
	}
	return &_BACnetConstructedDataMaintenanceRequiredBuilder{_BACnetConstructedDataMaintenanceRequired: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaintenanceRequired) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMaintenanceRequired) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAINTENANCE_REQUIRED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaintenanceRequired) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaintenanceRequired) GetMaintenanceRequired() BACnetMaintenanceTagged {
	return m.MaintenanceRequired
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMaintenanceRequired) GetActualValue() BACnetMaintenanceTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetMaintenanceTagged(m.GetMaintenanceRequired())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaintenanceRequired(structType any) BACnetConstructedDataMaintenanceRequired {
	if casted, ok := structType.(BACnetConstructedDataMaintenanceRequired); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaintenanceRequired); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaintenanceRequired) GetTypeName() string {
	return "BACnetConstructedDataMaintenanceRequired"
}

func (m *_BACnetConstructedDataMaintenanceRequired) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maintenanceRequired)
	lengthInBits += m.MaintenanceRequired.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMaintenanceRequired) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMaintenanceRequired) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMaintenanceRequired BACnetConstructedDataMaintenanceRequired, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaintenanceRequired"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaintenanceRequired")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maintenanceRequired, err := ReadSimpleField[BACnetMaintenanceTagged](ctx, "maintenanceRequired", ReadComplex[BACnetMaintenanceTagged](BACnetMaintenanceTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maintenanceRequired' field"))
	}
	m.MaintenanceRequired = maintenanceRequired

	actualValue, err := ReadVirtualField[BACnetMaintenanceTagged](ctx, "actualValue", (*BACnetMaintenanceTagged)(nil), maintenanceRequired)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaintenanceRequired"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaintenanceRequired")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMaintenanceRequired) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMaintenanceRequired) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaintenanceRequired"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaintenanceRequired")
		}

		if err := WriteSimpleField[BACnetMaintenanceTagged](ctx, "maintenanceRequired", m.GetMaintenanceRequired(), WriteComplex[BACnetMaintenanceTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maintenanceRequired' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaintenanceRequired"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaintenanceRequired")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMaintenanceRequired) IsBACnetConstructedDataMaintenanceRequired() {}

func (m *_BACnetConstructedDataMaintenanceRequired) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMaintenanceRequired) deepCopy() *_BACnetConstructedDataMaintenanceRequired {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMaintenanceRequiredCopy := &_BACnetConstructedDataMaintenanceRequired{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.MaintenanceRequired.DeepCopy().(BACnetMaintenanceTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMaintenanceRequiredCopy
}

func (m *_BACnetConstructedDataMaintenanceRequired) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
