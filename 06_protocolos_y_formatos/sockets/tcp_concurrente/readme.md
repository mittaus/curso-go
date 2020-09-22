# Test the Concurrent TCP Server

In this section, you will test the concurrent TCP server using the netcat command line utility.


1.  Run your concurrent TCP server. From the directory containing the `servidorTCPConcurrente.go` file, run the following command:

    `go run servidorTCPConcurrente.go 1234`

    The command creates a TCP server that listens on port number `1234`. You can use any port number, however, ensure it is not already in use and that you have the required privileges. Reference the list of well-known TCP and UDP ports, if needed. https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports

2.  Use netcat to establish a connection with the TCP server. By default, netcat will establish a TCP connection with a remote host on the specified port number.:

    `nc 127.0.0.1 1234`


3.  After issuing the previous command, you will not see any change in your output. Type `Hello!` to send a packet to the TCP server:

    > Hello!

    The TCP server will return the number of current client connections as its response. Since this is your first connection established with the TCP server, you should expect an output of `1`.

    > Hello!
    >
    > Reply: 1

    If you’d like, you can open a new shell session and use netcat to establish a second connection with the TCP server by repeating Step 2. When you send the server a second `Hello!`, you should receive a response of `2` this time.

4.  You can also connect to the TCP server using the TCP client you created in the Create the TCP Client section of the guide. Ensure you are in the directory containing the tcpC.go file and issue the following command:

    > `go run tcpcliente.go 127.0.0.1:1234`


5. You will see a `>>` prompt waiting for you to enter some text. Type in Hello! to receive a response from the TCP server:

    > Hello!

    You should see a similar output indicating `3` client connections:

    > \>> Hello!
    >
    > ->: 3

6.  Send the `STOP`  command to exit the TCP client:

    You should see a similar output on the client:

    > \>> STOP
    >
    > ->: TCP client exiting...

    The output on the TCP server side will be as follows:


    > .Hello!
    >
    > .Hello!
    >
    > .Hello!

> 
> **Note**
> From the shell session running the TCP server, type **CTRL-c** to interrupt program execution and then, **CTRL-D** to close all client connections and to stop the TCP server.
