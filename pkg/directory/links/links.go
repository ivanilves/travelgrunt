// Package links provides utilities for resolving and validating file paths,
// particularly for collecting directory paths relative to a root directory.
// It supports tilde (~) expansion for the current user's home directory and
// ensures paths are absolute and valid directories.
package links

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// ExpandPath expands a path starting with "~" to the current user's home directory.
// It supports "~" (home dir) and "~/path" but not "~user/path".
// Returns the expanded path or an error if the expansion fails.
func ExpandPath(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}

	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("failed to get current user: %w", err)
	}

	if path == "~" {
		return usr.HomeDir, nil
	}

	if len(path) > 1 && (path[1] == '/' || path[1] == filepath.Separator) {
		return filepath.Join(usr.HomeDir, path[2:]), nil
	}

	return "", fmt.Errorf("user home directory expansion (~user) is not supported")
}

// GetAbsPath converts a path to an absolute path, resolving relative to rootPath.
// It expands tilde, ensures the result is absolute, and cleans the path.
// Returns an error for empty paths or invalid expansions.
func GetAbsPath(path, rootPath string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("path cannot be empty")
	}

	// Expand tilde first
	expandedPath, err := ExpandPath(path)
	if err != nil {
		return "", err
	}

	// If already absolute, clean and return (no traversal check needed)
	if filepath.IsAbs(expandedPath) {
		return filepath.Clean(expandedPath), nil
	}

	// For relative paths: ensure rootPath is absolute
	absRoot, err := filepath.Abs(rootPath)
	if err != nil {
		return "", fmt.Errorf("failed to resolve absolute root path: %w", err)
	}

	// Join and clean
	absPath := filepath.Join(absRoot, expandedPath)
	absPath = filepath.Clean(absPath)

	// Prevent path traversal outside root (only for relative paths)
	if !strings.HasPrefix(absPath, absRoot+string(filepath.Separator)) {
		return "", fmt.Errorf("path %q resolves outside root %q", path, absRoot)
	}

	return absPath, nil
}

// ValidatePath checks if a path exists and is a directory.
// Returns specific errors for non-existence, permission issues, or non-directory paths.
func ValidatePath(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("path %q does not exist", path)
		}
		if os.IsPermission(err) {
			return fmt.Errorf("permission denied for path %q", path)
		}
		return fmt.Errorf("cannot access path %q: %w", path, err)
	}

	if !info.IsDir() {
		return fmt.Errorf("path %q is not a directory", path)
	}

	return nil
}

// Collect processes a list of link paths, resolving and validating them as directories.
// It returns a map of original to absolute paths, a slice of valid original paths,
// and an error if validation fails or no valid paths remain.
// All paths are resolved relative to rootPath, which is made absolute.
func Collect(rootPath string, linkPaths []string) (map[string]string, []string, error) {
	if len(linkPaths) == 0 {
		return nil, nil, fmt.Errorf("no links defined in config")
	}

	linkToAbsPath := make(map[string]string, len(linkPaths))
	validLinks := make([]string, 0, len(linkPaths))
	var validationErrors []string

	for _, link := range linkPaths {
		if link == "" {
			continue // Skip empty links
		}

		// Resolve to absolute path
		absPath, err := GetAbsPath(link, rootPath)
		if err != nil {
			validationErrors = append(validationErrors, fmt.Sprintf("link %q: %v", link, err))
			continue
		}

		// Validate path
		if err := ValidatePath(absPath); err != nil {
			validationErrors = append(validationErrors, fmt.Sprintf("link %q (resolved to %q): %v", link, absPath, err))
			continue
		}

		linkToAbsPath[link] = absPath
		validLinks = append(validLinks, link)
	}

	if len(validationErrors) > 0 {
		return nil, nil, fmt.Errorf("link validation failed:\n  - %s", strings.Join(validationErrors, "\n  - "))
	}

	if len(validLinks) == 0 {
		return nil, nil, fmt.Errorf("no valid links found in config")
	}

	return linkToAbsPath, validLinks, nil
}
