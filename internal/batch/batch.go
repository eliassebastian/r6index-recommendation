package batch

import (
	"log"
	"sync"
	"time"
)

type BatchPipelineCallback func([]interface{}) error

type BatchPipeline struct {
	// ...
	maxSize    int
	maxWait    time.Duration
	mutex      *sync.RWMutex
	data       []interface{}
	executeFnc BatchPipelineCallback
	flushChan  chan struct{}
}

func NewBatchPipeline(maxSize int, maxWait time.Duration, executeFnc BatchPipelineCallback) *BatchPipeline {
	batch := &BatchPipeline{
		maxSize:    maxSize,
		maxWait:    maxWait,
		mutex:      &sync.RWMutex{},
		data:       make([]interface{}, 0, maxSize),
		executeFnc: executeFnc,
		flushChan:  make(chan struct{}),
	}

	// start a goroutine to flush the data pipeline every maxWait
	go batch.flushAfterDeadline()

	return batch
}

func (bp *BatchPipeline) Add(data interface{}) {
	bp.mutex.Lock()
	defer bp.mutex.Unlock()

	bp.data = append(bp.data, data)

	// if the data pipeline is full or over limit (error running callback), execute the callback function
	if len(bp.data) >= bp.maxSize {
		err := bp.executeAndFlush()
		if err != nil {
			log.Println(err)
		}
	}
}

func (bp *BatchPipeline) executeAndFlush() error {
	// ...
	err := bp.executeFnc(bp.data)

	// reset the data pipeline if execution was successful
	if err == nil {
		// flush the data pipeline and keep allocated memory
		bp.data = bp.data[:0]
	}

	return err
}

func (bp *BatchPipeline) flushAfterDeadline() {
	for {
		select {
		case <-bp.flushChan:
			bp.flushChan <- struct{}{}
			return
		case <-time.After(bp.maxWait):
			bp.mutex.Lock()
			err := bp.executeAndFlush()
			if err != nil {
				log.Println(err)
			}
			bp.mutex.Unlock()
		}
	}
}
