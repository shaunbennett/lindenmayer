<p align="center"><img src="https://github.com/shaunbennett/lucis/raw/master/render/lsystem.png" width="256"></p>


# lindenmayer
A basic implementation of [Lindenmayer systems](https://en.wikipedia.org/wiki/L-system) in Golang.

This was used to generate L-System trees in my raytracer project [lucis](https://github.com/shaunbennett/lucis). See `main.go` for the lua code generation based off the L-System.

## Usage
```go
// Create a new L-System with axiom "S"
system := lsystem.New("S")

// Add basic production rules
system.AddProduction('S', "QX")
system.AddProduction('X', "B[LX[RX]B][RX[LX]B]X")

// Add two production rules with weight 10
// Weights are used as a probability of selecting that rule
// Since both weights are equal, there is a 50% chance for each rule
system.AddWeightedProduction('B', "BB", 10)
system.AddWeightedProduction('B', "B", 10)

// Iterate the system 8 times and print the result
fmt.Println(system.Iterate(8))
```
