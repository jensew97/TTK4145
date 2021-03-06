package DriveElevator

import (
	"./EventManager"
	"fmt"
	"time"
)

var InternalQueue [3][4]int

type Button struct {
	Dir   int
	Floor int
}

func InternalQueueAddNewOrder(backup *Network.Backup) {
	InternalQueue = backup.Orders
}

func InternalQueueDeleteOrder(currentFloor int) {
	for i := 0; i < 3; i++ {
		InternalQueue[i][currentFloor] = 0
	}
}

func InternalQueueDeleteQueue() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			InternalQueue[i][j] = 0
		}
	}
}

func InternalQueueCheckOrdersAbove(currentFloor int) int {
	for i := currentFloor + 1; i < 4; i++ {
		if InternalQueue[0][i] == 1 || InternalQueue[1][i] == 1 || InternalQueue[2][i] == 1 {
			return 1
		}
	}
	return 0
}

func InternalQueueCheckOrdersBlow(currentFloor int) int {
	for i := currentFloor - 1; i > -1; i-- {
		if InternalQueue[0][i] == 1 || InternalQueue[1][i] == 1 || InternalQueue[2][i] == 1 {
			return 1
		}
	}
	return 0
}

func Internal_queue_should_stop(dir int, currentFloor int) int {
	if InternalQueue[2][currentFloor] == 1 {
		return 1
	}
	if dir == UP {
		if InternalQueue[0][currentFloor] == 1 || InternalQueueCheckOrdersAbove(currentFloor) == 0 {
			return 1
		}
	}
	if dir == DOWN {
		if InternalQueue[1][currentFloor] == 1 || InternalQueueCheckOrdersBelow(currentFloor) == 0 {
			return 1
		}
	}
	if dir == STOP {
		if InternalQueue[0][currentFloor] == 1 || InternalQueue[1][currentFloor] == 1 || InternalQueue[2][currentFloor] == 1 {
			return 1
		}
	}
	return 0
}

func InternalQueueChooseDir() int {
	switch MotorDir {
	case UP:
		if InternalQueueCheckOrdersAbove(ElevatorFloor) == 1 {
			return UP
		} else if InternalQueueCheckOrdersBelow(ElevatorFloor) == 1 {
			return DOWN
		} else {
			return STOP
		}
	case DOWN:
		if InternalQueueCheckOrdersBelow(ElevatorFloor) == 1 {
			return DOWN
		} else if InternalQueueCheckOrdersAbove(ElevatorFloor) == 1 {
			return UP
		} else {
			return STOP
		}

	case STOP:
		if InternalQueueCheckOrdersAbove(ElevatorFloor) == 1 {
			return UP
		} else if InternalQueueCheckOrdersBelow(ElevatorFloor) == 1 {
			return DOWN
		} else {
			return STOP
		}
	default:
		return STOP
	}

}

func InternalQueuePollButtons(newHWOrderCh chan Button) {

	var buttonStatus [3][4]bool
	for {
		for i := 0; i < 3; i++ {
			for j := 0; j < 4; j++ {
				if !(i == 1 && j == 0) && !(i == 0 && j == 3) {
					if EventManager.ElevatorGetButtonSignal(i, j) == 1 && !buttonStatus[i][j] {
						button := Button{i, j}
						buttonStatus[i][j] = true
						newHWOrderCh <- button
					} else if EventManager.ElevatorGetButtonSignal(i, j) == 0 {
						buttonStatus[i][j] = false
					}
				}
			}
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func InternalQueuePollFloorSensors(floorSensorCh chan int) {
	var prevFloor int
	for {
		floor := EventManager.ElevatorGetFloorSensorSignal()
		if floor != -1 && prevFloor != floor {
			floorSensorCh <- floor
			fmt.Println("floor sensor...", floor)
			prevFloor = floor
		}
		time.Sleep(50 * time.Millisecond)
	}
}
