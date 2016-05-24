package main

import (
	"fmt"
	"github.com/scipipe/scipipe"
)

const (
	fastq_base_url  = "http://bioinfo.perdanauniversity.edu.my/tein4ngs/ngspractice/"
	fastq_base_name = "NA06984.ILLUMINA.low_coverage.17q_%s.fq"
)

func main() {
	// Generator for the two reads
	pairsGen := scipipe.NewStringGenerator("1", "2")

	// Download FastQ component
	dlFastq := scipipe.Shell("dl_fastq", "wget -O {o:fastq} "+fastq_base_url+fmt.Sprintf(fastq_base_name, "{p:pair}"))
	dlFastq.PathFormatters["fastq"] = func(t *scipipe.SciTask) string {
		return fmt.Sprintf(fastq_base_name, t.Params["pair"])
	}

	// Sink component
	sink := scipipe.NewSink()

	// Specify data flow
	dlFastq.ParamPorts["pair"] = pairsGen.Out
	sink.In = dlFastq.OutPorts["fastq"]

	// Set up and run
	pipeRun := scipipe.NewPipelineRunner()
	pipeRun.AddProcs(pairsGen, dlFastq, sink)
	pipeRun.Run()
}
