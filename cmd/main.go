package main

import (
	"crud/pkg/postgres"
	"crud/src/apiserver"
	"fmt"
	"os"
)

func main() {

	///// Pid Kaydı yanlış kapama yardımcısı
	logfile, fileerr := os.OpenFile("../src/apiserver/pid.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if fileerr != nil {
		fmt.Println("Dosya açılamadı DİKKAT")
	}
	logfile.WriteString("server pid : " + fmt.Sprint(os.Getpid()) + "\n")
	logfile.Close()
	///////////////////////////////

	err := apiserver.New(
		apiserver.WithPort("8080"),
		apiserver.WithDatabase(postgres.ConnectPQ()),
	)
	if err != nil {
		fmt.Println("test çıktı hatası ", "\"", err, "\"")
		return
	}
	fmt.Println("Server düzgün bir şekilde kapandı")

}
