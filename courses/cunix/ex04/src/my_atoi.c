#include <stdlib.h>
#include <math.h>

int my_atoi(const char *nptr)
{
    unsigned int i = 0;
    int number = 0;

    while (nptr[i] != '\0' || (i == 0 && nptr[i] == '-'))
    {
        // is the symbol dash?
        if (nptr[i] == '-')
        {
            i++;
            continue;
        }
        if (nptr[i] < '0' || nptr[i] > '9')
        {
            break;
        }
        number = number * 10 + nptr[i] - '0';
        i++;
    }
    // whether the first symbol is dash we return negative number or not
    return nptr[0] == '-' ? -number : number;
}

int digits(int num)
{
    unsigned int n = 1;
    while (num >= 10)
    {
        num = num / 10;
        n++;
    }
    return n;
}

const char *my_itoa(int number)
{
    /* char *buff = malloc(sizeof(char)*32); */
    static char buff[32];
    unsigned int i = 0;
   
    if (number == 0)
    {
        buff[i++] = '0';
        buff[i] = '\0';
        return buff;
    }
    if (number < 0)
    {
        buff[i++] = '-';
        number = -number;
    }

    int num_digits = digits(number);
    int p = num_digits;
    int b, value;
    while (p > 0)
    {
        b = (int)pow(10, (p--) - 1);
        value = number / b;
        buff[i++] = '0' + value;
        number -= value * b;
    }
    buff[i] = '\0';
    return buff;
}
