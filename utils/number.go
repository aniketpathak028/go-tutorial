package utils

import "math/rand"

// the Number() function is exported
// (i.e. its name begins with a capital letter).
// This means it will be visible to any code which imports
//  the random package.

func Number() int {
	return rand.Intn(100)
}
