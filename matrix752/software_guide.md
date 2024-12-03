# Matrix-752 Software Guide

## Access Ethernet Ports
Matrix-752 come two 10/100Mbit Ethernet ports, the default network settings are shown below:  
|Port Label|Device mapping|IP mode|IP address|
|---|---|---|---|
|LAN1|eth0|DHCP|auto|
|LAN2|eth1|static|192.168.2.127|

## SSH Connection via LAN2
Users need to configure their PC/Notebook's network settings properly to start an SSH connection via **LAN2**.

### Login with root account
The default root password is `root`.
```
$ ssh root@192.168.2.127
root@192.168.1.112's password: 
Last login: Sun Aug  4 15:15:09 2024 from 192.168.2.60

    ___    ____ ____________    ___ 
   /   |  / __ |_  __/  _/ /   /   |
  / /| | / /_/ // /  / // /   / /| |
 / ___ |/ __ _// / _/ // /___/ ___ |
/_/  |_/_/ |_|/_/ /___/_____/_/  |_|

                                    
http://www.artila.com

```
## Basic System Information
### Check Linux Kernel Version
```
root@matrix752:~# uname -a
Linux matrix752 6.1.19-rt8 #1 PREEMPT_RT Mon Mar 13 09:21:32 UTC 2023 armv7l armv7l armv7l GNU/Linux
```
### Check LSB Information
```
root@matrix752:~# lsb_release -a
LSB Version:	core-5.0-noarch:core-5.0-arm
Distributor ID:	poky
Description:	Poky (Yocto Project Reference Distro) 3.3.2
Release:	3.3.2
Codename:	hardknott
```

### Check File System Information 
The Matrix-752 comes with 16GB on-board eMMC Flash memory, which contains boot loader, Linux kernel, root file system and user disk (/home).  

`lsblk` - check block device information
```
root@matrix752:~# lsblk
NAME         MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
mmcblk1      179:0    0 14.6G  0 disk 
|-mmcblk1p1  179:1    0    2G  0 part 
`-mmcblk1p2  179:2    0 12.6G  0 part /
mmcblk1boot0 179:8    0    4M  1 disk 
mmcblk1boot1 179:16   0    4M  1 disk 
```

`df -h` - check file system information
```
Filesystem      Size  Used Avail Use% Mounted on
/dev/root        13G  813M   11G   7% /
devtmpfs        248M     0  248M   0% /dev
tmpfs            40K     0   40K   0% /mnt/.psplash
tmpfs           248M  140K  248M   1% /run
tmpfs           248M  228K  248M   1% /var/volatile
```

`ls -F /` - list all files in root directory

bin/   dev/  gpio/  lib/	 media/  proc/	sbin/	  sys/	usr/
boot/  etc/  home/  lost+found/  mnt/	 run/	swapfile  tmp@	var/

## Configure System Time
### Using `date` Command

`date` - display system date and time  
`date MMDDhhmmYYYY` - set system date and time  

### Write System Time to RTC
`hwclock -w` - write system time to RTC

### Using NTP Serverto Synchronize System Time
Please refer to the following steps to synchronize system time with NTP server.

`apt-get update` - update package list  
`apt-get install ntpdate` - install ntpdate package  
`ntpdate 0.pool.ntp.org` - synchronize system time with NTP server