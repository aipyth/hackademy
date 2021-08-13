#include <stddef.h>
#include <stdlib.h>

void *ft_memmove(void *dest, const void *src, size_t n)
{
    unsigned char *arr = (unsigned char *)malloc(n);

    size_t i = 0;
    while (i < n)
    {
        arr[i] = *(unsigned char *)(src + i);
        i++;
    }

    i = 0;
    while (i < n)
    {
        *(unsigned char *)(dest + i) = arr[i];
        i++;
    }

    free(arr);

    return dest;
}
