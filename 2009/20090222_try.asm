  org 0x100
  dd  _main
global	_main
section .text
_main:
	mov ax,0123H
	mov bx,0456H
	add ax,bx
	add ax,ax
  mov ax,4c00H
	int 21H
