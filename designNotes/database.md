# Players
| Field      | Description                             |
|------------|-----------------------------------------|
| PlayerID   | Sequence number                         |
| DiscoId    | discord user ID (immutable)             |


# Snails

| Field              | Description                                 |
|--------------------|---------------------------------------------|
| PlayerID           | Players FK                                  |
| SnailId            | Sequence number                             |
| Updated            | unix epoch of last record update            |
| SnailName          | game name                                   |
| Club               | which club, if any, they're a member of     |
| Server             | which server name they're in (enum)         |
| ServerNum          | which numeric server they're in             |
| Leadership         | number                                      |
| Art                | number                                      |
| Fth                | number                                      |
| Fame               | number                                      |
| Civ                | number                                      |
| Tech               | number                                      |
| Hp                 | number                                      |
| Atk                | number                                      |
| Rush               | number                                      |
| Def                | number                                      |
| TotalPower         | accept string (eg 13.0M), convert to number |
| ZombieForm         | number (0 means not unlocked)               |
| DemonForm          | number (0 means not unlocked)               |
| AngelForm          | number (0 means not unlocked)               |
| MutantForm         | number (0 means not unlocked)               |
| MechaForm          | number (0 means not unlocked)               |
| DragonForm         | number (0 means not unlocked)               |
| SpeciesWarEssences | number                                      |

# MinionStats
TBD

# Other Things
TBDs