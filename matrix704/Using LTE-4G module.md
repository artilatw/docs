# Using LTE miniPCIe Module

The Matrix-704 comes with a miniPCIe slot which supports SIMCom LTE module.

## Driver Installation
First, update the repository.
```bash
apt update
```

Install require packages:
```bash
apt -y install kernel-module-option kernel-module-sim7500-sim7600-wwan kernel-module-usbnet
```
## Configuration
### Set wwan0 up
```bash
ifconfig wwan0 up
```
### Dial up via module
```bash
echo -e "AT\$QCRMCALL=1,1\r" > /dev/ttyUSB2
```

### Get an IP Address via DHCP
```bash
udhcpc -i wwan0
```

## Check the WAN connection
### Check the WAN IP address
```bash
ifconfig wwan0
```
### Ping
```bash
ping -I wwan0 www.google.com
```
