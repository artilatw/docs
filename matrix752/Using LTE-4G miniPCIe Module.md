# Using LTE-4G miniPCIe Module

The Matrix-752 comes with a miniPCIe solt which supports SIMCom LTE module.

## Using module SIM7600E
### Driver Installation
- `apt install -y kernel-module-option`
- `apt install -y kernel-module-sim7500-sim7600-wwan`
- `apt install -y kernel-module-usbnet`

### Activate the Driver
- `modprobe sim7500_sim7600_wwan` - add /dev/wwan0
- `modprobe usbnet`
- `ifconfig wwan0 up` - activate /dev/wwan0
- `echo -ne "AT\$QCRMCALL=1,1\r" | eval /usr/bin/microcom -t 500 /dev/ttyUSB2`
- `udhcpc -i wwan0` - ask for IP address

### Check theWAN connection
- `ifconfig wwan0` - check the IP address
- `ping -I wwan0 www.google.com` - try pinging Google

## Using module SIM7100E
To be continued...


