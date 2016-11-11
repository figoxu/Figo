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

//mtool connect  -conn 127.0.0.1:6000 -id id.txt # make connections, read id from id.txt
//mtool connect  -conn 127.0.0.1:6000 -n 100 # make 100 connections, use sequence as id
//mtool genid -n 100 # generate and print 100 sequence id
//mtool exportid -aerospike AEROSPIKE_URL -tp a -app 53bd01a459ba0740be0000dc # export dv from aerospike
//result: id, bundle, token, tag, av, sdk, lg, la, mt
//mtool online -conn 127.0.0.1:6001 -mongo MONGO_URL -app 53bd01a459ba0740be0000dc
//mtool rtest -redis 127.0.0.1:6379 -n 10 -data abcd
//mtool push_stat -mongo MONGO_URL -app 53bd01a459ba0740be0000dc -start 20150617 -end 20150618
//result: app push_id start_date end_date os total sent read
//mtool dv_stat -aerospike AEROSPIKE_URL -app 53bd01a459ba0740be0000dc -date 20150617
//result: app date android_all android_mau android_dau ios_all ios_mau ios_dau
//
//mtool dv_sync -aerospike AEROSPIKE_URL -ro=<true/flase> -mongo MONGO_URL -redis REDIS_HOST -redis_pass REDIS_PASSWORD
//remark: -ro :readOnly option ,use to check the read speed of aerospike without write to redis ,Default value is false
//
//
//mtool dv_syncAs2Ca -aerospike AEROSPIKE_URL -ca CA_URL -mongo MONGO_URL
//
//mtool clear_redis -aerospike AEROSPIKE_URL -ro=<true/flase> -mongo MONGO_URL -redis REDIS_HOST -redis_pass REDIS_PASSWORD
//remark: -ro :readOnly option ,use to check the read speed of aerospike without dequeue from redis ,Default value is false
//
//mtool exportdv -app 551221b84ab88d261700b673,5386a17859ba0740ce00001b,551221404ab88d261700b66f -start 2015011900 -end 20160120024 -redis 10.10.216.166:6379 -redis_pass testpwd
//
//
//
//mtool dailyReport -aerospike AEROSPIKE_URL -mongo MONGO_URL
