assume cs:codesg

codesg segment
start: mov ax,0123
       mov bx,0456
       add ax,bx
       add ax,ax
       mov ax,4c00H
       int 21
codesg ends
end