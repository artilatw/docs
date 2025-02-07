# Using LTE miniPCIe Module

The Matrix-752U comes with a miniPCIe solt which supports SIMCom LTE module.

### Upgrde the kernel
The default kernel version is 6.1.19, which is not compatible with the LTE module. We need to upgrade the kernel to 6.1.46.

- `apt update;apt full-upgrade;apt dist-upgrade` - please wait for the upgrade to finish
- `rm /etc/alternatives/zImage`
- `ln -s /boot/zImage-6.1.46-rt13 /etc/alternatives/zImage`
- `reboot` - reboot the Matrix-752U to activate the new kernel

### Driver Installation

#### Required Packages:
- kernel-module-option
- kernel-module-qmi-wwan
- kernel-module-cdc-wdm
- kernel-module-usbnet
- udhcpc

Install all required packages:

- `apt update`
- `apt install -y kernel-module-option kernel-module-qmi-wwan kernel-module-cdc-wdm kernel-module-usbnet udhcpc`


- `reboot` - reboot the Matrix-752U to activate the driver

### Get an IP Address via DHCP
- `ifconfig wwan0 up` - activate /dev/wwan0
- `echo -e "AT\$QCRMCALL=1,1\r" > /dev/ttyUSB2`
- `udhcpc -i wwan0` - ask for IP address

### Check the WAN connection
- `ifconfig wwan0` - check the IP address
- `ping -I wwan0 www.google.com` - try pinging Google
