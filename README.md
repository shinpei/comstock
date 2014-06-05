## Comstock: store your favorite command

[![Build Status](https://drone.io/github.com/shinpei/comstock/status.png)](https://drone.io/github.com/shinpei/comstock/latest)
[![GoDoc](https://godoc.org/github.com/shinpei/comstock?status.png)](https://godoc.org/github.com/shinpei/comstock)

Comstock is a cloud-based command stocking tool. On this repository, we're providing cli application. comstock-cli can be used as normal tool without network connection. It'll sync when it gets internet connection.

## Usage
```
$ git diff HEAD^ --name-only
$ comstock save
saved command 'git diff HEAD^ --name-only'
$ comstock list
1: git diff HEAD^ --name-only
$ comstock get 1
git diff HEAD^ --name-only
```

## Sync with Cloud

```
$ comstock login
Your registered email? : shinpei@mail.com
Password for shinpei@mail.com?:
Authentification success.
$ git diff HEAD^ --name-only
$ comstock save
saved command 'git diff HEAD^ --name-only'
```

### Getting started
## Install
If you have installed Go, type following.
```
go get github.com/shinpei/comstock
```
Or, you can install comstock from homebrew.
```
$ brew tap shinpei/comstock
$ brew install comstock
```
## Create your account
Open register website from command.
```
$ comstock open
```

## Lisence, contact info, contribute
It's under [ASL2.0](http://www.apache.org/licenses/LICENSE-2.0). If you find bug or improvement request, please contact me through twitter, @shinpeintk. And always welcoming heartful pull request.

cheers, :coffee: :moyai:




