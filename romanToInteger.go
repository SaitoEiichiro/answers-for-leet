package main

func romanToInt(s string) int {
	ans := 0
	beforeMessage := ""
	for _, i := range s {
		strI := string(i)
		if strI == "M" {
			if beforeMessage == "C" {
				ans += 800
			} else {
				ans += 1000
			}
			beforeMessage = "M"
		}
		if strI == "D" {
			if beforeMessage == "C" {
				ans += 300
			} else {
				ans += 500
			}
			beforeMessage = "D"
		}
		if strI == "C" {
			if beforeMessage == "X" {
				ans += 80
			} else {
				ans += 100
			}
			beforeMessage = "C"
		}
		if strI == "L" {
			if beforeMessage == "X" {
				ans += 30
			} else {
				ans += 50
			}
			beforeMessage = "L"
		}
		if strI == "X" {
			if beforeMessage == "I" {
				ans += 8
			} else {
				ans += 10
			}
			beforeMessage = "X"
		}
		if strI == "V" {
			if beforeMessage == "I" {
				ans += 3
			} else {
				ans += 5
			}
			beforeMessage = "V"
		}
		if strI == "I" {
			ans += 1
			beforeMessage = "I"
		}
	}
	return ans
}
