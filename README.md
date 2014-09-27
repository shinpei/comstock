## Comstock: store your favorite command

[![Build Status](https://drone.io/github.com/shinpei/comstock/status.png)](https://drone.io/github.com/shinpei/comstock/latest)
[![GoDoc](https://godoc.org/github.com/shinpei/comstock?status.png)](https://godoc.org/github.com/shinpei/comstock)

Comstock is a cloud-based command stocking tool. Copy and pasting your often used command to the text editor, or cloud-based editor like evernote was a dull for me. History command also won't allow me to use favorite commands in the new environment. What Comstock provides is storing your commands to the cloud, and easily use them from anywhere. 

## Demo
![](https://github.com/shinpei/comstock/blob/master/comstock-demo.gif)

## Install

####via `curl`
`
curl -L https://github.com/shinpei/comstock/raw/master/dist/install.sh | sudo env HOME=$HOME sh
`

####via `wget`
```
wget --no-check-certificate https://github.com/shinpei/comstock/raw/master/dist/install.sh -O - | sudo env HOME=$HOME sh
```

#### via `homebrew`
```
$ brew tap shinpei/comstock
$ brew install comstock
```

If you have installed Go, type following.
```
go get github.com/shinpei/comstock
```
## Create your account
Open register website from command.
```
$ comstock open
```

## Lisence, contact info, contribute
It's under [ASL2.0](http://www.apache.org/licenses/LICENSE-2.0). If you find bug or improvement request, please contact me through twitter, @shinpeintk. And always welcoming heartful pull request.

cheers, :coffee: :moyai:




