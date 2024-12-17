
Program: 2,4,1,1,7,5,1,5,4,2,5,5,0,3,3,0

```
 0 BST 4 : A % 8    > B
 2 BXL 1 : B ^ 1    > B
 4 CDV 5 : A / 2**B > C
 6 BXL 5 : B ^ 5    > B
 8 BXC 2 : B ^ C    > B
10 OUT 5 : B % 8    > STDOUT
12 ADV 3 : A / 8    > A
14 JNZ 0 : rinse repeat
```

- 12 (ADV 3) is the only instruction writing to A. Diving A by 8.
- So the program is run 8log(x) times.
