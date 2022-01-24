package main

import "fmt"

func main() {
	myBill := newBill("gamo's bill")

	myBill.addItem("onion soup", 4.50)
	myBill.addItem("coffee", 3.25)
	myBill.updateTip(10)

	fmt.Println(myBill.format())
}