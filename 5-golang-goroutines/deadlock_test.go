package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	Mutex   sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)
	// user1.Unlock()

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user1.Name)
	user2.Change(amount)
	// user2.Unlock()

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "John",
		Balance: 100000,
	}
	user2 := UserBalance{
		Name:    "Jane",
		Balance: 200000,
	}

	fmt.Println("User1 ", user1.Name, ", Balance", user1.Balance)
	fmt.Println("User2 ", user2.Name, ", Balance", user2.Balance)

	go Transfer(&user1, &user2, 10000)
	go Transfer(&user2, &user1, 35000)

	time.Sleep(3 * time.Second)

	fmt.Println("After Transfer User1 ", user1.Name, ", Balance", user1.Balance)
	fmt.Println("After Transfer User2 ", user2.Name, ", Balance", user2.Balance)
}
