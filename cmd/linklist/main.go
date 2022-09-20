package main

import (
	"algo/internal/linklist"
	"fmt"
)

func main() {
	var list = linklist.LinkList[int]{}
	list.Init()
	var err error

	err = list.HeadAppend(2)
	if err != nil {
		panic(err)
	}
	elem, err := list.GetElem(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("elem: %+v, type: %T\n", elem, elem)

	err = list.HeadAppend(1)
	if err != nil {
		panic(err)
	}
	elem, err = list.GetElem(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("elem: %+v, type: %T\n", elem, elem)

	err = list.TailAppend(3)
	if err != nil {
		panic(err)
	}
	elem, err = list.GetElem(3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("elem: %+v, type: %T\n", elem, elem)

	err = list.Delete(2)
	if err != nil {
		panic(err)
	}
	elem, err = list.GetElem(2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("elem: %+v, type: %T\n", elem, elem)

	locate, err := list.LocateElem(3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("locate: %+v\n", locate)

	err = list.Insert(list.Length()+1, 4)
	if err != nil {
		panic(err)
	}
	locate, err = list.LocateElem(4)
	if err != nil {
		panic(err)
	}
	fmt.Printf("locate: %+v\n", locate)

	err = list.HeadDeduct()
	if err != nil {
		panic(err)
	}
	elem, err = list.GetElem(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("elem: %+v, type: %T\n", elem, elem)

	err = list.TailDeduct()
	if err != nil {
		panic(err)
	}
	elem, err = list.GetElem(list.Length())
	if err != nil {
		panic(err)
	}
	fmt.Printf("elem: %+v, type: %T\n", elem, elem)
}
