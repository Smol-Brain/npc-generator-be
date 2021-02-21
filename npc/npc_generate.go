package npc

import (
	generator "github/Smol-Brain/npc-generator-be"
	"log"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func generateNpc() generator.Npc {
	rand.Seed(time.Now().Unix())
	npc := generator.Npc{}

	// Passing in slices for now as opposed to reading from db or file
	// TODO: Replace with external sources
	npc.Gender = generateOneFromChoices([]string{"Male", "Female", "Non-Binary"})
	npc.Height = generateOneFromChoices([]string{"Tiny", "Short", "Average", "Tall", "Giant"})
	npc.Languages = generateManyFromChoices([]string{"Dwarvish", "Elvish", "Giant", "Gnomish", "Goblin", "Halfling", "Orc"}, 3)
	npc.LifeStage = generateOneFromChoices([]string{"Young Adult", "Adult", "Elderly", "Ancient"})
	npc.Race = generateOneFromChoices([]string{"Dragonborn", "Dwarf", "Elf", "Gnome", "Half-Elf", "Half-Orc", "Halfling", "Human", "Tiefling"})
	npc.Wealth = generateOneFromChoices([]string{"Poor", "Average", "Well-off", "Wealthy"})

	// Tie pronouns and names to gender (for now?)
	// TODO: Replace with percent likelihoods based on gender (e.g. Females have a 90% chance of She/Her, 10% chance of other pronouns)
	if npc.Gender == "Male" {
		npc.Pronouns = "He/Him"
		npc.FirstName = generateOneFromFile("./files/names_male.csv")
	} else if npc.Gender == "Female" {
		npc.Pronouns = "She/Her"
		npc.FirstName = generateOneFromFile("./files/names_female.csv")
	} else if npc.Gender == "Non-Binary" {
		npc.Pronouns = "They/Them"
		npc.FirstName = generateOneFromFile("./files/names_neutral.csv")
	} else {
		npc.Pronouns = "They/Them"
		npc.FirstName = generateOneFromFile("./files/names_neutral.csv")
	}

	// Since Common is a widespread language, weigh it more heavily or add it always if no other languages present
	addCommon := rand.Intn(10)
	if addCommon > 0 || len(npc.Languages) == 0 {
		npc.Languages = append(npc.Languages, "Common")
	}

	npc.Quirk = generateOneFromFile("./files/quirks.csv")
	npc.Job = generateOneFromFile("./files/jobs.csv")
	npc.NegativeTraits = generateManyFromFile("./files/negative_traits.csv", 3)
	npc.NeutralTraits = generateManyFromFile("./files/neutral_traits.csv", 3)
	npc.PositiveTraits = generateManyFromFile("./files/positive_traits.csv", 3)

	return npc
}

// For small, self-contained list of choices
func generateOneFromChoices(choices []string) string {
	index := rand.Intn(len(choices))
	return choices[index]
}

// Return a slice of at most max selections of choices
func generateManyFromChoices(choices []string, max int) []string {
	if max < 0 || max > len(choices) {
		max = rand.Intn(len(choices))
	} else {
		max = rand.Intn(max)
	}

	rand.Shuffle(len(choices), func(i, j int) {
		choices[i], choices[j] = choices[j], choices[i]
	})

	return choices[:max]
}

// For larger data sets
func generateOneFromFile(file string) string {
	shuf := exec.Command("shuf", "-n 1", file)
	out, err := shuf.Output()

	if err != nil {
		log.Fatal(err)
	}

	// return string(out)
	return strings.ReplaceAll(string(out), "\n", "")
}

func generateManyFromFile(file string, max int) []string {
	// We want at least one choice returned
	max = rand.Intn(max) + 1

	shuf := exec.Command("shuf", "-n", strconv.Itoa(max), file)
	out, err := shuf.Output()

	if err != nil {
		log.Fatal(err)
	}

	fields := strings.Split(string(out), "\n")
	// Hacky way of removing last empty string element from split
	return fields[:len(fields)-1]
}
