// Copyright 2018, Goomba project Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package namegenerator

var (
	// ADJECTIVES ...
	ADJECTIVES = []string{"abundant", "acidic", "aged", "aggressive", "agreeable", "alive", "ambitious", "ancient",
		"angry", "ashy", "attractive", "autumn", "bald", "beautiful", "better", "bewildered",
		"big", "billions", "billowing", "bitter", "black", "blue", "bold", "brave",
		"breezy", "brief", "broad", "broken", "bumpy", "calm", "careful", "chilly",
		"chubby", "clean", "clever", "clumsy", "cold", "colossal", "cool", "crashing",
		"creamy", "crimson", "crooked", "cuddly", "curved", "damaged", "damp", "dark",
		"dawn", "dazzling", "dead", "deafening", "deep", "defeated", "delicate", "delicious",
		"delightful", "dirty", "disgusting", "divine", "drab", "dry", "eager", "early",
		"easy", "echoing", "elegant", "embarrassed", "empty", "enough", "faint", "faithful",
		"falling", "famous", "fancy", "fast", "fat", "few", "fierce", "fit",
		"flabby", "flaky", "flat", "floral", "fluffy", "fragrant", "freezing", "fresh",
		"frosty", "full", "future", "gentle", "gifted", "gigantic", "glamorous", "gorgeous",
		"gray", "greasy", "great", "green", "grumpy", "hallowed", "handsome", "happy",
		"harsh", "helpful", "helpless", "hidden", "high", "hissing", "hollow", "holy",
		"hot", "howling", "huge", "hundreds", "icy", "immense", "important", "incalculable",
		"inexpensive", "itchy", "jealous", "jolly", "juicy", "kind", "large", "late",
		"lazy", "lemon", "limited", "lingering", "little", "lively", "long", "loose",
		"loud", "low", "magnificent", "mammoth", "mango", "many", "massive", "mealy",
		"melodic", "melted", "microscopic", "millions", "miniature", "misty", "modern", "moldy",
		"morning", "most", "muddy", "muscular", "mushy", "mysterious", "nameless", "narrow",
		"nervous", "nice", "noisy", "numerous", "nutritious", "nutty", "obedient", "obnoxious",
		"odd", "old", "old-fashioned", "orange", "panicky", "patient", "petite", "pitiful",
		"plain", "plump", "polished", "polite", "poor", "powerful", "prehistoric", "prickly",
		"proud", "puny", "purple", "purring", "putrid", "quaint", "quick", "quiet",
		"rancid", "rapid", "rapping", "raspy", "red", "refined", "repulsive", "restless",
		"rhythmic", "rich", "ripe", "rotten", "rough", "round", "salmon", "salty",
		"savory", "scarce", "scary", "scrawny", "screeching", "scruffy", "shaggy", "shallow",
		"shapely", "sharp", "short", "shrilling", "shy", "silent", "silly", "skinny",
		"slimy", "slow", "small", "snowy", "solitary", "some", "sour", "sparkling",
		"sparse", "spicy", "spoiled", "spring", "square", "squeaking", "stale", "steep",
		"sticky", "still", "stocky", "straight", "strong", "substantial", "summer", "sweet",
		"swift", "tall", "tangy", "tart", "tasteless", "tasty", "teeny", "tender",
		"thankful", "thoughtless", "thousands", "throbbing", "thundering", "tight", "tinkling", "tiny",
		"twilight", "ugly", "uneven", "unimportant", "uninterested", "unkempt", "unsightly", "uptight",
		"vast", "victorious", "wailing", "wandering", "warm", "weak", "weathered", "wet",
		"whining", "whispering", "white", "wide", "wild", "winter", "wispy", "withered",
		"witty", "wonderful", "wooden", "worried", "wrong", "yellow", "young", "yummy", "zealous"}

	// NOUNS ...
	NOUNS = []string{"ability", "accident", "activity", "actor", "ad", "addition", "administration", "advertising",
		"advice", "affair", "agency", "agreement", "airport", "alcohol", "ambition", "analysis",
		"analyst", "anxiety", "apartment", "appearance", "apple", "application", "appointment", "area",
		"argument", "army", "arrival", "art", "article", "aspect", "assignment", "assistance",
		"assistant", "association", "assumption", "atmosphere", "attention", "attitude", "audience", "awareness",
		"baseball", "basis", "basket", "bath", "bathroom", "bedroom", "beer", "bird",
		"birthday", "blood", "bonus", "boyfriend", "bread", "breath", "breeze", "brook",
		"bush", "butterfly", "buyer", "cabinet", "camera", "cancer", "candidate", "category",
		"celebration", "cell", "championship", "chapter", "charity", "cheek", "chemistry", "cherry",
		"chest", "child", "childhood", "chocolate", "church", "cigarette", "city", "classroom",
		"client", "climate", "clothes", "cloud", "coffee", "collection", "college", "combination",
		"committee", "communication", "community", "comparison", "competition", "complaint", "computer", "concept",
		"conclusion", "confusion", "connection", "consequence", "construction", "context", "contract", "contribution",
		"control", "conversation", "cookie", "country", "county", "courage", "cousin", "criticism",
		"currency", "customer", "dad", "darkness", "data", "database", "dawn", "dealer",
		"death", "debt", "decision", "definition", "delivery", "department", "departure", "depression",
		"depth", "description", "desk", "development", "device", "dew", "diamond", "difference",
		"difficulty", "dinner", "direction", "director", "dirt", "disaster", "discussion", "disease",
		"disk", "distribution", "drama", "drawer", "drawing", "dream", "driver", "dust",
		"ear", "economics", "editor", "education", "efficiency", "effort", "election", "elevator",
		"emotion", "emphasis", "employee", "employer", "employment", "energy", "engine", "engineering",
		"entertainment", "enthusiasm", "entry", "environment", "equipment", "error", "establishment", "estate",
		"event", "exam", "examination", "excitement", "explanation", "expression", "extent", "fact",
		"failure", "family", "farmer", "feather", "feedback", "field", "finding", "fire",
		"firefly", "fishing", "flight", "flower", "fog", "food", "football", "forest",
		"fortune", "foundation", "freedom", "friendship", "frog", "frost", "funeral", "garbage",
		"gate", "gene", "girl", "girlfriend", "glade", "glitter", "goal", "government",
		"grandmother", "grass", "grocery", "growth", "guest", "guidance", "guitar", "hair",
		"hall", "hat", "haze", "health", "hearing", "heart", "height", "highway",
		"hill", "historian", "history", "homework", "honey", "hospital", "hotel", "housing",
		"idea", "imagination", "importance", "impression", "improvement", "income", "independence", "indication",
		"industry", "inflation", "information", "initiative", "injury", "insect", "inspection", "inspector",
		"instance", "instruction", "insurance", "intention", "interaction", "internet", "introduction", "investment",
		"judgment", "king", "knowledge", "lab", "ladder", "lady", "lake", "language",
		"law", "leader", "leadership", "leaf", "length", "library", "literature", "location",
		"loss", "love", "magazine", "maintenance", "mall", "management", "manager", "manufacturer",
		"map", "marketing", "marriage", "math", "meadow", "meal", "meaning", "measurement",
		"meat", "media", "medicine", "member", "membership", "memory", "menu", "message",
		"method", "midnight", "mixture", "mode", "mom", "moment", "month", "mood",
		"moon", "morning", "mountain", "movie", "mud", "music", "nation", "nature",
		"negotiation", "news", "newspaper", "night", "obligation", "office", "operation", "opinion",
		"opportunity", "orange", "organization", "outcome", "oven", "owner", "painting", "paper",
		"passenger", "passion", "patience", "payment", "penalty", "people", "percentage", "perception",
		"performance", "permission", "person", "personality", "perspective", "philosophy", "phone", "photo",
		"physics", "piano", "pie", "pine", "pizza", "platform", "player", "poem",
		"poet", "poetry", "police", "policy", "politics", "pollution", "pond", "population",
		"possession", "possibility", "potato", "power", "preference", "preparation", "presence", "presentation",
		"president", "priority", "problem", "procedure", "product", "profession", "professor", "promotion",
		"property", "proposal", "protection", "psychology", "quality", "quantity", "queen", "rain",
		"ratio", "reaction", "reading", "reality", "reception", "recipe", "recognition", "recommendation",
		"recording", "reflection", "refrigerator", "region", "relation", "relationship", "replacement", "republic",
		"reputation", "requirement", "resolution", "resonance", "resource", "response", "responsibility", "restaurant",
		"revenue", "revolution", "river", "road", "role", "safety", "salad", "sample",
		"satisfaction", "scene", "science", "sea", "secretary", "sector", "security", "selection",
		"series", "session", "setting", "shadow", "shape", "shirt", "shopping", "signature",
		"significance", "silence", "singer", "sir", "sister", "situation", "skill", "sky",
		"smoke", "snow", "snowflake", "society", "software", "solution", "son", "song",
		"sound", "soup", "speaker", "speech", "star", "statement", "steak", "storage",
		"story", "stranger", "strategy", "student", "studio", "success", "suggestion", "sun",
		"sunset", "supermarket", "surf", "surgery", "sympathy", "system", "tale", "tea",
		"teacher", "teaching", "technology", "television", "temperature", "tennis", "tension", "thanks",
		"theory", "thing", "thought", "throat", "thunder", "tongue", "tooth", "topic",
		"town", "tradition", "trainer", "transportation", "tree", "truth", "two", "uncle",
		"understanding", "union", "unit", "university", "user", "variation", "variety", "vehicle",
		"version", "video", "village", "violet", "virus", "voice", "volume", "waterfall",
		"warning", "water", "wave", "way", "weakness", "wealth", "wedding", "week",
		"wife", "wildflower", "wind", "winner", "woman", "wood", "worker", "world",
		"writer", "writing", "year", "youth"}
)
