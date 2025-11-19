#include <stdio.h>
#include <math.h>

/* Проверка делимости без %, / */
int is_divisible(long long n, long long d) {
    if (d == 0) return 0;

    long long t = n;
    while (t > 0) {
        t -= d;
    }
    return t == 0;  // если ушли ровно в 0 — делится
}

/* Проверка числа на простоту без %, / */
int is_prime(long long x) {
    if (x < 2) return 0;
    if (x == 2 || x == 3) return 1;

    // Проверяем делители от 2 до x-1 (оптимизация запрещена без sqrt)
    for (long long d = 2; d < x; d++) {
        if (is_divisible(x, d))
            return 0;
    }
    return 1;
}

int main() {
    long long n;
    printf("Enter number: ");
    scanf("%lld", &n);

    // Проверка отрицательного
    if (n < 0) n = -n;

    // Проверка на 0, 1, 2
    if (n == 0) {
        printf("0 has no prime divisors.\n");
        return 0;
    }
    if (n == 1) {
        printf("1 has no prime divisors.\n");
        return 0;
    }
    if (n == 2) {
        printf("Largest prime divisor: 2\n");
        return 0;
    }

    long long largest = 1;

    // Поиск делителей
    for (long long d = 2; d <= n; d++) {
        if (is_divisible(n, d) && is_prime(d)) {
            largest = d;
        }
    }

    printf("Largest prime divisor: %lld\n", largest);

    return 0;
}
