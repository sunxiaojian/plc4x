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
package org.apache.plc4x.java.cbus.readwrite;

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

public class CALDataStatusExtended extends CALData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final StatusCoding coding;
  protected final ApplicationIdContainer application;
  protected final short blockStart;
  protected final List<StatusByte> statusBytes;
  protected final List<LevelInformation> levelInformation;

  // Arguments.
  protected final RequestContext requestContext;

  public CALDataStatusExtended(
      CALCommandTypeContainer commandTypeContainer,
      CALData additionalData,
      StatusCoding coding,
      ApplicationIdContainer application,
      short blockStart,
      List<StatusByte> statusBytes,
      List<LevelInformation> levelInformation,
      RequestContext requestContext) {
    super(commandTypeContainer, additionalData, requestContext);
    this.coding = coding;
    this.application = application;
    this.blockStart = blockStart;
    this.statusBytes = statusBytes;
    this.levelInformation = levelInformation;
    this.requestContext = requestContext;
  }

  public StatusCoding getCoding() {
    return coding;
  }

  public ApplicationIdContainer getApplication() {
    return application;
  }

  public short getBlockStart() {
    return blockStart;
  }

  public List<StatusByte> getStatusBytes() {
    return statusBytes;
  }

  public List<LevelInformation> getLevelInformation() {
    return levelInformation;
  }

  public byte getNumberOfStatusBytes() {
    return (byte)
        ((((((getCoding()) == (StatusCoding.BINARY_BY_THIS_SERIAL_INTERFACE))
                || ((getCoding()) == (StatusCoding.BINARY_BY_ELSEWHERE))))
            ? ((commandTypeContainer.getNumBytes()) - (3))
            : (0)));
  }

  public byte getNumberOfLevelInformation() {
    return (byte)
        ((((((getCoding()) == (StatusCoding.LEVEL_BY_THIS_SERIAL_INTERFACE))
                || ((getCoding()) == (StatusCoding.LEVEL_BY_ELSEWHERE))))
            ? ((((commandTypeContainer.getNumBytes()) - (3))) / (2))
            : (0)));
  }

  @Override
  protected void serializeCALDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CALDataStatusExtended");

    // Simple Field (coding)
    writeSimpleEnumField(
        "coding",
        "StatusCoding",
        coding,
        new DataWriterEnumDefault<>(
            StatusCoding::getValue, StatusCoding::name, writeByte(writeBuffer, 8)));

    // Simple Field (application)
    writeSimpleEnumField(
        "application",
        "ApplicationIdContainer",
        application,
        new DataWriterEnumDefault<>(
            ApplicationIdContainer::getValue,
            ApplicationIdContainer::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (blockStart)
    writeSimpleField("blockStart", blockStart, writeUnsignedShort(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    byte numberOfStatusBytes = getNumberOfStatusBytes();
    writeBuffer.writeVirtual("numberOfStatusBytes", numberOfStatusBytes);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    byte numberOfLevelInformation = getNumberOfLevelInformation();
    writeBuffer.writeVirtual("numberOfLevelInformation", numberOfLevelInformation);

    // Array Field (statusBytes)
    writeComplexTypeArrayField("statusBytes", statusBytes, writeBuffer);

    // Array Field (levelInformation)
    writeComplexTypeArrayField("levelInformation", levelInformation, writeBuffer);

    writeBuffer.popContext("CALDataStatusExtended");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CALDataStatusExtended _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (coding)
    lengthInBits += 8;

    // Simple field (application)
    lengthInBits += 8;

    // Simple field (blockStart)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // Array field
    if (statusBytes != null) {
      int i = 0;
      for (StatusByte element : statusBytes) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= statusBytes.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Array field
    if (levelInformation != null) {
      int i = 0;
      for (LevelInformation element : levelInformation) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= levelInformation.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static CALDataBuilder staticParseCALDataBuilder(
      ReadBuffer readBuffer,
      CALCommandTypeContainer commandTypeContainer,
      RequestContext requestContext)
      throws ParseException {
    readBuffer.pullContext("CALDataStatusExtended");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    StatusCoding coding =
        readEnumField(
            "coding",
            "StatusCoding",
            new DataReaderEnumDefault<>(StatusCoding::enumForValue, readByte(readBuffer, 8)));

    ApplicationIdContainer application =
        readEnumField(
            "application",
            "ApplicationIdContainer",
            new DataReaderEnumDefault<>(
                ApplicationIdContainer::enumForValue, readUnsignedShort(readBuffer, 8)));

    short blockStart = readSimpleField("blockStart", readUnsignedShort(readBuffer, 8));
    byte numberOfStatusBytes =
        readVirtualField(
            "numberOfStatusBytes",
            byte.class,
            (((((coding) == (StatusCoding.BINARY_BY_THIS_SERIAL_INTERFACE))
                    || ((coding) == (StatusCoding.BINARY_BY_ELSEWHERE))))
                ? ((commandTypeContainer.getNumBytes()) - (3))
                : (0)));
    byte numberOfLevelInformation =
        readVirtualField(
            "numberOfLevelInformation",
            byte.class,
            (((((coding) == (StatusCoding.LEVEL_BY_THIS_SERIAL_INTERFACE))
                    || ((coding) == (StatusCoding.LEVEL_BY_ELSEWHERE))))
                ? ((((commandTypeContainer.getNumBytes()) - (3))) / (2))
                : (0)));

    List<StatusByte> statusBytes =
        readCountArrayField(
            "statusBytes",
            new DataReaderComplexDefault<>(() -> StatusByte.staticParse(readBuffer), readBuffer),
            numberOfStatusBytes);

    List<LevelInformation> levelInformation =
        readCountArrayField(
            "levelInformation",
            new DataReaderComplexDefault<>(
                () -> LevelInformation.staticParse(readBuffer), readBuffer),
            numberOfLevelInformation);

    readBuffer.closeContext("CALDataStatusExtended");
    // Create the instance
    return new CALDataStatusExtendedBuilderImpl(
        coding, application, blockStart, statusBytes, levelInformation, requestContext);
  }

  public static class CALDataStatusExtendedBuilderImpl implements CALData.CALDataBuilder {
    private final StatusCoding coding;
    private final ApplicationIdContainer application;
    private final short blockStart;
    private final List<StatusByte> statusBytes;
    private final List<LevelInformation> levelInformation;
    private final RequestContext requestContext;

    public CALDataStatusExtendedBuilderImpl(
        StatusCoding coding,
        ApplicationIdContainer application,
        short blockStart,
        List<StatusByte> statusBytes,
        List<LevelInformation> levelInformation,
        RequestContext requestContext) {
      this.coding = coding;
      this.application = application;
      this.blockStart = blockStart;
      this.statusBytes = statusBytes;
      this.levelInformation = levelInformation;
      this.requestContext = requestContext;
    }

    public CALDataStatusExtended build(
        CALCommandTypeContainer commandTypeContainer,
        CALData additionalData,
        RequestContext requestContext) {
      CALDataStatusExtended cALDataStatusExtended =
          new CALDataStatusExtended(
              commandTypeContainer,
              additionalData,
              coding,
              application,
              blockStart,
              statusBytes,
              levelInformation,
              requestContext);
      return cALDataStatusExtended;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CALDataStatusExtended)) {
      return false;
    }
    CALDataStatusExtended that = (CALDataStatusExtended) o;
    return (getCoding() == that.getCoding())
        && (getApplication() == that.getApplication())
        && (getBlockStart() == that.getBlockStart())
        && (getStatusBytes() == that.getStatusBytes())
        && (getLevelInformation() == that.getLevelInformation())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getCoding(),
        getApplication(),
        getBlockStart(),
        getStatusBytes(),
        getLevelInformation());
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
