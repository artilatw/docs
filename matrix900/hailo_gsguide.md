# Hailo Getting Started Guide
The [Hailo-8 M.2 Module](https://hailo.ai/products/ai-accelerators/hailo-8-m2-ai-acceleration-module/) is an AI accelerator module for AI applications. The Artila Matrix-900 provides support for the Hailo-8 Module in the M.2 B-key form factor.

## Configuration
First, enable M.2 power to modify configure file /boot/firmware/config.txt

```text
#Enable M.2 power
gpio=16=op,dl
```

Then power-on the Matrix-900, and using command ***lspci*** to verify the Hailo module information.

```bash
guest@matrix900:~$ lspci | grep Co-processor
01:00.0 Co-processor: Hailo Technologies Ltd. Hailo-8 AI Processor (rev 01)
```

The ***dmesg*** command shows boot-up messages.

```bash
guest@matrix900:~$ dmesg|grep -i hailo
[    6.477186] hailo: Init module. driver version 4.17.0
[    6.479609] hailo 0000:01:00.0: Probing on: 1e60:2864...
[    6.479640] hailo 0000:01:00.0: Probing: Allocate memory for device extension, 11600
[    6.479695] hailo 0000:01:00.0: enabling device (0000 -> 0002)
[    6.479716] hailo 0000:01:00.0: Probing: Device enabled
[    6.479791] hailo 0000:01:00.0: Probing: mapped bar 0 - 00000000ecc10176 16384
[    6.479819] hailo 0000:01:00.0: Probing: mapped bar 2 - 000000007d528a76 4096
[    6.479838] hailo 0000:01:00.0: Probing: mapped bar 4 - 00000000167f8dcc 16384
[    6.479855] hailo 0000:01:00.0: Probing: Setting max_desc_page_size to 4096, (page_size=4096)
[    6.479888] hailo 0000:01:00.0: Probing: Enabled 64 bit dma
[    6.479900] hailo 0000:01:00.0: Probing: Using userspace allocated vdma buffers
[    6.479916] hailo 0000:01:00.0: Disabling ASPM L0s
[    6.479947] hailo 0000:01:00.0: Successfully disabled ASPM L0s
[    6.487008] hailo 0000:01:00.0: Firmware file not found (/lib/firmware/hailo/hailo8_fw.bin), please upload the firmware manually
[    6.514548] hailo 0000:01:00.0: Probing: Added board 1e60-2864, /dev/hailo0
```

The default hailo_pci driver version in 6.6.31+rpt-rpi-v8.
```
guest@matrix900:~$ modinfo hailo_pci
filename:       /lib/modules/6.6.31+rpt-rpi-v8/kernel/drivers/media/pci/hailo/hailo_pci.ko.xz
version:        4.17.0
license:        GPL v2
description:    Hailo PCIe driver
author:         Hailo Technologies Ltd.
srcversion:     C57A34FC4F4C6985BE6AF31
alias:          pci:v00001E60d000043A2sv*sd*bc*sc*i*
alias:          pci:v00001E60d000045C4sv*sd*bc*sc*i*
alias:          pci:v00001E60d00002864sv*sd*bc*sc*i*
depends:
intree:         Y
name:           hailo_pci
vermagic:       6.6.31+rpt-rpi-v8 SMP preempt mod_unload modversions aarch64
parm:           o_dbg:int
parm:           no_power_mode:Disables automatic D0->D3 PCIe transactions (invbool)
parm:           force_allocation_from_driver:Determines whether to force buffer allocation from driver or userspace (int)
parm:           force_desc_page_size:Determines the maximum DMA descriptor page size (must be a power of 2) (int)
```

### Update driver to v4.20

```
sudo apt -y install hailo-dkms hailort
```
Restart system, then check driver information.



```bash
guest@matrix900:~$ modinfo hailo_pci
filename:       /lib/modules/6.6.31+rpt-rpi-v8/updates/dkms/hailo_pci.ko.xz
version:        4.20.0
license:        GPL v2
description:    Hailo PCIe driver
author:         Hailo Technologies Ltd.
import_ns:      DMA_BUF
srcversion:     BD5E76800A08DB052331F58
alias:          pci:v00001E60d000043A2sv*sd*bc*sc*i*
alias:          pci:v00001E60d000045C4sv*sd*bc*sc*i*
alias:          pci:v00001E60d00002864sv*sd*bc*sc*i*
depends:
name:           hailo_pci
vermagic:       6.6.31+rpt-rpi-v8 SMP preempt mod_unload modversions aarch64
sig_id:         PKCS#7
signer:         DKMS module signing key
sig_key:        40:8B:66:4E:BC:FF:14:EE:69:A1:B9:19:78:D5:F8:FF:38:3F:21:3A
sig_hashalgo:   sha512
signature:      A3:47:1B:88:94:69:CC:E9:31:E7:42:42:ED:E1:4C:7B:2E:61:F5:C8:
                BC:9D:B2:AC:09:BD:D4:4A:5C:04:14:B4:A1:1D:75:6D:6F:69:B9:89:
                D2:AA:E6:AB:7E:1C:87:86:26:F3:40:E3:64:3A:7E:6A:28:46:61:4F:
                10:F8:D2:16:0C:64:C5:E7:F3:5D:91:FB:7B:5D:22:1A:5F:D3:0C:34:
                34:F2:41:80:C8:9B:AC:3D:30:1B:6C:B7:B6:4C:34:45:16:F0:8F:E1:
                47:EC:3F:8F:2D:81:2D:A4:7D:F6:1C:D2:FE:50:7D:83:50:28:54:3E:
                8E:80:E5:8C:14:75:7D:ED:4F:3C:A6:E9:56:C4:2E:CB:49:AE:3F:0B:
                0E:ED:EB:E8:4C:94:65:E9:3A:ED:A6:C9:02:4C:4B:A5:0C:2C:E9:9E:
                F4:A1:89:EC:74:E0:3B:A1:95:7C:FD:6E:0A:31:FE:1D:2F:15:41:4D:
                5B:94:75:D0:CF:E4:78:A3:E3:2F:D7:CB:ED:5F:2B:6D:9C:7E:65:BB:
                35:3C:19:6D:6E:63:35:01:10:F1:FD:F9:9C:AF:86:67:50:6A:13:BB:
                85:52:96:60:7D:79:FE:38:9B:8C:32:A8:B6:F8:83:40:C0:B5:9F:02:
                2E:5D:FA:52:2C:B2:8E:95:93:71:24:84:11:EC:CB:A1
parm:           o_dbg:int
parm:           no_power_mode:Disables automatic D0->D3 PCIe transactions (invbool)
parm:           force_allocation_from_driver:Determines whether to force buffer allocation from driver or userspace (int)
parm:           force_desc_page_size:Determines the maximum DMA descriptor page size (must be a power of 2) (int)
parm:           force_hailo10h_legacy_mode:Forces work with Hailo10h in legacy mode(relevant for emulators) (bool)
parm:           force_boot_linux_from_eemc:Boot the linux image from eemc (Requires special Image) (bool)
parm:           support_soft_reset:enables driver reload to reload a new firmware as well (bool)
```
Using ***hailortcli scan*** to validate that the device is identified.

```bash
guest@matrix900:~$ hailortcli scan
Hailo Devices:
[-] Device: 0000:01:00.0
```

Identifying Device’s Serial Number
```bash
guest@matrix900:~$ hailortcli fw-control identify
Executing on device: 0000:01:00.0
Identifying board
Control Protocol Version: 2
Firmware Version: 4.20.0 (release,app,extended context switch buffer)
Logger Version: 0
Board Name: Hailo-8
Device Architecture: HAILO8
Serial Number: HLLWMBA224900240
Part Number: HM218B1C2LAE
Product Name: HAILO-8 AI ACC M.2 B+M KEY MODULE EXT TEMP
```

## Python API
Install HailoRT Python API

```bash
sudo apt install python3-hailort
```


## References
- https://hailo.ai/developer-zone/documentation/hailort-v4-24-0/
- [Python Standalone Examples](https://github.com/hailo-ai/hailo-apps/tree/main/hailo_apps/python/standalone_apps)