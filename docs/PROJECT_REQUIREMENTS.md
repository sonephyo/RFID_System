# Class Attendance Tracking System - Project Requirements

## Overview
An RFID-based attendance system that eliminates remote attendance fraud by requiring physical presence for check-in.

## Hardware
- Arduino Uno (RFID reading)
- ESP32 (WiFi & communication)
- RFID-RC522 module
- LCD 1602 module (user interface display)
- Navigation input ( joystick module)
- Admin & Student RFID cards

## System Modes

### 1. Configuration Mode (USB Connected)
- Admin card authentication
- Access personalized dashboard
- Define classes (name, schedule, time)
- Manage class settings

### 2. Attendance Mode (Standalone)
- WiFi-connected operation
- Professor selects active class
- Students scan cards to check-in
- Real-time HTTP requests to backend
- Store attendance in database

### 3. Viewing Mode (Web Interface)
- Admin login to frontend
- View attendance records
- Filter by class, date, student
- Export/report generation

## Technical Components

### Device Firmware
- LCD 1602 menu system for navigation
- Button/joystick input handling
- USB serial communication for configuration
- WiFi connectivity management
- RFID card reading & validation
- HTTP client for API requests
- Class selection interface

### Backend API
- Admin authentication endpoints
- Class management (CRUD)
- Attendance recording endpoint
- Attendance retrieval/query endpoints
- Database integration

### Frontend Dashboard
- Admin login
- Display attendance data
- Class management interface
- Reports and analytics

### Database Schema
- Admins/Professors table
- Classes table
- Students table
- Attendance records table

## Key Features
- Physical presence verification via RFID
- Multi-professor support
- Scheduled class sessions
- Real-time attendance logging
- Historical attendance viewing
- Fraud prevention through hardware requirement