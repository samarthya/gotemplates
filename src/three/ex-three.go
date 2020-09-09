package three

// Uses {{block}}
/**
 * {{block "name" pipeline}} T1 {{end}}
 *	A block is shorthand for defining a template
 *		`{{define "name"}} T1 {{end}}`
 *	and then executing it in place
 *		`{{template "name" pipeline}}`
 *	The typical use is to define a set of root templates that are
 *	then customized by redefining the block templates within.
 *
 * {{range pipeline}} T1 {{end}}
 *	  The value of the pipeline must be an array, slice, map, or channel.
 *
 *	  If the value of the pipeline has length zero, nothing is output;
 *	  otherwise, dot is set to the successive elements of the array,
 *	  slice, or map and T1 is executed. If the value is a map and the
 *	  keys are of basic type with a defined order, the elements will be
 *	  visited in sorted key order.
 */
const (
	// Master Block Master
	Master = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`

	// Overlay Block Overlay
	Overlay = `{{define "list"}}{{join . ", "}}{{end}} `
)
