package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/tokopedia/gripmock/stub"
)

const (
	defaultImportsPath = "/protobuf"
	defaultStubPath = "/stubs"
	defaultProtoPath = "/proto/"
)

func main() {
	outputPointer := flag.String("o", "", "directory to output server.go. Default is $GOPATH/src/grpc/")
	grpcPort := flag.String("grpc-port", "4770", "Port of gRPC tcp server")
	grpcBindAddr := flag.String("grpc-listen", "", "Adress the gRPC server will bind to. Default to localhost, set to 0.0.0.0 to use from another machine")
	adminport := flag.String("admin-port", "4771", "Port of stub admin server")
	adminBindAddr := flag.String("admin-listen", "", "Adress the admin server will bind to. Default to localhost, set to 0.0.0.0 to use from another machine")
	stubPath := flag.String("stub", defaultStubPath, "Path where the stub files are (Optional)")
	imports := flag.String("imports", defaultImportsPath, "comma separated imports path. default path /protobuf is where gripmock Dockerfile install WKT protos")
	// for backwards compatibility
	if os.Args[1] == "gripmock" {
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}

	flag.Parse()
	fmt.Println("Starting GripMock")

	output := *outputPointer
	if output == "" {
		if os.Getenv("GOPATH") == "" {
			log.Fatal("output is not provided and GOPATH is empty")
		}
		output = os.Getenv("GOPATH") + "/src/grpc"
	}

	// for safety
	output += "/"
	if _, err := os.Stat(output); os.IsNotExist(err) {
		os.Mkdir(output, os.ModePerm)
	}

	// run admin stub server
	stub.RunStubServer(stub.Options{
		StubPath: *stubPath,
		Port:     *adminport,
		BindAddr: *adminBindAddr,
	})

	// parse proto files
	protoPaths := flag.Args()
	if len(protoPaths) == 0 {
		protoPaths = append(protoPaths, defaultProtoPath)
	} else if len(protoPaths) > 1 {
		log.Fatal("Need only one proto path")
	}

	importDirs := strings.Split(*imports, ",")
	if !strings.Contains(*imports, defaultImportsPath) {
		importDirs = append(importDirs, defaultImportsPath)
	}

	// generate pb.go and grpc server based on proto
	generateProtoc(protocParam{
		protoPath:   protoPaths[0],
		adminPort:   *adminport,
		grpcAddress: *grpcBindAddr,
		grpcPort:    *grpcPort,
		output:      output,
		imports:     importDirs,
	})

	// build the server
	buildServer(output, protoPaths[0])

	// and run
	run, runerr := runGrpcServer(output)

	var term = make(chan os.Signal)
	signal.Notify(term, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	select {
	case err := <-runerr:
		log.Fatal(err)
	case <-term:
		fmt.Println("Stopping gRPC Server")
		run.Process.Kill()
	}
}

func getProtoNameFromFilename(filename string) string {
	return strings.Split(filename, ".")[0]
}

type protocParam struct {
	protoPath   string
	adminPort   string
	grpcAddress string
	grpcPort    string
	output      string
	imports     []string
}

func generateProtoc(param protocParam) {
	if !strings.HasSuffix(param.protoPath, "/") {
		param.protoPath += "/"
	}

	args := []string{"-I", param.protoPath}
	// include well-known-types
	for _, i := range param.imports {
		args = append(args, "-I", i)
	}

	files, err := ioutil.ReadDir(param.protoPath)
	if err != nil {
		log.Fatalf("Can't read proto from dir %s. %v\n", param.protoPath, err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".proto") {
			continue
		}

		log.Printf("found proto file %s", file.Name())
		args = append(args, param.protoPath+file.Name())
	}

	args = append(args, "--go_out=plugins=grpc:"+param.output)
	args = append(args, fmt.Sprintf("--gripmock_out=admin-port=%s,grpc-address=%s,grpc-port=%s:%s",
		param.adminPort, param.grpcAddress, param.grpcPort, param.output))
	log.Printf("run protoc with args %v", args)
	protoc := exec.Command("protoc", args...)
	protoc.Stdout = os.Stdout
	protoc.Stderr = os.Stderr
	err = protoc.Run()
	if err != nil {
		log.Fatal("Fail on protoc ", err)
	}
	log.Print("protoc executed successfully")

	// change package to "main" on generated code
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".proto") {
			continue
		}

		protoName := getProtoNameFromFilename(file.Name()) + ".pb.go"
		log.Printf("go file %s from proto %s", protoName, file.Name())
		sedArgs := []string{"-i", `s/^package \w*$/package main/`, param.output + protoName}
		log.Printf("run sed with args %v", sedArgs)
		sed := exec.Command("sed", sedArgs...)
		sed.Stderr = os.Stderr
		sed.Stdout = os.Stdout
		err = sed.Run()
		if err != nil {
			log.Fatal("Fail on sed")
		}
	}
	log.Print("change package to \"main\" done")
}

func buildServer(output string, protoPath string) {
	exec.Command("go", "get", "gitlab.ozon.ru/map/types").Run()

	args := []string{"build", "-o", output + "grpcserver", output + "server.go"}

	files, _ := ioutil.ReadDir(protoPath)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".proto") {
			continue
		}
		args = append(args, output+getProtoNameFromFilename(file.Name())+".pb.go")
	}
	build := exec.Command("go", args...)
	build.Stdout = os.Stdout
	build.Stderr = os.Stderr
	err := build.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func runGrpcServer(output string) (*exec.Cmd, <-chan error) {
	run := exec.Command(output + "grpcserver")
	run.Stdout = os.Stdout
	run.Stderr = os.Stderr
	err := run.Start()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("grpc server pid: %d\n", run.Process.Pid)
	runerr := make(chan error)
	go func() {
		runerr <- run.Wait()
	}()
	return run, runerr
}
