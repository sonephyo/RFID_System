void setUpWifi() {
  WiFi.begin(SSID, PASS);
  while (WiFi.status() != WL_CONNECTED)
  {
    delay(500);
    Serial.print(".");
  }
  Serial.println("\n WiFi connected!");
  Serial.print("IP: ");
  Serial.println(WiFi.localIP());
}

void reconnectWifi()
{
  if (WiFi.status() != WL_CONNECTED)
  {
    Serial.println("Reconnecting WiFi...");
    WiFi.reconnect();
    delay(2000);
    return;
  }
}

bool connectBackend(WiFiClient &client)
{
  if (!client.connect(HOST, HTTPPORT))
  {
    Serial.println("Connection failed");
    delay(10000);
    return false;
  }
  return true;
}