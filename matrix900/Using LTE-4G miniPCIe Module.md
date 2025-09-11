# Using LTE miniPCIe Module

The Matrix-900 device includes a miniPCIe slot compatible with SIMCom LTE modules, such as the SIM7600G-H.

## Module Setup: SIMCom SIM7600G-H
### Required Utilities
Install the necessary tools:
```bash
sudo apt -y install socat udhcpc
```

### Enable miniPCIe Power and Configure GPIO
Activate power to the miniPCIe slot via GPIO-13 and ensure GPIO-11 is set to active low:
```bash
pinctrl set 11 op dl
pinctrl set 13 op dh
```

### Verify Module Detection
Use the following commands to confirm the module is recognized:
- Check module information using the commands lsusb and usb-devices:
```bash
lsusb
```
Expected output:
```bash
Bus 001 Device 003: ID 1e0e:9001 Qualcomm / Option SimTech, Incorporated
```

- View detailed USB device info:
```bash
usb-devices
```
Expected output includes:
```
T:  Bus=01 Lev=02 Prnt=02 Port=00 Cnt=01 Dev#=  3 Spd=480 MxCh= 0
D:  Ver= 2.00 Cls=00(>ifc ) Sub=00 Prot=00 MxPS=64 #Cfgs=  1
P:  Vendor=1e0e ProdID=9001 Rev=03.18
S:  Manufacturer=SimTech, Incorporated
S:  Product=SimTech, Incorporated
S:  SerialNumber=0123456789ABCDEF
C:  #Ifs= 6 Cfg#= 1 Atr=a0 MxPwr=500mA
I:  If#= 0 Alt= 0 #EPs= 2 Cls=ff(vend.) Sub=ff Prot=ff Driver=option
E:  Ad=01(O) Atr=02(Bulk) MxPS= 512 Ivl=0ms
E:  Ad=81(I) Atr=02(Bulk) MxPS= 512 Ivl=0ms
I:  If#= 1 Alt= 0 #EPs= 3 Cls=ff(vend.) Sub=00 Prot=00 Driver=option
E:  Ad=02(O) Atr=02(Bulk) MxPS= 512 Ivl=0ms
E:  Ad=82(I) Atr=02(Bulk) MxPS= 512 Ivl=0ms
E:  Ad=83(I) Atr=03(Int.) MxPS=  10 Ivl=32ms
I:  If#= 2 Alt= 0 #EPs= 3 Cls=ff(vend.) Sub=00 Prot=00 Driver=option
E:  Ad=03(O) Atr=02(Bulk) MxPS= 512 Ivl=0ms
E:  Ad=84(I) Atr=02(Bulk) MxPS= 512 Ivl=0ms
E:  Ad=85(I) Atr=03(Int.) MxPS=  10 Ivl=32ms
I:  If#= 3 Alt= 0 #EPs= 3 Cls=ff(vend.) Sub=00 Prot=00 Driver=option
E:  Ad=04(O) Atr=02(Bulk) MxPS= 512 Ivl=0ms
E:  Ad=86(I) Atr=02(Bulk) MxPS= 512 Ivl=0ms
E:  Ad=87(I) Atr=03(Int.) MxPS=  10 Ivl=32ms
I:  If#= 4 Alt= 0 #EPs= 3 Cls=ff(vend.) Sub=00 Prot=00 Driver=option
E:  Ad=05(O) Atr=02(Bulk) MxPS= 512 Ivl=0ms
E:  Ad=88(I) Atr=02(Bulk) MxPS= 512 Ivl=0ms
E:  Ad=89(I) Atr=03(Int.) MxPS=  10 Ivl=32ms
I:  If#= 5 Alt= 0 #EPs= 3 Cls=ff(vend.) Sub=ff Prot=ff Driver=simcom_wwan
E:  Ad=06(O) Atr=02(Bulk) MxPS= 512 Ivl=0ms
E:  Ad=8a(I) Atr=02(Bulk) MxPS= 512 Ivl=0ms
E:  Ad=8b(I) Atr=03(Int.) MxPS=   8 Ivl=32ms
```

### Identify Device Ports
Run the following to locate the assigned USB serial ports:
```bash
dmesg | grep GSM
```
Example output:
```
usb 1-1.1: GSM modem (1-port) converter now attached to ttyUSB4
usb 1-1.1: GSM modem (1-port) converter now attached to ttyUSB5
...
usb 1-1.1: GSM modem (1-port) converter now attached to ttyUSB8
```
### Network Interface Setup
Bring up the WWAN interface and initiate a data connection:
```bash
sudo ifconfig wwan0 up
sudo sh -c 'echo "AT\$QCRMCALL=1,1\r" |eval socat -T .5 - /dev/ttyUSB6,crnl'
sudo udhcpc -i wwan0
```

### WAN Connectivity Test
- Check IP assignment:
```bash
ifconfig wwan0
```
- Test internet access:
```bash
ping -I wwan0 www.google.com
```
### Troubleshooting & Diagnostics
#### Check SIM card status
```bash
sudo sh -c 'echo "AT+CPIN?" |eval socat -T .5 - /dev/ttyUSB6,crnl'
```

#### Check Signal Quality
```bash
sudo sh -c 'echo "AT+CSQ?" |eval socat -T .5 - /dev/ttyUSB6,crnl'
```