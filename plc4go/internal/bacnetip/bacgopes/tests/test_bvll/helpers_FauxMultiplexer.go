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

package test_bvll

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comm"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comp"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/pdu"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/vlan"
)

//go:generate plc4xGenerator -type=FauxMultiplexer -prefix=helpers_
type FauxMultiplexer struct {
	DefaultRFormatter `ignore:"true"`
	ClientContract
	ServerContract

	address        *Address
	unicastTuple   *AddressTuple[string, uint16]
	broadcastTuple *AddressTuple[string, uint16]

	node *IPNode

	log zerolog.Logger
}

func NewFauxMultiplexer(localLog zerolog.Logger, addr *Address, network *IPNetwork) (*FauxMultiplexer, error) {
	f := &FauxMultiplexer{
		DefaultRFormatter: NewDefaultRFormatter(),
		address:           addr,
		log:               localLog,
	}
	var err error
	f.ClientContract, err = NewClient(localLog) // TODO: do we need to pass ids?
	if err != nil {
		return nil, errors.Wrap(err, "error creating client")
	}
	f.ServerContract, err = NewServer(localLog) // TODO: do we need to pass ids?
	if err != nil {
		return nil, errors.Wrap(err, "error creating server")
	}

	// get the unicast and broadcast tuples
	f.unicastTuple = addr.AddrTuple
	f.broadcastTuple = addr.AddrBroadcastTuple

	// make an internal node and bind to it, this takes the place of
	// both the direct port and broadcast port of the real UDPMultiplexer
	f.node, err = NewIPNode(localLog, addr, network)
	if err != nil {
		return nil, errors.Wrap(err, "error creating ip node")
	}
	if err := Bind(localLog, f, f.node); err != nil {
		return nil, errors.Wrap(err, "error binding")
	}
	return f, nil
}

func (s *FauxMultiplexer) Indication(args Args, kwargs KWArgs) error {
	s.log.Debug().Stringer("Args", args).Stringer("KWArgs", kwargs).Msg("Indication")

	pdu := GA[PDU](args, 0)

	var dest *Address
	// check for a broadcast message
	if pdu.GetPDUDestination().AddrType == LOCAL_BROADCAST_ADDRESS {
		var err error
		dest, err = NewAddress(NA(NA(s.broadcastTuple)))
		if err != nil {
			return errors.Wrap(err, "error creating address")
		}
		s.log.Debug().Stringer("dest", dest).Msg("Requesting local broadcast")
	} else if pdu.GetPDUDestination().AddrType == LOCAL_STATION_ADDRESS {
		var err error
		dest, err = NewAddress(NA(NA(pdu.GetPDUDestination().AddrAddress)))
		if err != nil {
			return errors.Wrap(err, "error creating address")
		}
		s.log.Debug().Stringer("dest", dest).Msg("Requesting local station")
	} else {
		return errors.New("unknown destination type")
	}

	unicast, err := NewAddress(NA(NA(s.unicastTuple)))
	if err != nil {
		return errors.Wrap(err, "error creating address")
	}
	return s.Request(NA(NewPDU(NoArgs, NKW(KWCompRootMessage, pdu, KWCPCISource, unicast, KWCPCIDestination, dest))), NoKWArgs)
}

func (s *FauxMultiplexer) Confirmation(args Args, kwargs KWArgs) error {
	s.log.Debug().Stringer("Args", args).Stringer("KWArgs", kwargs).Msg("Indication")
	pdu := GA[PDU](args, 0)

	// the PDU source and destination are tuples, convert them to Address instances
	src := pdu.GetPDUSource()

	broadcast, err := NewAddress(NA(s.broadcastTuple))
	if err != nil {
		return errors.Wrap(err, "error creating address")
	}
	var dest *Address
	// see if the destination was our broadcast address
	if pdu.GetPDUDestination().Equals(broadcast) {
		dest = NewLocalBroadcast(nil)
	} else {
		dest, err = NewAddress(NA(pdu.GetPDUDestination().AddrAddress))
		if err != nil {
			return errors.Wrap(err, "error creating address")
		}
	}

	return s.Response(NA(NewPDU(NoArgs, NKW(KWCompRootMessage, pdu, KWCPCISource, src, KWCPCIDestination, dest))), NoKWArgs)
}
