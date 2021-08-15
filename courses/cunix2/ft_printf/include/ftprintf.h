#include <stddef.h>
#include <stdlib.h>
#include <unistd.h>
#include <stdarg.h>

size_t ft_strlen(char *s);

char *ft_strprepend(char *p, char *s);
char *ft_strprepend_r(char *p, char *s);

char *ft_strappend(char *p, char *s);
char *ft_strappend_r(char *p, char *s);

char *ft_strcut(char *s, size_t cs);
char *ft_strcut_r(char *s, size_t cs);

char *ft_strrepeat(char *s, int n);

char *ft_itoa(int number);
int ft_atoi(const char *nptr);
char *ft_utoa(unsigned int number);

int ft_printf(char *, ...);
int ft_sprintf(char *buff, char *format, ...);
