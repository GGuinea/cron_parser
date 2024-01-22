package consts

const (
	ASTERIKS        string = "*"
	STEP_OPERATOR          = "/"
	RANGE_OPERATOR         = "-"
	LISTING_OPRATOR        = ","
)

type Value interface {
	GetAllowedValues() []int
	GetMinValue() int
	GetMaxValue() int
}

type Minutes struct {
}

type Hours struct {
}

type DayOfMonth struct {
}

type Month struct {
}

type DayOfWeek struct {
}

func (m *Minutes) GetAllowedValues() []int {
	res := make([]int, len(allowed_minutes_values))
	copy(res, allowed_minutes_values)
	return res
}

func (m *Minutes) GetMinValue() int {
	return allowed_minutes_values[0]
}

func (m *Minutes) GetMaxValue() int {
	return allowed_minutes_values[len(allowed_minutes_values)-1]
}

func (d *DayOfWeek) GetAllowedValues() []int {
	res := make([]int, len(allowed_day_of_week_values))
	copy(res, allowed_day_of_week_values)
	return res
}

func (d *DayOfWeek) GetMinValue() int {
	return allowed_day_of_week_values[0]
}

func (d *DayOfWeek) GetMaxValue() int {
	return allowed_day_of_week_values[len(allowed_day_of_week_values)-1]
}

func (h *Hours) GetAllowedValues() []int {
	res := make([]int, len(allowed_hours_values))
	copy(res, allowed_hours_values)
	return res
}

func (h *Hours) GetMinValue() int {
	return allowed_hours_values[0]
}

func (h *Hours) GetMaxValue() int {
	return allowed_hours_values[len(allowed_hours_values)-1]
}

func (d *DayOfMonth) GetAllowedValues() []int {
	res := make([]int, len(allowed_day_of_month_values))
	copy(res, allowed_day_of_month_values)
	return res
}

func (d *DayOfMonth) GetMinValue() int {
	return allowed_day_of_month_values[0]
}

func (d *DayOfMonth) GetMaxValue() int {
	return allowed_day_of_month_values[len(allowed_day_of_month_values)-1]
}

func (m *Month) GetAllowedValues() []int {
	res := make([]int, len(allowed_month_values))
	copy(res, allowed_month_values)
	return res
}

func (m *Month) GetMinValue() int {
	return allowed_month_values[0]
}

func (m *Month) GetMaxValue() int {
	return allowed_month_values[len(allowed_month_values)-1]
}

var allowed_minutes_values = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 58, 59}

var allowed_hours_values = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}

var allowed_day_of_month_values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31}

var allowed_month_values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

var allowed_day_of_week_values = []int{0, 2, 3, 4, 5, 6}
