#include <LiquidCrystal.h>

LiquidCrystal lcd(22, 23, 5, 18, 19, 21);

void setupLCD() {
  lcd.begin(16, 2);
}

void clearLCD() {
  lcd.clear();
}

void clearLine(int row) {
  lcd.setCursor(0, row);
  lcd.print("                ");
}

void displayAt(int col, int row, String text) {
  lcd.setCursor(col, row);
  lcd.print(text);
}

void displayLine1(String text) {
  clearLine(0);
  lcd.setCursor(0, 0);
  lcd.print(text.substring(0, 16));
}

void displayLine2(String text) {
  clearLine(1);
  lcd.setCursor(0, 1);
  lcd.print(text.substring(0, 16));
}

void displayBothLines(String line1, String line2) {
  displayLine1(line1);
  displayLine2(line2);
}

void displayMenu(String option1, String option2, int selected) {
  clearLCD();
  lcd.setCursor(0, 0);
  lcd.print(selected == 0 ? "> " : "  ");
  lcd.print(option1.substring(0, 14));
  
  lcd.setCursor(0, 1);
  lcd.print(selected == 1 ? "> " : "  ");
  lcd.print(option2.substring(0, 14));
}

void scrollText(int row, String text) {
  if (text.length() <= 16) {
    clearLine(row);
    lcd.setCursor(0, row);
    lcd.print(text);
    return;
  }
  
  String padded = "                " + text;
  for (int i = 0; i < padded.length() - 16; i++) {
    lcd.setCursor(0, row);
    lcd.print(padded.substring(i, i + 16));
    delay(250);
  }
}