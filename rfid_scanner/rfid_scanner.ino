#define RXp2 16
#define TXp2 17
#include <WiFi.h>
#include <ArduinoJson.h>
#include "secrets.h"

String lastCard = "";

void setup()
{
  Serial.begin(115200);
  Serial2.begin(9600, SERIAL_8N1, RXp2, TXp2);
  Serial.println("Connecting to WiFi...");
  setUpWifi();
}

void loop()
{
  if (Serial2.available())
  {
    Serial.println(Serial2.readStringUntil('\n'));
  }

  reconnectWifi();
  WiFiClient client;
  bool isConnected = connectBackend(client);
  if (isConnected)
  {
    getUser(client);
  }
  else
  {
    Serial.println("Backend cannot be connected.");
  }

  delay(20000);
}

