#include <stddef.h>
void *ft_memcpy(void *dest, const void *src, size_t n)
{
    while (n-- > 0)
    {
        *(unsigned char *)(dest++) = *(unsigned char *)(src++);
    }
    return (void *)(dest - n);
}
