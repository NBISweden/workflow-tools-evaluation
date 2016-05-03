#!/bin/bash
echo "Downloading VCF file for Human Chromosome 17 ..."
wget http://ftp.1000genomes.ebi.ac.uk/vol1/ftp/phase1/analysis_results/integrated_call_sets/ALL.chr17.integrated_phase1_v3.20101123.snps_indels_svs.genotypes.vcf.gz
echo "Downloading DNA Fasta file for Human Chromosome 17 ..."
wget http://ftp.ensembl.org/pub/release-75//fasta/homo_sapiens/dna/Homo_sapiens.GRCh37.75.dna.chromosome.17.fa.gz
