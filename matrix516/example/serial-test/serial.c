/********************************************************************
 *  Copyright (C) Artila Electronics Co., Ltd. All Rights Reserved.
 *
 *	Description: 
 *		Example for serial test.
 *	
 *********************************************************************/
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <termios.h>
#include <errno.h>
#include <fcntl.h>
#include <unistd.h>
// #include <stropts.h>

#include <sys/ioctl.h>
#include <asm/ioctls.h>
#include <linux/ioctl.h>

#include <sys/stat.h>
#include <sys/time.h>
#include <sys/types.h>

#define PROGRAM "serial-test"

void usage()
{
    printf("Usage: " PROGRAM " [OPTION]\n"
           "\n"
           " -h         display this help and exit\n"
           " -p         uart port number\n"
           " -t         test times, default 10 times\n"
    );
    printf("\nExamples:\n"
           "  setuart -p 1                      send 10 times data to port 1\n"
           "  setuart -p 1 -t 100               send 100 times data to port 1\n"
    );
    exit(0);
}

int main(int argc, char *argv[]){
	int fd = -1;
    int port = -1, count = -1;
    int i = 0, opt;
    char buff[512], devicename[32];
    char *messages = "Hello world!!\n";
    //char port[] = "/dev/ttyS1";

    static const char* short_options = "hp:t:";
    
    while((opt = getopt(argc, argv, short_options)) != -1) 
    {
        switch (opt)
        {
            case 'h':
                usage();
                break;
            case 'p':
                port = atoi(optarg);
                break;
            case 't':
                count = atoi(optarg);
                break;
            default:
                usage();
                break;
        }
    }
    
    if( port != -1 )
    {
        sprintf(devicename, "/dev/ttyS%d", port);
        printf("Open %s\n", devicename);
        fd = open(devicename, O_RDWR | O_NOCTTY | O_NDELAY);
        
        if (fd == -1) {
            printf("Unable to open port: %s\n", devicename);
            exit(1);
        }
        
        if( count == -1 ) count = 10;
        
        // for(i=0; i<count; i++)
        while( 1 )
        {
            // send data
            printf("Send[%d]: %s", i+1, messages);
            write(fd, messages, strlen(messages));
            usleep(1000);
            i++;
            
            if( count != 0 && i >= count ) break;            
        }
      	close(fd);
        return 0;
    }
    else 
    {
        usage();
    }
}
