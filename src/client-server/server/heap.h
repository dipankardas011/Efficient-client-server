#define HEAP_ARR_SIZE 100

struct HeapDS {
  char character;
  int frequency;
  struct HeapDS *left;
  struct HeapDS *right;
};


void heapify(struct HeapDS *, int);

void buildHeap(struct HeapDS *);

struct HeapDS popHeap(struct HeapDS *);

void pushHeap(struct HeapDS *, struct HeapDS);

void swapNodeHeap(struct HeapDS *, struct HeapDS *);