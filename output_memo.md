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

```
$ go run . sample_asm_3
.text
mov eax, [ebx]
$ go run . sample_asm_3_win64.exe
mov eax, [ebx]
```