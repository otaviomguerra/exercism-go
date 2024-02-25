// A package that contains a function IsLeapYear
// to check weather a given year is a leap year or not.

package leap

// IsLeapYear checks if a given year is a leap year.
func IsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	} else {
		return false
	}
}
