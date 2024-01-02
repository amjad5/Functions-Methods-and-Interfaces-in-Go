package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println(c.name + " eat: grass")
}

func (c Cow) Move() {
	fmt.Println(c.name + "'s movement: walk")
}

func (c Cow) Speak() {
	fmt.Println(c.name + " Speaks: moo")
}

type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println(b.name + " eats: worms")
}

func (b Bird) Move() {
	fmt.Println(b.name + "'s movement: fly")
}

func (b Bird) Speak() {
	fmt.Println(b.name + " Speaks: peep")
}

type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Println(s.name + " eat: mice")
}

func (s Snake) Move() {
	fmt.Println(s.name + "'s movement: slither")
}

func (s Snake) Speak() {
	fmt.Println(s.name + ", speaks: " + " hsss")
}

func creationMessage() {
	fmt.Println("Created it!")
}

func actionNotFound() {
	fmt.Println("Action not found!")
}

func main() {
	var animals = make([]Animal, 0)
	for {
		var command, animalNameInput, animalTypeOrActionInput string
		fmt.Println("")
		fmt.Println("Please enter commands as following or x to exit")

		fmt.Println("newanimal cobra snake")
		fmt.Println("query cobra move")
		fmt.Println("")
		_, err := fmt.Scanf("%s %s %s %s", &command, &animalNameInput, &animalTypeOrActionInput)
		_ = err

		if command == "newanimal" {
			var ani Animal
			switch animalTypeOrActionInput {
			case "cow":
				var aniObject = Cow{animalNameInput}
				ani = aniObject
				animals = append(animals, ani)
				creationMessage()
			case "bird":
				var aniObject = Bird{animalNameInput}
				ani = aniObject
				animals = append(animals, ani)
				creationMessage()
			case "snake":
				var snakeObject = Snake{animalNameInput}
				ani = snakeObject
				animals = append(animals, ani)
				creationMessage()
			default:
			}
		} else if command == "query" {
			for k := range animals {
				cow, ok := animals[k].(Cow)
				if ok && cow.name == animalNameInput {
					switch animalTypeOrActionInput {
					case "eat":
						cow.Eat()
					case "move":
						cow.Move()
					case "speak":
						cow.Speak()
					default:
						actionNotFound()
					}
				}

				bird, ok := animals[k].(Bird)
				if ok && bird.name == animalNameInput {
					switch animalTypeOrActionInput {
					case "eat":
						bird.Eat()
					case "move":
						bird.Move()
					case "speak":
						bird.Speak()
					default:
						actionNotFound()
					}
				}

				snake, ok := animals[k].(Snake)
				if ok && snake.name == animalNameInput {
					switch animalTypeOrActionInput {
					case "eat":
						snake.Eat()
					case "move":
						snake.Move()
					case "speak":
						snake.Speak()
					default:
						actionNotFound()
					}
				}

				if !ok {
					fmt.Println("incorrect query!")
				}
			}

			if len(animals) == 0 {
				fmt.Println("No animals added yet")
			}
		} else if command == "x" {
			break
		}

	}
}
