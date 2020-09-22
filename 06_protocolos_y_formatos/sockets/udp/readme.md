# Test the UDP Client and Server

You can now test your UDP client and server. You will need to execute the UDP server first so that the UDP client has somewhere it can connect to.

1.  Run your UDP server. From the directory containing the `udpservidor.go` file, run the following command:

    `go run udpservidor.go 1234`

    The server will listen on port number `1234`. You will not see any output as a result of this command.

2.  Open a second shell session to execute the UDP client and to interact with the UDP server. Run the following command:

    `go run udpcliente.go 127.0.0.1:1234`


3.  You will see a `>>` prompt waiting for you to enter some text. Type in `Hello!` to receive a response from the UDP server:

    > Hello!

    You should see a similar output:

    > The UDP server is 127.0.0.1:1234
    >
    > \>> Hello!
    >
    > Reply: 82

4.  Send the `STOP` command to exit the UDP client and server:

    > STOP

    You should see a similar output on the client side:

    > \>> STOP
    >
    > ->: Exiting UDP client!

    The output on the UDP server side will be as follows:

    > -> STOP
    >
    > Exiting UDP server!
