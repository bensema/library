package redis

import (
	"context"
	"github.com/gomodule/redigo/redis"
)

type TraceConn struct {
	ctx  context.Context
	conn redis.Conn
	hs   hooks
}

func (tc *TraceConn) Close() error {
	return tc.conn.Close()
}
func (tc *TraceConn) Err() error {
	return tc.conn.Err()
}

func (tc *TraceConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	var ctx context.Context
	cmd := NewCmd(tc.ctx, commandName, args...)

	ctx, err = tc.hs.beforeProcess(tc.ctx, cmd)
	if err != nil {
		cmd.SetErr(err)
		return
	}

	reply, err = tc.conn.Do(commandName, args...)

	if err = tc.hs.afterProcess(ctx, cmd); err != nil {
		cmd.SetErr(err)
	}
	return
}
func (tc *TraceConn) Send(commandName string, args ...interface{}) error {

	cmd := NewCmd(tc.ctx, commandName, args...)
	ctx, err := tc.hs.beforeProcess(tc.ctx, cmd)
	if err != nil {
		cmd.SetErr(err)
		return err
	}

	err = tc.conn.Send(commandName, args...)

	if err = tc.hs.afterProcess(ctx, cmd); err != nil {
		cmd.SetErr(err)
	}
	return err
}

func (tc *TraceConn) Flush() error {
	return tc.conn.Flush()
}

func (tc *TraceConn) Receive() (reply interface{}, err error) {
	return tc.conn.Receive()
}
