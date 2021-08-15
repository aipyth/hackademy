#include "../include/ftprintf.h"

char *ft_strrepeat(char *s, int n)
{
    size_t slength = ft_strlen(s);
    char *ss = (char *)malloc(sizeof(char) * (slength * n + 1));
    size_t i = 0;
    size_t j = 0;
    while (n-- > 0)
    {
        j = 0;
        while (s[j] != '\0')
        {
            ss[i++] = s[j++];
        }
    }
    ss[i] = '\0';
    return ss;
}
