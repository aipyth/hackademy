TARGET := libft.a

SRCS := $(wildcard *.c)

OBJS := $(SRCS:.c=.o)

CFLAGS := -Iinclude -Wall -Wextra -Werror -g -std=c99

all: $(TARGET)
	echo $(SRCS)
	echo $(OBJS)

$(TARGET): $(OBJS)
	ar -crs $@ $(OBJS)

style:
	style50 -o split $(SRCS)

clean:
	rm -rf $(OBJS)

fclean: clean
	rm -rf $(TARGET)

.PHONY: clean fclean
