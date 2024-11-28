# Using LTE miniPCIe Module

The Matrix-752 comes with a miniPCIe solt which supports SIMCom LTE module.

## Using SIMCom SIM7600E
### Driver Installation
- `apt install -y kernel-module-option`
- `apt install -y kernel-module-sim7500-sim7600-wwan`
- `apt install -y kernel-module-usbnet`
- `reboot` - reboot the PAC-6070 to activate the driver

### Get an IP Address via DHCP
- `ifconfig wwan0 up` - activate /dev/wwan0
- `echo -ne "AT\$QCRMCALL=1,1\r" | eval /usr/bin/microcom -t 500 /dev/ttyUSB2`
- `udhcpc -i wwan0` - ask for IP address

### Check the WAN connection
- `ifconfig wwan0` - check the IP address
- `ping -I wwan0 www.google.com` - try pinging Google
