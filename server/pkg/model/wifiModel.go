package model

type RouterSettings struct {
	CoordinatesOfRouter       CoordinatesPoints `json:"coordinates_of_router"`
	TransmitterPower          float64           `json:"transmitter_power"`
	GainOfTransmittingAntenna float64           `json:"gain_of_transmitting_antenna"`
	GainOfReceivingAntenna    float64           `json:"gain_of_receiving_antenna"`
	Speed                     int               `json:"speed"`
	SignalLossTransmitting    float64           `json:"signal_loss_transmitting"`
	SignalLossReceiving       float64           `json:"signal_loss_receiving"`
	NumberOfChannels          int               `json:"number_of_channels"`
	Scale                     float64           `json:"scale"`
	Thickness                 float64           `json:"thickness"`
	COM                       float64           `json:"com"`
}

type Wifi struct {
	User   int64
	Router []RouterSettings
	Path   string
}

type CoordinatesPoints struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Response struct {
}
