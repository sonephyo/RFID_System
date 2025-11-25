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