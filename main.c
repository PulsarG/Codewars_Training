#include <stdio.h>

int main(void)
{
    int value;
    int result;

    while (1)
    {
        printf("Enter number: ");

        result = scanf("%d", &value);

        if (result == 1)
        {
            printf("OK: %d\n", value);
            break;
        }
        else
        {
            printf("Error: not number!\n");

            int c;
            while ((c = getchar()) != 10 && c != EOF);
        }
    }

    return 0;
}
