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

void init_swap(int* s, int* e) {
    // could be as simple as swap(s, s + rand() % (e - s));
    for (int i = 0; i < 3; i++) {
        swap(s + i, s + i + rand() % (e - s - i));
    } // elect 3 elements and put them at positions: s, s+1, s+2
    if (*s < *(s + 1)) {
        swap(s, s + 1);
    }
    if (*(s + 2) < *(s + 1)) {
        swap(s + 1, s + 2);
    } // s + 1 has the minimum of the 3
    if (*(s + 2) < *s) {
        swap(s, s + 2);
    } // s, the pivot, has the median of the 3
    swap(e - 1, s + 2); // e - 1, has the maximum of the 3
}

int* part_original(int* s, int*e) {
    swap(s, s + rand() % (e - s));
    int* const pivot = s;
    while(true) {
         while(s < e && *++s < *pivot);  // find the next element >= pivot
         while(pivot < e && *--e > *pivot); // find the prev element <= pivot
         if (s < e) {
             swap(s, e); //swap the out-of-place elements
         } else {
             swap(pivot, e);
             return e;
         }
    }
}

int* part_thin_loop(int* s, int*e) {
    swap(s, s + rand() % (e - s));
    int* const pivot = s++;
    e--;
    while(true) {
         if (s < e && *s < *pivot) {  // find the next element >= pivot
             ++s;
         } else if (pivot < e && *e > *pivot) { // find the prev element <= pivot
             --e;
         } else if (s < e) {
             swap(s++, e--); //swap the out-of-place elements
         } else {
             swap(pivot, e);
             return e;
         }
    }
}

int* part_sentinel(int* s, int*e) {
    init_swap(s, e);
    int* const pivot = s;
    while(true) {
         while(*++s < *pivot);   // find the next element >= pivot
         while(*--e > *pivot); // find the prev element <= pivot
         if (s < e) {
             swap(s, e); //swap the out-of-place elements
         } else {
             swap(pivot, e);
             return e;
         }
    }
}

int* part1(int* s, int*e) {
    swap(s, s + rand() % (e - s));
    int * p = s;
    // printf("Pivot: %d\n", *p);
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
     if (e - s == 2) {
         if (*(e - 1) < *s)
             swap(s, e - 1);
         return;
     }
     int* p = part_thin_loop(s, e);
     // printf("whole arr: ");
     // print(s, e);
     // printf("left: ");
     // print(s, p);
     // printf("right: ");
     // print(p,e);
     qSort(s, p);
     qSort(p + 1, e);
}

bool is_sorted(int* s, int* e) {
    for(s++; s < e; s++) {
        if (*(s-1) > *s) return false;
    }
    return true;
}

void assert(bool claim, const char* msg) {
    if (!claim) {
        printf("%s", msg);
        exit(1);
    }
}

const char* msg = "Error: array not sorted properly\n";

void unit_test_0() {
    int test[] = {}; //
    const int size = 0;
    printf("Unit test: ");
    print(test, size);
    qSort(test, test + size);
    assert(is_sorted(test, test + size), msg);
}

void unit_test_1() {
    int test[] = {-1}; //
    const int size = 1;
    printf("Unit test: ");
    print(test, size);
    qSort(test, test + size);
    assert(is_sorted(test, test + size), msg);
}

void unit_test_2() {
    int test[] = {-1, -15}; //
    const int size = 2;
    printf("Unit test: ");
    print(test, size);
    qSort(test, test + size);
    assert(is_sorted(test, test + size), msg);
}

void unit_test_3() {
    int test[] = {15, -1, -15}; //
    const int size = 3;
    printf("Unit test: ");
    print(test, size);
    qSort(test, test + size);
    assert(is_sorted(test, test + size), msg);
}

void unit_test_big_cross() {
    int test[] = {-1, -15, -15, -15, -15, -15, -15, -15}; //
    const int size = 8;
    printf("Unit test: ");
    print(test, size);
    qSort(test, test + size);
    assert(is_sorted(test, test + size), msg);
}

void unit_test_eq() {
    int test[] = {-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,}; //
    const int size = 64;
    printf("Unit test: ");
    print(test, size);
    qSort(test, test + size);
    assert(is_sorted(test, test + size), msg);
}

void unit_test_4() {
    int test[] = {-15, 15, -1, -15}; //
    const int size = 4;
    printf("Unit test: ");
    print(test, size);
    qSort(test, test + size);
    assert(is_sorted(test, test + size), msg);
}

void unit_test_stepping() {
    const int size = 100;
    int test[size];
    for (int i = 0; i < size; ++i) {
        test[i] = rand();
    }
    printf("Unit test: ");
    print(test, size);
    qSort(test, test + size);
    assert(is_sorted(test, test + size), msg);
}

void unit_test_happy() {
    int test[] = {15,-1,-1,5, 10,1, 5, 1, 5, -10, 3, -4, 7, 77, -58, 5}; //
    const int size = 16;
    printf("Unit test: ");
    print(test, size);
    qSort(test, test + size);
    assert(is_sorted(test, test + size), msg);
}

int main() {
    unit_test_0();
    unit_test_1();
    unit_test_2();
    unit_test_3();
    unit_test_4();
    unit_test_stepping();
    unit_test_big_cross();
    unit_test_eq();
    unit_test_happy();
}
