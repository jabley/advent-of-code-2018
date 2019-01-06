# bind register 2 as the instruction pointer
#ip 2

# load
     0 addi 2 16 2  r2 := r2 + 16 # goto 17 (main)

OUTER LOOP:
for r4 in 1..r2 
     1 seti 1 2 4   r4 := 1

INNER LOOP:
for r1 in 1..r2
     2 seti 1 8 1   r1 := 1
INNER LOOP:
if r3 = r4 * r1 goto 7 else goto 6
     3 mulr 4 1 5   r5 := r4 * r1
     4 eqrr 5 3 5   r5 := 1 if (r5 == r3) else 0
     5 addr 5 2 2   r2 := r5 + r2 # if r3 = r4 * r1 goto 7
     6 addi 2 1 2   r2 := r2 + 1 # goto 8
r0 += r4
     7 addr 4 0 0   r0 := r4 + r0

     8 addi 1 1 1   r1 := r1 + 1
     9 gtrr 1 3 5   r5 := 1 if (r1 > r3) else 0
    10 addr 2 5 2   r2 := r2 + r5 # if r1 > r3 goto 12 
    11 seti 2 6 2   r2 := 2 # goto 3

    12 addi 4 1 4   r4 := r4 + 1
    13 gtrr 4 3 5   r5 := 1 if (r4 > r3) else 0
    14 addr 5 2 2   r2 := r5 + r2 # if r4 > r3 goto 16
    15 seti 1 2 2   r2 := 1 # goto 2

PROGRAM END
    16 mulr 2 2 2   r2 := r2 * r2 (pc = 16) # goto 256

MAIN
r3 = (2 * 2) * 19 * 11 < 836 >
    17 addi 3 2 3   r3 := r3 + 2 (2)
    18 mulr 3 3 3   r3 := r3 * r3 (4)
    19 mulr 2 3 3   r3 := r2 * r3 (76)
    20 muli 3 11 3  r3 := r3 * 11 (836)
r5 = 2 * 22 + 8 < 52 > 
    21 addi 5 2 5   r5 := r5 + 2 (2)
    22 mulr 5 2 5   r5 := r5 * r2 (44)
    23 addi 5 8 5   r5 := r5 + 8 (52)
r3 = r3 + r5 < 888 >
    24 addr 3 5 3   r3 := r3 + r5 = 836 + 52 = 888
    25 addr 2 0 2   r2 := r2 + r0 = 25 + 0 = 25 # goto 26 + r0
    26 seti 0 4 2   r2 := 0 # goto 1 (OUTER LOOP)

MOAR MAIN: # if r0 == 1 at program start time
r5 = ((27 * 28) + 29) * 30 * 14 * 32 < 10550400 >
    27 setr 2 5 5   r5 := r2 (27)       # r5 = 27
    28 mulr 5 2 5   r5 := r5 * r2 (756)  # r5 = 27*28
    29 addr 2 5 5   r5 := r2 + r5 (785)  # r5 = 27*28+29
    30 mulr 2 5 5   r5 := r2 * r5 (23550)  # r5 = (27*28+29)*30
    31 muli 5 14 5  r5 := r5 * 14 (329700)  # r5 = (27*28+29)*30*14
    32 mulr 5 2 5   r5 := r5 * r2 (10550400)  # r5 = (27*28+29)*30*14*32
r3 = 888 + 10550400 < 10551288 >
    33 addr 3 5 3   r3 := r3 + r5
    34 seti 0 8 0   r0 := 0 # clear r0 from the 1 that it initially contained
    35 seti 0 5 2   r2 := 0 # goto 1
