package main

import "fmt"

// 医生
type Docter struct{}

// 治疗眼睛
func (d *Docter) treatEye() {
	fmt.Println("医生治疗眼睛")
}

// 治疗鼻子
func (d *Docter) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// 治疗眼睛的病单
type CommandTreatEye struct {
	docter *Docter
}

func (cmd *CommandTreatEye) Treat() {
	cmd.docter.treatEye()
}

type CommandTreatNose struct {
	docter *Docter
}

func (cmd *CommandTreatNose) Treat() {
	cmd.docter.treatNose()
}

func main() {
	docker := new(Docter)

	cmdEye := CommandTreatEye{docker}
	cmdEye.Treat()

	cmdNose := CommandTreatNose{docker}
	cmdNose.Treat()
}
