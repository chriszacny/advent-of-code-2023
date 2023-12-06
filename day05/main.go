package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	lowerBound int
	upperBound int
	offset     int
}

type SourceToDestinationRangeMap struct {
	source     Range
	sourceName string
	dest       Range
	destName   string
}

type SectionMapLine struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

type SectionMap struct {
	name string
	data []SectionMapLine
}

/***********************************************************************************
 * Struct Name: ParsedSectionMap
 *
 ***********************************************************************************/
type ParsedSectionMap struct {
	name string
	next string
	data []SourceToDestinationRangeMap
}

func (p *ParsedSectionMap) GetDestId(sourceid int) int {
	idToUse := sourceid
	for _, rm := range p.data {
		if idToUse >= rm.source.lowerBound && idToUse <= rm.source.upperBound {
			idToUse = rm.dest.lowerBound + (idToUse - rm.source.lowerBound)
			break
		}
	}
	return idToUse
}

/***********************************************************************************
 * End of ParsedSectionMap
 ***********************************************************************************/

/***********************************************************************************
 * Struct Name: InputParser
 *
 ***********************************************************************************/
type InputParser struct {
	rawInput          string
	splitIn           []string
	seeds             []int
	sectionMaps       map[string]SectionMap
	parsedSectionMaps map[string]ParsedSectionMap
}

func NewInputParser(input string) *InputParser {
	ip := InputParser{}
	ip.rawInput = input
	ip.splitIn = strings.Split(input, "\n")
	ip.sectionMaps = make(map[string]SectionMap)
	ip.parsedSectionMaps = make(map[string]ParsedSectionMap)
	return &ip
}

func (p *InputParser) GetSeeds() []int {
	if len(p.seeds) == 0 {
		for _, line := range p.splitIn {
			if strings.HasPrefix(line, "seeds:") {
				seedString := strings.Split(line, ":")[1]
				seedStrings := strings.Split(seedString, " ")
				for _, s := range seedStrings {
					if s == "" {
						continue
					}
					seedInt, err := strconv.Atoi(s)
					if err != nil {
						panic("")
					}
					p.seeds = append(p.seeds, seedInt)
				}
				break
			}
		}
	}
	return p.seeds
}

func (p *InputParser) GetSeedsPartTwo() []int {
	if len(p.seeds) == 0 {
		for _, line := range p.splitIn {
			if strings.HasPrefix(line, "seeds:") {
				seedString := strings.Split(line, ":")[1]
				seedStrings := strings.Split(seedString, " ")
				sanitized := []string{}
				for _, s := range seedStrings {
					if s != "" {
						sanitized = append(sanitized, s)
					}
				}

				seedIds := []int{}
				seedRangeLengths := []int{}
				//fmt.Printf("STRS: %v", sanitized)

				for i, s := range sanitized {
					seedInt, err := strconv.Atoi(s)
					if err != nil {
						panic("")
					}
					if i%2 == 0 {
						seedIds = append(seedIds, seedInt)
					} else {
						seedRangeLengths = append(seedRangeLengths, seedInt)
					}
				}

				for i := 0; i < len(seedIds); i++ {
					for j := 0; j < seedRangeLengths[i]; j++ {
						p.seeds = append(p.seeds, seedIds[i]+j)
					}
				}

				break
			}
		}
	}
	fmt.Printf("LEN OF SEEDS: %d", len(p.seeds))
	return p.seeds
}

func (p *InputParser) GetSectionString(s string) string {
	toReturn := ""
	inSection := false
	for _, line := range p.splitIn {
		//fmt.Printf("DEBUG: %v\n", line)
		if strings.HasPrefix(line, s) || inSection {
			if line == "" || (!strings.HasPrefix(line, s) && strings.Contains(line, "map")) {
				break
			}
			inSection = true
			toReturn += fmt.Sprintf("%s\n", line)
		}
	}
	return toReturn
}

