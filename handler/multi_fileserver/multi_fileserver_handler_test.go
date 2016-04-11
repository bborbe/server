package multi_fileserver

import (
	"net/http"
	"testing"
	server_mock "github.com/bborbe/server/mock"

	. "github.com/bborbe/assert"
	"os"
	"io/ioutil"
	"path"
	"fmt"
)

func TestImplementsHandler(t *testing.T) {
	r := NewMultiFileserverHandler(http.Dir("/tmp"))
	var i *http.Handler
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestServeHTTP404(t *testing.T) {
	dir1, _ := ioutil.TempDir("", "dir1")
	defer os.RemoveAll(dir1)
	dir2, _ := ioutil.TempDir("", "dir2")
	defer os.RemoveAll(dir2)
	writeFile(dir1, "a.txt", "a1")
	writeFile(dir2, "b.txt", "b2")
	h := NewMultiFileserverHandler(http.Dir(dir1), http.Dir(dir2))
	response := server_mock.NewHttpResponseWriterMock()
	request, err := server_mock.NewHttpRequestMock("http://www.example.com/foo.txt")
	if err != nil {
		t.Fatal(err)
	}
	h.ServeHTTP(response, request)
	if err = AssertThat(response.Status(), Is(404)); err != nil {
		t.Fatal(err)
	}
}

func TestServeHTTPDir1(t *testing.T) {
	dir1, _ := ioutil.TempDir("", "dir1")
	defer os.RemoveAll(dir1)
	dir2, _ := ioutil.TempDir("", "dir2")
	defer os.RemoveAll(dir2)
	writeFile(dir1, "a.txt", "a1")
	writeFile(dir2, "b.txt", "b2")
	h := NewMultiFileserverHandler(http.Dir(dir1), http.Dir(dir2))
	response := server_mock.NewHttpResponseWriterMock()
	request, err := server_mock.NewHttpRequestMock("http://www.example.com/a.txt")
	if err != nil {
		t.Fatal(err)
	}
	h.ServeHTTP(response, request)
	if err = AssertThat(response.Status(), Is(200)); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(string(response.Bytes()), Is("a1")); err != nil {
		t.Fatal(err)
	}
}

func TestServeHTTPDir2(t *testing.T) {
	dir1, _ := ioutil.TempDir("", "dir1")
	defer os.RemoveAll(dir1)
	dir2, _ := ioutil.TempDir("", "dir2")
	defer os.RemoveAll(dir2)
	writeFile(dir1, "a.txt", "a1")
	writeFile(dir2, "b.txt", "b2")
	h := NewMultiFileserverHandler(http.Dir(dir1), http.Dir(dir2))
	response := server_mock.NewHttpResponseWriterMock()
	request, err := server_mock.NewHttpRequestMock("http://www.example.com/b.txt")
	if err != nil {
		t.Fatal(err)
	}
	h.ServeHTTP(response, request)
	if err = AssertThat(response.Status(), Is(200)); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(string(response.Bytes()), Is("b2")); err != nil {
		t.Fatal(err)
	}
}

func TestServeHTTPFirstWins(t *testing.T) {
	dir1, _ := ioutil.TempDir("", "dir1")
	defer os.RemoveAll(dir1)
	dir2, _ := ioutil.TempDir("", "dir2")
	defer os.RemoveAll(dir2)
	writeFile(dir1, "a.txt", "a1")
	writeFile(dir2, "a.txt", "a2")
	h := NewMultiFileserverHandler(http.Dir(dir1), http.Dir(dir2))
	response := server_mock.NewHttpResponseWriterMock()
	request, err := server_mock.NewHttpRequestMock("http://www.example.com/a.txt")
	if err != nil {
		t.Fatal(err)
	}
	h.ServeHTTP(response, request)
	if err = AssertThat(response.Status(), Is(200)); err != nil {
		t.Fatal(err)
	}
	if err = AssertThat(string(response.Bytes()), Is("a1")); err != nil {
		t.Fatal(err)
	}
}

func writeFile(dir string, name string, content string) error {
	fmt.Printf("write dir %s name %s\n", dir, name)
	filename := path.Join(dir, name)
	fmt.Printf("write file %s\n", filename)
	return ioutil.WriteFile(filename, []byte(content), 0644)
}
