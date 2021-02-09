package main

import (
	"./listlist"
)


func main() {
	ll := listadelistas.NewLL()
	ll.Insert(2,"1")
	ll.Insert(1,"11")
	ll.Insert(2,"21")
	ll.Insert(3,"31")
	ll.Insert(3,"41")
	ll.Print()
}