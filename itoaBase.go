package layi

//  writing a code that can convert a number to any base

func ItoaBase(value, base int) string {
	result := ""
	if base >= 2 && base <= 10 {
		result = Itoa(value, base)
	} else if base > 10 && base <= 16 {
		result = ItoaHex(value, base)
	} else if base < 2 || base > 16 {
		result = "Not a valid base"
	}
	return result
}

// I would write my own string convert i.e my own itoa function for numbers between 2 and 10

func Itoa(value int, base int) string {
	result := ""
	sign := ""
	if value < 0 {
		sign = "-"
		value = -value
	}
	var digits []string
	for value > 0 {
		digit := value % base
		digits = append(digits, string(digit+'0'))
		value = value / base
	}
	for i := len(digits) - 1; i >= 0; i-- {
		result += digits[i]
	}
	return sign + result
}

func ItoaHex(value int, base int) string {
	result := ""
	sign := ""
	if value < 0 {
		sign = "-"
		value = -value
	}
	var digits []string
	for value > 0 {
		switch base {
		case 11:
			digit := value % base
			if digit == 10 {
				digits = append(digits, "A")
			} else {
				digits = append(digits, string(digit+'0'))
			}
		case 12:
			digit := value % base
			switch digit {
			case 10:
				digits = append(digits, "A")
			case 11:
				digits = append(digits, "B")
			default:
				digits = append(digits, string(digit+'0'))
			}
		case 13:
			digit := value % base
			switch digit {
			case 10:
				digits = append(digits, "A")
			case 11:
				digits = append(digits, "B")
			case 12:
				digits = append(digits, "C")
			default:
				digits = append(digits, string(digit+'0'))
			}
		case 14:
			digit := value % base
			switch digit {
			case 10:
				digits = append(digits, "A")
			case 11:
				digits = append(digits, "B")
			case 12:
				digits = append(digits, "C")
			case 13:
				digits = append(digits, "D")
			default:
				digits = append(digits, string(digit+'0'))
			}
		case 15:
			digit := value % base
			switch digit {
			case 10:
				digits = append(digits, "A")
			case 11:
				digits = append(digits, "B")
			case 12:
				digits = append(digits, "C")
			case 13:
				digits = append(digits, "D")
			case 14:
				digits = append(digits, "E")
			default:
				digits = append(digits, string(digit+'0'))
			}
		case 16:
			digit := value % base
			switch digit {
			case 10:
				digits = append(digits, "A")
			case 11:
				digits = append(digits, "B")
			case 12:
				digits = append(digits, "C")
			case 13:
				digits = append(digits, "D")
			case 14:
				digits = append(digits, "E")
			case 15:
				digits = append(digits, "F")
			default:
				digits = append(digits, string(digit+'0'))
			}
		}

		value = value / base
	}
	for i := len(digits) - 1; i >= 0; i-- {
		result += digits[i]
	}
	return sign + result
}
