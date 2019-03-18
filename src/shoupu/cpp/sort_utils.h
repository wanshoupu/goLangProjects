#include <stdlib.h>
#include <stdio.h>

void swap(int* a, int* b) {
    // printf("Swap: %d <--> %d\n", *a, *b);
    int t = *a;
    *a = *b;
    *b = t;
}

bool is_sorted(int* s, int* e) {
    for(int* r = s; r < e; ++r)
        printf("%d (%d), ", r - s, *r);
    printf("\n");
    for(s++; s < e; s++) {
        if (*(s - 1) > *s) return false;
    }
    return true;
}

void assert(bool claim, const char* msg) {
    if (!claim) {
        printf("%s", msg);
        exit(1);
    }
}

void print(int* s, int* e) {
    printf("Array(%ld): \n", e - s);
    for(; s < e; s++) {
        printf("%d, ", *s);
    }
    printf("\n");
}

void print(int a[], const int s) {
    print(a,a + s);
}
