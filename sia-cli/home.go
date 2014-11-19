package main

import (
	"fmt"

	"github.com/NebulousLabs/Andromeda/sia"
)

func toggleMining(env *walletEnvironment) {
	go env.state.ToggleMining(env.wallets[0].SpendConditions.CoinAddress())
}

// printWalletAddresses prints out all of the addresses that are spendable by
// this cli.
func printWalletAddresses(env *walletEnvironment) {
	fmt.Println("General Information:")

	// Dispaly whether or not the miner is mining.
	if env.state.Mining {
		fmt.Println("\tMining Status: ON - Wallet is mining on one thread.")
	} else {
		fmt.Println("\tMining Status: OFF - Wallet is not mining.")
	}
	fmt.Println()

	fmt.Println("\tCurrent Block Height:", env.state.Height())
	fmt.Println("\tCurrent Block Depth:", env.state.Depth())
	fmt.Println()

	fmt.Println("\tPrinting all valid wallet addresses.")
	for _, wallet := range env.wallets {
		fmt.Printf("\t\t%x\n", wallet.SpendConditions.CoinAddress())
	}
}

// sendCoinsWalkthorugh uses the wallets in the environment to send coins to an
// address that is provided through the command line.
func sendCoinsWalkthrough(env *walletEnvironment) (err error) {
	fmt.Println("Send Coins Walkthrough:")

	fmt.Print("Amount to send: ")
	var amount int
	_, err = fmt.Scanln(&amount)
	if err != nil {
		return
	}

	fmt.Print("Amount to use as Miner Fee: ")
	var minerFee int
	_, err = fmt.Scanln(&minerFee)
	if err != nil {
		return
	}

	fmt.Print("Address of Receiving Wallet: ")
	var addressBytes []byte
	_, err = fmt.Scanf("%x", &addressBytes)
	if err != nil {
		return
	}

	// Convert the address to a sia.CoinAddress
	var address sia.CoinAddress
	copy(address[:], addressBytes)

	// Use the wallet api to send. ==> Only uses wallets[0] for the time being.
	fmt.Printf("Sending %v coins with miner fee of %v to address %x", amount, minerFee, address[:])
	env.wallets[0].SpendCoins(sia.Currency(amount), sia.Currency(minerFee), address, env.state)

	return
}

// displayHomeHelp lists all of the options available at the home screen, with
// descriptions.
func displayHomeHelp() {
	fmt.Println(
		" h:\tHelp - display this message\n",
		"q:\tQuit - quit the program\n",
		"m:\tMine - turn mining on or off\n",
		"p\tPrint - list all of the wallets, plus some stats about the program\n",
		"s:\tSend - send coins to another wallet\n",
	)
}

// pollHome repeatedly querys the user for input.
func pollHome(env *walletEnvironment) {
	var input string
	var err error
	for {
		fmt.Println()
		fmt.Print("(Home) Please enter a command: ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			continue
		}
		fmt.Println()

		switch input {
		default:
			fmt.Println("unrecognized command")

		case "?", "h", "help":
			displayHomeHelp()

		case "q", "quit":
			return

		/*
			case "H", "host", "store", "advertise", "storage":
				becomeHostWalkthrough(env)
		*/

		case "m", "mine", "toggle", "mining":
			toggleMining(env)

		case "p", "print":
			printWalletAddresses(env)

		/*
			case "r", "rent":
				becomeClientWalkthrough(env)
		*/

		/*
			case "S", "save":
				saveWalletEnvironmentWalkthorugh()
		*/

		case "s", "send":
			err = sendCoinsWalkthrough(env)
		}

		if err != nil {
			fmt.Println("Error:", err)
			err = nil
		}
	}
}