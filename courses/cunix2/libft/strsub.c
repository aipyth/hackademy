#include <stddef.h>
#include <stdlib.h>
char *ft_strsub(char const *s, unsigned int start, size_t len)
{
    size_t slength = 0;
    while (s[slength++] != '\0') {}
    if (slength < start)
    {
        char *str = (char *)malloc(sizeof(char));
        str[0] = '\0';
        return str;
    }

    size_t n = 0;   // length from start to the end of string
    size_t i = start;
    while (s[i++] != '\0')
    {
        n++;
    }

    size_t length = n < len ? n : len;
    char *substr;
    if (!(substr = (char *)malloc(sizeof(char) * length + 1)))
    {
        return NULL;
    }

    i = 0;
    while (i < length)
    {
        if ((substr[i] = s[start + i]) == '\0')
        {
            break;
        }
        i++;
    }
    substr[i] = 0;
    return substr;
}
