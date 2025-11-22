// this sample code provided by www.programmingboss.com
#define RXp2 16
#define TXp2 17
void setup() {
  // put your setup code here, to run once:
  Serial.begin(115200);
  Serial2.begin(9600, SERIAL_8N1, RXp2, TXp2);
}
void loop() {
  String idCode = Serial2.readString();
  if (!idCode.isEmpty()) {
    Serial.printf("ID detected: %s", idCode);
  }
}