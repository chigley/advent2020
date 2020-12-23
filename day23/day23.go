package day23

func Part1(digits []int) string {
	cups := NewCups(digits)
	for i := 0; i < 100; i++ {
		cups.Move()
	}
	return cups.ResultString()
}
