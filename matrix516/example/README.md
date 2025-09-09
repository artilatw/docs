# Matrix-516 Development Environment Setup and Sample Program Guide

This document provides instructions for setting up a cross-compilation development environment for the Matrix-516 device on Ubuntu 24.04.2 or later. It supports Matrix-516 kernel version 5.12.x and newer, and includes steps to compile and deploy a simple example program to the device.

## 1. SDK Download

Download the Matrix-516 Software Development Kit (SDK) from Artila’s official website:

[arm-artila-linux-gnueabi_sdk-buildroot.tar.gz](https://www.artila.com/download/9G20/Linux/toolchain/arm-artila-linux-gnueabi_sdk-buildroot.tar.gz)

## 2. Development Environment Requirements

This guide assumes the use of Ubuntu 24.04.2 as the development platform and basic familiarity with Linux command-line operations.

### 2.1 Install Required Packages

Run the following command to install essential development tools:

```bash
sudo apt -y install net-tools ssh make binutils
```

Optionally, install the full build toolchain:

```
sudo apt -y install build-essential
```

## 3. SDK Installation and Configuration

### 3.1 Extract the SDK Archive

Unpack the downloaded SDK file:
```
tar zxf arm-artila-linux-gnueabi_sdk-buildroot.tar.gz
```

### 3.2 Set Up Toolchain Environment Variables
Configure the environment variables to include the SDK toolchain:
```
export toolchain=[Your directory]/arm-artila-linux-gnueabi_sdk-buildroot/bin
export PATH=$toolchain:$PATH
```

## 4. Sample Program Compilation and Deployment

### 4.1 Create the Sample Program hello.c
Create a file named hello.c with the following content:
```
#include <stdio.h>

int main(void) {
    printf("Hello, World!\n");
    return 0;
}
```

### 4.2 Compile the Program Using the Cross-Compiler
Use the SDK’s cross-compiler to build the program:
```
arm-artila-linux-gnueabi-gcc hello.c -o hello
```

### 4.3 Transfer the Program to Matrix-516
Use scp to copy the compiled binary to the Matrix-516 device:
```
scp hello root@192.168.2.127:
```

Ensure that the Matrix-516 is accessible via SSH and connected to the same network as your development host.



