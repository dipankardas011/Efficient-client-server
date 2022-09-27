#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>

#include "heap.c"


#define HASHMAP_CHARACTER_SIZE (123-32)

#define IS_VALID_HASHMAP_CHARACTER(character) character == 32 || (character >= 65 && character <= 90) || (character >= 97 && character <= 122)

void execute_heap(int *hashMap) {
  /**
    * @brief HeapArr - It's a 1 based array of characters
    */
  struct HeapDS HeapArr[HEAP_ARR_SIZE];
  for (int i = 0; i < HASHMAP_CHARACTER_SIZE; i++) {
    if ((IS_VALID_HASHMAP_CHARACTER(i+32)) && hashMap[i] != 0) {
      HeapSize++;
    }
  }

  // printf("%d\n", HeapSize);
}



void getFreq(char* str, int* hashMap) {
  // Assume ASCII of lowercase characters, upper case characters, digits and spaces
  for (int i=0; str[i] != '\0'; i++) {
    if (IS_VALID_HASHMAP_CHARACTER(str[i])) {
      hashMap[str[i]-32]++;
    }
  }
}





int main(int argc, char **argv){
  HeapSize = 1;

  char *msg = "My name is Dipankar Das";
  int hashMap[HASHMAP_CHARACTER_SIZE]={0};
  getFreq(msg, hashMap);

  for (int i=0; i< HASHMAP_CHARACTER_SIZE; i++) {
    printf("%c -> %d\n", (i+32), hashMap[i]);
  }

  execute_heap(hashMap);

  return 0;
}