package actions

import "github.com/bwmarrin/discordgo"

// Gives info about the six seven event
func SixEvenInfo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	interactionTextResponse("67 event: Triggers when a user includes "+"67"+" in their message.\n"+
		"*Drop rates:*\n🟩 Common: 69.5%\n🟨 Uncommon: 20%\n🟧 Rare: 8%\n🟥 Epic: 2%\n🟪 Legendary: 0.49%\n🟫 Mythic (0.01%): VERY rare global event",
		s, i)
}
