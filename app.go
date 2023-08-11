package main

import (
	"context"

	"fmt"
	"io"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// import nodal degree file
func (a *App) ImportNodalDegree() string {
	var err error
	fileLocation, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title:           "Open Nodal Degree Json file",
		ShowHiddenFiles: true,
	})
	NodalDegreeFile, err := os.Open(fileLocation)
	defer NodalDegreeFile.Close()
	byteNodalDegreeData, _ := io.ReadAll(NodalDegreeFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(byteNodalDegreeData)
}
