// ============================================================================
// = build.go
// = 	Description		various array or slice building functions
// = 	Date			2025.01.28
// ============================================================================

package utils

import "math/big"

// enables usage of big.Int by initializing a slice of len
func CreateSlice(len int64) []*big.Int {
	slice := make([]*big.Int, 0)
	for i := int64(0); i < len; i++ {
		slice = append(slice, big.NewInt(0))
	}
	return slice
}
