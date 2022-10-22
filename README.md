# Efficient-client-server

A Efficient way to communicate between client-server.

**Client-Server** denotes a relationship between cooperating programs in an application,composed of clients initiating requests for services and servers providing that function or service.Basically is a distributed application framework dividing tasks between servers and clients, which either reside in the same system or communicate through a computer network or the Internet.

> `Client` what requests a service or resource from server.The server can be located on or off premises.

> `Server` a computer program that provides service to another computer program and its users.

## What is the need?
Efficient meaning working in such a way that gets the response from server with **least amount of bandwidth requirement**.

## How it accomplishes
Similarly this client server paradigm uses a bit encoding due to which message to be send between server and client reduces significantly. The client and server both uses a bit encoding technique by which messsage sent by client is encoded and sent to server, server decodes understands and again sends a response in encoded format then client recoving it decodes and displays the message to user

## Architecture
![image](https://user-images.githubusercontent.com/65275144/197322355-2cdbc655-2ad7-4987-abb4-2ecff90a1b9c.png)

# Running Application
![image](https://user-images.githubusercontent.com/65275144/197322418-3a48b8fc-723f-4b7c-b2fd-535a69f50ec5.png)

## How to run

```bash
cd src
docker build -t eff-server-client . --no-cache
docker run --rm -d -p 80:8080 eff-server-client
docker ps
```



