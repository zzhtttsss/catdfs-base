package util

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"strings"
	"sync"
	"tinydfs-base/common"
	"tinydfs-base/protocol/pb"
)

type Queue[T fmt.Stringer] struct {
	data []T
	mu   *sync.RWMutex
}

func NewQueue[T fmt.Stringer]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
		mu:   &sync.RWMutex{},
	}
}

func (q *Queue[T]) Len() int {
	return len(q.data)
}

func (q *Queue[T]) Push(x T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.data = append(q.data, x)
}

func (q *Queue[T]) Pop() (t T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.data) == 0 {
		return
	}
	t = q.data[0]
	q.data = q.data[1:]
	return
}

func (q *Queue[T]) BatchPop(num int) (t []T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.data) < num {
		return
	}
	t = q.data[0:num]
	q.data = q.data[num:]
	return
}

func (q *Queue[T]) Top() (t T) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	if len(q.data) == 0 {
		return
	}
	t = q.data[0]
	return
}

func (q *Queue[T]) BatchTop(num int) (t []T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.data) < num {
		return
	}
	t = q.data[0:num]
	return
}

func (q *Queue[T]) String() string {
	// Make sure there is no write operations
	q.mu.RLock()
	defer q.mu.RUnlock()
	// Calling this function represents the leader has been changed, so the #{state} is unavailable
	sb := strings.Builder{}
	for q.Len() != 0 {
		x := q.Pop()
		sb.WriteString(fmt.Sprintf("%s%s", x, common.DollarDelimiter))
	}
	return sb.String()
}

// ChunkTaskResult represents the result of a chunk task.
type ChunkTaskResult struct {
	ChunkId          string   `json:"chunk_id"`
	FailDataNodes    []string `json:"fail_data_nodes"`
	SuccessDataNodes []string `json:"success_data_nodes"`
	SendType         int      `json:"send_type"`
}

// ConvReply2SingleResult converts the reply to ChunkTaskResult.
func ConvReply2SingleResult(reply *pb.TransferChunkReply, dataNodeIds []string,
	adds []string, sendType int) *ChunkTaskResult {
	singleSendResult := &ChunkTaskResult{
		ChunkId:  reply.ChunkId,
		SendType: sendType,
	}
	failDataNodes := make([]string, 0, len(adds))
	successDataNodes := make([]string, 0, len(adds))
	for i, address := range adds {
		isFail := false
		for _, a := range reply.FailAdds {
			if a == address {
				isFail = true
				break
			}
		}
		if isFail {
			failDataNodes = append(failDataNodes, dataNodeIds[i])
			continue
		}
		successDataNodes = append(successDataNodes, dataNodeIds[i])
	}
	singleSendResult.SuccessDataNodes = successDataNodes
	singleSendResult.FailDataNodes = failDataNodes
	return singleSendResult
}

// GetProgressBar returns a progressbar instance.
func GetProgressBar(size int64, description string, itsString string) *progressbar.ProgressBar {
	return progressbar.NewOptions64(size, progressbar.OptionSetDescription(description),
		progressbar.OptionEnableColorCodes(true), progressbar.OptionSetItsString(itsString),
		progressbar.OptionShowIts(), progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
}
