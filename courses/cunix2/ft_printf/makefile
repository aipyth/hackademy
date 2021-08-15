TARGET := libftprintf.a

SRCS := $(wildcard src/*.c)

OBJS := $(SRCS:.c=.o)

CFLAGS := -Iinclude -Wall -Wextra -Werror -g -std=c99

all: $(TARGET)

bonus: $(TARGET)

$(TARGET): $(OBJS)
	ar -crs $@ $(OBJS)

style:
	style50 -o split $(SRCS)

clean:
	rm -rf $(OBJS)

fclean: clean
	rm -rf $(TARGET)

.PHONY: clean fclean
