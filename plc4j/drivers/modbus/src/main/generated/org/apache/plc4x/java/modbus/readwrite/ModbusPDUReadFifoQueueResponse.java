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
package org.apache.plc4x.java.modbus.readwrite;

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

public class ModbusPDUReadFifoQueueResponse extends ModbusPDU implements Message {

  // Accessors for discriminator values.
  public Boolean getErrorFlag() {
    return (boolean) false;
  }

  public Byte getFunctionFlag() {
    return (byte) 0x18;
  }

  public Boolean getResponse() {
    return (boolean) true;
  }

  // Properties.
  protected final List<Integer> fifoValue;

  public ModbusPDUReadFifoQueueResponse(List<Integer> fifoValue) {
    super();
    this.fifoValue = fifoValue;
  }

  public List<Integer> getFifoValue() {
    return fifoValue;
  }

  @Override
  protected void serializeModbusPDUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ModbusPDUReadFifoQueueResponse");

    // Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    int byteCount = (int) ((((COUNT(getFifoValue())) * (2))) + (2));
    writeImplicitField("byteCount", byteCount, writeUnsignedInt(writeBuffer, 16));

    // Implicit Field (fifoCount) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    int fifoCount = (int) ((((COUNT(getFifoValue())) * (2))) / (2));
    writeImplicitField("fifoCount", fifoCount, writeUnsignedInt(writeBuffer, 16));

    // Array Field (fifoValue)
    writeSimpleTypeArrayField("fifoValue", fifoValue, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("ModbusPDUReadFifoQueueResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ModbusPDUReadFifoQueueResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (byteCount)
    lengthInBits += 16;

    // Implicit Field (fifoCount)
    lengthInBits += 16;

    // Array field
    if (fifoValue != null) {
      lengthInBits += 16 * fifoValue.size();
    }

    return lengthInBits;
  }

  public static ModbusPDUBuilder staticParseModbusPDUBuilder(
      ReadBuffer readBuffer, Boolean response) throws ParseException {
    readBuffer.pullContext("ModbusPDUReadFifoQueueResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int byteCount = readImplicitField("byteCount", readUnsignedInt(readBuffer, 16));

    int fifoCount = readImplicitField("fifoCount", readUnsignedInt(readBuffer, 16));

    List<Integer> fifoValue =
        readCountArrayField("fifoValue", readUnsignedInt(readBuffer, 16), fifoCount);

    readBuffer.closeContext("ModbusPDUReadFifoQueueResponse");
    // Create the instance
    return new ModbusPDUReadFifoQueueResponseBuilderImpl(fifoValue);
  }

  public static class ModbusPDUReadFifoQueueResponseBuilderImpl
      implements ModbusPDU.ModbusPDUBuilder {
    private final List<Integer> fifoValue;

    public ModbusPDUReadFifoQueueResponseBuilderImpl(List<Integer> fifoValue) {
      this.fifoValue = fifoValue;
    }

    public ModbusPDUReadFifoQueueResponse build() {
      ModbusPDUReadFifoQueueResponse modbusPDUReadFifoQueueResponse =
          new ModbusPDUReadFifoQueueResponse(fifoValue);
      return modbusPDUReadFifoQueueResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ModbusPDUReadFifoQueueResponse)) {
      return false;
    }
    ModbusPDUReadFifoQueueResponse that = (ModbusPDUReadFifoQueueResponse) o;
    return (getFifoValue() == that.getFifoValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getFifoValue());
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
