package msg

// MaliciousCategoryLong is a map from malicious category to long string names
var MaliciousCategoryLong = map[MaliciousCategory]string{
	MaliciousCategory_BOTS:    "Botnet",
	MaliciousCategory_CANDC:   "Command and Control Centers",
	MaliciousCategory_COMPR:   "Compromised & Links To Malware",
	MaliciousCategory_MALHOME: "Malware Call-Home",
	MaliciousCategory_MAL:     "Malware Distribution Point",
	MaliciousCategory_FRAUD:   "Phishing/Fraud",
	MaliciousCategory_SPAM:    "Spam URLs",
	MaliciousCategory_SPYWARE: "Spyware & Questionable Software",
	MaliciousCategory_UNK:     "Unknown",
}
