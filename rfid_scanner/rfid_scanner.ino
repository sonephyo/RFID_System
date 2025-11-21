// #include <WiFi.h>
// #include "secrets.h"

// const char *host = "api.thingspeak.com";
// const int httpPort = 80;
// const String writeApiKey = "JG4S1IKAS91AE0CN";

// int field1 = 0;

// void setup() {
//   Serial.begin(115200);
//   Serial.println("Connecting to WiFi...");
//   WiFi.begin(SSID, PASS);
//   while (WiFi.status() != WL_CONNECTED) {
//     delay(500);
//     Serial.print(".");
//   }
//   Serial.println("\nWiFi connected!");
//   Serial.print("IP: ");
//   Serial.println(WiFi.localIP());
// }

// void loop() {
//   if (WiFi.status() != WL_CONNECTED) {
//     Serial.println("Reconnecting WiFi...");
//     WiFi.reconnect();
//     delay(2000);
//     return;
//   }

//   WiFiClient client;
//   if (!client.connect(host, httpPort)) {
//     Serial.println("Connection failed");
//     delay(10000);
//     return;
//   }

//   String url = "/update?api_key=" + writeApiKey + "&field1=" + String(field1);

//   // Properly formatted HTTP GET request
//   client.print(String("GET ") + url + " HTTP/1.1\r\n" +
//                "Host: " + host + "\r\n" +
//                "Connection: close\r\n\r\n");

//   Serial.println("📤 Sending to ThingSpeak: field1 = " + String(field1));

//   // Read server response
//   while (client.connected() || client.available()) {
//     if (client.available()) {
//       String line = client.readStringUntil('\n');
//       Serial.println(line);
//     }
//   }

//   delay(20000);
// }


/*
  Rui Santos & Sara Santos - Random Nerd Tutorials
  Complete project details at https://RandomNerdTutorials.com/esp32-mfrc522-rfid-reader-arduino/
  Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files.
  The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
*/

#include <MFRC522v2.h>
#include <MFRC522DriverSPI.h>
//#include <MFRC522DriverI2C.h>
#include <MFRC522DriverPinSimple.h>
#include <MFRC522Debug.h>

// Learn more about using SPI/I2C or check the pin assigment for your board: https://github.com/OSSLibraries/Arduino_MFRC522v2#pin-layout
MFRC522DriverPinSimple ss_pin(5);

MFRC522DriverSPI driver{ss_pin}; // Create SPI driver
//MFRC522DriverI2C driver{};     // Create I2C driver
MFRC522 mfrc522{driver};         // Create MFRC522 instance

void setup() {
  Serial.begin(115200);  // Initialize serial communication
  while (!Serial);       // Do nothing if no serial port is opened (added for Arduinos based on ATMEGA32U4).

  mfrc522.PCD_Init();    // Init MFRC522 board.
  MFRC522Debug::PCD_DumpVersionToSerial(mfrc522, Serial);	// Show details of PCD - MFRC522 Card Reader details.
  Serial.println(F("Scan PICC to see UID, SAK, type, and data blocks..."));
}

void loop() {
  // Reset the loop if no new card present on the sensor/reader. This saves the entire process when idle.
  if (!mfrc522.PICC_IsNewCardPresent()) {
    return;
  }

  // Select one of the cards.
  if (!mfrc522.PICC_ReadCardSerial()) {
    return;
  }

  // Dump debug info about the card; PICC_HaltA() is automatically called.
  MFRC522Debug::PICC_DumpToSerial(mfrc522, Serial, &(mfrc522.uid));

  delay(2000);
}
