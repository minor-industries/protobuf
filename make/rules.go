package main

import (
	"fmt"
	"github.com/minor-industries/grm"
	"path/filepath"
)

var rules = map[string]func(rule string){
	"protos": func(rule string) {
		protoFiles, err := filepath.Glob("protos/*/*.proto")
		noError(err)

		args := []string{
			"/bin/protoc",
			"--proto_path=./protos/heads",
			"-I/build/include",
			"--go_out=plugins=grpc,paths=source_relative:./gen/go/heads",
		}

		for _, file := range protoFiles {
			// this may run into trouble if there are two proto files with the same name in
			// different directories
			base := filepath.Base(file)
			opt := fmt.Sprintf("--go_opt=M%s=github.com/minor-industries/theheads/gen/go/heads", base)
			args = append(args, opt)
		}

		args = append(args, protoFiles...)

		grm.RunDocker("heads-protoc", args...)
	},

	"heads-protoc": func(rule string) {
		grm.Cd("common", func() {
			grm.Run(nil, "docker", "build", "--tag", "heads-protoc", "protos")
		})
	},
}

func noError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	grm.Main(rules)
}
