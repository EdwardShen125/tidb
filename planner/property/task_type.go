// Copyright 2018 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package property

// TaskType is the type of execution task.
type TaskType int

const (
	// RootTaskType stands for the tasks that executed in the TiDB layer.
	RootTaskType TaskType = iota

	// CopSingleReadTaskType stands for the TableScan or IndexScan tasks
	// executed in the coprocessor layer.
	CopSingleReadTaskType

	// CopMultiReadTaskType stands for the IndexLookup tasks executed in the
	// coprocessor layer.
	CopMultiReadTaskType

	// MppTaskType stands for task that would run on Mpp nodes, currently meaning the tiflash node.
	MppTaskType
)

// String implements fmt.Stringer interface.
func (t TaskType) String() string {
	switch t {
	case RootTaskType:
		return "rootTask"
	case CopSingleReadTaskType:
		return "copSingleReadTask"
	case CopMultiReadTaskType:
		return "copMultiReadTask"
	case MppTaskType:
		return "mppTask"
	}
	return "UnknownTaskType"
}
