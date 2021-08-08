int my_strcmp(const char *str1, const char *str2)
{
    unsigned int i = 0;
    while (str1[i] != '\0' && str2[i] != '\0')
    {
        if (str1[i] == '\0')
        {
            return -1;
        }
        else if (str2[i] == '\0')
        {
            return 1;
        }

        if (str1[i] > str2[i])
        {
            return 1;
        }
        else if (str1[i] < str2[i])
        { 
            return -1;
        }
        i++;
    }
    return 0;
}
