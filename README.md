# Daily Coding Problem: Problem #787 [Medium]

This problem was asked by Facebook.

Given an integer n,
find the next biggest integer with the same number of 1-bits on.
For example,
given the number 6 (0110 in binary), return 9 (1001).

## Analysis

What size is this integer? 32 bits? 16? 64?
I'm doing this in Go, so I'm going to use `int` or `uint` types.
That's 64 bits.

I found 2 algorithms.

### Algorithm 1

1. Find number of 1-bits in input integer
2. Count up from input integer
3. For each number counted up, find number of 1-bits
4. If that number of 1-bits is the same as in the input integer,
that's the answer.

This runs in O(n) time, where n is the value of the input integer.
If you have an input like 8 or 16 or 268435456,
it counts from n to 2\*n, doing n increments,
and finding the number of bits in n integers.
This later part makes the constant part of the run-time large.

### Algorithm 2

1. Find the lowest 1-bit with a 0-bit to it's left.
2. Put a zero in the position of the 1-bit,
and a one in the position of the 0-bit.
3. If the least significant bit is zero:
    1. Find the 1-bit with the greatest significance
    below the 1-bit found in (1).
    2. Zero that position.
    3. Set the least significant bit to 1.

This is complicated, but runs in O(n) time,
where n is the position of the highest 1-bit with a 0-bit
in the next highest significant place.
It's quite a bit faster than Algorithm 1.
