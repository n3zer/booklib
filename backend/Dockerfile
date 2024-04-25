FROM golang:1.22

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest
COPY . .

RUN go get github.com/gin-gonic/gin
RUN go get gorm.io/gorm
RUN go get gorm.io/driver/postgres

RUN go mod tidy