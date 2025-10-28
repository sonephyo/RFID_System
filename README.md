# RFID Scanner System

In order to facilitate an easier workflow for places that require attendance, the RFID scanner plays a role in minimizing the time to take attendance by utilizing different waves of frequencies to read information.

## How to set up project

```
git clone https://github.com/sonephyo/RFID_System
mv ./secrets.h.example ./rfid_scanner/secrets.h
```

Configure ./rfid_scanner/secrets.h with the right credentials

Install "Arduino_JSON" by Arduino by going to Sketch > Include Library > ESP32

## How to run backend

Get into the directory
```
cd rfid-backend
```

Setup env file (Note: Fill in the secret variables)
```
mv .env.example .env
```

Starting backend in development
```
CompileDaemon -command="./bin/rfid-backend" -build="go build -o ./bin"
```


## Configuration
If the tables need to be reset, run the following. (**Caution**: This operation can delete all existing data created)
```
go run migrate/migrate.go
```