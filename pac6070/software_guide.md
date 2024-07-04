# PAC-6070

## Linux-ready Cortex-A7 Highly Integrated Industrial IoT Gateway
- NXP i.MX6ULL Cortex-A7 CPU, Up to 800MHz
- Support Ubuntu 22.04
- Toolchain: gcc 9.3.0 + glibc 2.31
- 512MB LvDDR3 SDRAM, 16GB eMMC
- 2 x 10/100Mbps Ethernet interface
- 1 x USB OTG port
- 1 x RS-485
- 1 x CAN port, 2 x Digital Input, 2 x Relay out
- 1 x full size miniPCIe socket inside,1 x Micro-SIM slot reserved, Two Antenna holes reserved
- 1 x Micro-SD slot reserved

## Setup analog input
### iio Path:  
/sys/bus/iio/devices/iio:device0/  
![iio_path](img/iio_path.png)  
### Formula:  
raw * scale + offset
```
root@pac6070:/sys/bus/iio/devices/iio:device0# cat in_voltage0-voltage1_raw
8387948
root@pac6070:/sys/bus/iio/devices/iio:device0# cat in_voltage0-voltage1_offset
-24038
root@pac6070:/sys/bus/iio/devices/iio:device0# cat in_voltage-voltage_scale
0.002865602
```
### Mode:  
- Single-end=0
- Differential=1
```
root@pac6070:~# cat /etc/modprobe.d/ad4111.conf
options ad4111 differential=1
```
reboot after modify the config file    

### Overview:  
`iio_info`
```
root@pac6070:~# iio_info
Library version: 0.23 (git tag: v0.23)
Compiled with backends: local xml ip usb
IIO context created with local backend.
Backend version: 0.23 (git tag: v0.23)
Backend description string: Linux pac6070 6.6.22 #1 Fri Mar 15 18:25:07 UTC 2024 armv7l
IIO context has 2 attributes:
        local,kernel: 6.6.22
        uri: local:
IIO context has 2 devices:
        iio:device0: ad4111 (buffer capable)
                5 channels found:
                        voltage0-voltage1:  (input, index: 0, format: be:u24/32>>0)
                        4 channel-specific attributes found:
                                attr  0: offset value: -24038
                                attr  1: raw value: 8387952
                                attr  2: scale value: 0.002865602
                                attr  3: scale_available value: 0.000000000
                        voltage2-voltage3:  (input, index: 1, format: be:u24/32>>0)
                        4 channel-specific attributes found:
                                attr  0: offset value: -24038
                                attr  1: raw value: 8387864
                                attr  2: scale value: 0.002865602
                                attr  3: scale_available value: 0.000000000
                        voltage4-voltage5:  (input, index: 2, format: be:u24/32>>0)
                        4 channel-specific attributes found:
                                attr  0: offset value: -24038
                                attr  1: raw value: 8387186
                                attr  2: scale value: 0.002865602
                                attr  3: scale_available value: 0.000000000
                        current0:  (input, index: 3, format: be:u24/32>>0)
                        4 channel-specific attributes found:
                                attr  0: offset value: 0
                                attr  1: raw value: 0
                                attr  2: scale value: 0.000002980
                                attr  3: scale_available value: 0.000000000
                        current1:  (input, index: 4, format: be:u24/32>>0)
                        4 channel-specific attributes found:
                                attr  0: offset value: 0
                                attr  1: raw value: 0
                                attr  2: scale value: 0.000002980
                                attr  3: scale_available value: 0.000000000
                3 device-specific attributes found:
                                attr  0: sampling_frequency value: 31250
                                attr  1: sampling_frequency_available value: 31250 31250 31250 31250 31250 31250 15625 10417 5208
                                attr  2: waiting_for_supplier value: 0
                3 buffer-specific attributes found:
                                attr  0: data_available value: 0
                                attr  1: direction value: in
                                attr  2: watermark value: 1
                Current trigger: trigger0(ad4111-dev0)
        trigger0: ad4111-dev0
                0 channels found:
                No trigger on this device
``` 

### Display - CLI:  
`lsadc`  
<img src="img/cli.png" width=600>
### Website
Login  
- account: admin
- password: admin  

<img src="img/login.png" width=600>  

### Dashboard
<img src="img/dashboard.png" width=600>  