package msgpb

import (
	"strings"

	"github.com/pkg/errors"
)

func Categories(group Group) []Category {
	cats, ok := groups[group]
	if !ok {
		return nil
	}

	ret := make([]Category, len(cats))
	copy(ret, cats)
	return ret
}

var groups = map[Group][]Category{
	GROUP_ADULT: {
		ABORTION_4,
		PROCHOICE_4,
		PROLIFE_4,
		PG13_4,
		LINGERIE_4,
		NUDITY_4,
		PORN_4,
		PROFANITY_4,
		RATEDR_4,
		SEX_4,
		EDSEX_4,
		GAMBLING_4,
		GAY_4,
		TOBACCO_4,
	},
	GROUP_AGGRESSIVE: {
		AGGRESSIVE_4,
		MIL_4,
		VIOLENCE_4,
		WEAPONS_4,
		TERRORISM_4,
	},
	GROUP_ARTS: {
		ART_4,
		FINEART_4,
	},
	GROUP_AUTOMOTIVE: {
		AUTOMOTIVE_4,
		AUTOPARTS_4,
		AUTOREPAIR_4,
		BUYSELLCAR_4,
		CARCULTURE_4,
		CPO_4,
		CARCONVERT_4,
		COUPE_4,
		CUV_4,
		DIESEL_4,
		EVEHICLE_4,
		HATCHBACK_4,
		HYBRIDCAR_4,
		LUXURYCAR_4,
		MINIVAN_4,
		MOTORCYCLE_4,
		OFFROAD_4,
		PERFCAR_4,
		PICKUP_4,
		ROADASSIST_4,
		SEDAN_4,
		TRUCK_4,
		VINTAGECAR_4,
		WAGON_4,
	},
	GROUP_BUSINESS: {
		BUSINESS_4,
		AG_4,
		BIOTECH_4,
		BIZSOFT_4,
		BUILD_4,
		FORESTRY_4,
		BIZGOV_4,
		NATURE_4,
		FURNITURE_4,
		HR_4,
		MFG_4,
		MKTG_4,
		METAL_4,
		PHYSEC_4,
		BIZPROD_4,
		RETIREHOME_4,
		SHIPPING_4,
	},
	GROUP_CAREERS: {
		CAREER_4,
		JOBADVICE_4,
		CAREERPLAN_4,
		COLJOB_4,
		FINAID_4,
		JOBFAIR_4,
		JOBSEARCH_4,
		NURSING_4,
		RESUME_4,
		SCHOLAR_4,
		TELCOMMUTE_4,
		MILJOBS_4,
	},
	GROUP_CRIMINAL_ACTIVITIES: {
		CRIM_4,
		CHILDABUSE_4,
		CRIMSKILL_4,
		HACKING_4,
		HATE_4,
		DRUGS_4,
		POT_4,
		WAREZ_4,
		EDUCHEAT_4,
		TORRENT_4,
		SELFHARM_4,
		TERRORISM_4,
		CRYPTOMINE_4,
	},
	GROUP_DYNAMIC: {
		ANON_4,
		CHAT_4,
		GROUPS_4,
		IM_4,
		LOGIN_4,
		BLOG_4,
		PHOTO_4,
		PROFNET_4,
		REDIR_4,
		SOCNET_4,
		SMS_4,
		TRANSLATOR_4,
		EMAIL_4,
		CARDS_4,
		API_4,
	},
	GROUP_EDUCATION: {
		EDU_4,
		EDU712_4,
		ADULTEDU_4,
		ARTHISTORY_4,
		COLADMIN_4,
		COLLIFE_4,
		DLEARNING_4,
		SCHOOL_4,
		EDSTUDIES_4,
		ENGLISH_4,
		GRADSCHOOL_4,
		HOMESCHOOL_4,
		HOMEWORK_4,
		K6EDU_4,
		LANGUAGE_4,
		BOOK_4,
		PRIVSCHOOL_4,
		REF_4,
		SPED_4,
		BIZEDU_4,
		TUTOR_4,
		WIKI_4,
	},
	GROUP_ENTERTAINMENT: {
		ENT_4,
		ENTNEWS_4,
		VENUE_4,
		HUMOR_4,
		MOVIE_4,
		MUSIC_4,
		AUDIO_4,
		VIDEO_4,
		TV_4,
	},
	GROUP_FAMILY_AND_PARENTING: {
		FAMILY_4,
		ADOPT_4,
		TODDLER_4,
		PRESCHOOL_4,
		ELDERCARE_4,
		FAMILYNET_4,
		K6PARENT_4,
		PARENTEEN_4,
		PREGNANCY_4,
		SNKPARENT_4,
	},
	GROUP_FASHION: {
		FASHION_4,
		FASHNMISC_4,
		ACCESSORY_4,
		BEAUTY_4,
		BODYART_4,
		CLOTHING_4,
		JEWELRY_4,
		SWIMSUIT_4,
	},
	GROUP_FINANCE: {
		FINANCE_4,
		ACCTG_4,
		BANK_4,
		BGNINVEST_4,
		LOAN_4,
		FINNEWS_4,
		FINPLAN_4,
		HEDGEFUND_4,
		INSURANCE_4,
		INVEST_4,
		MUTUALFUND_4,
		QUOTES_4,
		FINANCEOPT_4,
		RETIREPLAN_4,
		STOCKS_4,
		TAX_4,
		CRYPTOCUR_4,
	},
	GROUP_FOOD_AND_DRINK: {
		FOOD_4,
		USFOOD_4,
		BBQ_4,
		CAJUNFOOD_4,
		ZHFOOD_4,
		COCKTAIL_4,
		COFFEE_4,
		CUISINE_4,
		DESSERT_4,
		DININGOUT_4,
		FUDALLERGY_4,
		FRFOOD_4,
		LOWFAT_4,
		ITFOOD_4,
		JAFOOD_4,
		MXFOOD_4,
		VEGAN_4,
		VEGETARIAN_4,
		WINE_4,
	},
	GROUP_HEALTH: {
		HEALTH_4,
		ADHD_4,
		AIDS_4,
		ALLERGY_4,
		ALTMED_4,
		ARTHRITIS_4,
		ASTHMA_4,
		AUTISM_4,
		BIPOLAR_4,
		TUMOR_4,
		CANCER_4,
		KIDHEALTH_4,
		CHOLESTER_4,
		FATIGUE_4,
		PAIN_4,
		FLU_4,
		PSURGERY_4,
		DEAFNESS_4,
		DENTAL_4,
		DEPRESSION_4,
		DERMA_4,
		DIABETES_4,
		DISORDER_4,
		EPILEPSY_4,
		EXERCISE_4,
		GERD_4,
		MIGRAINE_4,
		CARDIO_4,
		HERB_4,
		HOLISTIC_4,
		IBS_4,
		ABUSE_4,
		INCNTNENCE_4,
		INFERTILE_4,
		MENHEALTH_4,
		DIET_4,
		ORTHO_4,
		ANXIETY_4,
		PEDIA_4,
		RX_4,
		PHYTHERAPY_4,
		PSYCH_4,
		ADDICTION_4,
		SRHEALTH_4,
		SEXUALITY_4,
		SLEEP_4,
		QUITSMOKE_4,
		VITAMINS_4,
		SYNDROME_4,
		THYROID_4,
		WEIGHTLOSS_4,
		FEMHEALTH_4,
	},
	GROUP_HOBBIES_AND_INTERESTS: {
		HOBBY_4,
		ARTTECH_4,
		ARTCRAFT_4,
		BEADWORK_4,
		BIRDWATCH_4,
		PUZZLE_4,
		CANDLE_4,
		CARDGAME_4,
		CARTOON_4,
		CHESS_4,
		CIGAR_4,
		COLLECT_4,
		COMIC_4,
		DRAW_4,
		WRITE_4,
		GENEALOGY_4,
		PUBLISH_4,
		GUITAR_4,
		HOMERECORD_4,
		PATENT_4,
		MAKEJEWEL_4,
		MAGIC_4,
		NEEDLEWORK_4,
		PAINT_4,
		PHOTOHOB_4,
		RADIO_4,
		RPG_4,
		SCIFI_4,
		SCRAPBOOK_4,
		SCREENWRT_4,
		STAMP_4,
		THEME_4,
		VIDEOGAME_4,
		WOODWORK_4,
	},
	GROUP_HOME_AND_GARDEN: {
		HOME_4,
		APPLIANCE_4,
		ENTHOME_4,
		ECOSAFETY_4,
		GARDEN_4,
		HOMEREPAIR_4,
		HOMECINEMA_4,
		INTERIOR_4,
		LANDSCAPE_4,
		REMODEL_4,
	},
	GROUP_KIDS: {
		KIDS_4,
		GAMES_4,
		KIDSPAGE_4,
		TOYS_4,
	},
	GROUP_LIFESTYLE: {
		LIFESTYLE_4,
		DATING_4,
		DIVORCE_4,
		ETHNIC_4,
		MARRIAGE_4,
		PARKS_4,
		SENIOR_4,
		TEENS_4,
		WEDDING_4,
	},
	GROUP_MALICIOUS: {
		ADFRAUD_4,
		BOTS_4,
		CANDC_4,
		COMPR_4,
		MALHOME_4,
		MAL_4,
		FRAUD_4,
		SPAM_4,
		SPYWARE_4,
		CRYPTOMINE_4,
		FAKENEWS_4,
	},
	GROUP_MISCELLANEOUS: {
		MISC_4,
		CDN_4,
		BLANK_4,
		PARKED_4,
		PRIVIP_4,
		BROKEN_4,
	},
	GROUP_NEWS_PORTAL_AND_SEARCH: {
		NEWS_4,
		INTLNEWS_4,
		LOCALNEWS_4,
		MAGAZINES_4,
		NATLNEWS_4,
		SEARCH_4,
		IMGSEARCH_4,
		PORTAL_4,
	},
	GROUP_ONLINE_ADS: {
		ADWARE_4,
		PAYTOSURF_4,
	},
	GROUP_PETS: {
		PETS_4,
		AQUARIUM_4,
		BIRD_4,
		CAT_4,
		DOG_4,
		BIGANIMALS_4,
		REPTILE_4,
		VETMED_4,
	},
	GROUP_PUBLIC_GOVERNMENT_AND_LAW: {
		GOV_4,
		GOVSPONSOR_4,
		IMMIGRANT_4,
		USRESOURCE_4,
		POLITICS_4,
		COMMENTARY_4,
		LEGAL_4,
		CLUBS_4,
		CHARITY_4,
		ADVOCACY_4,
	},
	GROUP_REAL_ESTATE: {
		PROP_4,
		APARTMENT_4,
		ARCHI_4,
		BNSHOME_4,
	},
	GROUP_RELIGION: {
		RELIGION_4,
		ALTRELGN_4,
		ATHEISM_4,
		BUDDHISM_4,
		CATHOLIC_4,
		CHRISTIAN_4,
		HINDU_4,
		ISLAM_4,
		JUDAISM_4,
		LDS_4,
		CULT_4,
		PAGAN_4,
	},
	GROUP_SCIENCE: {
		SCIENCE_4,
		ANATOMY_4,
		ASTROLOGY_4,
		BIOLOGY_4,
		BOTANY_4,
		CHEMISTRY_4,
		GEOGRAPHY_4,
		GEOLOGY_4,
		PARANORMAL_4,
		PHYSICS_4,
		ASTRONOMY_4,
		WEATHER_4,
	},
	GROUP_SHOPPING: {
		SHOP_4,
		P2PSHOP_4,
		CATALOG_4,
		SURVEY_4,
		COUPON_4,
		SHOPENGINE_4,
		ONLINESHOP_4,
		SHOPSEARCH_4,
	},
	GROUP_SPORTS: {
		SPORTS_4,
		AUTORACE_4,
		BASEBALL_4,
		BICYCLING_4,
		BODYBUILD_4,
		BOXING_4,
		KAYAK_4,
		CHEERLEAD_4,
		CLIMB_4,
		CRICKET_4,
		FIGSKATE_4,
		FLYFISHING_4,
		FOOTBALL_4,
		FRESHFISH_4,
		GAMEFISH_4,
		GOLF_4,
		HORSERACE_4,
		HORSES_4,
		INSKATE_4,
		MARTIALART_4,
		BIKING_4,
		NASCAR_4,
		OLYMPICS_4,
		PAINTBALL_4,
		MOTORACE_4,
		PROBASKET_4,
		ICEHOCKEY_4,
		RODEO_4,
		RUGBY_4,
		RUN_4,
		SAIL_4,
		SALTFISH_4,
		SCUBA_4,
		SKATEBOARD_4,
		SKI_4,
		SNOWBOARD_4,
		HUNT_4,
		SURF_4,
		SWIM_4,
		PINGPONG_4,
		TENNIS_4,
		VOLLEYBALL_4,
		WALK_4,
		WATERSKI_4,
		SOCCER_4,
	},
	GROUP_TECHNOLOGY: {
		TECH_4,
		GRAPHICS3D_4,
		ANIMATION_4,
		ANTIVIRUS_4,
		CLANG_4,
		CAMERA_4,
		COMPCERT_4,
		COMPNET_4,
		PERIPHERAL_4,
		COMPREVIEW_4,
		DATABASE_4,
		DESKPUB_4,
		DESKVID_4,
		REPOS_4,
		GRAPHICS_4,
		DVD_4,
		COMPSEC_4,
		VOIP_4,
		INTERNET_4,
		JAVA_4,
		JAVASCRIPT_4,
		LINUX_4,
		MAC_4,
		MACSUPPORT_4,
		MOBILE_4,
		MP3_4,
		NETCON_4,
		NETBEG_4,
		NETSEC_4,
		PDA_4,
		PCSUPPORT_4,
		P2P_4,
		STORAGE_4,
		TECHENT_4,
		PORTABLE_4,
		VPN_4,
		FREEWARE_4,
		UNIX_4,
		UTILITIES_4,
		VBASIC_4,
		CLIPART_4,
		WEBDESIGN_4,
		ISP_4,
		WINDOWS_4,
		OFFICE_4,
		BLOCKCHAIN_4,
		IOT_4,
		AI_ML_4,
	},
	GROUP_TRAVEL: {
		TRAVEL_4,
		ADVENTURE_4,
		AFRICA_4,
		AIRTRAVEL_4,
		AUSTRALIA_4,
		BNB_4,
		BDGTTRAVEL_4,
		BIZTRAVEL_4,
		USTRAVEL_4,
		CAMP_4,
		CANADA_4,
		CARIBBEAN_4,
		CRUISE_4,
		EASTEUROPE_4,
		EUROPE_4,
		FRANCE_4,
		GREECE_4,
		GETAWAY_4,
		HOTEL_4,
		ITALY_4,
		JAPAN_4,
		MEXICO_4,
		NATLPARKS_4,
		NAVIGATION_4,
		SAMERICA_4,
		SPA_4,
		THEMEPARK_4,
		KIDTRAVEL_4,
		UKTRAVEL_4,
	},
}

