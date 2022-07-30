# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.18.4-alpine3.16

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy all the go files to the root
COPY . .

# This container exposes port 8080 to the outside world
EXPOSE 8080

CMD [ "/product-stock-web-service" ]