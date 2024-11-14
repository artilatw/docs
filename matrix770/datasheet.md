# Matrix-770
### Ubuntu-core based Cortex-A9 IIoT Gateway

<img alt="Matrix-770" src="https://www.artila.com/images/news/Matrix-770.jpg" width=480>

## H/W Specifications

### CPU
- Artila UM20A910 Cortex-A9
- Cores: single-core
- Frequency：800MHz

### DRAM
- Spec.：DDR3-1600
- Size：1GB (up to 2GB)

### eMMC
- Size: 16GB

### Ethernet
- No. of ports: two, independent
- Speed: Giga/100M/10M, auto

### miniPCIe expansion slot
- Signals： supports USB
- SIM socket：supports nano-SIM card
- LED：x1, SMD type, green

### micro-SD card slot
- Spec.：supports SD 2.0
- Capacity：up to 128GB

### Real-time Clock (RTC)
- Yes
- Backup 1: on-board supercapacitor
- Backup 2: connector for external button cell battery

### USB 2.0 Host
- No. of ports: two
- Connector: Type A

### Status LEDs
- Power: green
- System ready: green


### Native Debug console
- Signals: RS-232 (Tx, Rx)
- Connector: 4-pin wafer header
- Speed: 115,200 baud, 8N1

### USB Serial Console
- Connector: Type C (**NO** PD support!)
- Speed: 115,200 baud, 8N1

### Serial Ports
- RS-485 ports: four
- RS-485 auto-direction control: Yes
- RS-232 ports: one
- RS-232 signals: Tx/Rx/RTS/CTS

### Serial Port Parameters
- Baud rate: up to 115,200 bps
- Data bits: 5, 6, 7, 8
- Parity: None, Even, Odd
- Stop bits: 1, 1.5, 2
- Flow control: RTS/CTS, XON/XOFF, None

### Power Input
- Input range: 9Vdc～48Vdc
- Input polarity protection: Yes
- Power consumption: 200mA@12Vdc typical
- Connector: terminal block with fixing screws

### Installation
- Dimensions (WxLxH): 140mm x 145mm x 25mm
- Wall-mountable
- Supports DIN rail mounti（optional）

### General
- Buzzer: yes
- Watchdog timer: yes (via RTC)
- Operating temperature: 0℃～70℃
- Humidity: 5%～95% (non-condensing)
- Regulation: FCC / CE class A

## Software Features

### Operating System
- Linux kernel: 6.6.x
- Ubuntu core (22.04.1 LTS)
- Supports uPnP protocol

### Toolchain
- GNU C/C++ native compiler, 13.2.0

### Pre-installed packages
- Node.js V22.10.0
- Artila backup/restore utilities
