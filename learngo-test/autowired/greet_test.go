package autowired

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestGreet(t *testing.T)  {
	buffer := bytes.Buffer{}
	Greet(&buffer, "chris")

	got := buffer.String()
	want := "hello, chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "hello, %s", name)
}

func main()  {
	Greet(os.Stdout, "Elodie")
}