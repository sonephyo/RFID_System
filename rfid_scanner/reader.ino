String lastCommand = "";

String readCommand() {
  if (Serial.available()) {
    String data = Serial.readStringUntil('\n');
    data.trim();

    if (data.length() > 0) {
      lastCommand = data;
      return lastCommand;
    }
  }
  return "";
}

void waitForConnection() {
  static bool connected = false;

  if (!connected) {

    static bool waitingConnectionDisplay = false;
    if (!waitingConnectionDisplay) {
      displayLine1("Waiting for");
      displayLine2("USB Connection");
      waitingConnectionDisplay = true;
    }

    String cmd = readCommand();
    if (cmd.startsWith("MESSAGE: Connected!")) {
      connected = true;
      displayLine1("Connected!");
      scrollText(1, "Waiting for Card Scan");
      delay(1000);
    }
    return;
  }
}