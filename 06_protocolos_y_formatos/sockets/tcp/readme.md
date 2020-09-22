# Test the TCP Client and Server

You can now test your TCP client and server. You will need to execute the TCP server first so that the TCP client has somewhere it can connect to.

1.  Run your TCP server. From the directory containing the tcpservidor.go file, run the following command:

    `go run tcpservidor.go 1234`

    The server will listen on port number `1234`. You will not see any output as a result of this command.

2.  Open a second shell session to execute the TCP client and to interact with the TCP server. Run the following command:

    `go run tcpcliente.go 127.0.0.1:1234`

    > **Note**
    >
    > <div>
    >
    > If the TCP server is not running on the expected TCP port, you will get the following error message from `tcpcliente.go`:
    >
    >     dial tcp [::1]:1234: connect: connection refused
    >
    > </div>

3.  You will see a `>>` prompt waiting for you to enter some text. Type in Hello! to receive a response from the TCP server:

    > Hello!

    You should see a similar output:

    > \>> Hello!
    >
    > ->: 2019-05-23T19:43:21+03:00

4.  Send the `STOP` command to exit the TCP client and server:

    > STOP

    You should see a similar output in the client:

    > \>> STOP
    >
    > ->: TCP client exiting...

    The output on the TCP server side will resemble the following:

    > -> Hello!
    >
    > Exiting TCP server!

> 
> **Note**
> The TCP server waits before writing back to the TCP client, whereas the client writes to the TCP server first and then waits to receive an answer. This behavior is part of the protocol definition that governs a TCP or a UDP connection. In this example, you have implemented an unofficial protocol that is based on TCP.
