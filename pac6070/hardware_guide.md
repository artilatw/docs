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
- Indicator: READY, LAN, RS-485, DI, Relay-out, A/D
- Dimensions (WxLxH): 121.8 x 204 x 50mm (4.8 x 8 x 1.96in)
- Net Weight: 360g (0.791b)
•	Operating Temperature: 0~70°C (32~158°F)
- Regulation: CE Class A, FCC Class A
- Installation: DIN-rail mounting


The behavior of the reset button depends on how long you press the reset button.

| Press and hold the reset button | Behavior                                          | Network settings after reboot                |
|---------------------------------|--------------------------------------------------|----------------------------------------------|
| < 3 seconds then release        | Re-boot the Matrix-752                           | Retains last user settings                   |
| 3~10 seconds then release       | Reset the network setting (The same setting as factory default) | eth0 IP: addr. by DHCP<br>eth1 IP: 192.168.2.127 |
| > 10 seconds then release       | Restore and back to factory default.<br>User's data may disappear | eth0 IP: addr. by DHCP<br>eth1 IP: 192.168.2.127 |
