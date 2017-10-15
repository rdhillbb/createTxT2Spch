package createTxT2Spch
//package main

import (
	"fmt"
	"log"
	"os"
	"io"
	"github.com/mediawen/watson-go-sdk"
	"github.com/rdhillbb/buildWatsonNumber"
)

func CreateWatAudio(direc,fileName,message,wToken,wPasswd string) {
	w := watson.New(wToken,wPasswd)
	a, err := w.Synthesize(message, "en-US_AllisonVoice", "mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer a.Close()
	
	f, err := os.Create(direc+"/"+fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	
	n, err := io.Copy(f, a)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("%s: %d bytes written\n", "ddd", n)
}

func main() {

message := buildWatsonNumber.CrMsg4wat("MBS in Singapore","178890")
fmt.Println("------------>",message)
CreateWatAudio("/tmp","Altert.mp3",message,"afa8b12a-fe4e-4cf3-8194-ce8238ae9834","LKD4D75Jg1GF")
}
