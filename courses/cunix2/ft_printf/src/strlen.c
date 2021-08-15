#include "../include/ftprintf.h"

size_t ft_strlen(char *s)
{
    size_t length = 0;
    while (s[length++] != '\0') {}
    return length - 1;
}
