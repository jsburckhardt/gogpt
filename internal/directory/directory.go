// Package directory used for working with directory tasks
package directory

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Scripts is a wrapper around the embedded scripts
var Scripts embed.FS

// Init home directory
func Init() error {
	homedir, err := Homedir()
	if err != nil {
		return err
	}

	gogptDir := filepath.Join(homedir, ".gogpt")
	err = os.RemoveAll(gogptDir)
	if err != nil {
		return err
	}
	err = os.Mkdir(gogptDir, 0755)
	if err != nil {
		return err
	}
	return nil
}

// Homedir returns the home directory of the user
func Homedir() (string, error) {
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	if home := os.Getenv("USERPROFILE"); home != "" {
		return home, nil
	}
	return "", errors.New("Unable to determine home directory (HOME/USERPROFILE)")
}

// initfolder initialises the gogptdir and copies the asset files to the given path
func initfolder(gogptDir, folder string) error {
	// func initfolder(gogptDir, folder string, logger *adapter.Logger) error {
	// copy the scripts to the ~/.gogpt directory
	gogptInitFolder := filepath.Join(gogptDir, folder)
	err := os.Mkdir(gogptInitFolder, 0755)
	if err != nil {
		return err
	}
	folderAsset := fmt.Sprintf("assets/%s", folder)
	files, err := Scripts.ReadDir(folderAsset)
	if err != nil {
		return err
	}
	for _, file := range files {
		// logger.Infof("copying script %s", file.Name())
		assetPath := filepath.Join(folderAsset, file.Name())
		scriptPath := filepath.Join(gogptInitFolder, file.Name())
		scriptData, err := Scripts.ReadFile(assetPath)
		if err != nil {
			return err
		}
		err = os.WriteFile(scriptPath, scriptData, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

// CleanPathArgument checks if it an absolute path or relatieve and removes /
func CleanPathArgument(packagePath string) (string, error) {
	if !filepath.IsAbs(packagePath) {
		currentDir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		packagePath = filepath.Join(currentDir, packagePath)
	}
	// remove / at the end of the path if it exists
	if packagePath[len(packagePath)-1:] == "/" {
		packagePath = packagePath[:len(packagePath)-1]
	}
	return packagePath, nil
}

// ValidateIfPathExists checks if the given path exists
func ValidateIfPathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// CreateDirectory creates a directory in the given path with 755 permissions
func CreateDirectory(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	return nil
}

// CreateFile create a file in the given path with 666 permissions
func CreateFile(path string) error {
	_, err := os.Create(path)
	if err != nil {
		return err
	}
	return nil
}

// PathWalk will walk up the path until it finds the folder you're looking for
func PathWalk(folderPath string) (string, error) {
	targetPath := folderPath

	var abs string

	for {

		newAbs, _ := filepath.Abs(targetPath)
		if newAbs == abs {
			return "", fmt.Errorf(fmt.Sprintf("Could not find path %v", folderPath))
		}

		abs = newAbs

		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			targetPath = filepath.Join("..", targetPath)
			continue
		}

		return targetPath, nil
	}
}

// FileWalk will walk up the path until it finds the file you're looking for
func FileWalk(folderPath string, fileName string) (string, error) {
	targetPath := folderPath

	for {

		searchFile := filepath.Join(targetPath, fileName)

		if _, err := os.Stat(searchFile); os.IsNotExist(err) {
			targetPath = filepath.Dir(targetPath)
			if targetPath == "/" {
				return "", fmt.Errorf(fmt.Sprintf("Could not find path %v", folderPath))
			}
			continue
		}

		return targetPath, nil
	}
}
