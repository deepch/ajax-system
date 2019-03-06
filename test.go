package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	CONN_HOST = ""
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen("tcp", ":33333")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		l, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		//http://s545463982.onlinehome.us/DC09Gen/
		sdata := string(buf[:l])
		//https://www.eca-vaud.ch/images/prevention/pdf/Standard_Transm/451170_complement_technique1_IP.pdf
		log.Println(string(buf[:l]))
		if strings.Contains(sdata, "Nri1/CL501") {
			log.Println("Поставленно на охрану")
		} else if strings.Contains(sdata, "Nri1/OP501") {
			log.Println("Снято с охраны")
		} else if strings.Contains(sdata, "Nri1/BA") {
			log.Println("Начало тревоги")

		} else if strings.Contains(sdata, "Nri1/BR") {
			log.Println("Конец тревоги")
		}

		responce := []byte{10}
		responce = append(responce, buf[0])
		responce = append(responce, []byte(`0B430026"ACK"0001L0#1111[]_10:36:48,03-06-2019`)...)
		responce = append(responce, []byte{13}...)

		conn.Write(responce)

		responce = []byte{10}
		responce = append(responce, []byte(`9D290026"NAK"0006L0#1111[]_10:46:32,03-06-2019`)...)
		responce = append(responce, []byte{13}...)

		conn.Write(responce)
		/*
			log.Println("================", string(buf[:l]))
			log.Println("LEN", l)
			log.Println("START", buf[0])
			log.Println("CRC16", string(buf[1:5]))
			log.Println("LEN", string(buf[5:9]))
			log.Println("ID Token", l)
			log.Println("Sequence Number", l)
			log.Println("Receiver Number", l)
			log.Println("Account Prefix", l)
			log.Println("Account Number", l)
			log.Println("Message Data", l)
			log.Println("Extended Data", l)
			log.Println("Timestamp", l)
			log.Println("END", buf[l])
			log.Println("================")
		*/
		//	conn.Write(<LF>A2940012"ACK"0001L0#1111[]<CR>

		//	time.Sleep(100 * time.Millisecond)

		//CL501 - поставленно на охрану
		//OP501 - снято с охраны

	}
}
