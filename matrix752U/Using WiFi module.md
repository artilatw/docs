# Using Wi-Fi miniPCIe Module

The Matrix-752U comes with a miniPCIe slot which supports Wi-Fi module.

## Using RYWDB00
### Driver Installation
- `apt -y install kernel-module-rsi-91x kernel-module-rsi-usb linux-firmware-rsi`
- `reboot` - reboot system for activate the driver

### Display USB info
- `lsusb`  
  ```
  Bus 001 Device 004: ID 1618:9116 Redpine Signals, Inc. Wireless USB Network Module
  ```
- `usb-devices`
  ```
  T:  Bus=01 Lev=01 Prnt=01 Port=00 Cnt=01 Dev#=  4 Spd=480 MxCh= 0
  D:  Ver= 2.00 Cls=ff(vend.) Sub=ff Prot=ff MxPS=64 #Cfgs=  1
  P:  Vendor=1618 ProdID=9116 Rev=00.02
  S:  Manufacturer=Redpine Signals, Inc.
  S:  Product=Wireless USB Network Module
  S:  SerialNumber=000000000001
  C:  #Ifs= 1 Cfg#= 1 Atr=c0 MxPwr=300mA
  /usr/bin/usb-devices: 91: printf: WLAN: expected numeric value
  I:  If#= 0 Alt= 0 #EPs= 6 Cls=ff(vend.) Sub=ff Prot=ff Driver=RSI-USB
  I:  If#= 0 Alt= 0 #EPs= 0 Cls=() Sub= Prot= Driver=
  E:  Ad=01(O) Atr=02(Bulk) MxPS= 512 Ivl=0ms
  E:  Ad=02(O) Atr=02(Bulk) MxPS= 512 Ivl=0ms
  E:  Ad=03(O) Atr=02(Bulk) MxPS= 512 Ivl=0ms
  E:  Ad=81(I) Atr=02(Bulk) MxPS= 512 Ivl=0ms
  E:  Ad=82(I) Atr=02(Bulk) MxPS= 512 Ivl=0ms
  E:  Ad=83(I) Atr=02(Bulk) MxPS= 512 Ivl=0ms
  ```

### List of Access Points info 
- `nmcli device wifi list`

### Connecting to Access Point
- `nmcli device wifi connect [AP] password [PASSWORD]`
  
### Display Wi-Fi interface info
- `nmcli device show wlx88da1a7632a4`

### Disconnect
- `nmcli device disconnect wlx88da1a7632a4`

