#include <stddef.h>
#include "libft.h"

char *ft_strnstr(const char *haystack, const char *needle, size_t n)
{
    int i = 0;
    int j = 0;
    int counter;
    while (1)
    {
        if (haystack[i] == needle[0])
        {
            j = 0;
            counter = n;
            while (1)
            {
                if (needle[j] == '\0')
                {
                    return (char *)(haystack + i);
                }
                if (haystack[i + j] != needle[j])
                {
                    break;
                }
                if (haystack[i + j] == '\0')
                {
                    break;
                }
                if (counter == 0)
                {
                    break;
                }
                counter--;
                j++;
            }
        }
        if (haystack[i] == '\0')
        {
            break;
        }
        i++;
    }
    return NULL;
}
