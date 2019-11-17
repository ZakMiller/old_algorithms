def get_inverse(string):
    new_string = ""
    for digit in string:
        if digit == '0':
            new_string = new_string + '1'
        else:
            new_string = new_string + '0'
    return new_string


def get_thue_morse(length):
    binary_string = "0"
    while len(binary_string) < length:
        inverse_string = get_inverse(binary_string)
        binary_string = binary_string + inverse_string

    return binary_string


def get_numbers_present(binary_string):
    numbers_present = dict()
    possible_new_numbers = [0]

    for window_length in xrange(len(binary_string), 2, -1):
        for i in xrange(0, len(binary_string) - window_length + 1):
            possible_new_numbers.append(
                int(binary_string[i: i + window_length], 2))

        for num in possible_new_numbers:
            if num not in numbers_present:
                numbers_present[num] = 1
    return numbers_present

thue_sequence = get_thue_morse(1000)
dictionary = get_numbers_present(thue_sequence)
keys = list(dictionary.keys())
keys.sort()
print keys[:100]
