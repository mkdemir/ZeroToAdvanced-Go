package error_handling

import (
	"fmt"
)

type borderException struct {
	parameter int
	message   string
}

func (b *borderException) Error() string { // pointer vererek performansı attırabilirsiniz.
	return fmt.Sprintf("%d---%s", b.parameter, b.message)
}

func TahminEt2(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", &borderException{tahmin, "Sınırların dışında kaldı"}
	}

	return "Bildiniz", nil
}
