package rd

import "log"

type Data []byte

type ConsumeArgs struct {
	QueueName string
}

type ConsumeRet struct {
	ReturnValue []Data
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

	Lock.RLock()
	defer Lock.RLock()

	retCh := PdQueue.Subscribe(args.QueueName)
	workQ[args.QueueName] = retCh

	return nil
}

//Consume : Read current data from server, it is non-channel code. So only read what we have for now
func (t *WorkQueue) Consume(args *ConsumeArgs, reply *ConsumeRet) error {
	Lock.RLock()
	defer Lock.RUnlock()
	var empty []Data
	if vSlice, exist := workSlice[args.QueueName]; exist {
		reply.ReturnValue = vSlice
		log.Println("[rd][consume]  total data len ", len(vSlice))
		workSlice[args.QueueName] = empty
		return nil
	}

	*reply = ConsumeRet{}
	return nil
}

func (t *WorkQueue) Publish(args *PArgs, quo *PubRet) error {
	PdQueue.Publish(args.QValue, args.QName)

	//Do something here
	someRet := []byte("VALUE")
	PdQueue.Publish(someRet, args.QResponseName)

	return nil
}

func (t *WorkQueue) ListCount(args *int, reply *int) error {
	Lock.RLock()
	defer Lock.RUnlock()

	*reply = len(PdQueue.ListTopics())
	return nil
}
