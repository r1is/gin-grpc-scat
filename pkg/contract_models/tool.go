package models

type Tool struct {
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	LongDescription  string `json:"long_description"`
	Version          string `json:"version"`
	Reference        string `json:"reference"`
}

var tools []Tool = []Tool{ToolMythril, ToolManticore, ToolOyente, ToolHoneyBadger, ToolSecurify1, ToolSecurify2, ToolSlither, ToolConkas}
var toolsMap map[string]Tool

var ToolMythril Tool = Tool{
	Name:             "Mythril",
	ShortDescription: "Mythril is a security analysis tool for EVM bytecode.",
	LongDescription:  "Mythril is a security analysis tool for EVM bytecode. It detects security vulnerabilities in smart contracts built for Ethereum, Hedera, Quorum, Vechain, Roostock, Tron and other EVM-compatible blockchains. It uses symbolic execution, SMT solving and taint analysis to detect a variety of security vulnerabilities. It's also used (in combination with other tools and techniques) in the MythX security analysis platform.",
	Version:          "0.22.35",
	Reference:        "https://github.com/ConsenSys/mythril",
}

var ToolManticore Tool = Tool{
	Name:             "Manticore",
	ShortDescription: "Manticore is a symbolic execution tool for analysis of smart contracts and binaries.",
	LongDescription:  "Manticore is a symbolic execution tool for analysis of smart contracts and binaries.Features:\n1.Program Exploration: Manticore can execute a program with symbolic inputs and explore all the possible states it can reach\n2.Input Generation: Manticore can automatically produce concrete inputs that result in a given program state\n3.Error Discovery: Manticore can detect crashes and other failure cases in binaries and smart contracts\n4.Instrumentation: Manticore provides fine-grained control of state exploration via event callbacks and instruction hooks\n5.Programmatic Interface: Manticore exposes programmatic access to its analysis engine via a Python API\nManticore can analyze the following types of programs:\n1.Ethereum smart contracts (EVM bytecode)\n2.Linux ELF binaries (x86, x86_64, aarch64, and ARMv7)\n3.WASM Modules",
	Version:          "0.3.6",
	Reference:        "https://github.com/trailofbits/manticore",
}

var ToolOyente Tool = Tool{
	Name:             "Oyente",
	ShortDescription: "An Analysis Tool for Smart Contracts.",
	LongDescription:  "An Analysis Tool for Smart Contracts.\nThe accompanying paper explaining the bugs detected by the tool can be found here(https://www.comp.nus.edu.sg/~prateeks/papers/Oyente.pdf).",
	Version:          "0.2.7",
	Reference:        "https://github.com/enzymefinance/oyente",
}

var ToolHoneyBadger Tool = Tool{
	Name:             "HoneyBadger",
	ShortDescription: "An analysis tool based on Oyente to detect honeypots in Ethereum smart contracts.",
	LongDescription:  "An analysis tool based on Oyente to detect honeypots in Ethereum smart contracts. Our paper can be found here(https://arxiv.org/pdf/1902.06976.pdf).",
	Version:          "0.0.1",
	Reference:        "https://github.com/christoftorres/HoneyBadger",
}

var ToolSecurify1 Tool = Tool{
	Name:             "Securify1",
	ShortDescription: "Securify is a security scanner for Ethereum smart contracts supported by the Ethereum Foundation and ChainSecurity. The core research behind Securify was conducted at the ICE Center at ETH Zurich.",
	LongDescription:  "Securify is a security scanner for Ethereum smart contracts supported by the Ethereum Foundation and ChainSecurity. The core research behind Securify was conducted at the ICE Center at ETH Zurich. It features an extensive list of security patterns commonly found in smart contracts:\nsome forms of the DAO bug (also known as reentrancy)\nlocked ether\nmissing input validation\ntransaction ordering-dependent amount, receiver and transfer\nunhandled exceptions\nunrestricted ether flow",
	Version:          "0.0.1",
	Reference:        "https://github.com/eth-sri/securify",
}

var ToolSecurify2 Tool = Tool{
	Name:             "Securify2",
	ShortDescription: "Securify 2.0 is a security scanner for Ethereum smart contracts supported by the Ethereum Foundation and ChainSecurity. The core research behind Securify was conducted at the Secure, Reliable, and Intelligent Systems Lab at ETH Zurich.",
	LongDescription:  "Securify 2.0 is a security scanner for Ethereum smart contracts supported by the Ethereum Foundation and ChainSecurity. The core research behind Securify was conducted at the Secure, Reliable, and Intelligent Systems Lab at ETH Zurich.\n\nIt is the successor of the popular Securify security scanner (you can find the old version here). Features\nSupports 37 vulnerabilities (see table below)\nImplements novel context-sensitive static analysis written in Datalog\nAnalyzes contracts written in Solidity >= 0.5.8",
	Version:          "0.0.1",
	Reference:        "https://github.com/eth-sri/securify2",
}

var ToolSlither Tool = Tool{
	Name:             "Slither",
	ShortDescription: "Slither is a Solidity static analysis framework written in Python 3.",
	LongDescription:  "Slither is a Solidity static analysis framework written in Python 3. It runs a suite of vulnerability detectors, prints visual information about contract details, and provides an API to easily write custom analyses. Slither enables developers to find vulnerabilities, enhance their code comprehension, and quickly prototype custom analyses.",
	Version:          "0.8.2",
	Reference:        "https://github.com/crytic/slither",
}

var ToolConkas Tool = Tool{
	Name:             "Conkas",
	ShortDescription: "Conkas is a modular static analysis tool for Ethereum Virtual Machine (EVM) based on symbolic execution.",
	LongDescription:  "Conkas is a modular static analysis tool for Ethereum Virtual Machine (EVM) based on symbolic execution. It is capable of analysing Ethereum Smart Contracts written in Solidity or the compiled runtime bytecode. Being a modular tool means that anyone can add easily their custom modules to analyse specific vulnerabilities. It uses Z3 as the SMT Solver and Rattle as the Intermediate Representation (IR). However, to fit Conkas needs a modified version of Rattle is needed and that version can be found here. Conkas is part of my master's thesis.",
	Version:          "0.0.1",
	Reference:        "https://github.com/nveloso/conkas",
}

func init() {
	toolsMap = make(map[string]Tool)
	for _, tool := range tools {
		toolsMap[tool.Name] = tool
	}
}

func GetToolByName(name string) Tool {
	return toolsMap[name]
}

func AllTools() []Tool {
	return tools
}
