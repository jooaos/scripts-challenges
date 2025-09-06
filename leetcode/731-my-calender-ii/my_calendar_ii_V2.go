package mycalenderii

import "sort"

type MyCalendarTwoV2 struct {
	bookingsLine map[int]int
	limit        int
}

func ConstructorV2() MyCalendarTwoV2 {
	return MyCalendarTwoV2{
		bookingsLine: map[int]int{},
		limit:        2,
	}
}

/**
Input: [5, 10], [15, 20], [8, 16], [9, 10]
	05: +1
	08: +1
	10: -1
	15: +1
	16: -1
	20: -1
**/

func (m *MyCalendarTwoV2) Book(startTime, endTime int) bool {
	m.bookingsLine[startTime]++
	m.bookingsLine[endTime]--

	keys := make([]int, 0, len(m.bookingsLine))
	for key := range m.bookingsLine {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	active := 0
	for _, t := range keys {
		active += m.bookingsLine[t]
		if active > m.limit {
			m.undo(startTime, endTime)
			return false
		}
	}

	return true
}

func (m *MyCalendarTwoV2) undo(startTime, endTime int) {
	m.bookingsLine[startTime]--
	m.bookingsLine[endTime]++

	if m.bookingsLine[startTime] == 0 {
		delete(m.bookingsLine, startTime)
	}
	if m.bookingsLine[endTime] == 0 {
		delete(m.bookingsLine, endTime)
	}
}
