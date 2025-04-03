# MyDisassembler
IA-32e向けのELF64, PE32+対応逆アセンブラの部分実装です

PrefixやModR/M, SIBの処理は出来てるはず

table.goのmapに対して機械的に命令を追加する工程がまだ残っています

デモのために逆アセンブルする対象のセクションを1つのみに指定しています、コード全体を逆アセンブルしたい場合は拡張してください

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
