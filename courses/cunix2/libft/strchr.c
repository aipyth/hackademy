#include <stddef.h>
char *ft_strchr(const char *s, int c)
{
    unsigned int i = 0;
    while (1)
    {
        if (s[i] == c % 256)
        {
            return (char *)(s + i);
        }
        if (s[i] == '\0')
        {
            break;
        }
        i++;
    }
    return NULL;
}
