FROM golang:1.18

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build cmd/main.go

EXPOSE 8080

RUN mv main /usr/local/bin/go-echo-server
CMD ["go-echo-server"]
