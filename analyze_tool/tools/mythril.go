package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	common "pkg/common"
	models "pkg/contract_models"
	"time"
)

type Mythril struct{}

var ToolMythril Tool = Mythril{}

func (m Mythril) Name() string {
	return models.ToolMythril.Name
}

func (m Mythril) ShortDescription() string {
	return models.ToolMythril.ShortDescription
}

func (m Mythril) LongDescription() string {
	return models.ToolMythril.LongDescription
}

func (m Mythril) Version() string {
	return models.ToolMythril.Version
}

func (m Mythril) Reference() string {
	return models.ToolMythril.Reference
}

func (m Mythril) AnalyzeByteCode(byteCode string) (models.ByteCodeAnalyzeResult, error) {
	// myth analyze -f 1.bytecode -o json
	result := models.ByteCodeAnalyzeResult{}
	result.AnalyzeAt = time.Now().Unix()

	type JsonStruct struct {
		Issues []struct {
			Function string `json:"function"`
			SwcID    string `json:"swc-id"`
		} `json:"issues"`
	}
	sourceFileName := common.GetRandomString() + ".sol"
	err := ioutil.WriteFile(sourceFileName, []byte(byteCode), 0644)
	if err != nil {
		return models.ByteCodeAnalyzeResult{}, fmt.Errorf("AnalyzeSourceCode: create new solidity file error")
	}
	output, err := exec.Command("myth", "analyze", "-f", "1.bytecode", "-o", "json").Output()
	if err != nil {
		return models.ByteCodeAnalyzeResult{}, fmt.Errorf("AnalyzeByteCode: exec mythril error")
	}
	var resultJson JsonStruct
	err = json.Unmarshal(output, &resultJson)
	if err != nil {
		return models.ByteCodeAnalyzeResult{}, fmt.Errorf("AnalyzeByteCode: result json error")
	}

	for _, issue := range resultJson.Issues {
		vulnerability := models.Vulnerability{
			Info: fmt.Sprintf("function: %s", issue.Function),
			Type: phaseVulnerabilityTypeMythril(issue.SwcID),
		}
		result.Vulnerabilities = append(result.Vulnerabilities, vulnerability)
	}
	_ = os.Remove(sourceFileName)
	return result, nil
}

func (m Mythril) AnalyzeSourceCode(sourceCode string) (models.SourceCodeAnalyzeResult, error) {
	// myth analyze 1.sol -o json --solv 0.4.25
	result := models.SourceCodeAnalyzeResult{}
	result.AnalyzeAt = time.Now().Unix()

	type JsonStruct struct {
		Issues []struct {
			Code     string `json:"code"`
			Function string `json:"function"`
			Lineno   int    `json:"lineno"`
			SwcID    string `json:"swc-id"`
		} `json:"issues"`
	}

	//落地的文件名
	sourceFileName := common.GetRandomString() + ".sol"
	err := ioutil.WriteFile(sourceFileName, []byte(sourceCode), 0644)
	if err != nil {
		return models.SourceCodeAnalyzeResult{}, fmt.Errorf("AnalyzeSourceCode: create new solidity file error")
	}
	output, _ := exec.Command("myth", "analyze", sourceFileName, "-o", "json").Output()
	//调试POST 传入参数是否能正常调用myth分析合约
	fmt.Println("分析结果: ", string(output), "err:", err)
	// if err != nil {
	// 	return models.SourceCodeAnalyzeResult{}, fmt.Errorf("AnalyzeSourceCode: exec mythril error")
	// }
	var resultJson JsonStruct
	err = json.Unmarshal(output, &resultJson)
	if err != nil {
		return models.SourceCodeAnalyzeResult{}, fmt.Errorf("AnalyzeSourceCode: result json error")
	}

	for _, issue := range resultJson.Issues {
		vulnerability := models.Vulnerability{
			Info: fmt.Sprintf("line: %v, function: %s, code: %s", issue.Lineno, issue.Function, issue.Code),
			Type: phaseVulnerabilityTypeMythril(issue.SwcID),
		}
		result.Vulnerabilities = append(result.Vulnerabilities, vulnerability)
	}
	_ = os.Remove(sourceFileName)
	return result, nil
}

func phaseVulnerabilityTypeMythril(s string) models.VulnerabilityType {
	switch s {
	case "100":
		return models.VulnerabilityTypeFunctionDefaultVisibility
	case "101":
		return models.VulnerabilityTypeIntegerOverflowAndUnderflow
	case "102":
		return models.VulnerabilityTypeOutdatedCompilerVersion
	case "103":
		return models.VulnerabilityTypeFloatingPragma
	case "104":
		return models.VulnerabilityTypeUncheckedCallReturnValue
	case "105":
		return models.VulnerabilityTypeUnprotectedEtherWithdrawal
	case "106":
		return models.VulnerabilityTypeUnprotectedSELFDESTRUCTInstruction
	case "107":
		return models.VulnerabilityTypeReentrancy
	case "108":
		return models.VulnerabilityTypeStateVariableDefaultVisibility
	case "109":
		return models.VulnerabilityTypeUninitializedStoragePointer
	case "110":
		return models.VulnerabilityTypeAssertViolation
	case "111":
		return models.VulnerabilityTypeUseOfDeprecatedSolidityFunctions
	case "112":
		return models.VulnerabilityTypeDelegatecallToUntrustedCallee
	case "113":
		return models.VulnerabilityTypeDoSWithFailedCall
	case "114":
		return models.VulnerabilityTypeTransactionOrderDependence
	case "115":
		return models.VulnerabilityTypeAuthorizationThroughTxOrigin
	case "116":
		return models.VulnerabilityTypeBlockValuesAsAProxyForTime
	case "117":
		return models.VulnerabilityTypeSignatureMalleability
	case "118":
		return models.VulnerabilityTypeIncorrectConstructorName
	case "119":
		return models.VulnerabilityTypeShadowingStateVariables
	case "120":
		return models.VulnerabilityTypeWeakSourcesOfRandomnessFromChainAttributes
	case "121":
		return models.VulnerabilityTypeMissingProtectionAgainstSignatureReplayAttacks
	case "122":
		return models.VulnerabilityTypeLackOfProperSignatureVerification
	case "123":
		return models.VulnerabilityTypeRequirementViolation
	case "124":
		return models.VulnerabilityTypeWriteToArbitraryStorageLocation
	case "125":
		return models.VulnerabilityTypeIncorrectInheritanceOrder
	case "126":
		return models.VulnerabilityTypeInsufficientGasGriefing
	case "127":
		return models.VulnerabilityTypeArbitraryJumpWithFunctionTypeVariable
	case "128":
		return models.VulnerabilityTypeDoSWithBlockGasLimit
	case "129":
		return models.VulnerabilityTypeTypographicalError
	case "130":
		return models.VulnerabilityTypeRightToLeftOverrideControlCharacter
	case "131":
		return models.VulnerabilityTypePresenceOfUnusedVariables
	case "132":
		return models.VulnerabilityTypeUnexpectedEtherBalance
	case "133":
		return models.VulnerabilityTypeHashCollisionsWithMultipleVariableLengthArguments
	case "134":
		return models.VulnerabilityTypeMessageCallWithHardcodedGasAmount
	case "135":
		return models.VulnerabilityTypeCodeWithNoEffects
	case "136":
		return models.VulnerabilityTypeUnencryptedPrivateDataOnChain
	default:
		return models.VulnerabilityTypeUnknown
	}
}
