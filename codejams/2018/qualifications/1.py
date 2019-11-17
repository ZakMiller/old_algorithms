import sys


def calculate_damage(instructions):
    totalDamage = 0
    multiplier = 1
    for c in instructions:
        if c == 'C':
            multiplier *= 2
        else:
            totalDamage += multiplier

    return totalDamage


def reduce(instructions):
    to_be_flipped = instructions.rfind("CS")
    return instructions[0:to_be_flipped] + "SC" + instructions[to_be_flipped + 2:]


def problem_one(problem_pieces):

    health = int(problem_pieces[0])
    instructions = problem_pieces[1].strip()
    swaps = 0

    total_instructions = len(instructions)
    charge_instructions = instructions.count('C')
    shoot_instructions = total_instructions - charge_instructions

    min_damage = shoot_instructions

    if min_damage > health:
        return "IMPOSSIBLE"

    current_damage = calculate_damage(instructions)

    while current_damage > health:
        instructions = reduce(instructions)
        current_damage = calculate_damage(instructions)
        swaps += 1

    return str(swaps)


problem_number = 1
infile = sys.stdin
next(infile)
for line in infile:
    problem_pieces = line.split(" ")
    print "Case #" + str(problem_number) + ": " + problem_one(problem_pieces)
    problem_number += 1
