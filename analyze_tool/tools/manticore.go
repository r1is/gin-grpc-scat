package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	common "pkg/common"
	models "pkg/contract_models"
	"pkg/log"
	"regexp"
	"strings"
	"time"
)

type Manticore struct{}

var ToolManticore Tool = Manticore{}

func (m Manticore) Name() string {
	return models.ToolManticore.Name
}

func (m Manticore) ShortDescription() string {
	return models.ToolManticore.ShortDescription
}

func (m Manticore) LongDescription() string {
	return models.ToolManticore.LongDescription
}

func (m Manticore) Version() string {
	return models.ToolManticore.Version
}

func (m Manticore) Reference() string {
	return models.ToolManticore.Reference
}

func (m Manticore) AnalyzeByteCode(byteCode string) (models.ByteCodeAnalyzeResult, error) {
	return models.ByteCodeAnalyzeResult{}, fmt.Errorf("暂不支持")
}

func (m Manticore) AnalyzeSourceCode(sourceCode string) (models.SourceCodeAnalyzeResult, error) {
	// manticore 1.sol
	result := models.SourceCodeAnalyzeResult{}
	result.AnalyzeAt = time.Now().Unix()

	re1 := regexp.MustCompile(`Results in [/a-z0-9_A-Z.]*/mcore_[/a-z0-9_]+`) // manticore 运行输出，查找results所在路径

	//落地的文件名
	sourceFileName := common.GetRandomString() + ".sol"
	err := ioutil.WriteFile(sourceFileName, []byte(sourceCode), 0644)
	if err != nil {
		return models.SourceCodeAnalyzeResult{}, fmt.Errorf("AnalyzeSourceCode: create new solidity file error")
	}
	log.LogInfo("exec manticore")
	output, err := exec.Command("manticore", sourceFileName).Output()
	log.LogInfo("exec manticore end")
	if err != nil {
		return models.SourceCodeAnalyzeResult{}, fmt.Errorf("AnalyzeSourceCode: exec manticore error: %v", err)
	}
	resultIn := re1.FindString(string(output))
	if resultIn == "" {
		return models.SourceCodeAnalyzeResult{}, fmt.Errorf("AnalyzeSourceCode: results in ... not found, output: %s ", string(output))
	}
	dirPath := re1.FindString(string(output))[11:]
	findings, err := ioutil.ReadFile(dirPath + "/global.findings")
	if err != nil {
		return models.SourceCodeAnalyzeResult{}, fmt.Errorf("AnalyzeSourceCode: manticore result %v not found, err=%v, output=%s", dirPath+"/global.findings", err, string(output))
	}
	lines := strings.Split(string(findings), "\n")
	for i := 0; 5*i+4 < len(lines); i++ {
		vulnerability := models.Vulnerability{
			Info: phaseVulnerabilityInfoManticore(lines[i+3]),
			Type: phaseVulnerabilityTypeManticore(lines[i]),
		}
		result.Vulnerabilities = append(result.Vulnerabilities, vulnerability)
	}
	_ = os.Remove(sourceFileName)
	return result, nil
}

// 从 manticore 检测结果中提取漏洞类型
func phaseVulnerabilityTypeManticore(s string) models.VulnerabilityType {
	if strings.Contains(s, "Delegatecall") || strings.Contains(s, "delegatecall") {
		return models.VulnerabilityTypeDelegatecallToUntrustedCallee
	}
	if strings.Contains(s, "Overflow") || strings.Contains(s, "overflow") {
		return models.VulnerabilityTypeIntegerOverflowAndUnderflow
	}
	if strings.Contains(s, "Reentrancy") || strings.Contains(s, "reentrancy") {
		return models.VulnerabilityTypeReentrancy
	}
	if strings.Contains(s, "SELFDESTRUCT") {
		return models.VulnerabilityTypeUnprotectedSELFDESTRUCTInstruction
	}
	if strings.Contains(s, "initialized") {
		return models.VulnerabilityTypeUninitializedStoragePointer
	}
	if strings.Contains(s, "Returned value") {
		return models.VulnerabilityTypeUncheckedCallReturnValue
	}
	return models.CustomizeVulnerabilityType(s, models.UnknownLevel)
}

// 从 manticore 检测结果中提取漏洞位置信息
func phaseVulnerabilityInfoManticore(s string) string {
	return "line:" + s
}
