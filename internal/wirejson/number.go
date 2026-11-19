package wirejson

import (
	"math"
	"strconv"
	"strings"
)

// FormatNumber rewrites JSON number text as JavaScript writes the same number (ECMAScript
// Number::toString): 300.0 becomes 300, 1E-07 becomes 1e-7, 1e21 becomes 1e+21 and
// 1.2345678901234568e20 becomes 123456789012345680000. A number a float64 can't hold exactly, such
// as 9007199254740993 (above 2^53) or 9007199254740992.5, is kept as written, so no value changes.
func FormatNumber(text string) string {
	value, err := strconv.ParseFloat(strings.TrimSpace(text), 64)
	if err != nil || math.IsInf(value, 0) || math.IsNaN(value) {
		return text
	}
	js := formatFloat(value)
	if sameValue(text, js) {
		return js
	}
	return text
}

// FormatFloat writes a finite float64 as JavaScript's Number.prototype.toString does.
func FormatFloat(value float64) string {
	return formatFloat(value)
}

func formatFloat(value float64) string {
	if value == 0 {
		return "0"
	}
	_, digits, point, _ := decompose(strconv.FormatFloat(math.Abs(value), 'e', -1, 64))
	k := len(digits)
	n := int(point)
	var body string
	switch {
	case k <= n && n <= 21:
		body = digits + strings.Repeat("0", n-k)
	case 0 < n && n <= 21:
		body = digits[:n] + "." + digits[n:]
	case -6 < n && n <= 0:
		body = "0." + strings.Repeat("0", -n) + digits
	default:
		sign := "+"
		if n-1 < 0 {
			sign = "-"
		}
		exponent := sign + strconv.Itoa(abs(n-1))
		if k == 1 {
			body = digits + "e" + exponent
		} else {
			body = digits[:1] + "." + digits[1:] + "e" + exponent
		}
	}
	if value < 0 {
		return "-" + body
	}
	return body
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// sameValue reports whether two number texts are the same decimal value, judged from every digit.
func sameValue(a, b string) bool {
	negA, digitsA, pointA, okA := decompose(a)
	negB, digitsB, pointB, okB := decompose(b)
	if !okA || !okB {
		return false
	}
	if digitsA == "" || digitsB == "" {
		// Zero, whatever its sign: JavaScript writes -0 as 0.
		return digitsA == digitsB
	}
	return negA == negB && digitsA == digitsB && pointA == pointB
}

// decompose splits number text into its sign, its significant digits (no leading or trailing zeros;
// empty for zero) and the decimal point position, so the value is 0.digits × 10^point.
func decompose(text string) (negative bool, digits string, point int64, ok bool) {
	s := strings.TrimSpace(text)
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
		negative = s[0] == '-'
		s = s[1:]
	}
	var exponent int64
	if e := strings.IndexAny(s, "eE"); e >= 0 {
		exp, err := strconv.ParseInt(s[e+1:], 10, 64)
		if err != nil {
			return false, "", 0, false
		}
		exponent = exp
		s = s[:e]
	}
	whole, fraction := s, ""
	if dot := strings.IndexByte(s, '.'); dot >= 0 {
		whole, fraction = s[:dot], s[dot+1:]
	}
	all := whole + fraction
	if all == "" {
		return false, "", 0, false
	}
	for i := 0; i < len(all); i++ {
		if all[i] < '0' || all[i] > '9' {
			return false, "", 0, false
		}
	}
	leading := 0
	for leading < len(all) && all[leading] == '0' {
		leading++
	}
	digits = strings.TrimRight(all[leading:], "0")
	point = int64(len(whole)-leading) + exponent
	return negative, digits, point, true
}
