package uzgo

import (
	"fmt"
	"config"
)

func main() {
	config.LoadConfig()
	fmt.Println("Hello, GO!")
}