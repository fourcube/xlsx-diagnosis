A simple demonstration of a problem with reopening xlsx files that have been saved with the xlsx library.

`go run main.go`

```
$ go run main.go
Opening original.xlsx
Saving after_write.xlsx
Opening after_write.xlsx
panic: runtime error: index out of range

goroutine 1 [running]:
main.main()
        f:/Development/Projects/GO/src/bitbucket.org/fourcube/xlsx-diagnosis/main.go:23 +0x29d
exit status 2
```
