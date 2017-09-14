//+build ignore

package main

import (
    "log"

    "github.com/golang/protobuf/proto"
    example "." 
)

func main() {
    test := &example.Test {
        Label: proto.String("hello"),
        Type:  proto.Int32(17),
        Reps:  []int64{1, 2, 3},
        Optionalgroup: &example.Test_OptionalGroup {
            RequiredField: proto.String("good bye"),
        },
    }
    data, err := proto.Marshal(test)
    if err != nil {
        log.Fatal("marshaling error: ", err)
    }
    newTest := &example.Test{}
    err = proto.Unmarshal(data, newTest)
    if err != nil {
        log.Fatal("unmarshaling error: ", err)
    }
    // Now test and newTest contain the same data.
    if test.GetLabel() != newTest.GetLabel() {
        log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
    }
    // etc.
}

/*
https://github.com/golang/protobuf

protoc --go_out=plugins=grpc,import_path=mypackage:. *.proto
import_prefix=xxx - a prefix that is added onto the beginning of all imports. Useful for things like generating protos in a subdirectory, or regenerating vendored protobufs in-place.
import_path=foo/bar - used as the package if no input files declare go_package. If it contains slashes, everything up to the rightmost slash is ignored.
plugins=plugin1+plugin2 - specifies the list of sub-plugins to load. The only plugin in this repo is grpc.
Mfoo/bar.proto=quux/shme - declares that foo/bar.proto is associated with Go package quux/shme. This is subject to the import_prefix parameter.

gRPC Support

If a proto file specifies RPC services, protoc-gen-go can be instructed to generate code compatible with gRPC (http://www.grpc.io/). To do this, pass the plugins parameter to protoc-gen-go; the usual way is to insert it into the --go_out argument to protoc:

protoc --go_out=plugins=grpc:. *.proto
*/