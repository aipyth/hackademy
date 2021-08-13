#include <stdlib.h>
#include <stdio.h>
char *ft_strmap(char const *s, char (*f)(char))
{
    size_t length = 0;
    while (s[length++] != '\0') {}

    char *r = (char *)malloc(sizeof(char) * length);

    size_t i = 0;
    while (i < length - 1)
    {
        r[i] = f(s[i]);
        i++;
    }
    r[i] = 0;

    return r;
}
