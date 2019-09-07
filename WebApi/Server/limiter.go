//流控

package main

import "log"

type ConnLimit struct {
	CurrentConnection int
	Bucket chan int
}

func NewConnLimiter(n int)*ConnLimit{
	return &ConnLimit{
		CurrentConnection: n,
		Bucket:            make(chan int,n),
	}
}

func GetConn(cl *ConnLimit)bool{
	if(len(cl.Bucket)>=cl.CurrentConnection){
		return false
	}
	cl.Bucket<-1
	return true
}

func ReleaseConn(cl *ConnLimit){
	log.Println("release:",<-cl.Bucket)
}