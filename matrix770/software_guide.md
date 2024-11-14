# Matrix-770

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