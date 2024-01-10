package main

import (
    "os"
    "log"
    "bufio"
    "strings"
)

type File struct {
    path string
    size int
    dir bool
    contents []File
}

func buildFileTree(fp string) map[string]File {
    file, err := os.Open(fp)

    fileTree := make(map[string]File)

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    currentDir := "/"

    for scanner.Scan() {
        currentLine := scanner.Text()
        splitLine := strings.Split(currentLine, " ")

        if splitLine[0] == "$" {
            if splitLine[1] == "cd" {
                currentDir = runCd(currentDir, splitLine)
                _, ok := fileTree[currentDir]
                if !ok {
                    fileTree[currentDir] = File{path: currentDir, dir: true, size: 0, contents: nil}
                }
            }
        } else {

        }
    }
}


func runCd(currentDir string, text []string) string {
    outputDir := ""
    if text[2] == ".." {
        splitDir := strings.Split(currentDir, "/")
        if len(splitDir) == 0 {
            return "/"
        }
        for i := 0; i < len(splitDir) - 1; i++ {
            outputDir += "/"
            outputDir += splitDir[i]
        }

        return outputDir
    }

    outputDir = currentDir + "/" + text[2]

    return outputDir
}

func main() {
}
