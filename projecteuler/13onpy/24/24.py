import itertools
permutations = list(map("".join, itertools.permutations('0123456789')))
permutations.sort()
print permutations[1000000-1]