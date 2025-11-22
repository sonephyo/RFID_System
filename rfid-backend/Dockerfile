
FROM golang:1.25-alpine3.22
 
WORKDIR /app

COPY . .
 
RUN go mod download
 
RUN go build -o ./bin/rfid-backend
 
EXPOSE 8080
 
CMD ["/app/bin/rfid-backend"]