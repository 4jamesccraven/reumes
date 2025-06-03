package resume

var templates map[string]Template

func init() {
	templates = LoadTemplates()
}

func GetTemplates() map[string]Template {
	return templates
}
