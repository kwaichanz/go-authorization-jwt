// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Single struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func main() {
	fmt.Println("Hello, 世界")
	showSingle()
}

var singles = []Single{
	{
		ID:     "1",
		Name:   "Koi",
		Author: "Aimer",
	},
	{
		ID:     "2",
		Name:   "Red",
		Author: "Taylor Swift",
	},
}

func showSingle() {
	for _, single := range singles {
		if single.ID == "1" {
			fmt.Println("ffff ", singles[0])
			return
		}
	}
}
