#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <netinet/ip.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <assert.h>
#include "server.c"

int main(int argc, char const *argv[])
{
  int fd = socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP);

  assert(fd != -1);

  struct sockaddr_in server;
  server.sin_family = AF_INET;
  server.sin_port = htons(5000);
  server.sin_addr.s_addr = inet_addr("0.0.0.0");
  server.sin_addr.s_addr = INADDR_ANY;

  int b = bind(fd, (struct sockaddr *)&server, sizeof(server));
  assert(b != -1);

  printf("Connecting");
  printf("[[[SERVER]]]\n");

  int choice =0;
  do {

    int size = sizeof(server);

    int rec_data = recvfrom(fd, &payload, sizeof(payload), 0, (struct sockaddr *)&server, &size);
    assert(fd != -1);

    printf("Message received from [client]: %s\n", payload.encode);

    printerHuffmanTable(payload.table, payload.sizeTable);

    char encoded_data[1000] = {'\0'};
    ///////////////////////////////
    //TODO: Added decoder of huffman encoding


    ////////////////////////////////
    // THIS IS FOR DECODING OF MESSAGE
    // HeapSize = 1;
    // char msgr[1000]={'\0'};

    // for (int i=0; payload.encode[i] != '\0'; i++)
    //   msgr[i] = payload.encode[i];

    // int hashMap[HASHMAP_CHARACTER_SIZE] = {0};
    // getFreq(msgr, hashMap);
    // execute_heap(msgr, hashMap);

    // sendto(fd, &payload, sizeof(payload), 0, (struct sockaddr *)&server, sizeof(server));

  }while (1);

  printf("Server disconnected");
  close(fd);
  return 0;
}
