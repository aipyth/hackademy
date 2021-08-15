#include "../include/ftprintf.h"

char *ft_strappend(char *p, char *s)
{
    size_t slength = ft_strlen(s);
    size_t plength = ft_strlen(p);

    char *ss = (char *)malloc(sizeof(char) * (slength + plength + 1));
    size_t i = 0;
    while (*s != '\0')
    {
        ss[i++] = *(s++);
    }
    while (*p != '\0')
    {
        ss[i++] = *(p++);
    }
    ss[i] = '\0';

    return ss;
}

char *ft_strappend_r(char *p, char *s)
{
    char *old_s = s;
    s = ft_strappend(p ,s);
    free(old_s);
    return s;
}
