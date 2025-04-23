FROM golang:latest
WORKDIR /usr/local/app

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/goshort
EXPOSE 8080

CMD ["app"]
