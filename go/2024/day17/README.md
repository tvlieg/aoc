
Program: 2,4,1,1,7,5,1,5,4,2,5,5,0,3,3,0

```
 0 BST 4 : A % 8    > B       | remainder of A (0-7) goes to B            |
 2 BXL 1 : B ^ 1    > B       | flip bit 1 for B                          |
 4 CDV 5 : A / 2**B > C       | remove last B bits from A and write to C  |
 6 BXL 5 : B ^ 5    > B       | flip bit 1 & 3 for B                      |
 8 BXC 2 : B ^ C    > B       | flip C bits for B                         |
10 OUT 5 : B % 8    > STDOUT  | 0                                         |
12 ADV 3 : A / 8    > A       | remove last 3 bits from A                 |
14 JNZ 0 : rinse repeat       | 
```





