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

  struct sockaddr_in client;
  client.sin_family = AF_INET;
  client.sin_port = htons(5000);
  client.sin_addr.s_addr = inet_addr("0.0.0.0");
  client.sin_addr.s_addr = INADDR_ANY;

  printf("Connecting");
  printf("[[[CLIENT]]]\n");

  int choice =0;
  do {

    char msgr[1000]={'\0'};
    printf("Enter the message\n");
    scanf(" %[^\n]s", msgr);

    HeapSize = 1;

    printf("Message received from [client]: %s\n", msgr);

    int hashMap[HASHMAP_CHARACTER_SIZE] = {0};
    getFreq(msgr, hashMap);
    execute_heap(msgr, hashMap);
    // printf("Message received from [server]: %s\n", payload.encode);


    sendto(fd, &payload, sizeof(payload), 0, (struct sockaddr *)&client, sizeof(client));


    printf("Waiting for servers response...\n");

    // int size = sizeof(client);
    // int rec_data = recvfrom(fd, &payload, sizeof(payload), 0, (struct sockaddr *)&client, &size);
    // assert(fd != -1);



    printf("Enter 0 to EXIT any other number to continue\n");
    scanf("%d", &choice);
    fflush(stdin);
    printf("\n");

  }while (choice != 0);

  printf("Client disconnected");
  close(fd);
  return 0;
}
