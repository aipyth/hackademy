#include "../include/ftprintf.h"

char *ft_strdup(const char *s)
{
    unsigned int length = 0;
    while (s[length++] != '\0') {}

    char *ret = (char *)malloc(sizeof(char) * length);

    unsigned int i = -1;
    while (++i < length)
    {
        ret[i] = s[i];
    }
    return ret;
}
