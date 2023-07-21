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

public class AnsiExtendedSymbolSegment extends DataSegmentType implements Message {

  // Accessors for discriminator values.
  public Byte getDataSegmentType() {
    return (byte) 0x11;
  }

  // Properties.
  protected final String symbol;
  protected final Short pad;

  public AnsiExtendedSymbolSegment(String symbol, Short pad) {
    super();
    this.symbol = symbol;
    this.pad = pad;
  }

  public String getSymbol() {
    return symbol;
  }

  public Short getPad() {
    return pad;
  }

  @Override
  protected void serializeDataSegmentTypeChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("AnsiExtendedSymbolSegment");

    // Implicit Field (dataSize) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    short dataSize = (short) (STR_LEN(getSymbol()));
    writeImplicitField("dataSize", dataSize, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (symbol)
    writeSimpleField("symbol", symbol, writeString(writeBuffer, (dataSize) * (8)));

    // Optional Field (pad) (Can be skipped, if the value is null)
    writeOptionalField("pad", pad, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("AnsiExtendedSymbolSegment");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AnsiExtendedSymbolSegment _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (dataSize)
    lengthInBits += 8;

    // Simple field (symbol)
    lengthInBits += (STR_LEN(getSymbol())) * (8);

    // Optional Field (pad)
    if (pad != null) {
      lengthInBits += 8;
    }

    return lengthInBits;
  }

  public static DataSegmentTypeBuilder staticParseDataSegmentTypeBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("AnsiExtendedSymbolSegment");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short dataSize = readImplicitField("dataSize", readUnsignedShort(readBuffer, 8));

    String symbol = readSimpleField("symbol", readString(readBuffer, (dataSize) * (8)));

    Short pad =
        readOptionalField(
            "pad", readUnsignedShort(readBuffer, 8), ((STR_LEN(symbol)) % (2)) != (0));

    readBuffer.closeContext("AnsiExtendedSymbolSegment");
    // Create the instance
    return new AnsiExtendedSymbolSegmentBuilderImpl(symbol, pad);
  }

  public static class AnsiExtendedSymbolSegmentBuilderImpl
      implements DataSegmentType.DataSegmentTypeBuilder {
    private final String symbol;
    private final Short pad;

    public AnsiExtendedSymbolSegmentBuilderImpl(String symbol, Short pad) {
      this.symbol = symbol;
      this.pad = pad;
    }

    public AnsiExtendedSymbolSegment build() {
      AnsiExtendedSymbolSegment ansiExtendedSymbolSegment =
          new AnsiExtendedSymbolSegment(symbol, pad);
      return ansiExtendedSymbolSegment;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AnsiExtendedSymbolSegment)) {
      return false;
    }
    AnsiExtendedSymbolSegment that = (AnsiExtendedSymbolSegment) o;
    return (getSymbol() == that.getSymbol())
        && (getPad() == that.getPad())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getSymbol(), getPad());
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
