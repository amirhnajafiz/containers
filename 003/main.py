import math


# create base functions
radius = lambda x, y : x**2 + y**2
degree = lambda x, y : math.atan2(y, x)
factor = lambda r, d : r * (d**2)


def convert_to_polar(sets: list) -> list:
  # create our factors
  array = []
  for item in sets:
    x, y = item
    r = radius(x, y)
    d = degree(x, y)
    d = d if d != 0 else 1
    array.append({
      "set": item,
      "factor": factor(r, d),
    })
    
  return array

if __name__ == "__main__":
  # input coordinates
  sets = [
    (1,1),(10,-1),(64,1),(5,0),(12,2),
    (-1,-1),(6,-6),(12,11),(100,-5),(50,20),
  ]

  # convert to polar
  p = convert_to_polar(sets)
