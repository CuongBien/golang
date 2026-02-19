package main

func swapPointers(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}