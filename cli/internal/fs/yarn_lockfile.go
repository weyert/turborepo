//go:build !windows
// +build !windows

package fs

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// func parseYarnLockfile(file string) (*YarnLockfile, error) {
// 	var lockfile YarnLockfile

// 	fileContents, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		return nil, fmt.Errorf("reading yarn.lock: %w", err)
// 	}

// 	lines := strings.Split(string(fileContents), "\n")

// 	r := regexp.MustCompile(`^[\w"]`)
// 	double := regexp.MustCompile(`\:\"\:`)
// 	l := regexp.MustCompile("\"|:\n$")
// 	o := regexp.MustCompile(`\"\s\"`)
// 	// deals with colons
// 	// integrity sha-... -> integrity: sha-...
// 	// "@apollo/client" latest -> "@apollo/client": latest
// 	// "@apollo/client" "0.0.0" -> "@apollo/client": "0.0.0"
// 	// apollo-client "0.0.0" -> apollo-client: "0.0.0"
// 	a := regexp.MustCompile(`(\w|\")\s(\"|\w)`)

// 	for i, line := range lines {
// 		if r.MatchString(line) {
// 			first := fmt.Sprintf("\"%v\":", l.ReplaceAllString(line, ""))
// 			lines[i] = double.ReplaceAllString(first, "\":")
// 		}
// 	}
// 	output := o.ReplaceAllString(strings.Join(lines, "\n"), "\": \"")

// 	next := a.ReplaceAllStringFunc(output, func(m string) string {
// 		parts := a.FindStringSubmatch(m)
// 		return fmt.Sprintf("%s: %s", parts[1], parts[2])
// 	})

// 	err = yaml.Unmarshal([]byte(next), &lockfile)
// 	if err != nil {
// 		return &YarnLockfile{}, fmt.Errorf("could not unmarshal lockfile: %w", err)
// 	}

// 	var prettyLockFile = YarnLockfile{}

// 	for key, val := range lockfile {
// 		if strings.Contains(key, ",") {
// 			for _, v := range strings.Split(key, ", ") {
// 				prettyLockFile[strings.TrimSpace(v)] = val
// 			}

// 		} else {
// 			prettyLockFile[key] = val
// 		}
// 	}

// 	print("prettyLockFile", prettyLockFile)

// 	return &prettyLockFile, nil
// }

var (
	yarnLocatorRegexp = regexp.MustCompile(`"?(?P<package>.+?)@(?:(?P<protocol>.+?):)?.+`)
	yarnVersionRegexp = regexp.MustCompile(`\s+"?version:?"?\s+"?(?P<version>[^"]+)"?`)
)

type Library struct {
	Name    string
	Version string
	License string `json:",omitempty"`
}
type LockFile struct {
	Dependencies map[string]Dependency
}
type Dependency struct {
	Version string
	// TODO : currently yarn can't recognize Dev flag.
	// That need to parse package.json for Dev flag
	Dev          bool
	Dependencies map[string]Dependency
}

func parsePackageLocator(target string) (packagename, protocol string, err error) {
	capture := yarnLocatorRegexp.FindStringSubmatch(target)
	if len(capture) < 2 {
		return "", "", fmt.Errorf("not package format")
	}

	fmt.Printf("\nparsePackageLocator() target: %s capture: %s", target, capture)

	for i, group := range yarnLocatorRegexp.SubexpNames() {
		switch group {
		case "package":
			packagename = capture[i]
		case "protocol":
			protocol = capture[i]
		}
	}
	return
}

func getVersion(target string) (version string, err error) {
	capture := yarnVersionRegexp.FindStringSubmatch(target)
	if len(capture) < 2 {
		return "", fmt.Errorf("not version")
	}

	return capture[len(capture)-1], nil
}

func validProtocol(protocol string) (valid bool) {
	switch protocol {
	// only scan npm packages
	case "npm", "":
		return true
	}
	return false
}

func parseYarnLockfile(file string) (*YarnLockfile, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("reading yarn.lock: %w", err)
	}

	libraries, err := readYarnLockfile(f)
	if err != nil {
		return nil, fmt.Errorf("reading yarn.lock: %w", err)
	}

	var yarnLockFile = YarnLockfile{}
	for _, dependency := range libraries {
		var packageVersion = fmt.Sprintf("%s@%s", dependency.Name, dependency.Version)
		yarnLockFile[packageVersion] = &LockfileEntry{
			Version:              dependency.Version,
			Resolved:             "",
			Integrity:            "",
			Dependencies:         nil,
			OptionalDependencies: nil,
		}
	}

	return &yarnLockFile, nil
}

func readYarnLockfile(r io.Reader) (libs []Library, err error) {
	scanner := bufio.NewScanner(r)
	unique := map[string]struct{}{}
	var lib Library
	var skipPackage bool
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 1 {
			continue
		}

		// parse version
		var version string
		if version, err = getVersion(line); err == nil {
			if skipPackage {
				continue
			}
			if lib.Name == "" {
				return nil, fmt.Errorf("invalid yarn.lock format")
			}
			// fetch between version prefix and last double-quote
			symbol := fmt.Sprintf("%s@%s", lib.Name, version)
			if _, ok := unique[symbol]; ok {
				lib = Library{}
				continue
			}

			lib.Version = version
			libs = append(libs, lib)
			lib = Library{}
			unique[symbol] = struct{}{}
			continue
		}
		// skip __metadata block
		if skipPackage = strings.HasPrefix(line, "__metadata"); skipPackage {
			continue
		}
		// packagename line start 1 char
		if line[:1] != " " && line[:1] != "#" {
			var name string
			var protocol string
			if name, protocol, err = parsePackageLocator(line); err != nil {
				continue
			}
			if skipPackage = !validProtocol(protocol); skipPackage {
				continue
			}
			lib.Name = name
		}
	}
	return libs, nil
}
