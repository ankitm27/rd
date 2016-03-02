package rd

import (
	"errors"

	pd "github.com/kkdai/pd"
)

type Args struct {
	A, B int
}

type PubRet struct {
	Quo, Rem int
}

type WorkQueue struct {
	queueMap map[string]pd.PD
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
	return nil
}

func (t *WorkQueue) Consume(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *WorkQueue) Publish(args *PArgs, quo *PubRet) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
