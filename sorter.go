/*
	Module: Sorter
	Author: Ryan
	Date: 2023/03/21
*/
package DeviceCore

import "net"

// StructSorter 分拣机结构体
type StructSorter struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Ip           string    `json:"ip"`
	Port         string    `json:"port"`
	Conn         *net.Conn `json:"conn"`
	CmdRun       string    `json:"cmd_run"`
	CmdStop      string    `json:"cmd_stop"`
	CmdTurnLeft  string    `json:"cmd_turn_left"`
	CmdTurnRight string    `json:"cmd_turn_right"`
	CmdLed1      string    `json:"cmd_led_1"`
	CmdLed2      string    `json:"cmd_led_2"`
	CmdLed3      string    `json:"cmd_led_3"`
	CmdLed4      string    `json:"cmd_led_4"`
	CmdLed5      string    `json:"cmd_led_5"`
	SignalIo     string    `json:"signal_io"`
	SignalIr     string    `json:"signal_ir"`
}

// 初始化设备
func (st *StructSorter) Init() (sorter StructSorter) {
	return sorter
}

// 运行设备
func (st *StructSorter) Run() (err error) {
	return err
}

// 停止设备
func (st *StructSorter) Stop() (err error) {
	return err
}

// 设备左转
func (st *StructSorter) TurnLeft() (err error) {
	return err
}

// 设备右转
func (st *StructSorter) TurnRight() (err error) {
	return err
}

// 运行设备左转
func (st *StructSorter) RunAndTurnLeft() (err error) {
	return err
}

// 运行设备右转
func (st *StructSorter) RunAndTurnRight() (err error) {
	return err
}

// 停止设备并回正
func (st *StructSorter) StopAndTurnCenter() (err error) {
	return err
}

//
