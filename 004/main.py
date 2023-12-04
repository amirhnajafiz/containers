matrix = [
  [1, 1, 4, 3, 6, 6],
  [2, 2, 3, 4, 5, 5]
]

# constant variables
N = 2
M = 6
LIMIT = N+M+1

row = 0
col = 0
stack = []

while row + col < LIMIT:
  current = matrix[row][col]
  flag = False
  
  for item in stack:
    if item == current:
      flag = True
      break

  if flag:
    index = row+1
    while index < N:
      pos = matrix[index][col]
      flag = True
      for item in stack:
        if item == pos:
          flag = False
          break

      if flag:
        tmp = matrix[row][col]
        matrix[row][col] = matrix[index][col]
        matrix[index][col] = tmp
        break

    index += 1
  
  stack.push(matrix[row][col])
  
  if col == N:
    stack = []
    row += 1
    col = 0


print(matrix)
