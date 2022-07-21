package rlimiter

import (
	"sync/atomic"
)

type Manager struct {
	inJob        chan interface{}
	outJob       chan interface{}
	finishJob    chan struct{}
	done         chan struct{}
	sender       Sender
	worker       uint8
	maxLimit     int64
	currentLimit int64
}

func NewManager(confs ...*Config) (*Manager, error) {
	conf, err := MergeConfigs(confs...)
	if err != nil {
		return nil, err
	}

	man := &Manager{
		inJob:        make(chan interface{}),
		outJob:       make(chan interface{}),
		finishJob:    make(chan struct{}),
		done:         make(chan struct{}),
		currentLimit: 0,
		sender:       conf.Sender,
		worker:       conf.Worker,
		maxLimit:     conf.MaxLimit,
	}

	return man, nil
}

func (m *Manager) Start() {
	go m.doWork()
	go m.receive()
}

// [TODO] Wait remains job to finish
func (m *Manager) Finish() {
	m.done <- struct{}{}
}

// Put msg to queue
func (m *Manager) Send(msg interface{}) {
	m.inJob <- msg
}

func (m *Manager) receive() {
	for {
		select {
		case <-m.done:
			return
		case msg := <-m.inJob:
			m.outJob <- msg

			m.tryToPause()
		}
	}
}

func (m *Manager) doSend(workerIdx int) {
	for {
		select {
		case <-m.done:
			return
		case msg := <-m.outJob:
			m.sender.Send(msg)

			go func() {
				m.finishJob <- struct{}{}
			}()
		}
	}
}

func (m *Manager) doWork() {
	for wrkIdx := 0; wrkIdx < int(m.worker); wrkIdx++ {
		go m.doSend(wrkIdx)
	}
}

func (m *Manager) waitBySize(size int) {
	for i := 0; i < size; i++ {
		<-m.finishJob
	}
}

func (m *Manager) isExceedLimit() bool {
	return m.currentLimit >= m.maxLimit-1
}

func (m *Manager) tryToPause() {
	if m.isExceedLimit() {
		m.waitBySize(int(m.maxLimit))
		m.resetLimit()
	}

	m.incLimit()
}

func (m *Manager) incLimit() {
	atomic.AddInt64(&m.currentLimit, 1)
}

func (m *Manager) resetLimit() {
	m.currentLimit = -1
}
