package rd

import (
	"errors"
	"log"
)

type Args struct {
	QueueName string
}

type PubRet struct {
	Quo, Rem int
}

type WorkQueue struct {
}

type PArgs struct {
	QName         string
	QValue        []byte
	QResponseName string
}
type QArgs struct {
	QueueName string
}

func (t *WorkQueue) QueueDeclare(args *QArgs, reply *int) error {
	log.Println("Got QueueDeclare: args=", *args, " total=", workQ)

	retCh := PdQueue.Subscribe(args.QueueName)
	workQ[args.QueueName] = retCh

	return nil
}

func (t *WorkQueue) Consume(args *Args, reply *[]byte) error {

	if retCh, exist := workQ[args.QueueName]; exist {
		*reply = <-retCh
		return nil
	}
	return errors.New("New err")
}

func (t *WorkQueue) Publish(args *PArgs, quo *PubRet) error {

	PdQueue.Publish(args.QValue, args.QName)

	//Do something here
	someRet := []byte("VALUE")
	PdQueue.Publish(someRet, args.QResponseName)

	return nil
}

func (t *WorkQueue) ListQueue(args *int, reply *int) error {
	*reply = len(PdQueue.ListTopics())
	return nil
}
