package interfacedesign

import (
	"errors"
	"fmt"
)

type monarchy interface {
	birth(child string, parent string) error
	death(name string) error
	getOrderOfSuccession() []string
}

type FamilyTree struct {
	root        *Person
	TreeMap     map[string]*Person
	livingCount int
}

type Person struct {
	name     string
	children []*Person
	isAlive  bool
}

func StartFamilyTree(name string) FamilyTree {

	root := Person{name: "Jake", isAlive: true}

	mon := FamilyTree{
		root:        &root,
		livingCount: 1,
		TreeMap:     map[string]*Person{},
	}

	mon.TreeMap[root.name] = &root

	return mon
}

func (ft *FamilyTree) birth(child string, parent string) error {

	parentNode, ok := ft.TreeMap[parent]

	if !ok {

		return errors.New("parent does not exist")
	}

	childNode := Person{
		name:     child,
		children: []*Person{},
		isAlive:  true,
	}

	parentNode.children = append(parentNode.children, &childNode)

	ft.TreeMap[childNode.name] = &childNode

	ft.livingCount++

	return nil
}

func (ft *FamilyTree) death(name string) error {

	targetNode, ok := ft.TreeMap[name]

	if !ok {

		errors.New("the person does not exist")
	}

	if !targetNode.isAlive {

		return errors.New("the person, sadly, is already dead :(")
	}

	targetNode.isAlive = false
	ft.livingCount--

	return nil
}

var gloabalCount int

func (ft *FamilyTree) getOrderOfSuccession() []string {

	orderOfSuccession := make([]string, ft.livingCount)

	gloabalCount = 0
	preOrderTraversal(ft.root, orderOfSuccession)

	return orderOfSuccession
}

func preOrderTraversal(node *Person, order []string) {

	if node.isAlive {

		order[gloabalCount] = node.name
		gloabalCount++
	}

	for _, child := range node.children {

		preOrderTraversal(child, order)
	}
}

func MonarchyCaller() {

	mon := StartFamilyTree("Jake")

	mon.birth("Catherine", "Jake")
	mon.birth("Tom", "Jake")
	mon.birth("Celine", "Jake")
	mon.birth("Peter", "Celine")
	mon.birth("Jane", "Catherine")
	mon.birth("Farah", "Jane")
	mon.birth("Mark", "Catherine")

	fmt.Println(mon.getOrderOfSuccession())

	mon.death("Jake")
	mon.death("Jane")

	fmt.Println(mon.getOrderOfSuccession())
}
