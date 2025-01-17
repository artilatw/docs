# System Backup and Restore Howto

## Backup master system
Backup (a.k.a clone) process will pack the Linux kernel and file system into a single image file.

Follow the steps below to backup the master system.

1. Prepare a USB drive with at least 16GB capacity.
2. Format the USB drive to FAT32.
3. Insert the USB drive into the master Matrix752U. Assume the USB drive is `/dev/sda1`.
4. Execute the following command to start the backup process.

```
$ backup /dev/sda1
Backup to /dev/sda1, Sure?(y/n)
y 
```

5. The backup process will take 20-30 minutes to complete.
6. The ready LED will be blinking during the backup process.
7. The master Matrix752U will automatically reboot after the backup process is complete.
8. After bootup, remove the USB drive and keep it safe.

## Restore from backup

1. Insert the USB drive into the target Matrix752U. Assume the USB drive is `/dev/sda1`.
2. Execute the following command to start the restore process.

```
$ restore /dev/sda1
Restore from /dev/sda1, Sure?(y/n)
y 
```

3. The restore process will take 20-30 minutes to complete.
4. The ready LED will be blinking during the restore process.
5. The target Matrix752U will automatically reboot after the restore process is complete.
6. After bootup, remove the USB drive and keep it safe.
