package november

import (
	"log"
	"testing"
)

//An interface can be embedded in another interface as well as it can be embedded in a struct.

// Embedding interface in another interface
// 在一个接口中嵌入另一个接口
type animal interface {
	breathe()
	walk()
}

// another interface named human which embeds the animal interface
type human interface {
	animal
	speak()
}

type employee struct {
}

func (e employee) speak() {
	log.Println("employee speak")
}
func (e employee) walk() {
	log.Println("employee walk")
}
func (e employee) breathe() {
	log.Println("employee breathe")
}

func TestEmbedding1(t *testing.T) {
	e1 := employee{}
	e1.speak()
	e1.walk()
	e1.breathe()
}

//Embedding interface in a struct
//在一个结构体中嵌入另一个结构体

type dog struct {
	name string
}

func (d dog) walk() {
	log.Printf("%v is walk\n", d.name)
}
func (d dog) breathe() {
	log.Printf("%v is beathe\n", d.name)
}

type pet struct {
	animal
}

type pet2 struct {
	a animal
}

func TestEmbedingInterfaceInStruct(t *testing.T) {
	dog1 := dog{name: "DD"}
	p1 := pet{animal(dog1)}
	p2 := pet2{a: dog1}
	p1.walk()
	p1.breathe()
	p1.animal.breathe()
	p1.animal.walk()

	p2.a.walk()
	p2.a.breathe()
	//p2.breathe()
	//p2.walk()

}
