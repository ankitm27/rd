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
	log.Println("[rd] Got QueueDeclare: args=", *args, " total=", queueMapChan)

	rlock.RLock()
	defer rlock.RLock()

	if _, exist := queueMapChan[args.QueueName]; exist {
		log.Println("[rd] Topic already exist .... ")
		*reply = 1
		return nil
	}
	retCh := pubsubObj.Subscribe(args.QueueName)
	queueMapChan[args.QueueName] = retCh
	*reply = 0

	return nil
}

//Consume : Read current data from server, it is non-channel code. So only read what we have for now
func (t *WorkQueue) Consume(args *ConsumeArgs, reply *ConsumeRet) error {
	rlock.RLock()
	defer rlock.RUnlock()
	if vSlice, exist := queueMapData[args.QueueName]; exist {
		reply.ReturnValue = vSlice
		log.Println("[rd][consume]  total data len ", len(vSlice))
		delete(queueMapData, args.QueueName)
		return nil
	}

	*reply = ConsumeRet{}
	return nil
}

//Publish Publish data to Specific Queue
func (t *WorkQueue) Publish(args *PublishArgs, reply *int) error {
	pubsubObj.Publish(args.QValue, args.QName)

	//Do something here
	return nil
}

//Count Display the count of Queues
func (t *WorkQueue) Count(args *int, reply *int) error {
	rlock.RLock()
	defer rlock.RUnlock()

	*reply = len(pubsubObj.ListTopics())
	return nil
}
