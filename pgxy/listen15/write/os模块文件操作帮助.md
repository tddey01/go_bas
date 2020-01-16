os.OpenFile(“output.dat”, os.O_WRONLY|os.O_CREATE, 0666)

第二个参数:文件打开模式:

1. os.O_WRONLY:只写
2. os.O_CREATE:创建文件
3. os.O_RDONLY:只读
4. os.O_RDWR:读写
5. os.O_TRUNC:清空
6. os. O_APPEND:追加

第三个参数:权限控制:

1. r ——> 004
2. w——> 002
3. x——> 001