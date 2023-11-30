# Question 001

Give an array of ints in size of ```N+1``` field with numbers in range ```(0,N)```, however there is duplicate number
in the array. How can we find that number?

## solution

The summation of numbers in range ```(0,N)``` is $N(N+1)/2$. Therefore, all we need to do is to sum every element
in the array and calculate the subtract of it from $N(N+1)/2$.

$$
f(N) = N(N+1)/2
$$

$$
output = SUM(array) - f(N)
$$
