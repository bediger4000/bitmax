# Daily Coding Problem: Problem #787 [Medium]

This problem was asked by Facebook.

Given an integer n,
find the next biggest integer with the same number of 1-bits on.
For example,
given the number 6 (0110 in binary), return 9 (1001).

## Build and Run

```sh
$  go build algorithm1.go bitcount.go
$  ./algorithm1 6
6 has 2 bits on
9
$ go build algorithm2.go bitsbytes.go                                              # /home/bediger/src/go/src/bitmax
$ ./algorithm2 127                                                                 # /home/bediger/src/go/src/bitmax
127 has 7 bits on
127     0000000000000000000000000000000000000000000000000000000001111111        7
191     0000000000000000000000000000000000000000000000000000000010111111        7
$ go build compare.go bitsbytes.go
./compare 2097150                                                                # /home/bediger/src/go/src/bitmax
2097150 has 20 bits on
2097150 0000000000000000000000000000000000000000000111111111111111111110        20
2621439 0000000000000000000000000000000000000000001001111111111111111111        20
2621439 0000000000000000000000000000000000000000001001111111111111111111        20
```

## Analysis

What size is this integer? 32 bits? 16? 64?
I'm doing this in Go, so I'm going to use `int` or `uint` types.
That's 64 bits.

I found 2 algorithms.

---

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

---

### Algorithm 2

1. Find the lowest 1-bit with a 0-bit to it's left
in the input value.
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

I don't have a convincing proof that this algorithm calculates the
next-biggest-number-same-1-bit-count.
It definitely finds a number bigger than the input
with the same 1-bit count.

---

## Interview Analysis

It took me a while to figure out Algorithm 2.
I think an interviewer can't expect a job candidate
to figure that out on a whiteboard.
Interviewers could probably expect that in a take-home problem.

The actual programming is minimal for Algorithm 1,
just a single loop, plus a population count function.
This isn't a "medium" solution.

The programming for Algorithm 2 is fiddly,
getting the indexes and end-conditions correct on the for-loops is difficult.
I'm not sure an inteviewer can expect a candidate to get all of those
starting and ending conditions correct.
They require experimentation,
when done on-the-job would lead to test cases.
Whiteboarding would probably not elicit that from the candidate.

It's possible to have cases, like 18446744073709551615 (64 1-bits),
that are difficult to get into the functions because string-to-integer
conversion fails, or gives negative numbers.
Getting this correct (I don't think I got it) will depend on knowledge
of how library code does string-to-integer conversion.

Barring the conceptual leap in discovering Algorithm 2,
this problem deserves the "medium" rating.
The interviewer should guard against irrationally high expectations.
