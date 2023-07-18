package queue

import (
	"ssv/internal/core"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOneValidatorOneDuty_RightOrder(t *testing.T) {
	queue := NewQueue()
	queue.AddTask(0, core.DutyList[core.PROPOSER], Task{Height: 3, Message: "a"})
	queue.AddTask(0, core.DutyList[core.PROPOSER], Task{Height: 2, Message: "b"})
	queue.AddTask(0, core.DutyList[core.PROPOSER], Task{Height: 1, Message: "c"})

	task := queue.GetTask(0, core.DutyList[core.PROPOSER])
	require.Equal(t, 1, task.Height)
	require.Equal(t, "c", task.Message)

	task = queue.GetTask(0, core.DutyList[core.PROPOSER])
	require.Equal(t, 2, task.Height)
	require.Equal(t, "b", task.Message)

	task = queue.GetTask(0, core.DutyList[core.PROPOSER])
	require.Equal(t, 3, task.Height)
	require.Equal(t, "a", task.Message)
}

func TestOneValidatorThreeDuties(t *testing.T) {
	queue := NewQueue()
	queue.AddTask(0, core.DutyList[core.PROPOSER], Task{Height: 3, Message: "a"})
	queue.AddTask(0, core.DutyList[core.ATTESTER], Task{Height: 2, Message: "b"})
	queue.AddTask(0, core.DutyList[core.SYNC_COMMITTEE], Task{Height: 1, Message: "c"})

	task := queue.GetTask(0, core.DutyList[core.PROPOSER])
	require.Equal(t, 3, task.Height)
	require.Equal(t, "a", task.Message)

	task = queue.GetTask(0, core.DutyList[core.ATTESTER])
	require.Equal(t, 2, task.Height)
	require.Equal(t, "b", task.Message)

	task = queue.GetTask(0, core.DutyList[core.SYNC_COMMITTEE])
	require.Equal(t, 1, task.Height)
	require.Equal(t, "c", task.Message)
}

func TestNoTasks(t *testing.T) {
	queue := NewQueue()

	task := queue.GetTask(0, core.DutyList[core.PROPOSER])
	require.Nil(t, task)
}
