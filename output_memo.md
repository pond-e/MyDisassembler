```
$ go run . sample_asm_3_win64.exe 200
[7 11]
parse operand loop
operand: 7
hoge4
eax
parse operand loop
operand: b
hoge3
modrm.modByte: 0
operand: b
[ebx]
mov eax, [ebx]
2025/01/31 19:57:24 Unknown combination of prefix, mnemonic and opcodeByte: (none, add, )
exit status 1
```

PEとELFの二つに対応
```
$ go run . sample_asm_3
.text
mov eax, [ebx]
$ go run . sample_asm_3_win64.exe
mov eax, [ebx]
```

複数行にも対応！
```
$ go run . sample_asm_8
.text
add eax, 0x1
mov ebx, eax
```

fix peread
```
$ go run . sample_asm_8_win64.exe 
sectionHeaderOffset: 0x1b0
Section 1: VirtualSize: 0x5, PointerToRawData: 0x200
Section 2: VirtualSize: 0x88, PointerToRawData: 0x400
add eax, 0x1
mov ebx, eax
```