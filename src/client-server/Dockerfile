FROM golang:1.18

WORKDIR /go/src

#COPY go.sum .

COPY go.mod .

RUN go get -d || true

COPY . .

RUN go build -o server.out server.go
RUN go build -o client.out client.go
RUN go build -o http.out http.go

#CMD ["./server.out &", "&&", "./client.out &", "&&", "./http.out"]
CMD ["./http.out"]

EXPOSE 8080
