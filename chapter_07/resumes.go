package main

import "fmt"

func pickResume(resumes []string) string {
	eliminate := "top"

	for len(resumes) > 1 {
		if eliminate == "top" {
			resumes = resumes[len(resumes)/2:]
			eliminate = "bot"
		} else if eliminate == "bot" {
			resumes = resumes[:len(resumes)/2]
			eliminate = "top"
		}
	}
	return resumes[0]
}

func pickResumeRECUR(resumes []string, eliminate string) string {
	if len(resumes) == 1 {
		return resumes[0]
	} else {
		if eliminate == "top" {
			return pickResumeRECUR(resumes[len(resumes)/2:], "bot")
		} else {
			return pickResumeRECUR(resumes[:len(resumes)/2], "top")
		}
	}
}

func main() {
	allResumes := []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg"}
	finalResume := pickResume(allResumes)
	finalResumeRECUR := pickResumeRECUR(allResumes, "top")
	fmt.Println("Final Resume:", finalResume)
	fmt.Println("RECUR Final Resume:", finalResumeRECUR)
}
