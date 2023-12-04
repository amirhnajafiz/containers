# Question 004

We were given an $(n,m)$ matrix. This matrix is filled with numbers of $(1,m)$, and every number
is repeated _n_ times. However, the matrix is out of order. Meaning that some rows might have duplicate
numbers in them. Our goal is to sort this matrix to have every number in range of $(1,m)$ in each of
the matrix rows. But, we can only swap numbers that are in a same column.

## example

In a $(2,2)$ matrix like this:

$$
\begin{array}{cc}
  1 & 1 \\
  2 & 2
\end{array}
$$

We need to convert it to:

$$
\begin{array}{cc}
  1 & 2 \\
  2 & 1
\end{array}
$$

## solution

We are going to use greedy dynamic programming to solve this problem. Imagine we are on the _i_ th row. Lets assume that
previous rows are sorted correctly. We are going to start from the first house and push the rows numbers in a stack while
we are iterating that row. In each house we are going to check to see if we have a duplicate number. If there was a duplicate
number, we are going to go down in its column until we reach to a number that we never had it in our stack. Then, we are
going to replace them in that column.
