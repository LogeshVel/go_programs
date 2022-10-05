package main

import "fmt"

type Cow struct {
	food       string
	locomotion string
	noise      string
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func main() {

	var animalMap = map[string]Animal{}
	for {
		fmt.Printf("> ")
		var req, animalName, typeOrInfo string
		fmt.Scanf("%s %s %s", &req, &animalName, &typeOrInfo)

		switch req {
		case "newanimal":
			// creation
			switch typeOrInfo {
			case "cow":
				animalMap[animalName] = Cow{food: "grass", locomotion: "walk", noise: "moo"}
			case "bird":
				animalMap[animalName] = Bird{food: "worms", locomotion: "fly", noise: "peep"}
			case "snake":
				animalMap[animalName] = Snake{food: "mice", locomotion: "slither", noise: "hsss "}
			default:
				continue
			}
			fmt.Println("Created it!")
		case "query":
			// query
			var animal Animal
			if a, ok := animalMap[animalName]; ok {
				animal = a
			} else {
				continue
			}
			switch typeOrInfo {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			}
		}
	}

}
