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


Client-Server denotes a relationship between cooperating programs in an application,composed of clients initiating requests for services and servers providing that function or service.Basically is a distributed application framework dividing tasks between servers and clients, which either reside in the same system or communicate through a computer network or the Internet.

Client:is what requests a service or resource from server.The server can be located on or off premises.
Server:is a computer program that provides service to another computer program and its users.

WORKING:
User runs the client and server terminal on the both sides and then sends a message from client side then server side sends the acknowledgement message to client and then a connection is established between both of them that is client and server.
The base concept is that User enters the URL(Uniform Resource Locator) of the website or file. The Browser then requests the DNS(DOMAIN NAME SYSTEM) Server.
DNS Server lookup for the address of the WEB Server. DNS Server responds with the IP address of the WEB Server. Browser sends over an HTTP/HTTPS request to WEB Server’s IP (provided by DNS server). Server sends over the necessary files of the website. Browser then renders the files and the website is displayed. 
