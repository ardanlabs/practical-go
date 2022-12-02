package main

import "fmt"

func main() {
	// Initializing structs
	var i1 Item
	fmt.Printf("i1: %#v\n", i1)

	i2 := Item{1, 2}
	fmt.Printf("i2: %#v\n", i2)

	i3 := Item{
		Y: 20,
		// X: 10,
	}
	fmt.Printf("i3: %#v\n", i3)

	fmt.Println(NewItem(100, 200))
	fmt.Println(NewItem(-100, 200))

	fmt.Println(Item{-1, -2})
	i3.Move(200, 500)
	fmt.Printf("i3 (move): %#v\n", i3)

	p1 := Player{
		Name: "Parzival",
		Item: Item{500, 300},
	}
	fmt.Println("p1.Name:", p1.Name)
	//	fmt.Println("p1.X:", p1.X)
	fmt.Println("p1.Item.X:", p1.Item.X)

	p1.Move(17, 530)
	fmt.Println("p1 (move):", p1)
	fmt.Println(p1.Found(Jade))   // nil
	fmt.Println(p1.Keys)          // [jade]
	fmt.Println(p1.Found(Jade))   // nil
	fmt.Println(p1.Keys)          // [jade]
	fmt.Println(p1.Found(Key(7))) // <error>
	fmt.Println(p1.Keys)          // [jade]

	ms := []Mover{
		&i1,
		&i2,
		&p1,
	}
	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Printf("moved: %#v\n", m)
	}

	/* How fmt checks if a type implement fmt.Stringer
	var i any = &p1
	_, ok := (i).(Mover)
	fmt.Println(ok)
	*/
}

/* Exercise:
- Add a Keys slice of string to Player
- Add a "Found(key string) error" method to player
	- Err if key is not one of "crystal", "copper", "jade"
	- Add the key to the Keys field only if it's not there already
*/

/*
type T struct {
	X float64
}
*/

/*
- Interface is a set of methods (and types)
- Interfaces say what we need (not what we provide)
- Rule of thumb: Accept interfaces, return types
*/
type Mover interface {
	Move(int, int)
}

func moveAll(ms []Mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

func (p *Player) Found(key Key) error {
	if !(key == Jade || key == Copper || key == Crystal) {
		return fmt.Errorf("unknown key: %#v", key)
	}

	if !p.hasKey(key) {
		p.Keys = append(p.Keys, key)
	}
	return nil
}

func (p *Player) hasKey(key Key) bool {
	for _, k := range p.Keys {
		if k == key {
			return true
		}
	}
	return false
}

// implement fmt.Stringer
func (k Key) String() string {
	switch k {
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	case Jade:
		return "jade"
	}

	return fmt.Sprintf("<Key %d>", k)
}

type Key byte

const (
	Crystal Key = iota + 1
	Copper
	Jade
)

type Player struct {
	Name string
	// X    float64 // will "shadow" Item.X
	Item // Player embeds Item, there's an implicit "Item" field in Player
	// T
	Keys []Key
}

// i is called "the receiver"
// i is a "pointer receiver"
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

// func NewItem(x, y int) Item {
// func NewItem(x, y int) *Item {
// func NewItem(x, y int) (Item, error) {
func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds (%d/%d)", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}
	return &i, nil // The Go compiler does escape analysis and will allocate i on the heap
}

const (
	maxX = 1000
	maxY = 600
)

// Item is an item in the game
type Item struct {
	X int
	Y int

	// height int
}

/*
func (i *Item) Height() int {
	return i.height
}

func (i *Item) SetHeight(h int) {
	i.height = h
}
*/

/*
Go
type Reader interface {
	Read(p []byte) (n int, err error)
}

Python
type Reader interface {
	Read(n int) (p []byte, err error)
}
*/

/* Design a Sortable interface */
type Sortable interface {
	Less(i, j int) bool
	Len() int
	Swap(i, j int)
}
