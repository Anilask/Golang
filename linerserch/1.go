/*package main
import (
	"fmt"
)
func main(){
	items:= []string{"Anil","kumar","Akula"}
	liner := linersearch(items,"Anil")
	fmt.Println(liner)
}
func linersearch(d []string, key string)bool{
	for _,items :=range d{
		if items ==key{
			return true
		}
	}
	return false
}*/
package main
import (
	f "fmt"
)
func main() {
item := []int{1,1,2,2,3,5,6}
a :=liner(item,5)
f.Println(a)
}
func liner(a []int,key int)int{
for _,items :=range a{
	if items ==key{
		return items
	}
}
return int(1)
}
