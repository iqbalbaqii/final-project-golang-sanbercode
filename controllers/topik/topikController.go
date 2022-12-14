package topik

import (
	cons "final-project-golang-sanbercode/controllers"
	MTopik "final-project-golang-sanbercode/model/topik"
	"fmt"
)

func GetActivePolling() {
	var today = cons.GetToday()
	fmt.Println(MTopik.GetActivePolling())
	_ = today
}
