FROM gcc:12-bullseye

WORKDIR /src/server

COPY . .

RUN gcc -o server.out server.c

CMD ["./server.out"]

EXPOSE 8080