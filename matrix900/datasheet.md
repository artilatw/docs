- [Matrix-900 Data Sheet](#matrix-900-data-sheet)
  - [Model Description](#model-description)
  - [H/W Specifications](#hw-specifications)
    - [SoM (Raspberry Pi Compute Module 4)](#som-raspberry-pi-compute-module-4)
    - [Dual Ethernet Interface](#dual-ethernet-interface)
    - [Dual USB 2.0 Interface](#dual-usb-20-interface)
    - [HDMI Display Port](#hdmi-display-port)
    - [Stereo Audio Output](#stereo-audio-output)
    - [Four RS-485 Serial Ports](#four-rs-485-serial-ports)
    - [M.2 Expansion slot](#m2-expansion-slot)
    - [miniPCIe Expansion slot](#minipcie-expansion-slot)
    - [SD Card slot](#sd-card-slot)
    - [Real Time Clock](#real-time-clock)
    - [Console/Debug Port](#consoledebug-port)
    - [General](#general)
  - [Order Information](#order-information)

# Matrix-900 Data Sheet
## Model Description
Raspberry Pi Compute Module 4 powered IIoT Gateway <br>Supporting AI inference accelerator (via M.2 slot)

## H/W Specifications

### SoM (Raspberry Pi Compute Module 4)
- Raspberry Pi Compute Module 4
- Processor: BCM2711, quad core Cortex-A72 (ARM v8) 64-bit SoC @ 1.5GHz
- DDR: 1/2/4/8GB
- eMMC: 8/16/32GB
- By design, the Matrix-900 does not support booting from an SD card

### Dual Ethernet Interface
- 1 x Gigabit Ethernet
- 1 x 10/100Mbps Ethernet
- Connector: RJ45 (with LED indicator)

### Dual USB 2.0 Interface
- 2 x USB 2.0 compliant host ports
- Supports 480Mbps high speed mode 
- Connector: Type A
- 
### HDMI Display Port
- 1 x HDMI port
- HDMI 2.0 compliant, upto 4K@60fps
- Connector: Type A

### Stereo Audio Output
- Codec: TI TLV320
- Connector: 3.5mm audio jack (4-pole)
- Mated plug: 3.5mm TRRS connector
- Standard: CTIA(MGRL) or OMPT(GMRL)? 

### Four RS-485 Serial Ports
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
- Key type: B Key
- Use case:
  - AI inference accelerator
  - 5G communication module
  - SSD module
- Nano-SIM socket: yes, on-board (SIM-M)
- Module dimension: supports upto 30x52mm
- Fixing screws: at 42mm and 52mm position

### miniPCIe Expansion slot
- Interface: USB 2.0
- Supports 480Mbps high speed mode
- Use case:
  - LTE/4G communication module
  - WiFi module
- Nano-SIM socket: yes, on-board (SIM-E)
- Module dimension: supports full-size module
- Fixing screws: for full-size modules

### SD Card slot
- 1 x microSD socket
- SD 3.0 compliant
- Storage capacity: SDXC

### Real Time Clock
- Yes, on-board
- Backup: external battery
- Connector: 2-pin 1.25mm wafer box

### Console/Debug Port
- RS-232 serial port (inside the box)
- Signal: Tx/Rx/GND
- Connector: 4-pin wafer
- Communication parameters: 115,200bps, N81

### General
- Operation temperature: -20~70&deg;C 

## Order Information