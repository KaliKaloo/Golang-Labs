package main

import (
	"container/list"
	"flag"
	"fmt"
	"math/rand"

	"strconv"
	"time"
)

var debug *bool

// An executor is a type of a worker goroutine that handles the incoming transactions.
func executor(bank *bank, executorID int, transactionQueue <-chan transaction, done chan<- bool) {
	for {
		t := <-transactionQueue

		from := bank.getAccountName(t.from)
		to := bank.getAccountName(t.to)

		fmt.Println("Executor\t", executorID, "attempting transaction from", from, "to", to)
		e := bank.addInProgress(t, executorID) // Removing this line will break visualisations.

		bank.execute(t, executorID)

		fmt.Println("Executor\t", executorId, "trying to unlock account", t.from)
		bank.unlockAccount(t.from, strconv.Itoa(executorID))
		fmt.Println("Executor\t", executorID, "unlocked account", from)
		
		fmt.Println("Executor\t", executorId, "trying to unlock account", t.to)
		bank.unlockAccount(t.to, strconv.Itoa(executorID))
		fmt.Println("Executor\t", executorID, "unlocked account", to)

		bank.removeCompleted(e, executorID) // Removing this line will break visualisations.
		
		done <- true
	}
}

func toChar(i int) rune {
	return rune('A' + i)
}

func manager(bank *bank, transactionQueue chan transaction, newQueue chan<- transaction) {
	for {
		select {
		case t := <-transactionQueue:
			if !(bank.accounts[t.from].locked && bank.accounts[t.to].locked) { //if both accounts not locked
				bank.lockAccount(t.from, "manager")
				fmt.Println("Manager\t", "locked account", t.from)
				bank.lockAccount(t.to, "manager")
				fmt.Println("Manager\t", "locked account", t.to)
				newQueue <- t
			} else {
				transactionQueue <- t
			}
		default:
			return
		}
	}
}

// main creates a bank and executors that will be handling the incoming transactions.
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	debug = flag.Bool("debug", false, "generate DOT graphs of the state of the bank")
	flag.Parse()

	bankSize := 6 // Must be even for correct visualisation.
	transactions := 1000

	accounts := make([]*account, bankSize)
	for i := range accounts {
		accounts[i] = &account{name: string(toChar(i)), balance: 1000}
	}

	bank := bank{
		accounts:               accounts,
		transactionsInProgress: list.New(),
		gen:                    newGenerator(),
	}

	startSum := bank.sum()

	transactionQueue := make(chan transaction, transactions)
	newQueue := make(chan transaction, transactions)
	expectedMoneyTransferred := 0
	for i := 0; i < transactions; i++ {
		t := bank.getTransaction()
		expectedMoneyTransferred += t.amount
		transactionQueue <- t
	}

	done := make(chan bool)

	go manager(&bank, transactionQueue, newQueue) //go routine for manager

	for i := 0; i < bankSize; i++ { //each account = 6 ececutor go routines
		go executor(&bank, i, newQueue, done)
	}

	for total := 0; total < transactions; total++ {
		fmt.Println("Completed transactions\t", total)
		<-done
	}

	fmt.Println()
	fmt.Println("Expected transferred", expectedMoneyTransferred)
	fmt.Println("Actual transferred", bank.moneyTransferred)
	fmt.Println("Expected sum", startSum)
	fmt.Println("Actual sum", bank.sum())
	if bank.sum() != startSum {
		panic("sum of the account balances does not much the starting sum")
	} else if len(transactionQueue) > 0 {
		panic("not all transactions have been executed")
	} else if bank.moneyTransferred != expectedMoneyTransferred {
		panic("incorrect amount of money was transferred")
	} else {
		fmt.Println("The bank works!")
	}
}
