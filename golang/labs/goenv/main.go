// Refer to a blog https://dave.cheney.net/2015/11/29/a-whirlwind-tour-of-gos-runtime-environment-variables
// use GOTRACEBACK to get more details when there is a panic
// more details about GOTRACEBACK change: https://github.com/golang/go/commit/bf1de1b141b6354874780401d4525b3b5a1ff6d5
/*
    GOTRACEBACK=none will suppress all tracebacks, you only get the panic message.
    GOTRACEBACK=single is the new default behaviour that prints only the goroutine believed to have caused the panic.
    GOTRACEBACK=all causes stack traces for all goroutines to be shown, but stack frames related to the runtime are suppressed.
    GOTRACEBACK=system is the same as the previous value, but frames related to the runtime are also shown, this will reveal goroutines started by the runtime itself.
    GOTRACEBACK=crash is unchanged from Go 1.5.
eg:
$go build .
$env GOTRACEBACK=all ./goenv
*/

/*
`go help environment` to get all golang environment variable:
`go env` to print all golang env
more details👉 https://pkg.go.dev/runtime#hdr-Environment_Variables
*/
package main

func main() {
	panic("kerboom")
}
