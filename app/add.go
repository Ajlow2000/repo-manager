package app

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const GIT_SSH_PREFIX = "git@"
const GIT_HTTPS_PREFIX = "https://"

// Build a name for the local checkout of the specified url.
// Namespaces repos under the owner's username (in lower case)
func buildLocalDirName(url string) string {
    userName := ""
    projectName := ""

    if strings.HasPrefix(url, GIT_SSH_PREFIX) {     // ssh
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

func buildFQDN(urlPrefix string, repo string) string {
    repoUrl := ""

    if ( (strings.HasPrefix(repo, GIT_SSH_PREFIX) || strings.HasPrefix(repo, GIT_HTTPS_PREFIX)) && strings.HasSuffix(repo, ".git")) {
        // Ignoring urlPrefix
        repoUrl = repo
    } else {
        if (strings.HasPrefix(urlPrefix, GIT_SSH_PREFIX) || strings.HasPrefix(urlPrefix, GIT_HTTPS_PREFIX)) {
            if (strings.HasSuffix(urlPrefix, "/")) {
                repoUrl = urlPrefix + repo + ".git"
            } else {
                repoUrl = urlPrefix + "/" + repo + ".git"
            }
        } else {
    	    fmt.Fprintln(os.Stderr, "Specified value for repo is not a valid address to a git repository")
            os.Exit(1)
        }
    }

    return repoUrl
}

func Add(urlPrefix string, repo string, path string) {
    path = os.Expand(path, os.Getenv)

    repoFQDN := buildFQDN(urlPrefix, repo)
    fmt.Println("FQDN: " + repoFQDN)
    fmt.Println()

	localName := buildLocalDirName(repoFQDN)

    // Clone repo
    cloneCmd := exec.Command("git", "-C", path, "clone", repoFQDN, localName)
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
    } else {
        fmt.Println("\nRegistered new repository with zoxide")
    }

}
