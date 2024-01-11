package main

import (
    "os"
    "log"
    "bufio"
    "strings"
    "strconv"
    "fmt"
)


const MAX_FOLDER_SIZE int = 100000

type File struct {
    path string
    size int
    dir bool
    contents []*File
}

func buildFileTree(fp string) map[string]*File {
    file, err := os.Open(fp)

    fileTree := make(map[string]*File)

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    currentDir := "/"

    for scanner.Scan() {
        currentLine := scanner.Text()
        splitLine := strings.Split(currentLine, " ")

        if splitLine[0] == "$" { // Is a command
            if splitLine[1] == "cd" {
                currentDir = runCd(currentDir, splitLine)
                _, ok := fileTree[currentDir]

                if !ok {
                    fileTree[currentDir] = &File{path: currentDir, dir: true, size: 0, contents: nil}
                }
            }

        } else { // Is Output
            newFilePath := currentDir + splitLine[1]
            if splitLine[0] == "dir" {
                newFilePath := newFilePath + "/"
                _, ok := fileTree[newFilePath]

                if !ok {
                    fileTree[newFilePath] = &File{path: newFilePath, dir: true, size: 0, contents: nil}
                    currentDirFile := fileTree[currentDir]
                    fmt.Printf("Appending to contents: %s\n", fileTree[newFilePath].path)
                    currentDirFile.contents = append(currentDirFile.contents, fileTree[newFilePath])
                    fmt.Printf("Contents length now for %s: %d\n", currentDirFile.path, len(currentDirFile.contents))
                    fileTree[currentDirFile.path] = currentDirFile
                }
            } else { // File
                size, err := strconv.Atoi(splitLine[0])
                if err != nil {
                    log.Fatal(err)
                }
                fileTree[newFilePath] = &File{path: newFilePath, dir: false, size: size, contents: nil}
                currentDirFile := fileTree[currentDir]
                fmt.Printf("Appending to contents: %s\n", fileTree[newFilePath].path)
                currentDirFile.contents = append(currentDirFile.contents, fileTree[newFilePath])
                fmt.Printf("Contents lenght now for %s: %d\n", currentDirFile.path, len(currentDirFile.contents))
                fileTree[currentDirFile.path] = currentDirFile
            }
        }
    }
    return fileTree
}

func runCd(currentDir string, text []string) string {
    fmt.Printf("Running runCd with at %s with command %s\n", currentDir, text)
    outputDir := ""
    if text[2] == "/" {
        fmt.Printf("runCd: file is %s, going to /\n", text[2])
        return "/"
    }

    if text[2] == ".." {
        fmt.Printf("runCd: Going up a dir\n")
        splitDir := strings.Split(currentDir, "/")
        if len(splitDir) == 0 {
            fmt.Printf("runCd: Going up a dir, no dirs left, returning /\n")
            return "/"
        }
        for i := 0; i < len(splitDir) - 2; i++ {
            outputDir += splitDir[i]
            outputDir += "/"
        }

        fmt.Printf("runCd: Going up a dir, returning %s\n", outputDir)
        return outputDir
    }

    outputDir = currentDir + text[2] + "/"
    fmt.Printf("runCd: cd into %s. New path: %s\n", text[2], outputDir)

    return outputDir
}

func printTree(tree map[string]*File) {
    for key, value := range tree {
        if value.dir {
            fmt.Printf("%s: %s Contents: ", key, value.path)
            for _, content := range value.contents {
                fmt.Printf("%s\t", content.path)
            }
            fmt.Printf("\n")
        } else {
            fmt.Printf("%s: %s size: %d\n", key, value.path, value.size)
        }

    }
    fmt.Printf("%s size: %d\n", tree["/"].path, tree["/"].size)
    _printTree(tree, tree["/"], "  ")
}
 
func _printTree(tree map[string]*File, currentDir *File, level string) {
    for _, file := range currentDir.contents {
        if file.dir {
            fmt.Printf("%s%s size: %d\n", level, file.path, file.size)
            _printTree(tree, file, level + "  ")
        } else {
            fmt.Printf("%s%s -- %d\n", level, file.path, file.size)
        }
    }
}

func calculateSizes(root *File) {
    sum := 0

    for _, file := range root.contents {
        if file.dir {
            calculateSizes(file)
        }
        sum += file.size
    }

    root.size = sum
}

func findSumOfRightSizes(tree map[string]*File) int {
    sum := 0
    for _, file := range tree {
        if file.dir {
            if file.size < MAX_FOLDER_SIZE {
                sum += file.size
            }
        }
    }

    return sum
}

func main() {
    fileTree := buildFileTree("input.txt")
    calculateSizes(fileTree["/"])
    printTree(fileTree)
    fmt.Printf("%d\n", findSumOfRightSizes(fileTree))
}
