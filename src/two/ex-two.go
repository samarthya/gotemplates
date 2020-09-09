package two

// Letter a template for the automatic reply.
const Letter = `
Dear {{.Name}},
{{/* Check if the person identified by .Name attended the event. */}}
{{if .Attended}}
It was a pleasure to meet you at the wedding.
{{- else}}
{{/* Person identified by .Name did not attended the event. */}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

// Recipient Identifies the recipient
type Recipient struct {
	Name, Gift string
	Attended   bool
}