func GroupHasFallback(group Group) bool {
	_, ok := groupHasFallback[group]
	return ok
}

var groupHasFallback = map[Group]struct{}{
	GROUP_AGGRESSIVE:                {},
	GROUP_ARTS:                      {},
	GROUP_AUTOMOTIVE:                {},
	GROUP_BUSINESS:                  {},
	GROUP_CAREERS:                   {},
	GROUP_CRIMINAL_ACTIVITIES:       {},
	GROUP_EDUCATION:                 {},
	GROUP_FAMILY_AND_PARENTING:      {},
	GROUP_FASHION:                   {},
	GROUP_FINANCE:                   {},
	GROUP_FOOD_AND_DRINK:            {},
	GROUP_HEALTH:                    {},
	GROUP_HOBBIES_AND_INTERESTS:     {},
	GROUP_HOME_AND_GARDEN:           {},
	GROUP_KIDS:                      {},
	GROUP_LIFESTYLE:                 {},
	GROUP_NEWS_PORTAL_AND_SEARCH:    {},
	GROUP_ONLINE_ADS:                {},
	GROUP_PETS:                      {},
	GROUP_PUBLIC_GOVERNMENT_AND_LAW: {},
	GROUP_REAL_ESTATE:               {},
	GROUP_SCIENCE:                   {},
	GROUP_SHOPPING:                  {},
	GROUP_SPORTS:                    {},
	GROUP_TECHNOLOGY:                {},
	GROUP_TRAVEL:                    {},
}

