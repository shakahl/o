package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Don't search for a corresponding header/source file for longer than ~0.5 seconds
var fileSearchMaxTime = 500 * time.Millisecond

// ExtFileSearch will search for a corresponding file, given a slice of extensions.
// This is useful for ie. finding a corresponding .h file for a .c file.
// The search starts in the current directory, then searches every parent directory in depth.
// TODO: Search sibling and parent directories named "include" first, then search the rest.
func ExtFileSearch(absCppFilename string, headerExtensions []string, maxTime time.Duration) (string, error) {
	cppBasename := filepath.Base(absCppFilename)
	searchPath := filepath.Dir(absCppFilename)
	ext := filepath.Ext(cppBasename)
	if ext == "" {
		return "", errors.New("filename has no extension: " + cppBasename)
	}
	firstName := cppBasename[:len(cppBasename)-len(ext)]

	// First search the same path as the given filename, without using Walk
	withoutExt := strings.TrimSuffix(absCppFilename, ext)
	for _, hext := range headerExtensions {
		if exists(withoutExt + hext) {
			return withoutExt + hext, nil
		}
	}

	var headerNames []string
	for _, ext := range headerExtensions {
		headerNames = append(headerNames, firstName+ext)
	}
	foundHeaderAbsPath := ""
	startTime := time.Now()
	for {
		err := filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
			basename := filepath.Base(info.Name())
			if err == nil {
				//logf("Walking %s\n", path)
				for _, headerName := range headerNames {
					if time.Since(startTime) > maxTime {
						return errors.New("file search timeout")
					}
					if basename == headerName {
						// Found the corresponding header!
						absFilename, err := filepath.Abs(path)
						if err != nil {
							continue
						}
						foundHeaderAbsPath = absFilename
						//logf("Found %s!\n", absFilename)
						return nil
					}
				}
			}
			// No result
			return nil
		})
		if err != nil {
			return "", errors.New("error when searching for a corresponding header for " + cppBasename + ":" + err.Error())
		}
		if len(foundHeaderAbsPath) == 0 {
			// Try the parent directory
			searchPath = filepath.Dir(searchPath)
			if len(searchPath) > 2 {
				continue
			}
		}
		break
	}
	if len(foundHeaderAbsPath) == 0 {
		return "", errors.New("found no corresponding header for " + cppBasename)
	}

	// Return the result
	return foundHeaderAbsPath, nil
}
