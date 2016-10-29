package main

import (
    "fmt"
    "net"
    "time"
    _ "strconv"
)

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}

func main() {
    ServerAddr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
    CheckError(err)

    LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
    CheckError(err)

    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)

    defer Conn.Close()

	values := make([]uint16, 3);
	values[0] = 100
	values[1] = 1024
	values[2] = 499
    for {

		buf := shortToBytes(values)
		_,err := Conn.Write(buf)
		if err != nil {
			fmt.Println("Unable to send UDP package", err)
		}

		/*
        msg := strconv.Itoa(i)
        i++
        buf := []byte(msg)
        _,err := Conn.Write(buf)
        if err != nil {
            fmt.Println(msg, err)
        }
        */
        time.Sleep(time.Second * 1)
    }
}

//conversion os C++ code from The ESP program
func shortToBytes(shorts []uint16) []byte {
	fmt.Println(shorts)
	l := len(shorts)
	bytes := make([]byte,l * 2)

	for i := 0; i < l; i++ {
		bytes[i*2] = byte(shorts[i] & 0x00FF)
		bytes[i*2 +1] = byte(shorts[i] >> 8)
	}
	return bytes
}
