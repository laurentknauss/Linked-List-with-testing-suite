
package main

import (
	"os"
	"fmt"
	"testing" 
	"strings"
)

func TestPrepend(t *testing.T) { 
	 
	var mylist = linkedList{} 
	var node1 = &node{data:48}

	mylist.prePend(node1) 

	if mylist.head != node1 { 
		t.Errorf("Prepend failed: expected head to be node1, got %+v", mylist.head)
	}
}

func TestDeleteWithValue(t *testing.T) { 
	var mylist  = linkedList{}  
	var node1  = &node{data:48} 
	var node2 = &node{data:18}

	mylist.prePend(node1) 
	mylist.prePend(node2) 

	mylist.deleteWithValue(18) 
	if mylist.head.data == 18 { 
		t.Errorf("DeleteWithValue failed - value 18 was not deleted rom the list")  

	}
}

func TestPrintListData(t *testing.T) { 
	var mylist = linkedList{} 
	var node1 = &node{data:48}
	var node2 = &node{data:18} 

	mylist.prePend(node1) 
	mylist.prePend(node2) 

	writer  := new(strings.Builder) 
	oldStdout  := os.Stdout 
	
	fmt.Fprintln(writer, "18 48") 


	os.Stdout = oldStdout



	var expectedOutput = "18 48"  
	var actualOutput = writer.String() 
	actualOutput = strings.TrimSuffix(actualOutput, "\n") 
	if actualOutput !=expectedOutput { 
		t.Errorf("PrintListData failed: expectedOutput '%s', but got '%s' ", expectedOutput, actualOutput)  
	}


}


// Run all the tests . 
func TestMain(m *testing.M) { 
	// Run tests 
	var result = m.Run()  

	// Exit with the result .  
	os.Exit(result) 

}