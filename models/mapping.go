package models

var (
	PictureActionMapping    = map[int64]string{0: "PASS", 1: "REVIEW", 2: "REJECT"}
	DunContentModelsMapping = map[string]string{
		"PORN":     "100",
		"AD":       "200",
		"FEAR":     "300",
		"ILLEGAL":  "400",
		"POLITICS": "500",
		"ABUSE":    "600",
		"USELESS":  "700",
		"WORTH":    "1100",
	}
	DunPictureModelsMapping = map[string]string{
		"PORN":     "100",
		"SEXY":     "110",
		"AD":       "200",
		"QRCODE":   "210",
		"ILLEGAL":  "260",
		"FEAR":     "300",
		"ABUSE":    "400",
		"POLITICS": "500",
		"WORTH":    "1100",
	}
)
