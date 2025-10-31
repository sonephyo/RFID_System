#include <WiFi.h>
#include <WiFiClientSecure.h>
#include "secrets.h"

void setup() {
  Serial.begin(115200);
  Serial.println("Connecting to WiFi...");
  WiFi.begin(SSID, PASS);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("\n WiFi connected!");
  Serial.print("IP: ");
  Serial.println(WiFi.localIP());
}

void loop() {
  reconnectWifi();
  WiFiClientSecure client;
  client.setInsecure();
  bool isConnected = connectBackend(client);
  if (isConnected) {
    getUser(client);
  } else {
    Serial.println("Backend cannot be connected.");
  }
  
  delay(20000);
}

void reconnectWifi() { 
  if (WiFi.status() != WL_CONNECTED) {
    Serial.println("Reconnecting WiFi...");
    WiFi.reconnect();
    delay(2000);
    return;
  }
}

bool connectBackend(WiFiClientSecure &client) {
  if (!client.connect(HOST, HTTPPORT)) {
    Serial.println("Connection failed");
    delay(10000);
    return false;
  }
  return true;
}
