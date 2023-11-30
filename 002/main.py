# creating our states
UP = {
    'value': "U",
    'callback': lambda x, y: (x-1, y)
}

DOWN = {
    'value': "D",
    'callback': lambda x, y: (x+1, y)
}

LEFT = {
    'value': "L",
    'callback': lambda x, y: (x, y-1)
}

RIGHT = {
    'value': "R",
    'callback': lambda x, y: (x, y+1)
}


# our state machine
class Machine(object):
    def __init__(self, n: int) -> None:
        self.limit = n
        self.matrix = [[0 for i in range(n)] for j in range(n)]
        self.direction = RIGHT
        self.value = 1
        self.x = n-1
        self.y = 0
        
    def output(self) -> list:
        return self.matrix
    
    def rules(self, x, y) -> bool:
        return x == self.limit or x == -1 or y == self.limit or y == -1 or self.matrix[x][y] != 0
    
    def change(self, direction: dict) -> dict:
        value = direction["value"]
        
        if value == UP["value"]:
            return LEFT
        elif value == LEFT["value"]:
            return DOWN
        elif value == DOWN["value"]:
            return RIGHT
        elif value == RIGHT["value"]:
            return UP
        else:
            return None
    
    def start(self) -> None:
        limit = self.limit**2
        while self.value <= limit:
            self.matrix[self.x][self.y] = self.value
            self.value += 1
            
            x, y = self.direction["callback"](self.x, self.y)
            if self.rules(x, y):
                self.direction = self.change(self.direction)
            
            self.x, self.y = self.direction["callback"](self.x, self.y)


if __name__ == "__main__":
    m = Machine(int(input("Enter (N): ")))
    m.start()
    
    output = m.output()
    for row in output:
        print(row)
