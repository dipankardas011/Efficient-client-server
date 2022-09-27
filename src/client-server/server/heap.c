#include <stdio.h>
#include "heap.h"

/**
 * @brief HeapSize - It is the current size of the heap
 */
int HeapSize;

void swapNodeHeap(struct HeapDS *a, struct HeapDS *b){
  struct HeapDS t;
  t.character = a->character;
  t.frequency = a->frequency;
  t.left = a->left;
  t.right = a->right;

  a->character = b->character;
  a->frequency = b->frequency;
  a->left = b->left;
  a->right = b->right;

  b->character = t.character;
  b->frequency = t.frequency;
  b->left = t.left;
  b->right = t.right;
}

void heapify(struct HeapDS **arr, int i) {
  int l = 2*i;
  int r = 2*i + 1;
  int smallest = i;
  if (l < HeapSize && arr[l]->frequency < arr[smallest]->frequency)
    smallest = l;
  if (r < HeapSize && arr[r]->frequency < arr[smallest]->frequency)
    smallest = r;

  if (smallest == i)
    return;

  swapNodeHeap(arr[smallest], arr[i]);
  heapify(arr, smallest);
}

void buildHeap(struct HeapDS **arr) {
  for (int i = HeapSize/2; i >= 1; i--)
    heapify(arr, i);
}

struct HeapDS* popHeap(struct HeapDS **arr) {
  struct HeapDS* ret = arr[1];
  arr[1] = arr[HeapSize-1];
  // free(arr[HeapSize-1]);
  HeapSize--;
  heapify(arr, 1);
  return ret;
}

void pushHeap(struct HeapDS **arr, struct HeapDS *item) {
  HeapSize += 1;
  int i = HeapSize - 1;

  while (i > 1 && item->frequency < arr[i/2]->frequency) {
    arr[i] = arr[i/2];
    i /= 2;
  }
  arr[i] = item;
}

struct HeapDS *peekHeap(struct HeapDS **arr) {
  return HeapSize > 0 ? arr[1] : NULL;
}