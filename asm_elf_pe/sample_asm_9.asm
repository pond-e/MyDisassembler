	BITS 64
	section .text
	global _start

_start:
	jmp [r9 + rdx * 4]