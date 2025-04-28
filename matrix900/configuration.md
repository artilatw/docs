# Artila Matrix-900(CM4) Configuration Guide

## config.txt
  - PATH: /boot/firmware/config.txt

### Enable Debug Console
  ```
  [cm4]
  enable_uart=1
  ```
 
### Enable Real Time Clock(RV-8803) support
  ```
  [cm4]
  dtoverlay=i2c-rtc,rv8803,i2c_csi_dsi
  ```

### Enable microSD support
  ```
  [cm4]
  dtoverlay=sdio,poll_once=off
  ```

### Enable Audio CODEC support
  ```
  [cm4]
  dtoverlay=tlv320aic3x-soundcard-csidsi,i2c_csi_dsi
  ```
- Copy tlv320aic3x-soundcard-csidsi.dtbo to /boot/firmware/overlays
  ```
  sudo cp tlv320aic3x-soundcard-csidsi.dtbo /boot/firmware/overlays/
  ```  

- Test  
  Set JP7 to 3,5 & 4,6 for Headphone Out, which is on the Matrix-900 board.
  ```
  
  ## list devices
  aplay -l
  
  card 0: vc4hdmi0 [vc4-hdmi-0], device 0: MAI PCM i2s-hifi-0 [MAI PCM i2s-hifi-0]
    Subdevices: 1/1
    Subdevice #0: subdevice #0
  card 1: vc4hdmi1 [vc4-hdmi-1], device 0: MAI PCM i2s-hifi-0 [MAI PCM i2s-hifi-0]
    Subdevices: 1/1
    Subdevice #0: subdevice #0
  card 2: tlv320aic3104so [tlv320aic3104-soundcard], device 0: fe203000.i2s-tlv320aic3x-hifi tlv320aic3x-hifi-0 [fe203000.i2s-tlv320aic3x-hifi tlv320aic3x-hifi-0]
    Subdevices: 0/1
    Subdevice #0: subdevice #0
  
  ## play audio on terminal
  cvlc --play-and-exit sample.mp3
  
  ## Set default audio output
  ##   Run Raspberry Pi Software Configuration Tool
  ##   System Options -> Audio
  sudo raspi-config
  
  ## set volume
  amixer set Master 50%
  ```  

### Set GPIO state
  ```
  [cm4]
  gpio=3=op,dl
  gpio=2,11,13,16=op,dh
  ```
  
### Disable Bluetooth
  ```
  [cm4]
  dtoverlay=disable-bt
  ```

### Disable Wi-Fi
  ```
  [cm4]
  dtoverlay=disable-wifi
  ```
## References
- [Raspberry Pi Documentation - Configuration](https://www.raspberrypi.com/documentation/computers/configuration.html)
- Original boot files
  - [config.txt](config.txt)
  - [cmdline.txt](cmdline.txt)
  
- [樹莓派與藍芽低功耗(BLE)](https://hackmd.io/@ShenTengTu/S1uEIKbKE)