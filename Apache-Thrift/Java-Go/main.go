package main

import (
	"fmt"
	"github.com/amirhallaji/My-Learning/Apache-Thrift"
	"github.com/amirhallaji/My-Learning/Apache-Thrift/thrift
//	"git.apache.org/thrift.git/lib/go/thrift"
//	"github.com/soofaloofa/taskmanager/gen-go/taskmanager"
)

// Maintain a global list of tasks to work on
var tasks []taskmanager.Task

// TaskManagerHandler processes incoming tasks
type TaskManagerHandler struct{}

// Create adds tasks to our global list
func (t *TaskManagerHandler) Create(task *taskmanager.Task) (r int64, err error) {
	// task is a struct as defined by our Thrift IDL definition.
	// here we simply add it to our list of work
	tasks = append(tasks, *task)

	// Return the response
	return int64(task.GetIdentifier()), nil
}

func main() {
	addr := "localhost:9090"

	// Declare the serialization protocol
	protocolFactory := thrift.NewTSimpleJSONProtocolFactory()

	// Declare the transport method to use
	transportFactory := thrift.NewTBufferedTransportFactory(8192)

	// Get a tansport
	transport, _ := thrift.NewTServerSocket(addr)

	// Implements the interface to our service
	handler := &TaskManagerHandler{}

	// Tell the Thrift processor which interface implementation to use
	processor := taskmanager.NewTaskManagerProcessor(handler)

	// Start the server
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", addr)
	server.Serve()
}
