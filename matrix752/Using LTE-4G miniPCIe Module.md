# Using LTE-4G miniPCIe Module

The Matrix-752 comes with a miniPCIe solt which supports SIMCom LTE module.

## Using module SIM7600E
### Driver Installation
- `apt install -y kernel-module-option`
- `apt install -y kernel-module-sim7500-sim7600-wwan`
- `apt install -y kernel-module-usbnet`
- `reboot`

### Get an IP Address via DHCP
- `ifconfig wwan0 up` - activate /dev/wwan0
- `echo -ne "AT\$QCRMCALL=1,1\r" | eval /usr/bin/microcom -t 500 /dev/ttyUSB2`
- `udhcpc -i wwan0` - ask for IP address

### Check theWAN connection
- `ifconfig wwan0` - check the IP address
- `ping -I wwan0 www.google.com` - try pinging Google

## Using module SIM7100E
### Driver Installation
- `apt install -y kernel-module-option`
- `apt install -y kernel-module-qmi-wwan`
- `apt install -y kernel-module-cdc-wdm`
- `apt install -y kernel-module-usbnet`
- `apt install -y libqmi-utils`
- `reboot`

### Add config file
/etc/qmi-network.conf
```
APN=internet
```

### Get an IP Address via DHCP
```
/usr/sbin/modprobe -r qmi_wwan cdc_wdm
/usr/sbin/modprobe cdc_wdm && /usr/sbin/modprobe qmi_wwan
ifconfig wwan0 up
qmi-network /dev/cdc-wdm0 start
dhclient wwan0
```

### Check theWAN connection
- `ifconfig wwan0` - check the IP address
- `ping -I wwan0 www.google.com` - try pinging Google


