package main

//Digit type
type Digit [5]string

//Digits type
type Digits []Digit

func generateSeprator() Digit {
	return Digit{
		"  ",
		"00",
		"  ",
		"00",
		"  ",
	}
}

//GenerateDigits func
func generateDigits() Digits {
	return Digits{
		{
			"000000",
			"00  00",
			"00  00",
			"00  00",
			"000000",
		},
		Digit{
			"0000  ",
			"  00  ",
			"  00  ",
			"  00  ",
			"000000",
		},
		Digit{
			"000000",
			"    00",
			"000000",
			"00    ",
			"000000",
		},
		Digit{
			"000000",
			"    00",
			"000000",
			"    00",
			"000000",
		}, {
			"00  00",
			"00  00",
			"000000",
			"    00",
			"    00",
		},
		Digit{
			"000000",
			"00    ",
			"000000",
			"    00",
			"000000",
		},
		Digit{
			"000000",
			"00    ",
			"000000",
			"00  00",
			"000000",
		},
		Digit{
			"000000",
			"    00",
			"    00",
			"    00",
			"    00",
		},
		Digit{
			"000000",
			"00  00",
			"000000",
			"00  00",
			"000000",
		},
		Digit{
			"000000",
			"00  00",
			"000000",
			"    00",
			"000000",
		},
	}
}

// CreateClock func
func CreateClock(h, m, s int) Digits {

	digits := generateDigits()

	return Digits{
		digits[h/10], digits[h%10],
		generateSeprator(),
		digits[m/10], digits[m%10],
		generateSeprator(),
		digits[s/10], digits[s%10],
	}
}
