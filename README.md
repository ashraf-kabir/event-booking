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
GOOS=linux GOARCH=amd64 go build -o build/linux/event-booking
```

### Give execute permissions to the binary:
```bash
cd build/linux
sudo chmod +x event-booking
```

### To install PM2 globally:
```bash
npm install -g pm2
```

### Serve using the PM2 process manager:
```bash
cd /path/to/your/project
pm2 start build/linux/event-booking --name event-booking
```

### To stop the PM2 process:
```bash
pm2 stop event-booking
```

### To restart the PM2 process:
```bash
pm2 restart event-booking
```

### To view PM2 logs:
```bash
pm2 logs event-booking
```

### To delete PM2 process:
```bash
pm2 delete event-booking
```

### To clean all PM2 logs:
```bash
pm2 flush
```

### NGINX Configuration:
```nginx
server {
    listen 80;
    server_name yourdomain.com;  # or public IP for testing

    location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection keep-alive;
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
```

### Check & reload NGINX configuration:
```bash
sudo nginx -t && sudo systemctl reload nginx
```

### SSL Configuration with Let's Encrypt and Certbot:
```bash
sudo certbot --nginx -d yourdomain.com
```
### To renew SSL certificates:
```bash
sudo certbot renew --dry-run
```
