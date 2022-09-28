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



#define printerHeapArr                                                 \
  for (int i = 1; i < HeapSize; i++)                                   \
    printf("%c ; %d\n", HeapArr[i]->character, HeapArr[i]->frequency); \
  printf("\n\n");

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
  printf("\n\n");

  int tableSize = HeapSize - 1;
  // struct TableHuffman tableHuff[tableSize];
  struct TableHuffman* tableHuff = (struct TableHuffman*)malloc(sizeof(struct TableHuffman) * tableSize);
  // printerHeapArr

  buildHeap(HeapArr);

  // printerHeapArr

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

  printf("%s\n", newMessage);

  payload.sizeTable = tableSize;
  payload.table = tableHuff;
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

// int main(int argc, char **argv)
// {
//   HeapSize = 1;

//   char *msg = "abaaabbcdhhjdch";
//   printf("Message: %s\n", msg);
//   int hashMap[HASHMAP_CHARACTER_SIZE] = {0};
//   getFreq(msg, hashMap);
//   execute_heap(msg, hashMap);

//   printf("%s",payload.encode);

//   return 0;
// }