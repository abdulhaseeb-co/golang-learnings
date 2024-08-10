package main

import "fmt"

func main() {
    var u1 uint32      //declare a variable and init with 0
	u1 = 32            //assign its value
	var u2 uint32 = 32 //declare a variable and assign its value at once
	//declare a new variable with defining data type:
	u3 := uint32(32)        //inside the function block this is equal to: var u3 uint32 = 32
	fmt.Println(u1, u2, u3) 
	u3 = 20
	fmt.Println(u1, u2, u3) 
}