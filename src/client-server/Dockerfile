FROM gcc:12-bullseye

WORKDIR /app

COPY . .

LABEL MAINTAINER="Dipankar Das <dipankardas0115@gmail.com>"

RUN cd server && gcc -O0 -o ../server.out main_server.c
RUN cd client && gcc -O0 -o ../client.out main_client.c



# CMD ["./server.out"]

# EXPOSE 8080