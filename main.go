package main
import "fmt"

func main(){
	cards := []string{"ploy", newCard()}
	cards = append(cards,"pen ba")
	
	for i,card := range cards{
		fmt.Println(i,card)
	}
}

func newCard() string {
	return "Chonnikan"
}