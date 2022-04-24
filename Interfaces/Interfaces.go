package interfaces

import "fmt"

type printer interface {
	printInfo()
}

type Book struct {
	Title string
	Price float32
}

type Drink struct {
	Name  string
	Price float32
}

func (b Book) printInfo() {
	fmt.Printf("Book: %s Price: %.2f\n", b.Title, b.Price)
}

func (d Drink) printInfo() {
	fmt.Printf("Drink: %s Price: %.2f\n", d.Name, d.Price)
}

func myType(i interface{}) {
	switch i.(type) {
	case Book:
		fmt.Printf("I'm a book: %s\n", i.(Book).Title)
	case Drink:
		fmt.Printf("I'm a drink: %s\n", i.(Drink).Name)
	}
}

func main() {
	gunslinger := Book{
		Title: "The Gunslinger",
		Price: 4.75,
	}

	coffe := Drink{
		Name:  "Coffe",
		Price: 2.50,
	}

	info := []printer{gunslinger, coffe}

	info[0].printInfo()
	info[1].printInfo()


	myType(gunslinger)
	myType(coffe)
}
