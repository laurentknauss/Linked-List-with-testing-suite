/* 
The following codebase defones 2 functions. 
The `prepend` function addsa node to the front of the linked list. 
The `deleteWithValue` function deletes a node witha given value from the linked list. 
 he  `main` function creates a linked list and then calls the `prepend()` and `deleteWithValue()` functions to add and remove 
 nodes from the Linked List . 
*/
package main 

import (
	"fmt"
)


// first , we need to describe a node . 
type node  struct { 
	data int // a node needs to have data and we will only use simple numbers.
	next *node  // this is a pointer to the  address of the next node . 
}

//§ second, we need to describe the list 
type linkedList struct { 
	head *node
	length int // this is indicating how long the linked list is .
}

// The `prepend' function to put the  node inside the linked list. 
// The receiver is going to be `prepend` . The receiver (*linkedList) is a pointer receiver this time .
func(l *linkedList ) prePend(n *node)  { // this is going to take a node to be added at the front. 
   second := l.head
	 l.head = n 
	 l.head.next = second
	 l.length++
}

// The receiver (linkList) is now a value receiver this time.
func (l linkedList) printListData() { 
		var toPrint = l.head
	for l.length != 0 { 
		// inside the loop, let's print out he first node . 
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next // we keep updating `toPrint` to the next  , this is how we update the loop . 
		l.length--
	}
	fmt.Printf("\n")
}

// We are now creating a function that can delete with a given value 
func (l *linkedList) deleteWithValue(value int) {  
	if l.length == 0 { // if the length of the list is zero 
		return 
	}
	if l.head.data == value { 
		l.head = l.head.next 
		l.length--
		return 
	}
	
	previousToDelete := l.head
	for previousToDelete.next.data != value { 
      if previousToDelete.next.next ==nil { 
				return 
			}

		previousToDelete = previousToDelete.next
	}
	previousToDelete.next =  previousToDelete.next.next 
	l.length--
}


func main() { 
	mylist := linkedList{} 
	node1 := &node{data:48} // creating some nodes  . 
	node2 := &node{data:18}  
	node3 := &node{data:16} 
	node4 := &node{data:11}
	node5 := &node{data:7} 
	node6 := &node{data:2} 
	
	mylist.prePend(node1)
	mylist.prePend(node2)
	mylist.prePend(node3)
	mylist.prePend(node4) 
	mylist.prePend(node5)
	mylist.prePend(node6)
	

	/*
	mylist.prepend(node7)
	mylist.prepend(node8)
  mylist.prepend(node9)
  mylist.prepend(node10)
  */ 

	mylist.printListData()


mylist.deleteWithValue(100) 
mylist.deleteWithValue(2)

mylist.printListData()
emptyList := linkedList{} // to handle the case when there is an empty list . 
emptyList.deleteWithValue(10)


fmt.Println("The first value printed out  is the address of the last node  and the second value printed out is the length of mylist :  " , mylist)  
}