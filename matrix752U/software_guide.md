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
root@pac6070:~# uname -a
Linux pac6070 6.6.22 #1 Fri Mar 15 18:25:07 UTC 2024 armv7l armv7l armv7l GNU/Linux
```

### Check Ubuntu Version
```
root@pac6070:~# lsb_release -a
No LSB modules are available.
Distributor ID: Ubuntu
Description:    Ubuntu 22.04.4 LTS
Release:        22.04
Codename:       jammy
```  
### Check File System Information 
The Matrix-752U comes with 16GB on-board eMMC Flash memory, which contains boot loader, Linux kernel, root file system and user disk (/home).  
```
root@pac6070:~# lsblk
NAME         MAJ:MIN RM  SIZE RO TYPE MOUNTPOINTS
mmcblk1      179:0    0 14.6G  0 disk
tqmmcblk1p1  179:1    0    2G  0 part
mqmmcblk1p2  179:2    0 12.6G  0 part /
mmcblk1boot0 179:8    0    4M  1 disk
mmcblk1boot1 179:16   0    4M  1 disk

root@pac6070:~# df -h
Filesystem      Size  Used Avail Use% Mounted on
/dev/root        13G  1.1G   11G   9% /
tmpfs           502M     0  502M   0% /dev/shm
tmpfs           201M   21M  180M  11% /run
tmpfs           5.0M     0  5.0M   0% /run/lock
tmpfs           101M     0  101M   0% /run/user/0

root@pac6070:~# ls -F /
bin@   etc/   lib@         mnt/   root/  srv/      tmp/
boot/  gpio/  lost+found/  opt/   run/   swapfile  usr/
dev/   home/  media/       proc/  sbin@  sys/      var/
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
The PAC-6070 comes with 8x opto-isolated digital inputs and 8x relay digital outputs. Below is the DIO mapping table:

|DI Number|Device Mapping|DO Number|Device Mapping|
|---|---|--|---|
|DI1|/gpio/DI1|DO1|/gpio/DO1|
|DI2|/gpio/DI2|DO2|/gpio/DO2|
|DI3|/gpio/DI3|DO3|/gpio/DO3|
|DI4|/gpio/DI4|DO4|/gpio/DO4|
|DI5|/gpio/DI5|DO5|/gpio/DO5|
|DI6|/gpio/DI6|DO6|/gpio/DO6|
|DI7|/gpio/DI7|DO7|/gpio/DO7|
|DI8|/gpio/DI8|DO8|/gpio/DO8|

```
root@pac6070:~# ls /gpio/
DI1  DI3  DI5  DI7  DO1  DO3  DO5  DO7  pciepower
DI2  DI4  DI6  DI8  DO2  DO4  DO6  DO8
```

### Read Digital Input
Example: read value of DI1  
```
root@pac6070:~# cat /gpio/DI1/value  
```
### Write Digital Output
Example: 
If DO1 relay is at NO (normally open) mode, the following command will let the relay close.
```  
root@pac6070:~# echo 1 > /gpio/DO1/value
```
Example: If DO1 relay is at NO (normally open) mode, the following command will let the relay open.
```  
root@pac6070:~# echo 0 > /gpio/DO1/value
```
## Access Analog Input
The PAC-6070 can measure voltage input (-10Vdc ~ +10Vdc) or current input (0mA ~ 20mA). 



### Voltage Input Wiring
Users can find voltage input terminals labeled as V1+, V1-, V2+, V2-, V3+, V3- and AGND. 

### Voltage Input Mode Setting
The voltage input supports differential mode or single-end mode.
User can modify **/etc/modprobe.d/ad4111.conf** file to set the voltage input mode.
```
options ad4111 differential=1
or
options ad4111 differential=0
```
differential=1 means differential mode.
differential=0 means single-end mode.

If the voltage input mode is changed, please reboot the system to activate the new settings.
```
root@pac6070:~# reboot
```

