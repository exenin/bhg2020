func main() {
 var (
 reader FooReader
 writer FooWriter
 )
 if _, err := io.Copy(&writer, &reader); err != nil {
 log.Fatalln("Unable to read/write data")
 }
}
