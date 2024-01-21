package consts

const (
	ASTERIKS        string = "*"
	STEP_OPERATOR          = "/"
	RANGE_OPERATOR         = "-"
	LISTING_OPRATOR        = ","
)

var allowed_minutes_values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 58, 59}

var allowed_hours_values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

var allowed_day_of_month_values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31}

var allowed_month_values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

var allowed_day_of_week_values = []int{0, 2, 3, 4, 5, 6}

func GetAllowedMinutesValues() []int {
	res := make([]int, len(allowed_minutes_values))
	copy(res, allowed_minutes_values)
	return res
}

func GetAllowedHoursValues() []int {
	res := make([]int, len(allowed_hours_values))
	copy(res, allowed_hours_values)
	return res
}

func GetAllowedDayOfMonthValues() []int {
	res := make([]int, len(allowed_day_of_month_values))
	copy(res, allowed_day_of_month_values)
	return res
}

func GetAllowedMonthValues() []int {
	res := make([]int, len(allowed_month_values))
	copy(res, allowed_month_values)
	return res
}

func GetAllowedDaysValues() []int {
	res := make([]int, len(allowed_day_of_week_values))
	copy(res, allowed_day_of_week_values)
	return res
}

func GetMinMinutesValue() int {
	return allowed_minutes_values[0]
}

func GetMaxMinutesValue() int {
	return allowed_minutes_values[len(allowed_minutes_values)-1]
}

func GetMinHoursValue() int {
	return allowed_hours_values[0]
}

func GetMaxHoursValue() int {
	return allowed_hours_values[len(allowed_hours_values)-1]
}

func GetMinDayOfMonthValue() int {
	return allowed_day_of_month_values[0]
}

func GetMaxDayOfMonthValue() int {
	return allowed_day_of_month_values[len(allowed_day_of_month_values)-1]
}

func GetMinMonthValue() int {
	return allowed_month_values[0]
}

func GetMaxMonthValue() int {
	return allowed_month_values[len(allowed_month_values)-1]
}

func GetMinDayOfWeekValue() int {
	return allowed_day_of_week_values[0]
}

func GetMaxDayOfWeekValue() int {
	return allowed_day_of_week_values[len(allowed_day_of_week_values)-1]
}
