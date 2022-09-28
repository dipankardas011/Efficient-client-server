#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>

#include "heap.c"

#define HASHMAP_CHARACTER_SIZE (123 - 32)

#define IS_VALID_HASHMAP_CHARACTER(character) character == 32 || (character >= 65 && character <= 90) || (character >= 97 && character <= 122)

struct TableHuffman
{
  char character;
  char encode[1000];
};

struct ServerClientPayload {
  struct TableHuffman *table;
  int sizeTable;
  char encode[1000];
} payload;

void getHuffmanCodes(struct HeapDS *root, char *str, struct TableHuffman *table, int *index)
{
  if (!root->left && !root->right)
  {
    table[*index].character = root->character;
    int i;
    for (i = 0; str[i] != '\0'; i++)
      table[*index].encode[i] = str[i];
    table[*index].encode[i] = '\0';
    *index = *index + 1;
    return;
  }

  char left[100] = {'\0'};
  char right[100] = {'\0'};
  int i;
  for (i = 0; str[i] != '\0'; i++)
    left[i] = str[i];
  left[i] = '0';
  for (i = 0; str[i] != '\0'; i++)
    right[i] = str[i];
  right[i] = '1';

  getHuffmanCodes(root->left, left, table, index);
  getHuffmanCodes(root->right, right, table, index);
}


void execute_heap(char *message, int *hashMap)
{
  /**
   * @brief HeapArr - It's a 1 based array of characters
   */
  struct HeapDS *HeapArr[HEAP_ARR_SIZE] = {NULL};
  int idx = 1;

  for (int i = 0; i < HASHMAP_CHARACTER_SIZE; i++)
  {
    if ((IS_VALID_HASHMAP_CHARACTER(i + 32)) && hashMap[i] != 0)
    {
      HeapSize++;
      HeapArr[idx] = (struct HeapDS *)malloc(sizeof(struct HeapDS));
      HeapArr[idx]->character = (char)(i + 32);
      HeapArr[idx]->frequency = hashMap[i];
      HeapArr[idx]->left = NULL;
      HeapArr[idx]->right = NULL;
      idx++;
    }
  }

  int tableSize = HeapSize - 1;
  struct TableHuffman* tableHuff = (struct TableHuffman*)malloc(sizeof(struct TableHuffman) * tableSize);

  buildHeap(HeapArr);

  while (HeapSize > 2)
  {
    struct HeapDS *x = popHeap(HeapArr);
    struct HeapDS *y = popHeap(HeapArr);

    struct HeapDS *z = (struct HeapDS *)malloc(sizeof(struct HeapDS));

    z->frequency = x->frequency + y->frequency;
    z->left = x;
    z->right = y;

    pushHeap(HeapArr, z);
  }
  struct HeapDS *root = peekHeap(HeapArr);

  char str[100];
  int index = 0;
  fflush(stdout);
  fflush(stdin);
  getHuffmanCodes(root, str, tableHuff, &index);
  char newMessage[10000] = {'\0'};
  int ii = 0;
  for (int i = 0; message[i] != '\0'; i++)
  {
    for (int j = 0; j < tableSize; j++)
    {
      if (message[i] == tableHuff[j].character) {
        for (int k = 0; tableHuff[j].encode[k] != '\0'; k++)
          newMessage[ii++] = tableHuff[j].encode[k];
        break;
      }

    }
  }

  payload.sizeTable = tableSize;
  payload.table = (struct TableHuffman *)malloc(sizeof(struct TableHuffman)*tableSize);
  for (int i=0; i<tableSize; i++) {
    payload.table[i].character = tableHuff[i].character;
    int j;
    for (j=0; tableHuff[i].encode[j] != '\0'; j++)
      payload.table[i].encode[j] = tableHuff[i].encode[j];
    payload.table[i].encode[j] = '\0';
  }
  // payload.table = tableHuff;
  for (int i=0; newMessage[i] !='\0'; i++)
    payload.encode[i] = newMessage[i];
}

void getFreq(char *str, int *hashMap)
{
  // Assume ASCII of lowercase characters, upper case characters, digits and spaces
  for (int i = 0; str[i] != '\0'; i++)
  {
    if (IS_VALID_HASHMAP_CHARACTER(str[i]))
    {
      hashMap[str[i] - 32]++;
    }
  }
}



void printerHuffmanTable(struct TableHuffman* table, int size) {
  printf("Character\tEncoded\n");
  for (int i = 0; i < size; i++) {
    printf("%c\t%s\n", table[i].character, table[i].encode);
  }
}