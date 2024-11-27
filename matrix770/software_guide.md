# Matrix-770 Software Guide
> Linux-Ready Cortex-A9 Industrial IoT Gateway 

## Trademarks 
The Artila logo is a registered trademark of Artila Inc. All other trademarks or 
registered marks in this manual belong to their respective manufacturers. 

## Disclaimer 
Information in this document is subject to change without notice and does not 
represent a commitment on the part of Artila. 

Artila provides this document as is, without warranty of any kind, either expressed or implied, including, but not limited to, its particular purpose. Artila reserves the right to make improvements and/or changes to this manual, or to the products and/or the programs described in this manual, at any time. 

Information provided in this manual is intended to be accurate and reliable. 
However, Artila assumes no responsibility for its use, or for any infringements on the rights of third parties that may result from its use. 

This product might include unintentional technical or typographical errors. 
Changes are periodically made to the information herein to correct such 
errors, and these changes are incorporated into new editions of the 
publication. 

## 1. Overview
This software guide applies to Artila’s Matrix-77x series Industrial IoT gateway. 

## System Backup/Restore
- Backup system to USB drive
  ```
  ## Usage
  root@matrix77x:~# backup
  Usage: backup device  - backup to device, eg. /dev/sda1
  
  root@matrix77x:~# backup /dev/sda1
  Backup to /dev/sda1, Sure ?(y/n)
  ```

- Restore system from USB drive
  ```
  ## Usage
  root@matrix77x:~# restore
  Usage: restore factory - restore to factoy
         restore device  - restore from device, eg. /dev/sda1
  
  root@matrix77x:~# restore /dev/sda1       
  Restore from /dev/sda1, Sure ?(y/n)
  ```

- Restore to default system
  ```
  root@matrix77x:~# restore factory
  Restore to factory, Sure ?(y/n)
  
  ```