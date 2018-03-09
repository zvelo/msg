package msg

import (
	"strconv"
	"strings"
)

// CategoryLong is a map from Category to long string names
var CategoryLong = map[Category]string{
	UNKNOWN_CATEGORY: "Unknown Category",
	ABORTION_4:       "Abortion",
	PROCHOICE_4:      "Abortion Pro Choice",
	PROLIFE_4:        "Abortion Pro Life",
	PG13_4:           "Child Inappropriate",
	GAMBLING_4:       "Gambling",
	GAY_4:            "Gay, Lesbian or Bisexual",
	LINGERIE_4:       "Lingerie, Suggestive & Pinup",
	NUDITY_4:         "Nudity",
	PORN_4:           "Pornography",
	PROFANITY_4:      "Profanity",
	RATEDR_4:         "R-Rated",
	SEX_4:            "Sex & Erotic",
	EDSEX_4:          "Sex Education",
	TOBACCO_4:        "Tobacco",
	MIL_4:            "Military",
	VIOLENCE_4:       "Violence",
	WEAPONS_4:        "Weapons",
	AGGRESSIVE_4:     "Aggressive - Other",
	FINEART_4:        "Fine Art",
	ART_4:            "Arts - Other",
	AUTOPARTS_4:      "Auto Parts",
	AUTOREPAIR_4:     "Auto Repair",
	BUYSELLCAR_4:     "Buying/Selling Cars",
	CARCULTURE_4:     "Car Culture",
	CPO_4:            "Certified Pre-Owned",
	CARCONVERT_4:     "Convertible",
	COUPE_4:          "Coupe",
	CUV_4:            "Crossover",
	DIESEL_4:         "Diesel",
	EVEHICLE_4:       "Electric Vehicle",
	HATCHBACK_4:      "Hatchback",
	HYBRIDCAR_4:      "Hybrid",
	LUXURYCAR_4:      "Luxury",
	MINIVAN_4:        "MiniVan",
	MOTORCYCLE_4:     "Motorcycles",
	OFFROAD_4:        "Off-Road Vehicles",
	PERFCAR_4:        "Performance Vehicles",
	PICKUP_4:         "Pickup",
	ROADASSIST_4:     "Road-Side Assistance",
	SEDAN_4:          "Sedan",
	TRUCK_4:          "Trucks & Accessories",
	VINTAGECAR_4:     "Vintage Cars",
	WAGON_4:          "Wagon",
	AUTOMOTIVE_4:     "Automotive - Other",
	AG_4:             "Agriculture",
	BIOTECH_4:        "Biotechnology",
	BIZSOFT_4:        "Business Software",
	BUILD_4:          "Construction",
	FORESTRY_4:       "Forestry",
	BIZGOV_4:         "Government",
	NATURE_4:         "Green Solutions & Conservation",
	FURNITURE_4:      "Home & Office Furnishings",
	HR_4:             "Human Resources",
	MFG_4:            "Manufacturing",
	MKTG_4:           "Marketing Services",
	METAL_4:          "Metals",
	PHYSEC_4:         "Physical Security",
	BIZPROD_4:        "Productivity",
	RETIREHOME_4:     "Retirement Homes & Assisted Living",
	SHIPPING_4:       "Shipping & Logistics",
	BUSINESS_4:       "Business - Other",
	JOBADVICE_4:      "Career Advice",
	CAREERPLAN_4:     "Career Planning",
	COLJOB_4:         "College",
	FINAID_4:         "Financial Aid",
	JOBFAIR_4:        "Job Fairs",
	JOBSEARCH_4:      "Job Search",
	NURSING_4:        "Nursing",
	RESUME_4:         "Resume Writing/Advice",
	SCHOLAR_4:        "Scholarships",
	TELCOMMUTE_4:     "Telecommuting",
	MILJOBS_4:        "U.S. Military",
	CAREER_4:         "Careers - Other",
	CHILDABUSE_4:     "Child Abuse Images",
	CRIMSKILL_4:      "Criminal Skills",
	HACKING_4:        "Hacking",
	HATE_4:           "Hate Speech",
	DRUGS_4:          "Illegal Drugs",
	POT_4:            "Marijuana",
	WAREZ_4:          "Piracy & Copyright Theft",
	EDUCHEAT_4:       "School Cheating",
	SELFHARM_4:       "Self Harm",
	TORRENT_4:        "Torrent Repository",
	CRIM_4:           "Criminal Activities - Other",
	ANON_4:           "Anonymizer",
	CHAT_4:           "Chat",
	GROUPS_4:         "Community Forums",
	IM_4:             "Instant Messenger",
	LOGIN_4:          "Login Screens",
	BLOG_4:           "Personal Pages & Blogs",
	PHOTO_4:          "Photo Sharing",
	PROFNET_4:        "Professional Networking",
	REDIR_4:          "Redirect",
	SOCNET_4:         "Social Networking",
	SMS_4:            "Text Messaging & SMS",
	TRANSLATOR_4:     "Translator",
	EMAIL_4:          "Web-based Email",
	CARDS_4:          "Web-based Greeting Cards",
	EDU712_4:         "7-12 Education",
	ADULTEDU_4:       "Adult Education",
	ARTHISTORY_4:     "Art History",
	COLADMIN_4:       "College Administration",
	COLLIFE_4:        "College Life",
	DLEARNING_4:      "Distance Learning",
	SCHOOL_4:         "Educational Institutions",
	EDSTUDIES_4:      "Educational Materials & Studies",
	ENGLISH_4:        "English as a 2nd Language",
	GRADSCHOOL_4:     "Graduate School",
	HOMESCHOOL_4:     "Homeschooling",
	HOMEWORK_4:       "Homework/Study Tips",
	K6EDU_4:          "K-6 Educators",
	LANGUAGE_4:       "Language Learning",
	BOOK_4:           "Literature & Books",
	PRIVSCHOOL_4:     "Private School",
	REF_4:            "Reference Materials & Maps",
	SPED_4:           "Special Education",
	BIZEDU_4:         "Studying Business",
	TUTOR_4:          "Tutoring",
	WIKI_4:           "Wikis",
	EDU_4:            "Education - Other",
	ENTNEWS_4:        "Entertainment News & Celebrity Sites",
	VENUE_4:          "Entertainment Venues & Events",
	HUMOR_4:          "Humor",
	MOVIE_4:          "Movies",
	MUSIC_4:          "Music",
	AUDIO_4:          "Streaming & Downloadable Audio",
	VIDEO_4:          "Streaming & Downloadable Video",
	TV_4:             "Television",
	ENT_4:            "Entertainment - Other",
	ADOPT_4:          "Adoption",
	TODDLER_4:        "Babies and Toddlers",
	PRESCHOOL_4:      "Daycare/Pre School",
	ELDERCARE_4:      "Eldercare",
	FAMILYNET_4:      "Family Internet",
	K6PARENT_4:       "Parenting",
	PARENTEEN_4:      "Parenting Teens",
	PREGNANCY_4:      "Pregnancy",
	SNKPARENT_4:      "Special Needs Kids",
	FAMILY_4:         "Family & Parenting - Other",
	ACCESSORY_4:      "Accessories",
	BEAUTY_4:         "Beauty",
	BODYART_4:        "Body Art",
	CLOTHING_4:       "Clothing",
	FASHNMISC_4:      "Fashion",
	JEWELRY_4:        "Jewelry",
	SWIMSUIT_4:       "Swimsuits",
	FASHION_4:        "Fashion - Other",
	ACCTG_4:          "Accounting",
	BANK_4:           "Banking",
	BGNINVEST_4:      "Beginning Investing",
	LOAN_4:           "Credit/Debt & Loans",
	FINNEWS_4:        "Financial News",
	FINPLAN_4:        "Financial Planning",
	HEDGEFUND_4:      "Hedge Fund",
	INSURANCE_4:      "Insurance",
	INVEST_4:         "Investing",
	MUTUALFUND_4:     "Mutual Funds",
	QUOTES_4:         "Online Financial Tools & Quotes",
	FINANCEOPT_4:     "Options",
	RETIREPLAN_4:     "Retirement Planning",
	STOCKS_4:         "Stocks",
	TAX_4:            "Tax Planning",
	FINANCE_4:        "Finance - Other",
	USFOOD_4:         "American Cuisine",
	BBQ_4:            "Barbecues & Grilling",
	CAJUNFOOD_4:      "Cajun/Creole",
	ZHFOOD_4:         "Chinese Cuisine",
	COCKTAIL_4:       "Cocktails/Beer",
	COFFEE_4:         "Coffee/Tea",
	CUISINE_4:        "Cuisine-Specific",
	DESSERT_4:        "Desserts & Baking",
	DININGOUT_4:      "Dining Out",
	FUDALLERGY_4:     "Food Allergies",
	FRFOOD_4:         "French Cuisine",
	LOWFAT_4:         "Health/Lowfat Cooking",
	ITFOOD_4:         "Italian Cuisine",
	JAFOOD_4:         "Japanese Cuisine",
	MXFOOD_4:         "Mexican Cuisine",
	VEGAN_4:          "Vegan",
	VEGETARIAN_4:     "Vegetarian",
	WINE_4:           "Wine",
	FOOD_4:           "Food & Drink - Other",
	ADHD_4:           "A.D.D.",
	AIDS_4:           "AIDS/HIV",
	ALLERGY_4:        "Allergies",
	ALTMED_4:         "Alternative Medicine",
	ARTHRITIS_4:      "Arthritis",
	ASTHMA_4:         "Asthma",
	AUTISM_4:         "Autism/PDD",
	BIPOLAR_4:        "Bipolar Disorder",
	TUMOR_4:          "Brain Tumor",
	CANCER_4:         "Cancer",
	KIDHEALTH_4:      "Children's Health",
	CHOLESTER_4:      "Cholesterol",
	FATIGUE_4:        "Chronic Fatigue",
	PAIN_4:           "Chronic Pain",
	FLU_4:            "Cold & Flu",
	PSURGERY_4:       "Cosmetic Surgery",
	DEAFNESS_4:       "Deafness",
	DENTAL_4:         "Dental Care",
	DEPRESSION_4:     "Depression",
	DERMA_4:          "Dermatology",
	DIABETES_4:       "Diabetes",
	DISORDER_4:       "Disorders",
	EPILEPSY_4:       "Epilepsy",
	EXERCISE_4:       "Exercise",
	GERD_4:           "GERD/Acid Reflux",
	MIGRAINE_4:       "Headaches/Migraines",
	CARDIO_4:         "Heart Disease",
	HERB_4:           "Herbs for Health",
	HOLISTIC_4:       "Holistic Healing",
	IBS_4:            "IBS/Crohn’s Disease",
	ABUSE_4:          "Incest/Abuse Support",
	INCNTNENCE_4:     "Incontinence",
	INFERTILE_4:      "Infertility",
	MENHEALTH_4:      "Men’s Health",
	DIET_4:           "Nutrition & Diet",
	ORTHO_4:          "Orthopedics",
	ANXIETY_4:        "Panic/Anxiety",
	PEDIA_4:          "Pediatrics",
	RX_4:             "Pharmaceuticals",
	PHYTHERAPY_4:     "Physical Therapy",
	PSYCH_4:          "Psychology/Psychiatry",
	ADDICTION_4:      "Self-help & Addiction",
	SRHEALTH_4:       "Senior Health",
	SEXUALITY_4:      "Sexuality",
	SLEEP_4:          "Sleep Disorders",
	QUITSMOKE_4:      "Smoking Cessation",
	VITAMINS_4:       "Supplements & Compounds",
	SYNDROME_4:       "Syndrome",
	THYROID_4:        "Thyroid Disease",
	WEIGHTLOSS_4:     "Weight Loss",
	FEMHEALTH_4:      "Women’s Health",
	HEALTH_4:         "Health - Other",
	ARTTECH_4:        "Art/Technology",
	ARTCRAFT_4:       "Arts & Crafts",
	BEADWORK_4:       "Beadwork",
	BIRDWATCH_4:      "Birdwatching",
	PUZZLE_4:         "Board Games/Puzzles",
	CANDLE_4:         "Candle & Soap Making",
	CARDGAME_4:       "Card Games",
	CARTOON_4:        "Cartoons & Anime",
	CHESS_4:          "Chess",
	CIGAR_4:          "Cigars",
	COLLECT_4:        "Collecting",
	COMIC_4:          "Comic Books",
	DRAW_4:           "Drawing/Sketching",
	WRITE_4:          "Freelance Writing",
	GENEALOGY_4:      "Genealogy",
	PUBLISH_4:        "Getting Published",
	GUITAR_4:         "Guitar",
	HOMERECORD_4:     "Home Recording",
	PATENT_4:         "Investors & Patents",
	MAKEJEWEL_4:      "Jewelry Making",
	MAGIC_4:          "Magic & Illusion",
	NEEDLEWORK_4:     "Needlework",
	PAINT_4:          "Painting",
	PHOTOHOB_4:       "Photography",
	RADIO_4:          "Radio",
	RPG_4:            "Roleplaying Games",
	SCIFI_4:          "Sci-Fi & Fantasy",
	SCRAPBOOK_4:      "Scrapbooking",
	SCREENWRT_4:      "Screenwriting",
	STAMP_4:          "Stamps & Coins",
	THEME_4:          "Themes",
	VIDEOGAME_4:      "Video & Computer Games",
	WOODWORK_4:       "Woodworking",
	HOBBY_4:          "Hobbies & Interests - Other",
	APPLIANCE_4:      "Appliances",
	ENTHOME_4:        "Entertaining",
	ECOSAFETY_4:      "Environmental Safety",
	GARDEN_4:         "Gardening",
	HOMEREPAIR_4:     "Home Repair",
	HOMECINEMA_4:     "Home Theater",
	INTERIOR_4:       "Interior Decorating",
	LANDSCAPE_4:      "Landscaping",
	REMODEL_4:        "Remodeling & Construction",
	HOME_4:           "Home & Garden - Other",
	GAMES_4:          "Games",
	KIDSPAGE_4:       "Kid’s Pages",
	TOYS_4:           "Toys",
	KIDS_4:           "Kids - Other",
	DATING_4:         "Dating & Relationships",
	DIVORCE_4:        "Divorce Support",
	ETHNIC_4:         "Ethnic Specific",
	MARRIAGE_4:       "Marriage",
	PARKS_4:          "Parks, Rec Facilities & Gyms",
	SENIOR_4:         "Senior Living",
	TEENS_4:          "Teens",
	WEDDING_4:        "Weddings",
	LIFESTYLE_4:      "Lifestyle - Other",
	ADFRAUD_4:        "Ad Fraud",
	BOTS_4:           "Botnet",
	CANDC_4:          "Command and Control Centers",
	COMPR_4:          "Compromised & Links To Malware",
	MALHOME_4:        "Malware Call-Home",
	MAL_4:            "Malware Distribution Point",
	FRAUD_4:          "Phishing/Fraud",
	SPAM_4:           "Spam URLs",
	SPYWARE_4:        "Spyware & Questionable Software",
	CDN_4:            "Content Server",
	BLANK_4:          "No Content Found",
	PARKED_4:         "Parked & For Sale Domains",
	PRIVIP_4:         "Private IP Address",
	BROKEN_4:         "Unreachable",
	MISC_4:           "Miscellaneous - Other",
	IMGSEARCH_4:      "Image Search",
	INTLNEWS_4:       "International News",
	LOCALNEWS_4:      "Local News",
	MAGAZINES_4:      "Magazines",
	NATLNEWS_4:       "National News",
	PORTAL_4:         "Portal Sites",
	SEARCH_4:         "Search Engines",
	NEWS_4:           "News, Portal & Search - Other",
	PAYTOSURF_4:      "Pay To Surf",
	ADWARE_4:         "Online Ads - Other",
	AQUARIUM_4:       "Aquariums",
	BIRD_4:           "Birds",
	CAT_4:            "Cats",
	DOG_4:            "Dogs",
	BIGANIMALS_4:     "Large Animals",
	REPTILE_4:        "Reptiles",
	VETMED_4:         "Veterinary Medicine",
	PETS_4:           "Pets - Other",
	ADVOCACY_4:       "Advocacy Groups & Trade Associations",
	COMMENTARY_4:     "Commentary",
	GOVSPONSOR_4:     "Government Sponsored",
	IMMIGRANT_4:      "Immigration",
	LEGAL_4:          "Legal Issues",
	CHARITY_4:        "Philanthropic Organizations",
	POLITICS_4:       "Politics",
	CLUBS_4:          "Social & Affiliation Organizations",
	USRESOURCE_4:     "U.S. Government Resources",
	GOV_4:            "Public, Government & Law - Other",
	APARTMENT_4:      "Apartments",
	ARCHI_4:          "Architects",
	BNSHOME_4:        "Buying/Selling Homes",
	PROP_4:           "Real Estate - Other",
	ALTRELGN_4:       "Alternative Religions",
	ATHEISM_4:        "Atheism & Agnosticism",
	BUDDHISM_4:       "Buddhism",
	CATHOLIC_4:       "Catholicism",
	CHRISTIAN_4:      "Christianity",
	HINDU_4:          "Hinduism",
	ISLAM_4:          "Islam",
	JUDAISM_4:        "Judaism",
	LDS_4:            "Latter-Day Saints",
	CULT_4:           "Non-traditional Religion & Occult",
	PAGAN_4:          "Pagan/Wiccan",
	RELIGION_4:       "Religion - Other",
	ANATOMY_4:        "Anatomy",
	ASTROLOGY_4:      "Astrology & Horoscopes",
	BIOLOGY_4:        "Biology",
	BOTANY_4:         "Botany",
	CHEMISTRY_4:      "Chemistry",
	GEOGRAPHY_4:      "Geography",
	GEOLOGY_4:        "Geology",
	PARANORMAL_4:     "Paranormal Phenomena",
	PHYSICS_4:        "Physics",
	ASTRONOMY_4:      "Space/Astronomy",
	WEATHER_4:        "Weather",
	SCIENCE_4:        "Science - Other",
	P2PSHOP_4:        "Auctions & Marketplaces",
	CATALOG_4:        "Catalogs",
	SURVEY_4:         "Contests & Surveys",
	COUPON_4:         "Coupons",
	SHOPENGINE_4:     "Engines",
	ONLINESHOP_4:     "Online Shopping",
	SHOPSEARCH_4:     "Product Reviews & Price Comparisons",
	SHOP_4:           "Shopping - Other",
	AUTORACE_4:       "Auto Racing",
	BASEBALL_4:       "Baseball",
	BICYCLING_4:      "Bicycling",
	BODYBUILD_4:      "Bodybuilding",
	BOXING_4:         "Boxing",
	KAYAK_4:          "Canoeing/Kayaking",
	CHEERLEAD_4:      "Cheerleeding",
	CLIMB_4:          "Climbing",
	CRICKET_4:        "Cricket",
	FIGSKATE_4:       "Figure Skating",
	FLYFISHING_4:     "Fly Fishing",
	FOOTBALL_4:       "Football",
	FRESHFISH_4:      "Freshwater Fishing",
	GAMEFISH_4:       "Game & Fish",
	GOLF_4:           "Golf",
	HORSERACE_4:      "Horse Racing",
	HORSES_4:         "Horses",
	INSKATE_4:        "Inline Skating",
	MARTIALART_4:     "Martial Arts",
	BIKING_4:         "Mountain Biking",
	NASCAR_4:         "NASCAR Racing",
	OLYMPICS_4:       "Olympics",
	PAINTBALL_4:      "Paintball",
	MOTORACE_4:       "Power & Motorcycles",
	PROBASKET_4:      "Pro Basketball",
	ICEHOCKEY_4:      "Pro Ice Hockey",
	RODEO_4:          "Rodeo",
	RUGBY_4:          "Rugby",
	RUN_4:            "Running/Jogging",
	SAIL_4:           "Sailing",
	SALTFISH_4:       "Saltwater Fishing",
	SCUBA_4:          "Scuba Diving",
	SKATEBOARD_4:     "Skateboarding",
	SKI_4:            "Skiing",
	SNOWBOARD_4:      "Snowboarding",
	HUNT_4:           "Sport Hunting",
	SURF_4:           "Surfing/Bodyboarding",
	SWIM_4:           "Swimming",
	PINGPONG_4:       "Table Tennis/Ping-Pong",
	TENNIS_4:         "Tennis",
	VOLLEYBALL_4:     "Volleyball",
	WALK_4:           "Walking",
	WATERSKI_4:       "Waterski/Wakeboard",
	SOCCER_4:         "World Soccer",
	SPORTS_4:         "Sports - Other",
	GRAPHICS3D_4:     "3-D Graphics",
	ANIMATION_4:      "Animation",
	ANTIVIRUS_4:      "Antivirus Software",
	CLANG_4:          "C/C++",
	CAMERA_4:         "Cameras & Camcorders",
	COMPCERT_4:       "Computer Certification",
	COMPNET_4:        "Computer Networking",
	PERIPHERAL_4:     "Computer Peripherals",
	COMPREVIEW_4:     "Computer Reviews",
	DATABASE_4:       "Databases",
	DESKPUB_4:        "Desktop Publishing",
	DESKVID_4:        "Desktop Video",
	TECHENT_4:        "Entertainment",
	REPOS_4:          "File Repositories",
	GRAPHICS_4:       "Graphics Software",
	DVD_4:            "Home Video/DVD",
	COMPSEC_4:        "Information Security",
	VOIP_4:           "Internet Phone & VOIP",
	INTERNET_4:       "Internet Technology",
	JAVA_4:           "Java",
	JAVASCRIPT_4:     "Javascript",
	LINUX_4:          "Linux",
	MAC_4:            "Mac OS",
	MACSUPPORT_4:     "Mac Support",
	MOBILE_4:         "Mobile Phones",
	MP3_4:            "MP3/MIDI",
	NETCON_4:         "Net Conferencing",
	NETBEG_4:         "Net for Beginners",
	NETSEC_4:         "Network Security",
	OFFICE_4:         "Online Information Management",
	PDA_4:            "Palmtops/PDAs",
	PCSUPPORT_4:      "PC Support",
	P2P_4:            "Peer-to-Peer",
	STORAGE_4:        "Personal Storage",
	PORTABLE_4:       "Portable",
	VPN_4:            "Remote Access",
	FREEWARE_4:       "Shareware/Freeware",
	UNIX_4:           "Unix",
	UTILITIES_4:      "Utilities",
	VBASIC_4:         "Visual Basic",
	CLIPART_4:        "Web Clip Art",
	WEBDESIGN_4:      "Web Design/HTML",
	ISP_4:            "Web Hosting, ISP & Telco",
	WINDOWS_4:        "Windows",
	TECH_4:           "Technology - Other",
	ADVENTURE_4:      "Adventure Travel",
	AFRICA_4:         "Africa",
	AIRTRAVEL_4:      "Air Travel",
	AUSTRALIA_4:      "Australia & New Zealand",
	BNB_4:            "Bed & Breakfast",
	BDGTTRAVEL_4:     "Budget Travel",
	BIZTRAVEL_4:      "Business Travel",
	USTRAVEL_4:       "By US Locale",
	CAMP_4:           "Camping",
	CANADA_4:         "Canada",
	CARIBBEAN_4:      "Caribbean",
	CRUISE_4:         "Cruises",
	EASTEUROPE_4:     "Eastern Europe",
	EUROPE_4:         "Europe",
	FRANCE_4:         "France",
	GREECE_4:         "Greece",
	GETAWAY_4:        "Honeymoons/Getaways",
	HOTEL_4:          "Hotels",
	ITALY_4:          "Italy",
	JAPAN_4:          "Japan",
	MEXICO_4:         "Mexico & Central America",
	NATLPARKS_4:      "National Parks",
	NAVIGATION_4:     "Navigation",
	SAMERICA_4:       "South America",
	SPA_4:            "Spas",
	THEMEPARK_4:      "Theme Parks",
	KIDTRAVEL_4:      "Traveling with Kids",
	UKTRAVEL_4:       "United Kingdom",
	TRAVEL_4:         "Travel - Other",
	TERRORISM_4:      "Terrorism",
	CRYPTOCUR_4:      "Cryptocurrency",
	CRYPTOMINE_4:     "Cryptocurrency Mining",
	BLOCKCHAIN_4:     "Blockchain",
	FAKENEWS_4:       "Fake News",
	API_4:            "APIs",
	IOT_4:            "Internet of Things",
	AI_ML_4:          "A.I. & M.L.",
}

// Long returns the category long name
func (c Category) Long() string {
	return CategoryLong[c]
}

// ParseCategory parses a category id or case insensitive short name and returns
// a Category
func ParseCategory(name string) Category {
	if cid, err := strconv.Atoi(name); err == nil {
		if _, ok := Category_name[int32(cid)]; ok {
			return Category(cid)
		}
	}

	for k, v := range Category_value {
		if strings.EqualFold(name, k) {
			return Category(v)
		}

		if strings.EqualFold(name+"_4", k) {
			return Category(v)
		}
	}

	return UNKNOWN_CATEGORY
}
