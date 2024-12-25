# Matrix-752 Technical Documentation

# 1. Introduction
Matrix-752, based on the ARM Cortex-A7 architecture, is a Linux-ready IoT gateway offering high integration and low power consumption. It serves as an ideal solution for various markets, including industrial control, automation, mobile gateways, and other applications.

## 1.1 Features
- **Processor:** NXP iMX6ULL 800MHz Cortex-A7  
- **OS Support:** Ubuntu/Linux kernel 5.10.x (or later)  
- **Toolchain:** gcc 9.3.0 + glibc 2.31  
- **Memory:** 512MB LvDDR3 SDRAM  
- **Networking:** Two 10/100Mbps Ethernet ports  
- **USB:** One USB OTG port  
- **Serial Ports:**  
  - 1 x RS-485/RS-232  
  - 1 x RS-232  
- **CAN Port:** One CAN Bus 2.0 A/B  
- **Digital I/O:**  
  - 2 x Digital Input channels  
  - 2 x Digital Output channels  
- **Expansion and Storage:**  
  - 1 x microSD socket  
  - 1 x full-size miniPCIe socket  
  - 1 x micro-SIM socket  
- **Other Features:**  
  - Two SMA-type Antenna holes reserved  
  - +9 to +48VDC power input  
  - Ultra-low power consumption  
  - Wall-mounting and optional DIN RAIL mounting adapter  

## 1.2 Specifications (Hardware)

### **CPU / Memory**
- **CPU:** NXP iMX6ULL Cortex-A7 MPCore, up to 800MHz  
- **Memory:** 512MB LvDDR3  

### **Network Interface**
- **Type:** 2 x 10/100Mbps Ethernet  
- **Connector:** RJ45 (with LED indicators)  

### **USB Interface**
- 1 x USB OTG Port (Micro-USB connector)  

### **TTY (Serial) Ports**
- **Ports:**  
  - 1 x RS-485/RS-232  
  - 1 x RS-232  
- **Connector:** Terminal block  
- **RS-485 Signals:** Data+, Data-  
- **RS-232 Signals:** TX, RX  

### **TTY Port Parameters**
- **Baud Rate:** Up to 921.6Kbps  
- **Parity:** None, Even, Odd, Mark, Space  
- **Data Bits:** 5, 6, 7, 8  
- **Stop Bits:** 1, 1.5, 2  
- **Flow Control:** XON/XOFF, None  

### **CAN Bus**
- **Compliance:** CAN Bus 2.0 A/B  
- **Speed:** Up to 1Mbps  
- **Isolation:** 2500Vrms  

### **Digital Input**
- 2 x Channels  
- **Protection:** 2500Vrms isolation  
- **Logical Levels:**  
  - High: 5~24VDC  
  - Low: 0~1.5VDC  

### **Digital Output**
- 2 x Channels (Solid State Relay)  
- **Contact Rating:** 80VDC @ 1.5A  

### **Console / Debug Ports**
- **Type:** Serial console port  
- **Connector:** 4PIN Box header  

### **SD Slot**
- **Socket:** 1 x microSD  
- **Standard:** SD 2.0 compliant (SDHC, up to 128G)  

### **Expansion Slot**
- 1 x full-size miniPCIe socket  
- 1 x micro-SIM card socket (USB interface)  
- 2 x SMA-type antenna holes reserved  

### **Power Requirement**
- **Input Voltage:** +9~+48VDC  
- **Typical Power:** 12VDC @ 250mA  

### **General**
- **Watchdog Timer:** Yes  
- **RTC:** Yes (backup by supercapacitor)  
- **Dimensions:** 89 x 112 x 30mm  
- **Weight:** 350g  
- **Operating Temperature:** 0~70°C  
- **Regulations:** CE Class A, FCC Class A  
- **Installation:** Wall or optional DIN-rail mount  



## 1.3 Specifications (Software)

### **Operating System**
- Supports Ubuntu/Linux kernel 5.10.x (or later)  
- Boot from eMMC or SD card  
- Backup/Restore from SD card or USB  
- **Bootloader:** U-Boot  
- **File System:** EXT4/EXT3/EXT2, VFAT/FAT, NFS  

### **Software Development**
- **Toolchain:** gcc 9.3.x + glibc 2.31  
- In-place C/C++ code compilation  

### **Package Management**
- **Repository:** Artila self-maintained  
- **Command:** Standard `apt-get`  

### **Popular Packages**
- **Web Servers:** Apache, Nginx, Lighttpd  
- **Databases:** MySQL, SQLite3, PostgreSQL  
- **Languages:** PHP, Python, Perl, Node.js  
- **Editors:** vim, nano, sed  
- **Administration:** Webmin  



## 1.4 Packing List
- **Matrix-752:** Linux-ready Cortex-A7 800MHz Industrial IoT Gateway with 512MB SDRAM and 16G eMMC  



## 1.5 Optional Accessories
- **DK-35A:** DIN Rail Mounting Kit  
- **PWR-12V-1A:** 110~240VAC to 12VDC 1A Power Adapter  
- **Console Cable:** 4Pin Wafer Box to DB9 Female (50cm)  



## 1.6 Optional Communication Modules
- 4G/LTE miniPCIe Module with Antenna  
- WiFi miniPCIe Module with Antenna  


# 2.	Layout        　　　　　　　        
## 2.1	Connector & LED Indicator
<img src="./images/01 Indicators.jpg" alt="alt text" width="500">


## 2.2	Dimension 
(Unit: mm)

<img src="./images/02 dimension.jpg" alt="alt text" width="400" height="450">

