FROM gcc:12-bullseye

WORKDIR /src/server

COPY . .

RUN gcc -o client.out client.c

CMD ["./client.out"]

EXPOSE 8081