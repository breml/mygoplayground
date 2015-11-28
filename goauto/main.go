package main

import (
	"path/filepath"

	"github.com/dshills/goauto"
	"github.com/dshills/goauto/gotask"
)

func main() {
	// Create a pipeline (Develop using Verbose, Change to Silent after testing)
	p := goauto.NewPipeline("Go Pipeline", goauto.Verbose)
	defer p.Stop()

	// Watch directories recursively, ignoring hidden directories
	wd := filepath.Join("src", "github.com", "breml", "gosampler")
	if err := p.WatchRecursive(wd, goauto.IgnoreHidden); err != nil {
		panic(err)
	}

	// Create a workflow
	wf := goauto.NewWorkflow(
		gotask.NewGoVetTask(),
		gotask.NewGoLintTask(),
		gotask.NewGoTestTask(),
		gotask.NewGoInstallTask())

	// Add a file pattern to match
	if err := wf.WatchPattern(".*\\.go$"); err != nil {
		panic(err)
	}

	// Add workflow to pipeline
	p.Add(wf)

	// start the pipeline, it will block
	p.Start()
}
