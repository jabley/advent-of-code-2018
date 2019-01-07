# bind register 4 as the instruction pointer
#ip 4

NUMERIC_CHECK:
     0 seti 123 0 1         r1 := 123
     1 bani 1 456 1         r1 := r1 & 456
     2 eqri 1 72 1          r1 := 1 if (r1 == 72) else 0
     3 addr 1 4 4           r4 := r1 + r4 # goto 5 if (r1 & 456) == 72)
     4 seti 0 0 4           r4 := 0 # goto 1 !!INFINITE LOOP IF bani doesn't work correctly!!

MAIN:
     5 seti 0 3 1           r1 := 0
     6 bori 1 65536 5       r5 := r1 | 65536 # r5=65536 first time through
     7 seti 8586263 3 1     r1 := 8586263
     8 bani 5 255 2         r2 := r5 & 255
     9 addr 1 2 1           r1 := r1 + r2
    10 bani 1 16777215 1    r1 := r1 & 16777215
    11 muli 1 65899 1       r1 := r1 * 65899
    12 bani 1 16777215 1    r1 := r1 & 16777215
    13 gtir 256 5 2         r2 := 1 if (256 > r5) else 0
    14 addr 2 4 4           r4 := r2 + r4 # goto 16 if (256 > r5)
    15 addi 4 1 4           r4 := r4 + 1 # goto 17
    16 seti 27 8 4          r4 := 27 # goto 28

LOOP:
while ((r2+1)*256) < r5:
  r2 := r2 + 1
    17 seti 0 1 2           r2 := 0
    18 addi 2 1 3           r3 := r2 + 1
    19 muli 3 256 3         r3 := r3 * 256
    20 gtrr 3 5 3           r3 := 1 if (r3 > r5) else 0
    21 addr 3 4 4           r4 := r3 + r4 # goto 23 if (r3 > r5)
    22 addi 4 1 4           r4 := r4 + 1 # goto 24
    23 seti 25 8 4          r4 := 25 # goto 26
    24 addi 2 1 2           r2 := r2 + 1
    25 seti 17 7 4          r4 := 17 # goto 18
    26 setr 2 0 5           r5 := r2
    27 seti 7 8 4           r4 := 7 # goto 8

TEST_FOR_TERMINATION:
    28 eqrr 1 0 2           r2 := 1 if (r1 == r0) else 0
    29 addr 2 4 4           r4 := r2 + r4 # goto 31 if (r1 == r0) !!PROGRAM END!!
    30 seti 5 4 4           r4 := 5 # goto 6
