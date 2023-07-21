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
package org.apache.plc4x.java.openprotocol.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public abstract class OpenProtocolMessageLinkLevelPositiveAcknowledge extends OpenProtocolMessage
    implements Message {

  // Accessors for discriminator values.
  public Mid getMid() {
    return Mid.LinkLevelPositiveAcknowledge;
  }

  // Abstract accessors for discriminator values.
  public abstract Integer getRevision();

  public OpenProtocolMessageLinkLevelPositiveAcknowledge(
      Integer midRevision,
      Short noAckFlag,
      Integer targetStationId,
      Integer targetSpindleId,
      Integer sequenceNumber,
      Short numberOfMessageParts,
      Short messagePartNumber) {
    super(
        midRevision,
        noAckFlag,
        targetStationId,
        targetSpindleId,
        sequenceNumber,
        numberOfMessageParts,
        messagePartNumber);
  }

  protected abstract void serializeOpenProtocolMessageLinkLevelPositiveAcknowledgeChild(
      WriteBuffer writeBuffer) throws SerializationException;

  @Override
  protected void serializeOpenProtocolMessageChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("OpenProtocolMessageLinkLevelPositiveAcknowledge");

    // Switch field (Serialize the sub-type)
    serializeOpenProtocolMessageLinkLevelPositiveAcknowledgeChild(writeBuffer);

    writeBuffer.popContext("OpenProtocolMessageLinkLevelPositiveAcknowledge");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpenProtocolMessageLinkLevelPositiveAcknowledge _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static OpenProtocolMessageBuilder staticParseOpenProtocolMessageBuilder(
      ReadBuffer readBuffer, Integer revision) throws ParseException {
    readBuffer.pullContext("OpenProtocolMessageLinkLevelPositiveAcknowledge");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    OpenProtocolMessageLinkLevelPositiveAcknowledgeBuilder builder = null;
    if (EvaluationHelper.equals(revision, (int) 1)) {
      builder =
          OpenProtocolMessageLinkLevelPositiveAcknowledgeRev1
              .staticParseOpenProtocolMessageLinkLevelPositiveAcknowledgeBuilder(
                  readBuffer, revision);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "revision="
              + revision
              + "]");
    }

    readBuffer.closeContext("OpenProtocolMessageLinkLevelPositiveAcknowledge");
    // Create the instance
    return new OpenProtocolMessageLinkLevelPositiveAcknowledgeBuilderImpl(builder);
  }

  public interface OpenProtocolMessageLinkLevelPositiveAcknowledgeBuilder {
    OpenProtocolMessageLinkLevelPositiveAcknowledge build(
        Integer midRevision,
        Short noAckFlag,
        Integer targetStationId,
        Integer targetSpindleId,
        Integer sequenceNumber,
        Short numberOfMessageParts,
        Short messagePartNumber);
  }

  public static class OpenProtocolMessageLinkLevelPositiveAcknowledgeBuilderImpl
      implements OpenProtocolMessage.OpenProtocolMessageBuilder {
    private final OpenProtocolMessageLinkLevelPositiveAcknowledgeBuilder builder;

    public OpenProtocolMessageLinkLevelPositiveAcknowledgeBuilderImpl(
        OpenProtocolMessageLinkLevelPositiveAcknowledgeBuilder builder) {
      this.builder = builder;
    }

    public OpenProtocolMessageLinkLevelPositiveAcknowledge build(
        Integer midRevision,
        Short noAckFlag,
        Integer targetStationId,
        Integer targetSpindleId,
        Integer sequenceNumber,
        Short numberOfMessageParts,
        Short messagePartNumber) {
      return builder.build(
          midRevision,
          noAckFlag,
          targetStationId,
          targetSpindleId,
          sequenceNumber,
          numberOfMessageParts,
          messagePartNumber);
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenProtocolMessageLinkLevelPositiveAcknowledge)) {
      return false;
    }
    OpenProtocolMessageLinkLevelPositiveAcknowledge that =
        (OpenProtocolMessageLinkLevelPositiveAcknowledge) o;
    return super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
