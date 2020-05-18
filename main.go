package main

import (
	"fmt"
	"log"

	"github.com/programmer-richa/protocol_buffer_code_separate_package/go_format"
	"google.golang.org/protobuf/proto"
)

func main() {
	richa := &go_format.Person{
		Name:  "Richa",
		Email: "programmer_richa@gmail.com",
		Id:    101,
		Followers: &go_format.SocialFollowers{
			Twitter: 1400,
			Youtube: 2000,
		},
	}
	data, err := proto.Marshal(richa)
	if err != nil {
		log.Fatal("Marshalling error", err)
	}
	fmt.Println(data)

	newPerson := &go_format.Person{}
	err = proto.Unmarshal(data, newPerson)
	if err != nil {
		log.Fatal("Unmarshalling error", err)
	}
	fmt.Println("New Person data ", newPerson)
	fmt.Println("ID:", newPerson.GetId())
	fmt.Println("Name:", newPerson.GetName())
	fmt.Println("E-mail:", newPerson.GetEmail())
	fmt.Println("Twitter:", newPerson.GetFollowers().GetTwitter())
	fmt.Println("Youtube:", newPerson.GetFollowers().GetYoutube())
}
