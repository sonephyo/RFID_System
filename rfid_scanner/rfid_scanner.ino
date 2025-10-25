#include <WiFi.h>
#include <WiFiClientSecure.h>
#include "secrets.h"

const char *host = "api.github.com";
const int httpPort = 443;

void setup() {
  Serial.begin(115200);
  Serial.println("Connecting to WiFi...");
  WiFi.begin(SSID, PASS);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("\n✅ WiFi connected!");
  Serial.print("IP: ");
  Serial.println(WiFi.localIP());
}

void loop() {
  if (WiFi.status() != WL_CONNECTED) {
    Serial.println("Reconnecting WiFi...");
    WiFi.reconnect();
    delay(2000);
    return;
  }
  Serial.println("Wifi is still in connection.");

  WiFiClientSecure client;

  client.setInsecure();

  if (!client.connect(host, httpPort)) {
    Serial.println("❌ Connection failed");
    delay(10000);
    return;
  }

  String url = "/users/sonephyo";

  // Properly formatted HTTP GET request
  client.print(String("GET ") + url + " HTTP/1.1\r\n" +
               "Host: " + host + "\r\n" +
               "User-Agent: ESP32HTTPClient/1.0\r\n" +  // ADD THIS!
               "Accept: application/json\r\n" +
               "Connection: close\r\n\r\n");

  // Read server response
  while (client.connected() || client.available()) {
    if (client.available()) {
      String line = client.readStringUntil('\n');
      Serial.println(line);
    }
  }
  
  delay(200000);
}
