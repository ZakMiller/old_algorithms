import sys


def all_full(grid, x, y):
    if grid[x - 1][y - 1] == '0' \
            or grid[x][y - 1] == '0' \
            or grid[x - 1][y] == '0' \
            or grid[x][y] == '0' \
            or grid[x + 1][y] == '0' \
            or grid[x][y + 1] == '0' \
            or grid[x + 1][y + 1] == '0' \
            or grid[x + 1][y - 1] == '0' \
            or grid[x - 1][y + 1] == '0':
        return False
    return True


def problem_three(area):
    grid = [['0' for i in range(10)] for j in range(10)]
    target_area = ((area + 8) / 9) * 9
    x = 2
    y = 2
    num_vertices = target_area / 9

    while True:
        print str(x) + " " + str(y)
        sys.stdout.flush()
        pair = raw_input().split(" ")
        placed_x = pair[0]
        placed_y = pair[1]
        if placed_x == '0' and placed_y == '0':
            return
        elif placed_x == '-1' and placed_y == '-1':
            return
        else:
            grid[int(placed_x)][int(placed_y)] = 'X'
            if all_full(grid, x, y):
                y += 3


test_cases = input()
line = input()
while line is not None:
    problem_three(line)
    test_cases -= 1

    if test_cases == 0:
        break
    line = input()
