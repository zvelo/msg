package msg

// Long returns the category long name
func (c Category) Long() string {
	switch c {
	case Category_UNKNOWN:
		return "Unknown Category"
	case Category_ABORTION_4:
		return "Abortion"
	case Category_PROCHOICE_4:
		return "Abortion Pro Choice"
	case Category_PROLIFE_4:
		return "Abortion Pro Life"
	case Category_PG13_4:
		return "Child Inappropriate"
	case Category_GAMBLING_4:
		return "Gambling"
	case Category_GAY_4:
		return "Gay, Lesbian or Bisexual"
	case Category_LINGERIE_4:
		return "Lingerie, Suggestive & Pinup"
	case Category_NUDITY_4:
		return "Nudity"
	case Category_PORN_4:
		return "Pornography"
	case Category_PROFANITY_4:
		return "Profanity"
	case Category_RATEDR_4:
		return "R-Rated"
	case Category_SEX_4:
		return "Sex & Erotic"
	case Category_EDSEX_4:
		return "Sex Education"
	case Category_TOBACCO_4:
		return "Tobacco"
	case Category_MIL_4:
		return "Military"
	case Category_VIOLENCE_4:
		return "Violence"
	case Category_WEAPONS_4:
		return "Weapons"
	case Category_AGGRESSIVE_4:
		return "Aggressive - Other"
	case Category_FINEART_4:
		return "Fine Art"
	case Category_ART_4:
		return "Arts - Other"
	case Category_AUTOPARTS_4:
		return "Auto Parts"
	case Category_AUTOREPAIR_4:
		return "Auto Repair"
	case Category_BUYSELLCAR_4:
		return "Buying/Selling Cars"
	case Category_CARCULTURE_4:
		return "Car Culture"
	case Category_CPO_4:
		return "Certified Pre-Owned"
	case Category_CARCONVERT_4:
		return "Convertible"
	case Category_COUPE_4:
		return "Coupe"
	case Category_CUV_4:
		return "Crossover"
	case Category_DIESEL_4:
		return "Diesel"
	case Category_EVEHICLE_4:
		return "Electric Vehicle"
	case Category_HATCHBACK_4:
		return "Hatchback"
	case Category_HYBRIDCAR_4:
		return "Hybrid"
	case Category_LUXURYCAR_4:
		return "Luxury"
	case Category_MINIVAN_4:
		return "MiniVan"
	case Category_MOTORCYCLE_4:
		return "Motorcycles"
	case Category_OFFROAD_4:
		return "Off-Road Vehicles"
	case Category_PERFCAR_4:
		return "Performance Vehicles"
	case Category_PICKUP_4:
		return "Pickup"
	case Category_ROADASSIST_4:
		return "Road-Side Assistance"
	case Category_SEDAN_4:
		return "Sedan"
	case Category_TRUCK_4:
		return "Trucks & Accessories"
	case Category_VINTAGECAR_4:
		return "Vintage Cars"
	case Category_WAGON_4:
		return "Wagon"
	case Category_AUTOMOTIVE_4:
		return "Automotive - Other"
	case Category_AG_4:
		return "Agriculture"
	case Category_BIOTECH_4:
		return "Biotechnology"
	case Category_BIZSOFT_4:
		return "Business Software"
	case Category_BUILD_4:
		return "Construction"
	case Category_FORESTRY_4:
		return "Forestry"
	case Category_BIZGOV_4:
		return "Government"
	case Category_NATURE_4:
		return "Green Solutions & Conservation"
	case Category_FURNITURE_4:
		return "Home & Office Furnishings"
	case Category_HR_4:
		return "Human Resources"
	case Category_MFG_4:
		return "Manufacturing"
	case Category_MKTG_4:
		return "Marketing Services"
	case Category_METAL_4:
		return "Metals"
	case Category_PHYSEC_4:
		return "Physical Security"
	case Category_BIZPROD_4:
		return "Productivity"
	case Category_RETIREHOME_4:
		return "Retirement Homes & Assisted Living"
	case Category_SHIPPING_4:
		return "Shipping & Logistics"
	case Category_BUSINESS_4:
		return "Business - Other"
	case Category_JOBADVICE_4:
		return "Career Advice"
	case Category_CAREERPLAN_4:
		return "Career Planning"
	case Category_COLJOB_4:
		return "College"
	case Category_FINAID_4:
		return "Financial Aid"
	case Category_JOBFAIR_4:
		return "Job Fairs"
	case Category_JOBSEARCH_4:
		return "Job Search"
	case Category_NURSING_4:
		return "Nursing"
	case Category_RESUME_4:
		return "Resume Writing/Advice"
	case Category_SCHOLAR_4:
		return "Scholarships"
	case Category_TELCOMMUTE_4:
		return "Telecommuting"
	case Category_MILJOBS_4:
		return "U.S. Military"
	case Category_CAREER_4:
		return "Careers - Other"
	case Category_CHILDABUSE_4:
		return "Child Abuse Images"
	case Category_CRIMSKILL_4:
		return "Criminal Skills"
	case Category_HACKING_4:
		return "Hacking"
	case Category_HATE_4:
		return "Hate Speech"
	case Category_DRUGS_4:
		return "Illegal Drugs"
	case Category_POT_4:
		return "Marijuana"
	case Category_WAREZ_4:
		return "Piracy & Copyright Theft"
	case Category_EDUCHEAT_4:
		return "School Cheating"
	case Category_SELFHARM_4:
		return "Self Harm"
	case Category_TORRENT_4:
		return "Torrent Repository"
	case Category_CRIM_4:
		return "Criminal Activities - Other"
	case Category_ANON_4:
		return "Anonymizer"
	case Category_CHAT_4:
		return "Chat"
	case Category_GROUPS_4:
		return "Community Forums"
	case Category_IM_4:
		return "Instant Messenger"
	case Category_LOGIN_4:
		return "Login Screens"
	case Category_BLOG_4:
		return "Personal Pages & Blogs"
	case Category_PHOTO_4:
		return "Photo Sharing"
	case Category_PROFNET_4:
		return "Professional Networking"
	case Category_REDIR_4:
		return "Redirect"
	case Category_SOCNET_4:
		return "Social Networking"
	case Category_SMS_4:
		return "Text Messaging & SMS"
	case Category_TRANSLATOR_4:
		return "Translator"
	case Category_EMAIL_4:
		return "Web-based Email"
	case Category_CARDS_4:
		return "Web-based Greeting Cards"
	case Category_EDU712_4:
		return "7-12 Education"
	case Category_ADULTEDU_4:
		return "Adult Education"
	case Category_ARTHISTORY_4:
		return "Art History"
	case Category_COLADMIN_4:
		return "College Administration"
	case Category_COLLIFE_4:
		return "College Life"
	case Category_DLEARNING_4:
		return "Distance Learning"
	case Category_SCHOOL_4:
		return "Educational Institutions"
	case Category_EDSTUDIES_4:
		return "Educational Materials & Studies"
	case Category_ENGLISH_4:
		return "English as a 2nd Language"
	case Category_GRADSCHOOL_4:
		return "Graduate School"
	case Category_HOMESCHOOL_4:
		return "Homeschooling"
	case Category_HOMEWORK_4:
		return "Homework/Study Tips"
	case Category_K6EDU_4:
		return "K-6 Educators"
	case Category_LANGUAGE_4:
		return "Language Learning"
	case Category_BOOK_4:
		return "Literature & Books"
	case Category_PRIVSCHOOL_4:
		return "Private School"
	case Category_REF_4:
		return "Reference Materials & Maps"
	case Category_SPED_4:
		return "Special Education"
	case Category_BIZEDU_4:
		return "Studying Business"
	case Category_TUTOR_4:
		return "Tutoring"
	case Category_WIKI_4:
		return "Wikis"
	case Category_EDU_4:
		return "Education - Other"
	case Category_ENTNEWS_4:
		return "Entertainment News & Celebrity Sites"
	case Category_VENUE_4:
		return "Entertainment Venues & Events"
	case Category_HUMOR_4:
		return "Humor"
	case Category_MOVIE_4:
		return "Movies"
	case Category_MUSIC_4:
		return "Music"
	case Category_AUDIO_4:
		return "Streaming & Downloadable Audio"
	case Category_VIDEO_4:
		return "Streaming & Downloadable Video"
	case Category_TV_4:
		return "Television"
	case Category_ENT_4:
		return "Entertainment - Other"
	case Category_ADOPT_4:
		return "Adoption"
	case Category_TODDLER_4:
		return "Babies and Toddlers"
	case Category_PRESCHOOL_4:
		return "Daycare/Pre School"
	case Category_ELDERCARE_4:
		return "Eldercare"
	case Category_FAMILYNET_4:
		return "Family Internet"
	case Category_K6PARENT_4:
		return "Parenting"
	case Category_PARENTEEN_4:
		return "Parenting Teens"
	case Category_PREGNANCY_4:
		return "Pregnancy"
	case Category_SNKPARENT_4:
		return "Special Needs Kids"
	case Category_FAMILY_4:
		return "Family & Parenting - Other"
	case Category_ACCESSORY_4:
		return "Accessories"
	case Category_BEAUTY_4:
		return "Beauty"
	case Category_BODYART_4:
		return "Body Art"
	case Category_CLOTHING_4:
		return "Clothing"
	case Category_FASHNMISC_4:
		return "Fashion"
	case Category_JEWELRY_4:
		return "Jewelry"
	case Category_SWIMSUIT_4:
		return "Swimsuits"
	case Category_FASHION_4:
		return "Fashion - Other"
	case Category_ACCTG_4:
		return "Accounting"
	case Category_BANK_4:
		return "Banking"
	case Category_BGNINVEST_4:
		return "Beginning Investing"
	case Category_LOAN_4:
		return "Credit/Debt & Loans"
	case Category_FINNEWS_4:
		return "Financial News"
	case Category_FINPLAN_4:
		return "Financial Planning"
	case Category_HEDGEFUND_4:
		return "Hedge Fund"
	case Category_INSURANCE_4:
		return "Insurance"
	case Category_INVEST_4:
		return "Investing"
	case Category_MUTUALFUND_4:
		return "Mutual Funds"
	case Category_QUOTES_4:
		return "Online Financial Tools & Quotes"
	case Category_FINANCEOPT_4:
		return "Options"
	case Category_RETIREPLAN_4:
		return "Retirement Planning"
	case Category_STOCKS_4:
		return "Stocks"
	case Category_TAX_4:
		return "Tax Planning"
	case Category_FINANCE_4:
		return "Finance - Other"
	case Category_USFOOD_4:
		return "American Cuisine"
	case Category_BBQ_4:
		return "Barbecues & Grilling"
	case Category_CAJUNFOOD_4:
		return "Cajun/Creole"
	case Category_ZHFOOD_4:
		return "Chinese Cuisine"
	case Category_COCKTAIL_4:
		return "Cocktails/Beer"
	case Category_COFFEE_4:
		return "Coffee/Tea"
	case Category_CUISINE_4:
		return "Cuisine-Specific"
	case Category_DESSERT_4:
		return "Desserts & Baking"
	case Category_DININGOUT_4:
		return "Dining Out"
	case Category_FUDALLERGY_4:
		return "Food Allergies"
	case Category_FRFOOD_4:
		return "French Cuisine"
	case Category_LOWFAT_4:
		return "Health/Lowfat Cooking"
	case Category_ITFOOD_4:
		return "Italian Cuisine"
	case Category_JAFOOD_4:
		return "Japanese Cuisine"
	case Category_MXFOOD_4:
		return "Mexican Cuisine"
	case Category_VEGAN_4:
		return "Vegan"
	case Category_VEGETARIAN_4:
		return "Vegetarian"
	case Category_WINE_4:
		return "Wine"
	case Category_FOOD_4:
		return "Food & Drink - Other"
	case Category_ADHD_4:
		return "A.D.D."
	case Category_AIDS_4:
		return "AIDS/HIV"
	case Category_ALLERGY_4:
		return "Allergies"
	case Category_ALTMED_4:
		return "Alternative Medicine"
	case Category_ARTHRITIS_4:
		return "Arthritis"
	case Category_ASTHMA_4:
		return "Asthma"
	case Category_AUTISM_4:
		return "Autism/PDD"
	case Category_BIPOLAR_4:
		return "Bipolar Disorder"
	case Category_TUMOR_4:
		return "Brain Tumor"
	case Category_CANCER_4:
		return "Cancer"
	case Category_KIDHEALTH_4:
		return "Children's Health"
	case Category_CHOLESTER_4:
		return "Cholesterol"
	case Category_FATIGUE_4:
		return "Chronic Fatigue"
	case Category_PAIN_4:
		return "Chronic Pain"
	case Category_FLU_4:
		return "Cold & Flu"
	case Category_PSURGERY_4:
		return "Cosmetic Surgery"
	case Category_DEAFNESS_4:
		return "Deafness"
	case Category_DENTAL_4:
		return "Dental Care"
	case Category_DEPRESSION_4:
		return "Depression"
	case Category_DERMA_4:
		return "Dermatology"
	case Category_DIABETES_4:
		return "Diabetes"
	case Category_DISORDER_4:
		return "Disorders"
	case Category_EPILEPSY_4:
		return "Epilepsy"
	case Category_EXERCISE_4:
		return "Exercise"
	case Category_GERD_4:
		return "GERD/Acid Reflux"
	case Category_MIGRAINE_4:
		return "Headaches/Migraines"
	case Category_CARDIO_4:
		return "Heart Disease"
	case Category_HERB_4:
		return "Herbs for Health"
	case Category_HOLISTIC_4:
		return "Holistic Healing"
	case Category_IBS_4:
		return "IBS/Crohn’s Disease"
	case Category_ABUSE_4:
		return "Incest/Abuse Support"
	case Category_INCNTNENCE_4:
		return "Incontinence"
	case Category_INFERTILE_4:
		return "Infertility"
	case Category_MENHEALTH_4:
		return "Men’s Health"
	case Category_DIET_4:
		return "Nutrition & Diet"
	case Category_ORTHO_4:
		return "Orthopedics"
	case Category_ANXIETY_4:
		return "Panic/Anxiety"
	case Category_PEDIA_4:
		return "Pediatrics"
	case Category_RX_4:
		return "Pharmaceuticals"
	case Category_PHYTHERAPY_4:
		return "Physical Therapy"
	case Category_PSYCH_4:
		return "Psychology/Psychiatry"
	case Category_ADDICTION_4:
		return "Self-help & Addiction"
	case Category_SRHEALTH_4:
		return "Senior Health"
	case Category_SEXUALITY_4:
		return "Sexuality"
	case Category_SLEEP_4:
		return "Sleep Disorders"
	case Category_QUITSMOKE_4:
		return "Smoking Cessation"
	case Category_VITAMINS_4:
		return "Supplements & Compounds"
	case Category_SYNDROME_4:
		return "Syndrome"
	case Category_THYROID_4:
		return "Thyroid Disease"
	case Category_WEIGHTLOSS_4:
		return "Weight Loss"
	case Category_FEMHEALTH_4:
		return "Women’s Health"
	case Category_HEALTH_4:
		return "Health - Other"
	case Category_ARTTECH_4:
		return "Art/Technology"
	case Category_ARTCRAFT_4:
		return "Arts & Crafts"
	case Category_BEADWORK_4:
		return "Beadwork"
	case Category_BIRDWATCH_4:
		return "Birdwatching"
	case Category_PUZZLE_4:
		return "Board Games/Puzzles"
	case Category_CANDLE_4:
		return "Candle & Soap Making"
	case Category_CARDGAME_4:
		return "Card Games"
	case Category_CARTOON_4:
		return "Cartoons & Anime"
	case Category_CHESS_4:
		return "Chess"
	case Category_CIGAR_4:
		return "Cigars"
	case Category_COLLECT_4:
		return "Collecting"
	case Category_COMIC_4:
		return "Comic Books"
	case Category_DRAW_4:
		return "Drawing/Sketching"
	case Category_WRITE_4:
		return "Freelance Writing"
	case Category_GENEALOGY_4:
		return "Genealogy"
	case Category_PUBLISH_4:
		return "Getting Published"
	case Category_GUITAR_4:
		return "Guitar"
	case Category_HOMERECORD_4:
		return "Home Recording"
	case Category_PATENT_4:
		return "Investors & Patents"
	case Category_MAKEJEWEL_4:
		return "Jewelry Making"
	case Category_MAGIC_4:
		return "Magic & Illusion"
	case Category_NEEDLEWORK_4:
		return "Needlework"
	case Category_PAINT_4:
		return "Painting"
	case Category_PHOTOHOB_4:
		return "Photography"
	case Category_RADIO_4:
		return "Radio"
	case Category_RPG_4:
		return "Roleplaying Games"
	case Category_SCIFI_4:
		return "Sci-Fi & Fantasy"
	case Category_SCRAPBOOK_4:
		return "Scrapbooking"
	case Category_SCREENWRT_4:
		return "Screenwriting"
	case Category_STAMP_4:
		return "Stamps & Coins"
	case Category_THEME_4:
		return "Themes"
	case Category_VIDEOGAME_4:
		return "Video & Computer Games"
	case Category_WOODWORK_4:
		return "Woodworking"
	case Category_HOBBY_4:
		return "Hobbies & Interests - Other"
	case Category_APPLIANCE_4:
		return "Appliances"
	case Category_ENTHOME_4:
		return "Entertaining"
	case Category_ECOSAFETY_4:
		return "Environmental Safety"
	case Category_GARDEN_4:
		return "Gardening"
	case Category_HOMEREPAIR_4:
		return "Home Repair"
	case Category_HOMECINEMA_4:
		return "Home Theater"
	case Category_INTERIOR_4:
		return "Interior Decorating"
	case Category_LANDSCAPE_4:
		return "Landscaping"
	case Category_REMODEL_4:
		return "Remodeling & Construction"
	case Category_HOME_4:
		return "Home & Garden - Other"
	case Category_GAMES_4:
		return "Games"
	case Category_KIDSPAGE_4:
		return "Kid’s Pages"
	case Category_TOYS_4:
		return "Toys"
	case Category_KIDS_4:
		return "Kids - Other"
	case Category_DATING_4:
		return "Dating & Relationships"
	case Category_DIVORCE_4:
		return "Divorce Support"
	case Category_ETHNIC_4:
		return "Ethnic Specific"
	case Category_MARRIAGE_4:
		return "Marriage"
	case Category_PARKS_4:
		return "Parks, Rec Facilities & Gyms"
	case Category_SENIOR_4:
		return "Senior Living"
	case Category_TEENS_4:
		return "Teens"
	case Category_WEDDING_4:
		return "Weddings"
	case Category_LIFESTYLE_4:
		return "Lifestyle - Other"
	case Category_ADFRAUD_4:
		return "Ad Fraud"
	case Category_BOTS_4:
		return "Botnet"
	case Category_CANDC_4:
		return "Command and Control Centers"
	case Category_COMPR_4:
		return "Compromised & Links To Malware"
	case Category_MALHOME_4:
		return "Malware Call-Home"
	case Category_MAL_4:
		return "Malware Distribution Point"
	case Category_FRAUD_4:
		return "Phishing/Fraud"
	case Category_SPAM_4:
		return "Spam URLs"
	case Category_SPYWARE_4:
		return "Spyware & Questionable Software"
	case Category_CDN_4:
		return "Content Server"
	case Category_BLANK_4:
		return "No Content Found"
	case Category_PARKED_4:
		return "Parked & For Sale Domains"
	case Category_PRIVIP_4:
		return "Private IP Address"
	case Category_BROKEN_4:
		return "Unreachable"
	case Category_MISC_4:
		return "Miscellaneous - Other"
	case Category_IMGSEARCH_4:
		return "Image Search"
	case Category_INTLNEWS_4:
		return "International News"
	case Category_LOCALNEWS_4:
		return "Local News"
	case Category_MAGAZINES_4:
		return "Magazines"
	case Category_NATLNEWS_4:
		return "National News"
	case Category_PORTAL_4:
		return "Portal Sites"
	case Category_SEARCH_4:
		return "Search Engines"
	case Category_NEWS_4:
		return "News, Portal & Search - Other"
	case Category_PAYTOSURF_4:
		return "Pay To Surf"
	case Category_ADWARE_4:
		return "Online Ads - Other"
	case Category_AQUARIUM_4:
		return "Aquariums"
	case Category_BIRD_4:
		return "Birds"
	case Category_CAT_4:
		return "Cats"
	case Category_DOG_4:
		return "Dogs"
	case Category_BIGANIMALS_4:
		return "Large Animals"
	case Category_REPTILE_4:
		return "Reptiles"
	case Category_VETMED_4:
		return "Veterinary Medicine"
	case Category_PETS_4:
		return "Pets - Other"
	case Category_ADVOCACY_4:
		return "Advocacy Groups & Trade Associations"
	case Category_COMMENTARY_4:
		return "Commentary"
	case Category_GOVSPONSOR_4:
		return "Government Sponsored"
	case Category_IMMIGRANT_4:
		return "Immigration"
	case Category_LEGAL_4:
		return "Legal Issues"
	case Category_CHARITY_4:
		return "Philanthropic Organizations"
	case Category_POLITICS_4:
		return "Politics"
	case Category_CLUBS_4:
		return "Social & Affiliation Organizations"
	case Category_USRESOURCE_4:
		return "U.S. Government Resources"
	case Category_GOV_4:
		return "Public, Government & Law - Other"
	case Category_APARTMENT_4:
		return "Apartments"
	case Category_ARCHI_4:
		return "Architects"
	case Category_BNSHOME_4:
		return "Buying/Selling Homes"
	case Category_PROP_4:
		return "Real Estate - Other"
	case Category_ALTRELGN_4:
		return "Alternative Religions"
	case Category_ATHEISM_4:
		return "Atheism & Agnosticism"
	case Category_BUDDHISM_4:
		return "Buddhism"
	case Category_CATHOLIC_4:
		return "Catholicism"
	case Category_CHRISTIAN_4:
		return "Christianity"
	case Category_HINDU_4:
		return "Hinduism"
	case Category_ISLAM_4:
		return "Islam"
	case Category_JUDAISM_4:
		return "Judaism"
	case Category_LDS_4:
		return "Latter-Day Saints"
	case Category_CULT_4:
		return "Non-traditional Religion & Occult"
	case Category_PAGAN_4:
		return "Pagan/Wiccan"
	case Category_RELIGION_4:
		return "Religion - Other"
	case Category_ANATOMY_4:
		return "Anatomy"
	case Category_ASTROLOGY_4:
		return "Astrology & Horoscopes"
	case Category_BIOLOGY_4:
		return "Biology"
	case Category_BOTANY_4:
		return "Botany"
	case Category_CHEMISTRY_4:
		return "Chemistry"
	case Category_GEOGRAPHY_4:
		return "Geography"
	case Category_GEOLOGY_4:
		return "Geology"
	case Category_PARANORMAL_4:
		return "Paranormal Phenomena"
	case Category_PHYSICS_4:
		return "Physics"
	case Category_ASTRONOMY_4:
		return "Space/Astronomy"
	case Category_WEATHER_4:
		return "Weather"
	case Category_SCIENCE_4:
		return "Science - Other"
	case Category_P2PSHOP_4:
		return "Auctions & Marketplaces"
	case Category_CATALOG_4:
		return "Catalogs"
	case Category_SURVEY_4:
		return "Contests & Surveys"
	case Category_COUPON_4:
		return "Coupons"
	case Category_SHOPENGINE_4:
		return "Engines"
	case Category_ONLINESHOP_4:
		return "Online Shopping"
	case Category_SHOPSEARCH_4:
		return "Product Reviews & Price Comparisons"
	case Category_SHOP_4:
		return "Shopping - Other"
	case Category_AUTORACE_4:
		return "Auto Racing"
	case Category_BASEBALL_4:
		return "Baseball"
	case Category_BICYCLING_4:
		return "Bicycling"
	case Category_BODYBUILD_4:
		return "Bodybuilding"
	case Category_BOXING_4:
		return "Boxing"
	case Category_KAYAK_4:
		return "Canoeing/Kayaking"
	case Category_CHEERLEAD_4:
		return "Cheerleeding"
	case Category_CLIMB_4:
		return "Climbing"
	case Category_CRICKET_4:
		return "Cricket"
	case Category_FIGSKATE_4:
		return "Figure Skating"
	case Category_FLYFISHING_4:
		return "Fly Fishing"
	case Category_FOOTBALL_4:
		return "Football"
	case Category_FRESHFISH_4:
		return "Freshwater Fishing"
	case Category_GAMEFISH_4:
		return "Game & Fish"
	case Category_GOLF_4:
		return "Golf"
	case Category_HORSERACE_4:
		return "Horse Racing"
	case Category_HORSES_4:
		return "Horses"
	case Category_INSKATE_4:
		return "Inline Skating"
	case Category_MARTIALART_4:
		return "Martial Arts"
	case Category_BIKING_4:
		return "Mountain Biking"
	case Category_NASCAR_4:
		return "NASCAR Racing"
	case Category_OLYMPICS_4:
		return "Olympics"
	case Category_PAINTBALL_4:
		return "Paintball"
	case Category_MOTORACE_4:
		return "Power & Motorcycles"
	case Category_PROBASKET_4:
		return "Pro Basketball"
	case Category_ICEHOCKEY_4:
		return "Pro Ice Hockey"
	case Category_RODEO_4:
		return "Rodeo"
	case Category_RUGBY_4:
		return "Rugby"
	case Category_RUN_4:
		return "Running/Jogging"
	case Category_SAIL_4:
		return "Sailing"
	case Category_SALTFISH_4:
		return "Saltwater Fishing"
	case Category_SCUBA_4:
		return "Scuba Diving"
	case Category_SKATEBOARD_4:
		return "Skateboarding"
	case Category_SKI_4:
		return "Skiing"
	case Category_SNOWBOARD_4:
		return "Snowboarding"
	case Category_HUNT_4:
		return "Sport Hunting"
	case Category_SURF_4:
		return "Surfing/Bodyboarding"
	case Category_SWIM_4:
		return "Swimming"
	case Category_PINGPONG_4:
		return "Table Tennis/Ping-Pong"
	case Category_TENNIS_4:
		return "Tennis"
	case Category_VOLLEYBALL_4:
		return "Volleyball"
	case Category_WALK_4:
		return "Walking"
	case Category_WATERSKI_4:
		return "Waterski/Wakeboard"
	case Category_SOCCER_4:
		return "World Soccer"
	case Category_SPORTS_4:
		return "Sports - Other"
	case Category_GRAPHICS3D_4:
		return "3-D Graphics"
	case Category_ANIMATION_4:
		return "Animation"
	case Category_ANTIVIRUS_4:
		return "Antivirus Software"
	case Category_CLANG_4:
		return "C/C++"
	case Category_CAMERA_4:
		return "Cameras & Camcorders"
	case Category_COMPCERT_4:
		return "Computer Certification"
	case Category_COMPNET_4:
		return "Computer Networking"
	case Category_PERIPHERAL_4:
		return "Computer Peripherals"
	case Category_COMPREVIEW_4:
		return "Computer Reviews"
	case Category_DATABASE_4:
		return "Databases"
	case Category_DESKPUB_4:
		return "Desktop Publishing"
	case Category_DESKVID_4:
		return "Desktop Video"
	case Category_TECHENT_4:
		return "Entertainment"
	case Category_REPOS_4:
		return "File Repositories"
	case Category_GRAPHICS_4:
		return "Graphics Software"
	case Category_DVD_4:
		return "Home Video/DVD"
	case Category_COMPSEC_4:
		return "Information Security"
	case Category_VOIP_4:
		return "Internet Phone & VOIP"
	case Category_INTERNET_4:
		return "Internet Technology"
	case Category_JAVA_4:
		return "Java"
	case Category_JAVASCRIPT_4:
		return "Javascript"
	case Category_LINUX_4:
		return "Linux"
	case Category_MAC_4:
		return "Mac OS"
	case Category_MACSUPPORT_4:
		return "Mac Support"
	case Category_MOBILE_4:
		return "Mobile Phones"
	case Category_MP3_4:
		return "MP3/MIDI"
	case Category_NETCON_4:
		return "Net Conferencing"
	case Category_NETBEG_4:
		return "Net for Beginners"
	case Category_NETSEC_4:
		return "Network Security"
	case Category_OFFICE_4:
		return "Online Information Management"
	case Category_PDA_4:
		return "Palmtops/PDAs"
	case Category_PCSUPPORT_4:
		return "PC Support"
	case Category_P2P_4:
		return "Peer-to-Peer"
	case Category_STORAGE_4:
		return "Personal Storage"
	case Category_PORTABLE_4:
		return "Portable"
	case Category_VPN_4:
		return "Remote Access"
	case Category_FREEWARE_4:
		return "Shareware/Freeware"
	case Category_UNIX_4:
		return "Unix"
	case Category_UTILITIES_4:
		return "Utilities"
	case Category_VBASIC_4:
		return "Visual Basic"
	case Category_CLIPART_4:
		return "Web Clip Art"
	case Category_WEBDESIGN_4:
		return "Web Design/HTML"
	case Category_ISP_4:
		return "Web Hosting, ISP & Telco"
	case Category_WINDOWS_4:
		return "Windows"
	case Category_TECH_4:
		return "Technology - Other"
	case Category_ADVENTURE_4:
		return "Adventure Travel"
	case Category_AFRICA_4:
		return "Africa"
	case Category_AIRTRAVEL_4:
		return "Air Travel"
	case Category_AUSTRALIA_4:
		return "Australia & New Zealand"
	case Category_BNB_4:
		return "Bed & Breakfast"
	case Category_BDGTTRAVEL_4:
		return "Budget Travel"
	case Category_BIZTRAVEL_4:
		return "Business Travel"
	case Category_USTRAVEL_4:
		return "By US Locale"
	case Category_CAMP_4:
		return "Camping"
	case Category_CANADA_4:
		return "Canada"
	case Category_CARRIBEAN_4:
		return "Caribbean"
	case Category_CRUISE_4:
		return "Cruises"
	case Category_EASTEUROPE_4:
		return "Eastern Europe"
	case Category_EUROPE_4:
		return "Europe"
	case Category_FRANCE_4:
		return "France"
	case Category_GREECE_4:
		return "Greece"
	case Category_GETAWAY_4:
		return "Honeymoons/Getaways"
	case Category_HOTEL_4:
		return "Hotels"
	case Category_ITALY_4:
		return "Italy"
	case Category_JAPAN_4:
		return "Japan"
	case Category_MEXICO_4:
		return "Mexico & Central America"
	case Category_NATLPARKS_4:
		return "National Parks"
	case Category_NAVIGATION_4:
		return "Navigation"
	case Category_SAMERICA_4:
		return "South America"
	case Category_SPA_4:
		return "Spas"
	case Category_THEMEPARK_4:
		return "Theme Parks"
	case Category_KIDTRAVEL_4:
		return "Traveling with Kids"
	case Category_UKTRAVEL_4:
		return "United Kingdom"
	case Category_TRAVEL_4:
		return "Travel - Other"
	default:
		return "Invalid Category"
	}
}
