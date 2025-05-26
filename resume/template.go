package resume

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
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

	// Do post process if necessary, or dump the template
	if self.PostProcess != "" {
		return self.DoPostProcess(out)
	}
	return self.DumpTemplate(out)
}

func (self *Template) DoPostProcess(out string) error {
	// Create a temporary file to be post-processed
	tmpFile, err := ReumesTempFile(out)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile)

	// Substitute the result of template...
	cmd := strings.ReplaceAll(self.PostProcess, "%template_result%", tmpFile)
	// ... and the output file
	cmd = strings.ReplaceAll(cmd, "%output_name%", self.OutputFile)

	// Post-process the file, using sh or cmd
	var childHandle *exec.Cmd
	if runtime.GOOS != "windows" {
		childHandle = exec.Command("sh", "-c", cmd)
	} else {
		childHandle = exec.Command("cmd", "/C", cmd)
	}

	// Get the output of the commands, reporting errors if necessary
	commandOutput, err := childHandle.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(commandOutput))
		return err
	}

	return nil
}

func (self *Template) DumpTemplate(out string) error {
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

	return nil
}

// Create a temporary file in the current dir to post process
func ReumesTempFile(content string) (string, error) {
	// Make a random string of digits for the file name
	bytes := make([]byte, 10)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	suffix := hex.EncodeToString(bytes)

	// Create the temporary file
	filename := filepath.Join(".", "reumes-"+suffix)
	file, err := os.Create(filename)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	// Write the contents to the file
	if _, err := file.WriteString(content); err != nil {
		return "", nil
	}

	return filename, nil
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
