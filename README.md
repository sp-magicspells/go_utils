# go_utils
Some tiny golang utilities which can be used for debugging purpose

```func main() {
	listener, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	defer listener.Close()
	port := listener.Addr().(*net.TCPAddr).Port
	fmt.Printf("port found: %d\n", port)
}
```

Above utily in main.go can be used to debug port not available issue most of developers experience. you can run it by `go run main.go` and if successful, it should print a random port number instead of failing with an error.

Some more utils will be added soon.
