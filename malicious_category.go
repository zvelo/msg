package msg

// MaliciousCategoryLong is a map from malicious category to long string names
var MaliciousCategoryLong = map[MaliciousCategory]string{
	MaliciousCategory_BOTS_4:    "Botnet",
	MaliciousCategory_CANDC_4:   "Command and Control Centers",
	MaliciousCategory_COMPR_4:   "Compromised & Links To Malware",
	MaliciousCategory_MALHOME_4: "Malware Call-Home",
	MaliciousCategory_MAL_4:     "Malware Distribution Point",
	MaliciousCategory_FRAUD_4:   "Phishing/Fraud",
	MaliciousCategory_SPAM_4:    "Spam URLs",
	MaliciousCategory_SPYWARE_4: "Spyware & Questionable Software",
	MaliciousCategory_UNK:       "Unknown",
}
