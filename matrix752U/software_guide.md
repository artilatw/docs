# Matrix-752U Software Guide

## Access Ethernet Ports
Matrix-752U come two 10/100Mbit Ethernet ports, the default network settings are shown below:  
|Port Label|Device mapping|IP mode|IP address|
|---|---|---|---|
|LAN1|eth0|DHCP|auto|
|LAN2|eth1|static|192.168.2.127|

## SSH Connection via LAN2
Users need to configure their PC/Notebook's network settings properly to start an SSH connection via **LAN2**.

### Login with guest account
By default, the Matrix-752U is not allowed to login as a root account via SSH. You can login as a guest account. The default guest user name is guest and the password is guest.
```
$ ssh guest@192.168.2.127
guest@192.168.2.127's password: 
Welcome to Ubuntu 22.04.2 LTS (GNU/Linux 6.1.19-rt8 armv7l)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

This system has been minimized by removing packages and content that are
not required on a system that users do not log into.

To restore this content, you can run the 'unminimize' command.
Last login: Fri Nov 29 11:29:37 2024 from 192.168.1.141
guest@matrix752:~$
```

### Switch to root account
The root account user name is **root** and the passwword is **root**.

```
guest@matrix752:~$ su -
Password: 
root@matrix752:~#
```
## Basic System Information
### Check Linux Kernel Version
```
root@matrix752:~# uname -a
Linux matrix752 6.1.19-rt8 #1 PREEMPT_RT Mon Mar 13 09:21:32 UTC 2023 armv7l armv7l armv7l GNU/Linux
```

### Check Ubuntu Version
```
root@matrix752:~# lsb_release -a
No LSB modules are available.
Distributor ID:	Ubuntu
Description:	Ubuntu 22.04.2 LTS
Release:	22.04
Codename:	jammy
```  
### Check File System Information 
The Matrix-752U comes with 16GB on-board eMMC Flash memory, which contains boot loader, Linux kernel, root file system and user disk (/home).  
```
root@matrix752:~# lsblk
NAME         MAJ:MIN RM  SIZE RO TYPE MOUNTPOINTS
mmcblk1      179:0    0 14.6G  0 disk
tqmmcblk1p1  179:1    0    2G  0 part
mqmmcblk1p2  179:2    0 12.6G  0 part /
mmcblk1boot0 179:8    0    4M  1 disk
mmcblk1boot1 179:16   0    4M  1 disk

root@matrix752:~# df -h
Filesystem      Size  Used Avail Use% Mounted on
/dev/root        13G  778M   11G   7% /
tmpfs           248M     0  248M   0% /dev/shm
tmpfs            99M  3.3M   96M   4% /run
tmpfs           5.0M     0  5.0M   0% /run/lock
tmpfs            50M     0   50M   0% /run/user/1000

root@matrix752:~# ls -F /
bin@   dev/  gpio/  lib@         media/  opt/   root/  sbin@  swapfile  tmp/  var/
boot/  etc/  home/  lost+found/  mnt/    proc/  run/   srv/   sys/      usr/

```
## Configure System Time
### Auto Synchronize with NTP Server
Matrix-752U supports `timedatectl` command to manage the Linux system time. By Default, the system time is synchronized with an NTP server.  

```
root@pac6070:~# timedatectl
               Local time: Thu 2024-07-18 13:47:56 CST
           Universal time: Thu 2024-07-18 05:47:56 UTC
                 RTC time: Thu 2024-07-18 05:47:56
                Time zone: Asia/Taipei (CST, +0800)
System clock synchronized: yes
              NTP service: active
          RTC in local TZ: no
```  
### Manually Set the System Time
If you want to manually set the system time, please follow the steps shown below: 
```
root@pac6070:~# timedatectl set-ntp no
root@pac6070:~# timedatectl set-time "2024-07-18 14:00:00"
root@pac6070:~# timedatectl
               Local time: Thu 2024-07-18 14:00:13 CST
           Universal time: Thu 2024-07-18 06:00:13 UTC
                 RTC time: Thu 2024-07-18 06:00:13
                Time zone: Asia/Taipei (CST, +0800)
System clock synchronized: no
              NTP service: inactive
          RTC in local TZ: no
```

## Access Digital I/O  
### DIO Mapping
The Matrix-752U comes with 2x opto-isolated digital inputs and 2x relay digital outputs. Below is the DIO mapping table:

|DI Number|Device Mapping|DO Number|Device Mapping|
|---|---|--|---|
|DI1|/gpio/DI1|DO1|/gpio/DO1|
|DI2|/gpio/DI2|DO2|/gpio/DO2|
```
root@matrix752:~# ls /gpio
DI1  DI2  DO1  DO2
```

### Read Digital Input
The Matrix-752U provides 2x opto-isolated digital inputs. The following command will read the input status (0 or 1) of DI1 and DI2.

`cat /gpio/DI1/value`  - read value of DI1  
`cat /gpio/DI2/value`  - read value of DI2  

### Write Digital Output
The Matrix-752U provides 2x solid-state relays, which are normally open. The following command will let the relay close.
  
`echo 1 > /gpio/DO1/value` - close DO1 relay  
`echo 0 > /gpio/DO1/value` - open DO1 relay  
`echo 1 > /gpio/DO2/value` - close DO2 relay  
`echo 0 > /gpio/DO2/value` - open DO2 relay  


