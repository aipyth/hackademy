char *my_strcpy(char *dest, const char *src)
{
    unsigned int i = 0;
    do
    {
        dest[i] = src[i];
    }
    while (src[i++] != '\0');
    return dest;
}
