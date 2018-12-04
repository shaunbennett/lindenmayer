package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var RANDOM = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	system := New("S")
	system.AddWeightedProduction('S', "QX", 10)
	system.AddWeightedProduction('B', "BB", 10)
	system.AddWeightedProduction('B', "B", 10)
	system.AddProduction('X', "B[LX[RX]B][RX[LX]B]X")
	lString := system.Iterate(8)

	XZ := RANDOM.Float64() * 0.05
	Y := RANDOM.Float64() * 0.2
	branch, _ := createBranch(lString, 0, false, 0.3-XZ, 1-Y, 0.3-XZ)
	fmt.Printf("return %s", branch)
}

func createBranch(lString string, si int, lastWasBranch bool, lastScaleX float64, lastScaleY float64, lastScaleZ float64) (string, int) {
	children := make([]string, 0)
	i := si

OUTER:
	for ; i < len(lString); i++ {
		switch lString[i] {
		case '[':
			var branchName string
			branchName, i = createBranch(lString, i+1, lastWasBranch, lastScaleX, lastScaleY, lastScaleZ)
			if len(children) == 0 {
				node := fmt.Sprintf("node%d", i)
				fmt.Printf("%s = rt.node('%s')\n", node, node)
				children = append(children, node)
			}
			if branchName != "" {
				fmt.Printf("%s:add_child(%s)\n", children[len(children)-1], branchName)
			}
		case ']':
			break OUTER
		case 'B':
			name := fmt.Sprintf("branch%d", i)
			children = append(children, name)
			fmt.Printf("%s = rt.cylinder('%s')\n", name, name)
			if lastWasBranch {
				fmt.Printf("%s:translate(0, 1, 0)\n", name)
			} else {
				XZ := (RANDOM.Float64() * 0.1) + 0.04
				lastScaleX = math.Max(0.05, lastScaleX-XZ)
				lastScaleY = math.Max(0.3, lastScaleY-(RANDOM.Float64()*0.1)-0.05)
				lastScaleZ = math.Max(0.05, lastScaleZ-XZ)
				fmt.Printf("%s:scale(%f, %f, %f)\n", name, lastScaleX, lastScaleY, lastScaleZ)
			}
			fmt.Printf("%s:set_material(tree)\n", name)
			lastWasBranch = true
		case 'R':
			name := fmt.Sprintf("branchrot%d", i)
			children = append(children, name)
			fmt.Printf("%s = rt.node('%s')\n", name, name)
			zRotation := RANDOM.Intn(20) + 25
			yRotation := RANDOM.Intn(160) - 80
			fmt.Printf("%s:rotate('z', %d)\n", name, zRotation)
			fmt.Printf("%s:rotate('y', %d)\n", name, yRotation)
			fmt.Printf("%s:scale(1/%f, 1/%f, 1/%f)\n", name, lastScaleX, lastScaleY, lastScaleZ)
			fmt.Printf("%s:translate(0, 0.9, 0)\n", name)
			lastWasBranch = false
		case 'L':
			name := fmt.Sprintf("branchrot%d", i)
			children = append(children, name)
			fmt.Printf("%s = rt.node('%s')\n", name, name)
			zRotation := -(RANDOM.Intn(20) + 25)
			yRotation := RANDOM.Intn(160) - 80
			fmt.Printf("%s:rotate('z', %d)\n", name, zRotation)
			fmt.Printf("%s:rotate('y', %d)\n", name, yRotation)
			fmt.Printf("%s:scale(1/%f, 1/%f, 1/%f)\n", name, lastScaleX, lastScaleY, lastScaleZ)
			fmt.Printf("%s:translate(0, 0.9, 0)\n", name)
			lastWasBranch = false
		case 'Q':
			name := fmt.Sprintf("branchrot%d", i)
			children = append(children, name)
			fmt.Printf("%s = rt.node('%s')\n", name, name)
			zRotation := RANDOM.Intn(20) - 10
			yRotation := RANDOM.Intn(160) - 80
			fmt.Printf("%s:rotate('z', %d)\n", name, zRotation)
			fmt.Printf("%s:rotate('y', %d)\n", name, yRotation)
			lastWasBranch = false
		}
	}

	// add children
	currIndex := len(children) - 1
	prev := ""
	for currIndex >= 0 {
		if prev != "" && children[currIndex] != "" {
			fmt.Printf("%s:add_child(%s)\n", children[currIndex], prev)
		}
		if children[currIndex] != "" {
			prev = children[currIndex]
		}
		currIndex--
	}
	if len(children) > 0 {
		return children[0], i
	} else {
		return "", i
	}
}
