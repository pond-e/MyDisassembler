## elf64
$ nasm -f elf64 -o sample/obj/sample_ch15_1.o sample/src/sample_ch15_1.asm
$ ld -s -o sample/bin/sample_ch15_1 sample/obj/sample_ch15_1.o
-s はシンボル情報が消えるから使わなくてもいい？

## win64
$ nasm -f win64 -o sample/obj/sample_ch15_1.o sample/src/sample_ch15_1.asm

sample_3_asm_win64.exe entry point
00000200: 678b 0300 0000 0000 0000 0000 0000 0000  g...............

```
C:\Program Files\Microsoft Visual Studio\2022\Community>cd C:\Users\harut\Downloads

C:\Users\harut\Downloads>link sample_asm_3_win64.o /subsystem:console /entry:_start
Microsoft (R) Incremental Linker Version 14.40.33811.0
Copyright (C) Microsoft Corporation.  All rights reserved.
```