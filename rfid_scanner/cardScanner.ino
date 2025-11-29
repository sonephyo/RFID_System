String lastScannedCard = "";

void setupCardScanner() {
  Serial2.begin(9600, SERIAL_8N1, 16, 17);
}

String readCard() {
  if (Serial2.available()) {
    String data = Serial2.readStringUntil('\n');
    data.trim();
    
    if (data.startsWith("Card:")) {
      lastScannedCard = data.substring(5);
      lastScannedCard.trim();
      return lastScannedCard;
    }
  }
  return "";
}

String getLastScannedCard() {
  return lastScannedCard;
}

void clearLastCard() {
  lastScannedCard = "";
}