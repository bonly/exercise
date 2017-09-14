/* OFFSET宏定义可取得指定结构体某成员在结构体内部的偏移 */
#define OFFSET(st, field)     (size_t)&(((st*)0)->field)
typedef struct{
    char  a;
    short b;
    char  c;
    int   d;
    char  e[3];
}T_Test;

int main(void){  
    printf("Size = %d\n  a-%d, b-%d, c-%d, d-%d\n  e[0]-%d, e[1]-%d, e[2]-%d\n",
           sizeof(T_Test), OFFSET(T_Test, a), OFFSET(T_Test, b),
           OFFSET(T_Test, c), OFFSET(T_Test, d), OFFSET(T_Test, e[0]),
           OFFSET(T_Test, e[1]),OFFSET(T_Test, e[2]));
    return 0;
}