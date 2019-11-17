import datetime

def get_all_days_between(start_date, end_date):
	date = start_date
	while date <= end_date:
		yield date
		date = date + datetime.timedelta(days=1)


start = datetime.date(1901, 1, 1)
end = datetime.date(2000, 12, 31)

num_first_of_month_and_sundays = 0
dates = get_all_days_between(start, end)
for date in dates:
	if date.weekday() == 6 and date.day == 1:
		num_first_of_month_and_sundays+=1
print num_first_of_month_and_sundays