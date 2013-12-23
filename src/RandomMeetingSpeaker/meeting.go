package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var bottle string = ` [=]
 | |
 }_{
/   \
:___:
|###|
|###|
|---|
'---'
`

func main() {

	users := make(map[int]string)

	users[0] = "Mike"
	users[1] = "Paul"
	users[2] = "Steve"
	users[3] = "Lawrence"
	users[4] = "Stephen"
	users[5] = "James"

	getNextSpeaker(users)

}

func getNextSpeaker(users map[int]string) {
	numUsers := len(users)
	list := randList(1, numUsers)
	clearCmd()

	for _, element := range list {
		spinTheBottle()
		fmt.Println(users[element-1] + " the bottle has landed on you!")
		fmt.Print(bottle)
		pressAnyKey()
		clearCmd()
	}

}

func randList(min, max int) []int {
	if max < min {
		min, max = max, min
	}
	length := max - min + 1
	t0 := time.Now()
	rand.Seed(int64(t0.Nanosecond()))
	list := rand.Perm(length)
	for index, _ := range list {
		list[index] += min
	}

	return list
}

func pressAnyKey() {
	fmt.Println("Press return to spin the bottle...")
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
}

func clearCmd() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func spinTheBottle() {
	anim := make(map[int]string)
	anim[0] = string('|')
	anim[1] = string('\\')
	anim[2] = string('-')
	anim[3] = string('/')
	anim[4] = anim[1]
	anim[5] = anim[2]
	anim[6] = anim[3]
	anim[7] = anim[0]
	anim[8] = anim[1]
	anim[9] = anim[2]
	anim[10] = anim[3]
	clearCmd()
	
	for i := 0; i < len(anim); i++ {
		fmt.Println("Round and round and round it goes where it stops nobody knows...")
		fmt.Println(anim[i])
		time.Sleep(500 * time.Millisecond)
		clearCmd()
	}
	
}
