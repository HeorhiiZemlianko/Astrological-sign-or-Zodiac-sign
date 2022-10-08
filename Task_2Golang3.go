package main

import (
	"fmt"
)

func main() {
	var day, month int8
	fmt.Print("Введите день: ")
	fmt.Scan(&day)

	fmt.Print("Введите месяц: ")
	fmt.Scan(&month)

	var astro_sign string
	// checks month and date within the
	// valid range of a specified zodiac
	if day >= 1 && day <= 30 {
		if month == 12 {
			if day < 22 {
				astro_sign = "Стрелец"
			} else {
				astro_sign = "Козерог"
			}
			fmt.Println("Знак зодиака: ", astro_sign)
		} else if month == 1 {
			if day < 20 {
				astro_sign = "Козерог"
			} else {
				astro_sign = "Водолей"
			}
			fmt.Println("Знак зодиака: ", astro_sign)
		} else if month == 2 {
			if day < 19 {
				astro_sign = "Водолей"
			} else {
				astro_sign = "Рыбы"
			}
			fmt.Println("Знак зодиака: ", astro_sign)
		} else if month == 3 {
			if day < 21 {
				astro_sign = "Рыбы"
			} else {
				astro_sign = "Овен"
			}
			fmt.Println("Знак зодиака: ", astro_sign)
		} else if month == 4 {
			if day < 20 {
				astro_sign = "Овен"
			} else {
				astro_sign = "Телец"
			}
			fmt.Println("Знак зодиака: ", astro_sign)
		} else if month == 5 {
			if day < 21 {
				astro_sign = "Телец"
			} else {
				astro_sign = "Близнецы"
			}
			fmt.Println("Знак зодиака: ", astro_sign)
		} else if month == 6 {
			if day < 22 {
				astro_sign = "Близнецы"
			} else {
				astro_sign = "Рак"
			}
			fmt.Println("Знак зодиака: ", astro_sign)
		} else if month == 7 {
			if day < 23 {
				astro_sign = "Рак"
			} else {
				astro_sign = "Лев"
			}
			fmt.Println("Знак зодиака: ", astro_sign)
		} else if month == 8 {
			if day < 23 {
				astro_sign = "Лев"
			} else {
				astro_sign = "Дева"
			}
			fmt.Println("Знак зодиака: ", astro_sign)
		} else if month == 9 {
			if day < 23 {
				astro_sign = "Дева"
			} else {
				astro_sign = "Весы"
			}
			fmt.Println("Знак зодиака: ", astro_sign)
		} else if month == 10 {
			if day < 23 {
				astro_sign = "Весы"
			} else {
				astro_sign = "Скорпион"
			}
			fmt.Println("Знак зодиака: ", astro_sign)
		} else if month == 11 {
			if day < 23 {
				astro_sign = "Скорпион"
			} else {
				astro_sign = "Стрелец"
			}
			fmt.Println("Знак зодиака: ", astro_sign)
		} else if month < 1 || month > 12 {
			fmt.Println("Введите корректные данные.")
		}
	} else {
		fmt.Println("Введите корректные данные.")
	}
}
