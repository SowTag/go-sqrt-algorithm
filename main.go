package main

import "fmt"

func Sqrt(x float64, maxIterations int) (z float64) {
	zIterations := 0
	z = 1 // Starting guess

	for zIterations < maxIterations {
		zIterations++

		error := x - z * z
		zSquared := z * z
		
		// Show full status only on low (<1000) iteration counts and one each 100 for higher iteration counts
		if maxIterations < 1000 || zIterations % 100 == 0 {
			fmt.Printf("iteration: #%d	 z: %g	error: %g	z^2: %g", zIterations, z, error, zSquared)
		}
		
		// z adjustment
		z -= (z*z - x) / (2 * z)
	}
	
	return
}

func main() {
	fmt.Println(Sqrt(16, 10)) // The second argument determines the amount of iterations z will go through.
}
