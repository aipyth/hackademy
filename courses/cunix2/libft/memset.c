#include <stddef.h>

void *ft_memset(void *s, int c, size_t n)
{
    while (n-- > 0)
    {
        *(unsigned char *)(s++) = c;
    }
    return s - n;
}
