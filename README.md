# Efficient-client-server

## Architecture
![Architecture](./architecture.png)

## Running complete Backend
![coverimage](./CoverImg.png)

## How to run

```bash
cd src/client-server
docker build -t abc .

docker run --rm -d -p 8080:8080 backend
# server

docker exec -it $(docker ps -q) bash
./server.out

# client
docker exec -it $(docker ps -q) bash
./client.out

```
