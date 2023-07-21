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
package org.apache.plc4x.java.eip.readwrite;

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

public class CipConnectedResponse extends CipService implements Message {

  // Accessors for discriminator values.
  public Byte getService() {
    return (byte) 0x52;
  }

  public Boolean getResponse() {
    return (boolean) true;
  }

  public Boolean getConnected() {
    return false;
  }

  // Properties.
  protected final short status;
  protected final short additionalStatusWords;
  protected final CIPDataConnected data;

  public CipConnectedResponse(short status, short additionalStatusWords, CIPDataConnected data) {
    super();
    this.status = status;
    this.additionalStatusWords = additionalStatusWords;
    this.data = data;
  }

  public short getStatus() {
    return status;
  }

  public short getAdditionalStatusWords() {
    return additionalStatusWords;
  }

  public CIPDataConnected getData() {
    return data;
  }

  @Override
  protected void serializeCipServiceChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CipConnectedResponse");

    // Reserved Field (reserved)
    writeReservedField("reserved", (short) 0x00, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (status)
    writeSimpleField("status", status, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (additionalStatusWords)
    writeSimpleField(
        "additionalStatusWords", additionalStatusWords, writeUnsignedShort(writeBuffer, 8));

    // Optional Field (data) (Can be skipped, if the value is null)
    writeOptionalField("data", data, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("CipConnectedResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CipConnectedResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 8;

    // Simple field (status)
    lengthInBits += 8;

    // Simple field (additionalStatusWords)
    lengthInBits += 8;

    // Optional Field (data)
    if (data != null) {
      lengthInBits += data.getLengthInBits();
    }

    return lengthInBits;
  }

  public static CipServiceBuilder staticParseCipServiceBuilder(
      ReadBuffer readBuffer, Boolean connected, Integer serviceLen) throws ParseException {
    readBuffer.pullContext("CipConnectedResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Short reservedField0 =
        readReservedField("reserved", readUnsignedShort(readBuffer, 8), (short) 0x00);

    short status = readSimpleField("status", readUnsignedShort(readBuffer, 8));

    short additionalStatusWords =
        readSimpleField("additionalStatusWords", readUnsignedShort(readBuffer, 8));

    CIPDataConnected data =
        readOptionalField(
            "data",
            new DataReaderComplexDefault<>(
                () -> CIPDataConnected.staticParse(readBuffer), readBuffer),
            (((serviceLen) - (4))) > (0));

    readBuffer.closeContext("CipConnectedResponse");
    // Create the instance
    return new CipConnectedResponseBuilderImpl(status, additionalStatusWords, data);
  }

  public static class CipConnectedResponseBuilderImpl implements CipService.CipServiceBuilder {
    private final short status;
    private final short additionalStatusWords;
    private final CIPDataConnected data;

    public CipConnectedResponseBuilderImpl(
        short status, short additionalStatusWords, CIPDataConnected data) {
      this.status = status;
      this.additionalStatusWords = additionalStatusWords;
      this.data = data;
    }

    public CipConnectedResponse build() {
      CipConnectedResponse cipConnectedResponse =
          new CipConnectedResponse(status, additionalStatusWords, data);
      return cipConnectedResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CipConnectedResponse)) {
      return false;
    }
    CipConnectedResponse that = (CipConnectedResponse) o;
    return (getStatus() == that.getStatus())
        && (getAdditionalStatusWords() == that.getAdditionalStatusWords())
        && (getData() == that.getData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getStatus(), getAdditionalStatusWords(), getData());
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
