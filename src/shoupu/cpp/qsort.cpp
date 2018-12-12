#include <stdlib.h>
#include <stdio.h>

void swap(int* a, int* b) {
    // printf("Swap: %d <--> %d\n", *a, *b);
    int t = *a;
    *a = *b;
    *b = t;
}

void print(int* s, int* e) {
    printf("Array(%ld): ", e-s);
    for(; s < e; s++) {
        printf("%d, ", *s);
    }
    printf("\n");
}

void print(int a[], const int s) {
    print(a,a+s);
}

int* part(int* s, int*e) {
    int* const pivot = s;
    swap(s, s + rand() % (e - s));
    while(true) {
         // the left partition is strictly less than the pivot
         while(s < e && *s < *pivot) {  // find the next element >= pivot
             ++s;
         }
         // the right partition is  greater than or equal to the pivot
         while(s < e && *--e > *pivot); // find the prev element <= pivot
         if (e > s) {
             swap(s, e); //swap the out-of-place elements
         } else {
             return e;
         }
    }
}

int* part1(int* s, int*e) {
    int * p = s + rand() % (e - s);
    swap(s, p);
    p = s;
    for(int* i = s; i < e; i++ ) {
        if (*i < *s) {
            swap(++p, i); //swap the out-of-place elements
        }
    }
    swap(s, p);
    return p;
}

void qSort(int* s, int* e) {
     if (e - s < 2) return;
     int* p = part(s, e);
     qSort(s, p);
     qSort(p + 1, e);
}

int main() {
    int test[] = {15,-1,-1,5, 10,1, 5, 1, 5, -10, 3, -4, 7, 77, -58, 5}; //
    const int size = 16;
    print(test, size);
    qSort(test, test + size);
    print(test, size);
}
