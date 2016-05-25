package main

import (
	"fmt"
	sp "github.com/scipipe/scipipe"
)

const (
	fastq_base_url = "http://bioinfo.perdanauniversity.edu.my/tein4ngs/ngspractice/"
	fastq_file     = "%s.ILLUMINA.low_coverage.4p_%s.fq"
	ref_base_url   = "http://ftp.ensembl.org/pub/release-75/fasta/homo_sapiens/dna/"
	ref_file       = "Homo_sapiens.GRCh37.75.dna.chromosome.17.fa"
	ref_file_gz    = "Homo_sapiens.GRCh37.75.dna.chromosome.17.fa.gz"
	vcf_base_url   = "http://ftp.1000genomes.ebi.ac.uk/vol1/ftp/phase1/analysis_results/integrated_call_sets/"
	vcf_file       = "ALL.chr17.integrated_phase1_v3.20101123.snps_indels_svs.genotypes.vcf.gz"
)

var (
	individuals = []string{"NA06984", "NA12489"}
	samples     = []string{"1", "2"}
)

func main() {
	//sp.InitLogDebug()
	gunzipCmdPat := "gunzip -c {i:in} > {o:out}"

	// --------------------------------------------------------------------------------
	// Initialize pipeline runner
	// --------------------------------------------------------------------------------
	pipeRun := sp.NewPipelineRunner()
	sink := sp.NewSink()

	// --------------------------------------------------------------------------------
	// Download Reference Genome
	// --------------------------------------------------------------------------------
	dlRefGz := sp.Shell("dl_gzipped",
		"wget -O {o:outfile} "+ref_base_url+ref_file_gz)
	pipeRun.AddProcess(dlRefGz)
	dlRefGz.SetPathFormatStatic("outfile", ref_file_gz)

	// --------------------------------------------------------------------------------
	// Unzip ref file
	// --------------------------------------------------------------------------------
	gunzipRef := sp.Shell("gunzipRef", gunzipCmdPat)
	pipeRun.AddProcess(gunzipRef)
	gunzipRef.SetPathFormatReplace("in", "out", ".gz", "")
	gunzipRef.InPorts["in"] = dlRefGz.OutPorts["outfile"]

	refFanOut := sp.NewFanOut()
	pipeRun.AddProcess(refFanOut)
	refFanOut.InFile = gunzipRef.OutPorts["out"]

	// --------------------------------------------------------------------------------
	// Index Reference Genome
	// --------------------------------------------------------------------------------
	idxRef := sp.Shell("Index Ref",
		"bwa index -a bwtsw {i:index}; echo done > {o:done}")
	idxRef.SetPathFormatExtend("index", "done", ".indexed")
	pipeRun.AddProcess(idxRef)
	idxRef.InPorts["index"] = refFanOut.GetOutPort("index_ref")

	idxRefDoneFanOut := sp.NewFanOut()
	pipeRun.AddProcess(idxRefDoneFanOut)
	idxRefDoneFanOut.InFile = idxRef.OutPorts["done"]

	outPorts := make(map[string]map[string]map[string]chan *sp.FileTarget)
	for _, individual := range individuals {
		outPorts[individual] = make(map[string]map[string]chan *sp.FileTarget)
		for _, sample := range samples {
			outPorts[individual][sample] = make(map[string]chan *sp.FileTarget)

			// --------------------------------------------------------------------------------
			// Download FastQ component
			// --------------------------------------------------------------------------------
			file_name := fmt.Sprintf(fastq_file, individual, sample)
			dlFastq := sp.Shell("dl_fastq",
				"wget -O {o:fastq} "+fastq_base_url+file_name)
			pipeRun.AddProcess(dlFastq)
			dlFastq.SetPathFormatStatic("fastq", file_name)
			fastQFanOut := sp.NewFanOut()
			pipeRun.AddProcess(fastQFanOut)
			fastQFanOut.InFile = dlFastq.OutPorts["fastq"]
			outPorts[individual][sample]["fastq"] = fastQFanOut.GetOutPort("merg")

			// --------------------------------------------------------------------------------
			// BWA Align
			// --------------------------------------------------------------------------------
			bwaAln := sp.Shell("bwa_aln",
				"bwa aln {i:ref} {i:fastq} > {o:sai} # {i:index_done}")
			pipeRun.AddProcess(bwaAln)
			bwaAln.SetPathFormatExtend("fastq", "sai", ".sai")
			// Connect
			bwaAln.InPorts["ref"] = refFanOut.GetOutPort("bwa_aln_" + individual + "_" + sample)
			bwaAln.InPorts["index_done"] = idxRefDoneFanOut.GetOutPort("bwa_aln_" + individual + "_" + sample)
			bwaAln.InPorts["fastq"] = fastQFanOut.GetOutPort("bwa_aln")
			// Store in map
			outPorts[individual][sample]["sai"] = bwaAln.OutPorts["sai"]
		}

		// --------------------------------------------------------------------------------
		// Merge
		// --------------------------------------------------------------------------------
		individualParamGen := sp.NewStringGenerator(individual)
		pipeRun.AddProcess(individualParamGen)

		merg := sp.Shell("merge_"+individual,
			"bwa sampe {i:ref} {i:sai1} {i:sai2} {i:fq1} {i:fq2} > {o:merged} # {p:individual}")
		pipeRun.AddProcess(merg)
		merg.PathFormatters["merged"] = func(t *sp.SciTask) string {
			return fmt.Sprintf("%s.merged.sam", t.Params["individual"])
		}
		merg.InPorts["ref"] = refFanOut.GetOutPort("merg_" + individual)
		merg.InPorts["index_done"] = idxRefDoneFanOut.GetOutPort("merg_" + individual)
		merg.InPorts["sai1"] = outPorts[individual]["1"]["sai"]
		merg.InPorts["sai2"] = outPorts[individual]["2"]["sai"]
		merg.InPorts["fq1"] = outPorts[individual]["1"]["fastq"]
		merg.InPorts["fq2"] = outPorts[individual]["2"]["fastq"]
		merg.ParamPorts["individual"] = individualParamGen.Out

		sink.InPorts[individual] = merg.OutPorts["merged"]
	}
	// --------------------------------------------------------------------------------
	// Run pipeline
	// --------------------------------------------------------------------------------
	pipeRun.AddProcess(sink)
	pipeRun.Run()
}
