package hub

// Notif is presented a event between gmoudles
type Notif int

const (
	Notif_app_Start = iota
	Notif_mainui_LevelStart
	Notif_LevelOver
	Notif_LevelEnd
)
