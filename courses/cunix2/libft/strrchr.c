#include <stddef.h>

char *ft_strrchr(const char *s, int c)
{
    unsigned int i = 0;
    char *last_occ = NULL;

    while (1)
    {
        if (s[i] == c % 256)
        {
            last_occ = (char *)(s + i);
        }
        if (s[i] == '\0')
        {
            break;
        }
        i++;
    }
    return last_occ;
}
