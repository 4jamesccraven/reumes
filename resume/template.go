package resume

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/noirbizarre/gonja"
)

type Template struct {
	Name        string `yaml:"name"`
	Text        string `yaml:"text"`
	OutputFile  string `yaml:"outputFile"`
	PostProcess string `yaml:"postProcess"`
}

func (self *Template) Render(res Resume) error {
	// Get the user template
	tpl, err := gonja.FromString(self.Text)
	if err != nil {
		return err
	}

	// Use the Resume to populate it
	out, err := tpl.Execute(res.ExportContext())
	if err != nil {
		return err
	}

	if self.PostProcess != "" {
		// Create a temporary file to be post-processed
		tmpFile, err := os.CreateTemp("", "reumes-*")
		if err != nil {
			return err
		}
		defer os.Remove(tmpFile.Name())

		if _, err := tmpFile.WriteString(out); err != nil {
			return err
		}

		// Result of template
		cmd := strings.ReplaceAll(self.PostProcess, "%template_result%", tmpFile.Name())
		// Output file
		cmd = strings.ReplaceAll(cmd, "%output_name%", self.OutputFile)

		// Post-process the file
		var childHandle *exec.Cmd
		if runtime.GOOS == "windows" {
			childHandle = exec.Command("cmd", "/C", cmd)
		} else {
			childHandle = exec.Command("sh", "-c", cmd)
		}

		_, err = childHandle.CombinedOutput()
		if err != nil {
			return err
		}
	} else {
		// Use the template's output file, or supply a default
		outputFile := self.OutputFile
		if outputFile == "" {
			outputFile = "reumes.out"
		}

		// Make the file
		file, err := os.Create(outputFile)
		if err != nil {
			return err
		}
		defer file.Close()

		// Write the output
		_, err = file.WriteString(out)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetTemplates() map[string]Template {
	templates := make(map[string]Template)

	paths := strings.Split(os.Getenv("REUMES_PATH"), ":")
	cwd, err := os.Getwd()
	if err == nil {
		paths = append(paths, cwd)
	}

	for _, path := range paths {
		entries, err := os.ReadDir(path)

		if err != nil {
			continue
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			path := filepath.Join(path, entry.Name())
			data, err := os.ReadFile(path)
			if err != nil {
				continue
			}

			var tmpl Template
			if err := yaml.Unmarshal(data, &tmpl); err != nil {
				continue
			}

			templates[tmpl.Name] = tmpl
		}
	}

	return templates
}
