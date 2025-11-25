# RFID Scanner System

An automated attendance system that uses RFID technology to quickly and efficiently track attendance, ready to set up out of the box for different events.

1. **Arduino Uno** reads the RFID card and extracts the card UID
2. **ESP32** receives the UID from Arduino Uno via serial communication
3. **Frontend** connects to ESP32 via Web Serial API and displays the scanned data
4. **Backend** stores and manages authorization and tracking. 

## System Architecture

```
RFID Card → Arduino Uno → ESP32 → Frontend
                                    ↓
                                 Backend → Database
```

- **Arduino Uno**: Reads RFID cards using MFRC522 module
- **ESP32**: Acts as a bridge between Arduino Uno and the frontend
- **Frontend**: React application that connects to ESP32 via Web Serial API
- **Backend**: Go/Gin REST API for data management
- **Database**: Postgresql for storing data

## 📦 Prerequisites

- **Arduino IDE** (latest version) - [Download](https://www.arduino.cc/en/software)
- **ESP32 Board Support** for Arduino IDE - [Installation Guide](https://docs.espressif.com/projects/arduino-esp32/en/latest/installing.html)
  - Settings > Additional Board Managers URLs
  - Put [https://espressif.github.io/arduino-esp32/package_esp32_index.json](https://espressif.github.io/arduino-esp32/package_esp32_index.json)
- **Node.js** (v18 or higher) - [Download](https://nodejs.org/)
- **Go** (v1.19 or higher) - [Download](https://go.dev/doc/install)
- **Git** - [Download](https://git-scm.com/downloads)

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/sonephyo/RFID_System
cd CSC473-MainProject
```

### 2. Arduino Uno Setup

The Arduino Uno reads RFID cards and sends the card UID to the ESP32 via serial communication.

#### Required Libraries

Install the following library in Arduino IDE:

1. Open Arduino IDE
2. Go to **Sketch** → **Include Library** → **Manage Libraries...**
3. Search for **"MFRC522"** by GithubCommunity and install
4. Search for **ArduinoJson** by Benoit Blanchon and install

#### Hardware Connections

Connect the MFRC522 RFID module to Arduino Uno:
- **SDA (SS)** → Pin 10
- **SCK** → Pin 13
- **MOSI** → Pin 11
- **MISO** → Pin 12
- **RST** → Pin 9
- **3.3V** → 3.3V
- **GND** → GND

#### Upload Code

1. Connect your Arduino Uno to your computer via USB
2. Open `arduino_uno/arduino_uno.ino` in Arduino IDE
3. Select **Tools** → **Board** → **Arduino Uno**
4. Select the correct **Port** under **Tools** → **Port**
5. Click **Upload** (or press `Ctrl+U` / `Cmd+U`)

The Arduino Uno will now read RFID cards and print the card UID to Serial at 9600 baud.

### 3. ESP32 Setup

The ESP32 reads data from the Arduino Uno and forwards it to the frontend via Web Serial API.

#### Required Libraries

The ESP32 code uses built-in libraries, so no additional installation is needed.

#### Hardware Connections

Connect Arduino Uno to ESP32:
- **Arduino Uno TX** → **ESP32 Pin 16 (RXp2)**
- **Arduino Uno RX** → **ESP32 Pin 17 (TXp2)**
- **GND** → **GND** (common ground)


<img src="https://github.com/user-attachments/assets/229e0cfe-6fa9-42f7-a987-3b080297a009" width="400">

#### Configuration

1. Copy the secrets template:
   ```bash
   cp secrets.h.example rfid_scanner/secrets.h
   ```

2. Edit `rfid_scanner/secrets.h` with your WiFi credentials (Use your mobile hospot for development):
   ```cpp
   #define SSID "<wifi_name>"
   #define PASS "<wifi_password>"
   ```

#### Upload Code

1. Connect your ESP32 to your computer via USB
2. Open `rfid_scanner/rfid_scanner.ino` in Arduino IDE
3. Select **Tools** → **Board** → **ESP32 Dev Module** (or your specific ESP32 board)
4. Select the correct **Port** under **Tools** → **Port**
5. Click **Upload** (or press `Ctrl+U` / `Cmd+U`)

The ESP32 will now read from Serial2 (connected to Arduino Uno) and print to Serial (USB) at 115200 baud.

### 4. Frontend Setup

The React frontend connects to the ESP32 via Web Serial API and displays scanned RFID data.

#### Installation

1. Navigate to the frontend directory:
   ```bash
   cd rfid-frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

#### Running the Frontend

Start the development server:
```bash
npm run dev
```

The frontend will be available at `http://localhost:5173` (or the port shown in the terminal).

**Note**: To connect to the ESP32, you'll need to use the Web Serial API in your browser. Make sure your frontend code implements the serial connection to read data from the ESP32.

### 5. Backend Setup

The Go backend provides REST API endpoints for managing users and attendance (currently not connected to the frontend).

#### Installation

1. Navigate to the backend directory:
   ```bash
   cd rfid-backend
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

#### Configuration

1. Create a `.env` file (if `.env.example` exists, copy it):
   ```bash
   cp .env.example .env
   ```

2. Edit `.env` with your database credentials and other configuration:
   ```env
   DB_HOST=localhost
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_NAME=rfid_db
   DB_PORT=5432
   ```

#### Database Migration

If you need to reset the database tables (**Caution**: This will delete all existing data):
```bash
go run migrate/migrate.go
```

#### Running the Backend

**Development mode** (with auto-reload):
```bash
CompileDaemon -command="./bin/rfid-backend" -build="go build -o ./bin"
```

**Production mode**:
```bash
go build -o ./bin/rfid-backend
./bin/rfid-backend
```

The backend will run on the default port (usually `:8080`).

#### API Documentation

Once the backend is running, access the Swagger documentation at:
```
http://localhost:8080/swagger/index.html
```

To compile/update the Swagger documentation:
```bash
./scripts/compile_docs.sh
```

**Note**: When developing new APIs, follow the [gin-swagger guidelines](https://github.com/swaggo/gin-swagger) to add proper documentation.

## 🔌 Hardware Connections

### Complete Wiring Diagram

```
RFID Module (MFRC522)          Arduino Uno          ESP32
─────────────────────          ───────────          ─────
SDA (SS)  ────────────→  Pin 10
SCK       ────────────→  Pin 13
MOSI      ────────────→  Pin 11
MISO      ────────────→  Pin 12
RST       ────────────→  Pin 9
3.3V      ────────────→  3.3V
GND       ────────────→  GND

Arduino Uno                  ESP32
───────────                  ─────
TX         ────────────→  Pin 16 (RXp2)
RX         ────────────→  Pin 17 (TXp2)
GND        ────────────→  GND
```

### LCD 1602 Display (Connected to ESP32)

Connect the LCD 1602 (16-pin version) to ESP32:

| LCD Pin | Connection |
|---------|------------|
| 1 (VSS) | GND |
| 2 (VDD) | 5V |
| 3 (V0)  | Potentiometer middle pin (or GND for max contrast) |
| 4 (RS)  | GPIO 22 |
| 5 (RW)  | GND |
| 6 (E)   | GPIO 23 |
| 7-10 (D0-D3) | Not connected |
| 11 (D4) | GPIO 5 |
| 12 (D5) | GPIO 18 |
| 13 (D6) | GPIO 19 |
| 14 (D7) | GPIO 21 |
| 15 (A)  | 5V |
| 16 (K)  | GND |

**Potentiometer (for contrast adjustment):**
- Outer pin 1 → 5V
- Outer pin 2 → GND
- Middle pin → LCD Pin 3 (V0)

**Note:** Each potentiometer leg must be in a different row on the breadboard to avoid short circuit.

## 🏃 Running the System

### Complete Workflow

1. **Upload Arduino Uno code** and ensure it's powered
2. **Upload ESP32 code** and connect it to Arduino Uno
3. **Start the frontend**:
   ```bash
   cd rfid-frontend
   npm run dev
   ```
4. **Start the backend** (optional, currently not connected):
   ```bash
   cd rfid-backend
   CompileDaemon -command="./bin/rfid-backend" -build="go build -o ./bin"
   ```
5. **Open the frontend** in your browser and connect to the ESP32 via Web Serial API
6. **Scan an RFID card** - the UID should appear in the frontend

## 📁 Project Structure

```
CSC473-MainProject/
├── arduino_uno/
│   └── arduino_uno.ino          # Arduino Uno code (reads RFID cards)
├── rfid_scanner/
│   ├── rfid_scanner.ino          # ESP32 code (bridges Arduino to frontend)
│   └── secrets.h                 # WiFi credentials (create from secrets.h.example)
├── rfid-frontend/
│   ├── src/
│   │   ├── App.tsx               # Main React component
│   │   ├── pages/                # Page components
│   │   └── components/           # Reusable components
│   ├── package.json              # Frontend dependencies
│   └── vite.config.ts            # Vite configuration
├── rfid-backend/
│   ├── main.go                   # Backend entry point
│   ├── controllers/              # API controllers
│   ├── models/                   # Data models
│   ├── initializers/             # Database and env initialization
│   ├── migrate/                  # Database migration scripts
│   └── docs/                     # Swagger documentation
├── secrets.h.example             # Template for secrets.h
└── README.md                     # This file
```

## 📝 Development Guidelines

### Code Style

- **Arduino/ESP32**: Follow Arduino coding conventions
- **Frontend**: Use TypeScript, follow React best practices
- **Backend**: Follow Go conventions and use `gofmt`

### API Documentation

When adding new API endpoints:
1. Add Swagger annotations to your handler functions
2. Run `./scripts/compile_docs.sh` to regenerate documentation
3. Verify the docs at `/swagger/index.html`

## 🔧 Troubleshooting

### Arduino Uno Issues

- Click RESET button on both UNO and ESP32 to ensure a refresh state
- When uploading code to either UNO or ESP32, disconnect the TX and RX to create a sucessful upload

### ESP32 Issues

- **Not receiving data from Arduino**:
  - Check Serial2 connections (TX/RX pins)
  - Verify both devices share common GND
  - Check baud rates match (9600)
- **Upload fails**:
  - Hold BOOT button during upload
  - Check USB cable supports data transfer
  - Verify correct board and port selected
- Close the Serial Monitor on ESP32 for it to be readible by the browser


### Frontend Issues

- **Cannot connect to ESP32**:
  - Check browser supports Web Serial API (Chrome/Edge recommended)
- **Dependencies not installing**:
  - Delete `node_modules` and `package-lock.json`, then run `npm install` again

### Backend Issues

- **Database connection fails**:
  - Verify `.env` file has correct credentials
  - Ensure database server is running
  - Check database exists
- **Swagger docs not updating**:
  - Run `./scripts/compile_docs.sh`
  - Check Swagger annotations are correct

## 🤝 Contributing

1. Create a new branch for your feature
2. Make your changes following the development guidelines
3. Test your changes thoroughly

4. Submit a pull request with a clear description
