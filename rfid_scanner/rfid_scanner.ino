#define RXp2 16
#define TXp2 17

String lastCard = "";

void setup() {
  Serial.begin(115200);
  Serial2.begin(9600, SERIAL_8N1, RXp2, TXp2);
}

void loop() {
  if (Serial2.available()) {
    Serial.println(Serial2.readStringUntil('\n'));
    // String cardID = Serial2.readStringUntil('\n');
    // cardID.trim();
    
    // if (cardID.length() > 0 && cardID != lastCard) {
    //   Serial.println(cardID);
    //   lastCard = cardID;
    //   delay(2000);
    // }
  }
} 