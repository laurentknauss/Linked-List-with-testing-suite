package main

import (
	
	"io"
	"os"
	"strings"
	"testing"
)

func TestPrepend(t *testing.T) {

	var mylist = linkedList{}
	var node1 = &node{data: 48}

	mylist.Prepend(node1)

	if mylist.head != node1 {
		t.Errorf("Prepend failed: expected head to be node1, got %+v", mylist.head)
	}
}

func TestDeleteWithValue(t *testing.T) {
	var mylist = linkedList{}
	var node1 = &node{data: 48}
	var node2 = &node{data: 18}

	mylist.Prepend(node1)
	mylist.Prepend(node2)

	mylist.DeleteWithValue(18)
	if mylist.head.data == 18 {
		t.Errorf("DeleteWithValue failed - value 18 was not deleted from the list")

	}
}

func TestPrintListData(t *testing.T) {
	var mylist = linkedList{}
	var node1 = &node{data: 48}
	var node2 = &node{data: 18}

	mylist.Prepend(node1)
	mylist.Prepend(node2) 

  // Temporarily re-direct os.Stdout to a pipe . 
	oldStdout := os.Stdout
	r,w , _ := os.Pipe()
	os.Stdout = w

	// Calling 'PrintListData' which should now write to our pipe . 
	mylist.PrintListData()


// Closing the write-end of the pipe, read the output, and restore os.Stdout . 
w.Close()
os.Stdout = oldStdout
capturedOutput, _ := io.ReadAll(r)
r.Close()

// Converting captured output to a string and trim any trailing newlines . 
actualOutput := strings.TrimSuffix(string(capturedOutput), "\n") 

// Defining the expected output . 
expectedOuput := "18 48 " 

// Comparing the actual output to the expected output . 
if actualOutput != expectedOuput { 
	t.Errorf("PrintListData failed: expected output '%s', but got '%s'", expectedOuput, actualOutput)
}
}