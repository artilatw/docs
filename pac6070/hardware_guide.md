## Introduction
The PAC-6070 is a cost-effective Linux computing platform powered by Cortex-A7 MPU. It offers a versatile range of on-board IOs, including voltage/current measurement, opto-isolated digital inputs, and high-power relay outputs.  With its comprehensive IO capabilities, the PAC-6070 enables efficient monitoring and control, making it the perfect choice for advanced Industrial IoT applications.  
PAC-6070 is powered by NXP i.MX6ULL Cortex-A7 Processor with 1GB SDRAM and 16GB MMC. PAC-6070 features 5 A/D channels：3 channels differential input mode (Voltage) and 2 channels single-end input mode (Current)，8 channels digital inputs, and 8 x relay outputs. In addition, PAC-6070 integrates two 10/100Mbps Ethernet, one isolated RS-485 ports, one USB host, and one MicroSD socket, one miniPCIe slot. The DIN-Rail also makes PAC-6070’s onsite installation flexible.

## Features
- Cost-Effective Linux Computing Platform
- Rugged Design for Industrial Automation Application
- NXP i.MX6ULL Arm Cortex-A7 CPU, up to 800MHz
- 1GB LvDDR3 SDRAM, 16GB eMMC
- 8 x Relay Outputs
- 8 x Opto-Isolated Digital Inputs
- 5 x 24-bits A/D Channels  
- 1 x Full Size miniPCIe Socket & 1 x Micro-SIM Slot
- 1 x Micro-SD Card Socket
- 2 x 10/100Mbps Ethernet Interface
- 1 x USB port
- 1 x RS-485 Serial Port
- +9~+48 VDC Wide-Range Power Input
- Dimension: 121.8mm x 204mm x 50mm (4.8" x 8" x1.96")
- Easy for DIN-Rail Mounting

## H/W Speciﬁcation
### CPU / Memory
- CPU: NXP i.MX6ULL Cortex-A7 MPCore, up to 800MHz
- SRAM: 1GB LvDDR3
- Flash: 16GB eMMC
### Network Interface
- Type: 2 x 10/100Mbps Ethernet
- Connector Type: RJ45
### Serial Port
- 1 x RS-485 (1500Vrms isolation)
- Termination Resistor: 120ohm (Disable as default)
- Signal: RS-485 (Data+, Data-)
- Connector: Terminal block
- LED Indicator: YES

### Digital Input
- 8 x digital input channels
- Isolation Protection: 5000Vrms (photo-coupler)
- Logical High: 5~24VDC
- Logical Low: 0~1.5VDC

### General
- Watchdog (WDT): YES
- Real-Time Clock(RTC): YES, backup by super capacitor
- Power Input: +9~+48 VDC (terminal block)
- Reset: Multi-function reset button
- Indicator: Reset, Power, READY, LAN, RS-485, DI, Relay-out
- Dimensions (WxLxH): 121.8 x 204 x 50mm (4.8 x 8 x 1.96in)
- Net Weight: 360g (0.791b)
- Operating Temperature: 0~+70°C (+32~+158°F)
- Regulation: CE Class A, FCC Class A
- Installation: DIN-rail mounting

### Console / Debug Port
- 1 x 4-Pin box connector
### SD Slot
- 1x Micro-SD socket, SD 2.0 compliant, supports SDHC
### Expansion
- 1 x miniPCIe slot, 1 x Micro-SIM socket
### USB Interface
- 1 x USB 2.0 Port, USB Type-A connector
### Relay Output
  - 8 x relay output channels
  - Contact Rating: 125VAC@0.5A / 30VDC@1.0A
  - Max. Switching Voltage: 125VAC / 60VDC
  - Max. Switching Current: 2A
### Analog Input
- 5 x 24-bits A/D channels
- Input Mode (Voltage): 3 channels (diﬀerential input )
- Input Range (Voltage): +/- 10VDC
- (Optional) 2 x Single-end channels, call for driver support
- Input Mode (Current): 2 channels (single-end) Input Current: 0~20mA

## S/W Specification
### Operation System
- Support Ubuntu and Embedded Linux, kernel 5.10.x (or up)
- Supports bootup from eMMC or SD card
- Supports backup from SD card or USB device

### Software Development
- Toolchain: gcc 9.3.0 + glibc 2.31
- Supports in-place C/C++ code compilation



### Ordering Information
### PAC-6070
- Linux-based Arm Cortex-A7 Industrial Automation Controller.

## Layout: Diagram and M/B photo

![alt text](<./6070images/01 PAC 6070 正面layout diagram圖-1.jpg>)

![alt text](<./6070images/02 PAC-6070 MB and dimension.jpg>)



## 	Pin Assignment and Definition
### 1. Reset Button
The PAC 6070 provides a multi-function reset button location shown below: 

![alt text](<./6070images/03 Reset Button.jpg>)

The behavior of the reset button depends on how long you press the reset button.

| Press and hold the reset button | Behavior                                          | Network settings after reboot                |
|---------------------------------|--------------------------------------------------|----------------------------------------------|
| < 3 seconds then release        | Re-boot the PAC-6070                           | Retains last user settings                   |
| 3~10 seconds then release       | Reset the network setting (The same setting as factory default) | eth0 IP: addr. by DHCP<br>eth1 IP: 192.168.2.127 |
| > 10 seconds then release       | Restore and back to factory default.<br>User's data may disappear | eth0 IP: addr. by DHCP<br>eth1 IP: 192.168.2.127 |

### 2. Power LED
The Power LED will show solid green if power is properly applied. 

![alt text](<./6070images/04 Power LED.jpg>)

### 3. Ready LED
After PAC-6070 is power on, the Ready LED will show solid orange if PAC-6070 complete system boots up.  If Ready LED is off during system boot up, please check if power input is correct.  Turn off the power and restart PAC-6070 again.  If Ready LED is still off, please contact the manufacturer for technical support. 

![alt text](<./6070images/05 Ready LED-1.jpg>)

### 4. LAN1 / LAN2 LED
When Ethernet ports are connected to the network, Link/Act will show solid green and if there is traffic in the Ethernet, this LED will flash.

![alt text](<./6070images/06 LAN1 and LAN2 LED.jpg>)

### 5. Serial Port LED
One dual color LEDs indicate the data traffic at the serial ports. Dual color LEDs indicate the data traffic at the serial ports. When the RXD line is high then green light is ON and when the TXD line is high, yellow light is ON.

<img src="./6070images/07 Serial Port LED.jpg" alt="alt text" width="320">

### 6. Ethernet Port (LAN1 / LAN2)

<img src="./6070images/081Ethernet Port.jpg" alt="alt text" width="350">
---

### 7 Serial Port

#### RS-485 (Data+, Data-)
-  **Data+**: Pull-up to V_ISO with a 1K Ohm resistor.
-  **Data-**: Pull-down to ground.

The PAC-6070 provides an on-board **120 Ohm termination resistor** for each RS-485 port. However, the termination resistor is **not enabled by default**.

- **To enable the termination resistor**: Adjust the associated jumper to short as shown in the follwing figure.
- Alternatively, the user can add a **120 Ohm resistor** shunting Data+ to Data- if necessary. 

![alt text](<./6070images/09 RS485 jumper settings-3.jpg>)

 ○ ● ● (1, 2, 3) — Default Settings  
 ● ● ○ (1, 2, 3) — Add a 120 Ohm resistor 


### 8. Serial Console Port ：Debug Port（ RS-232 with RxD, TxD ）
There is a 4-pin wafer box header (J4) inside the PAC-6070 features as serial console port that used for locally accessing PAC-6070 system via console port.

#### **Pin Debug Port**

| Pin | Debug Port  |
|-----|-------------|
| 1   | RXD         |
| 2   | TXD         |
| 3   | VPERI_3.3V  |
| 4   | GND         |


![alt text](<./6070images/10 Console port 1.jpg>)	

### 9. Power Input Connector (CN1)
PAC-6070 uses +9VDC to 48VDC power and input from CN1 connector.

![alt text](<./6070images/11 9V48VDC.jpg>)

### 10. Relay Output
-  8 x relay output channels
-  Contact Rating: 125VAC@0.5A / 30VDC@1.0A
-  Max. Switching Voltage: 125VAC / 60VDC
-  Max. Switching Current: 2A

<img src="./6070images/12 Relay Output Layout.jpg" alt="alt text" width="570">

#### Pin Assignment of Digital Output :

| Pin   | Pin1  | Pin2  | Pin3  | Pin4  | Pin5  | Pin6  | Pin7  | Pin8  |
|-------|-------|-------|-------|-------|-------|-------|-------|-------|
| Signal| COM1  | DO1   | COM2  | DO2   | COM3  | DO3   | COM4  | DO4   |



| Pin    | Pin9  | Pin10 | Pin11 | Pin12 | Pin13 | Pin14 | Pin15 | Pin16 |
|--------|-------|-------|-------|-------|-------|-------|-------|-------|
| Signal | COM5  | DO5   | COM6  | DO6   | COM7  | DO7   | COM8  | DO8   |


#### Reference Circuit as follows:


<img src="./6070images/13 Relay Output Circuit 123-1.jpg" alt="alt text" width="220">

### 11. Digital Input Connector
The 8 channel isolated input are equipped with 5000Vrms photo coupler isolator.	Four of the channels form a group and share the same common ground.	The specifications of the isolated input channels are:

- Logical High: 5~24Vdc 
- Logical Low: 0~1.5Vdc
- Input resistance: 1.8KOhms @0.5W 
- Response time: 20µs
- Isolation: 5000Vrms

<img src="./6070images/13 Digital Input Connector.jpg" alt="alt text" width="600">	

#### Pin Assignment of Digital Input:


| **Pin**   | **Pin1** | **Pin2** | **Pin3** | **Pin4** | **Pin5**  |
|-----------|----------|----------|----------|----------|-----------|
| **Signal**| DI1      | DI2      | DI3      | DI4      | ICOM1     |

| **Pin**   | **Pin6** | **Pin7** | **Pin8** | **Pin9** | **Pin10** |
|-----------|----------|----------|----------|----------|-----------|
| **Signal**| DI5      | DI6      | DI7      | DI8      | ICOM2     |


<img src="./6070images/14 DI simple.jpg" alt="alt text" width="300">	

- **DIx: Isolated digital input channels. COM: Common ground.**

### 12. Analog Input Connector

Each of the 5 channels of isolated analog input can be configured for various input ranges. The common features are as follows:

- **Effective Resolution:** 5x24-bit A/D channels  
- **Input Mode(Voltage)s:** 3 channels (differential input)
- **Input Range (Voltage):** ±10VDC  
- **(Optional):** 6 x Single-end channels, call for driver support  
- **Input Mode (Current):** 2 channels (single-end) 
- **Input Current:** 0~20mA 
  

<img src="./6070images/15 AI connector and Jumper settings.jpg" alt="alt text" width="400">

### 13. MiniPCIe Slot

The PAC-6070 includes a miniPCIe slot, accessible by following the direction of the red arrow shown below.

#### Wireless Connectivity Options

The PAC-6070 supports wireless connectivity options. The device is compatible with both WiFi and LTE modules. Certified modules include:

<img src="./6070images/16 miniPCIe 1.jpg" alt="alt text" width="450">

- **SIM7600** (LTE)  
- **RYWDB00** (WiFi + Bluetooth)

For detailed specifications and compatibility information, please refer to the following:  
- [RYWDB00｜REYAX TECHNOLOGY](https://www.reyax.com)  
- [SIM7600X Module｜SIMCom Wireless Solutions](https://www.simcom.com)



### 14. Nano-SIM Card Socket

The PAC-6070 includes an internal Nano-SIM card socket for use with an LTE/4G module. Insert the SIM card into the tray holder and follow the instructions provided.

<img src="./6070images/18 Nano_SIM Card.jpg" alt="alt text" width="650">


### 15. SD Card Socket

The PAC-6070 includes an internal micro-SD card socket for data storage, accessible by following the red arrow direction shown below.

<img src="./6070images/19 SD card.jpg" alt="alt text" width="200">

<img src="./6070images/20 SD Card Jumper Settings-1.jpg" alt="alt text" width="350">