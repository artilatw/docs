# CC = gcc
CC = arm-artila-linux-gnueabi-gcc
STRIP = arm-artila-linux-gnueabi-strip

all: serial-test

serial-test: serial.c
	${CC} serial.c -o serial-test
	${STRIP} serial-test

clean:
	rm -f serial-test
