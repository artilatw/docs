#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>

#define BUZZER "/sys/class/leds/buzzer/brightness"

int main()
{
  int fd;
  char buf[255];
  int delay = 0;
  int count = 200;

  sprintf(buf, BUZZER);

  fd = open(buf, O_WRONLY);

  for(int i=0; i<count; i++)
  {
    write(fd, "1", 1); 
    usleep(delay);
    
    write(fd, "0", 1); 
    usleep(delay);
  }

  close(fd);
  return 0;
}
