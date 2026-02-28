package main

import (
	"container/heap"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Selected struct {
	residentID int
	rank       int
}
type SelectedHeap []Selected

func (sh SelectedHeap) Len() int {
	return len(sh)
}
func (sh SelectedHeap) Less(i, j int) bool {
	return sh[i].rank > sh[j].rank
}
func (sh SelectedHeap) Swap(i, j int) {
	sh[i], sh[j] = sh[j], sh[i]
}

func (sh *SelectedHeap) Push(x interface{}) {
	*sh = append(*sh, x.(Selected))
}

func (sh *SelectedHeap) Pop() interface{} {
	temp := *sh
	n := len(temp)
	v := temp[n-1]
	*sh = temp[:n-1]
	return v

}

// The Resident data type
type Resident struct {
	residentID     int
	firstname      string
	lastname       string
	rol            []string // resident rank order list
	matchedProgram string   // will be "" for unmatched resident

	rolIndex int
}

// The Program data type
type Program struct {
	programID  string
	name       string
	nPositions int   // number of positions available (quota)
	rol        []int // program rank order list
	// TO ADD: a data structure for the selected resident IDs
	selectedResidents SelectedHeap
}

func offer(rID int, residents map[int]*Resident, programs map[string]*Program) {
	r := residents[rID]
	//resident matched or no programs left
	if r.matchedProgram != "" || r.rolIndex >= len(r.rol) {

		return
	}

	pID := r.rol[r.rolIndex]
	r.rolIndex++
	evaluate(rID, pID, residents, programs)
}
func evaluate(rID int, pID string, residents map[int]*Resident, programs map[string]*Program) {

	p := programs[pID]
	r := residents[rID]

	// find residents rank
	rank := -1
	for i, id := range p.rol {
		if id == rID {
			rank = i
			break
		}
	}

	// exit if unranked
	if rank == -1 {
		offer(rID, residents, programs)
		return
	}

	//accept if positions are available
	if p.selectedResidents.Len() < p.nPositions {

		heap.Push(&p.selectedResidents, Selected{residentID: rID, rank: rank})
		r.matchedProgram = pID
		return
	}

	// compare to ranked resident
	//Top of heap holds the worst resident (Less() ordering)
	worst := p.selectedResidents[0]

	// if new resident has better rank
	if rank < worst.rank {

		// remove worst (top of heap)
		removed := heap.Pop(&p.selectedResidents).(Selected)

		// removed resident resets
		residents[removed.residentID].matchedProgram = ""

		// push new resident
		heap.Push(&p.selectedResidents, Selected{residentID: rID, rank: rank})
		r.matchedProgram = pID
		//removed resident proposes again
		offer(removed.residentID, residents, programs)

	} else {
		// new resident trys next program
		offer(rID, residents, programs)
	}
}

func availablePositions(programs map[string]*Program) int {
	emptyPositions := 0
	for _, p := range programs {
		open := p.nPositions - len(p.selectedResidents)
		emptyPositions += open

	}
	return emptyPositions
}

func PrintMatches(residents map[int]*Resident, programs map[string]*Program) {
	fmt.Println("lastname,firstname,residentID,programID,name")
	unmatchedCount := 0

	for _, r := range residents {

		pID := r.matchedProgram
		if pID == "" {
			fmt.Printf("%s,%s,%d,XXX,NOT_MATCHED\n", r.lastname, r.firstname, r.residentID)
			unmatchedCount++
			continue
		}
		p := programs[pID]
		fmt.Printf("%s,%s,%d,%s,%s\n", r.lastname, r.firstname, r.residentID, pID, p.name)
	}

	fmt.Println("Number of unmatched residents:", unmatchedCount)
	fmt.Println("Number of positions available:", availablePositions(programs))
}

// Parse a resident's ROL
func parseRol(s string) []string {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	if s == "" {
		return []string{}
	}
	parts := strings.Split(s, ",")
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return parts
}

// Parse a program's ROL
func parseIntRol(s string) []int {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	if s == "" {
		return []int{}
	}
	parts := strings.Split(s, ",")
	var ints []int
	for _, part := range parts {
		pid, _ := strconv.Atoi(strings.TrimSpace(part))
		ints = append(ints, pid)
	}
	return ints
}

// ReadCSV reads a CSV file into a map of Resident
func ReadResidentsCSV(filename string) (map[int]*Resident, error) {

	// map to store residents by ID
	residents := make(map[int]*Resident)

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %w", err)
	}

	// Skip header if present (assuming it is)
	for i, record := range records {
		if i == 0 && record[0] == "id" {
			continue
		}
		if len(record) < 4 {
			return nil, fmt.Errorf("invalid record at line %d: %v", i+1, record)
		}

		// Parse ID
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("invalid ID at line %d: %w", i+1, err)
		}

		if _, exists := residents[id]; exists {
			fmt.Println(id)
		}

		residents[id] = &Resident{
			residentID:     id,
			firstname:      record[1],
			lastname:       record[2],
			rol:            parseRol(record[3]),
			matchedProgram: "",
		}
	}

	return residents, nil
}

// reads a CSV file into a map of Program
func ReadProgramsCSV(filename string) (map[string]*Program, error) {

	// map to store programs by ID
	programs := make(map[string]*Program)

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %w", err)
	}

	// Skip header if present (assuming it is)
	for i, record := range records {
		if i == 0 && record[0] == "id" {
			continue
		}
		if len(record) < 4 {
			return nil, fmt.Errorf("invalid record at line %d: %v", i+1, record)
		}

		// Parse number of positions
		np, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, fmt.Errorf("invalid number at line %d: %w", i+1, err)
		}

		programs[record[0]] = &Program{
			programID:  record[0],
			name:       record[1],
			nPositions: np,
			rol:        parseIntRol(record[3]),
		}
		heap.Init(&programs[record[0]].selectedResidents)

	}

	return programs, nil
}

// Example usage
func main() {

	// read residents
	residents, err := ReadResidentsCSV("residentsLARGE.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//for _, p := range residents {
	//	fmt.Printf("ID: %d, Name: %s %s, Rol: %v\n", p.residentID, p.firstname, p.lastname, p.rol)
	//}

	programs, err := ReadProgramsCSV("programsLARGE.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//everything commented below is not needed for desired output

	//for _, p := range programs {
	//	fmt.Printf("ID: %s, Name: %s, Number of pos: %d, Number of applicants: %d\n", p.programID, p.name, p.nPositions, len(p.rol))
	//}
	//
	//fmt.Printf("\nNMD: %v", programs["NMD"])

	for id := range residents {
		offer(id, residents, programs)
	}

	PrintMatches(residents, programs)
}
