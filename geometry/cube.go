package geometry

import "errors"

// CubeVolume calculates the volume of a cube where n is the cube length
func CubeVolume(n int) (int, error) {
	if n == 0 {
		return 0, errors.New("Zero is not allowed")
	}
	return n * n * n, nil
}
