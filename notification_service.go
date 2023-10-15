package main
import (
	"fmt"
	"os"
	"bufio"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Mess struct{
	Status string `json:"status"`
	Info string `json:"info"`
}

func handlerSendMess(w http.ResponseWriter, r *http.Request){
	fmt.Println("Reading input message")
	respBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var ms Mess
	err = json.Unmarshal(respBody,&ms)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("<"+ms.Info+">")
	fmt.Println("Reading input message correctly complited")
}

func main(){
	fmt.Println("Notification server start to work")
	defer fmt.Println("Notification server stop wroking")
	http.HandleFunc("/", handlerSendMess)
	go http.ListenAndServe("127.0.0.1:8084", nil)
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if text == "exit\n"{
			return
		}
	}
}