package stack

import "errors"

/**
 * Стек
 **/
type Stack []interface{}

/**
 * Получение размера стека
 **/
func (stack Stack) Len() int {
	return len(stack)
}

/**
 * Положить значение на стек
 **/
func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

/**
 * Забрать значение со стека
 **/
func (stack *Stack) Pop() (interface{}, error) {
	currentStack := *stack

	if len(currentStack) == 0 {
		return nil, errors.New("Empty stack")
	}

	value := currentStack[len(currentStack)-1]
	*stack = currentStack[:len(currentStack)-1]

	return value, nil
}
