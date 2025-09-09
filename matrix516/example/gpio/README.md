# Example GPIO

1. Modify Makefile to change SDK directory

2. Compiler and build  
execute "make" will create executable file

2. Transfer to Matrix-516  
use scp or USB/SD to put the executable file into Matrix-516

3. Execute example  
first change file access permissions:  
chmod +x chip-info  
and then execute:  
./chip-info

4. For more information, please see the references

## References
- [libgpiod](https://git.kernel.org/pub/scm/libs/libgpiod/libgpiod.git/tree/?h=v2.2.x)
- [libgpiod mirror](https://github.com/brgl/libgpiod)