package style

import (
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
)

var HStyleMap map[string]document.Run

func init() {
	HStyleMap = make(map[string]document.Run)

	h1paragraph := Doc.AddParagraph()
	h1Run := h1paragraph.AddRun()
	h1Run.Properties().SetBold(true)
	h1Run.Properties().SetSize(measurement.Distance(H1FontSize))
	HStyleMap["h1"] = h1Run

	h2paragraph := Doc.AddParagraph()
	h2Run := h2paragraph.AddRun()
	h2Run.Properties().SetBold(true)
	h2Run.Properties().SetSize(measurement.Distance(H2FontSize))
	HStyleMap["h2"] = h2Run

	h3paragraph := Doc.AddParagraph()
	h3Run := h3paragraph.AddRun()
	h3Run.Properties().SetBold(true)
	h3Run.Properties().SetSize(measurement.Distance(H3FontSize))
	HStyleMap["h3"] = h3Run

	h4paragraph := Doc.AddParagraph()
	h4Run := h4paragraph.AddRun()
	h4Run.Properties().SetBold(true)
	h4Run.Properties().SetSize(measurement.Distance(H4FontSize))
	HStyleMap["h4"] = h4Run

	h5paragraph := Doc.AddParagraph()
	h5Run := h5paragraph.AddRun()
	h5Run.Properties().SetBold(true)
	h5Run.Properties().SetSize(measurement.Distance(H5FontSize))
	HStyleMap["h5"] = h5Run
}

func SetH(text string, tag string) {
	run := HStyleMap[tag]
	run.AddText(text)
	run.AddBreak()
}
