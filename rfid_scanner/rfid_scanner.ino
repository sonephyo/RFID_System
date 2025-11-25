#define RXp2 16
#define TXp2 17
#include <WiFi.h>
#include <ArduinoJson.h>
#include "secrets.h"

String lastCard = "";

#define MODE_ATTENDANCE 0
#define MODE_CONFIG 1

int currentMode = MODE_ATTENDANCE;

void setup() {
  Serial.begin(115200);
  Serial2.begin(9600, SERIAL_8N1, RXp2, TXp2);
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
  if (Serial2.available()) {
    Serial.println(Serial2.readStringUntil('\n'));
  }

  if (currentMode == MODE_ATTENDANCE) {
    attendanceOperation();
  } else {
    configOperation();
  }
}

void configOperation() {
  Serial.print("X: ");
  Serial.print(getJoystickX());
  Serial.print(" | Y: ");
  Serial.print(getJoystickY());
  Serial.print(" | Button: ");
  Serial.print(isButtonPressed() ? "PRESSED" : "---");
  Serial.print(" | ");
  Serial.println(getDirection());

  delay(200);
}

void attendanceOperation() {
  reconnectWifi();
  WiFiClient client;
  bool isConnected = connectBackend(client);
  if (isConnected) {
    getUser(client);
  } else {
    Serial.println("Backend cannot be connected.");
  }
}
