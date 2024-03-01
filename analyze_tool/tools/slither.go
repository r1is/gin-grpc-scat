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

type Slither struct{}

var ToolSlither Tool = Slither{}

func (s Slither) Name() string {
	return models.ToolSlither.Name
}

func (s Slither) ShortDescription() string {
	return models.ToolSlither.ShortDescription
}

func (s Slither) LongDescription() string {
	return models.ToolSlither.LongDescription
}

func (s Slither) Version() string {
	return models.ToolSlither.Version
}

func (s Slither) Reference() string {
	return models.ToolSlither.Reference
}

func (s Slither) AnalyzeByteCode(byteCode string) (models.ByteCodeAnalyzeResult, error) {
	return models.ByteCodeAnalyzeResult{}, fmt.Errorf("AnalyzeByteCode error: slither do not support bytecode input")
}

func (s Slither) AnalyzeSourceCode(sourceCode string) (models.SourceCodeAnalyzeResult, error) {
	// slither 1.sol --exclude-informational --exclude-low --json -
	result := models.SourceCodeAnalyzeResult{}
	result.AnalyzeAt = time.Now().Unix()

	type JsonStruct struct {
		Success bool        `json:"success"`
		Error   interface{} `json:"error"`
		Results struct {
			Detectors []struct {
				Description string `json:"description"`
				Check       string `json:"check"`
			} `json:"detectors"`
		} `json:"results"`
	}
	//落地的文件名
	sourceFileName := common.GetRandomString() + ".sol"
	err := ioutil.WriteFile(sourceFileName, []byte(sourceCode), 0644)
	if err != nil {
		return models.SourceCodeAnalyzeResult{}, fmt.Errorf("AnalyzeSourceCode: create new solidity file error")
	}
	output, err := exec.Command("slither", sourceFileName, "--exclude-informational", "--exclude-low", "--json", "-").Output()
	if err != nil {
		return models.SourceCodeAnalyzeResult{}, fmt.Errorf("AnalyzeSourceCode: exec slither error")
	}
	var resultJson JsonStruct
	err = json.Unmarshal(output, &resultJson)
	if err != nil {
		return models.SourceCodeAnalyzeResult{}, fmt.Errorf("AnalyzeSourceCode: result json error")
	}
	for _, issue := range resultJson.Results.Detectors {
		vulnerability := models.Vulnerability{
			Info: issue.Description,
			Type: phaseVulnerabilityTypeSlither(issue.Check),
		}
		result.Vulnerabilities = append(result.Vulnerabilities, vulnerability)
	}
	_ = os.Remove(sourceFileName)
	return result, nil
}

func phaseVulnerabilityTypeSlither(s string) models.VulnerabilityType {
	switch s {
	case "abiencoderv2-array":
		return models.VulnerabilityTypeABIEncoderV2Array
	case "array-by-reference":
		return models.VulnerabilityTypeArrayByReference
	case "incorrect-shift":
		return models.VulnerabilityTypeIncorrectShift
	case "multiple-constructors":
		return models.VulnerabilityTypeMultipleConstructors
	case "name-reused":
		return models.VulnerabilityTypeNameReused
	case "public-mappings-nested":
		return models.VulnerabilityTypePublicMappingsNested
	case "rtlo":
		return models.VulnerabilityTypeRightToLeftOverrideControlCharacter
	case "shadowing-state":
		return models.VulnerabilityTypeShadowingStateVariables
	case "suicidal":
		return models.VulnerabilityTypeUnprotectedSELFDESTRUCTInstruction
	case "uninitialized-state":
		return models.VulnerabilityTypeUninitializedStoragePointer
	case "uninitialized-storage":
		return models.VulnerabilityTypeUninitializedStoragePointer
	case "unprotected-upgrade":
		return models.VulnerabilityTypeUnprotectedUpgrade
	case "arbitrary-send":
		return models.VulnerabilityTypeUnprotectedEtherWithdrawal
	case "controlled-array-length":
		return models.VulnerabilityTypeArrayLengthAssignment
	case "controlled-delegatecall":
		return models.VulnerabilityTypeDelegatecallToUntrustedCallee
	case "delegatecall-loop":
		return models.VulnerabilityTypeReentrancy
	case "msg-value-loop":
		return models.VulnerabilityTypeMsgValueInsideLoop
	case "reentrancy-eth":
		return models.VulnerabilityTypeReentrancy
	case "storage-array":
		return models.VulnerabilityTypeStorageSignedIntegerArray
	case "unchecked-transfer":
		return models.VulnerabilityTypeUncheckedCallReturnValue
	case "weak-prng":
		return models.VulnerabilityTypeWeakSourcesOfRandomnessFromChainAttributes
	case "enum-conversion":
		return models.VulnerabilityTypeDangerousEnumConversion
	case "erc20-interface":
		return models.VulnerabilityTypeIncorrectERC20Interface
	case "erc721-interface":
		return models.VulnerabilityTypeIncorrectERC721Interface
	case "incorrect-equality":
		return models.VulnerabilityTypeDangerousStrictEqualities
	case "locked-ether":
		return models.VulnerabilityTypeContractsThatLockEther
	case "mapping-deletion":
		return models.VulnerabilityTypeMappingDeletion
	case "shadowing-abstract":
		return models.VulnerabilityTypeShadowingStateVariables
	case "tautology":
		return models.VulnerabilityTypeTautologyOrContradiction
	case "write-after-write":
		return models.VulnerabilityTypeWriteAfterWrite
	case "boolean-cst":
		return models.VulnerabilityTypeMisuseOfBooleanConstant
	case "constant-function-asm":
		return models.VulnerabilityTypeConstantFunctionAsm
	case "constant-function-state":
		return models.VulnerabilityTypeConstantFunctionsState
	case "divide-before-multiply":
		return models.VulnerabilityTypeDivideBeforeMultiply
	case "reentrancy-no-eth":
		return models.VulnerabilityTypeReentrancy
	case "reused-constructor":
		return models.VulnerabilityTypeReusedBaseConstructors
	case "tx-origin":
		return models.VulnerabilityTypeAuthorizationThroughTxOrigin
	case "unchecked-lowlevel":
		return models.VulnerabilityTypeUncheckedCallReturnValue
	case "unchecked-send":
		return models.VulnerabilityTypeUncheckedCallReturnValue
	case "uninitialized-local":
		return models.VulnerabilityTypeUninitializedStoragePointer
	case "unused-return":
		return models.VulnerabilityTypeUnusedReturn
	default:
		return models.CustomizeVulnerabilityType(s, models.UnknownLevel)
	}
}
