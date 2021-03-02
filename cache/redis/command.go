package redis

import (
	"context"
	"time"
)

type Cmder interface {
	Name() string
	//FullName() string
	Args() []interface{}
	//String() string
	//stringArg(int) string

	//readTimeout() *time.Duration
	//readReply(rd *proto.Reader) error

	SetErr(error)
	Err() error
}

type baseCmd struct {
	ctx     context.Context
	cmdName string
	args    []interface{}
	err     error

	_readTimeout *time.Duration
}

func (cmd *baseCmd) Name() string {
	return cmd.cmdName
}

func (cmd *baseCmd) Args() []interface{} {
	return cmd.args
}

func (cmd *baseCmd) SetErr(e error) {
	cmd.err = e
}

func (cmd *baseCmd) Err() error {
	return cmd.err
}

type Cmd struct {
	baseCmd
	val interface{}
}

func NewCmd(ctx context.Context, cmdName string, args ...interface{}) *Cmd {
	return &Cmd{
		baseCmd: baseCmd{
			cmdName: cmdName,
			ctx:     ctx,
			args:    args,
		},
	}
}
