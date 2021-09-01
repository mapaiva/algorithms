package list

type SortList struct {
	items []int
}

func (s *SortList) Add(value int) {
	s.items = append(s.items, value)
	s.sort(len(s.items) - 1)
}

func (s *SortList) Median() float64 {
	isEven := len(s.items)%2 == 0
	index := len(s.items) / 2

	if isEven {
		return (float64(s.items[index-1]) + float64(s.items[index])) / 2
	}
	return float64(s.items[index])
}

func (s *SortList) sort(index int) {
	for s.items[previous(index)] > s.items[index] {
		s.swap(previous(index), index)
		index = previous(index)
	}
}

func (s *SortList) swap(index1, index2 int) {
	s.items[index1], s.items[index2] = s.items[index2], s.items[index1]
}

func previous(index int) int {
	if index == 0 {
		return 0
	}
	return index - 1
}
