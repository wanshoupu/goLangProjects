#include "sort_utils.h"

void init_swap(int* s, int* e) {
    // could be as simple as swap(s, s + rand() % (e - s));
    // elect 3 elements at position s, m, l = e-1, respectively
    // put the median at s, small at m, and large at l
    int* m = s + (e - s) / 2;
    int* l = e - 1;
    if (*l < *m) {
        swap(l, m);
    }
    if (*l < *s) {
        swap(l, s);
    }
    if (*s < *m) {
        swap(s, m);
    }
}

void rand_swap(int* s, int* e) {
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

int* part_nested_loop(int* s, int* e) {
    swap(s, s + rand() % (e - s));
    int* const pivot = s;
    while (true) {
        while (s < e && *++s < *pivot); // find the next element >= pivot
        while (pivot < e && *pivot < *--e); // find the prev element <= pivot
        if (s < e) {
            swap(s, e); //swap the out-of-place elements
        } else {
            swap(pivot, e);
            return e;
        }
    }
}

int* part_nested_loop_var3(int* s, int* e) {
    swap(s, s + rand() % (e - s));
    int* const pivot = s;
    while (true) {
        // '<=' or '<' is not critical. But '<=' brings more saving on unnecessary swaps.
        while (s < e && *++s <= *pivot);
        while (pivot < e && *pivot <= *--e);
        if (s < e) {
            swap(s, e);
        } else {
            swap(pivot, e);
            return e;
        }
    }
}

/**
 *Found another variation of partition scheme based on part_nested_loop
 *Diff wise, it lies between part_nested_loop and part_nested_loop_var
 *- the order of the two while loops is identical to part_nested_loop
 *- the boundary checking is identical to part_nested_loop_var
 *- the return value is shifted toward the left by one unit
 */
int* part_nested_loop_var2(int* s, int* e) {
    swap(s, s + rand() % (e - s));
    int* const pivot = s;
    while (true) {
        while (s < e && *++s <= *pivot); // find the next element >= pivot
        while (s < e && *pivot <= *--e); // find the prev element <= pivot
        if (s < e) {
            swap(s, e); //swap the out-of-place elements
        } else {
            swap(pivot, s - 1);
            return s - 1;
        }
    }
}

int* part_thin_loop_var2(int* s, int* e) {
    swap(s, s + rand() % (e - s));
    int* const pivot = s;
    ++s;
    while (true) {
        --e;
        while (s < e && *s <= *pivot) // find the next element >= pivot
        {
            ++s;
        }
        while (s < e && *pivot <= *e) {
            --e;
        }
        if (s < e) {
            swap(s++, e); //swap the out-of-place elements
        } else {
            swap(pivot, s - 1);
            return s - 1;
        }
    }
}

/**
 *Variation of partition scheme from part_nested_loop.
 *The two differs in three ways
 *- order of the two while loops
 *- increment condition for equal element pointer
 *- return pointer choice between s and e
 */
int* part_nested_loop_var(int* s, int* e) {
    swap(s, s + rand() % (e - s));
    int* const pivot = s;
    while (true) {
        while (s < e && *pivot <= *--e); // find the prev element <= pivot
        while (s < e && *++s <= *pivot); // find the next element >= pivot
        if (s < e) {
            swap(s, e); //swap the out-of-place elements
        } else {
            swap(pivot, s);
            return s;
        }
    }
}

/*
Find yet another variation of implementation where the second nested while-loop
is made post-incremental.
This will remove the difficulty encountered in loop-thinning.
*/
int* part_nested_loop_var4(int* s, int* e) {
    swap(s, s + rand() % (e - s));
    int* const pivot = s;
    while (true) {
        while (s < e && *pivot <= *--e); // find the prev element <= pivot
        while (s < e && *s <= *pivot) ++s; // find the next element >= pivot
        if (s < e) {
            swap(s, e); //swap the out-of-place elements
        } else {
            swap(pivot, s);
            return s;
        }
    }
}

/*
loop-thinned version of part_nested_loop_var4
*/
int* part_thin_loop_var4(int* s, int* e) {
    swap(s, s + rand() % (e - s));
    int* const pivot = s;
    --e;
    while (true) {
        if (s < e && *pivot <= *e) --e; // find the prev element <= pivot
        else if (s < e && *s <= *pivot) ++s; // find the next element >= pivot
        else if (s < e) {
            swap(s, e--); //swap the out-of-place elements
        } else {
            swap(pivot, s);
            return s;
        }
    }
}

int* part_thin_loop(int* s, int* e) {
    swap(s, s + rand() % (e - s));
    int* const pivot = s++;
    e--;
    while (true) {
        if (s < e && *s <= *pivot) { // find the next element >= pivot
            ++s;
        } else if (pivot < e && *pivot <= *e) { // find the prev element <= pivot
            --e;
        } else if (s < e) {
            swap(s++, e--); //swap the out-of-place elements
        } else {
            swap(pivot, e);
            return e;
        }
    }
}

int* part_sentinel(int* s, int* e) {
    init_swap(s, e);
    int* const pivot = s;
    while (true) {
        while ( *++s < *pivot); // find the next element >= pivot
        while ( *--e > *pivot); // find the prev element <= pivot
        if (s < e) {
            swap(s, e); //swap the out-of-place elements
        } else {
            swap(pivot, e);
            return e;
        }
    }
}

int* part(int* s, int* e) {
    init_swap(s, e);
    int* const pivot = s++;
    --e;
    while (true) {
        if (*s < *pivot) { // find the next element >= pivot
            ++s;
        } else if (*e > *pivot) { // find the prev element <= pivot
            --e;
        } else if (s < e) {
            swap(s++, e--); //swap the out-of-place elements
        } else {
            swap(pivot, e);
            return e;
        }
    }
}

/**
This is the Lomuto's partitioning scheme
*/
int* part1(int* s, int* e) {
    swap(s, s + rand() % (e - s));
    int* p = s;
    // printf("Pivot: %d\n", *p);
    for (int* i = s; i < e; i++) {
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
    int* p = part_nested_loop_var4(s, e);
    qSort(s, p);
    qSort(p + 1, e);
}

const char *msg = "Error: array not sorted properly\n";

void unit_test(int* test, int size) {
    print(test, size);
    qSort(test, test + size);
    assert(is_sorted(test, test + size), msg);
    printf("Unit test success!\n========================\n");
}

void unit_test_0() {
    int test[] = {}; //
    const int size = 0;
    unit_test(test, size);
}

void unit_test_1() {
    int test[] = {
        -1
    }; //
    const int size = 1;
    unit_test(test, size);
}

void unit_test_2() {
    int test[] = {
        -1,
        -15
    }; //
    const int size = 2;
    unit_test(test, size);
}

void unit_test_3() {
    int test[] = {
        15,
        -1,
        -15
    }; //
    const int size = 3;
    unit_test(test, size);
}

void unit_test_big_cross() {
    int test[] = {-1, -15, -15, -15, -15, -15, -15, -15}; //
    const int size = 8;
    unit_test(test, size);
}

void unit_test_eq() {
    int test[] = {-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,-15, -15, -15,}; //
    const int size = 64;
    unit_test(test, size);
}

void unit_test_4() {
    int test[] = {
        -15,
        15,
        -1,
        -15
    }; //
    const int size = 4;
    unit_test(test, size);
}

void unit_test_stepping() {
    const int size = 100;
    int test[size];
    for (int i = 0; i < size; ++i) {
        test[i] = rand();
    }
    unit_test(test, size);
}

void unit_test_happy() {
    int test[] = {15,-1,-1,5, 10,1, 5, 1, 5, -10, 3, -4, 7, 77, -58, 5}; //
    const int size = 16;
    unit_test(test, size);
}

void unit_left_bound_exception() {
    int test[] = {
        15,
        -77,
        7,
        -58,
        5
    }; //
    const int size = 5;
    unit_test(test, size);
}

void unit_test_limited_rand() {
    const int size = 100;
    int test[size];
    for (int i = 0; i < size; ++i) {
        test[i] = rand() % 3;
    }
    unit_test(test, size);
}

int main() {
    printf("unit_left_bound_exception");
    unit_left_bound_exception();
    printf("unit_test_0");
    unit_test_0();
    printf("unit_test_1");
    unit_test_1();
    printf("unit_test_2");
    unit_test_2();
    printf("unit_test_3");
    unit_test_3();
    printf("unit_test_4");
    unit_test_4();
    printf("unit_test_stepping");
    unit_test_stepping();
    printf("unit_test_big_cross");
    unit_test_big_cross();
    printf("unit_test_eq");
    unit_test_eq();
    printf("unit_test_limited_rand");
    unit_test_limited_rand();
    printf("unit_test_happy");
    unit_test_happy();
}
