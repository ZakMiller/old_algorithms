def get_num_ways_make_change(coins, sum):
	ways = 0
	coins = [value for value in coins if value <= sum]

	if len(coins) == 1 and sum % coins[0] == 0:
		return ways+1

	for coin in coins:
		remaining = sum - coin
		if remaining == 0:
			ways += 1
		else:
			ways += get_num_ways_make_change([value for value in coins if value <= coin], remaining)
	return ways

print get_num_ways_make_change([1, 2, 5, 10, 20, 50, 100, 200], 200)