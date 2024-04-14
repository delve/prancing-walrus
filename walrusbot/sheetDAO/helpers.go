package sheetDAO

func GetPlayerClubMemberships(discordName string) (clubs []*Club, err error) {
	clubs = []*Club{}
	err = nil
	player, err := GetPlayerByDiscoId(discordName)
	if err != nil {
		return nil, err
	}

	// ignore the 'none' club
	clubFilter := func(snail *Snail) bool { return snail.Club > 1 }
	snails, err := player.GetSnails(SnailFilter(clubFilter))
	if err != nil {
		return nil, err
	}
	if len(snails) == 0 {
		return nil, nil
	}

	for _, snail := range snails {
		cb, err := GetClub(snail.Club)
		if err != nil {
			return nil, err
		}

		clubs = append(clubs, cb)
	}
	return
}
