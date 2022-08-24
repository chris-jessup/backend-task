package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
)

type IdNameRecord struct {
    Id      string
    Name    string
}


func readCsvToRecords(filename string) []IdNameRecord {

    // open file
    f, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }

    // close later
    defer f.Close()

    // read csv values using csv.Reader
    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    var records []IdNameRecord

    // Convert to IdNameRecord
    for i, line := range data {
        if i > 0 { // Skip Header
            var record IdNameRecord
            for j, field := range line {
                if j == 0 {
                    record.Id = field
                } else if j == 1 {
                    record.Name = field
                }
            }
            records = append(records, record)
        }
    }

    return records
}

func common_ids_with_bug(left []IdNameRecord, right []IdNameRecord) []string {

    var common_ids []string

    for _, element := range left {
        for _, rightElement := range right {
            if element == rightElement {
                common_ids = append(common_ids, element.Id)
            }
        }
    }

    return common_ids
}

func main() {

    if len(os.Args) < 3 {
        log.Fatal("Usage: " + os.Args[0] + " <left.csv> <right.csv>")
    }

    leftFileName := os.Args[1]
    rightFileName := os.Args[2]

    leftRecords := readCsvToRecords(leftFileName)
    rightRecords := readCsvToRecords(rightFileName)

    common := common_ids_with_bug(leftRecords, rightRecords)

    fmt.Println("IDs that are common to", leftFileName, "and", rightFileName)
    for _, id := range common {
        fmt.Println("    ", id)
    }

    // // print the array
    // fmt.Printf("%+v\n", shoppingList)
}
