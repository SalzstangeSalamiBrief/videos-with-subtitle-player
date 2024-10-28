package models

type TestData[I any, E any] struct {
	Title    string
	Input    I
	Expected E
}

type ValueErrorCombination[V any] struct {
	Value    V
	HasError bool
}
