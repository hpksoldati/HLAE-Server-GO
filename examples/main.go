package main

import (
	"fmt"
	"github.com/FlowingSPDG/HLAE-Server-GO"
	"github.com/c-bata/go-prompt"
)

var (
	hlaeserver = &mirvpgl.HLAEServer{}
)

func init() {
	hlaeserver = &mirvpgl.HLAEServer{}
}

// ExampleHandler for HLAE Server
func ExampleHandler(cmd string) {
	fmt.Printf("Received %s\n", cmd)
}

// ExampleCamHandler for cam datas
func ExampleCamHandler(cam *mirvpgl.CamData) {
	fmt.Printf("Received cam data %v\n", cam)
}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func main() {
	hlaeserver.RegisterHandler(ExampleHandler)
	hlaeserver.RegisterCamHandler(ExampleCamHandler)
	go hlaeserver.Start(":65535", "/mirv")
	for {
		cmd := prompt.Input("CSGO >>> ", completer)
		hlaeserver.SendRCON(cmd)
	}
}
