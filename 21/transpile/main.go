// +build ignore

package main

func main() {
	var r0 int
	var r5 int
	var r2 int
	var r3 int
	var r1 int
	var r4 int

	r1 = 123
	for {
		r1 = r1 & 456
		if r1 == 72 {
			r1 = 1
		} else {
			r1 = 0
		}
		if r1 == 1 {
			break
		}
	}
	r1 = 0
	for {
		r5 = r1 | 65536
		r1 = 8586263
		for {
			r2 = r5 & 255
			r1 = r1 + r2
			r1 = r1 & 16777215
			r1 = r1 * 65899
			r1 = r1 & 16777215
			if 256 >= r5 {
				r2 = 1
			} else {
				r2 = 0
			}
			r4 = r2 + r4
			r4 = r4 + 1
			r4 = 27
			r2 = 0
			for {
				r3 = r2 + 1
				r3 = r3 * 256
				if r3 >= r5 {
					r3 = 1
				} else {
					r3 = 0
				}
				r4 = r3 + r4
				r4 = r4 + 1
				r4 = 25
				r2 = r2 + 1
			}
			r5 = r2
		}
		if r1 == r0 {
			r2 = 1
		} else {
			r2 = 0
		}
		if r2 == 1 {
			break
		}
	}
}
