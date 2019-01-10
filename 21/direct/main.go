package main

import (
	"fmt"
	"time"
)

func main() {
	r1, r5 := 0, 0
	prev := 0
	seen := make(map[int]struct{})

	start := time.Now()

	for {
		r1 = 123
		if (r1 & 456) == 72 {
			break
		}
	}

	r1 = 0

done:
	for {
		prev = r1
		r5 = r1 | 65536
		r1 = 8586263

		for {
			// fmt.Printf("r1: %d, r5: %d\n", r1, r5)
			r1 = (((r1 + (r5 & 255)) & 16777215) * 65899) & 16777215
			if 256 > r5 {
				if len(seen) == 0 {
					// first termination: dump the value of r1
					fmt.Printf("Part 1 in %v: %d\n", time.Since(start), r1)
				}
				if _, ok := seen[r1]; ok {
					// First duplicate: dump the previous value of r1
					fmt.Printf("Part 2 in %v: %d\n", time.Since(start), prev)
					break done
				} else {
					seen[r1] = struct{}{}
				}
				break
			} else {
				// r2 = 0
				// for {
				// 	r3 = r2 + 1
				// 	r3 *= 256
				// 	if r3 > r5 {
				// 		r5 = r2
				// 		break
				// 	} else {
				// 		r2++
				// 	}
				// }
				r5 = r5 / 256
			}
		}
	}
}
