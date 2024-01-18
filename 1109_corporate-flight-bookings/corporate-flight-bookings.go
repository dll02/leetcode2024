package _109_corporate_flight_bookings

func corpFlightBookings(bookings [][]int, n int) []int {
	arr := make([]int, n+1)
	for _, book := range bookings {
		arr[book[0]-1] += book[2]
		arr[book[1]] -= book[2]
	}
	ans := make([]int, n)
	ans[0] = arr[0]
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] + arr[i]
	}
	return ans
}
