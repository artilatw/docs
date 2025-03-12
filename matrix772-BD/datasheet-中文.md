# HIVE-M05 規格書

## 摘要
HIVE-M05 是一款基於弘訊 TX7芯片 (Cortex-A9 架構) 設計的小型物聯網網關，內建 Ubuntu Linux 作業系統，方便客戶進行二次開發。

之前 HIVE-M05 叫做 TPD-M05，航航於2025/03/07 通知改名。HIVE-M05 的硬件，部署 TPD 軟件後，是要取代 iD201 (OPC-UA) 的。

> HIVE-M05 本身具備兩個網口 LAN1，LAN2。LAN1 接口是直接從 TX7 芯片引出的；LAN2 接口則是接到板上的網路交換器芯片（Realtek RTL8211F）後再引出。

> HIVE-M05 的 PCB 上已經預留孔位，可以加裝一個 4口的網路接頭，提供額外的四個網口，LAN3 ~ LAN6，其中 LAN2～LAN6 是掛在同一個網路交換器芯片上。

## 硬件規格

### CPU
- TX7: UM20A910（Cortex-A9）
- Cores: single-core，單核心
- Frequency：800MHz

### DRAM
- Spec.：DDR3-1600
- Size：1GB (up to 2GB)

### eMMC
- Size: 16GB

### Ethernet

#### Model: HIVE-M05 
- No. of ports: two, independent
- Speed: 100M/10M, auto

#### Model: HIVE-M05-HUB
- WAN port: x1, 100M/10M, auto
- LAN switch: 5 ports, 100M/10M, auto
- Switch IC：Realtek RTL8211F

### micro-SD card slot
- Spec.：supports SD 2.0
- Capacity：up to 128GB

### Real-time Clock (RTC)
- Yes
- Backup : connector for external button cell battery

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

### miniPCIe expansion slot- Signals： supports USB
- SIM socket：supports nano-SIM card
- LED：x1, SMD type, green

### Power Input
- Input range: 9Vdc～30Vdc
- Input polarity protection: Yes
- Power consumption: 300mA@12Vdc typical
- Connector: terminal block with fixing screws

### Installation
- Dimensions (WxLxH): 125mm x 85mm
- PCB thickness: 1.6mm
- Mounting holes: 4x M3

### General
- Real-time clock（RTC）: yes
- Watchdog timer: yes (via RTC)
- Operating temperature: 0℃～70℃
- Humidity: 5%～95% (non-condensing)
- Regulation: FCC / CE class A

## 軟件特色

### Operating System
- Linux kernel: 6.6.x
- Ubuntu core (24.04.1 LTS)
- Supports uPnP protocol

### Toolchain
- GNU C/C++ native compiler, 13.2.0

### Pre-installed packages
- Node.js v22.11.0
- Artila backup/restore utilities
