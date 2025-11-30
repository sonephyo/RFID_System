#include <WiFi.h>
#include <HTTPClient.h>

void setUpWifi() {
    WiFi.begin(SSID, PASS);
    while (WiFi.status() != WL_CONNECTED) {
        delay(500);
        Serial.print(".");
    }
    Serial.println("\nWiFi connected!");
    Serial.print("IP: ");
    Serial.println(WiFi.localIP());
}

void reconnectWifi() {
    if (WiFi.status() != WL_CONNECTED) {
        Serial.println("Reconnecting WiFi...");
        WiFi.reconnect();
        delay(2000);
    }
}

bool checkWifi() {
    return WiFi.status() == WL_CONNECTED;
}