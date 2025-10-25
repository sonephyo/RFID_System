#include <WiFi.h>
#include "secrets.h"

const char *host = "api.thingspeak.com";
const int httpPort = 80;
const String writeApiKey = "JG4S1IKAS91AE0CN";

int field1 = 0;

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

  WiFiClient client;
  if (!client.connect(host, httpPort)) {
    Serial.println("❌ Connection failed");
    delay(10000);
    return;
  }

  String url = "/update?api_key=" + writeApiKey + "&field1=" + String(field1);

  // Properly formatted HTTP GET request
  client.print(String("GET ") + url + " HTTP/1.1\r\n" +
               "Host: " + host + "\r\n" +
               "Connection: close\r\n\r\n");

  Serial.println("📤 Sending to ThingSpeak: field1 = " + String(field1));

  // Read server response
  while (client.connected() || client.available()) {
    if (client.available()) {
      String line = client.readStringUntil('\n');
      Serial.println(line);
    }
  }
  
  delay(20000);
}
