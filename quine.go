package main
import "fmt"
func main(){
	s := `package main
import "fmt"
func main(){
	s := %c%v%c
	fmt.Printf(s,0x60,s,0x60)
}`
	fmt.Printf(s,0x60,s,0x60)
}