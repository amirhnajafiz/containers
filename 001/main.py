# creating f(N)
f = lambda n : n * (n+1) / 2

# getting the input list
array = [ int(x) for x in input().split(",") ]

# print the output
print(sum(array) - f(len(array) - 1) )