## Software Package Management
The Matrix-752U uses Ubuntu's APT (Advanced Package Tool) for software package management. Here are the common package management commands:

### Basic Package Operations
- `apt install <package>` - Install a new package
- `apt remove <package>` - Remove an installed package 
- `apt purge <package>` - Remove package and its configuration files
- `apt autoremove` - Remove automatically installed dependencies that are no longer needed

### Package Information
- `apt search <keyword>` - Search for packages matching keyword
- `apt show <package>` - Show detailed package information
- `apt list --installed` - List all installed packages

### System Updates
- `apt update` - Update package index from repositories
- `apt upgrade` - Upgrade all upgradeable packages
- `apt full-upgrade` - Upgrade packages with auto-handling of dependencies
- `apt dist-upgrade` - Smart upgrade that handles changing dependencies

Note: Most apt commands require root privileges. Use sudo if not logged in as root.

## Using SD Card
The Matrix-752U supports storage expansion through its built-in micro SD card slot. When you insert an SD card, the system will detect it automatically. You can use the `lsblk` command to view the device identifier (usually mmcblk0), and then mount it to any directory (e.g. /media) using the `mount` command. The SD card must be properly formatted with a supported filesystem like FAT32 or ext4 before mounting.

Before SD Insertion
```
root@matrix752:~# lsblk
NAME         MAJ:MIN RM  SIZE RO TYPE MOUNTPOINTS
mmcblk1      179:0    0 14.6G  0 disk 
├─mmcblk1p1  179:1    0    2G  0 part 
└─mmcblk1p2  179:2    0 12.6G  0 part /
mmcblk1boot0 179:8    0    4M  1 disk 
mmcblk1boot1 179:16   0    4M  1 disk 
```  
After SD Insertion
```
root@matrix752:~# lsblk
NAME         MAJ:MIN RM  SIZE RO TYPE MOUNTPOINTS
mmcblk1      179:0    0 14.6G  0 disk
├─mmcblk1p1  179:1    0    2G  0 part
└─mmcblk1p2  179:2    0 12.6G  0 part /
mmcblk1boot0 179:8    0    4M  1 disk
mmcblk1boot1 179:16   0    4M  1 disk
mmcblk0      179:24   0  7.3G  0 disk
└─mmcblk0p1  179:25   0  7.3G  0 part
```  
Mount mmcblk0 to /media.  
```
root@matrix752:~# mount /dev/mmcblk0p1 /media/
root@matrix752:~# lsblk
NAME         MAJ:MIN RM  SIZE RO TYPE MOUNTPOINTS
mmcblk1      179:0    0 14.6G  0 disk
├─mmcblk1p1  179:1    0    2G  0 part
└─mmcblk1p2  179:2    0 12.6G  0 part /
mmcblk1boot0 179:8    0    4M  1 disk
mmcblk1boot1 179:16   0    4M  1 disk
mmcblk0      179:24   0  7.3G  0 disk
└─mmcblk0p1  179:25   0  7.3G  0 part /media
```  
Unmount /media
```
root@matrix752:~# umount /media/
```  

## RS-485/232 Serial Port Settings
The Matrix-752U comes with two serial ports P1 and P2. P1 can be configured as RS-485 (default) or RS-232. P2 is always RS-232. 

|Serial port|Interface|Mapping|
|---|---|---|
|P1|RS-485|/dev/ttymxc1|
|P2|RS-232|/dev/ttymxc2|

If you need to modify the serial port settings, use the **setuart** command.
```
root@matrix752:~# setuart -h
Artila utility: setuart
Usage: setuart [OPTION]

 -h         display this help and exit
 -v         print version number and exit
 -p         uart port number
 -t         uart interface type [232, 485]
 -b         set baud rate, up to 921600bps

Examples:
  setuart -p 1                      display port 1 type and baud rate
  setuart -p 1 -t 485 -b 115200     set port 1 type RS-485 and baud rate to 115200
  setuart -p 1 -t 232 -b 9600       set port 1 type to RS-232 and baud rate to 9600
```
***Caution***  
The serial port’s mode and associated communication parameters will go back to factory default after system reboot.  

### Default Serial Port Settings
The default serial port settings are shown below:
```
root@matrix752:~# setuart -p 1
Port 1 ==> type: RS485
baud: 9600
root@matrix752:~# setuart -p 2
Port 2 ==> type: RS232
baud: 9600
```

## Network Interface Settings
The Matrix-752U comes with two 10/100Mbit Ethernet ports, the default network settings are shown below:  
|Port Label|Device mapping|IP mode|IP address|
|---|---|---|---|
|LAN1|eth0|DHCP|auto|
|LAN2|eth1|static|192.168.2.127|

### Display the Network Settings
`ip a show eth0` - display the network settings of eth0  
`ip a show eth1` - display the network settings of eth1  

### Modify the Network Settings
To modify network settings according to your LAN environment, follow these steps:

1. `vi /etc/network/interfaces` - modify the network interface configuration file
2. `systemctl restart NetworkManager`   # Restart the network manager service
3. `ifconfig eth0 down`                 # Shutdown the network interface
4. `ifconfig eth0 up`                   # Start up the network interface

Note: Replace eth0 with the appropriate interface name (eth0 or eth1) that you want to configure.
```
root@pac6070:~# systemctl restart NetworkManager
root@pac6070:~# ifconfig eth0 down && ifconfig eth0 up
root@pac6070:~# ip a show eth0
```


