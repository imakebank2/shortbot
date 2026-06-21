package responses

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/imakebank2/shortbot/utils"
)

// SIX SEVENNNNN
// Based from ChatGPT's creativity: https://chatgpt.com/s/t_69e6046a51b8819180de7d7605fb1327

var triggers = []string{
	"67",
}

var common = []string{
	"Not equal. Not fair. Just 67.",
	"This is a 67 situation.",
	"67 initiated. Repayment pending.",
	"Close enough to confuse both parties.",
	"You thought this was mutual. It’s a 67.",
	"Balance updated: still 67.",
	"Somewhere between a deal and a mistake.",
}

var uncommon = []string{
	"Terms accepted. Terms not understood.",
	"You owe one (1) favor. This is now recorded.",
	"Payment deferred indefinitely.",
	"This handshake comes with consequences.",
	"A 67 has been established between both parties.",
	"This agreement is emotionally unstable.",
}

var rare = []string{
	"You expected 69. You received 67 customer support.",
	"Effort: high. Outcome: 67.",
	"That was almost impressive. Almost.",
	"We were so close to something meaningful.",
	"This feels like a legally questionable agreement.",
	"You almost achieved balance. The system rejected it.",
}

var epic = []string{
	"ERROR: 67 detected. Emotional stability reduced.",
	"This 67 has been reviewed and deemed unstable.",
	"You didn’t break the system. The system became 67.",
	"67 protocol has overridden standard interaction rules.",
	"This is no longer a meme. It’s a transaction",
}

var legendary = []string{
	"67 is now the default state of this interaction.",
	"You are no longer participating in 67. You are inside it.",
	"This was supposed to be a joke. It became infrastructure.",
	"When balance fails, 67 is what remains.",
	"There is no opposite of 67. Only echoes.",
	"This 67 will outlive the conversation that created it.",
	"We attempted fairness. The result was 67.",
}

var mythic = []string{
	"SYSTEM OVERRIDE: 67 has replaced the concept of agreement.",
	"You did not trigger 67. 67 has been waiting for you.",
	"This interaction has been archived as a permanent 67.",
	"All outcomes converge here. The result is 67.",
	"There is no randomness left. Only 67.",
	"You have witnessed the origin point of 67.",
	"This is the last stable version of meaning.",
}

func SixSevenEvent(s *discordgo.Session, m *discordgo.MessageCreate) {

	var key string
	var value string

	content := strings.ToLower(m.Content)

	for _, trigger := range triggers {
		if strings.Contains(content, trigger) {

			n := rand.IntN(10000)

			switch {
			case n < 6950:
				key = "🟩 COMMON — 69.5%"
				value, _ = utils.GetRandomElement(common)

			case n < 8950:
				key = "🟨 UNCOMMON — 20%"
				value, _ = utils.GetRandomElement(uncommon)

			case n < 9750:
				key = "🟧 RARE — 8%"
				value, _ = utils.GetRandomElement(rare)

			case n < 9950:
				key = "🟥 EPIC — 2%"
				value, _ = utils.GetRandomElement(epic)

			case n < 9999:
				key = "🟪 LEGENDARY EXTREMELY RARE — 0.49%"
				value, _ = utils.GetRandomElement(legendary)

			default:
				key = "🟫 MYTHIC / 0.01% (Ultra-Secret Drop 🔥)"
				value, _ = utils.GetRandomElement(mythic)
			}

			replyToUserWithoutPing(
				fmt.Sprintf("%s triggered a 67.\n***Rarity:***  **%s**\n*\"%s\"*", m.Author.DisplayName(), key, value),
				s, m,
			)

			break
		}
	}
}
