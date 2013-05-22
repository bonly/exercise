;; 在NASM的語法中, 分號後面的部份為註解
;;
GLOBAL main ; 宣告 main 為全域符號(symbol)
EXTERN printf,scanf ; 宣告 printf 和 scanf 為外部符號

SECTION .data ; 從以下開始為data節區
pfmt DB '%d', 0x0A, 0
sfmt DB '%d', 0

SECTION .bss ; 從以下開始為bss節區
buf RESD 1

SECTION .text ; 從以下開始為text節區
main: ; 主函式
push ebp
mov ebp, esp
push buf
push sfmt
call scanf ; 呼叫 scanf 函式
add esp, 8 ; 清除堆疊中的參數
mov eax, DWORD [buf]
push eax
push pfmt
call printf ; 呼叫 printf 函式
add esp, 8
mov eax, 0
leave
ret ; 主函式結束 