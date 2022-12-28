package main

import "fmt"

func main() {
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "Aguest",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Println(Q2)
	fmt.Println(summer)

	// getting slices capacity
	fmt.Println(cap(Q2))
	fmt.Println(cap(summer))

	// getting slice and array element address
	fmt.Println(&months[0])
	fmt.Println(&Q2[0])
}
