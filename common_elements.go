package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
)

type IdNameRecord struct {
    id      string
    name    string
}


func readCsvToIdNameRecords(filename string) []IdNameRecord {

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
                    record.id = field
                } else if j == 1 {
                    record.name = field
                }
            }
            records = append(records, record)
        }
    }

    return records
}

func common_ids_with_bug(left []IdNameRecord, right []IdNameRecord) []string {
    /*
    Given two lists of records with the following shape:
        {
            id:string,
            name:string
        }
    return a tuple of all the IDs that are common to both lists

    Warning: This version has a bug!!!
        1. What is the bug?
        2. How would you solve it?
        3. How would you test it?
    */

    var commonElements []IdNameRecord

    for _, leftElement := range left {
        for _, rightElement := range right {
            if leftElement == rightElement {
                commonElements = append(commonElements, leftElement)
            }
        }
    }

    var commonIds []string
    for _, element := range commonElements {
        commonIds = append(commonIds, element.id)
    }
    return commonIds
}

func main() {

    if len(os.Args) < 3 {
        log.Fatal("Usage: " + os.Args[0] + " <left.csv> <right.csv>")
    }

    leftFileName := os.Args[1]
    rightFileName := os.Args[2]

    leftRecords := readCsvToIdNameRecords(leftFileName)
    rightRecords := readCsvToIdNameRecords(rightFileName)

    common := common_ids_with_bug(leftRecords, rightRecords)

    fmt.Println("IDs that are common to", leftFileName, "and", rightFileName)
    for _, id := range common {
        fmt.Println("    ", id)
    }

}
