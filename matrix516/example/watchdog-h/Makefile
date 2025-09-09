CC = arm-artila-linux-gnueabi-gcc
STRIP = arm-artila-linux-gnueabi-strip
TARGET = watchdog.o


all: watchdog

watchdog: $(TARGET)
	$(CC) -o $@ $(TARGET)
	$(STRIP) $@
	
clean:
	rm -rf *~ $(TARGET) watchdog