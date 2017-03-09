package Master

import (
	"../DriveElevator"
)
/*

type Slave struct {
	IP         string
	Alive      bool
	Descendant int
	//Last_floor int
	//Direction  int
}

type Master struct {
	Slaves       map[Slave]time.Time
	IP           string
	Participants int
	Last_floor   int
	Direction    int
}

func (master *Master) master_init() {
	fmt.Println("Master init...")
	master.IP,_ = Network.Udp_get_local_ip()
	master.Participants = 0
	for {
		Network.Udp_broadcast(master.IP)
		time.Sleep(50 * time.Millisecond)
	}
}

func Master_detect_slave(chan_rec_msg chan []byte, chan_is_alive chan string, port_nr string, chan_error chan error, master_p *Master) {
	go Network.Udp_receive_is_alive(chan_rec_msg, chan_is_alive, port_nr, chan_error)
	master := *master_p
	master.master_init()

	for {
		msg := <-chan_is_alive
		is_updated := false

		if msg != "" {
			for s := range master.Slaves {
				if s.IP == msg {
					master.Slaves[s] = time.Now()
				}
			}
			if is_updated == false {
				new_slave := Slave{msg, true, master.Participants + 1}
				master.Slaves[new_slave] = time.Now()
				master.Participants += 1
			}
		}

		const DEADLINE = 1 * time.Second
		descendant := -1
		for slave, last_time := range master.Slaves {
			if time.Since(last_time) > DEADLINE {
				descendant = slave.Descendant
				delete(master.Slaves, slave)
				master.Participants -= 1
			}
		}
		for s := range master.Slaves {
			if s.Descendant > descendant {
				s.Descendant -= 1
				Network.Udp_send_descendant_nr(s.Descendant, s.IP)
			}
		}
	}

}*/

func Master_write_backup(backup_p *Network.Backup, order DriveElevator.Button, source_ip string) [][]int {
	backup = *backup_p
	var set_lights [][]int

	for order := range backup.MainQueue{
		if backup.MainQueue[order] == source_ip {
			order.Orders[button.dir][button.floor] = 1
		}
		for i := 0; i < 3; i++ {
			for j := 0; i < 4; i++ {
				if order.Orders[i][j] == 1 {
					set_lights[i][j] = 1
				}
			}
		}
	}
	return set_lights
}

func Master_test_drive(chan_new_hw_order chan DriveElevator.Button, chan_new_master_order chan DriveElevator.Button) {
	for {
		select {
		case new_hw_order := <- chan_new_hw_order:
			chan_new_master_order <- new_hw_order
			Master_write_backup()
		}
	}
}
/*
func Master_drive_elevator(backup_p *Network.Backup, chan_new_order chan Network.NewOrder, chan_source_ip chan string, chan_received_msg chan []byte, port_nr string, chan_error chan error, master_p *Master, chan_order_executed chan int, chan_new_hw_order chan DriveElevator.Button, chan_new_master_order chan DriveElevator.Button) {
	go Network.Udp_receive_new_order(chan_new_order, chan_received_msg, port_nr, chan_error, chan_source_ip)
	//go DriveElevator.Driveelevator_get_new_order(chan_elev_order)
	go Network.Udp_receive_order_status(chan_new_order, chan_received_msg, port_nr, chan_error, chan_source_ip)
	backup := *backup_p
	//master := *master_p
	local_ip,_ := Network.Udp_get_local_ip() 

	for {
		select {
		case new_slave_order := <-chan_new_order:
			source := <-chan_source_ip

			if new_slave_order.Direction == 2 { //inside order
				for order := range backup.MainQueue {
					if backup.MainQueue[order] == source {
						order.Orders[new_slave_order.Direction][new_slave_order.Floor] = 1
					}
				}
			} else {
				elevator := local_ip
				//elevator := master.Master_queue_delegate_order(new_slave_order)
				if elevator != local_ip {
					Network.Udp_send_new_order(new_slave_order, elevator)
				} else {
					order := DriveElevator.Button{new_slave_order.Floor, new_slave_order.Direction}
					chan_new_master_order <- order
				}
				for order := range backup.MainQueue {
					if backup.MainQueue[order] == elevator {
						order.Orders[new_slave_order.Direction][new_slave_order.Floor] = 1
					}
				}
			}
		case new_hw_order := <-chan_new_hw_order:
			elevator,_ := Network.Udp_get_local_ip()

			if new_hw_order.Dir == 2 {
				chan_new_master_order <- new_hw_order
			} else {
				order := Network.NewOrder{new_hw_order.Floor, new_hw_order.Dir, 1, 0, 0}
				//elevator = master.Master_queue_delegate_order(order)
				elevator := local_ip
				if elevator != local_ip {
					Network.Udp_send_new_order(order, elevator)
				} else {
					chan_new_master_order <- new_hw_order
				}
			}
			//add order to main queue
			for order := range backup.MainQueue {
				if backup.MainQueue[order] == elevator {
					order.Orders[new_hw_order.Dir][new_hw_order.Floor] = 1
				}
			}
		case executed := <- chan_order_executed:
			source_ip := <-chan_source_ip
				//slette fra hovedkø med floor og alle buttons.
			for order := range backup.MainQueue {
				if backup.MainQueue[order] == source_ip {
					for i := 0; i < 3; i++ {
						order.Orders[i][executed] = 0
					}
				}
			}
		}
	}

}

//Må oppdatere direction og floor variablene hos master og slaver.
*/