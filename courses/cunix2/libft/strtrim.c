#include <stddef.h>
#include <stdlib.h>

int whitespace(char c)
{
    if (c == ' ' ||
        c == '\n' ||
        c == '\t')
    {
        return 1;
    }
    return 0;
}

char *ft_strtrim(char const *s)
{
    size_t slength = 0;
    while (s[slength++] != '\0') {}

    size_t i = 0;
    // count whitespaces at the beginning
    size_t ws_start = 0;
    while (whitespace(s[i++]))
    {
        ws_start++;
    }
    // and at the end
    i = slength == 1 ? 0 : slength - 2;
    size_t ws_end = 0;
    while (whitespace(s[i--]))
    {
        ws_end++;
    }

    size_t start = ws_start;
    size_t end = slength - ws_end - 1;
    size_t length = start < end ? end - start : 0;
    
    char *str;
    if (!(str = (char *)malloc(sizeof(char) * length + 1)))
    {
        return NULL;
    }

    i = start;
    size_t j = 0;
    while (i < end)
    {
        str[j++] = s[i++];
    }
    str[j] = '\0';

    return str;
}
