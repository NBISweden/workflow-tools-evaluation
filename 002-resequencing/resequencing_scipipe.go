package main

import (
	"fmt"
	"github.com/scipipe/scipipe"
)

const (
	fastq_base_url = "http://bioinfo.perdanauniversity.edu.my/tein4ngs/ngspractice/"
	fastq_file     = "NA06984.ILLUMINA.low_coverage.17q_%s.fq"
	ref_base_url   = "http://ftp.ensembl.org/pub/release-75/fasta/homo_sapiens/dna/"
	ref_file       = "Homo_sapiens.GRCh37.75.dna.chromosome.17.fa.gz"
	vcf_base_url   = "http://ftp.1000genomes.ebi.ac.uk/vol1/ftp/phase1/analysis_results/integrated_call_sets/"
	vcf_file       = "ALL.chr17.integrated_phase1_v3.20101123.snps_indels_svs.genotypes.vcf.gz"
)

func main() {
	// --------------------------------------------------------------------------------
	// Initialize pipeline runner
	// --------------------------------------------------------------------------------
	pipeRunner := scipipe.NewPipelineRunner()

	// --------------------------------------------------------------------------------
	// Generator for the two reads
	// --------------------------------------------------------------------------------
	pairsGen := scipipe.NewStringGenerator("1", "2")
	pipeRunner.AddProcess(pairsGen)

	// --------------------------------------------------------------------------------
	// Download Reference Genome
	// --------------------------------------------------------------------------------
	dlGzipped := scipipe.Shell("dl_gzipped", "wget -O {o:downloaded} {p:url} # {p:outfile}")
	pipeRunner.AddProcess(dlGzipped)
	// Path format
	dlGzipped.PathFormatters["downloaded"] = func(t *scipipe.SciTask) string {
		return t.Params["outfile"]
	}
	// Feed the download gzipped component with parameters
	go func() {
		defer close(dlGzipped.ParamPorts["url"])
		defer close(dlGzipped.ParamPorts["outfile"])

		dlGzipped.ParamPorts["url"] <- ref_base_url + ref_file
		dlGzipped.ParamPorts["outfile"] <- ref_file
		dlGzipped.ParamPorts["url"] <- vcf_base_url + vcf_file
		dlGzipped.ParamPorts["outfile"] <- vcf_file
	}()

	// --------------------------------------------------------------------------------
	// Unzip ref file
	// --------------------------------------------------------------------------------
	gunzip := scipipe.Shell("gunzip", "gunzip -c {i:in} > {o:out}")
	pipeRunner.AddProcess(gunzip)
	// Path format
	gunzip.SetPathFormatReplace("in", "out", ".gz", "")
	// Data flow
	gunzip.InPorts["in"] = dlGzipped.OutPorts["downloaded"]

	// --------------------------------------------------------------------------------
	// Download FastQ component
	// --------------------------------------------------------------------------------
	dlFastq := scipipe.Shell("dl_fastq", "wget -O {o:fastq} "+fastq_base_url+fmt.Sprintf(fastq_file, "{p:pair}"))
	pipeRunner.AddProcess(dlFastq)
	// Path format
	dlFastq.PathFormatters["fastq"] = func(t *scipipe.SciTask) string {
		return fmt.Sprintf(fastq_file, t.Params["pair"])
	}
	// Data flow
	dlFastq.ParamPorts["pair"] = pairsGen.Out

	// --------------------------------------------------------------------------------
	// Sink component
	// --------------------------------------------------------------------------------
	sink := scipipe.NewSink()
	pipeRunner.AddProcess(sink)
	// Data flow
	sink.InPorts["fastq"] = dlFastq.OutPorts["fastq"]
	sink.InPorts["gunzip"] = gunzip.OutPorts["out"]

	// --------------------------------------------------------------------------------
	// Run pipeline
	// --------------------------------------------------------------------------------
	pipeRunner.Run()
}
