package main

func main() {
	// Build a tree
	// B -> Branch segment
	// L -> Leaf segment
	// [ -> Rotate Left
	// ] -> Rotate Right
	system := New("L")
	system.AddProduction('L', "B[L]L")
	system.AddProduction('B', "BB")

	//system.AddOutput('A', func() {
	//	println("A")
	//})
	//system.AddOutput('B', func() {
	//	println("Goodbye")
	//})

	println(system.Iterate(10))
}
