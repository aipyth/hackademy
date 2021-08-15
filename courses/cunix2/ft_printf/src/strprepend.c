#include "../include/ftprintf.h"

char *ft_strprepend(char *p, char *s)
{
    size_t slength = ft_strlen(s);
    size_t plength = ft_strlen(p);

    char *ss = (char *)malloc(sizeof(char) * (slength + plength + 1));
    size_t i = 0;
    while (*p != '\0')
    {
        ss[i++] = *(p++);
    }
    while (*s != '\0')
    {
        ss[i++] = *(s++);
    }
    if (s[0] == '\0' && s[1] == '\0')
    {
        ss[i++] = '\0';
    }
    ss[i] = '\0';

    return ss;
}

char *ft_strprepend_r(char *p, char *s)
{
    char *old_s = s;
    s = ft_strprepend(p ,s);
    free(old_s);
    return s;
}
