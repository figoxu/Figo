package main

import (
	"flag"
	"fmt"
	"github.com/figoxu/Figo"
	"github.com/gogap/errors"
	"github.com/quexer/utee"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		info()
		return
	}
	log.Println("hello")
	cmd, leftArgs := os.Args[1], os.Args[2:]
	if Figo.Exist(cmd, "hdfs") {
		log.Println("HDFS Command Found")
		fs := flag.NewFlagSet("connect", flag.ContinueOnError)
		nameServer := fs.String("nameserver", "", "The HDFS Name Node Address")
		user := fs.String("user", "", "The HDFS Visitor User")
		read := fs.String("read", "", "The HDFS File Path")
		write := fs.String("write","","The File Which Will Write To HDFS")
		hdfsPath := fs.String("hdfsPath","","The HDFS  Which Will Write To HDFS")
		fs.Parse(leftArgs)
		if *nameServer == "" {
			utee.Chk(errors.New("nameserver Param Is Need"))
		}
		if *user == "" {
			utee.Chk(errors.New("user Param Is Need"))
		}
		client := Figo.NewHDFSClient(*nameServer, *user)
		if *read != "" {
			d, e := client.Read(*read)
			utee.Chk(e)
			log.Println("Data I Read From HDFS Is :")
			log.Println(string(d))
		}else if *write!="" {
			if *hdfsPath==""{
				utee.Chk(errors.New("hdfsPath Param Is Need"))
			}
			f,e:=Figo.FileOpen(*write)
			utee.Chk(e)
			client.WriteFile(*hdfsPath,f)
		}
	} else {
		log.Println("Sorry Sir,Command '", cmd, "' Not Found.")
	}
}

func info() {
	{
		fmt.Println("Action Needed")
		fmt.Println(`examples:
		figo hdfs -nameserver 172.17.0.2:9000 -user root -read /user/root/crc_logfile.2016-11-09
		figo hdfs -nameserver 172.17.0.2:9000 -user root -write /user/root/crc_logfile.2016-11-09 -hdfsPath /user/root/test.txt
	`)
	}

}
