import math


MIN_POSITIVE = 0.00000000001


# create base functions
radius = lambda x, y : x**2 + y**2
degree = lambda x, y : math.atan2(y, x)
factor = lambda r, d : r * math.sqrt(abs(d))


def convert_to_polar(sets: list) -> list:
  # create our factors
  array = []
  for item in sets:
    x, y = item
    r = radius(x, y)
    d = degree(x, y)
    d = d if d != 0 else MIN_POSITIVE
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
  
  # get input
  coordinates = [int(x) for x in input("(x, y) > ").split(",")]
    
  # convert to polar
  p = convert_to_polar(sets)
  
  # convert input to polar
  pc = convert_to_polar([(coordinates[0], coordinates[1])])[0]
  
  print(p, pc)
  
  # calculate items
  items = []
  
  for item in p:
    items.append({
      "set": item["set"],
      "delta": abs(item["factor"] - pc["factor"])
    })
    
  print(sorted(items, key=lambda d: d["delta"])[0])
