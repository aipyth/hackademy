#include <stddef.h>
#include <stdlib.h>
char *ft_strjoin(char const *s1, char const *s2)
{
    size_t l1 = 0;
    size_t l2 = 0;
    while (s1[l1++] != '\0') {}
    while (s2[l2++] != '\0') {}

    char *s = (char *)malloc(sizeof(char) * (l1 + l2 - 1));
    
    size_t i = 0;
    size_t j = 0;
    while (j < l1 - 1)
    {
        s[i++] = s1[j++];
    }
    j = 0;
    while (j < l2 - 1)
    {
        s[i++] = s2[j++];
    }
    s[i] = 0;
    return s;
}
