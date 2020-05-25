# wetterstats-cli

Easy way to get informations from wetterstats.de.

Available locations are:
  - Finowfurt
  - Berlin Pankow
  - Frankfurt (Oder)
  - Müllrose
  - Großziethen
  
### Installation

```sh
$ go build wetterstats-cli.go
```

### Mapping

| Location | Endpoint |
| ------ | ------ |
| Finowfurt | finowfurt |
| Berlin Pankow | pankow |
| Frankfurt (Oder) | frankfurt |
| Müllrose | muellrose |
| Großziethen | grossziethen |

### Usage

```sh
$ ./wetterstats-cli -endpoint=pankow
```