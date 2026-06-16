#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <time.h>

#include "merge_two_sorted_lists.c"

#define RUN_TEST(name, condition)    \
    if (condition)                   \
    {                                \
        printf("[PASS] %s\n", name); \
    }                                \
    else                             \
    {                                \
        printf("[FAIL] %s\n", name); \
    }

struct ListNode *arrayToList(int *nums, int size)
{
    if (size == 0 || nums == NULL)
    {
        return NULL;
    }

    struct ListNode dummy;
    struct ListNode *current = &dummy;

    for (int i = 0; i < size; i++)
    {
        current->next = (struct ListNode *)malloc(sizeof(struct ListNode));
        current->next->val = nums[i];
        current->next->next = NULL;
        current = current->next;
    }
    return dummy.next;
}

bool verifyAndFreeList(struct ListNode *node, int *expected, int expected_size)
{
    bool isEqual = true;
    int count = 0;

    struct ListNode *current = node;
    while (current != NULL)
    {
        if (count >= expected_size || current->val != expected[count])
        {
            isEqual = false;
        }

        struct ListNode *temp = current;
        current = current->next;
        free(temp);
        count++;
    }

    if (count != expected_size)
    {
        isEqual = false;
    }

    return isEqual;
}

typedef struct
{
    const char *name;
    int *list1;
    int size1;
    int *list2;
    int size2;
    int *expected;
    int expected_size;
} TestCase;

void TestMergeTwoLists()
{
    printf("--- RUNNING TESTS ---\n");

    TestCase tests[] = {
        {"Both lists empty",
         NULL, 0,
         NULL, 0,
         NULL, 0},
        {"First list empty",
         NULL, 0,
         (int[]){0}, 1,
         (int[]){0}, 1},
        {"Second list empty",
         (int[]){1, 2, 3}, 3,
         NULL, 0,
         (int[]){1, 2, 3}, 3},
        {"Lists of same size with overlapping values",
         (int[]){1, 2, 4}, 3,
         (int[]){1, 3, 4}, 3,
         (int[]){1, 1, 2, 3, 4, 4}, 6},
        {"Lists of different sizes (L1 larger)",
         (int[]){1, 5, 9, 10}, 4,
         (int[]){2, 3}, 2,
         (int[]){1, 2, 3, 5, 9, 10}, 6},
        {"Lists of different sizes (L2 larger)",
         (int[]){5}, 1,
         (int[]){1, 2, 4, 6}, 4,
         (int[]){1, 2, 4, 5, 6}, 5},
        {"Lists with negative numbers",
         (int[]){-9, -3, 0}, 3,
         (int[]){-5, 2, 8}, 3,
         (int[]){-9, -5, -3, 0, 2, 8}, 6}};

    int num_tests = sizeof(tests) / sizeof(tests[0]);

    for (int i = 0; i < num_tests; i++)
    {
        struct ListNode *l1 = arrayToList(tests[i].list1, tests[i].size1);
        struct ListNode *l2 = arrayToList(tests[i].list2, tests[i].size2);

        struct ListNode *gotList = mergeTwoLists(l1, l2);

        bool passed = verifyAndFreeList(gotList, tests[i].expected, tests[i].expected_size);
        RUN_TEST(tests[i].name, passed);
    }
}

void generateSortedArray(int *arr, int size, int start, int step)
{
    for (int i = 0; i < size; i++)
    {
        arr[i] = start + i * step;
    }
}

void BenchmarkMergeTwoLists()
{
    printf("\n--- RUNNING BENCHMARK ---\n");

    int size = 10000;
    int *s1 = (int *)malloc(size * sizeof(int));
    int *s2 = (int *)malloc(size * sizeof(int));

    generateSortedArray(s1, size, 0, 2);
    generateSortedArray(s2, size, 1, 2);

    int iterations = 1000;

    clock_t start_time = clock();

    for (int i = 0; i < iterations; i++)
    {
        struct ListNode *l1 = arrayToList(s1, size);
        struct ListNode *l2 = arrayToList(s2, size);

        struct ListNode *merged = mergeTwoLists(l1, l2);

        verifyAndFreeList(merged, NULL, 0);
    }

    clock_t end_time = clock();
    double time_taken = ((double)(end_time - start_time)) / CLOCKS_PER_SEC;

    printf("BenchmarkMergeTwoLists\t%d iteracoes\t%.4f segundos totais\n", iterations, time_taken);

    free(s1);
    free(s2);
}

int main()
{
    TestMergeTwoLists();
    BenchmarkMergeTwoLists();

    return 0;
}