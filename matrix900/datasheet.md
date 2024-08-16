# Matrix-900
Linux-ready Cortex-A72(ARM v8)

## Features
- Broadcom BCM2711 quad-core Cortex-A72(ARM v8) 64-bit 1.5GHz
- Support 1GB, 2GB, 4GB or 8GB LPDDR4
- Support 8GB, 16GB or 32GB eMMC
- Support Raspberry Pi OS
- 1 x full size HDMI 2.0
- [1 x Gigabit and 1 x 10/100Mbps Ethernet](#network-interface)
- 2 x USB 2.0 host ports
- [4 x RS-485 serial port](#serial-ports)
- 1 x Audio Line out/in
- 1 x MIPI CSI camera interface
- 1 x MIPI DSI display interface
- 1 x mPCIe socket
- 1 x M.2 B key socket(3042/3052)
- 1 x microSD socket
  
## H/W Specifications
### SoM (System On Module)
- Raspberry Pi Compute Module 4
- Processor: BCM2711, quad core Cortex-A72 (ARM v8) 64-bit SoC @ 1.5GHz
- DDR: 1/2/4/8GB
- eMMC: 8/16/32GB
- By design, the Matrix-900 does not support booting from an SD card

### Network Interface
- 1 x Gigabit Ethernet
- 1 x 10/100Mbps Ethernet
- Connector: RJ45 (with LED indicator)

### USB Interface
- 2 x USB 2.0 compliant host ports
- Supports 480Mbps high speed mode 
- Connector: Type A
- 
### Display Port
- 1 x HDMI port
- HDMI 2.0 compliant, upto 4K@60fps
- Connector: Type A

### Serial Ports
- 4 x RS-485 serial ports
- Connector: Terminal block
- Signal: Data+, Data-
- Auto direction control: yes
- Baudrate: upto 12Mbps
- Data bits: 7, 8
- Stop bits: 1, 2
- Parity: Odd/Even/Mark/Space/None

### M.2 Expansion slot
- Interface: PCIe (1-lane, Gen2)
- Use case:
  - AI inference accelerator
  - 5G communication module
  - SSD module
- Module dimension: supports upto 3052

### miniPCIe Expansion slot
- Interface: USB 2.0
- Supports 480Mbps high speed mode
- Use case:
  - LTE/4G communication module
  - WiFi module
 
### Console/Debug Port
- RS-232 serial port (inside the box)
- Signal: Tx/Rx/GND
- Connector: 4-pin wafer
- Communication parameters: 115,200bps, N81

### SD Card
- 1 x microSD socket
- SD 3.0 compliant
- Storage capacity: SDXC

### Real Time Clock
- Yes, on-board
- Backup: external battery
- Connector: 2-pin 1.25mm wafer box

### General
- Operation temperature: -20~70&deg;C 

### Order Information