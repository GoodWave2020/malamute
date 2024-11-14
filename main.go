package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	malamuteDir = ".malamute"
)

var hooks = []string{"pre-commit", "pre-push"}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: malamute <commmand>")
		fmt.Println("Commands:")
		fmt.Println("  init   Initialize malamute and set up Git hooks")
		fmt.Println("  reset  Reset Git hooks to default and remove .malamute directory")
		return
	}

	command := os.Args[1]
	switch command {
	case "init":
		err := setupMalamute()
		if err != nil {
			fmt.Println("Error setting up Malamute:", err)
		} else {
			fmt.Println("Malamute setup completed successfully.")
		}
	case "reset":
		err := resetMalamute()
		if err != nil {
			fmt.Println("Error resetting Malamute:", err)
		} else {
			fmt.Println("Malamute reset completed successfully.")
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

func setupMalamute() error {
	if _, err := os.Stat(malamuteDir); os.IsNotExist(err) {
		err := os.Mkdir(malamuteDir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create %s directory: %w", malamuteDir, err)
		}
		fmt.Printf("Created %s directory\n", malamuteDir)
	}

	err := createHooks(malamuteDir, hooks)
	if err != nil {
		return err
	}

	err = setGitHooksPath(malamuteDir)
	if err != nil {
		return fmt.Errorf("failed to set Git hooks path: %w", err)
	}

	return nil
}

func createHooks(dir string, hooks []string) error {
	for _, hook := range hooks {
		hookPath := filepath.Join(dir, hook)
		if _, err := os.Stat(hookPath); os.IsNotExist(err) {
			file, err := os.Create(hookPath)
			if err != nil {
				return fmt.Errorf("falled to create %s: %w", hookPath, err)
			}
			file.WriteString("#!/bin/sh\n# " + hook + " hook\n")
			file.Close()
			os.Chmod(hookPath, 0755)
			fmt.Printf("Created %s script\n", hookPath)
		}
	}
	return nil
}

func setGitHooksPath(path string) error {
	cmd := exec.Command("git", "config", "core.hooksPath", path)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error setting core.hooksPath: %w", err)
	}
	fmt.Printf("Set Git hooks path to %s\n", path)
	return nil
}

func resetGitHooksPath() error {
	cmd := exec.Command("git", "config", "--unset", "core.hooksPath")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error unsetting core.hooksPath: %w", err)
	}
	fmt.Println("Unset Git hooks path")
	return nil
}

func removeMalamuteDir() error {
	err := os.RemoveAll(malamuteDir)
	if err != nil {
		return fmt.Errorf("failed to remove %s directory: %w", malamuteDir, err)
	}
	fmt.Printf("Removed %s directory\n", malamuteDir)
	return nil
}

func resetMalamute() error {
	err := resetGitHooksPath()
	if err != nil {
		return err
	}
	err = removeMalamuteDir()
	if err != nil {
		return err
	}
	return nil
}
