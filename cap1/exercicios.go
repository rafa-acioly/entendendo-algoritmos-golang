package cap1

import "errors"

// exercicio da pagina 27
func pesquisaBinaria(list []int, item int) (int, error) {
	baixo, alto := 0, len(list)-1

	for baixo <= alto {
		meio := (baixo + alto) / 2
		chute := list[meio]
		if chute == item {
			return meio, nil
		} else if chute > item {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}
	return 0, errors.New("Not found")
}