### Differential Input Channels
For differential voltage input, the first channel is V1+ and V1-, the second channel is V2+ and V2-, the third channel is V3+ and V3-. 
|Differential Channel|Terminals|
|---|---|
|CH1|V1+ and V1-|
|CH2|V2+ and V2-| 
|CH3|V3+ and V3-|

### Differential Input Calculation
The file paths of the voltage inputs are located in /sys/bus/iio/devices/iio:device0/

|Channel|Raw Value Path|Offset Value Path|
|---|---|---|
|CH1|in_voltage0_voltage1_raw|in_voltage0_voltage1_offset|
|CH2|in_voltage2_voltage3_raw|in_voltage2_voltage3_offset|
|CH3|in_voltage4_voltage5_raw|in_voltage4_voltage5_offset|

In the same directory, there is an **in_voltage_voltage_scale** file stores the scale value which needs to multiply with raw data to get the voltage value.

The common formula to calculate the voltage is:
```
voltage = (raw * scale) + offset
```

Below are the formulas to calculate the voltage value of each channel:
```
CH1_voltage = (in_voltage0_voltage1_raw * in_voltage_voltage_scale) + in_voltage0_voltage1_offset

CH2_voltage = (in_voltage2_voltage3_raw * in_voltage_voltage_scale) + in_voltage2_voltage3_offset

CH3_voltage = (in_voltage4_voltage5_raw * in_voltage_voltage_scale) + in_voltage4_voltage5_offset
```

### Single-end Input Channels
For single-end voltage input, the first channel is V1+ and AGND, the second channel is V1- and AGND, and so on, please refer to the following table:

|Single-end Channel|Terminals|
|---|---|
|CH1|V1+ and AGND|
|CH2|V1- and AGND|
|CH3|V2+ and AGND|
|CH4|V2- and AGND|
|CH5|V3+ and AGND|
|CH6|V3- and AGND|

### Single-end Input Calculation
The file paths of the voltage inputs are located in /sys/bus/iio/devices/iio:device0/

|Channel|Raw Value Path|Offset Value Path|
|---|---|---|
|CH1|in_voltage0_raw|in_voltage0_offset|
|CH2|in_voltage1_raw|in_voltage1_offset|
|CH3|in_voltage2_raw|in_voltage2_offset|
|CH4|in_voltage3_raw|in_voltage3_offset|
|CH5|in_voltage4_raw|in_voltage4_offset|
|CH6|in_voltage5_raw|in_voltage5_offset|

In the same directory, there is an **in_voltage_voltage_scale** file stores the scale value which needs to multiply with raw data to get the voltage value.

The common formula to calculate the voltage is:
```
voltage = (raw * scale) + offset
```

Below are the formulas to calculate the voltage value of each channel:
```
CH1_voltage = (in_voltage0_raw * in_voltage_voltage_scale) + in_voltage0_offset

CH2_voltage = (in_voltage1_raw * in_voltage_voltage_scale) + in_voltage1_offset

CH3_voltage = (in_voltage2_raw * in_voltage_voltage_scale) + in_voltage2_offset

CH4_voltage = (in_voltage3_raw * in_voltage_voltage_scale) + in_voltage3_offset

CH5_voltage = (in_voltage4_raw * in_voltage_voltage_scale) + in_voltage4_offset

CH6_voltage = (in_voltage5_raw * in_voltage_voltage_scale) + in_voltage5_offset
```

### Current Input Wiring
Users can find current input terminals labeled as AI1, AI2 and AGND.

### Current Input Channels
|Current Channel|Terminals|
|---|---|
|CH1|AI1 and AGND|
|CH2|AI2 and AGND|

### Current Input Calculation
The file paths of the current inputs are located in /sys/bus/iio/devices/iio:device0/

|Channel|Raw Value Path|Offset Value Path|
|---|---|---|
|CH1|in_current0_raw|in_current0_offset|
|CH2|in_current1_raw|in_current1_offset|

In the same directory, there is an **in_current_current_scale** file stores the scale value which needs to multiply with raw data to get the current value.

The common formula to calculate the current is:
```
current = (raw * scale) + offset
```

