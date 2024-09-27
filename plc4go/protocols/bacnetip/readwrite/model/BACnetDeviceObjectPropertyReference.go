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

// BACnetDeviceObjectPropertyReference is the corresponding interface of BACnetDeviceObjectPropertyReference
type BACnetDeviceObjectPropertyReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() BACnetContextTagObjectIdentifier
	// IsBACnetDeviceObjectPropertyReference is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetDeviceObjectPropertyReference()
	// CreateBuilder creates a BACnetDeviceObjectPropertyReferenceBuilder
	CreateBACnetDeviceObjectPropertyReferenceBuilder() BACnetDeviceObjectPropertyReferenceBuilder
}

// _BACnetDeviceObjectPropertyReference is the data-structure of this message
type _BACnetDeviceObjectPropertyReference struct {
	ObjectIdentifier   BACnetContextTagObjectIdentifier
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
	DeviceIdentifier   BACnetContextTagObjectIdentifier
}

var _ BACnetDeviceObjectPropertyReference = (*_BACnetDeviceObjectPropertyReference)(nil)

// NewBACnetDeviceObjectPropertyReference factory function for _BACnetDeviceObjectPropertyReference
func NewBACnetDeviceObjectPropertyReference(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger, deviceIdentifier BACnetContextTagObjectIdentifier) *_BACnetDeviceObjectPropertyReference {
	if objectIdentifier == nil {
		panic("objectIdentifier of type BACnetContextTagObjectIdentifier for BACnetDeviceObjectPropertyReference must not be nil")
	}
	if propertyIdentifier == nil {
		panic("propertyIdentifier of type BACnetPropertyIdentifierTagged for BACnetDeviceObjectPropertyReference must not be nil")
	}
	return &_BACnetDeviceObjectPropertyReference{ObjectIdentifier: objectIdentifier, PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex, DeviceIdentifier: deviceIdentifier}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetDeviceObjectPropertyReferenceBuilder is a builder for BACnetDeviceObjectPropertyReference
type BACnetDeviceObjectPropertyReferenceBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged) BACnetDeviceObjectPropertyReferenceBuilder
	// WithObjectIdentifier adds ObjectIdentifier (property field)
	WithObjectIdentifier(BACnetContextTagObjectIdentifier) BACnetDeviceObjectPropertyReferenceBuilder
	// WithObjectIdentifierBuilder adds ObjectIdentifier (property field) which is build by the builder
	WithObjectIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetDeviceObjectPropertyReferenceBuilder
	// WithPropertyIdentifier adds PropertyIdentifier (property field)
	WithPropertyIdentifier(BACnetPropertyIdentifierTagged) BACnetDeviceObjectPropertyReferenceBuilder
	// WithPropertyIdentifierBuilder adds PropertyIdentifier (property field) which is build by the builder
	WithPropertyIdentifierBuilder(func(BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyIdentifierTaggedBuilder) BACnetDeviceObjectPropertyReferenceBuilder
	// WithArrayIndex adds ArrayIndex (property field)
	WithOptionalArrayIndex(BACnetContextTagUnsignedInteger) BACnetDeviceObjectPropertyReferenceBuilder
	// WithOptionalArrayIndexBuilder adds ArrayIndex (property field) which is build by the builder
	WithOptionalArrayIndexBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetDeviceObjectPropertyReferenceBuilder
	// WithDeviceIdentifier adds DeviceIdentifier (property field)
	WithOptionalDeviceIdentifier(BACnetContextTagObjectIdentifier) BACnetDeviceObjectPropertyReferenceBuilder
	// WithOptionalDeviceIdentifierBuilder adds DeviceIdentifier (property field) which is build by the builder
	WithOptionalDeviceIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetDeviceObjectPropertyReferenceBuilder
	// Build builds the BACnetDeviceObjectPropertyReference or returns an error if something is wrong
	Build() (BACnetDeviceObjectPropertyReference, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetDeviceObjectPropertyReference
}

// NewBACnetDeviceObjectPropertyReferenceBuilder() creates a BACnetDeviceObjectPropertyReferenceBuilder
func NewBACnetDeviceObjectPropertyReferenceBuilder() BACnetDeviceObjectPropertyReferenceBuilder {
	return &_BACnetDeviceObjectPropertyReferenceBuilder{_BACnetDeviceObjectPropertyReference: new(_BACnetDeviceObjectPropertyReference)}
}

type _BACnetDeviceObjectPropertyReferenceBuilder struct {
	*_BACnetDeviceObjectPropertyReference

	err *utils.MultiError
}

var _ (BACnetDeviceObjectPropertyReferenceBuilder) = (*_BACnetDeviceObjectPropertyReferenceBuilder)(nil)

func (m *_BACnetDeviceObjectPropertyReferenceBuilder) WithMandatoryFields(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged) BACnetDeviceObjectPropertyReferenceBuilder {
	return m.WithObjectIdentifier(objectIdentifier).WithPropertyIdentifier(propertyIdentifier)
}

func (m *_BACnetDeviceObjectPropertyReferenceBuilder) WithObjectIdentifier(objectIdentifier BACnetContextTagObjectIdentifier) BACnetDeviceObjectPropertyReferenceBuilder {
	m.ObjectIdentifier = objectIdentifier
	return m
}

func (m *_BACnetDeviceObjectPropertyReferenceBuilder) WithObjectIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetDeviceObjectPropertyReferenceBuilder {
	builder := builderSupplier(m.ObjectIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	m.ObjectIdentifier, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return m
}

func (m *_BACnetDeviceObjectPropertyReferenceBuilder) WithPropertyIdentifier(propertyIdentifier BACnetPropertyIdentifierTagged) BACnetDeviceObjectPropertyReferenceBuilder {
	m.PropertyIdentifier = propertyIdentifier
	return m
}

func (m *_BACnetDeviceObjectPropertyReferenceBuilder) WithPropertyIdentifierBuilder(builderSupplier func(BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyIdentifierTaggedBuilder) BACnetDeviceObjectPropertyReferenceBuilder {
	builder := builderSupplier(m.PropertyIdentifier.CreateBACnetPropertyIdentifierTaggedBuilder())
	var err error
	m.PropertyIdentifier, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetPropertyIdentifierTaggedBuilder failed"))
	}
	return m
}

func (m *_BACnetDeviceObjectPropertyReferenceBuilder) WithOptionalArrayIndex(arrayIndex BACnetContextTagUnsignedInteger) BACnetDeviceObjectPropertyReferenceBuilder {
	m.ArrayIndex = arrayIndex
	return m
}

func (m *_BACnetDeviceObjectPropertyReferenceBuilder) WithOptionalArrayIndexBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetDeviceObjectPropertyReferenceBuilder {
	builder := builderSupplier(m.ArrayIndex.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	m.ArrayIndex, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetDeviceObjectPropertyReferenceBuilder) WithOptionalDeviceIdentifier(deviceIdentifier BACnetContextTagObjectIdentifier) BACnetDeviceObjectPropertyReferenceBuilder {
	m.DeviceIdentifier = deviceIdentifier
	return m
}

func (m *_BACnetDeviceObjectPropertyReferenceBuilder) WithOptionalDeviceIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetDeviceObjectPropertyReferenceBuilder {
	builder := builderSupplier(m.DeviceIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	m.DeviceIdentifier, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return m
}

func (m *_BACnetDeviceObjectPropertyReferenceBuilder) Build() (BACnetDeviceObjectPropertyReference, error) {
	if m.ObjectIdentifier == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'objectIdentifier' not set"))
	}
	if m.PropertyIdentifier == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'propertyIdentifier' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetDeviceObjectPropertyReference.deepCopy(), nil
}

func (m *_BACnetDeviceObjectPropertyReferenceBuilder) MustBuild() BACnetDeviceObjectPropertyReference {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetDeviceObjectPropertyReferenceBuilder) DeepCopy() any {
	return m.CreateBACnetDeviceObjectPropertyReferenceBuilder()
}

// CreateBACnetDeviceObjectPropertyReferenceBuilder creates a BACnetDeviceObjectPropertyReferenceBuilder
func (m *_BACnetDeviceObjectPropertyReference) CreateBACnetDeviceObjectPropertyReferenceBuilder() BACnetDeviceObjectPropertyReferenceBuilder {
	if m == nil {
		return NewBACnetDeviceObjectPropertyReferenceBuilder()
	}
	return &_BACnetDeviceObjectPropertyReferenceBuilder{_BACnetDeviceObjectPropertyReference: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDeviceObjectPropertyReference) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetDeviceObjectPropertyReference) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetDeviceObjectPropertyReference) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *_BACnetDeviceObjectPropertyReference) GetDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.DeviceIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetDeviceObjectPropertyReference(structType any) BACnetDeviceObjectPropertyReference {
	if casted, ok := structType.(BACnetDeviceObjectPropertyReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDeviceObjectPropertyReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDeviceObjectPropertyReference) GetTypeName() string {
	return "BACnetDeviceObjectPropertyReference"
}

func (m *_BACnetDeviceObjectPropertyReference) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits(ctx)
	}

	// Optional Field (deviceIdentifier)
	if m.DeviceIdentifier != nil {
		lengthInBits += m.DeviceIdentifier.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetDeviceObjectPropertyReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDeviceObjectPropertyReferenceParse(ctx context.Context, theBytes []byte) (BACnetDeviceObjectPropertyReference, error) {
	return BACnetDeviceObjectPropertyReferenceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetDeviceObjectPropertyReferenceParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDeviceObjectPropertyReference, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDeviceObjectPropertyReference, error) {
		return BACnetDeviceObjectPropertyReferenceParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetDeviceObjectPropertyReferenceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDeviceObjectPropertyReference, error) {
	v, err := (&_BACnetDeviceObjectPropertyReference{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetDeviceObjectPropertyReference) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetDeviceObjectPropertyReference BACnetDeviceObjectPropertyReference, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDeviceObjectPropertyReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDeviceObjectPropertyReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	propertyIdentifier, err := ReadSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", ReadComplex[BACnetPropertyIdentifierTagged](BACnetPropertyIdentifierTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyIdentifier' field"))
	}
	m.PropertyIdentifier = propertyIdentifier

	var arrayIndex BACnetContextTagUnsignedInteger
	_arrayIndex, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayIndex' field"))
	}
	if _arrayIndex != nil {
		arrayIndex = *_arrayIndex
		m.ArrayIndex = arrayIndex
	}

	var deviceIdentifier BACnetContextTagObjectIdentifier
	_deviceIdentifier, err := ReadOptionalField[BACnetContextTagObjectIdentifier](ctx, "deviceIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceIdentifier' field"))
	}
	if _deviceIdentifier != nil {
		deviceIdentifier = *_deviceIdentifier
		m.DeviceIdentifier = deviceIdentifier
	}

	if closeErr := readBuffer.CloseContext("BACnetDeviceObjectPropertyReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDeviceObjectPropertyReference")
	}

	return m, nil
}

func (m *_BACnetDeviceObjectPropertyReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDeviceObjectPropertyReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetDeviceObjectPropertyReference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDeviceObjectPropertyReference")
	}

	if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
	}

	if err := WriteSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", m.GetPropertyIdentifier(), WriteComplex[BACnetPropertyIdentifierTagged](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'propertyIdentifier' field")
	}

	if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", GetRef(m.GetArrayIndex()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'arrayIndex' field")
	}

	if err := WriteOptionalField[BACnetContextTagObjectIdentifier](ctx, "deviceIdentifier", GetRef(m.GetDeviceIdentifier()), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'deviceIdentifier' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDeviceObjectPropertyReference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDeviceObjectPropertyReference")
	}
	return nil
}

func (m *_BACnetDeviceObjectPropertyReference) IsBACnetDeviceObjectPropertyReference() {}

func (m *_BACnetDeviceObjectPropertyReference) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetDeviceObjectPropertyReference) deepCopy() *_BACnetDeviceObjectPropertyReference {
	if m == nil {
		return nil
	}
	_BACnetDeviceObjectPropertyReferenceCopy := &_BACnetDeviceObjectPropertyReference{
		m.ObjectIdentifier.DeepCopy().(BACnetContextTagObjectIdentifier),
		m.PropertyIdentifier.DeepCopy().(BACnetPropertyIdentifierTagged),
		m.ArrayIndex.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.DeviceIdentifier.DeepCopy().(BACnetContextTagObjectIdentifier),
	}
	return _BACnetDeviceObjectPropertyReferenceCopy
}

func (m *_BACnetDeviceObjectPropertyReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
