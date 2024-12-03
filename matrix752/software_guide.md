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
root@192.168.2.127's password: 
Last login: Tue Dec  3 07:35:12 2024 from 192.168.2.60

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
```
bin/   dev/  gpio/  lib/	 media/  proc/	sbin/	  sys/	usr/
boot/  etc/  home/  lost+found/  mnt/	 run/	swapfile  tmp@	var/
```

## Configure System Time
### Using `date` Command

`date` - display system date and time  
`date MMDDhhmmYYYY` - set system date and time  

### Write System Time to RTC
`hwclock -w` - write system time to RTC

### Using NTP Serverto Synchronize System Time
Please refer to the following steps to synchronize system time with NTP server.

`apt update` - update package list  
`apt install ntpdate` - install ntpdate package  
`ntpdate 0.pool.ntp.org` - synchronize system time with NTP server

## Insert Kernel Modules
`lsmod` - list all installed kernel modules

To load additional kernel modules during the system boot-up, you can modify the file:
`vi /etc/modules`

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
The Matrix-752 supports storage expansion through its built-in micro SD card slot. When you insert an SD card, the system will detect it automatically. You can use the `lsblk` command to view the device identifier (usually mmcblk0), and then mount it to any directory (e.g. /media) using the `mount` command. The SD card must be properly formatted with a supported filesystem like FAT32 or ext4 before mounting.

### Before SD Insertion
`lsblk` - check block device information
```
NAME         MAJ:MIN RM  SIZE RO TYPE MOUNTPOINTS
mmcblk1      179:0    0 14.6G  0 disk 
├─mmcblk1p1  179:1    0    2G  0 part 
└─mmcblk1p2  179:2    0 12.6G  0 part /
mmcblk1boot0 179:8    0    4M  1 disk 
mmcblk1boot1 179:16   0    4M  1 disk 
```  
### After SD Insertion
`lsblk` - check block device information
```
NAME         MAJ:MIN RM  SIZE RO TYPE MOUNTPOINTS
mmcblk1      179:0    0 14.6G  0 disk
├─mmcblk1p1  179:1    0    2G  0 part
└─mmcblk1p2  179:2    0 12.6G  0 part /
mmcblk1boot0 179:8    0    4M  1 disk
mmcblk1boot1 179:16   0    4M  1 disk
mmcblk0      179:24   0  7.3G  0 disk
└─mmcblk0p1  179:25   0  7.3G  0 part
```  
### Mount mmcblk0 to /media
`mount /dev/mmcblk0p1 /media/` - mount mmcblk0 to /media
```
NAME         MAJ:MIN RM  SIZE RO TYPE MOUNTPOINTS
mmcblk1      179:0    0 14.6G  0 disk
├─mmcblk1p1  179:1    0    2G  0 part
└─mmcblk1p2  179:2    0 12.6G  0 part /
mmcblk1boot0 179:8    0    4M  1 disk
mmcblk1boot1 179:16   0    4M  1 disk
mmcblk0      179:24   0  7.3G  0 disk
└─mmcblk0p1  179:25   0  7.3G  0 part /media
```  
### Unmount /media
`umount /media/` - unmount /media  
