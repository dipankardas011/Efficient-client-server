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

struct ServerClientPayload {
  struct TableHuffman *table;
  int sizeTable;
  char encode[1000];
} payload;

int main(int argc, char const *argv[])
{
  int fd = socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP);

  assert(fd != -1);

  struct sockaddr_in client;
  client.sin_family = AF_INET;
  client.sin_port = htons(5000);
  client.sin_addr.s_addr = inet_addr("0.0.0.0");
  // client.sin_addr.s_addr = INADDR_ANY;

  char msgr[1000] = "abaaabbcdhhjdch";
  printf("[[[CLIENT]]]\n");

  sendto(fd, msgr, strlen(msgr), 0, (struct sockaddr *)&client, sizeof(client));

  // char msg[1000] = {'\0'};
  int size = sizeof(client);
  int rec_data = recvfrom(fd, &payload, sizeof(payload), 0, (struct sockaddr *)&client, &size);
  assert(fd != -1);

  printf("Message received from [server]: %s\n", payload.encode);
  close(fd);
  return 0;
}
