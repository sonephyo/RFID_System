#define VRX 34
#define VRY 35
#define SW 32

void setupJoystick() {
  pinMode(SW, INPUT_PULLUP);
}

int getJoystickX() {
  return analogRead(VRX);
}

int getJoystickY() {
  return analogRead(VRY);
}

bool isButtonPressed() {
  return digitalRead(SW) == 0;
}

String getDirection() {
  int x = getJoystickX();
  int y = getJoystickY();

  if (x < 1000) return "LEFT";
  if (x > 3000) return "RIGHT";
  if (y < 1000) return "UP";
  if (y > 3000) return "DOWN";
  return "CENTER";
}
void debugJoyStick() {
  Serial.print("X: ");
  Serial.print(getJoystickX());
  Serial.print(" | Y: ");
  Serial.print(getJoystickY());
  Serial.print(" | Button: ");
  Serial.print(isButtonPressed() ? "PRESSED" : "---");
  Serial.print(" | ");
  Serial.println(getDirection());
}

int selectMode() {
  int selected = 0;
  displayMenu("Attendance", "Config", selected);
  delay(300);

  while (true) {
    String dir = getDirection();

    if (dir == "UP" && selected != 0) {
      selected = 0;
      displayMenu("Attendance", "Config", selected);
      delay(200);
    }

    if (dir == "DOWN" && selected != 1) {
      selected = 1;
      displayMenu("Attendance", "Config", selected);
      delay(200);
    }

    if (isButtonPressed()) {
      delay(200);
      return selected;
    }
  }
}
