# Resequencing analysis example

The idea here is to implement the [Resequencing analysis
pipeline](http://uppnex.se/twiki/do/view/Courses/NgsIntro1502/ResequencingAnalysis.html)
from Scilifelab's NGS intro course, as has [partly already been done in other
tools](https://gist.github.com/samuell/6da9a7c1e03912fde62e).

Note that there are a number of pre-existing community-made tool definitions in
[this repository](https://github.com/common-workflow-language/workflows/tree/master/tools)
that could most probably be re-used in this demo.

## Tools included

* `bwa aln` ([CWL Tool](https://github.com/common-workflow-language/workflows/blob/master/tools/bwa-aln.cwl))
* `bwa index` ([CWL Tool](https://github.com/common-workflow-language/workflows/blob/master/tools/bwa-index.cwl))
* `bwa sampe`
* `Picard/AddOrReplaceReadGroups.jar`
* `Picard/BuildBamIndex.jar` ([CWL Tool](https://github.com/BILS/workflows/blob/master/tools/picard-BuildBamIndex.cwl))
* `Picard/MarkDuplicates.jar` ([CWL Tool](https://github.com/BILS/workflows/blob/master/tools/picard-MarkDuplicates.cwl))
* `GATK/GenomeAnalysisTK.jar -T CountCovariates`
* `GATK/GenomeAnalysisTK.jar -T IndelRealigner` ([CWL Tool](https://github.com/common-workflow-language/workflows/blob/master/tools/GATK-RealignTargetCreator.cwl))
* `GATK/GenomeAnalysisTK.jar -T RealignerTargetCreator` ([CWL Tool](https://github.com/common-workflow-language/workflows/blob/master/tools/GATK-RealignTargetCreator.cwl))
* `GATK/GenomeAnalysisTK.jar -T TableRecalibration`

### How to install on Ubuntu

#### BWA

```bash
sudo apt-get install bwa
```

#### Picard

```bash
sudo apt-get install picard-tools
```

#### GATK
See [this page](https://www.broadinstitute.org/gatk/download/) for how to download.

## Steps included

1.  Create a BAM index
2.  Mapping - Making Single Read Alignments for each of the reads in the paired end data
3.  Merging Alignments and Making SAM Files
4.  Creating a BAM File
5.  Processing Reads with GATK
6.  Variant Calling
7.  Looking at Your Data with IGV

## Raw data needed:

* [ALL.chr17.integrated_phase1_v3.20101123.snps_indels_svs.genotypes.vcf.gz](http://ftp.1000genomes.ebi.ac.uk/vol1/ftp/phase1/analysis_results/integrated_call_sets/ALL.chr17.integrated_phase1_v3.20101123.snps_indels_svs.genotypes.vcf.gz)
* [Homo_sapiens.GRCh37.75.dna.chromosome.17.fa.gz](http://ftp.ensembl.org/pub/release-75//fasta/homo_sapiens/dna/Homo_sapiens.GRCh37.75.dna.chromosome.17.fa.gz)
* [NA06984.ILLUMINA.low_coverage.17q_1.fq](http://bioinfo.perdanauniversity.edu.my/tein4ngs/ngspractice/NA06984.ILLUMINA.low_coverage.17q_1.fq)
