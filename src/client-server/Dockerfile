FROM gcc:12-bullseye

WORKDIR /app

COPY . .

LABEL MAINTAINER="Dipankar Das <dipankardas0115@gmail.com>"

RUN cd server && gcc -O1 -o ../server.out main.c
RUN cd client && gcc -O1 -o ../client.out client.c



# CMD ["./server.out"]

# EXPOSE 8080