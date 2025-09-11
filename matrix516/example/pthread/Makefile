CC = arm-artila-linux-gnueabi-gcc
STRIP = arm-artila-linux-gnueabi-strip
SDK = /home/david/m502-workspace/arm-artila-linux-gnueabi_sdk-buildroot
INC = $(SDK)/arm-artila-linux-gnueabi/sysroot/usr/include
LIB = $(SDK)/arm-artila-linux-gnueabi/sysroot/usr/lib

all: helloworld
helloworld: helloworld.c
	$(CC) -o $@ helloworld.c -lpthread
	$(STRIP) $@

clean:
	rm -f helloworld

