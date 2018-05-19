package depone

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type MyDepOne string

func (m MyDepOne) Hello() {
	fmt.Printf("Hello from MyDepOne, %s", m)
	logrus.Println("called: MyDepOne.Hello()")
}
