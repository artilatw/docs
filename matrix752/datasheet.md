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

## H/W Specifications
### CPU / Memory / Flash
- NXP i.MX6ULL Cortex-A7 @ 800MHz
- 512MB LPDDR3 SDRAM
- 16GB eMMC Flash

### Network Interface
- 2x independent 10/100Mbps
- Connector: RJ45 with LED indicators

### USB Interface
- 1x USB 2.0 OTG port
- Connector: micro-USB

### Serial Communication Ports
- 1x RS-485/232 (two-in-one)
- 1x RS-232
- RS-485 Signals: Data+, Data-
- RS-232 Signals: TX, RX, GND
- Connector: Terminal block

### Serial Communication Parameters
- Baud Rate: Up to 921.6Kbps
- Parity: None, Even, Odd, Mark, Space
- Data Bits: 5, 6, 7, 8
- Stop Bits: 1, 1.5, 2
- Flow Control: XON/XOFF, None

### CAN Bus Port
- 1x CAN Bus 2.0 A/B
- Speed: Up to 1Mbps

### Digital Input
- 2x Optical Isolated Inputs
- Isolation: 2500Vrms
- Logic High: 5~24VDC
- Logic Low: 0~1.5VDC

### Digital Output
- 2x Solid State Relay Outputs
- Normal Open
- Contact Rating: 80VDC @ 1.5A

### Native Serial Console
- Interface: RS-232
- Signals: TX, RX, GND
- Baud Rate: 115.2Kbps
- Parameters: 8N1
- Connector: 4-pin box header

### Expansion Slots
- 1x full-size miniPCIe socket
- Supports USB interface
- Supports WiFi, LTE modules
- 1x micro-SIM socket

### SD Card
- 1x microSD socket
- SDXC compliant, up to 2TB

### LED Indicators
- 1x System Ready LED
- 2x Serial Communication LED
- 2x Ethernet Status LED (on RJ45 connector)

### Power Requirements
- Power input range: 9 ~ 48Vdc 
- Typical power consumption: 12Vdc @ 250mA

### Real Time Clock
- Yes
- Backup 1: by onboard super capacitor
- Backup 2: supports external 3.3V battery
- Watchdog output: can be enabled to trigger system reboot

### General
- Operating temperature: 0 ~ 70℃
- Regulations: CE Class A, FCC Class A
- Dimensions: 89 x 112 x 30mm
- Weight: 350g
- Wall mount: yes
- DIN Rail mount: yes (35mm DIN rail mounting kit needs to be ordered separately)

## S/W Specifications
### Linux OS:
- Linux Kernel 5.10.x
- Matrix-752: Yocto Linux 2.7
- Matrix-752U: Ubuntu Linux 20.04 LTS

### User Application Development
- Supports in-system C/C++ compilation
- Toolchain: gcc 9.3.0, glibc 2.3.1
- Can install preferred script languages ( python, nodejs, etc.) and related libraries from software package repository.

### Software Package Repository
- Matrix-752: maintained by Artila
- Matrix-752U: maintained by Ubuntu
- Supports standard apt package management commands (apt install/remove/update/upgrade)

### System Backup and Restore
- Artila provides backup/restore commands to ease the cloning of kernel and file system from a master device to multiple target devices
- Supported media: microSD, USB drive

### Multi-function Reset Button
- Press less than 3 seconds: Reboot the Matrix-752
- Press 3-10 seconds: Reset the network settings (factory default)
- Press more than 10 seconds: Restore and return to factory default. User data are erased.
