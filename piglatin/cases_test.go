package piglatin

type piglatinTest struct {
	input    string
	expected string
}

var stringTestCases = []piglatinTest{
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
		input:    "the",
		expected: "ethay",
	},
	{
		input:    "Привет Мир!",
		expected: "Привет Мир!",
	},
	{
		input:    "dddddd",
		expected: "dddddday",
	},
}
