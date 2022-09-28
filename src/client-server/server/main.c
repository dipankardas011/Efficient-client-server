#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <netinet/ip.h>
#include <arpa/inet.h>
#include <unistd.h>
#include "server.c"

int main(int argc, char const *argv[])
{
  int fd = socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP);

  assert(fd != -1);

  struct sockaddr_in server;
  server.sin_family = AF_INET;
  server.sin_port = htons(5000);
  server.sin_addr.s_addr = inet_addr("0.0.0.0");
  printf("[[[SERVER]]]\n");

  int b = bind(fd, (struct sockaddr *)&server, sizeof(server));

  char msg[1000] = {'\0'};
  int size = sizeof(server);
  int rec_data = recvfrom(fd, msg, sizeof(msg), 0, (struct sockaddr *)&server, &size);
  assert(rec_data != -1);

  HeapSize = 1;

  printf("Message received from [server]: %s\n", msg);

  int hashMap[HASHMAP_CHARACTER_SIZE] = {0};
  getFreq(msg, hashMap);
  execute_heap(msg, hashMap);

  printf("Encoded: %s\nTableSize: %d\n",payload.encode, payload.sizeTable);

  sendto(fd, &payload, sizeof(payload), 0, (struct sockaddr *)&server, sizeof(server));


  return 0;
}