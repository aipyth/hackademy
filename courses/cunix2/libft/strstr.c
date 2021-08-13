#include <stddef.h>
char *ft_strstr(const char *haystack, const char *needle)
{
    int i = 0;
    int j = 0;
    while (1)
    {
        if (haystack[i] == needle[0])
        {
            j = 0;
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
