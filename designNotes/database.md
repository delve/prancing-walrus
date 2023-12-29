# Snails
| Field      | Description                             |
|------------|-----------------------------------------|
| UserID     | Sequence number                         |
| DiscoId    | discord user ID (immutable)             |
| GameName   | game name                               |
| Club       | which club, if any, they're a member of |
| Server     | which server name they're in (enum)     |
| ServerNum  | which numeric server they're in         |


# SnailStats

| Field        | Description                                 |
|--------------|---------------------------------------------|
| UserId       | Snails FK                                   |
| SnailStatsId | Sequence number                             |
| Updated      | unix epoch of last record update            |
| Leadership   | number                                      |
| Art          | number                                      |
| Fth          | number                                      |
| Fame         | number                                      |
| Civ          | number                                      |
| Tech         | number                                      |
| Hp           | number                                      |
| Atk          | number                                      |
| Rush         | number                                      |
| Def          | number                                      |
| TotalPower   | accept string (eg 13.0M), convert to number |
| ZombieForm   | number (0 means not unlocked)               |
| DemonForm    | number (0 means not unlocked)               |
| AngelForm    | number (0 means not unlocked)               |
| MutantForm   | number (0 means not unlocked)               |
| MechaForm    | number (0 means not unlocked)               |
| DragonForm   | number (0 means not unlocked)               |

# MinionStats
TBD

# Other Things
TBDs