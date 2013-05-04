package server


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

    backlog := make( chan []byte, 10 );

    go LogicHandle( conn, backlog );

    header := make( []byte, 1 );

    for{

        n, err := io.ReadFull( conn, header );

        if n == 0 && err == io.EOF{
            fmt.Printf("client closed \n");
            break
        }else if err != nil {
            fmt.Println("error receving header:", err)
			break
        }

        backlog <- header

    }
}

func LogicHandle( conn net.Conn, backlog chan []byte ){

    for{
        select{
        case <- backlog :
            res := []byte("hi\n");
            conn.Write( res );

        }
    }

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

