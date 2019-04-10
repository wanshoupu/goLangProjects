#include <stdlib.h>
#include <stdio.h>

void swap(int* a, int* b) {
    // printf("Swap: %d <--> %d\n", *a, *b);
    int t = *a;
    *a = *b;
    *b = t;
}

void assert(bool claim, const char* msg='\0') {
    if (!claim) {
        printf("%s", msg);
        exit(1);
    }
}

void print(int* s, int* e) {
    printf("\nIndexs: ");
    for(int* r = s; r < e; ++r)
        printf("%6ld, ", r);
    printf("\nValues: ");
    for(; s < e; s++) {
        printf("%6d, ", *s);
    }
    printf("\n");
}

void print(int a[], const int s) {
    print(a,a + s);
}

bool is_sorted(int* s, int* e) {
    print(s, e);
    for(s++; s < e; s++) {
        if (*(s - 1) > *s) return false;
    }
    return true;
}
