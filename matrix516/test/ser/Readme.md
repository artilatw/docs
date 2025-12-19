# Matrix-516 Serial Test

## Device Info
### Kernel: 6.12.62

## Test-1
### Test Conditions 
- Baud: 19200
- Data length: 128byte
- Data Header(start byte): 5A A5
- Refresh rate: 10Hz(10 times per second)
- Example data:
  ```
  5A A5 04 4E 7E 03 E8 00 6C 19 0C 0D 5C B0 5B 02 05 00 64 00 00 65 D4 32 11 A7 65 30 A7 65 30 D0 C9 CD 16 CD 16 02 00 01 04 00 66 FB 4E FC 19 02 00 D9 51 BD 02 0A 00 68 00 F6 7D 05 02 02 06 00 69 00 03 FF E5 00 11 2A 02 00 6A 9C 3D 06 1B 12 7F 00 34 02 14 00 6B 05 8B 00 00 02 00 67 1A AC 2C E1 33 41 D5 F8 02 00 00 09 1D 00 00 09 1A 00 09 00 00 01 1A AC 3E A6 33 41 D9 6C 02 0A 20 5B
  ```

### Result
- Test: send 100 times test data from PC
- Raw date log: [serial1-test.log](serial1-test.log)
- Elapsed time: 6405msec

## Test-2
### Test Conditions 
- Baud: 9600
- Data length: 18byte
- Data Header(start byte): 02 5C
- Refresh rate: 50Hz(50 times per second)
- Example data:
  ```
  02 5C 5F 61 40 40 40 42 60 40 40 41 74 4E 40 40 44 03
  ```

### Result
- Test: send 100 times test data from PC
- Raw date log: [serial2-test.log](serial2-test.log)
- Elapsed time: 2976msec