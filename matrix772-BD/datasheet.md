# Matrix-772-BD-2E/6E
### Cortex-A9 based Linux Single Board Computer

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

#### Model: Matrix-722-BD-2E
- No. of ports: two, independent
- Speed: Giga/100M/10M, auto

#### Model: Matrix-722-BD-6E
- WAN port: x1, 100M/10M, auto
- LAN port Hub: x5, 100M/10M, auto

### micro-SD card slot
- Spec.：supports SD 2.0
- Capacity：up to 128GB

### Real-time Clock (RTC)
- Yes
- Backup 1: connector for external button cell battery
- Backup 2: on-board supercapacitor (want spec?)

### USB 2.0 Host
- No. of ports: one
- Connector: Type A

### Status LEDs
- Power: green
- System ready: green
- User LED: yellow

### Native Debug console
- Signals: RS-232 (Tx, Rx)
- Connector: 4-pin wafer header
- Speed: 115,200 baud, 8N1

### miniPCIe expansion slot（want spec? 有可能因空間不足而移除）
- Signals： supports USB
- SIM socket：supports nano-SIM card
- LED：x1, SMD type, green

### Power Input
- Input range: 9Vdc～30Vdc
- Input polarity protection: Yes
- Power consumption: 200mA@12Vdc typical
- Connector: terminal block with fixing screws

### Installation
- Dimensions (WxLxH): 125mm x 85mm
- PCB thickness: 1.6mm
- Mounting holes: 4xM3

### General
- Real-time clock（RTC）: yes
- Watchdog timer: yes (via RTC)
- Operating temperature: 0℃～70℃
- Humidity: 5%～95% (non-condensing)
- Regulation: FCC / CE class A

## Software Features

### Operating System
- Linux kernel: 6.6.x
- Ubuntu core (24.04.1 LTS)
- Supports uPnP protocol

### Toolchain
- GNU C/C++ native compiler, 13.2.0

### Pre-installed packages
- Node.js v22.11.0
- Artila backup/restore utilities
