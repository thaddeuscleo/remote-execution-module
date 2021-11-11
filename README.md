# Remote Execution Module

Remote Execution Module is a lightweight psexec based executor with additional features. 
This module can be also used as a cli tool for managing rooms, you can see the usage example below.

## Compilation

For compiling you can use the go build command as follows:

```
go build main.go  
```

## Features
✅ Executing Shared Batch Script

✅ Wake On Lan

✅ Shutdown

✅ Restart

✅ Deep

✅ Undeep

❌ Log

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