package npc

import (
	generator "github/Smol-Brain/npc-generator-be"
	wr "github.com/mroth/weightedrand"
	"log"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func generateNpc() generator.Npc {
	rand.Seed(time.Now().UTC().UnixNano())
	npc := generator.Npc{}

	// Passing in slices for now as opposed to reading from db or file
	// TODO: Replace with external sources
	npc.Gender = generateWeightedOneFromChoices(wr.Choice{Item: "Male", Weight: 4}, wr.Choice{Item: "Female", Weight: 4}, wr.Choice{Item: "Non-Binary", Weight: 1})
	npc.Height = generateOneFromChoices([]string{"Tiny", "Short", "Average", "Tall", "Giant"})
	npc.Languages = generateManyFromChoices([]string{"Dwarvish", "Elvish", "Giant", "Gnomish", "Goblin", "Halfling", "Orc"}, 3)
	npc.LifeStage = generateOneFromChoices([]string{"Young Adult", "Adult", "Elderly", "Ancient"})
	npc.Race = generateOneFromChoices([]string{"Dragonborn", "Dwarf", "Elf", "Gnome", "Half-Elf", "Half-Orc", "Halfling", "Human", "Tiefling"})
	npc.Wealth = generateOneFromChoices([]string{"Poor", "Average", "Well-off", "Wealthy"})

	// Tie pronouns and names to gender
	if npc.Gender == "Male" {
		npc.Pronouns = generateWeightedOneFromChoices(wr.Choice{Item: "He/Him", Weight: 4}, wr.Choice{Item: "He/They", Weight: 1})
		npc.FirstName = generateOneFromFile("./files/names_male.csv")
	} else if npc.Gender == "Female" {
		npc.Pronouns = generateWeightedOneFromChoices(wr.Choice{Item: "She/Her", Weight: 4}, wr.Choice{Item: "She/They", Weight: 1})
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

func generateWeightedOneFromChoices(choices ...wr.Choice) string {
	chooser, _ := wr.NewChooser(choices)

	result := chooser.Pick().(string)
	return result
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
