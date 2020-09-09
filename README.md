# Templates
This project will help talk about go templates as I try and learn how to use it. The initial objective was learning HELM which uses GO templates, and since it was obvious that I already know GO but have not worked on templates. It was natural to see what it is and utilise it as a pet project.

## Package template
Templates are executed by applying them to a data structure. 

Package template implements data-driven templates for generating textual output.

```bash
To generate HTML output, see package html/template, which has the same interface as this package but automatically secures HTML output against certain attacks.
```

Annotations in the template refer to elements of the data structure (typically a field of a struct or a key in a map) to control execution and derive values to be displayed. 

### Example

```
{{.Count}}
```

## Helpful links
- https://golang.org/pkg/text/template/
- https://golang.org/pkg/text/template/#hdr-Actions