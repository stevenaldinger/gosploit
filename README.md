# Metasploit, but in golang...that's it!
![alt text](https://allg.one/G4zb)

# Build

```sh
docker build -t stevenaldinger/gosploit:latest .
```

# Run

```sh
docker run --rm -it stevenaldinger/gosploit:latest
```

# Dev

```sh
docker run --rm -it -v "$(pwd)":/go/src/github.com/stevenaldinger/gosploit stevenaldinger/gosploit:latest bash
```
