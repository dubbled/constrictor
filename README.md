# constrictor

Facilitates throttling of the following pipelines:

    1. From io.Reader to []byte
    
    2. From io.Reader to io.Writer

#### Examples
(excluded error checking)

###### Reader to Writer
```
func main() {
    response, _ := http.Get("google.com")
    defer response.Close()

    file, _ := os.OpenFile("google.html", os.O_WRONLY, 0644)
    defer file.Close()

    // read 1000 bytes per second
    constrict.NewReader(response.Body, 1000).WriteTo(file)
}
```

###### Reader to Byte Slice ([]byte)
```
func main() {
    // read 3000 bytes per second
    input := make([]byte, 1024 * 1024)
    constrict.NewReader(os.Stdin, 3000).Read(input)
}
```
