import io_manager
import parser

all_numbers = io_manager.get_line()
substrs = parser.parse(all_numbers, 100)

total = 0
for substr in substrs:
	total += int(substr)

print total

