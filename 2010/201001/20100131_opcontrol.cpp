
int fast_multiply(int x,  int y) 
{
    return x * y;
}
int slow_multiply(int x, int y) 
{
    int i, j, z;
    for (i = 0, z = 0; i < x; i++) 
        z = z + y;
    return z;
}
int main()
{
    int i,j;
    int x,y;
    for (i = 0; i < 200; i ++) {
        for (j = 0; j < 30 ; j++) {
            x = fast_multiply(i, j);
        y = slow_multiply(i, j);
    }
    }
    return 0;
}
/*
 * http://www.ibm.com/developerworks/cn/linux/l-pow-oprofile/index.html
 */
