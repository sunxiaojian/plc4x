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

// BACnetConstructedDataIPv6DHCPLeaseTimeRemaining is the corresponding interface of BACnetConstructedDataIPv6DHCPLeaseTimeRemaining
type BACnetConstructedDataIPv6DHCPLeaseTimeRemaining interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetIpv6DhcpLeaseTimeRemaining returns Ipv6DhcpLeaseTimeRemaining (property field)
	GetIpv6DhcpLeaseTimeRemaining() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataIPv6DHCPLeaseTimeRemaining is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIPv6DHCPLeaseTimeRemaining()
	// CreateBuilder creates a BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder
	CreateBACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder() BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder
}

// _BACnetConstructedDataIPv6DHCPLeaseTimeRemaining is the data-structure of this message
type _BACnetConstructedDataIPv6DHCPLeaseTimeRemaining struct {
	BACnetConstructedDataContract
	Ipv6DhcpLeaseTimeRemaining BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataIPv6DHCPLeaseTimeRemaining = (*_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining)(nil)

// NewBACnetConstructedDataIPv6DHCPLeaseTimeRemaining factory function for _BACnetConstructedDataIPv6DHCPLeaseTimeRemaining
func NewBACnetConstructedDataIPv6DHCPLeaseTimeRemaining(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, ipv6DhcpLeaseTimeRemaining BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining {
	if ipv6DhcpLeaseTimeRemaining == nil {
		panic("ipv6DhcpLeaseTimeRemaining of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataIPv6DHCPLeaseTimeRemaining must not be nil")
	}
	_result := &_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Ipv6DhcpLeaseTimeRemaining:    ipv6DhcpLeaseTimeRemaining,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder is a builder for BACnetConstructedDataIPv6DHCPLeaseTimeRemaining
type BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(ipv6DhcpLeaseTimeRemaining BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder
	// WithIpv6DhcpLeaseTimeRemaining adds Ipv6DhcpLeaseTimeRemaining (property field)
	WithIpv6DhcpLeaseTimeRemaining(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder
	// WithIpv6DhcpLeaseTimeRemainingBuilder adds Ipv6DhcpLeaseTimeRemaining (property field) which is build by the builder
	WithIpv6DhcpLeaseTimeRemainingBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder
	// Build builds the BACnetConstructedDataIPv6DHCPLeaseTimeRemaining or returns an error if something is wrong
	Build() (BACnetConstructedDataIPv6DHCPLeaseTimeRemaining, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIPv6DHCPLeaseTimeRemaining
}

// NewBACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder() creates a BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder
func NewBACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder() BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder {
	return &_BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder{_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining: new(_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining)}
}

type _BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder struct {
	*_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining

	err *utils.MultiError
}

var _ (BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder) = (*_BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder)(nil)

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder) WithMandatoryFields(ipv6DhcpLeaseTimeRemaining BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder {
	return m.WithIpv6DhcpLeaseTimeRemaining(ipv6DhcpLeaseTimeRemaining)
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder) WithIpv6DhcpLeaseTimeRemaining(ipv6DhcpLeaseTimeRemaining BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder {
	m.Ipv6DhcpLeaseTimeRemaining = ipv6DhcpLeaseTimeRemaining
	return m
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder) WithIpv6DhcpLeaseTimeRemainingBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder {
	builder := builderSupplier(m.Ipv6DhcpLeaseTimeRemaining.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	m.Ipv6DhcpLeaseTimeRemaining, err = builder.Build()
	if err != nil {
		if m.err == nil {
			m.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		m.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return m
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder) Build() (BACnetConstructedDataIPv6DHCPLeaseTimeRemaining, error) {
	if m.Ipv6DhcpLeaseTimeRemaining == nil {
		if m.err == nil {
			m.err = new(utils.MultiError)
		}
		m.err.Append(errors.New("mandatory field 'ipv6DhcpLeaseTimeRemaining' not set"))
	}
	if m.err != nil {
		return nil, errors.Wrap(m.err, "error occurred during build")
	}
	return m._BACnetConstructedDataIPv6DHCPLeaseTimeRemaining.deepCopy(), nil
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder) MustBuild() BACnetConstructedDataIPv6DHCPLeaseTimeRemaining {
	build, err := m.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder) DeepCopy() any {
	return m.CreateBACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder()
}

// CreateBACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder creates a BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder
func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) CreateBACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder() BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder {
	if m == nil {
		return NewBACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder()
	}
	return &_BACnetConstructedDataIPv6DHCPLeaseTimeRemainingBuilder{_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining: m.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_DHCP_LEASE_TIME_REMAINING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetIpv6DhcpLeaseTimeRemaining() BACnetApplicationTagUnsignedInteger {
	return m.Ipv6DhcpLeaseTimeRemaining
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetIpv6DhcpLeaseTimeRemaining())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPv6DHCPLeaseTimeRemaining(structType any) BACnetConstructedDataIPv6DHCPLeaseTimeRemaining {
	if casted, ok := structType.(BACnetConstructedDataIPv6DHCPLeaseTimeRemaining); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6DHCPLeaseTimeRemaining); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetTypeName() string {
	return "BACnetConstructedDataIPv6DHCPLeaseTimeRemaining"
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (ipv6DhcpLeaseTimeRemaining)
	lengthInBits += m.Ipv6DhcpLeaseTimeRemaining.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIPv6DHCPLeaseTimeRemaining BACnetConstructedDataIPv6DHCPLeaseTimeRemaining, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6DHCPLeaseTimeRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6DHCPLeaseTimeRemaining")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ipv6DhcpLeaseTimeRemaining, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "ipv6DhcpLeaseTimeRemaining", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ipv6DhcpLeaseTimeRemaining' field"))
	}
	m.Ipv6DhcpLeaseTimeRemaining = ipv6DhcpLeaseTimeRemaining

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), ipv6DhcpLeaseTimeRemaining)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6DHCPLeaseTimeRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6DHCPLeaseTimeRemaining")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6DHCPLeaseTimeRemaining"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6DHCPLeaseTimeRemaining")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "ipv6DhcpLeaseTimeRemaining", m.GetIpv6DhcpLeaseTimeRemaining(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ipv6DhcpLeaseTimeRemaining' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6DHCPLeaseTimeRemaining"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6DHCPLeaseTimeRemaining")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) IsBACnetConstructedDataIPv6DHCPLeaseTimeRemaining() {
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) deepCopy() *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIPv6DHCPLeaseTimeRemainingCopy := &_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.Ipv6DhcpLeaseTimeRemaining.DeepCopy().(BACnetApplicationTagUnsignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIPv6DHCPLeaseTimeRemainingCopy
}

func (m *_BACnetConstructedDataIPv6DHCPLeaseTimeRemaining) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
