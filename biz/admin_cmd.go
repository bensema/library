package model

type AdminCmd int

const (
	_ AdminCmd = iota
	AdvertiseAdd
	AdvertiseDel
	AdvertiseUpdate
	AdvertiseQuery
)
