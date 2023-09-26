package tilde

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func GetFloatsIn2Decimals[F constraints.Float](f F) string {
	return fmt.Sprintf("%.2f", f)
}
