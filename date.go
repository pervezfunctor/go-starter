package main

type Month int8

const (
	January Month = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

type Day int8
type Year int

type WeekDay int8

const (
	Sunday WeekDay = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func IsLeap(year Year) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}

func DaysInMonth(month Month, year Year) Day {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case February:
		if IsLeap(year) {
			return 29
		}
		return 28

	default:
		panic("Invalid month")
	}
}

func IsValid(month Month, day Day, year Year) bool {
	return day >= 1 && day <= DaysInMonth(month, year)
}
