import os
import prime

x = 1024124214811
print prime.isPrime(x)


def do():
    with open('input.txt') as input_file:
        testCases = input_file.readlines()
    numTestCases = testCases[0]
    testCases.remove(numTestCases)
    output_filename = 'output.txt'
    if os.path.exists(output_filename):
        os.remove(output_filename)
    tidy_numbers = []
    for test_case in testCases:
        n = int(test_case)
        closest_tidy = get_closest_tidy(n)
        tidy_numbers.append(closest_tidy)

    with open(output_filename, 'a') as output_file:
        for i in xrange(0, len(tidy_numbers)):
            output_file.write('Case #' + str(i + 1) + ': ' +
                              str(tidy_numbers[i]) + '\n')
