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
	GetRaw() ([]byte, error)
	GetScanner()
	GetTextScanner()
	Write([]byte) error
	Persist() error
	GetFlowContext() FlowContext // Should this be moved to flowbase instead?
	Tags() [string]string        // Should this be kept in the FlowContext instead?
}

func main() {
	net := flowbase.NewNetwork()

	net.AddProc()

	net.Run()
}
