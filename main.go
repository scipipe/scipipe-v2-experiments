package main

import (
	"github.com/flowbase/flowbase"
)

type PipelineNode interface {
	flowbase.Node
	// Things to add to flowbase.Node
	UpstreamProcs() []PipelineNode
	DownstreamProcs() []PipelineNode

	// New things specific to PipelineNode
	StageInputs() error
	OutputsExist() (bool, error)
	Persist() error
}

type PipelinePacket interface {
	flowbase.Packet
	ID() []string
	GetURI() []string
	GetRaw() ([]byte, err)
	GetScanner()
	GetTextScanner()
	Tags() [string]string
	GetFlowContext() FlowContext
}

func main() {
	net := flowbase.NewNetwork()

	net.AddProc()

	net.Run()
}
