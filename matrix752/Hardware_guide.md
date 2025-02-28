# Matrix-752 Hardware Guide

# 1. Introduction
Matrix-752, based on the ARM Cortex-A7 architecture, is a Linux-ready IoT gateway offering high integration and low power consumption. It serves as an ideal solution for various markets, including industrial control, automation, mobile gateways, and other applications.

## 1.1 Packing List
- **Matrix-752:** Linux-ready Cortex-A7 800MHz Industrial IoT Gateway with 512MB SDRAM and 16G eMMC  

## 1.2 Optional Accessories
- **DK-35A:** DIN Rail Mounting Kit  
- **PWR-12V-1A:** 110~240VAC to 12VDC 1A Power Adapter  
- **Console Cable:** 4Pin Wafer Box to DB9 Female (50cm)  

## 1.3 Optional Communication Modules
- 4G/LTE miniPCIe Module with Antenna  
- WiFi miniPCIe Module with Antenna  


# 2.	Layout        　　　　　　　        
## 2.1	Connector & LED Indicator
![Connector & LED Indicator](images/01%20Indicators.jpg)


## 2.2	Dimension 
(Unit: mm)

![Dimension](images/02%20dimension.jpg)

# 3. Pin Assignment and Definitions

## 3.1 Multi-function Reset Button

The Matrix-752 provides a multi-function reset button located on the side of the chassis as shown below:

![Reset Button](images/03%20reset%20button.jpg)

The behavior of the reset button depends on how long you press the reset button.
| **Press and Hold**              | **Behavior**                              | **Network Settings After Reboot**         |
|---------------------------------|-------------------------------------------|-------------------------------------------|
| < 3 seconds then release         | Reboot the Matrix-752                    | Retains last user settings                |
| 3–10 seconds then release        | Reset the network settings (factory default) | `eth0` IP: addr. by DHCP<br>`eth1` IP: 192.168.2.127 |
| > 10 seconds then release        | Restore and return to factory default. <br>User data may be erased. | `eth0` IP: addr. by DHCP<br>`eth1` IP: 192.168.2.127 |



## 3.2 LED Indicators

The LED indicators provide operational information for the Matrix-752:

![LED Indicators](images/04%20LED%20Indicators.jpg)

- **“Ready” LED**: Stays ON when the system is ready for operation.
- **“SERIAL 1” & “SERIAL 2” LEDs**: Dual-color LEDs indicate data traffic:
  - **Green**: RXD line is high.
  - **Yellow**: TXD line is high.



## 3.3 Ethernet LAN Port

There are two 10/100Mbps Ethernet ports using RJ45 connectors with LED indicators. 

  ![LAN Port Indicators](images/04%20LED%20Indicators.jpg)

![RJ45 Pin Assignment](images/05%20RJ45%20Pin%20Assignment.jpg)
)



## 3.4 Power Connector

The Matrix-752 supports a +9 to +48VDC power line connected via a terminal block.


![Power Input](images/06%20power%20input-1.jpg)


## 3.5 USB OTG Port

The Matrix-752 is equipped with one USB OTG port using a micro-USB connector.


![USB OTG](images/07%20USB%20OTG.jpg)


## 3.6 Serial Port

The Matrix-752 has two serial ports:

- **Serial Port 1**: RS-485/RS-232 (Default: RS-485).
- **Serial Port 2**: RS-232 (TX/RX).

RS-485 is designed without isolation that automatically direction controlled via software
### Serial Port Pin Assignments

| **Port No.**       | **Pin 1** | **Pin 2** | **Pin 3** | **Pin 4** | **Pin 5** |
|---------------------|-----------|-----------|-----------|-----------|-----------|
| **Serial Port 1** RS-485  | D+ | D- | --        | --        | GND       |
| **Serial Port 1** RS-232 (1)  |TX1 | RX1 | --        | --        | GND       |
| **Serial Port 2** RS-232 (2)| --        | --        | TX2 | RX2 | GND      |

![Serial Port](images/08%20Serial%20Port.jpg)

### Enable/Disable RS-485 Termination Resistor (JP4):
The Matrix-752 provides on-board 120Ohm termination resistor for each RS-485 port. To enable the termination resistor, please remove the upper cover of the Matrix-752, and the adjust the associated jumper to short as below:

![RS486 JP4-1](images/09%20RS486%20JP4-1.jpg)

