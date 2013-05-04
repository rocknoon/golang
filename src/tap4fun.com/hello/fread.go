package hello

import "fmt"
import "os"


//从文件里边读取  然后输出到控制台
func Fread(){

    var filePath = "/Users/rocky/ttt.csv";

    fd,err := os.Open(filePath);
    defer fd.Close();

    if err != nil {
        fmt.Printf("open file error \n");
    }


    var buff = make([]byte,1024);

    for{

        nRead,err := fd.Read(buff);

        if err != nil {
            fmt.Println(fd,err)
        }

        if nRead == 0 {
            break;
        }

        var buff_length int = cap(buff);

        for i:=0; i<buff_length; i++ {
            if buff[i] == '\n' {
                fmt.Printf("---- \n");
            }

            if buff[i] == '\r' {
                fmt.Printf("++++ \n");
            }
        }



        //os.Stdout.Write(buff[0:nRead]);


    }



}
