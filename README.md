# constrictor
======

Facilitates throttling of the following pipelines:
    1. From io.Reader to []byte
    2. From io.Reader to io.Writer

======

#### Examples
(excluded error checking)

###### ReadConstrictor
```
func main() {
    response, _ := http.Get("google.com")
    defer response.Close()

    // read 1000 bytes per second
    res, _ := constrict.NewReader(response.Body, 1000)
}
```

###### WriteConstrictor
```
func main() {
    // read 3000 bytes per second
    _ := constrict.NewWriter(os.Stderr, os.DevNull, 3000)
}
```
