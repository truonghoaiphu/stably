# stably

 How it works
follow steps: 1 -> 2

## 1. Prepare tools
Download and install golang from 

```
https://golang.org/doc/install
```


## 2. Run Service

```

For macos
➜ make build

For linux
➜ make linux

➜ make serve

```

## Testing
```
➜  make test
```

## Example REST URL
```
POST http://localhost:8080/find-highest-prime-number
{
    "number": 10000
}
```