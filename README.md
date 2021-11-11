# Roomnetman CLI

Roomnetman CLI is a tool for dealing with room management written in GO.

## Compilation

For compiling you can use the go build command as follows:

```
go build main.go  
```

## Features
✅ Executing Shared Batch Script

✅ Wake On Lan

❌ Log

❌ Shutdown

❌ Restart

❌ Deep

❌ Undeep



## Usage

The usage examples:

```bash
# executing batch script with the (run) command
main.exe run -p [password] -r 626 -x "shared\main.bat"

# wake up computers with (wake) command
main.exe wake -r 626

# shutdown computers with shutdown computer
main.exe shutdown -r 626
```