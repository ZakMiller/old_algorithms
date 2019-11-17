def factor(value):
    factors = []
    for i in range(1, int(value**0.5) + 1):
        if value % i == 0:
            factors.append((i, value / i))
    return factors

possibleNumber = 646

while possibleNumber < 650:
    print factors(possibleNumber)
    possibleNumber += 1
