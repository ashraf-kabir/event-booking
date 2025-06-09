# event-booking

### Install Dependencies:
```go
go get -u github.com/gin-gonic/gin
go get github.com/mattn/go-sqlite3
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v5
```

### To serve the project
```go
go run .
```

### To build the project for development:
```bash
cd scripts
./build.sh
```

### To serve the project in development mode:
```bash
cd scripts
./serve.sh
```

### Build the project for production (Linux amd64):
```bash
GOOS=linux GOARCH=amd64 go build -o build/event-booking-linux
```

### Give execute permissions to the binary:
```bash
sudo chmod +x event-booking
```

### Serve using the PM2 process manager:
```bash
cd /path/to/your/project
pm2 start event-booking --name event-booking
```
### To stop the PM2 process:
```bash
pm2 stop event-booking
```

### To restart the PM2 process:
```bash
pm2 restart event-booking
```