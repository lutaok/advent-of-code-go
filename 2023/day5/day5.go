package aoc2023day5

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type SeedTransformer struct {
	destination int
	source      int
	seedRange   int
}

func LowestSeedNumber() int {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	textInput, err := os.ReadFile(currentDir + "/2023/day5/day5-input.txt")

	if err != nil {
		log.Fatal(err)
	}

	buffer := textInput

	var seeds []int

	seedsRegexp := regexp.MustCompile(`seeds: `)
	seedsToSoilRegexp := regexp.MustCompile(`seed-to-soil map:`)
	soilToFertilizerRegexp := regexp.MustCompile(`soil-to-fertilizer map:`)
	fertilizerToWaterRegexp := regexp.MustCompile(`fertilizer-to-water map:`)
	waterToLightRegexp := regexp.MustCompile(`water-to-light map:`)
	lightToTemperatureRegexp := regexp.MustCompile(`light-to-temperature map:`)
	temperatureToHumidityRegexp := regexp.MustCompile(`temperature-to-humidity map:`)
	humidityToLocationRegexp := regexp.MustCompile(`humidity-to-location map:`)

	var soilTransformers []SeedTransformer
	var fertilizerTransformers []SeedTransformer
	var waterTransformers []SeedTransformer
	var lightTransformers []SeedTransformer
	var temperatureTransformers []SeedTransformer
	var humidityTransformers []SeedTransformer
	var locationTransformers []SeedTransformer

	seedsParsing := false
	soilParsing := false
	fertilizerParsing := false
	waterParsing := false
	lightParsing := false
	temperatureParsing := false
	humidityParsing := false

	for {
		advance, token, err := bufio.ScanLines(buffer, true)

		if err != nil {
			log.Fatal(err)
		}

		if advance == 0 {
			break
		}

		if advance <= len(buffer) {
			if seedsRegexp.FindString(string(token)) != "" {
				seedsText := strings.Split(seedsRegexp.ReplaceAllString(string(token), ""), " ")

				seeds = mapStringToInt(seedsText)
			}

			if seedsToSoilRegexp.FindString(string(token)) != "" {
				seedsParsing = true
				buffer = buffer[advance:]
				continue
			}

			if soilToFertilizerRegexp.FindString(string(token)) != "" {
				seedsParsing = false
				soilParsing = true
				buffer = buffer[advance:]
				continue
			}

			if fertilizerToWaterRegexp.FindString(string(token)) != "" {
				soilParsing = false
				seedsParsing = false
				fertilizerParsing = true
				buffer = buffer[advance:]
				continue
			}

			if waterToLightRegexp.FindString(string(token)) != "" {
				soilParsing = false
				seedsParsing = false
				fertilizerParsing = false
				waterParsing = true
				buffer = buffer[advance:]
				continue
			}
			if lightToTemperatureRegexp.FindString(string(token)) != "" {
				soilParsing = false
				seedsParsing = false
				fertilizerParsing = false
				waterParsing = false
				lightParsing = true
				buffer = buffer[advance:]
				continue
			}
			if temperatureToHumidityRegexp.FindString(string(token)) != "" {
				soilParsing = false
				seedsParsing = false
				fertilizerParsing = false
				waterParsing = false
				lightParsing = false
				temperatureParsing = true
				buffer = buffer[advance:]
				continue
			}
			if humidityToLocationRegexp.FindString(string(token)) != "" {
				soilParsing = false
				seedsParsing = false
				fertilizerParsing = false
				waterParsing = false
				lightParsing = false
				temperatureParsing = true
				humidityParsing = true
				buffer = buffer[advance:]
				continue
			}

			if seedsParsing {
				seedsToSoilText := strings.Split(string(token), " ")

				seedsToSoilInput := mapStringToInt(seedsToSoilText)

				soilTransformers = append(soilTransformers, mapNumbersToSeedTransformer(seedsToSoilInput))
			}

			if soilParsing {
				soilToFertilizerText := strings.Split(string(token), " ")

				soilToFertilizerInput := mapStringToInt(soilToFertilizerText)

				fertilizerTransformers = append(fertilizerTransformers, mapNumbersToSeedTransformer(soilToFertilizerInput))
			}

			if fertilizerParsing {
				fertilizerToWaterText := strings.Split(string(token), " ")

				fertilizerToWaterInput := mapStringToInt(fertilizerToWaterText)

				waterTransformers = append(waterTransformers, mapNumbersToSeedTransformer(fertilizerToWaterInput))
			}

			if waterParsing {
				waterToLightText := strings.Split(string(token), " ")

				waterToLightInput := mapStringToInt(waterToLightText)

				lightTransformers = append(lightTransformers, mapNumbersToSeedTransformer(waterToLightInput))
			}

			if lightParsing {
				lightToTemperatureText := strings.Split(string(token), " ")

				lightToTemperatureInput := mapStringToInt(lightToTemperatureText)

				temperatureTransformers = append(temperatureTransformers, mapNumbersToSeedTransformer(lightToTemperatureInput))
			}

			if temperatureParsing {
				temperatureToHumidityText := strings.Split(string(token), " ")

				temperatureToHumidityInput := mapStringToInt(temperatureToHumidityText)

				humidityTransformers = append(humidityTransformers, mapNumbersToSeedTransformer(temperatureToHumidityInput))
			}

			if humidityParsing {
				humidityToLocationText := strings.Split(string(token), " ")
				humidityToLocationInput := mapStringToInt(humidityToLocationText)

				locationTransformers = append(locationTransformers, mapNumbersToSeedTransformer(humidityToLocationInput))
			}

			buffer = buffer[advance:]
		}
	}

	var seedNumbers []int
	var seedLengths []int

	for x, value := range seeds {
		if x%2 == 0 {
			seedNumbers = append(seedNumbers, value)
		} else {
			seedLengths = append(seedLengths, value)
		}
	}

	lowestSeedLocation := -1
	for i, value := range seedNumbers {
		// For part 1 you just need to iterate on seeds and compute mapping on that value
		j := value
		maxRange := seedLengths[i]
		for j < value+maxRange {
			transformedSoil := j

			// Soil
			soilMapperIndex := findSeedTransformerIndex(transformedSoil, soilTransformers)

			if soilMapperIndex > -1 {
				transformedSoil = computeSeedMappedValue(soilTransformers[soilMapperIndex], transformedSoil)
			}

			// Fertilizer
			fertilizerMapperIndex := findSeedTransformerIndex(transformedSoil, fertilizerTransformers)

			if fertilizerMapperIndex > -1 {
				transformedSoil = computeSeedMappedValue(fertilizerTransformers[fertilizerMapperIndex], transformedSoil)
			}

			// Water
			waterMapperIndex := findSeedTransformerIndex(transformedSoil, waterTransformers)

			if waterMapperIndex > -1 {
				transformedSoil = computeSeedMappedValue(waterTransformers[waterMapperIndex], transformedSoil)
			}

			// Light
			lightMapperIndex := findSeedTransformerIndex(transformedSoil, lightTransformers)

			if lightMapperIndex > -1 {
				transformedSoil = computeSeedMappedValue(lightTransformers[lightMapperIndex], transformedSoil)
			}

			// Temperature
			temperatureMapperIndex := findSeedTransformerIndex(transformedSoil, temperatureTransformers)

			if temperatureMapperIndex > -1 {
				transformedSoil = computeSeedMappedValue(temperatureTransformers[temperatureMapperIndex], transformedSoil)
			}

			// Temperature
			humidityMapperIndex := findSeedTransformerIndex(transformedSoil, humidityTransformers)

			if humidityMapperIndex > -1 {
				transformedSoil = computeSeedMappedValue(humidityTransformers[humidityMapperIndex], transformedSoil)
			}

			// Temperature
			locationMapperIndex := findSeedTransformerIndex(transformedSoil, locationTransformers)

			if locationMapperIndex > -1 {
				transformedSoil = computeSeedMappedValue(locationTransformers[locationMapperIndex], transformedSoil)
			}

			if transformedSoil < lowestSeedLocation || lowestSeedLocation == -1 {
				lowestSeedLocation = transformedSoil
			}

			j++
		}
	}

	return lowestSeedLocation
}

func mapStringToInt(arr []string) []int {
	var arrNumbers []int

	for _, v := range arr {
		valNumber, _ := strconv.Atoi(v)

		arrNumbers = append(arrNumbers, valNumber)
	}

	return arrNumbers
}

func mapNumbersToSeedTransformer(arr []int) SeedTransformer {
	var seedTransformer SeedTransformer
	if len(arr) == 3 {
		seedTransformer.destination = arr[0]
		seedTransformer.source = arr[1]
		seedTransformer.seedRange = arr[2]
	}

	return seedTransformer
}

func findSeedTransformerIndex(currentValue int, arr []SeedTransformer) int {
	return slices.IndexFunc(arr, func(st SeedTransformer) bool {
		return currentValue >= st.source && currentValue < st.source+st.seedRange
	})
}

func computeSeedMappedValue(st SeedTransformer, currentValue int) int {
	return st.destination + currentValue - st.source
}
