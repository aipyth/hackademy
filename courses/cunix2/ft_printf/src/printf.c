#include "../include/ftprintf.h"

typedef struct
{
    char was_num;
    char num_positive;
    char sign_num;
    char space;
    char inversed_pagination;
    char pagination_placeholder;
    size_t pagination_width;
} context;

int is_format_char(char c)
{
    if (c == 'c' ||
        c == 'd' ||
        c == 'i' ||
        c == 's')
    {
        return 1;
    }
    return 0;
}

char *get_flag_content(context *ctx, char **format, va_list *args)
{
    int d;
    unsigned int u;
    char *content;
    switch (**format)
    {

        case 's':
            content = va_arg(*args, char *);
            if (content == NULL)
            {
                content = (char *)malloc(sizeof(char) * 7);
                content = "(null)";
            }
            break;
        case 'i':
        case 'd':
            d = va_arg(*args, int);
            content = ft_itoa(d);
            ctx->was_num = 1;
            ctx->num_positive = d >= 0;
            if (ctx->space && d >= 0)
            { 
                content = ft_strprepend_r(" ", content);
            }
            if (ctx->sign_num && d >= 0)
            {
                content = ft_strprepend_r("+", content);
            }
            break;
        case 'u':
            u = va_arg(*args, unsigned int);
            content = ft_utoa(u);
            ctx->was_num = 1;
            ctx->num_positive = 1;
            if (ctx->space)
            { 
                content = ft_strprepend_r(" ", content);
            }
            if (ctx->sign_num)
            {
                content = ft_strprepend_r("+", content);
            }
            break;
        case 'c':
            char c = (char)va_arg(*args, int);
            content = (char *)malloc(sizeof(char) * 2);
            content[0] = c;
            content[1] = '\0';
            break;
        case '%':
            content = (char *)malloc(sizeof(char) * 2);
            content[0] = '%';
            content[1] = '\0';
            break;
    }
    return content;
}

char *paginate_number(context *ctx, char *content)
{
    size_t content_length = ft_strlen(content);

    char *(*pagger)(char *, char *) = ctx->inversed_pagination ? &ft_strappend_r : ft_strprepend_r;

    if (ctx->pagination_placeholder == '0')
    {
        if (content[0] == ' ')
        {
            content = ft_strcut_r(content, 1);
        }
        if (ctx->sign_num || content[0] == '-')
        {
            content = ft_strcut_r(content, 1);

            char *to_paste = ft_strrepeat("0", ctx->pagination_width - content_length);
            content = pagger(to_paste, content);
            free(to_paste);

            char *sign = ctx->num_positive ? "+" : "-";
            content = ft_strprepend_r(sign, content);
        }
        else
        {
            char *to_paste = ft_strrepeat("0", ctx->pagination_width - content_length);
            content = pagger(to_paste, content);
            free(to_paste);
        }
        if (ctx->space)
        {
            content = ft_strprepend_r(" ", content);
        }
    }
    else
    {
        char *to_paste = ft_strrepeat(" ", ctx->pagination_width - content_length);
        content = pagger(to_paste, content);
        free(to_paste);
    }
    return content;
}

char *paginate_string(context *ctx, char *content)
{
    char *(*pagger)(char *, char *) = ctx->inversed_pagination ? &ft_strappend : ft_strprepend;
    size_t content_length = ft_strlen(content);
    char *to_paste = ft_strrepeat(" ", ctx->pagination_width - content_length);
    content = pagger(to_paste, content);
    free(to_paste);
    return content;
}

char *process_complex_format(char **format, va_list *args)
{
    context *ctx = (context *)malloc(sizeof(context *));
    ctx->was_num = 0;
    ctx->num_positive = 0;
    ctx->sign_num = 0;
    ctx->inversed_pagination = 0;
    ctx->pagination_placeholder = ' ';
    ctx->pagination_width = 0;

    if (**format == ' ')
    {
        ctx->space = 1;
        (*format)++;
    }
    if (**format == '+')
    {
        ctx->sign_num = 1;
        (*format)++;
    }
    if (**format == '-')
    {
        ctx->inversed_pagination = 1;
        (*format)++;
    }
    if (**format == '0')
    {
        ctx->pagination_placeholder = '0';
        (*format)++;
    }
    if (**format >= '1' && **format <='9')
    {
        ctx->pagination_width = ft_atoi(*format);
        while (**format >= '0' && **format <= '9')
        {
            (*format)++;
        }
    }

    char *content = get_flag_content(ctx, format, args);
    size_t content_length = ft_strlen(content);

    if (ctx->pagination_width > content_length)
    {
        if (ctx->was_num)
        {
            content = paginate_number(ctx, content);
        }
        else if (content[0] != '%')
        {
            content = paginate_string(ctx, content);
        }
    }
    
    return  content;
}

int ft_printf(char *format, ...)
{
    va_list args;
    size_t written = 0;

    char *content = (char *)malloc(sizeof(char));
    content[0] = '\0';
    char c[] = {0, 0};

    va_start(args, format);
    while (*format)
    {
        switch (*format)
        {
            case '%':
                format++;
                content = ft_strappend_r(process_complex_format(&format, &args), content);
                break;
            default:
                c[0] = *format;
                content = ft_strappend_r(c, content);
                break;
        }
        format++;
    }

    va_end(args);
    written = write(1, content, ft_strlen(content));
    free(content);
    return written;
}

int ft_sprintf(char *buff, char *format, ...)
{
    va_list args;
    size_t written = 0;

    char *content = (char *)malloc(sizeof(char));
    content[0] = '\0';
    char c[] = {0, 0};

    va_start(args, format);
    while (*format)
    {
        switch (*format)
        {
            case '%':
                format++;
                content = ft_strappend_r(process_complex_format(&format, &args), content);
                break;
            default:
                c[0] = *format;
                content = ft_strappend_r(c, content);
                break;
        }
        format++;
    }

    va_end(args);

    size_t i = 0;
    while ((buff[i++] = content[i]) != '\0') {}
    
    free(content);
    return written;

}

