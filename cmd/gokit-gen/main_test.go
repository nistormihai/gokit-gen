package main

import (
	"bytes"
	"log"
	"testing"
)

func Test(t *testing.T) {

	const (
		templatesDir = "./../../templates"
		deployTo     = "/home/badu/go.workspace/go-kit-gen-pub/src/github.com/badu/gokit-gen/example/poc/gen/"
		protoFile    = "/home/badu/go.workspace/go-kit-gen-pub/src/github.com/badu/gokit-gen/example/poc/pb/example.proto"
		packageName  = "github.com/badu/gokit-gen/example/poc/gen"
	)
	/**
	const (
		templatesDir = "./../../templates"
		deployTo     = "/home/badu/go.workspace/go-kit-gen-pub/src/github.com/badu/gokit-gen/example/http/gen/"
		protoFile    = "/home/badu/go.workspace/go-kit-gen-pub/src/github.com/badu/gokit-gen/example/http/pb/factory.proto"
		packageName  = "github.com/badu/gokit-gen/example/http/gen"
	)
	const (
		templatesDir = "./../../templates"
		deployTo     = "/home/badu/go.workspace/go-kit-gen-pub/src/github.com/badu/gokit-gen/example/route_guide/gen/"
		protoFile    = "/home/badu/go.workspace/go-kit-gen-pub/src/github.com/badu/gokit-gen/example/route_guide/routeguide/route_guide.proto"
		packageName  = "github.com/badu/gokit-gen/example/route_guide/gen"
	)
	const (
		templatesDir = "./../../templates"
		deployTo     = "/home/badu/go.workspace/go-kit-gen-pub/src/github.com/badu/gokit-gen/example/poc/gen/"
		protoFile    = "/home/badu/go.workspace/go-kit-gen-pub/src/github.com/badu/gokit-gen/example/poc/pb/example.proto"
		packageName  = "github.com/badu/gokit-gen/example/poc/gen"
	)
	**/
	log.SetFlags(0)
	log.SetPrefix("gokit-gen: ")

	args := []string{"", protoFile, templatesDir, packageName, deployTo}

	var stdout bytes.Buffer
	//log.Print("Run")
	err := run(args, &stdout)
	if err != nil {
		t.Fatalf("error : %v", err)
	}
	//log.Print("Done")
}
