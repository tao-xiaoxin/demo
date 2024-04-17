package uploads

import (
	"fmt"
	"time"
)

func Add(a int, b int) int {
	return a + b
}
func Mul(a int, b int) int {
	return a * b
}
func calculateCost(daysUsed int) float64 {
	const handlingFee = 10.0
	const officialPrice = 700.0
	const daysInYear = 365.0

	cost := handlingFee + ((officialPrice / daysInYear) * float64(daysUsed))
	return cost
}
func calculateDays(date1 string, date2 string) (int, error) {
	layout := "2006-01-02"
	t1, err1 := time.Parse(layout, date1)
	t2, err2 := time.Parse(layout, date2)

	if err1 != nil || err2 != nil {
		return 0, fmt.Errorf("invalid date format, use YYYY-MM-DD")
	}

	duration := t2.Sub(t1)
	days := int(duration.Hours() / 24)

	return days, nil
}

func main() {

	days, err := calculateDays("2023-12-13", "2024-04-17")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The difference between the two dates is %d days\n", days)
	}
	cost := calculateCost(days)
	fmt.Printf("The cost for %d days is %.2f\n", days, cost)
	//fmt.Println(0.13 * float64(days))
	//fmt.Println(48 - 27)

}
