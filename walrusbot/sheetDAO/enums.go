package sheetDAO

// ServerName is what server the user is in
type ServerName string

// Server Enumeration. Need to replace this with a lookup on a table in the DB to avoid rebuilding whenever there's a new server ¬.¬
const (
	Unknown      ServerName = "Unknown"
	Hotdowog     ServerName = "Hotdowog"
	Americanowo  ServerName = "Americanowo"
	PittedOlives ServerName = "Pitted Olives"
	HalfChips    ServerName = "Half Chips"
	DeepFried    ServerName = "Deep Fried"
	AvoToast     ServerName = "Avo Toast"
	Potatowo     ServerName = "Potatowo"
	Hamborger    ServerName = "Hamborger"
	LorgFries    ServerName = "Lorg Fries"
	Cheestborger ServerName = "Cheestborger"
	ExtraCrispy  ServerName = "Extra Crispy"
	MashePotato  ServerName = "Mashed Potato"
)

func NewServerName(name string) (server ServerName, err error) {

	return ServerName(name), nil
}

func (s ServerName) String() string {
	return string(s)
}
