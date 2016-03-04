package rd

import "log"

//Data Basic type of Queue data
type Data []byte

//ConsumeArgs Argment of Consume RPC call
type ConsumeArgs struct {
	QueueName string
}

//ConsumeRet Return Value of Consume RPC call
type ConsumeRet struct {
	ReturnValue []Data
}

//PublishArgs Argment of Publish RPC call
type PublishArgs struct {
	QName         string
	QValue        []byte
	QResponseName string
}

//QueryArgs Argment of Query
type QueryArgs struct {
	QueueName string
}

//WorkQueue RPC call instance
type WorkQueue struct {
}

//QueueDeclare Query the Queue if exist, if not will declare a new one
func (t *WorkQueue) QueueDeclare(args *QueryArgs, reply *int) error {
	log.Println("[rd] Got QueueDeclare: args=", *args, " total=", workQ)

	Lock.RLock()
	defer Lock.RLock()

	if _, exist := workQ[args.QueueName]; exist {
		log.Println("[rd] Topic already exist .... ")
		*reply = 1
		return nil
	}
	retCh := PdQueue.Subscribe(args.QueueName)
	workQ[args.QueueName] = retCh
	*reply = 0

	return nil
}

//Consume : Read current data from server, it is non-channel code. So only read what we have for now
func (t *WorkQueue) Consume(args *ConsumeArgs, reply *ConsumeRet) error {
	Lock.RLock()
	defer Lock.RUnlock()
	if vSlice, exist := workSlice[args.QueueName]; exist {
		reply.ReturnValue = vSlice
		log.Println("[rd][consume]  total data len ", len(vSlice))
		delete(workSlice, args.QueueName)
		return nil
	}

	*reply = ConsumeRet{}
	return nil
}

//Publish Publish data to Specific Queue
func (t *WorkQueue) Publish(args *PublishArgs, reply *int) error {
	PdQueue.Publish(args.QValue, args.QName)

	//Do something here
	return nil
}

//Count Display the count of Queues
func (t *WorkQueue) Count(args *int, reply *int) error {
	Lock.RLock()
	defer Lock.RUnlock()

	*reply = len(PdQueue.ListTopics())
	return nil
}
