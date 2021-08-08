#include <unistd.h>

int my_puts(const char *s)
{
    // get s length
    int n = 0;
    while (s[n++] != '\0') {}
    int ret = write(1, s, n - 1);
    write(1, "\n", 1);
    return ret;
}
