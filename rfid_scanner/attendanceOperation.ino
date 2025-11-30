#include <HTTPClient.h>

#define JOY_NONE 0
#define JOY_UP 1
#define JOY_DOWN 2
#define JOY_LEFT 3
#define JOY_RIGHT 4
#define JOY_CLICK 5

int selectedClassId = 0;
String selectedClassName = "";
bool classSelected = false;

void selectClass() {
    if (!checkWifi()) {
        reconnectWifi();
        if (!checkWifi()) {
            displayBothLines("WiFi", "Disconnected");
            delay(2000);
            return;
        }
    }

    displayBothLines("Loading", "classes...");
    
    HTTPClient http;
    http.begin(String(BACKEND_URL) + "/api/classes/today");
    int responseCode = http.GET();
    
    if (responseCode != 200) {
        displayBothLines("Error", "No classes");
        delay(10000);
        return;
    }
    
    String response = http.getString();
    http.end();
    
    JsonDocument doc;
    deserializeJson(doc, response);
    JsonArray classes = doc["data"];
    
    int classCount = classes.size();
    
    if (classCount == 0) {
        displayBothLines("No classes", "today");
        delay(2000);
        return;
    }
    
    int classIds[10];
    String classNames[10];
    
    for (int i = 0; i < classCount && i < 10; i++) {
        classIds[i] = classes[i]["ID"];
        classNames[i] = classes[i]["name"].as<String>();
    }
    
    int selected = 0;
    displayBothLines("Select class:", classNames[selected]);
    
    while (true) {
        int joy = readJoystick();
        
        if (joy == JOY_UP && selected > 0) {
            selected--;
            displayLine2(classNames[selected]);
            delay(200);
        }
        if (joy == JOY_DOWN && selected < classCount - 1) {
            selected++;
            displayLine2(classNames[selected]);
            delay(200);
        }
        if (joy == JOY_CLICK) {
            selectedClassId = classIds[selected];
            selectedClassName = classNames[selected];
            classSelected = true;
            break;
        }
    }
}

void postAttendance(String cardId) {
    if (!checkWifi()) {
        reconnectWifi();
        if (!checkWifi()) {
            displayBothLines("WiFi", "Disconnected");
            delay(2000);
            return;
        }
    }

    displayBothLines("Checking...", cardId);
    
    HTTPClient http;
    http.begin(String(BACKEND_URL) + "/api/attendance/");
    http.addHeader("Content-Type", "application/json");
    
    String payload = "{\"cardId\":\"" + cardId + "\",\"classId\":" + String(selectedClassId) + "}";
    http.POST(payload);
    String response = http.getString();
    http.end();
    
    JsonDocument doc;
    deserializeJson(doc, response);
    
    String name = doc["name"] | "Unknown";
    bool success = doc["success"] | false;
    String error = doc["error"] | "";
    
    if (success) {
        displayBothLines("Welcome!", name);
    } else {
        displayLine1(name);
        displayLine2(error);
    }
    
    delay(2000);
    displayLine1(selectedClassName);
    displayLine2("Scan card...");
}

void attendanceOperation() {
    if (!classSelected) {
        selectClass();
        if (!classSelected) return;
        displayLine1(selectedClassName);
        displayLine2("Scan card...");
    }
    
    String card = readCard();
    
    if (card != "") {
        postAttendance(card);
    }
}