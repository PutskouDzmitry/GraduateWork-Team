package model

type RequestRouters struct {
	Id       int64           `json:"id"`
	Coords   CoordsRouters   `json:"coords"`
	Settings SettingsRouters `json:"settings"`
}

type RequestFlux struct {
	Steps     []CoordsForFlux `json:"steps"`
	AcsParsed []AcsParsed     `json:"acsParsed"`
}

type CoordsForFlux struct {
	Id     int64         `json:"id"`
	Coords CoordsRouters `json:"coords"`
}

type AcsParsed struct {
	Id      int64           `json:"id"`
	Signals []SignalsOnFlux `json:"signals"`
}

type SignalsOnFlux struct {
	Id  int64        `json:"id"`
	Obj SignalOnFlux `json:"obj"`
}

type SignalOnFlux struct {
	Id                 int64  `json:"id"`
	AdId               string `json:"AT_ID"`
	MAC                string `json:"MAC"`
	LastSignalStrength string `json:"LastSignalStrength"`
}

type SettingsRouters struct {
	TransmitterPower          string `json:"transmitterPower"`
	GainOfTransmittingAntenna string `json:"gainOfTransmittingAntenna"`
	GainOfReceivingAntenna    string `json:"gainOfReceivingAntenna"`
	Speed                     string `json:"speed"`
	SignalLossTransmitting    string `json:"signalLossTransmitting"`
	SignalLossReceiving       string `json:"signalLossReceiving"`
	NumberOfChannels          string `json:"numberOfChannels"`
}

type CoordsRouters struct {
	X float64 `json:"left"`
	Y float64 `json:"top"`
}

type RouterSettings struct {
	RouterName                string            `json:"router_name"`
	CoordinatesOfRouter       CoordinatesPoints `json:"coordinates_of_router"`
	TransmitterPower          float64           `json:"transmitter_power"`
	GainOfTransmittingAntenna float64           `json:"gain_of_transmitting_antenna"`
	GainOfReceivingAntenna    float64           `json:"gain_of_receiving_antenna"`
	Speed                     int               `json:"speed"`
	SignalLossTransmitting    float64           `json:"signal_loss_transmitting"`
	SignalLossReceiving       float64           `json:"signal_loss_receiving"`
	TypeOfSignal              float64           `json:"type_of_signal"`
	NumberOfChannels          int               `json:"number_of_channels"`
	Scale                     float64           `json:"scale"`
	COM                       float64           `json:"com"`
}

type Wifi struct {
	User       int64
	Router     []RouterSettings
	PathInput  string
	PathOutput string
}

type WifiResponse struct {
	User       int64
	Router     []RouterSettings
	PathInput  string
	PathOutput string
}

type CoordinatesPoints struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type RoutersSettingForMigrator struct {
	Coordinates              CoordinatesPoints
	RoutersSettingsMigration []RouterSettingForMigrator
}

type RouterSettingForMigrator struct {
	Name  string
	Power float64
	MAC   string
}

type ResponseOfGettingStatisticsOnPoint struct {
	Name          string  `json:"name"`
	MAC           string  `json:"mac"`
	SignalStrange float64 `json:"signal_strange"`
	Frequency     float64 `json:"frequency"`
	MaxSpeed      float64 `json:"max_speed"`
}

type RequestAcrylicPicture struct {
	Steps         []CoordsForFlux `json:"steps"`
	AcrylicParsed []AcrylicParsed `json:"acrylicParsed"`
}

type AcrylicParsed struct {
	Id         int64  `json:"id"`
	ParsedText string `json:"parsedText"`
}
