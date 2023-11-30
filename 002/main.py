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


class Track(object):
    def __init__(self) -> None:
        self.path = []
    
    def add(self, value: str) -> None:
        self.path.append(value)
        
    def accumulate(self) -> list:
        output = []
        current = None
        count = 0
        
        for index, item in enumerate(self.path):
            if current == item:
                count += 1
            else:
                if current is not None:
                    output.append((current, count))
                current = item
                count = 1
            
            if index == len(self.path) - 1:
                output.append((current, count))
        
        return output


# our state machine
class Machine(object):
    def __init__(self, n: int, targets: tuple) -> None:
        self.limit = n
        self.matrix = [[0 for _ in range(n)] for _ in range(n)]
        self.direction = RIGHT
        self.value = 1
        self.x = n-1
        self.y = 0
        self.start_target, self.stop_target = targets
        self.storage = Track()
        
    def shape(self) -> list:
        return self.matrix
        
    def output(self) -> list:
        return self.storage.accumulate()
    
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
        flag = False
        
        while self.value <= limit:
            if self.value == self.start_target:
                flag = True
            
            self.matrix[self.x][self.y] = self.value
            self.value += 1
            
            x, y = self.direction["callback"](self.x, self.y)
            if self.rules(x, y):
                self.direction = self.change(self.direction)
            
            self.x, self.y = self.direction["callback"](self.x, self.y)
            
            if flag == True:
                self.storage.add(self.direction["value"])
            
            if self.value == self.stop_target:
                flag = False


if __name__ == "__main__":
    n0 = int(input("Enter (N): "))
    t1 = int(input("Enter 1st target: "))
    t2 = int(input("Enter 2nd target: "))
    
    m = Machine(n0, (t1, t2))
    m.start()
    
    for row in m.shape():
        print(row)
    
    output = m.output()
    print(', '.join([ f'({a}, {b})' for a, b in output ]))
