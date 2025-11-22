void getUser(WiFiClientSecure &client) {

  Serial.println("Get Request User");
  String url = "/users/sonephyo";

  // Properly formatted HTTP GET request
  client.print(String("GET ") + url + " HTTP/1.1\r\n" +
               "Host: " + HOST + "\r\n" +
               "User-Agent: ESP32HTTPClient/1.0\r\n" +
               "Accept: application/json\r\n" +
               "Connection: close\r\n\r\n");

  unsigned long timeout = millis();
  while (client.available() == 0) {
    if (millis() - timeout > 5000) {
      Serial.println(">>> Client Timeout!");
      client.stop();
      return;
    }
  }

  // Read server response
  while (client.connected() || client.available()) {
    if (client.available()) {
      String line = client.readStringUntil('\n');
      Serial.println(line);
    }
  }
}