package mycalenderii

type MyCalendarTwoV1 struct {
	bookings       [][]int
	doubleBookings [][]int
}

func ConstructorV1() MyCalendarTwoV1 {
	return MyCalendarTwoV1{
		bookings:       [][]int{},
		doubleBookings: [][]int{},
	}
}

/**
Input: [5, 10], [15, 20], [8, 16]
	|-------|
			 |-------|
|-------|
	 |-------|

			|-------|
			 			|-------|
|-------|
!(En <= So OR Eo <= Sn)
En > So AND Eo > Sn
**/

func (m *MyCalendarTwoV1) Book(startTime int, endTime int) bool {
	for _, i := range m.doubleBookings {
		if endTime > i[0] && i[1] > startTime {
			return false
		}
	}

	for _, i := range m.bookings {
		if endTime > i[0] && i[1] > startTime {
			m.doubleBookings = append(m.doubleBookings, []int{
				max(startTime, i[0]),
				min(endTime, i[1]),
			})
		}
	}

	m.bookings = append(m.bookings, []int{startTime, endTime})
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
