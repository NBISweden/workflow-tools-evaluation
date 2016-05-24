package main

import (
	"fmt"
	"github.com/scipipe/scipipe"
)

const (
	ref_base_url   = "http://ftp.ensembl.org/pub/release-75/fasta/homo_sapiens/dna/"
	ref_file       = "Homo_sapiens.GRCh37.75.dna.chromosome.17.fa.gz"
	fastq_base_url = "http://bioinfo.perdanauniversity.edu.my/tein4ngs/ngspractice/"
	fastq_file     = "NA06984.ILLUMINA.low_coverage.17q_%s.fq"
)

func main() {
	// Init logging
	scipipe.InitLogAudit()

	// Generator for the two reads
	pairsGen := scipipe.NewStringGenerator("1", "2")

	// Download Reference Genome
	dlRef := scipipe.Shell("dl_ref", "wget -O {o:ref} "+ref_base_url+ref_file)
	dlRef.SetPathFormatStatic("ref", ref_file)

	// Unzip ref file
	gunzip := scipipe.Shell("gunzip", "gunzip -c {i:in} > {o:out}")
	gunzip.SetPathFormatReplace("in", "out", ".gz", "")

	// Download FastQ component
	dlFastq := scipipe.Shell("dl_fastq", "wget -O {o:fastq} "+fastq_base_url+fmt.Sprintf(fastq_file, "{p:pair}"))
	dlFastq.PathFormatters["fastq"] = func(t *scipipe.SciTask) string {
		return fmt.Sprintf(fastq_file, t.Params["pair"])
	}

	// Sink component
	sink := scipipe.NewSink()

	// Specify data flow
	dlFastq.ParamPorts["pair"] = pairsGen.Out
	gunzip.InPorts["in"] = dlRef.OutPorts["ref"]

	sink.InPorts["fastq"] = dlFastq.OutPorts["fastq"]
	sink.InPorts["ref"] = gunzip.OutPorts["out"]

	// Set up and run
	pipeRun := scipipe.NewPipelineRunner()
	pipeRun.AddProcs(pairsGen, dlFastq, dlRef, gunzip, sink)
	pipeRun.Run()
}
