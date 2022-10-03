FROM golang:1.18

WORKDIR /go/src

#COPY go.sum .

COPY go.mod .

RUN go get -d || true

COPY . .

RUN go build -o server.out server.go
RUN go build -o client.out client.go
