package main

type Storage interface {
	Open ();
	GetReader();
	GetWriter();
}

type LocalStorage interface{ }

type RemoteStorage interface {} 

type FileStorage struct {

}
