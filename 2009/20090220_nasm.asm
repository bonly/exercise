;; ��NASM���Z����, ��̖����Ĳ��ݞ��]��
;;
GLOBAL main ; ���� main ��ȫ���̖(symbol)
EXTERN printf,scanf ; ���� printf �� scanf ���ⲿ��̖

SECTION .data ; �������_ʼ��data���^
pfmt DB '%d', 0x0A, 0
sfmt DB '%d', 0

SECTION .bss ; �������_ʼ��bss���^
buf RESD 1

SECTION .text ; �������_ʼ��text���^
main: ; ����ʽ
push ebp
mov ebp, esp
push buf
push sfmt
call scanf ; ���� scanf ��ʽ
add esp, 8 ; ����ѯB�еą���
mov eax, DWORD [buf]
push eax
push pfmt
call printf ; ���� printf ��ʽ
add esp, 8
mov eax, 0
leave
ret ; ����ʽ�Y�� 