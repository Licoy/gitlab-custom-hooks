package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type CommitType string

const (
	FEAT     CommitType = "feat"
	BUILD    CommitType = "build"
	CI       CommitType = "ci"
	FIX      CommitType = "fix"
	DOCS     CommitType = "docs"
	STYLE    CommitType = "style"
	REFACTOR CommitType = "refactor"
	TEST     CommitType = "test"
	PERF     CommitType = "perf"
)
const CommitMessagePattern = "^(?:fixup!\\s*)?(\\w*)(\\(([\\w_\u4e00-\u9fa5\\$\\.\\*/-]*)\\))?\\: (.*)|^Merge\\ branch(.*)|Merge remote-tracking branch(.*)"

const checkFailedMeassge = `##############################################################################
##                                                                          
## Commit Message 样式不符合设定规则，请检查其内容！								
##                                                                          
## Commit Message 样式应符合以下正则规则且末尾不能有符号:                      				    
##   ^(?:fixup!\s*)?(\w*)(\(([\w_\u4e00-\u9fa5\$\.\*/-].*)\ 				
##   ))?\: (.*)|^Merge\ branch(.*)|Merge remote-tracking branch(.*)         
##                                                                          
## 例如:                                                       	           
##   fix(账号): 无法接口验证码				                                
##   feat(设备模块): 增加基础设备信息											
##                                                                          
##############################################################################`

const strictMode = true

var commitMsgReg = regexp.MustCompile(CommitMessagePattern)

var symbols = []string{",", ".", "，", "。", ";", "；", "/", "?", "？"}

func main() {

	input, _ := ioutil.ReadAll(os.Stdin)
	param := strings.Fields(string(input))

	// allow branch/tag delete
	if param[1] == "0000000000000000000000000000000000000000" {
		os.Exit(0)
	}

	commitMsg := getCommitMsg(param[0], param[1])
	for _, tmpStr := range commitMsg {
		commitTypes := commitMsgReg.FindAllStringSubmatch(tmpStr, -1)

		if len(commitTypes) != 1 {
			checkFailed()
		} else {
			switch commitTypes[0][1] {
			case string(FEAT):
			case string(FIX):
			case string(DOCS):
			case string(STYLE):
			case string(REFACTOR):
			case string(TEST):
			case string(BUILD):
			case string(PERF):
			case string(CI):
			default:
				if !strings.HasPrefix(tmpStr, "Merge branch") && !strings.HasPrefix(tmpStr, "Merge remote-tracking branch") {
					checkFailed()
				}
				for _, v := range symbols {
					if strings.HasSuffix(tmpStr, v) {
						checkFailed()
					}
				}
			}
		}
		if !strictMode {
			os.Exit(0)
		}
	}

}

func getCommitMsg(odlCommitID, commitID string) []string {

	getCommitMsgCmd := exec.Command("git", "log", odlCommitID+".."+commitID, "--pretty=format:%s")

	if odlCommitID == "0000000000000000000000000000000000000000" {
		getCommitMsgCmd = exec.Command("git", "log", "-1", "--pretty=format:%s")
	}

	getCommitMsgCmd.Stdin = os.Stdin
	getCommitMsgCmd.Stderr = os.Stderr
	b, err := getCommitMsgCmd.Output()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	commitMsg := strings.Split(string(b), "\n")
	return commitMsg
}

func checkFailed() {
	_, _ = fmt.Fprintln(os.Stderr, checkFailedMeassge)
	os.Exit(1)
}
