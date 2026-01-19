### Simple gRPC example

## how to use
for windows just make sure you have make installed and golang then go to the terminal and run

```bash
    make
```
it will generate a bin directory that has `server.exe` and `client.exe`.

to start all you need to do is :
```bash
    .\bin\server.exe

    # then run in a seperate terminal

    .\bin\client.exe -msg="whatever you want"
```

the proto file is in `services` directory and the code is in the cmd directory
