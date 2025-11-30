#include <WiFi.h>
#include <ArduinoJson.h>
#include "secrets.h"

String lastCard = "";

#define MODE_ATTENDANCE 0
#define MODE_CONFIG 1

int currentMode = MODE_ATTENDANCE;

void setup() {
  Serial.begin(115200);
  setupCardScanner();
  setupLCD();
  setupJoystick();

  displayBothLines("RFID System", "Select mode...");
  delay(2000);

  currentMode = selectMode();
  clearAllLines();

  if (currentMode == MODE_ATTENDANCE) {
    scrollText(0, "Connecting to Wifi...");
    setUpWifi();
    displayLine1("Connected");
    delay(2000);
    displayBothLines("ATTENDANCE", "Scan card...");
  } else {
    displayLine1("CONFIG");
    scrollText(0, "Waiting to connect PC");
  }
}

void loop() {
  // if (Serial2.available()) {
  //   Serial.println(Serial2.readStringUntil('\n'));
  // }

  if (currentMode == MODE_ATTENDANCE) {
    attendanceOperation();
  } else {
    configOperation();
  }
}

void configOperation() {

  waitForConnection();
  
  String card = readCard();

  if (card != "") {
    Serial.println("Card: " + card);  // Sending to frontend
    displayLine1("Card scanned:");
    displayLine2(card);
  }
}

void attendanceOperation() {
  // reconnectWifi();
  // WiFiClient client;
  // bool isConnected = connectBackend(client);
  // if (isConnected) {
  //   getUser(client);
  // } else {
  //   Serial.println("Backend cannot be connected.");
  // }
}
