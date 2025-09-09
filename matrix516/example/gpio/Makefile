CC = arm-artila-linux-gnueabi-gcc
STRIP = arm-artila-linux-gnueabi-strip
SDK = /home/david/m502-workspace/arm-artila-linux-gnueabi_sdk-buildroot
INC = $(SDK)/arm-artila-linux-gnueabi/sysroot/usr/include
LIB = $(SDK)/arm-artila-linux-gnueabi/sysroot/usr/lib

all:	chip_info line_info
chip_info: 	get_chip_info.c
	$(CC) -o chip-info get_chip_info.c -I$(INC) -L$(LIB) -lgpiod
	$(STRIP) chip-info

line_info:  get_line_info.c
	$(CC) -o line-info get_line_info.c -I$(INC) -L$(LIB) -lgpiod
	$(STRIP) line-info

clean:
	rm -f chip-info line-info

