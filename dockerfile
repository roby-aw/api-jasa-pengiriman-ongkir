FROM golang:1.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go test ./...
RUN go mod verify
RUN go build -o /dist

EXPOSE 8080

CMD ["/dist"]