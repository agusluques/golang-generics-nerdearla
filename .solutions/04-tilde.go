package results

import "fmt"

type Float interface {
	~float32 | ~float64
}

func GetFloatsIn2Decimals[F Float](f F) string {
	return fmt.Sprintf("%.2f", f)
}
