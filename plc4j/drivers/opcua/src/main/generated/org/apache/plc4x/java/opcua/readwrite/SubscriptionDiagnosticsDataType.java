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
package org.apache.plc4x.java.opcua.readwrite;

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

public class SubscriptionDiagnosticsDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "876";
  }

  // Properties.
  protected final NodeId sessionId;
  protected final long subscriptionId;
  protected final short priority;
  protected final double publishingInterval;
  protected final long maxKeepAliveCount;
  protected final long maxLifetimeCount;
  protected final long maxNotificationsPerPublish;
  protected final boolean publishingEnabled;
  protected final long modifyCount;
  protected final long enableCount;
  protected final long disableCount;
  protected final long republishRequestCount;
  protected final long republishMessageRequestCount;
  protected final long republishMessageCount;
  protected final long transferRequestCount;
  protected final long transferredToAltClientCount;
  protected final long transferredToSameClientCount;
  protected final long publishRequestCount;
  protected final long dataChangeNotificationsCount;
  protected final long eventNotificationsCount;
  protected final long notificationsCount;
  protected final long latePublishRequestCount;
  protected final long currentKeepAliveCount;
  protected final long currentLifetimeCount;
  protected final long unacknowledgedMessageCount;
  protected final long discardedMessageCount;
  protected final long monitoredItemCount;
  protected final long disabledMonitoredItemCount;
  protected final long monitoringQueueOverflowCount;
  protected final long nextSequenceNumber;
  protected final long eventQueueOverFlowCount;

  public SubscriptionDiagnosticsDataType(
      NodeId sessionId,
      long subscriptionId,
      short priority,
      double publishingInterval,
      long maxKeepAliveCount,
      long maxLifetimeCount,
      long maxNotificationsPerPublish,
      boolean publishingEnabled,
      long modifyCount,
      long enableCount,
      long disableCount,
      long republishRequestCount,
      long republishMessageRequestCount,
      long republishMessageCount,
      long transferRequestCount,
      long transferredToAltClientCount,
      long transferredToSameClientCount,
      long publishRequestCount,
      long dataChangeNotificationsCount,
      long eventNotificationsCount,
      long notificationsCount,
      long latePublishRequestCount,
      long currentKeepAliveCount,
      long currentLifetimeCount,
      long unacknowledgedMessageCount,
      long discardedMessageCount,
      long monitoredItemCount,
      long disabledMonitoredItemCount,
      long monitoringQueueOverflowCount,
      long nextSequenceNumber,
      long eventQueueOverFlowCount) {
    super();
    this.sessionId = sessionId;
    this.subscriptionId = subscriptionId;
    this.priority = priority;
    this.publishingInterval = publishingInterval;
    this.maxKeepAliveCount = maxKeepAliveCount;
    this.maxLifetimeCount = maxLifetimeCount;
    this.maxNotificationsPerPublish = maxNotificationsPerPublish;
    this.publishingEnabled = publishingEnabled;
    this.modifyCount = modifyCount;
    this.enableCount = enableCount;
    this.disableCount = disableCount;
    this.republishRequestCount = republishRequestCount;
    this.republishMessageRequestCount = republishMessageRequestCount;
    this.republishMessageCount = republishMessageCount;
    this.transferRequestCount = transferRequestCount;
    this.transferredToAltClientCount = transferredToAltClientCount;
    this.transferredToSameClientCount = transferredToSameClientCount;
    this.publishRequestCount = publishRequestCount;
    this.dataChangeNotificationsCount = dataChangeNotificationsCount;
    this.eventNotificationsCount = eventNotificationsCount;
    this.notificationsCount = notificationsCount;
    this.latePublishRequestCount = latePublishRequestCount;
    this.currentKeepAliveCount = currentKeepAliveCount;
    this.currentLifetimeCount = currentLifetimeCount;
    this.unacknowledgedMessageCount = unacknowledgedMessageCount;
    this.discardedMessageCount = discardedMessageCount;
    this.monitoredItemCount = monitoredItemCount;
    this.disabledMonitoredItemCount = disabledMonitoredItemCount;
    this.monitoringQueueOverflowCount = monitoringQueueOverflowCount;
    this.nextSequenceNumber = nextSequenceNumber;
    this.eventQueueOverFlowCount = eventQueueOverFlowCount;
  }

  public NodeId getSessionId() {
    return sessionId;
  }

  public long getSubscriptionId() {
    return subscriptionId;
  }

  public short getPriority() {
    return priority;
  }

  public double getPublishingInterval() {
    return publishingInterval;
  }

  public long getMaxKeepAliveCount() {
    return maxKeepAliveCount;
  }

  public long getMaxLifetimeCount() {
    return maxLifetimeCount;
  }

  public long getMaxNotificationsPerPublish() {
    return maxNotificationsPerPublish;
  }

  public boolean getPublishingEnabled() {
    return publishingEnabled;
  }

  public long getModifyCount() {
    return modifyCount;
  }

  public long getEnableCount() {
    return enableCount;
  }

  public long getDisableCount() {
    return disableCount;
  }

  public long getRepublishRequestCount() {
    return republishRequestCount;
  }

  public long getRepublishMessageRequestCount() {
    return republishMessageRequestCount;
  }

  public long getRepublishMessageCount() {
    return republishMessageCount;
  }

  public long getTransferRequestCount() {
    return transferRequestCount;
  }

  public long getTransferredToAltClientCount() {
    return transferredToAltClientCount;
  }

  public long getTransferredToSameClientCount() {
    return transferredToSameClientCount;
  }

  public long getPublishRequestCount() {
    return publishRequestCount;
  }

  public long getDataChangeNotificationsCount() {
    return dataChangeNotificationsCount;
  }

  public long getEventNotificationsCount() {
    return eventNotificationsCount;
  }

  public long getNotificationsCount() {
    return notificationsCount;
  }

  public long getLatePublishRequestCount() {
    return latePublishRequestCount;
  }

  public long getCurrentKeepAliveCount() {
    return currentKeepAliveCount;
  }

  public long getCurrentLifetimeCount() {
    return currentLifetimeCount;
  }

  public long getUnacknowledgedMessageCount() {
    return unacknowledgedMessageCount;
  }

  public long getDiscardedMessageCount() {
    return discardedMessageCount;
  }

  public long getMonitoredItemCount() {
    return monitoredItemCount;
  }

  public long getDisabledMonitoredItemCount() {
    return disabledMonitoredItemCount;
  }

  public long getMonitoringQueueOverflowCount() {
    return monitoringQueueOverflowCount;
  }

  public long getNextSequenceNumber() {
    return nextSequenceNumber;
  }

  public long getEventQueueOverFlowCount() {
    return eventQueueOverFlowCount;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SubscriptionDiagnosticsDataType");

    // Simple Field (sessionId)
    writeSimpleField("sessionId", sessionId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (subscriptionId)
    writeSimpleField("subscriptionId", subscriptionId, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (priority)
    writeSimpleField("priority", priority, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (publishingInterval)
    writeSimpleField("publishingInterval", publishingInterval, writeDouble(writeBuffer, 64));

    // Simple Field (maxKeepAliveCount)
    writeSimpleField("maxKeepAliveCount", maxKeepAliveCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (maxLifetimeCount)
    writeSimpleField("maxLifetimeCount", maxLifetimeCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (maxNotificationsPerPublish)
    writeSimpleField(
        "maxNotificationsPerPublish",
        maxNotificationsPerPublish,
        writeUnsignedLong(writeBuffer, 32));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 7));

    // Simple Field (publishingEnabled)
    writeSimpleField("publishingEnabled", publishingEnabled, writeBoolean(writeBuffer));

    // Simple Field (modifyCount)
    writeSimpleField("modifyCount", modifyCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (enableCount)
    writeSimpleField("enableCount", enableCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (disableCount)
    writeSimpleField("disableCount", disableCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (republishRequestCount)
    writeSimpleField(
        "republishRequestCount", republishRequestCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (republishMessageRequestCount)
    writeSimpleField(
        "republishMessageRequestCount",
        republishMessageRequestCount,
        writeUnsignedLong(writeBuffer, 32));

    // Simple Field (republishMessageCount)
    writeSimpleField(
        "republishMessageCount", republishMessageCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (transferRequestCount)
    writeSimpleField(
        "transferRequestCount", transferRequestCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (transferredToAltClientCount)
    writeSimpleField(
        "transferredToAltClientCount",
        transferredToAltClientCount,
        writeUnsignedLong(writeBuffer, 32));

    // Simple Field (transferredToSameClientCount)
    writeSimpleField(
        "transferredToSameClientCount",
        transferredToSameClientCount,
        writeUnsignedLong(writeBuffer, 32));

    // Simple Field (publishRequestCount)
    writeSimpleField(
        "publishRequestCount", publishRequestCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (dataChangeNotificationsCount)
    writeSimpleField(
        "dataChangeNotificationsCount",
        dataChangeNotificationsCount,
        writeUnsignedLong(writeBuffer, 32));

    // Simple Field (eventNotificationsCount)
    writeSimpleField(
        "eventNotificationsCount", eventNotificationsCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (notificationsCount)
    writeSimpleField("notificationsCount", notificationsCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (latePublishRequestCount)
    writeSimpleField(
        "latePublishRequestCount", latePublishRequestCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (currentKeepAliveCount)
    writeSimpleField(
        "currentKeepAliveCount", currentKeepAliveCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (currentLifetimeCount)
    writeSimpleField(
        "currentLifetimeCount", currentLifetimeCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (unacknowledgedMessageCount)
    writeSimpleField(
        "unacknowledgedMessageCount",
        unacknowledgedMessageCount,
        writeUnsignedLong(writeBuffer, 32));

    // Simple Field (discardedMessageCount)
    writeSimpleField(
        "discardedMessageCount", discardedMessageCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (monitoredItemCount)
    writeSimpleField("monitoredItemCount", monitoredItemCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (disabledMonitoredItemCount)
    writeSimpleField(
        "disabledMonitoredItemCount",
        disabledMonitoredItemCount,
        writeUnsignedLong(writeBuffer, 32));

    // Simple Field (monitoringQueueOverflowCount)
    writeSimpleField(
        "monitoringQueueOverflowCount",
        monitoringQueueOverflowCount,
        writeUnsignedLong(writeBuffer, 32));

    // Simple Field (nextSequenceNumber)
    writeSimpleField("nextSequenceNumber", nextSequenceNumber, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (eventQueueOverFlowCount)
    writeSimpleField(
        "eventQueueOverFlowCount", eventQueueOverFlowCount, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("SubscriptionDiagnosticsDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SubscriptionDiagnosticsDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (sessionId)
    lengthInBits += sessionId.getLengthInBits();

    // Simple field (subscriptionId)
    lengthInBits += 32;

    // Simple field (priority)
    lengthInBits += 8;

    // Simple field (publishingInterval)
    lengthInBits += 64;

    // Simple field (maxKeepAliveCount)
    lengthInBits += 32;

    // Simple field (maxLifetimeCount)
    lengthInBits += 32;

    // Simple field (maxNotificationsPerPublish)
    lengthInBits += 32;

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (publishingEnabled)
    lengthInBits += 1;

    // Simple field (modifyCount)
    lengthInBits += 32;

    // Simple field (enableCount)
    lengthInBits += 32;

    // Simple field (disableCount)
    lengthInBits += 32;

    // Simple field (republishRequestCount)
    lengthInBits += 32;

    // Simple field (republishMessageRequestCount)
    lengthInBits += 32;

    // Simple field (republishMessageCount)
    lengthInBits += 32;

    // Simple field (transferRequestCount)
    lengthInBits += 32;

    // Simple field (transferredToAltClientCount)
    lengthInBits += 32;

    // Simple field (transferredToSameClientCount)
    lengthInBits += 32;

    // Simple field (publishRequestCount)
    lengthInBits += 32;

    // Simple field (dataChangeNotificationsCount)
    lengthInBits += 32;

    // Simple field (eventNotificationsCount)
    lengthInBits += 32;

    // Simple field (notificationsCount)
    lengthInBits += 32;

    // Simple field (latePublishRequestCount)
    lengthInBits += 32;

    // Simple field (currentKeepAliveCount)
    lengthInBits += 32;

    // Simple field (currentLifetimeCount)
    lengthInBits += 32;

    // Simple field (unacknowledgedMessageCount)
    lengthInBits += 32;

    // Simple field (discardedMessageCount)
    lengthInBits += 32;

    // Simple field (monitoredItemCount)
    lengthInBits += 32;

    // Simple field (disabledMonitoredItemCount)
    lengthInBits += 32;

    // Simple field (monitoringQueueOverflowCount)
    lengthInBits += 32;

    // Simple field (nextSequenceNumber)
    lengthInBits += 32;

    // Simple field (eventQueueOverFlowCount)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("SubscriptionDiagnosticsDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId sessionId =
        readSimpleField(
            "sessionId",
            new DataReaderComplexDefault<>(() -> NodeId.staticParse(readBuffer), readBuffer));

    long subscriptionId = readSimpleField("subscriptionId", readUnsignedLong(readBuffer, 32));

    short priority = readSimpleField("priority", readUnsignedShort(readBuffer, 8));

    double publishingInterval = readSimpleField("publishingInterval", readDouble(readBuffer, 64));

    long maxKeepAliveCount = readSimpleField("maxKeepAliveCount", readUnsignedLong(readBuffer, 32));

    long maxLifetimeCount = readSimpleField("maxLifetimeCount", readUnsignedLong(readBuffer, 32));

    long maxNotificationsPerPublish =
        readSimpleField("maxNotificationsPerPublish", readUnsignedLong(readBuffer, 32));

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 7), (byte) 0x00);

    boolean publishingEnabled = readSimpleField("publishingEnabled", readBoolean(readBuffer));

    long modifyCount = readSimpleField("modifyCount", readUnsignedLong(readBuffer, 32));

    long enableCount = readSimpleField("enableCount", readUnsignedLong(readBuffer, 32));

    long disableCount = readSimpleField("disableCount", readUnsignedLong(readBuffer, 32));

    long republishRequestCount =
        readSimpleField("republishRequestCount", readUnsignedLong(readBuffer, 32));

    long republishMessageRequestCount =
        readSimpleField("republishMessageRequestCount", readUnsignedLong(readBuffer, 32));

    long republishMessageCount =
        readSimpleField("republishMessageCount", readUnsignedLong(readBuffer, 32));

    long transferRequestCount =
        readSimpleField("transferRequestCount", readUnsignedLong(readBuffer, 32));

    long transferredToAltClientCount =
        readSimpleField("transferredToAltClientCount", readUnsignedLong(readBuffer, 32));

    long transferredToSameClientCount =
        readSimpleField("transferredToSameClientCount", readUnsignedLong(readBuffer, 32));

    long publishRequestCount =
        readSimpleField("publishRequestCount", readUnsignedLong(readBuffer, 32));

    long dataChangeNotificationsCount =
        readSimpleField("dataChangeNotificationsCount", readUnsignedLong(readBuffer, 32));

    long eventNotificationsCount =
        readSimpleField("eventNotificationsCount", readUnsignedLong(readBuffer, 32));

    long notificationsCount =
        readSimpleField("notificationsCount", readUnsignedLong(readBuffer, 32));

    long latePublishRequestCount =
        readSimpleField("latePublishRequestCount", readUnsignedLong(readBuffer, 32));

    long currentKeepAliveCount =
        readSimpleField("currentKeepAliveCount", readUnsignedLong(readBuffer, 32));

    long currentLifetimeCount =
        readSimpleField("currentLifetimeCount", readUnsignedLong(readBuffer, 32));

    long unacknowledgedMessageCount =
        readSimpleField("unacknowledgedMessageCount", readUnsignedLong(readBuffer, 32));

    long discardedMessageCount =
        readSimpleField("discardedMessageCount", readUnsignedLong(readBuffer, 32));

    long monitoredItemCount =
        readSimpleField("monitoredItemCount", readUnsignedLong(readBuffer, 32));

    long disabledMonitoredItemCount =
        readSimpleField("disabledMonitoredItemCount", readUnsignedLong(readBuffer, 32));

    long monitoringQueueOverflowCount =
        readSimpleField("monitoringQueueOverflowCount", readUnsignedLong(readBuffer, 32));

    long nextSequenceNumber =
        readSimpleField("nextSequenceNumber", readUnsignedLong(readBuffer, 32));

    long eventQueueOverFlowCount =
        readSimpleField("eventQueueOverFlowCount", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("SubscriptionDiagnosticsDataType");
    // Create the instance
    return new SubscriptionDiagnosticsDataTypeBuilderImpl(
        sessionId,
        subscriptionId,
        priority,
        publishingInterval,
        maxKeepAliveCount,
        maxLifetimeCount,
        maxNotificationsPerPublish,
        publishingEnabled,
        modifyCount,
        enableCount,
        disableCount,
        republishRequestCount,
        republishMessageRequestCount,
        republishMessageCount,
        transferRequestCount,
        transferredToAltClientCount,
        transferredToSameClientCount,
        publishRequestCount,
        dataChangeNotificationsCount,
        eventNotificationsCount,
        notificationsCount,
        latePublishRequestCount,
        currentKeepAliveCount,
        currentLifetimeCount,
        unacknowledgedMessageCount,
        discardedMessageCount,
        monitoredItemCount,
        disabledMonitoredItemCount,
        monitoringQueueOverflowCount,
        nextSequenceNumber,
        eventQueueOverFlowCount);
  }

  public static class SubscriptionDiagnosticsDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId sessionId;
    private final long subscriptionId;
    private final short priority;
    private final double publishingInterval;
    private final long maxKeepAliveCount;
    private final long maxLifetimeCount;
    private final long maxNotificationsPerPublish;
    private final boolean publishingEnabled;
    private final long modifyCount;
    private final long enableCount;
    private final long disableCount;
    private final long republishRequestCount;
    private final long republishMessageRequestCount;
    private final long republishMessageCount;
    private final long transferRequestCount;
    private final long transferredToAltClientCount;
    private final long transferredToSameClientCount;
    private final long publishRequestCount;
    private final long dataChangeNotificationsCount;
    private final long eventNotificationsCount;
    private final long notificationsCount;
    private final long latePublishRequestCount;
    private final long currentKeepAliveCount;
    private final long currentLifetimeCount;
    private final long unacknowledgedMessageCount;
    private final long discardedMessageCount;
    private final long monitoredItemCount;
    private final long disabledMonitoredItemCount;
    private final long monitoringQueueOverflowCount;
    private final long nextSequenceNumber;
    private final long eventQueueOverFlowCount;

    public SubscriptionDiagnosticsDataTypeBuilderImpl(
        NodeId sessionId,
        long subscriptionId,
        short priority,
        double publishingInterval,
        long maxKeepAliveCount,
        long maxLifetimeCount,
        long maxNotificationsPerPublish,
        boolean publishingEnabled,
        long modifyCount,
        long enableCount,
        long disableCount,
        long republishRequestCount,
        long republishMessageRequestCount,
        long republishMessageCount,
        long transferRequestCount,
        long transferredToAltClientCount,
        long transferredToSameClientCount,
        long publishRequestCount,
        long dataChangeNotificationsCount,
        long eventNotificationsCount,
        long notificationsCount,
        long latePublishRequestCount,
        long currentKeepAliveCount,
        long currentLifetimeCount,
        long unacknowledgedMessageCount,
        long discardedMessageCount,
        long monitoredItemCount,
        long disabledMonitoredItemCount,
        long monitoringQueueOverflowCount,
        long nextSequenceNumber,
        long eventQueueOverFlowCount) {
      this.sessionId = sessionId;
      this.subscriptionId = subscriptionId;
      this.priority = priority;
      this.publishingInterval = publishingInterval;
      this.maxKeepAliveCount = maxKeepAliveCount;
      this.maxLifetimeCount = maxLifetimeCount;
      this.maxNotificationsPerPublish = maxNotificationsPerPublish;
      this.publishingEnabled = publishingEnabled;
      this.modifyCount = modifyCount;
      this.enableCount = enableCount;
      this.disableCount = disableCount;
      this.republishRequestCount = republishRequestCount;
      this.republishMessageRequestCount = republishMessageRequestCount;
      this.republishMessageCount = republishMessageCount;
      this.transferRequestCount = transferRequestCount;
      this.transferredToAltClientCount = transferredToAltClientCount;
      this.transferredToSameClientCount = transferredToSameClientCount;
      this.publishRequestCount = publishRequestCount;
      this.dataChangeNotificationsCount = dataChangeNotificationsCount;
      this.eventNotificationsCount = eventNotificationsCount;
      this.notificationsCount = notificationsCount;
      this.latePublishRequestCount = latePublishRequestCount;
      this.currentKeepAliveCount = currentKeepAliveCount;
      this.currentLifetimeCount = currentLifetimeCount;
      this.unacknowledgedMessageCount = unacknowledgedMessageCount;
      this.discardedMessageCount = discardedMessageCount;
      this.monitoredItemCount = monitoredItemCount;
      this.disabledMonitoredItemCount = disabledMonitoredItemCount;
      this.monitoringQueueOverflowCount = monitoringQueueOverflowCount;
      this.nextSequenceNumber = nextSequenceNumber;
      this.eventQueueOverFlowCount = eventQueueOverFlowCount;
    }

    public SubscriptionDiagnosticsDataType build() {
      SubscriptionDiagnosticsDataType subscriptionDiagnosticsDataType =
          new SubscriptionDiagnosticsDataType(
              sessionId,
              subscriptionId,
              priority,
              publishingInterval,
              maxKeepAliveCount,
              maxLifetimeCount,
              maxNotificationsPerPublish,
              publishingEnabled,
              modifyCount,
              enableCount,
              disableCount,
              republishRequestCount,
              republishMessageRequestCount,
              republishMessageCount,
              transferRequestCount,
              transferredToAltClientCount,
              transferredToSameClientCount,
              publishRequestCount,
              dataChangeNotificationsCount,
              eventNotificationsCount,
              notificationsCount,
              latePublishRequestCount,
              currentKeepAliveCount,
              currentLifetimeCount,
              unacknowledgedMessageCount,
              discardedMessageCount,
              monitoredItemCount,
              disabledMonitoredItemCount,
              monitoringQueueOverflowCount,
              nextSequenceNumber,
              eventQueueOverFlowCount);
      return subscriptionDiagnosticsDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SubscriptionDiagnosticsDataType)) {
      return false;
    }
    SubscriptionDiagnosticsDataType that = (SubscriptionDiagnosticsDataType) o;
    return (getSessionId() == that.getSessionId())
        && (getSubscriptionId() == that.getSubscriptionId())
        && (getPriority() == that.getPriority())
        && (getPublishingInterval() == that.getPublishingInterval())
        && (getMaxKeepAliveCount() == that.getMaxKeepAliveCount())
        && (getMaxLifetimeCount() == that.getMaxLifetimeCount())
        && (getMaxNotificationsPerPublish() == that.getMaxNotificationsPerPublish())
        && (getPublishingEnabled() == that.getPublishingEnabled())
        && (getModifyCount() == that.getModifyCount())
        && (getEnableCount() == that.getEnableCount())
        && (getDisableCount() == that.getDisableCount())
        && (getRepublishRequestCount() == that.getRepublishRequestCount())
        && (getRepublishMessageRequestCount() == that.getRepublishMessageRequestCount())
        && (getRepublishMessageCount() == that.getRepublishMessageCount())
        && (getTransferRequestCount() == that.getTransferRequestCount())
        && (getTransferredToAltClientCount() == that.getTransferredToAltClientCount())
        && (getTransferredToSameClientCount() == that.getTransferredToSameClientCount())
        && (getPublishRequestCount() == that.getPublishRequestCount())
        && (getDataChangeNotificationsCount() == that.getDataChangeNotificationsCount())
        && (getEventNotificationsCount() == that.getEventNotificationsCount())
        && (getNotificationsCount() == that.getNotificationsCount())
        && (getLatePublishRequestCount() == that.getLatePublishRequestCount())
        && (getCurrentKeepAliveCount() == that.getCurrentKeepAliveCount())
        && (getCurrentLifetimeCount() == that.getCurrentLifetimeCount())
        && (getUnacknowledgedMessageCount() == that.getUnacknowledgedMessageCount())
        && (getDiscardedMessageCount() == that.getDiscardedMessageCount())
        && (getMonitoredItemCount() == that.getMonitoredItemCount())
        && (getDisabledMonitoredItemCount() == that.getDisabledMonitoredItemCount())
        && (getMonitoringQueueOverflowCount() == that.getMonitoringQueueOverflowCount())
        && (getNextSequenceNumber() == that.getNextSequenceNumber())
        && (getEventQueueOverFlowCount() == that.getEventQueueOverFlowCount())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getSessionId(),
        getSubscriptionId(),
        getPriority(),
        getPublishingInterval(),
        getMaxKeepAliveCount(),
        getMaxLifetimeCount(),
        getMaxNotificationsPerPublish(),
        getPublishingEnabled(),
        getModifyCount(),
        getEnableCount(),
        getDisableCount(),
        getRepublishRequestCount(),
        getRepublishMessageRequestCount(),
        getRepublishMessageCount(),
        getTransferRequestCount(),
        getTransferredToAltClientCount(),
        getTransferredToSameClientCount(),
        getPublishRequestCount(),
        getDataChangeNotificationsCount(),
        getEventNotificationsCount(),
        getNotificationsCount(),
        getLatePublishRequestCount(),
        getCurrentKeepAliveCount(),
        getCurrentLifetimeCount(),
        getUnacknowledgedMessageCount(),
        getDiscardedMessageCount(),
        getMonitoredItemCount(),
        getDisabledMonitoredItemCount(),
        getMonitoringQueueOverflowCount(),
        getNextSequenceNumber(),
        getEventQueueOverFlowCount());
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
