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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public class ApduDataExtAuthorizeRequest extends ApduDataExt implements Message {

  // Accessors for discriminator values.
  public Byte getExtApciType() {
    return (byte) 0x11;
  }

  // Properties.
  protected final short level;
  protected final byte[] data;

  public ApduDataExtAuthorizeRequest(short level, byte[] data) {
    super();
    this.level = level;
    this.data = data;
  }

  public short getLevel() {
    return level;
  }

  public byte[] getData() {
    return data;
  }

  @Override
  protected void serializeApduDataExtChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ApduDataExtAuthorizeRequest");

    // Simple Field (level)
    writeSimpleField("level", level, writeUnsignedShort(writeBuffer, 8));

    // Array Field (data)
    writeByteArrayField("data", data, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("ApduDataExtAuthorizeRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ApduDataExtAuthorizeRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (level)
    lengthInBits += 8;

    // Array field
    if (data != null) {
      lengthInBits += 8 * data.length;
    }

    return lengthInBits;
  }

  public static ApduDataExtBuilder staticParseApduDataExtBuilder(
      ReadBuffer readBuffer, Short length) throws ParseException {
    readBuffer.pullContext("ApduDataExtAuthorizeRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short level = readSimpleField("level", readUnsignedShort(readBuffer, 8));

    byte[] data = readBuffer.readByteArray("data", Math.toIntExact(4));

    readBuffer.closeContext("ApduDataExtAuthorizeRequest");
    // Create the instance
    return new ApduDataExtAuthorizeRequestBuilderImpl(level, data);
  }

  public static class ApduDataExtAuthorizeRequestBuilderImpl
      implements ApduDataExt.ApduDataExtBuilder {
    private final short level;
    private final byte[] data;

    public ApduDataExtAuthorizeRequestBuilderImpl(short level, byte[] data) {
      this.level = level;
      this.data = data;
    }

    public ApduDataExtAuthorizeRequest build() {
      ApduDataExtAuthorizeRequest apduDataExtAuthorizeRequest =
          new ApduDataExtAuthorizeRequest(level, data);
      return apduDataExtAuthorizeRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ApduDataExtAuthorizeRequest)) {
      return false;
    }
    ApduDataExtAuthorizeRequest that = (ApduDataExtAuthorizeRequest) o;
    return (getLevel() == that.getLevel())
        && (getData() == that.getData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getLevel(), getData());
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
