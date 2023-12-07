/* 
The following codebase defines 2 functions. 
The `prepend` function adds a node to the front of the linked list. 
The `deleteWithValue` function deletes a node with a given value from the linked list. 
The  `main` function creates a linked list and then calls the `prepend()` and `deleteWithValue()` functions
to add and remove nodes from the Linked List . 

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

// second, we need to describe the list 
type linkedList struct { 
	head *node
	length int // this is indicating how long the linked list is .
}


// The 'string' function provides a string representation of the linked list.
func (l linkedList) String() string { 
	return fmt.Sprintf("Length: %d, Head: %v", l.length, l.head) 
}


// The `prepend' function to put the  node inside the linked list. 
// The receiver is going to be `prepend` . The receiver (*linkedList) is a pointer receiver this time .

// This is going to take a node to be added at the front. 
func(l *linkedList ) Prepend(n *node)  { 
   second := l.head
	 l.head = n 
	 l.head.next = second
	 l.length++
}

// The 'printListData' function  prints the data of all nodes in the linked list.
func (l linkedList) PrintListData() { 
	toPrint := l.head
	for l.length != 0 { 
		// Inside the loop, let's print out he first node . 
		fmt.Printf("%d ", toPrint.data)
		// We keep updating `toPrint` to the next  , this is how we update the loop .
		toPrint = toPrint.next 
		l.length--
	}
	fmt.Printf("\n")
}

// The 'DeleteWithValue' function removes the node with a given value . 
func (l *linkedList) DeleteWithValue(value int) {  
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
	// creating some nodes . 
	node1 := &node{data:48} 
	node2 := &node{data:18}  
	node3 := &node{data:16} 
	node4 := &node{data:11}
	node5 := &node{data:7} 
	node6 := &node{data:2} 
	
	mylist.Prepend(node1)
	mylist.Prepend(node2)
	mylist.Prepend(node3)
	mylist.Prepend(node4) 
	mylist.Prepend(node5)
	mylist.Prepend(node6)
	
fmt.Println("Before Deletion : ")
mylist.PrintListData()
fmt.Println()

mylist.DeleteWithValue(100) 
mylist.DeleteWithValue(2)

fmt.Println("After Deletion : ") 
mylist.PrintListData()


fmt.Println("Linked List Details : ", mylist) 


}