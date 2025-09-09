CC = arm-artila-linux-gnueabi-gcc
STRIP = arm-artila-linux-gnueabi-strip
EXEC = hello
SRC = hello.c

all:	$(EXEC)

$(EXEC): $(OBJS)
	$(CC) -o $(EXEC) $(SRC)
	$(STRIP) $(EXEC)

clean:
	rm -f $(EXEC)
