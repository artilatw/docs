# Matrix-752 Datasheet

## Introduction
Matrix-752 is a Linux-ready computing platform with a 800MHz Cortex-A7 processor, 512MB SDRAM, and 16GB eMMC. It is designed for industrial IoT applications and supports a wide range of communication interfaces and peripherals.

## Features
- Linux Kernel 5.10.x
- Supports Yocto Linux or Ubuntu Linux
- NXP i.MX6ULL Cortex-A7 @ 800MHz
- 512MB LPDDR3
- 16GB eMMC Flash
- 2x Ethernet ports
- 1x USB OTG port
- 1x RS-485/232
- 1x RS-232
- 1x CAN bus port
- 2x optical isolated inputs
- 2x solid state relay outputs
- 1x microSD socket
- 1x miniPCIe socket


## Specifications (Hardware)
### Processor
>NXP i.MX6ULL Cortex-A7 @ 800MHz

### SDRAM 
>512MB LPDDR3

### Flash memory
>16G eMMC 

### Linux Kernel
>Version: 5.10.x (or later)

### Network Interface
>2x independent 10/100Mbps<br>
Connector: RJ45 with LED indicators

### USB Interface
>1x USB 2.0OTG port<br>
Connector: micro-USB

### Serial Ports
>1x RS-485/232 software selectable<br>
1x RS-232<br>
RS-485 Signals: Data+, Data-<br>
RS-232 Signals: TX, RX, GND<br>
Connector: Terminal block<br>

### Communication parameters
>Baud Rate: Up to 921.6Kbps<br>
Parity: None, Even, Odd, Mark, Space<br>
Data Bits: 5, 6, 7, 8<br>
Stop Bits: 1, 1.5, 2<br>
Flow Control: XON/XOFF, None<br>

### CAN Port
>1x CAN Bus 2.0 A/B<br>
Speed: Up to 1Mbps<br>

### Digital Input
>2x Optical Isolated Inputs<br>
Isolation: 2500Vrms<br>
Logical High: 5~24VDC<br>
Logical Low: 0~1.5VDC<br>

### Digital Output
  - 2x Optical Isolated Inputs
  - Isolation: 2500Vrms
  - Logical High: 5~24VDC
  - Logical Low: 0~1.5VDC
**Digital Output:**  
  - 2x Solid State Relay Outputs
  - Normal Open
  - Contact Rating: 80VDC @ 1.5A
**Serial Console:**
  - Type: RS-232
  - Signals: TX, RX, GND
  - Baud Rate: 115.2Kbps
  - Parameters: 8N1
  - Connector: 4-pin box header
**Expansion Slots:**  
  - 1x full-size miniPCIe socket, supports USB interface
  - Supports WiFi, LTE modules
  - 1x micro-SIM socket
**SD Card:**
  - 1x microSD socket
  - SDXC compliant, up to 2TB
**Power Input:**
  - 9 ~ 48Vdc 
**Power Consumption:**
  - 12Vdc @ 250mA typical
**Watchdog Timer:**
  - Yes
**RTC:**
  - Yes
  - Backup: by supercap
**Temperature Range:**
  - -20 ~ 85℃
**Regulations:**
  - CE Class A
  - FCC Class A
**Installation:**
  - Wall or optional 35mm rail mount
**Dimensions:**
  - 89 x 112 x 30mm
**Weight:**
  - 350g