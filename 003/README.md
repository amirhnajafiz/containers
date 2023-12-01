# Question 003

Given a set of coordinates as $(x_i, y_i)$ on a __2D__ map. Each time we set on of the existing coordinates
as our source. After that, we want to find the closest points to that source. Our output should be a list
of coordinates sorted by their distance from the source coordinate.

## solution

Let's set $(0,0)$ point as the center of our map. Then we are going to create a Polar coordinate system
from the given coordinates. For each entry, we have (find more [details](https://en.wikipedia.org/wiki/Polar_coordinate_system)):

$$
radius = \sqrt{x^2+y^2}
$$

$$
degree = atan2(y,x)
$$

Now we are going to set a new factor for each of our entities (if _degree_ value is 0 then we round it to 1):

$$
factor = radius * degree^2
$$

Finally, every time that we set a new source, we can sort our list based on the subtraction of these factors.
