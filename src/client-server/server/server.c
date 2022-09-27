#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>

#include "heap.c"

#define HASHMAP_CHARACTER_SIZE (123 - 32)

#define IS_VALID_HASHMAP_CHARACTER(character) character == 32 || (character >= 65 && character <= 90) || (character >= 97 && character <= 122)

#define printerHeapArr                                                 \
  for (int i = 1; i < HeapSize; i++)                                   \
    printf("%c ; %d\n", HeapArr[i]->character, HeapArr[i]->frequency); \
  printf("\n\n");



void execute_heap(int *hashMap)
{
  /**
   * @brief HeapArr - It's a 1 based array of characters
   */
  struct HeapDS *HeapArr[HEAP_ARR_SIZE] = {NULL};
  int idx = 1;

  for (int i = 0; i < HASHMAP_CHARACTER_SIZE; i++) {
    if ((IS_VALID_HASHMAP_CHARACTER(i + 32)) && hashMap[i] != 0) {
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
  getHuffmanCodes(root, str);
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

void getHuffmanCodes(struct HeapDS *root, char *str) {
  if (!root->left && !root->right) {
    printf("Code for %c -> %s\n", root->character, str);
    return;
  }

  char left[100]={'\0'};
  char right[100] = {'\0'};
  int i;
  for (i=0; str[i] != '\0'; i++) left[i] = str[i];
  left[i] = '0';
  for (i=0; str[i] != '\0'; i++) right[i] = str[i];
  right[i] = '1';

  getHuffmanCodes(root->left, left);
  getHuffmanCodes(root->right, right);
}


int main(int argc, char **argv)
{
  HeapSize = 1;

  // char msg[10000] = {'\0'};
  // printf("Enter the message\n");
  // scanf("%s",msg);
  char *msg = "My name is Dipankar";
  printf("Message: %s\n",msg);
  int hashMap[HASHMAP_CHARACTER_SIZE] = {0};
  getFreq(msg, hashMap);
  execute_heap(hashMap);

  return 0;
}