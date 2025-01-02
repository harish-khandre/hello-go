package main

import "fmt"

// List represents a singly-linked list that holds values of comparable types.
type List[T comparable] struct {
	next *List[T] // Pointer to the next node
	val  T        // Value stored in this node
}

// Append adds a value to the end of the list.
func (l *List[T]) Append(value T) {
	if l == nil {
		panic("Cannot append to a nil list")
	}
	curr := l
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &List[T]{val: value}
}

// Prepend adds a value to the beginning of the list.
func (l *List[T]) Prepend(value T) *List[T] {
	return &List[T]{val: value, next: l}
}

// Length returns the number of elements in the list.
func (l *List[T]) Length() int {
	count := 0
	curr := l
	for curr != nil {
		count++
		curr = curr.next
	}
	return count
}

// Find searches for a value in the list and returns true if found.
func (l *List[T]) Find(value T) bool {
	curr := l
	for curr != nil {
		if curr.val == value {
			return true
		}
		curr = curr.next
	}
	return false
}

// Delete removes the first occurrence of a value from the list.
func (l *List[T]) Delete(value T) *List[T] {
	if l == nil {
		return nil
	}
	if l.val == value { // This works because T is constrained to be comparable
		return l.next // Skip the current node
	}
	curr := l
	for curr.next != nil {
		if curr.next.val == value { // This works because T is constrained to be comparable
			curr.next = curr.next.next // Remove the node
			break
		}
		curr = curr.next
	}
	return l
}

// Print displays the list elements.
func (l *List[T]) Print() {
	curr := l
	for curr != nil {
		fmt.Printf("%v -> ", curr.val)
		curr = curr.next
	}
	fmt.Println("nil")
}

func binaryTree() {
	// Example usage with comparable types
	list := &List[int]{val: 1}            // Initialize the list with a single node
	list.Append(2)                        // Append 2
	list.Append(3)                        // Append 3
	list = list.Prepend(0)                // Prepend 0
	list.Print()                          // Print the list: 0 -> 1 -> 2 -> 3 -> nil
	fmt.Println("Length:", list.Length()) // Length: 4

	fmt.Println("Find 2:", list.Find(2)) // Find 2: true
	fmt.Println("Find 4:", list.Find(4)) // Find 4: false

	list = list.Delete(2) // Delete 2
	list.Print()          // Print the list: 0 -> 1 -> 3 -> nil
}
