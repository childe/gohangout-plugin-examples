# build

```
go build -buildmode=plugin -o file.so file_output.go
```

# config example

```
inputs:
    - Stdin: {}
filters:
    - Grok:
        match:
            - '^(?P<appid>\S+) (?P<level>\S+)'
outputs:
    - file.so:
        overfail_path: overfail.log
        path: '%{appid}/%{level}.log'
```
