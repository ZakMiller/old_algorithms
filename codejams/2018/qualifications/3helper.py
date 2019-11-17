from random import randint


def run_scenario():
    numbers = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
    count = 0
    while len(numbers) > 0:
        random_selection = randint(0, 9)
        count += 1
        if random_selection in numbers:
            numbers.remove(random_selection)
    return count


total = 0
for i in xrange(0, 10000):
    total += run_scenario()

average = total / 10000

print average
