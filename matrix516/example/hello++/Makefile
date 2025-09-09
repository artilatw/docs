CC = arm-artila-linux-gnueabi-g++
STRIP = arm-artila-linux-gnueabi-strip
EXEC = hello++
SRC = hello.cpp
CFLAGS = -Wno-deprecated

all:	$(EXEC)

$(EXEC): $(OBJS)
	$(CC) $(SRC) -o $(EXEC)
	$(STRIP) $(EXEC)

clean:
	rm -f $(EXEC)