func GroupHasRollup(group Group) bool {
	_, ok := groupHasRollup[group]
	return ok
}

var groupHasRollup = map[Group]struct{}{
	GROUP_AUTOMOTIVE:     {},
	GROUP_EDUCATION:      {},
	GROUP_ENTERTAINMENT:  {},
	GROUP_FASHION:        {},
	GROUP_FINANCE:        {},
	GROUP_FOOD_AND_DRINK: {},
	GROUP_HEALTH:         {},
	GROUP_LIFESTYLE:      {},
	GROUP_ONLINE_ADS:     {},
	GROUP_PETS:           {},
	GROUP_REAL_ESTATE:    {},
	GROUP_RELIGION:       {},
	GROUP_SCIENCE:        {},
	GROUP_SPORTS:         {},
	GROUP_TECHNOLOGY:     {},
	GROUP_TRAVEL:         {},
}

func GroupHasOther(group Group) bool {
	_, ok := groupHasOther[group]
	return ok
}

var groupHasOther = map[Group]struct{}{
	GROUP_ARTS:                   {},
	GROUP_AUTOMOTIVE:             {},
	GROUP_BUSINESS:               {},
	GROUP_CAREERS:                {},
	GROUP_ENTERTAINMENT:          {},
	GROUP_FAMILY_AND_PARENTING:   {},
	GROUP_FINANCE:                {},
	GROUP_FOOD_AND_DRINK:         {},
	GROUP_HEALTH:                 {},
	GROUP_HOBBIES_AND_INTERESTS:  {},
	GROUP_NEWS_PORTAL_AND_SEARCH: {},
	GROUP_ONLINE_ADS:             {},
	GROUP_PETS:                   {},
	GROUP_REAL_ESTATE:            {},
	GROUP_RELIGION:               {},
	GROUP_SCIENCE:                {},
	GROUP_SPORTS:                 {},
	GROUP_TECHNOLOGY:             {},
	GROUP_TRAVEL:                 {},
}

func ParseGroup(name string) (Group, error) {
	if gid, ok := Group_value[strings.ToUpper(name)]; ok {
		return Group(gid), nil
	}

	return GROUP_UNKNOWN, errors.Errorf("invalid group: %s", name)
}
