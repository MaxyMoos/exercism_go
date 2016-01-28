// This iteration has an average Convert speed of 717ns/op on my machine
// Intel Core i5-4570 Quad-Core 3.2GHz, 4GB RAM, Ubuntu 14.04 LTS
package raindrops

import "strconv"

const TestVersion = 1

func Convert(num int) string {
	result := ""
	if num%3 == 0 {
		result += "Pling"
	}
	if num%5 == 0 {
		result += "Plang"
	}
	if num%7 == 0 {
		result += "Plong"
	}
	// strconv.Itoa is faster than fmt.Sprintf (approx. 300ns/op)
	if len(result) == 0 {
		result = strconv.Itoa(num)
	}
	return result
}
