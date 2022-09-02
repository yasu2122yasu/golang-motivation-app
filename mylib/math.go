package

func Average(s []int)int{
	total := 0
	for _, i := range s{
		total += 1
	}
	return int(total/len(s))
}