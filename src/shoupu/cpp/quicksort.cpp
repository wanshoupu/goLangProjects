#include "sort_utils.h"

void quicksort_simple(int* arr, int left, int right) {
    if (left >= right)
       return;

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

    quicksort_simple(arr, left, i - 1);
    quicksort_simple(arr, i + 1, right);
}

void quicksort(int* arr, int left, int right) {
    if (left > right)
       return;

    int pivot = arr[left];
    int i = left;
    for (int j = right; i != j;) {
        // note the order: right first, then left
        while (arr[j] >= pivot && i < j)
            j--;
        while (arr[i] <= pivot && i < j)
            i++;
        if (i < j) {
            int t = arr[i];
            arr[i] = arr[j];
            arr[j] = t;
        }
    }
    arr[left] = arr[i];
    arr[i] = pivot;

    quicksort(arr, left, i - 1);
    quicksort(arr, i + 1, right);
}

const char* msg = "Error: array not sorted properly\n";

void unit_test(int* test, int size) {
    printf("Unit test: ");
    print(test, size);
    quicksort_simple(test, 0, size - 1);
    assert(is_sorted(test + 1, test + size), msg);
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