package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/helunxing/energy-consumption-instrumentation/go/method"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "1966"
	CONN_TYPE = "tcp"
	PRO_ADDR  = "/Users/h/code/energy-consumption-instrumentation/"
)

var IO_time_list []float64

func main() {
	var cPU_time_list []int64
	jsonFile, err := os.Open(pathGen("control.json"))
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var control map[string]map[string]float32
	json.Unmarshal([]byte(byteValue), &control)

	end := len(control) - 1
	for idx := 0; idx < end; idx++ {
		num := control[strconv.Itoa(idx)]["times"]
		cPU_time := method.ExeAndPrintCPUusage(server, int(num))
		// fmt.Println("CPU using time (CPU_time): ", cPU_time)
		cPU_time_list = append(cPU_time_list, cPU_time)
		// fmt.Println()
		fmt.Println("go server done no. ", idx)
		time.Sleep(time.Second)
	}
	fmt.Println(IO_time_list)
	fmt.Println(cPU_time_list)
}

func server(num int) {

	start := time.Now()
	jsonFile, err := os.Open(pathGen("datas.json"))
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	elapsed := time.Since(start)
	// fmt.Println("reading datas file using time (IO_time):", elapsed)
	IO_time_list = append(IO_time_list, float64(elapsed))

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var datas map[string]string
	json.Unmarshal([]byte(byteValue), &datas)

	config := &net.ListenConfig{Control: reusePort}
	l, err := config.Listen(context.Background(), CONN_TYPE, CONN_HOST+":"+CONN_PORT)

	if err != nil {
		fmt.Println("net.Listen err=", err.Error())
		os.Exit(1)
	}

	for i := 0; i < num; i++ {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("net.Listen.Accept", err.Error())
			os.Exit(1)
		}
		handle(conn, datas)
	}

	defer l.Close()

}

func handle(conn net.Conn, datas map[string]string) {
	buf := make([]byte, 1050)

	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	tmp := string(buf[:reqLen])
	index := strings.Index(tmp, "\r\n")

	paths := strings.Split(tmp[:index], "/")
	reqTar := paths[len(paths)-2]
	var res string = "20 "
	res += datas[reqTar]
	res += "\r\n"
	conn.Write([]byte(res))

	defer conn.Close()
}

func pathGen(s string) string {
	return filepath.Join(PRO_ADDR, s)
}

func reusePort(network, address string, conn syscall.RawConn) error {
	return conn.Control(func(descriptor uintptr) {
		syscall.SetsockoptInt(int(descriptor), syscall.SOL_SOCKET, syscall.SO_REUSEPORT, 1)
	})
}
