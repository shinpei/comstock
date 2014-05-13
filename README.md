## Comstock: store your favorite command

Comstock is a cloud-based command stocking tool. On this repository, we're providing cli application. comstock-cli can be used as normal tool without network connection. It'll sync when it gets internet connection.

<!--
## Motivation
We have a plenty of convenient command line tools nowadays, such as `git`, `brew`, `chef`,  Thanks to github, providing commands become a fame for developers, making good tools is now a orner.
-->

## Use as local tool
```
$ brew doctor
$ comstock save
saved command 'brew doctor'
$ comstock list
1: brew doctor
$ comstock run 1
```

<!--
## Stock command to the cloud
```
$ comstock login
$ comstock push
```
-->

## Install

```
go get github.com/shinpei/comstock
```



## Lisence, contact info, contribute
It's under [ASL2.0](http://www.apache.org/licenses/LICENSE-2.0). If you find bug or improvement request, please contact me through twitter, @shinpeintk. And always welcoming heartful pull request.

cheers, :coffee: :moyai:




