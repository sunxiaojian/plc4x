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

public class OpenProtocolMessageAlarmRev1 extends OpenProtocolMessageAlarm implements Message {

  // Accessors for discriminator values.
  public Integer getRevision() {
    return (int) 1;
  }

  // Constant values.
  public static final Integer BLOCKIDERRORCODE = 1;
  public static final Integer BLOCKIDCONTROLLERREADYSTATUS = 2;
  public static final Integer BLOCKIDTOOLREADYSTATUS = 3;
  public static final Integer BLOCKIDTIME = 4;

  // Properties.
  protected final String errorCode;
  protected final NokOk controllerReadyStatus;
  protected final NokOk toolReadyStatus;
  protected final String alarmTime;

  public OpenProtocolMessageAlarmRev1(
      Integer midRevision,
      Short noAckFlag,
      Integer targetStationId,
      Integer targetSpindleId,
      Integer sequenceNumber,
      Short numberOfMessageParts,
      Short messagePartNumber,
      String errorCode,
      NokOk controllerReadyStatus,
      NokOk toolReadyStatus,
      String alarmTime) {
    super(
        midRevision,
        noAckFlag,
        targetStationId,
        targetSpindleId,
        sequenceNumber,
        numberOfMessageParts,
        messagePartNumber);
    this.errorCode = errorCode;
    this.controllerReadyStatus = controllerReadyStatus;
    this.toolReadyStatus = toolReadyStatus;
    this.alarmTime = alarmTime;
  }

  public String getErrorCode() {
    return errorCode;
  }

  public NokOk getControllerReadyStatus() {
    return controllerReadyStatus;
  }

  public NokOk getToolReadyStatus() {
    return toolReadyStatus;
  }

  public String getAlarmTime() {
    return alarmTime;
  }

  public int getBlockIdErrorCode() {
    return BLOCKIDERRORCODE;
  }

  public int getBlockIdControllerReadyStatus() {
    return BLOCKIDCONTROLLERREADYSTATUS;
  }

  public int getBlockIdToolReadyStatus() {
    return BLOCKIDTOOLREADYSTATUS;
  }

  public int getBlockIdTime() {
    return BLOCKIDTIME;
  }

  @Override
  protected void serializeOpenProtocolMessageAlarmChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("OpenProtocolMessageAlarmRev1");

