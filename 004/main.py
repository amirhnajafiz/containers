matrix = [
  [1, 1, 4, 3, 6, 6],
  [2, 2, 3, 4, 6, 5],
  [1, 2, 3, 4, 5, 5]
]


# constant variables
N = len(matrix)
M = len(matrix[0])
LIMIT = N+M


if __name__ == "__main__":
  row = 0
  col = 0
  stack = []


  while row < N and col < M:
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
    
    stack.append(matrix[row][col])
    
    col += 1
    
    if col == M:
      stack = []
      row += 1
      col = 0


  print(matrix)
