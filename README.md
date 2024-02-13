# Multithreaded Matrix Rotation

Program that performs a matrix rotation using multithreading.

The rotation works by evaluating the matrix as a set of squares with one
surrounding the other such that the inner square is the smallest and the outer
square is the largest. For each of the squares a thread is created that will
create additional threads for performing the rotation of sets of four values
equidistant from each other.

Consider the following matrix:

```text
1 2 3
4 5 6
7 8 9
```

The matrix is made up of two squares. The inner square is a 1x1 matrix with the
single value 5 and the outer square is the remaining numbers. The outer square
thread will spawn two threads. One will rotate the corner values *{1, 3, 9, 7}*
and the other will rotate *{2, 6, 8, 4}*. For each square the number of rotation
sets is equal to one less than the length of one side. In other words, for a
square of length three, only two rotations are needed.

For another example, consider the following matrix:

```text
A B C D E F
T a b c d G
S l 0 1 e H
R k 3 2 f I
Q j i h g J
P O N M L K
```

Here, imagine three nested squares. One Go routine will spawn per square. For
each square of length *n*, *n-1* threads are spawned to handle the rotation of
four equidistant points. Example rotation sets include *{A, F, K, P}*, *{C, H,
M, R}*, *{c, f, i, l}*, and *{0, 1, 2, 3}*.
