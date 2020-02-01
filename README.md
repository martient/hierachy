# ToDefine

OUEP

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/ToDefine
```

### Create new csv file

``` bash
hierachy create
```

Generate 'hierachy.csv'
but if you want an custom file name

``` bash
hierachy create -n theBestTeam
```

### Append someone to your hierachy file

``` bash
hierachy append -n theBestTeam.csv "FirstName LastName" POSITION "OFFICE PLACE" email@demo.com
```

for example:

``` bash
hierachy append -f theBestTeam.csv "Stephane Durondpoints" CEO "Office 1, floor 2" stephane.durondpoints@bestcompany.com
```

#### If your guy have manager

``` bash
hierachy append -f theBestTeam.csv "FirstName LastName" POSITION "OFFICE PLACE" email@demo.com -m "MANAGER FirstName LastName"
```

for example:

``` bash
hierachy append -f theBestTeam.csv "Kevin Sapindhotel" "software engineer intern" "Open Space 2, floor 1" kevin.sapindhotel@bestcompany.com -m "Stephane Durondpoints"
```

### Testing

``make test``