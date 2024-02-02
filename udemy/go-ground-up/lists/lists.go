package lists

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func lists() {

	arr := [2]string{}
	fmt.Println(arr)
	hobbiesArr := [3]string{"hobbie1", "hobbie2", "hobbie3"}
	fmt.Println(hobbiesArr)
	fmt.Println(hobbiesArr[0])
	fmt.Println(hobbiesArr[1:3])
	slice1 := hobbiesArr[0:2]
	slice2 := hobbiesArr[1:3]
	fmt.Println(slice1)
	fmt.Println(slice2)
	dyn := []string{"goal1", "goal2"}
	fmt.Println(dyn)
	dyn[1] = "changedGoal"
	dyn2 := append(dyn, "goal3")
	fmt.Println(dyn2)
	products := []Product{Product{}, Product{}}
	fmt.Println(products)
	products = append(products, Product{})
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
