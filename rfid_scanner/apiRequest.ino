void getUser(WiFiClient &client) {

  Serial.println("Get Request User");
  String url = "/users/";

  client.print(String("GET ") + url + " HTTP/1.1\r\n" +
               "Host: " + HOST + "\r\n" +
               "User-Agent: ESP32HTTPClient/1.0\r\n" +
               "Accept: application/json\r\n" +
               "Connection: close\r\n\r\n");

  waitForResponse(client);

  JsonDocument doc = parseJson(client);

  if (doc.isNull()) {
    return;
  }
  
  JsonArray dataArr = doc["data"];
  for (JsonObject person : dataArr) {
    int id = person["ID"];
    String name = person["Name"];
    int age = person["Age"];
    String cardID = person["CardID"];
    
    Serial.println("-------------------");
    Serial.print("ID: ");
    Serial.println(id);
    Serial.print("Name: ");
    Serial.println(name);
    Serial.print("Age: ");
    Serial.println(age);
    Serial.print("Card ID: ");
    Serial.println(cardID);
  }
}

JsonDocument parseJson(WiFiClient &client) {
  String jsonResponse = "";
  bool headersEnded = false;

  while (client.connected() || client.available()) {
    if (client.available()) {
      String line = client.readStringUntil('\n');
      if (!headersEnded) {
        if (line == "\r" || line.length() <= 1) {
          headersEnded = true;
        }
      } else {
        jsonResponse += line;
      }
    }
  }

  JsonDocument doc;
  DeserializationError error = deserializeJson(doc, jsonResponse);

  if (error) {
    Serial.print("JSON parsing failed: ");
    Serial.println(error.c_str());
  }

  return doc;
}

void waitForResponse(WiFiClient &client) {
  unsigned long timeout = millis();
  while (client.available() == 0) {
    if (millis() - timeout > 5000) {
      Serial.println(">>> Client Timeout!");
      client.stop();
      return;
    }
  }
}