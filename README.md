<h1 align="center">Astrological sign or Zodiac sign
<p target="_blank">Astronomy is the knowledge of the stars <img src="https://github.com/blackcater/blackcater/raw/main/images/Hi.gif" height="32"/></p></h1>
<h3 align="center">The program shows which sign of the Zodiac, with the entered parameters: day and month.
</h3>
<img src="https://badges.frapsoft.com/os/v1/open-source.svg?v=103" >

## The goal of the work
Program to display Astrological sign or Zodiac sign for given date of birth.

## Task statement
Two integers are given: D (day) and M (month), which determine the correct date. Derive the sign of the Zodiac that corresponds to this date: "Aquarius" (20.1–18.2), "Pisces" (19.2–20.3), "Aries" (21.3–19.4), "Taurus" (20.4–20.5), "Gemini" (21.5 –21.6), “Cancer” (22.6–22.7), “Leo” (23.7–22.8), “Virgo” (23.8–22.9), “Libra” (23.9–22.10), “Scorpio” (23.10–22.11), “ Sagittarius" (23.11-21.12), "Capricorn" (22.12-19.1).

# Additional Information
For given date of birth, this program displays an astrological sign or Zodiac sign.
- ## Examples :
``` 
Input : Day = 10, Month = December
Output : Sagittarius
Explanation :
People born on this date have a zodiac Sagittarius.

Input : Day = 7, Month = September
Output : Virgo
``` 
- ## Approach :
Although the exact dates can shift plus or minus a day, depending on the year, here are the general zodiac sign dates used by Western (or Tropical) astrology :
``` 
WESTERN ASTROLOGY STAR SIGN DATES :

Aries (March 21-April 19)
Taurus (April 20-May 20)
Gemini (May 21-June 20)
Cancer (June 21-July 22)
Leo (July 23-August 22)
Virgo (August 23-September 22)
Libra (September 23-October 22)
Scorpio (October 23-November 21)
Sagittarius (November 22-December 21)
Capricorn (December 22-January 19)
Aquarius (January 20-February 18)
Pisces (February 19-March 20) 
``` 
We need to check our mentioned date and month and thus find its equivalent zodiac, i.e which zodiac fits in that particular date as well as month and print its corresponding zodiac sign.

## Notes
This program is implemented in the **Go** language, the following libraries were used: **`fmt`** and **`math`**

## Notation of variables
This table shows the short designation of the variable in the code and its description

| Name       | Type   | Description                      |
| ---------- | ------ | -------------------------------- |
| `day` | int8 | the day of the astrological sign |
| `month` | int8 | the month of the astrological sign  |
| `astro_sign` | string | The astrological sign |

## Condition code in the program
- Сhecks month and date within the valid range of a specified zodiac
``` Go
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
```

## Result
Screenshot that displays the result of the entered data
<p align="center">
<img  src="https://github.com/HeorhiiZemlianko/Calculate-BMI-and-risk-category/blob/main/task2golang/Снимок.PNG"  width="350" alt="Calculate-BMI-and-risk-category"/>
</p>
