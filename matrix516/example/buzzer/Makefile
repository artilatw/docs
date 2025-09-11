CC = arm-artila-linux-gnueabi-gcc
STRIP = arm-artila-linux-gnueabi-strip

all: beep

beep: beep.c
	$(CC) -o beep beep.c
	$(STRIP) beep    
    
clean:
	rm -f beep
