package models

type TestData[T any, K any] struct {
	Input    T
	Expected K
}

type ValueErrorCombination[T any] struct {
	Value    T
	HasError bool
}
