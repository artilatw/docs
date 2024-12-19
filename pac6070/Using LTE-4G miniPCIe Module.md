# Using LTE miniPCIe Module

The Matrix-752 comes with a miniPCIe solt which supports SIMCom LTE module.

## Using SIMCom SIM7600E
### Driver Installation
- `apt install -y kernel-module-option`
- `apt install -y kernel-module-qmi-wwan`
- `apt install -y kernel-module-cdc-wdm`
- `apt install -y kernel-module-usbnet`
- `apt install -y udhcpc`
- `reboot` - reboot the Matrix-752 to activate the driver

### Get an IP Address via DHCP
- `ifconfig wwan0 up` - activate /dev/wwan0
- `echo -e "AT\$QCRMCALL=1,1\r" > /dev/ttyUSB2`
- `udhcpc -i wwan0` - ask for IP address

### Check the WAN connection
- `ifconfig wwan0` - check the IP address
- `ping -I wwan0 www.google.com` - try pinging Google
