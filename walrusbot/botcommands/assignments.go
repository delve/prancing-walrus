package botcommands

var assignments map[string]assignment

type assignment struct {
	gameName        string
	role            string
	gather          string
	canUseClamMagic bool
}

func init() {
	assignments = map[string]assignment{}
	assignments["gaze3"] = assignment{
		gameName:        "Gaze",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["iknyc"] = assignment{
		gameName:        "Chaos",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: true,
	}
	assignments["bionic_turkey"] = assignment{
		gameName:        "BionicTurkey",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["mehhhhhhhhhhhhhhhhhhhhhhhhhhhhh"] = assignment{
		gameName:        "Mehh",
		role:            "Prospector",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["na000"] = assignment{
		gameName:        "Traveler000",
		role:            "Prospector",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["satyrical1"] = assignment{
		gameName:        "Satyrislug",
		role:            "Prospector",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["twentygauge37"] = assignment{
		gameName:        "gooberr",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["roo1987"] = assignment{
		gameName:        "RooRoo",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["shiviro"] = assignment{
		gameName:        "Shiviro",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["sevensinlegion"] = assignment{
		gameName:        "Seven",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["goodie5624"] = assignment{
		gameName:        "Goodie",
		role:            "Prospector",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["discoDarthBella"] = assignment{
		gameName:        "DarthBella",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["darkincubus"] = assignment{
		gameName:        "DarkIncubus",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["thepokemonmattster"] = assignment{
		gameName:        "Pwese23",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["calbehe"] = assignment{
		gameName:        "Traveleryee",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["rexromos"] = assignment{
		gameName:        "Rexromos",
		role:            "Prospector",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["the.inTIMidator24015"] = assignment{
		gameName:        "TheTraveler",
		role:            "Prospector",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["dannyspeed_"] = assignment{
		gameName:        "DannySpeed",
		role:            "Prospector",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["discoRydia"] = assignment{
		gameName:        "Rydia",
		role:            "Prospector",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["waffleduck4990"] = assignment{
		gameName:        "WaffleDuck",
		role:            "Vanguard",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["discowordnerd"] = assignment{
		gameName:        "wordnerd",
		role:            "Prospector",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["king_za215"] = assignment{
		gameName:        "GoldSpaceman",
		role:            "Prospector",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["agentmeep"] = assignment{
		gameName:        "StSteve",
		role:            "Prospector",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["sea19anzomathew"] = assignment{
		gameName:        "anzo6910",
		role:            "Vanguard",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["slipstreamx10x"] = assignment{
		gameName:        "Slipstream",
		role:            "Vanguard",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["whitessmaro"] = assignment{
		gameName:        "0069",
		role:            "Vanguard",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["Taytertot"] = assignment{
		gameName:        "Taytertot",
		role:            "Vanguard",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["discomz13"] = assignment{
		gameName:        "mz13",
		role:            "Prospector",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["calyxalex."] = assignment{
		gameName:        "Alex",
		role:            "Vanguard",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["th31nd00rs"] = assignment{
		gameName:        "TH31NDOORS",
		role:            "Vanguard",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["benjaminc247"] = assignment{
		gameName:        "benc247",
		role:            "Vanguard",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["ikeepitnacho"] = assignment{
		gameName:        "stoopoodoopoo",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["lishabourne"] = assignment{
		gameName:        "LishaBourne",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["Venomn0us"] = assignment{
		gameName:        "Venom",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["Rye the Quiet"] = assignment{
		gameName:        "sluginthemidst",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["justcallmeyd"] = assignment{
		gameName:        "justcallmeyd",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["knivesnchains"] = assignment{
		gameName:        "niagrafalls",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["seraphic9"] = assignment{
		gameName:        "Aszuzsa",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["discourmumsdad"] = assignment{
		gameName:        "urmumsdad",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["arkane_x6"] = assignment{
		gameName:        "Megümin",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["discoLemon"] = assignment{
		gameName:        "Lemon",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["JHZ1"] = assignment{
		gameName:        "JHZ1",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["discoWslim"] = assignment{
		gameName:        "Wslim",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["discoTimmy"] = assignment{
		gameName:        "Timmy",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["kurama215"] = assignment{
		gameName:        "kurama215",
		role:            "Vanguard",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["vinnydev"] = assignment{
		gameName:        "lucifer",
		role:            "Vanguard",
		gather:          "γ",
		canUseClamMagic: false,
	}
}
