#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <time.h>

#include "invert_binary_tree.c"

#define RUN_TEST(name, condition)    \
    if (condition)                   \
    {                                \
        printf("[PASS] %s\n", name); \
    }                                \
    else                             \
    {                                \
        printf("[FAIL] %s\n", name); \
    }

struct TreeNode *createNode(int val)
{
    struct TreeNode *node = (struct TreeNode *)malloc(sizeof(struct TreeNode));
    node->val = val;
    node->left = NULL;
    node->right = NULL;
    return node;
}

void freeTree(struct TreeNode *root)
{
    if (root == NULL)
        return;
    freeTree(root->left);
    freeTree(root->right);
    free(root);
}

bool compareTrees(struct TreeNode *t1, struct TreeNode *t2)
{
    if (t1 == NULL && t2 == NULL)
        return true;
    if (t1 == NULL || t2 == NULL)
        return false;
    if (t1->val != t2->val)
        return false;

    return compareTrees(t1->left, t2->left) && compareTrees(t1->right, t2->right);
}

void TestInvertTree()
{
    printf("--- RUNNING TESTS ---\n");

    struct TreeNode *t1 = NULL;
    struct TreeNode *expected1 = NULL;
    struct TreeNode *got1 = invertTree(t1);
    RUN_TEST("Empty tree", compareTrees(got1, expected1));
    freeTree(got1);
    freeTree(expected1);

    struct TreeNode *t2 = createNode(1);
    struct TreeNode *expected2 = createNode(1);
    struct TreeNode *got2 = invertTree(t2);
    RUN_TEST("Single node", compareTrees(got2, expected2));
    freeTree(got2);
    freeTree(expected2);

    struct TreeNode *t3 = createNode(1);
    t3->left = createNode(2);

    struct TreeNode *expected_t3 = createNode(1);
    expected_t3->right = createNode(2);

    struct TreeNode *got3 = invertTree(t3);
    RUN_TEST("Left-heavy tree", compareTrees(got3, expected_t3));
    freeTree(got3);
    freeTree(expected_t3);

    struct TreeNode *t4 = createNode(4);
    t4->left = createNode(2);
    t4->right = createNode(7);
    t4->left->left = createNode(1);
    t4->left->right = createNode(3);
    t4->right->left = createNode(6);
    t4->right->right = createNode(9);

    struct TreeNode *expected_t4 = createNode(4);
    expected_t4->left = createNode(7);
    expected_t4->right = createNode(2);
    expected_t4->left->left = createNode(9);
    expected_t4->left->right = createNode(6);
    expected_t4->right->left = createNode(3);
    expected_t4->right->right = createNode(1);

    struct TreeNode *got4 = invertTree(t4);
    RUN_TEST("Full binary tree", compareTrees(got4, expected_t4));
    freeTree(got4);
    freeTree(expected_t4);
}

struct TreeNode *createBalancedTree(int depth)
{
    if (depth == 0)
        return NULL;
    struct TreeNode *node = createNode(depth);
    node->left = createBalancedTree(depth - 1);
    node->right = createBalancedTree(depth - 1);
    return node;
}

void BenchmarkInvertTree()
{
    printf("\n--- RUNNING BENCHMARK ---\n");

    int iterations = 10000;
    int depth = 10;

    clock_t start_time = clock();

    for (int i = 0; i < iterations; i++)
    {
        struct TreeNode *root = createBalancedTree(depth);
        struct TreeNode *inverted = invertTree(root);

        freeTree(inverted);
    }

    clock_t end_time = clock();
    double time_taken = ((double)(end_time - start_time)) / CLOCKS_PER_SEC;

    printf("BenchmarkInvertTree\t%d iteracoes\t%.4f segundos totais\n", iterations, time_taken);
}

int main()
{
    TestInvertTree();
    BenchmarkInvertTree();

    return 0;
}