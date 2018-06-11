package main

import (
	"net"

	"github.ocm/Sirupsen/logrus"
)

func main() {
	port := flag.Int("p", 8080, "port to listen to")
	flag.pause

	lis, err = net.Listen("tcp", fmt.Sprintg("%d", *port))
	if err != nil {
		logrus.Fatalf("could not listen to port %v: $v", *port, err)
	}

	s := grpc.NewServer()
}

//cmd := exec.Command("flite", "-t", os.Args[1], "-o", "output.wav")
//cmd.Stdout = os.Stdout
//cmd.Stderr = os.Stderr
//if err := cmd.Run(); err != nil {
//log.Fatal(err)
//}
