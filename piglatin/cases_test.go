package piglatin

type piglatinTest struct {
	input    string
	expected string
}
type alpahaTest struct {
	input    string
	expected bool
}

var encryptTestCases = []piglatinTest{
	{
		input:    "paris world",
		expected: "arispay orldway",
	},
	{
		input:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		expected: "Oremlay ipsumway olorday itsay ametway, onsecteturcay adipiscingway elitway, edsay oday eiusmodway emportay incididuntway utway aborelay etway oloreday agnamay aliquaway.",
	},
	{
		input:    "",
		expected: "",
	},
	{
		input:    "    ",
		expected: "    ",
	},
	{
		input:    "Привет Мир!",
		expected: "Привет Мир!",
	},
}

var encryptSingleWordTestCases = []piglatinTest{
	{
		input:    "paris",
		expected: "arispay",
	},
	{
		input:    "amazon",
		expected: "amazonway",
	},
	{
		input:    "pig",
		expected: "igpay",
	},
	{
		input:    "latin",
		expected: "atinlay",
	},
	{
		input:    "smile",
		expected: "ilesmay",
	},
	{
		input:    "the",
		expected: "ethay",
	},
	{
		input:    "dddddd",
		expected: "dddddday",
	},
}
var alpha = []alpahaTest{
	{
		input:    "a",
		expected: false,
	},
	{
		input:    "b",
		expected: true,
	},
	{
		input:    "c",
		expected: true,
	},
	{
		input:    "d",
		expected: true,
	},
	{
		input:    "e",
		expected: false,
	},
	{
		input:    "f",
		expected: true,
	},
	{
		input:    "g",
		expected: true,
	},
	{
		input:    "h",
		expected: true,
	},
	{
		input:    "i",
		expected: false,
	},
	{
		input:    "j",
		expected: true,
	},
	{
		input:    "k",
		expected: true,
	},
	{
		input:    "m",
		expected: true,
	},
	{
		input:    "n",
		expected: true,
	},
	{
		input:    "o",
		expected: false,
	},
	{
		input:    "p",
		expected: true,
	},
	{
		input:    "q",
		expected: true,
	},
	{
		input:    "r",
		expected: true,
	},
	{
		input:    "s",
		expected: true,
	},
	{
		input:    "t",
		expected: true,
	},
	{
		input:    "u",
		expected: false,
	},
	{
		input:    "v",
		expected: true,
	},
	{
		input:    "w",
		expected: true,
	},
	{
		input:    "x",
		expected: true,
	},
	{
		input:    "y",
		expected: true,
	},
	{
		input:    "z",
		expected: true,
	},
}
