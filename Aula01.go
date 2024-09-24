package main

import (
	"errors" // Corrigido o pacote de erros
	"fmt"
)

// Interface da lista
type List interface {
	Get(index int) (int, error)
	Size() int
	Add(e int)
	AddAt(e int, index int) error
	Remove(index int) error
}

//Estrtura abstratas

// Definindo a estrutura ArrayList
type ArrayList struct {
	v []int
	size int
}

// Inicializando o ArrayList com um tamanho inicial
func (l *ArrayList) Init(size int) error {
	if size > 0 {
		l.v = make([]int, size)
		//l.size = 0 // Inicializamos o size em 0
		return nil
	} else {
		return errors.New("Size deve ser maior que 0")
	}
}

// Método que retorna o tamanho da lista(Size)
func (l *ArrayList) Size() int {
	return l.size
}

// Método que obtém um elemento da lista(Get)
func (l *ArrayList) Get(index int) (int, error) {
	if index >= 0 && index < l.size {
		return l.v[index], nil
	} else {
		return 0, errors.New("Index fora dos limites da lista")
	}
}

// Método para adicionar um elemento no final da lista
func (l *ArrayList) Add(e int) {
	if l.size == len(l.v) {
		// Redimensionar o slice se estiver cheio
		l.v = append(l.v, e)
	} else {
		l.v[l.size] = e
	}
	l.size++
}

// Método para adicionar um elemento em uma posição específica
func (l *ArrayList) AddAt(e int, index int) error {
	if index < 0 || index > l.size {
		return errors.New("Índice fora do limite")
	}
	// Realocando elementos à direita se necessário
	l.v = append(l.v[:index+1], l.v[index:]...)
	l.v[index] = e
	l.size++
	return nil
}

// Método para remover um elemento em um índice específico
func (l *ArrayList) Remove(index int) error {
	if index < 0 || index >= l.size {
		return errors.New("Índice fora do limite")
	}
	// Realocar os elementos à esquerda para preencher o espaço removido
	l.v = append(l.v[:index], l.v[index+1:]...)
	l.size--
	return nil
}

func main() {
	l := &ArrayList{}
	err := l.Init(5)
	if err != nil {
		fmt.Println("Erro ao inicializar lista:", err)
		return
	}

	// Adicionando elementos
	l.Add(10)
	l.Add(20)
	l.Add(30)

	// Imprimindo elementos
	for i := 0; i < l.Size(); i++ {
		val, _ := l.Get(i)
		fmt.Println("Elemento", i, ":", val)
	}

	// Remover elemento
	l.Remove(1)
	fmt.Println("Após remover o índice 1:")
	for i := 0; i < l.Size(); i++ {
		val, _ := l.Get(i)
		fmt.Println("Elemento", i, ":", val)
	}
}
