package model

type CoordinatesAllSchemes struct {
	CoordinatesRouter         CoordinatesPoints               `json:"coordinates-router"`
	CoordinatesForCalculate   []CoordinatesPointsForCalculate `json:"coordinates_for_calculate"`
	TransmitterPower          float64                         `json:"transmitter_power"`
	GainOfTransmittingAntenna float64                         `json:"gain_of_transmitting_antenna"`
	GainOfReceivingAntenna    float64                         `json:"gain_of_receiving_antenna"`
	Speed                     int                             `json:"speed"`
	SignalLossTransmitting    float64                         `json:"signal_loss_transmitting"`
	SignalLossReceiving       float64                         `json:"signal_loss_receiving"`
	NumberOfChannels          int                             `json:"number_of_channels"`
}

type CoordinatesPoints struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type CoordinatesPointsForCalculate struct {
	Coordinates   CoordinatesPoints `json:"coordinates"`
	NumberOfWalls uint64            `json:"number_of_walls"`
	Distance      float64           `json:"distance"`
}

type Response struct {
}
