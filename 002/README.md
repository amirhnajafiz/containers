# Question 002

Given a number as ```N``` and two target places in a __2D__ array that is like
a snail shell maze. For example, if $N = 5$ then the output will be like this:

$$
\begin{array}{cc}
  13 & 12 & 11 & 10 & 9 \\
  14 & 23 & 22 & 21 & 8 \\
  15 & 24 & 25 & 20 & 7 \\
  16 & 17 & 18 & 19 & 6 \\
  1 & 2 & 3 & 4 & 5
\end{array}
$$

We need to find the moves that should be taken to go from first target to second one.
For example in above matrix, for reaching house number 9 from house number 4, we
need to go one right and four ups.

## solution

First we need to create a state machine for our direction changings.

- If you were going ```RIGHT```, then you need to go ```UP```
- If you were going ```UP```, then you need to go ```LEFT```
- If you were going ```LEFT```, then you need to go ```DOWN```
- If you were going ```DOWN```, then you need to go ```RIGHT```

Now we are going to create a $N \times N$ matrix with 0 value for each house.
Also, we need is to set some rules for direction changing. Changing in direction is needed
when:

- $x_i$ or $y_i$ goes out of range $(0,N)$
- the value of the next house is greater than 0.

Now we are going to set action functions for each of our directions.

- ```RIGHT``` : $(x_i, y_i) \to (x_i, y_i + 1)$
- ```UP``` : $(x_i, y_i) \to (x_i - 1, y_i)$
- ```LEFT``` : $(x_i, y_i) \to (x_i, y_i - 1)$
- ```DOWN``` : $(x_i, y_i) \to (x_i + 1, y_i)$

After that we are going to start from $(N-1,0)$ ( as $(x_i, y_i)$ ) house with initial ```RIGHT```
direction. If we hit one of the first target we are going to store our moves in a stack until
we reach the second target.
