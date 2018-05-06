package p2p

import (
	"net"
	"github.com/cyyber/go-QRL/log"
	"sync"
	"io"
	"github.com/cyyber/go-QRL/generated"
	"github.com/golang/protobuf/proto"
	"time"
)

type Peer struct {
	conn		net.Conn
	inbound		bool

	wg			sync.WaitGroup
	closed		chan struct{}
	log			log.Logger
}

func newPeer(conn *net.Conn, inbound bool, log *log.Logger) *Peer {
	p := &Peer {
		conn: *conn,
		inbound: inbound,
		log: *log,
	}
	return p
}

func (p *Peer) WriteMsg(msg Msg) error {
	return nil
}

func (p *Peer) ReadMsg() (msg Msg, err error){
	buf := make([]byte, 4)
	if _, err := io.ReadFull(p.conn, buf); err != nil {
		return msg, err
	}
	size := convertBytesToLong(buf)
	buf = make([]byte, size)
	if _, err := io.ReadFull(p.conn, buf); err != nil {
		return msg, err
	}
	message := &generated.LegacyMessage{}
	err = proto.Unmarshal(buf, message)

	return msg, err
}

func (p *Peer) readLoop(errc chan<- error) {
	defer p.wg.Done()
	for {
		msg, err := p.ReadMsg()
		if err != nil {
			errc <- err
			return
		}
		msg.ReceivedAt = time.Now()
		if err = p.handle(msg); err != nil {
			errc <- err
			return
		}
	}
}

func (p *Peer) pingLoop() {
	defer p.wg.Done()

}

func (p* Peer) handle(msg Msg) error {
	switch msg.msg.FuncName {
	case generated.LegacyMessage_VE:
		p.log.Debug("Received VE MSG")
	case generated.LegacyMessage_PL:
		p.log.Debug("Received PL MSG")
	case generated.LegacyMessage_PONG:
		p.log.Debug("Received PONG MSG")
	case generated.LegacyMessage_MR:
	case generated.LegacyMessage_SFM:
	case generated.LegacyMessage_BK:
	case generated.LegacyMessage_FB:
	case generated.LegacyMessage_PB:
	case generated.LegacyMessage_BH:
	case generated.LegacyMessage_TX:
	case generated.LegacyMessage_LT:
	case generated.LegacyMessage_EPH:
	case generated.LegacyMessage_MT:
	case generated.LegacyMessage_TK:
	case generated.LegacyMessage_TT:
	case generated.LegacyMessage_SL:
	case generated.LegacyMessage_SYNC:
	case generated.LegacyMessage_CHAINSTATE:
	case generated.LegacyMessage_HEADERHASHES:
	case generated.LegacyMessage_P2P_ACK:
	}
	return nil
}

func (p *Peer) run() (remoteRequested bool, err error) {
	var (
		writeStart = make(chan struct{}, 1)
		writeErr = make(chan error, 1)
		readErr	 = make(chan error, 1)
	)
	p.wg.Add(2)
	go p.readLoop(readErr)
	go p.pingLoop()

loop:
	for {
		select {
		case err = <-writeErr:
			if err != nil {
				break loop
			}
			writeStart <- struct{}{}
		case err = <-readErr:
			if err != nil {
				break loop
			}
		}
	}
}

func convertBytesToLong(b []byte) uint32 {
	return uint32(b[0]) << 24 | uint32(b[1]) << 16 | uint32(b[2]) << 8 | uint32(b[3])
}