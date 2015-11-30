package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

func main() {
	if err := Main(); err != nil {
		panic(err)
	}
}

func Main() error {
	// f, err := os.Create("README.md")
	// if err != nil {
	// 	return err
	// }
	f := os.Stdout
	defer f.Close()

	b, err := Asset("readme.template")
	if err != nil {
		return err
	}

	a, err := Arg()
	if err != nil {
		return err
	}

	t := template.Must(template.New("README.md").Parse(string(b)))
	if err := t.Execute(f, a); err != nil {
		return err
	}
	return nil
}

type TemplateArg struct {
	Name  string
	IsGo  bool
	GoURL string
	Make  bool
}

func Arg() (*TemplateArg, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	_, name := filepath.Split(path)

	isGo, goURL := IsGo(path)

	return &TemplateArg{
		Name:  name,
		IsGo:  isGo,
		GoURL: goURL,
		Make:  FileExist("Makefile"),
	}, nil
}

func IsGo(pwd string) (bool, string) {
	pathes := SplitPATH(os.Getenv("GOPATH"))
	for _, path := range pathes {
		if filepath.HasPrefix(pwd, path) {
			return true, strings.TrimPrefix(pwd, filepath.Join(path, "src")+"/")
		}
	}
	return false, ""
}

// /bin:/usr/bin => {"/bin", "/usr/bin"}
func SplitPATH(path string) []string {
	var sp string
	if runtime.GOOS == "windows" {
		sp = ";"
	} else {
		sp = ":"
	}
	return strings.Split(path, sp)
}

func FileExist(fname string) bool {
	_, err := os.Stat(fname)
	return err == nil
}
