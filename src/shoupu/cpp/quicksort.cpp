#include "sort_utils.h"

int part_simple(int* arr, int left, int right) {
    int pivot = arr[left];
    int i = left;
    for (int j = right; i < j;) {
        if (arr[j] >= pivot)
            j--;
        else if (arr[i] <= pivot)
            i++;
        else {
            int t = arr[i];
            arr[i] = arr[j];
            arr[j] = t;
        }
    }
    arr[left] = arr[i];
    arr[i] = pivot;
    return i;
}

int* part_nested_closed(int* left, int* right) {
    int* const pivot = left;
    while(true) {
        // note the order: right first, then left
        while (left < right && *right >= *pivot)
            right--;
        while (left < right && *left <= *pivot)
            left++;
        if (left < right) {
            swap(left, right);
        } else {
            swap(left, pivot);
            return left;
        }
    }
}

int* part_unnested_closed(int* left, int* right) {
    int* const pivot = left;
    while(true) {
        // note the order: right first, then left
        if (left < right && *right >= *pivot)
            right--;
        else if (left < right && *left <= *pivot)
            left++;
        else if (left < right)
            swap(left, right);
        else {
            swap(left, pivot);
            return left;
        }
    }
}

/*
do quicksort on array bound by two pointers left and right, both inclusive
*/
void quicksort(int* left, int* right) {
    if (left >= right)
       return;
    int* i = part_unnested_closed(left, right);
    quicksort(left, i - 1);
    quicksort(i + 1, right);
}

const char* msg = "Error: array not sorted properly\n";

void unit_test(int* test, int size) {
    print(test, size);
    quicksort(test, test + size - 1);
    assert(is_sorted(test + 1, test + size), msg);
    printf("Unit test success!\n========================\n");
}

void unit_test_0() {
    int test[] = {}; //
    const int size = 0;
    unit_test(test, size);
}

void unit_test_1() {
    int test[] = {-1}; //
    const int size = 1;
    unit_test(test, size);
}

void unit_test_2() {
    int test[] = {-1, -15}; //
    const int size = 2;
    unit_test(test, size);
}

void unit_test_3() {
    int test[] = {15, -1, -15}; //
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
    int test[] = {-15, 15, -1, -15}; //
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