# 3. Pin Assignment and Definitions

## 3.1 Multi-function Reset Button

The Matrix-752 provides a multi-function reset button located on the side of the chassis as shown below:

<img src="./images/03 reset button.jpg" alt="alt text" width="400">

The behavior of the reset button depends on how long you press the reset button.
| **Press and Hold**              | **Behavior**                              | **Network Settings After Reboot**         |
|---------------------------------|-------------------------------------------|-------------------------------------------|
| < 3 seconds then release         | Reboot the Matrix-752                    | Retains last user settings                |
| 3–10 seconds then release        | Reset the network settings (factory default) | `eth0` IP: addr. by DHCP<br>`eth1` IP: 192.168.2.127 |
| > 10 seconds then release        | Restore and return to factory default. <br>User data may be erased. | `eth0` IP: addr. by DHCP<br>`eth1` IP: 192.168.2.127 |



## 3.2 LED Indicators

The LED indicators provide operational information for the Matrix-752:

<img src="./images/04 LED Indicators.jpg" alt="alt text" width="400">

- **“Ready” LED**: Stays ON when the system is ready for operation.
- **“SERIAL 1” & “SERIAL 2” LEDs**: Dual-color LEDs indicate data traffic:
  - **Green**: RXD line is high.
  - **Yellow**: TXD line is high.



## 3.3 Ethernet LAN Port

There are two 10/100Mbps Ethernet ports using RJ45 connectors with LED indicators. 

<img src="./images/4 LAN Port Indicators.jpg" alt="alt text" width="300">

<img src="./images/05 RJ45 Pin Assignment.jpg" alt="alt text" width="400">



## 3.4 Power Connector

The Matrix-752 supports a +9 to +48VDC power line connected via a terminal block.


<img src="./images/06 power input-1.jpg" alt="alt text" width="300">


## 3.5 USB OTG Port

The Matrix-752 is equipped with one USB OTG port using a micro-USB connector.


<img src="./images/07 USB OTG.jpg" alt="alt text" width="400">


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

<img src="./images/08 Serial Port.jpg" alt="alt text" width="450">

### Enable/Disable RS-485 Termination Resistor (JP4):
The Matrix-752 provides on-board 120Ohm termination resistor for each RS-485 port. To enable the termination resistor, please remove the upper cover of the Matrix-752, and the adjust the associated jumper to short as below:

<img src="./images/09 RS486 JP4-1.jpg" alt="alt text" width="450">

### Set Serial Port 1 to RS232 port (JP3 & JP6)
The Serial Port on Matrix-752, default is RS-485 at JP3 (setting Pin 3 and Pin4)
To Enable RS-232 port, setting JP3 at Pin 1 and Pin 2

<img src="./images/10 RS485 & 232 jumper settings.jpg" alt="alt text" width="450">

In the meantime, it should set D+/TX1 definition at JP6.
Default setting is for RS-485 (D+) at JP6/Pin 2 and Pin 3.
To Enable RS232 / TX1, setting JP6 at Pin 1 and Pin 2

<img src="./images/11 RS485 & 232 TX1 jumper settings.jpg" alt="alt text" width="450">


## 3.7 CAN Bus Port

The Matrix-752 has one CAN Bus port:

<img src="./images/12 Can Bus Port.jpg" alt="alt text" width="300">


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

<img src="./images/13 CAN Bus JP5.jpg" alt="alt text" width="450">


## 3.8 Digital Input

Two digital input channels with 5000Vrms photocoupler isolation.

<img src="./images/14 Digital Input.jpg" alt="alt text" width="300">

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

<img src="./images/15 DI Simple figure.jpg" alt="alt text" width="300">

- **DIx:** Isolated digital input channels.

- **COM:** Common ground


## 3.9 Digital Output

Two digital output channels using solid-state relays.

-  Normal Open (NO)**:Normal Open (NO) Type
-  Contact Rating:-  80VDC @ 1.5A  

<img src="./images/16 DO digital output.jpg" alt="alt text" width="300">

### Pin Assignment of Digital Input:

| **Pin** | **Signal** |
|---------|------------|
| 12      | DO1        |
| 13      | COM        |
| 14      | DO2        |
| 15      | COM        |

Reference Circuit of following:

![alt text](<./images/17 DO reference circuit-1.jpg>)



## 3.10 Serial Console Port

There is a 4-pin wafer box header (JP2) inside the Matrix-752 features as serial console port that used for locally accessing Matrix-752 system via console port.

<img src="./images/18 Consol Port.jpg" alt="alt text" width="400">

<img src="./images/19 Console Pin Assignment.jpg" alt="alt text" width="400">



## 3.11 External Battery Connection

There is a 2Pin wafer (1.2mm pitch) reserved that can be connected to external battery for RTC

<img src="./images/20 External Battery Connection.jpg" alt="alt text" width="480">


## 3.12 SD Card Socket

The micro-SD card socket inside the Matrix-752 can be accessed after removing the top cover. It supports additional data storage.

<img src="./images/21 SD CArd Socket.jpg" alt="alt text" width="500">


## 3.13 miniPCIe Slot

The Matrix-752 comes with one miniPCIe slot and dual antenna holes reserved for communication/networking modules.

<img src="./images/22 miniPCIe Slot-1.jpg" alt="alt text" width="400">


## 3.14 Micro-SIM Card Socket

The Matrix-752 includes a micro-SIM card socket inside. After removing the top cover, a micro-SIM card can be inserted for LTE/4G module functionality.

![alt text](<./images/23 SIM Card socket.jpg>)






