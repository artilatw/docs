# Using LTE miniPCIe Module

The Matrix-900 features a miniPCIe slot that supports SIMCom LTE modules.

## Using SIMCom SIM7600G-H
Installation require utilities
```
sudo apt -y install socat udhcpc
```

Open mPCIe Power and verify GPIO-11 is active low
```
pinctrl set 11 op dl
pinctrl set 13 op dh
```

Check module information using the commands lsusb and usb-devices:
- lsusb displays
```
Bus 001 Device 003: ID 1e0e:9001 Qualcomm / Option SimTech, Incorporated
```

- usb-devices displays detailed information
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

Retrieve device port information using the following commands
```
guest@matrix900:~$ dmesg|grep GSM
[  493.913438] usbserial: USB Serial support registered for GSM modem (1-port)
[  493.922675] option 1-1.1:1.0: GSM modem (1-port) converter detected
[  493.923446] usb 1-1.1: GSM modem (1-port) converter now attached to ttyUSB4
[  493.923792] option 1-1.1:1.1: GSM modem (1-port) converter detected
[  493.924253] usb 1-1.1: GSM modem (1-port) converter now attached to ttyUSB5
[  493.925102] option 1-1.1:1.2: GSM modem (1-port) converter detected
[  493.926180] usb 1-1.1: GSM modem (1-port) converter now attached to ttyUSB6
[  493.926438] option 1-1.1:1.3: GSM modem (1-port) converter detected
[  493.926986] usb 1-1.1: GSM modem (1-port) converter now attached to ttyUSB7
[  493.927373] option 1-1.1:1.4: GSM modem (1-port) converter detected
[  493.927811] usb 1-1.1: GSM modem (1-port) converter now attached to ttyUSB8
```

Set up the wwan0 interface and obtain an IP address
```
sudo ifconfig wwan0 up
sudo sh -c 'echo "AT\$QCRMCALL=1,1\r" |eval socat -T .5 - /dev/ttyUSB6,crnl'
sudo udhcpc -i wwan0
```

## Testing WAN connection
- `ifconfig wwan0` - check the IP address  
- `ping -I wwan0 www.google.com` - try pinging Google
