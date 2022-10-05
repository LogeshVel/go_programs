package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}
func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss "}

	for {
		fmt.Printf("> ")
		var selectedAnimal Animal
		var a, i string
		fmt.Scanf("%s %s", &a, &i)
		switch a {
		case "cow":
			selectedAnimal = cow
		case "bird":
			selectedAnimal = bird
		case "snake":
			selectedAnimal = snake
		}
		switch i {
		case "eat":
			selectedAnimal.Eat()
		case "move":
			selectedAnimal.Move()
		case "speak":
			selectedAnimal.Speak()
		}
	}

}
