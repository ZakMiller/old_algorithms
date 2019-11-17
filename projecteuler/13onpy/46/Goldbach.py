import prime
prime_finder = prime.prime('bil')
primes = list(prime_finder.prime_dictionary.keys())
primes.sort()


def can_be_written(x):
    square_base = 1
    while 2 * (square_base ** 2) + 2 <= x:
        for prime in primes:
            if 2 * (square_base ** 2) + prime > x:
                break
            elif 2 * (square_base ** 2) + prime == x:
                return True

        square_base += 1
    return False


def get_goldberg_number():
    x = 33
    while True:
        if x % 2 == 1 and prime_finder.is_prime(x) == False:
            if can_be_written(x) == False:
                return x
        x += 1

goldberg_number = get_goldberg_number()
print goldberg_number
