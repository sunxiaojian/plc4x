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

public abstract class CipService implements Message {

  // Abstract accessors for discriminator values.
  public abstract Boolean getConnected();

  public abstract Boolean getResponse();

  public abstract Byte getService();

  public CipService() {
    super();
  }

  protected abstract void serializeCipServiceChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CipService");

    // Discriminator Field (response) (Used as input to a switch field)
    writeDiscriminatorField("response", getResponse(), writeBoolean(writeBuffer));

    // Discriminator Field (service) (Used as input to a switch field)
    writeDiscriminatorField("service", getService(), writeUnsignedByte(writeBuffer, 7));

    // Switch field (Serialize the sub-type)
    serializeCipServiceChild(writeBuffer);

    writeBuffer.popContext("CipService");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    CipService _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (response)
    lengthInBits += 1;

    // Discriminator Field (service)
    lengthInBits += 7;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static CipService staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 2)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 2, but got " + args.length);
    }
    Boolean connected;
    if (args[0] instanceof Boolean) {
      connected = (Boolean) args[0];
    } else if (args[0] instanceof String) {
      connected = Boolean.valueOf((String) args[0]);
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type Boolean or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    Integer serviceLen;
    if (args[1] instanceof Integer) {
      serviceLen = (Integer) args[1];
    } else if (args[1] instanceof String) {
      serviceLen = Integer.valueOf((String) args[1]);
    } else {
      throw new PlcRuntimeException(
          "Argument 1 expected to be of type Integer or a string which is parseable but was "
              + args[1].getClass().getName());
    }
    return staticParse(readBuffer, connected, serviceLen);
  }

  public static CipService staticParse(ReadBuffer readBuffer, Boolean connected, Integer serviceLen)
      throws ParseException {
    readBuffer.pullContext("CipService");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    boolean response = readDiscriminatorField("response", readBoolean(readBuffer));

    byte service = readDiscriminatorField("service", readUnsignedByte(readBuffer, 7));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    CipServiceBuilder builder = null;
    if (EvaluationHelper.equals(service, (byte) 0x01)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          GetAttributeAllRequest.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x01)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          GetAttributeAllResponse.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x02)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          SetAttributeAllRequest.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x02)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          SetAttributeAllResponse.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x03)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          GetAttributeListRequest.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x03)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          GetAttributeListResponse.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x04)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          SetAttributeListRequest.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x04)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          SetAttributeListResponse.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x0A)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          MultipleServiceRequest.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x0A)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          MultipleServiceResponse.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x0E)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          GetAttributeSingleRequest.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x0E)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          GetAttributeSingleResponse.staticParseCipServiceBuilder(
              readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x10)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          SetAttributeSingleRequest.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x10)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          SetAttributeSingleResponse.staticParseCipServiceBuilder(
              readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x4C)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = CipReadRequest.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x4C)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = CipReadResponse.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x4D)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder = CipWriteRequest.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x4D)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder = CipWriteResponse.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x4E)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          CipConnectionManagerCloseRequest.staticParseCipServiceBuilder(
              readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x4E)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          CipConnectionManagerCloseResponse.staticParseCipServiceBuilder(
              readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x52)
        && EvaluationHelper.equals(response, (boolean) false)
        && EvaluationHelper.equals(connected, (boolean) false)) {
      builder =
          CipUnconnectedRequest.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x52)
        && EvaluationHelper.equals(response, (boolean) false)
        && EvaluationHelper.equals(connected, (boolean) true)) {
      builder = CipConnectedRequest.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x52)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          CipConnectedResponse.staticParseCipServiceBuilder(readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x5B)
        && EvaluationHelper.equals(response, (boolean) false)) {
      builder =
          CipConnectionManagerRequest.staticParseCipServiceBuilder(
              readBuffer, connected, serviceLen);
    } else if (EvaluationHelper.equals(service, (byte) 0x5B)
        && EvaluationHelper.equals(response, (boolean) true)) {
      builder =
          CipConnectionManagerResponse.staticParseCipServiceBuilder(
              readBuffer, connected, serviceLen);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "service="
              + service
              + " "
              + "response="
              + response
              + " "
              + "connected="
              + connected
              + "]");
    }

    readBuffer.closeContext("CipService");
    // Create the instance
    CipService _cipService = builder.build();
    return _cipService;
  }

  public interface CipServiceBuilder {
    CipService build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CipService)) {
      return false;
    }
    CipService that = (CipService) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
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