    // Const Field (blockIdErrorCode)
    writeConstField(
        "blockIdErrorCode",
        BLOCKIDERRORCODE,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (errorCode)
    writeSimpleField(
        "errorCode", errorCode, writeString(writeBuffer, 32), WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdControllerReadyStatus)
    writeConstField(
        "blockIdControllerReadyStatus",
        BLOCKIDCONTROLLERREADYSTATUS,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (controllerReadyStatus)
    writeSimpleEnumField(
        "controllerReadyStatus",
        "NokOk",
        controllerReadyStatus,
        new DataWriterEnumDefault<>(
            NokOk::getValue, NokOk::name, writeUnsignedShort(writeBuffer, 8)),
        WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdToolReadyStatus)
    writeConstField(
        "blockIdToolReadyStatus",
        BLOCKIDTOOLREADYSTATUS,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (toolReadyStatus)
    writeSimpleEnumField(
        "toolReadyStatus",
        "NokOk",
        toolReadyStatus,
        new DataWriterEnumDefault<>(
            NokOk::getValue, NokOk::name, writeUnsignedShort(writeBuffer, 8)),
        WithOption.WithEncoding("ASCII"));

    // Const Field (blockIdTime)
    writeConstField(
        "blockIdTime",
        BLOCKIDTIME,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (alarmTime)
    writeSimpleField(
        "alarmTime", alarmTime, writeString(writeBuffer, 152), WithOption.WithEncoding("ASCII"));

    writeBuffer.popContext("OpenProtocolMessageAlarmRev1");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpenProtocolMessageAlarmRev1 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (blockIdErrorCode)
    lengthInBits += 16;

    // Simple field (errorCode)
    lengthInBits += 32;

    // Const Field (blockIdControllerReadyStatus)
    lengthInBits += 16;

    // Simple field (controllerReadyStatus)
    lengthInBits += 8;

    // Const Field (blockIdToolReadyStatus)
    lengthInBits += 16;

    // Simple field (toolReadyStatus)
    lengthInBits += 8;

    // Const Field (blockIdTime)
    lengthInBits += 16;

    // Simple field (alarmTime)
    lengthInBits += 152;

    return lengthInBits;
  }

  public static OpenProtocolMessageAlarmBuilder staticParseOpenProtocolMessageAlarmBuilder(
      ReadBuffer readBuffer, Integer revision) throws ParseException {
    readBuffer.pullContext("OpenProtocolMessageAlarmRev1");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int blockIdErrorCode =
        readConstField(
            "blockIdErrorCode",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageAlarmRev1.BLOCKIDERRORCODE,
            WithOption.WithEncoding("ASCII"));

    String errorCode =
        readSimpleField("errorCode", readString(readBuffer, 32), WithOption.WithEncoding("ASCII"));

    int blockIdControllerReadyStatus =
        readConstField(
            "blockIdControllerReadyStatus",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageAlarmRev1.BLOCKIDCONTROLLERREADYSTATUS,
            WithOption.WithEncoding("ASCII"));

    NokOk controllerReadyStatus =
        readEnumField(
            "controllerReadyStatus",
            "NokOk",
            new DataReaderEnumDefault<>(NokOk::enumForValue, readUnsignedShort(readBuffer, 8)),
            WithOption.WithEncoding("ASCII"));

    int blockIdToolReadyStatus =
        readConstField(
            "blockIdToolReadyStatus",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageAlarmRev1.BLOCKIDTOOLREADYSTATUS,
            WithOption.WithEncoding("ASCII"));

    NokOk toolReadyStatus =
        readEnumField(
            "toolReadyStatus",
            "NokOk",
            new DataReaderEnumDefault<>(NokOk::enumForValue, readUnsignedShort(readBuffer, 8)),
            WithOption.WithEncoding("ASCII"));

    int blockIdTime =
        readConstField(
            "blockIdTime",
            readUnsignedInt(readBuffer, 16),
            OpenProtocolMessageAlarmRev1.BLOCKIDTIME,
            WithOption.WithEncoding("ASCII"));

    String alarmTime =
        readSimpleField("alarmTime", readString(readBuffer, 152), WithOption.WithEncoding("ASCII"));

    readBuffer.closeContext("OpenProtocolMessageAlarmRev1");
    // Create the instance
    return new OpenProtocolMessageAlarmRev1BuilderImpl(
        errorCode, controllerReadyStatus, toolReadyStatus, alarmTime);
  }

  public static class OpenProtocolMessageAlarmRev1BuilderImpl
      implements OpenProtocolMessageAlarm.OpenProtocolMessageAlarmBuilder {
    private final String errorCode;
    private final NokOk controllerReadyStatus;
    private final NokOk toolReadyStatus;
    private final String alarmTime;

    public OpenProtocolMessageAlarmRev1BuilderImpl(
        String errorCode, NokOk controllerReadyStatus, NokOk toolReadyStatus, String alarmTime) {
      this.errorCode = errorCode;
      this.controllerReadyStatus = controllerReadyStatus;
      this.toolReadyStatus = toolReadyStatus;
      this.alarmTime = alarmTime;
    }

    public OpenProtocolMessageAlarmRev1 build(
        Integer midRevision,
        Short noAckFlag,
        Integer targetStationId,
        Integer targetSpindleId,
        Integer sequenceNumber,
        Short numberOfMessageParts,
        Short messagePartNumber) {
      OpenProtocolMessageAlarmRev1 openProtocolMessageAlarmRev1 =
          new OpenProtocolMessageAlarmRev1(
              midRevision,
              noAckFlag,
              targetStationId,
              targetSpindleId,
              sequenceNumber,
              numberOfMessageParts,
              messagePartNumber,
              errorCode,
              controllerReadyStatus,
              toolReadyStatus,
              alarmTime);
      return openProtocolMessageAlarmRev1;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenProtocolMessageAlarmRev1)) {
      return false;
    }
    OpenProtocolMessageAlarmRev1 that = (OpenProtocolMessageAlarmRev1) o;
    return (getErrorCode() == that.getErrorCode())
        && (getControllerReadyStatus() == that.getControllerReadyStatus())
        && (getToolReadyStatus() == that.getToolReadyStatus())
        && (getAlarmTime() == that.getAlarmTime())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getErrorCode(),
        getControllerReadyStatus(),
        getToolReadyStatus(),
        getAlarmTime());
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
