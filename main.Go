package main

import "fmt"

func main() {
  fmt.println("matter")
}
**************************************************************************************

package main
import "fmt"

func main() {
  var colour string = "Green"
}
OUTPUT: go run main.go----> Green
*****************************************************************************************

var name string = "bangalore"
fmt.printf("The city name is %v", name)
OUTPUT: The city name is bangalore

var score int = 95
fmt.printf("Marks: %d", score)
OUTPUT: Marks: 95

fmt.print("Hey, %v! you have scored %i/100 in pollution", name,score)
OUTPUT: Hey, bangalore! you have scored 95/100 in pollution.

*********************************************************************************************

package main
import "fmt"

func main() {
  var(
  name string = "bangalore"
  marks int = 95)
  fmt.println(name)
  fmt.println(marks)
}
OUTPUT: bangalore
        95
*************************************************************************************************
package main
import "fmt"

func main() {
country:= "UK"
{
      city: "London"
      fmt.println(city)
      fmt.println(country)
}
fmt.println(country)
fmt.println(city)
}
OUTPUT: ERROR. Because local var(line:56) city cannot be declared outside.
// Correct code
func main() {
country:= "UK"
{
      city: "London"
      fmt.println(city)
      fmt.println(country)
}
fmt.println(country)
}
OUTPUT: LONDON
        UK
        UK

*************************************************************************************************
package main
import"fmt"
func main() {
  var colour string = "blue"
  var marks int = 95
  var valid bool = true

fmt.printf("shirt colour = %v is a type of %T \n", colour,colour)
fmt.printf("marks scored = %v is a type of %T \n", marks,marks)
fmt.printf("validation = %v is a type if %T \n", valid,valid)
}

OUTPUT: shirt coloue = blue is a type of string
        marks scored = 95 is a type of int
        validation = true is a type of bool

//OR WE CAN WRITE IT USING reflect.TypeOf()

package main
import{
"fmt"
"reflect"
}
func main() {
  fmt.printf("Type: %v \n", reflect.TypeOf(100))
  fmt.printf("Type: %v \n", reflect.TypeOf("Ahmed"))
}

OUTPUT: Type: int
        Type: string

*****************************************************************************************
package main
import (
  "fmt"
  "strconv")
func main() {
 var i int = 45
 var s string = strconv.itoa(i) // itoa means convert integer to string
 fmt.printf("%q", s)
}
OUTPUT: "45"
