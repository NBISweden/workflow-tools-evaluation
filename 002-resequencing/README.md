# Resequencing analysis example

The idea here is to implement the [Resequencing analysis
pipeline](http://uppnex.se/twiki/do/view/Courses/NgsIntro1502/ResequencingAnalysis.html)
from Scilifelab's NGS intro course, as has [partly already been done in other
tools](https://gist.github.com/samuell/6da9a7c1e03912fde62e).

Note that there are a number of pre-existing community-made tool definitions in
[this repository](https://github.com/common-workflow-language/workflows/tree/master/tools)
that could most probably be re-used in this demo.

## Tools included

* `bwa index`
* `bwa aln`
* `bwa sampe`
* `Picard/AddOrReplaceReadGroups.jar`
* `Picard/BuildBamIndex.jar`
* `Picard/MarkDuplicates.jar`
* `GATK/GenomeAnalysisTK.jar -T RealignerTargetCreator`
* `GATK/GenomeAnalysisTK.jar -T IndelRealigner`
* `GATK/GenomeAnalysisTK.jar -T CountCovariates`
* `GATK/GenomeAnalysisTK.jar -T TableRecalibration`

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
* [Homo_sapiens.GRCh37.75.dna.chromosome.17.fa.gz](ftp://ftp.ensembl.org/pub/release-75//fasta/homo_sapiens/dna/Homo_sapiens.GRCh37.75.dna.chromosome.17.fa.gz)
