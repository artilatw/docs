# Using LTE miniPCIe Module

The Matrix-772 comes with a miniPCIe slot which supports SIMCom LTE module.

## Using SIMCom SIM7600E

### Update the Kernel
- Extract vmlinuz-6.6.zip to `/boot`.  
- `reboot` - reboot the Matrix-772 to activate the new kernel.

### Utility Installation
- `apt install -y udhcpc`

### Get an IP Address via DHCP
- `echo Y > /sys/class/net/wwan0/qmi/raw_ip` - set the raw IP mode  
- `ifconfig wwan0 up` - activate /dev/wwan0  
- `echo -e "AT\$QCRMCALL=1,1\r" > /dev/ttyUSB2`  
- `udhcpc -i wwan0` - ask for IP address

### Check the WAN connection
- `ifconfig wwan0` - check the IP address  
- `ping -I wwan0 www.google.com` - try pinging Google
