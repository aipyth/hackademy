#include "../include/ftprintf.h"

char *ft_strcut(char *s, size_t cs)
{
    size_t slength = ft_strlen(s);
    char *ss;
    if (cs < slength)
    {
        if (!(ss = (char *)malloc(sizeof(char) * (slength - cs + 1))))
        {
            return NULL;
        }
        size_t i = 0;
        while (s[cs] != '\0')
        {
            ss[i++] = s[cs++];
        }
        ss[i] = '\0';
    }
    return ss;
}

char *ft_strcut_r(char *s, size_t n)
{
    char *old_s = s;
    s = ft_strcut(s, n);
    free(old_s);
    return s;
}
