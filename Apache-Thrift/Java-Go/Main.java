package com.soofaloofa.alice;

import com.soofaloofa.taskmanager.Task;
import com.soofaloofa.taskmanager.TaskManager;
import org.apache.thrift.TException;
import org.apache.thrift.protocol.TProtocol;
import org.apache.thrift.protocol.TSimpleJSONProtocol;
import org.apache.thrift.transport.TSocket;
import org.apache.thrift.transport.TTransport;

public class Main {

  public static void main(String[] args) throws TException {
    // Open a transport to the server
    TTransport transport = new TSocket("localhost", 9090);
    transport.open();

    // Declare the protocol to use
    TProtocol protocol = new TSimpleJSONProtocol(transport);

    // Create a client using the protocol
    TaskManager.Client client = new TaskManager.Client(protocol);

    // Create a task
    Task task = new Task();
    task.setIdentifier(42);
    task.setDescription("Close Jira ticket PR-12345");

    // Send it to the server
    long result = client.create(task);
    System.out.println("Processed task " + result);

    transport.close();
  }
}