### Set Serial Port 1 to RS232 port (JP3 & JP6)
The Serial Port on Matrix-752, default is RS-485 at JP3 (setting Pin 3 and Pin4)
To Enable RS-232 port, setting JP3 at Pin 1 and Pin 2

![RS485 & 232 jumper settings](images/10%20RS485%20&%20232%20jumper%20settings.jpg)

In the meantime, it should set D+/TX1 definition at JP6.
Default setting is for RS-485 (D+) at JP6/Pin 2 and Pin 3.
To Enable RS232 / TX1, setting JP6 at Pin 1 and Pin 2

![RS485 & 232 TX1 jumper settings](images/11%20RS485%20&%20232%20TX1%20jumper%20settings.jpg)


## 3.7 CAN Bus Port

The Matrix-752 has one CAN Bus port:

![Can Bus Port](images/12%20Can%20Bus%20Port.jpg)


Users can open the CAN bus port as network sockets, the socket names are 'can0'
| **Port Label** | **Socket Mapping** |
|----------------|---------------------|
| CAN            |   can0              |

### Pin Assignment of CAN Bus port:

| **Pin** | **Signal**      |
|---------|-----------------|
| 6       | CAN_Hi (CAN+)   |
| 7       | CAN_Lo (CAN-)   |


### Enable/Disable Termination resistor for CAN bus (JP5)
The Matrix-752 provides on-board 120Ohm termination resistor for each CAN port. 
Default setting is “Disable” the terminal resistor for CAN bus.
To enable the termination resistor, please remove the upper cover of the Matrix-752, and the adjust the associated jumper to short position1 and position 2, shown below:

![CAN Bus JP5](images/13%20CAN%20Bus%20JP5.jpg)


## 3.8 Digital Input

Two digital input channels with 5000Vrms photocoupler isolation.

![Digital Input](images/14%20Digital%20Input.jpg)

### Pin Assignment of Digital Input:

| **Pin** | **Signal** |
|---------|------------|
| 8       | DI1        |
| 9       | COM        |
| 10      | DI2        |
| 11      | COM        |

### Specifications of the isolated input channels:

- **Logical High**: 5~24VDC  
- **Logical Low**: 0~1.5VDC  
- **Input Resistance**: 1.8KΩ @ 0.32W  
- **Response Time**: 20µs  
- **Isolation**: 5000Vrms  

![DI Simple figure](images/15%20DI%20Simple%20figure.jpg)

- **DIx:** Isolated digital input channels.

- **COM:** Common ground


## 3.9 Digital Output

Two digital output channels using solid-state relays.

-  Solid State Relay: Normal Open (NO) Type
-  Contact Rating: 80VDC @ 1.5A  

![DO digital output](images/16%20DO%20digital%20output.jpg)

### Pin Assignment of Digital Input:

| **Pin** | **Signal** |
|---------|------------|
| 12      | DO1        |
| 13      | COM        |
| 14      | DO2        |
| 15      | COM        |

Reference Circuit of following:

![DO reference circuit-1](images/17%20DO%20reference%20circuit-1.jpg)



## 3.10 Serial Console Port

There is a 4-pin wafer box header (JP2) inside the Matrix-752 features as serial console port that used for locally accessing Matrix-752 system via console port.

![Consol Port](images/18%20Consol%20Port.jpg)

![Console Pin Assignment](images/19%20Console%20Pin%20Assignment.jpg)



## 3.11 External Battery Connection

There is a 2Pin wafer (1.2mm pitch) reserved that can be connected to external battery for RTC

![External Battery Connection](images/20%20External%20Battery%20Connection.jpg)


## 3.12 SD Card Socket

The micro-SD card socket inside the Matrix-752 can be accessed after removing the top cover. It supports additional data storage.

![SD CArd Socket](images/21%20SD%20CArd%20Socket.jpg)


## 3.13 miniPCIe Slot

The Matrix-752 comes with one miniPCIe slot and dual antenna holes reserved for communication/networking modules.

![miniPCIe Slot-1](images/22%20miniPCIe%20Slot-1.jpg)


## 3.14 Micro-SIM Card Socket

The Matrix-752 includes a micro-SIM card socket inside. After removing the top cover, a micro-SIM card can be inserted for LTE/4G module functionality.

![SIM Card socket](images/23%20SIM%20Card%20socket.jpg)






