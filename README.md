# Efficient-client-server

## Architecture
![Architecture](./architecture.png)

## Running complete Backend
![coverimage](./CoverImg.png)

## How to run

```bash
cd src
docker build -t eff-server-client . --no-cache
docker run --rm -d -p 80:8080 eff-server-client
docker ps
```
