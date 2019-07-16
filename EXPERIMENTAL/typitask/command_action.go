package typitask

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/internal/util"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typienv"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typigen"
	"gopkg.in/urfave/cli.v1"
)

func (t *TypicalTask) buildBinary(ctx *cli.Context) {
	typienv.GenerateProjectEnvIfNotExist(t.Context)

	typigen.AppSideEffects(t.Context)

	binaryPath := typienv.BinaryPath(t.BinaryNameOrDefault())
	mainPackage := typienv.MainPackage(t.AppPkgOrDefault())

	log.Printf("Build the Binary for '%s' at '%s'", mainPackage, binaryPath)
	util.RunOrFatal(util.GoCommand(), "build", "-o", binaryPath, mainPackage)
}

func (t *TypicalTask) runBinary(ctx *cli.Context) {
	if !ctx.Bool("no-build") {
		t.buildBinary(ctx)
	}

	binaryPath := typienv.BinaryPath(t.BinaryNameOrDefault())
	log.Printf("Run the Binary '%s'", binaryPath)
	util.RunOrFatal(binaryPath, []string(ctx.Args())...)
}

func (t *TypicalTask) runTest(ctx *cli.Context) {
	log.Println("Run the Test")
	args := []string{"test"}
	args = append(args, t.ArcheType.GetTestTargets()...)
	args = append(args, "-coverprofile=cover.out")
	util.RunOrFatal(util.GoCommand(), args...)
}

func (t *TypicalTask) releaseDistribution(ctx *cli.Context) {
	fmt.Println("Not implemented")
}

func (t *TypicalTask) generateMock(ctx *cli.Context) {
	util.RunOrFatal(util.GoCommand(), "get", "github.com/golang/mock/mockgen")
	mockPkg := t.MockPkgOrDefault()

	if ctx.Bool("new") {
		log.Printf("Clean mock package '%s'", mockPkg)
		os.RemoveAll(mockPkg)
	}

	for _, mockTarget := range t.ArcheType.GetMockTargets() {
		dest := mockPkg + "/" + mockTarget[strings.LastIndex(mockTarget, "/")+1:]

		log.Printf("Generate mock for '%s' at '%s'", mockTarget, dest)
		util.RunOrFatal(util.GoBinary("mockgen"),
			"-source", mockTarget,
			"-destination", dest,
			"-package", mockPkg)
	}
}

func (t *TypicalTask) generateReadme(ctx *cli.Context) (err error) {
	readmeFile := t.ReadmeFileOrDefault()
	readmeTemplate := t.ReadmeTemplateOrDefault()

	templ, err := template.New("readme").Parse(readmeTemplate)

	if err != nil {
		return
	}

	file, err := os.Create(readmeFile)
	if err != nil {
		return
	}

	log.Printf("Generate ReadMe Document at '%s'", readmeFile)
	err = templ.Execute(file, Readme{
		Context: t.Context,
	})
	return nil
}

func (t *TypicalTask) cleanProject(ctx *cli.Context) {
	log.Println("Remove bin folder")
	os.RemoveAll(typienv.Bin())

	log.Println("Trigger go clean")
	os.Setenv("GO111MODULE", "off") // NOTE:XXX: https://github.com/golang/go/issues/28680
	util.RunOrFatal(util.GoCommand(), "clean", "-x", "-testcache", "-modcache")
}