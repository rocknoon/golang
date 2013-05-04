package http


import "net"
import "fmt"
import "io"
import "os"

const PORT = ":8080"

func Start(){

    tcpAddr,err := net.ResolveTCPAddr( "tcp4" , PORT );

    checkError( err );

    tcpListener,err := net.ListenTCP( "tcp" , tcpAddr );

    checkError( err );

    fmt.Printf("cool server start \n");

    for{

        conn , err := tcpListener.Accept();

        fmt.Printf( "accpet one connection! \n"  );


        checkError( err );

        go ConnHandle( conn  )


    }



}

func ConnHandle( conn net.Conn  ){

    defer conn.Close();
    defer fmt.Printf("close one client \n");

    header := make( []byte, 5000 );

    n, err := conn.Read(header);

    if n == 0{
        fmt.Printf("no header \n");
        return;
    }

    if err != nil && err != io.EOF {
        fmt.Println("error receving header:", err)
		return;
    }

    response := []byte("HTTP/1.1 200 OK\r\nServer: Golang/1.2.1\r\nDate: Sat, 04 May 2013 03:08:51 GMT\r\n\r\nHello World\r\n");

    conn.Write(response);

}


func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

