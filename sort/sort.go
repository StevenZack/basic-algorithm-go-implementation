package algorithm

func QuickSort(areas []int) {
	if len(areas) < 2 {
		return
	}
	i := 1
	head, tail := 0, len(areas)-1
	for head < tail {
		if areas[i] > areas[head] {
			areas[i], areas[tail] = areas[tail], areas[i]
			tail--
		} else {
			areas[i], areas[head] = areas[head], areas[i]
			i++
			head++
		}
	}
	sort(areas[:head])
	sort(areas[head+1:])
}
