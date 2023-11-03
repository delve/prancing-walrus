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
	assignments["discoDustinJ"] = assignment{
		gameName:        "DustinJ",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["discoChaos"] = assignment{
		gameName:        "Chaos",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: true,
	}
	assignments["discoBionicTurkey"] = assignment{
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
	assignments["discoTraveler000"] = assignment{
		gameName:        "Traveler000",
		role:            "Prospector",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["discoSatyrislug"] = assignment{
		gameName:        "Satyrislug",
		role:            "Prospector",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["discogooberr"] = assignment{
		gameName:        "gooberr",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["discoRooRoo"] = assignment{
		gameName:        "RooRoo",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["discoShiviro"] = assignment{
		gameName:        "Shiviro",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["discoSeven"] = assignment{
		gameName:        "Seven",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["discoGoodie"] = assignment{
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
	assignments["discoDarkIncubus"] = assignment{
		gameName:        "DarkIncubus",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["discoPwese23"] = assignment{
		gameName:        "Pwese23",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["discoTraveleryee"] = assignment{
		gameName:        "Traveleryee",
		role:            "Vanguard",
		gather:          "a",
		canUseClamMagic: false,
	}
	assignments["discoRexromos"] = assignment{
		gameName:        "Rexromos",
		role:            "Prospector",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["discoTheTraveler"] = assignment{
		gameName:        "TheTraveler",
		role:            "Prospector",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["discoDannySpeed"] = assignment{
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
	assignments["discoWaffleDuck"] = assignment{
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
	assignments["discoGoldSpaceman"] = assignment{
		gameName:        "GoldSpaceman",
		role:            "Prospector",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["discoStSteve"] = assignment{
		gameName:        "StSteve",
		role:            "Prospector",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["discoanzo6910"] = assignment{
		gameName:        "anzo6910",
		role:            "Vanguard",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["discoSlipstream"] = assignment{
		gameName:        "Slipstream",
		role:            "Vanguard",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["disco0069"] = assignment{
		gameName:        "0069",
		role:            "Vanguard",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["discoTaytertot"] = assignment{
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
	assignments["discoAlex"] = assignment{
		gameName:        "Alex",
		role:            "Vanguard",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["discoRorstark"] = assignment{
		gameName:        "Rorstark",
		role:            "Vanguard",
		gather:          "b",
		canUseClamMagic: false,
	}
	assignments["discobenc247"] = assignment{
		gameName:        "benc247",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["discostoopoodoopoo"] = assignment{
		gameName:        "stoopoodoopoo",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["discoLishaBourne"] = assignment{
		gameName:        "LishaBourne",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["discoVenom"] = assignment{
		gameName:        "Venom",
		role:            "Vanguard",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["discosluginthemidst"] = assignment{
		gameName:        "sluginthemidst",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["discojustcallmeyd"] = assignment{
		gameName:        "justcallmeyd",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["disconiagrafalls"] = assignment{
		gameName:        "niagrafalls",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["discoAszuzsa"] = assignment{
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
	assignments["discoMegümin"] = assignment{
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
	assignments["discoJHZ1"] = assignment{
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
		role:            "Vanguard",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["discokurama215"] = assignment{
		gameName:        "kurama215",
		role:            "Vanguard",
		gather:          "γ",
		canUseClamMagic: false,
	}
	assignments["discolucifer"] = assignment{
		gameName:        "lucifer",
		role:            "Prospector",
		gather:          "γ",
		canUseClamMagic: false,
	}
}