func (p *InputParser) GetSectionMap(s string) SectionMap {
	m, ok := p.sectionMaps[s]
	if !ok {
		// create and return it
		p.sectionMaps[s] = SectionMap{name: s}
		sectionString := p.GetSectionString(s)
		//fmt.Printf("DEBUG: %v\n", sectionString)
		//fmt.Printf("DEBUG: ***********************\n")
		sectionLineStrs := strings.Split(sectionString, "\n")
		for _, lineStr := range sectionLineStrs {
			if lineStr == "" || lineStr == "\n" || strings.HasPrefix(lineStr, s) {
				continue
			}
			sectionLine := SectionMapLine{}
			dataStrs := strings.Split(lineStr, " ")

			destRangeStart, err := strconv.Atoi(dataStrs[0])

			if err != nil {
				panic("a")
			}
			sectionLine.destinationRangeStart = destRangeStart

			sourceRangeStart, err := strconv.Atoi(dataStrs[1])
			if err != nil {
				panic("b")
			}
			sectionLine.sourceRangeStart = sourceRangeStart

			rangeLength, err := strconv.Atoi(dataStrs[2])
			if err != nil {
				panic("c")
			}
			sectionLine.rangeLength = rangeLength

			if entry, ok := p.sectionMaps[s]; ok {
				entry.data = append(p.sectionMaps[s].data, sectionLine)
				p.sectionMaps[s] = entry
			}
		}
		return p.sectionMaps[s]
	}
	return m
}

func (p *InputParser) GetNext(s string) string {
	if s == "seed-to-soil" {
		return "soil-to-fertilizer"
	} else if s == "soil-to-fertilizer" {
		return "fertilizer-to-water"
	} else if s == "fertilizer-to-water" {
		return "water-to-light"
	} else if s == "water-to-light" {
		return "light-to-temperature"
	} else if s == "light-to-temperature" {
		return "temperature-to-humidity"
	} else if s == "temperature-to-humidity" {
		return "humidity-to-location"
	}
	return ""
}

func (p *InputParser) GetParsedSectionMap(s string) ParsedSectionMap {
	m, ok := p.parsedSectionMaps[s]
	if !ok {
		p.parsedSectionMaps[s] = ParsedSectionMap{name: s, next: p.GetNext(s)}
		splitName := strings.Split(s, "-")
		sm := p.GetSectionMap(s)
		for _, d := range sm.data {
			sourceRangeStart := d.sourceRangeStart
			sourceRangeEnd := d.sourceRangeStart + d.rangeLength - 1
			destRangeStart := d.destinationRangeStart
			destRangeEnd := d.destinationRangeStart + d.rangeLength - 1
			offset := d.destinationRangeStart - d.sourceRangeStart // offset from source to destination (add offset to source to get dest)

			newSourceRange := Range{lowerBound: sourceRangeStart, upperBound: sourceRangeEnd, offset: offset}
			newDestRange := Range{lowerBound: destRangeStart, upperBound: destRangeEnd, offset: offset}

			newRM := SourceToDestinationRangeMap{}
			newRM.source = newSourceRange
			newRM.sourceName = splitName[0]
			newRM.dest = newDestRange
			newRM.destName = splitName[2]

			if entry, ok := p.parsedSectionMaps[s]; ok {
				entry.data = append(p.parsedSectionMaps[s].data, newRM)
				p.parsedSectionMaps[s] = entry
			}
		}
		return p.parsedSectionMaps[s]
	}
	return m
}

func (p *InputParser) Traverse() []int {
	seeds := p.GetSeeds()
	toReturn := []int{}
	for _, seedId := range seeds {
		locationId := p.TraversePath(seedId, "seed", "soil")
		toReturn = append(toReturn, locationId)
	}
	return toReturn
}

func (p *InputParser) TraversePartTwo() []int {
	seeds := p.GetSeedsPartTwo()
	toReturn := []int{}
	for i, seedId := range seeds {
		locationId := p.TraversePath(seedId, "seed", "soil")
		toReturn = append(toReturn, locationId)
		if i%1000000 == 0 {
			fmt.Printf("1,000,000 complete\n")
		}
	}
	return toReturn
}

func (p *InputParser) TraversePath(thingId int, sourceName string, destName string) int {
	mapname := sourceName + "-to-" + destName
	psm := p.GetParsedSectionMap(mapname)
	destId := psm.GetDestId(thingId)
	if psm.next == "" {
		return destId
	} else {
		return p.TraversePath(destId, destName, strings.Split(psm.next, "-")[2])
	}
}

/***********************************************************************************
 * End of InputParser
 ***********************************************************************************/

func main() {
	// Boilerplate setup
	in, file := getInput()
	defer file.Close()

	p := NewInputParser(in)
	results := p.Traverse()
	slices.Sort(results)
	fmt.Println(results[0])
	// Part two ///////////////////////
	p2 := NewInputParser(in)
	results = p2.TraversePartTwo()
	slices.Sort(results)
	fmt.Println(results[0])
}