Below are the formulas to calculate the current value of each channel:
```
CH1_current = (in_current0_raw * in_current_current_scale) + in_current0_offset

CH2_current = (in_current1_raw * in_current_current_scale) + in_current1_offset
```

### Analog Input Utility
The PAC-6070 provides a command line utility `lsadc` which displays real-time analog input readings in both differential and single-end modes. Below are example screenshots showing the output format and measurement values:

```
root@pac6070:~# lsadc
```

- Differential mode
<img src="img/cli_diff.png" width=600> 

- Single-end mode       
<img src="img/cli_single.png" width=600> 

## Built-in Web-based Tool
The PAC-6070 comes with a web-based tool to monitor the voltage and current inputs. Besides, users can also monitor all the digital I/O status and control the digital outputs.

### Access the Web-based Tool
The web-based tool is hosted on the PAC-6070 at http://192.168.2.127 (please adjust the IP address according to your network settings).

### Login  
- account: admin
- password: admin  

<img src="img/login.png" width=600>  

### Analog Input Dashboard
- Differential mode
<img src="img/dashboard_diff.png" width=600>

- Single-end mode
<img src="img/dashboard_single.png" width=600>

### Digital I/O Status
<img src="img/io_status.png" width=600>  

## Software Package Management
The PAC-6070 uses Ubuntu's APT (Advanced Package Tool) for software package management. Here are the common package management commands:

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
The PAC-6070 supports storage expansion through its built-in micro SD card slot. When you insert an SD card, the system will detect it automatically. You can use the `lsblk` command to view the device identifier (usually mmcblk0), and then mount it to any directory (e.g. /media) using the `mount` command. The SD card must be properly formatted with a supported filesystem like FAT32 or ext4 before mounting.

Before SD Insertion
```
root@pac6070:~# lsblk
NAME         MAJ:MIN RM  SIZE RO TYPE MOUNTPOINTS
mmcblk1      179:0    0 14.6G  0 disk 
├─mmcblk1p1  179:1    0    2G  0 part 
└─mmcblk1p2  179:2    0 12.6G  0 part /
mmcblk1boot0 179:8    0    4M  1 disk 
mmcblk1boot1 179:16   0    4M  1 disk 
```  
After SD Insertion
```
root@pac6070:~# lsblk
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
root@pac6070:~# mount /dev/mmcblk0p1 /media/
root@pac6070:~# lsblk
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
root@pac6070:~# umount /media/
```  










## RS-485 Serial Port Settings
### Port Mapping
PAC-6070 come with one RS-485 port.

|Port Number|Device Mapping|
|---|---|
|1|/dev/ttymxc1|

### Configure the Serial Port
If you need to modify the serial port settings, use the **setuart** command.
```
root@pac6070:~# setuart -h
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
root@pac6070:~# setuart -p1
Port 1 ==> type: RS485
baud: 115200
```

## Network Interface Settings
PAC-6070 come two 10/100Mbit Ethernet ports, the default network settings are shown below:  
|Port Label|Device mapping|IP mode|IP address|
|---|---|---|---|
|LAN1|eth0|DHCP|auto|
|LAN2|eth1|static|192.168.2.127|

### Display the Network Settings
```
root@pac6070:~# ip a show eth0
root@pac6070:~# ip a show eth1
```

### Modify the Network Settings
To modify network settings according to your LAN environment, follow these steps:

1. Edit the network interface configuration file:
   ```
   vi /etc/network/interfaces
   ```

2. After saving your changes, apply the new settings by executing these commands:
   ```
   systemctl restart NetworkManager   # Restart the network manager service
   ifconfig eth0 down                 # Shutdown the network interface
   ifconfig eth0 up                   # Start up the network interface
   ```

Note: Replace eth0 with the appropriate interface name (eth0 or eth1) that you want to configure.
```
root@pac6070:~# systemctl restart NetworkManager
root@pac6070:~# ifconfig eth0 down && ifconfig eth0 up
root@pac6070:~# ip a show eth0
```


