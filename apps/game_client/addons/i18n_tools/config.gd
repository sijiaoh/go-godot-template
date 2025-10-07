const languages = ["ja"]

const IGNORE_DIRS = [
	"res://addons/",
]

const EXTRA_SCAN_RULES = [
	# {"dir": "res://foo", "suffix": "_data.tres"},
	{"dir": "res://", "suffix": ".gd"},
	{"dir": "res://", "suffix": ".tscn"},
]

const NODE_PROPERTY_NAMES = ["_text", "text", "placeholder_text"]

const RESOURCE_PROPERTY_NAMES = ["text", "name"]
