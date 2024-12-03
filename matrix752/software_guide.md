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
