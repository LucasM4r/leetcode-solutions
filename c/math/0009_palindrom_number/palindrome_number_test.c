#include <stdio.h>
#include <stdbool.h>
#include "palindrome_number.c"

#define RUN_TEST(name, condition)    \
    if (condition)                   \
    {                                \
        printf("[PASS] %s\n", name); \
    }                                \
    else                             \
    {                                \
        printf("[FAIL] %s\n", name); \
    }

int main()
{
    printf("--- Running C Tests ---\n");

    RUN_TEST("Example 1 (Odd length)", isPalindrome(121) == true);
    RUN_TEST("Example 2 (Negative)", isPalindrome(-121) == false);
    RUN_TEST("Example 3 (Ends with 0)", isPalindrome(10) == false);
    RUN_TEST("Single digit zero", isPalindrome(0) == true);

    return 0;
}