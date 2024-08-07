package app

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Build a name for the local checkout of the specified url.
// Namespaces repos under the owner's username (in lower case)
func buildLocalDirName(url string) string {
    userName := ""
    projectName := ""

    if strings.HasPrefix(url, "git@") {     // ssh
        url = strings.Split(url, ":")[1]
        params := strings.Split(url, "/")
        userName = params[0]
        projectName = params[1]
    } else {                                // https
        params := strings.Split(url, "/")
        userName = params[3]
        projectName = params[4]
    }

    userName = strings.ToLower(userName)
    projectName, _ = strings.CutSuffix(projectName, ".git")

    return userName + "_" + projectName;
}

func Main(url string, path string) {
    path = os.Expand(path, os.Getenv)

	localName := buildLocalDirName(url)

    // Clone repo
    cloneCmd := exec.Command("git", "-C", path, "clone", url, localName)
    cloneCmd.Stdout = os.Stdout
    cloneCmd.Stderr = os.Stderr
    err := cloneCmd.Run()
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        return
    }

    // Build destination with proper / divider for linux
    destination := path
    if !strings.HasSuffix(destination, "/") {
        destination = destination + "/"
    }
    destination = destination + localName

    // Add .envrc
    f, err := os.Create(destination + "/" + ".envrc")
	if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	_, err = f.WriteString("use flake")
	if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        f.Close()
		return
	}
	err = f.Close()
	if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
		return
	}

    // Register with zoxide
    registerCmd := exec.Command("zoxide", "add", destination)
    registerCmd.Stdout = os.Stdout
    registerCmd.Stderr = os.Stderr
    err = registerCmd.Run()
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        return
    }

}
