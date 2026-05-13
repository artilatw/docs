# System Backup and Restore using USB Type-C port

The backup process packs the Linux kernel and filesystem into a single image file on a USB drive.

> [!CAUTION]
> Backup and restore operations are only supported between Matrix-800 units with **identical Linux kernel versions**. Before restoring, verify both the source and target devices report the same kernel version using `uname -r`. Restoring an image across different kernel versions may result in an unbootable or unstable system.

## Backup

1. Prepare a USB drive with at least **16 GB** capacity, formatted as **FAT32**.

2. Insert the USB drive into the Matrix-800. Use an adapter if the USB Drive doesn't have any USB Type-C head.

3. Execute the following command to start the backup:

```
    $ backup /dev/sda1
    Backup to /dev/sda1, Sure?(y/n)
    y
```

4. The backup process takes approximately 10–15 minutes. The **Ready LED** blinks during the process.

5. The Matrix-800 automatically reboots once the backup is complete.

6. After the device has booted, remove the USB drive and store it in a safe location.


## Restore

1. **Verify that the target device have the same Linux kernel version as the source**, check `uname -r` on both units before proceeding.

2. Insert the USB drive containing a previously created backup image into the target Matrix-800.

3. Execute the following command to start the restore:

```
    $ restore /dev/sda1
    Restore from /dev/sda1, Sure?(y/n)
    y
```

4. The restore process takes approximately 10–15 minutes. The **Ready LED** blinks during the process.

5. The Matrix-800 automatically reboots once the restore is complete.

6. After the device has booted, remove the USB drive and store it in a safe location.