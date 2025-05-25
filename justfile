TEMPLATE_DIR := "templates"

md:
    REUMES_PATH={{TEMPLATE_DIR}} go run . ./schema.yaml -t markdown && bat -pp Resume.md
