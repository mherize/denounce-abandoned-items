package main

import (
	"denounce-abandoned-items/clients"
	"denounce-abandoned-items/utils"
	"embed"
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"

	"strconv"
	"sync"
)

//go:embed dump-items-nw.csv
//go:embed dump-items-ow.csv
var dump embed.FS

func main() {
	var wg sync.WaitGroup
	fileNW, err := dump.Open("dump-items-nw.csv")
	if err != nil {
		fmt.Println(err)
	}
	fileOW, err := dump.Open("dump-items-ow.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer fileNW.Close()
	defer fileOW.Close()

	wg.Add(1)
	go handleNewWorldFile(fileNW, &wg)
	//go handleOldWorldFile(fileOW, &wg)

	fmt.Printf("Proceso finalizado!!! :D")
}

func handleNewWorldFile(file fs.File, wg *sync.WaitGroup) {
	defer wg.Done()
	errorFile, _ := os.Create("errors-items-nw.csv")
	defer errorFile.Close()

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		itemID := record[0]
		userID, _ := strconv.Atoi(record[1])
		status, err := clients.DenounceItem(itemID)
		if err != nil {
			log.Printf("FAILED POST NW itemID: %s, status: %d error: %v\n", itemID, status, err)
			line := fmt.Sprintf("%s,%d\n", itemID, userID)
			errorFile.WriteString(line)
			continue
		}
		//newWorldUsers = append(newWorldUsers, userID)
	}
	errorFile.Sync()
	//filteredUsers := utils.RemoveDuplicateUsers(newWorldUsers)
	//sendMailStep(filteredUsers)
}

func handleOldWorldFile(file fs.File, wg *sync.WaitGroup) {
	defer wg.Done()
	var oldWorldUsers []int
	errorFile, _ := os.Create("errors-items-ow.csv")
	defer errorFile.Close()

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		itemID := record[0]
		userID, _ := strconv.Atoi(record[1])
		status, err := clients.PauseItemOW(itemID)
		if err != nil {
			log.Printf("FAILED POST OW itemID: %s, status: %d error: %v\n", itemID, status, err)
			line := fmt.Sprintf("%s,%d\n", itemID, userID)
			errorFile.WriteString(line)
			continue
		}
		oldWorldUsers = append(oldWorldUsers, userID)
	}
	errorFile.Sync()
	filteredUsers := utils.RemoveDuplicateUsers(oldWorldUsers)
	sendMailStep(filteredUsers)
}

func sendMailStep(users []int) {
	errorFile, _ := os.Create("errors-emails.csv")
	defer errorFile.Close()

	for _, usr := range users {
		email := utils.BuildEmail(usr)
		status, err := clients.SendMail(email)
		if err != nil || status != http.StatusCreated {
			fmt.Printf("FAILED SEND EMAIL userID: %d, status: %d, error: %v\n", usr, status, err)
			line := fmt.Sprintf("%d\n", usr)
			errorFile.WriteString(line)
		}
	}

	errorFile.Sync()
}
