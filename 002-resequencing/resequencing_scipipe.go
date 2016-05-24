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
	dlFastq := scipipe.Shell("dl_fastq", "wget -O {o:fastq} "+fastq_base_url+fmt.Sprintf(fastq_base_name, "{p:pair}"))
	dlFastq.PathFormatters["fastq"] = func(t *scipipe.SciTask) string {
		return fmt.Sprintf(fastq_base_name, t.Params["pair"])
	}
	sink := scipipe.NewSink()

	sink.In = dlFastq.OutPorts["fastq"]

	go func() {
		defer close(dlFastq.ParamPorts["pair"])
		for i := 1; i <= 2; i++ {
			dlFastq.ParamPorts["pair"] <- fmt.Sprintf("%d", i)
		}
	}()

	pipeRun := scipipe.NewPipelineRunner()
	pipeRun.AddProcs(dlFastq, sink)
	pipeRun.Run()
}
