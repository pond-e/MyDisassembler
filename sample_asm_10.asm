	BITS 64
	section .text
	global _start

_start:
	add [r9 + rdx * 4], eax