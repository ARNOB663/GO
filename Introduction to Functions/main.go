package main 

import "fmt"

func add(a int,b int) int{
	 sum := a+b

	 return sum
}

func getusername() string {
	var name string
	fmt.Println("Enter your name")
	fmt.Scan(&name)
	return name

}
func printSomething(){
	
	
fmt.Print("Education must be free")

}

func main() {
   var a, b int 
	fmt.Scan(&a)
	fmt.Scan(&b)
  
   //result := add(a,b)
  // fmt.Println(result)
  var name string = getusername()
   fmt.Println("Hello",name)
}