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
* `GATK/GenomeAnalysisTK.jar -T CountCovariates`
* `GATK/GenomeAnalysisTK.jar -T IndelRealigner` ([CWL Tool](https://github.com/common-workflow-language/workflows/blob/master/tools/GATK-RealignTargetCreator.cwl))
* `GATK/GenomeAnalysisTK.jar -T RealignerTargetCreator` ([CWL Tool](https://github.com/common-workflow-language/workflows/blob/master/tools/GATK-RealignTargetCreator.cwl))
* `GATK/GenomeAnalysisTK.jar -T TableRecalibration`
* `GATK/GenomeAnalysisTK.jar -T UnifiedGenotyper`
* `GATK/GenomeAnalysisTK.jar -T VariantFiltration`
* `Picard/AddOrReplaceReadGroups.jar`
* `Picard/BuildBamIndex.jar` ([CWL Tool](https://github.com/BILS/workflows/blob/master/tools/picard-BuildBamIndex.cwl))
* `Picard/MarkDuplicates.jar` ([CWL Tool](https://github.com/BILS/workflows/blob/master/tools/picard-MarkDuplicates.cwl))
* `Picard/MergeSamFiles.jar`
* `samtools faidx` ([CWL Tool](https://github.com/common-workflow-language/workflows/blob/master/tools/samtools-faidx.cwl))

### How to install on Ubuntu

#### BWA

```bash
sudo apt-get install bwa
```

#### Picard

```bash
sudo apt-get install picard-tools
```

#### Toil

Install the version of CWLTool that Toil depends on:
```bash
pip uninstall cwltool
pip install cwltool==1.0.20160425140546
```

Install toil:
```
git clone https://github.com/BD2KGenomics/toil.git
cd toil
pip install .
```

(Then, use the `cwltoil` tool, as a replacement for `cwl-runner` or `cwltool`)

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

## Outstanding issues

* [Provide a means to calculate the relative directory of input paths](https://github.com/common-workflow-language/common-workflow-language/issues/213)

## Current Status

Tried running the workflow ending with the bwa aln step, but is running into strange issues:

```bash
[samuel 002-resequencing]$ python --version
Python 2.7.11
[samuel 002-resequencing]$ cwltool --version
/home/samuel/.pyenv/versions/2.7.11/bin/cwltool 1.0.20160425140546
[samuel 002-resequencing]$ cwltoil --version
3.2.0a2
[samuel 002-resequencing]$ ./run_resequencing_workflow_testjob.sh 
2016-05-23 17:30:24,586 INFO:toil.lib.bioio: Logging set at level: INFO
2016-05-23 17:30:36,024 INFO:toil.lib.bioio: Logging set at level: INFO
2016-05-23 17:30:36,025 INFO:toil.jobStores.fileJobStore: Jobstore directory is: /tmp/tmpnhlStz
2016-05-23 17:30:36,026 INFO:toil.jobStores.abstractJobStore: The workflow ID is: 'ddb828e0-cc6d-48db-963a-952d4c586938'
2016-05-23 17:30:36,029 INFO:toil.jobStores.abstractJobStore: Unable to import 'toil.jobStores.azureJobStore'
2016-05-23 17:30:36,030 INFO:toil.jobStores.abstractJobStore: Unable to import 'toil.jobStores.aws.jobStore'
2016-05-23 17:30:36,115 INFO:toil.common: Using the single machine batch system
2016-05-23 17:30:36,115 WARNING:toil.batchSystems.singleMachine: Limiting maxCores to CPU count of system (4).
2016-05-23 17:30:36,115 WARNING:toil.batchSystems.singleMachine: Limiting maxMemory to physically available memory (8277901312).
2016-05-23 17:30:36,116 INFO:toil.batchSystems.singleMachine: Setting up the thread pool with 40 workers, given a minimum CPU fraction of 0.100000 and a maximum CPU value of 4.
2016-05-23 17:30:36,121 INFO:toil.common: Written the environment for the jobs to the environment file
2016-05-23 17:30:36,121 INFO:toil.common: Caching all jobs in job store
2016-05-23 17:30:36,122 INFO:toil.common: 0 jobs downloaded.
2016-05-23 17:30:36,543 INFO:toil.realtimeLogger: Real-time logging disabled
2016-05-23 17:30:36,549 INFO:toil.leader: (Re)building internal scheduler state
2016-05-23 17:30:36,550 INFO:toil.leader: Checked batch system has no running jobs and no updated jobs
2016-05-23 17:30:36,550 INFO:toil.leader: Found 1 jobs to start and 0 jobs with successors to run
2016-05-23 17:30:36,555 INFO:toil.leader: Starting the main loop
2016-05-23 17:30:36,558 INFO:toil.batchSystems.singleMachine: Executing command: '_toil_worker /tmp/tmpnhlStz p/5/jobvF8POb'.
INFO:toil.common:Created the workflow directory at /tmp/toil-ddb828e0-cc6d-48db-963a-952d4c586938
2016-05-23 17:30:37,425 INFO:toil.batchSystems.singleMachine: Executing command: '_toil_worker /tmp/tmpnhlStz b/8/jobuYbpXG'.
2016-05-23 17:30:37,425 INFO:toil.batchSystems.singleMachine: Executing command: '_toil_worker /tmp/tmpnhlStz Y/M/jobr5Bzt8'.
2016-05-23 17:30:37,425 INFO:toil.batchSystems.singleMachine: Executing command: '_toil_worker /tmp/tmpnhlStz g/V/jobqUg36A'.
Exception in thread Thread-1:
Traceback (most recent call last):
  File "/home/samuel/.pyenv/versions/2.7.11/lib/python2.7/threading.py", line 801, in __bootstrap_inner
    self.run()
  File "/home/samuel/.pyenv/versions/2.7.11/lib/python2.7/threading.py", line 754, in run
    self.__target(*self.__args, **self.__kwargs)
  File "/home/samuel/.pyenv/versions/2.7.11/lib/python2.7/site-packages/toil/job.py", line 528, in asyncWrite
    raise RuntimeError("The termination flag is set, exiting")
RuntimeError: The termination flag is set, exiting
Exception in thread Thread-2:
Traceback (most recent call last):
  File "/home/samuel/.pyenv/versions/2.7.11/lib/python2.7/threading.py", line 801, in __bootstrap_inner
    self.run()
  File "/home/samuel/.pyenv/versions/2.7.11/lib/python2.7/threading.py", line 754, in run
    self.__target(*self.__args, **self.__kwargs)
  File "/home/samuel/.pyenv/versions/2.7.11/lib/python2.7/site-packages/toil/job.py", line 528, in asyncWrite
    raise RuntimeError("The termination flag is set, exiting")
RuntimeError: The termination flag is set, exiting


Exception RuntimeError: RuntimeError('cannot join current thread',) in <bound method FileStore.__del__ of <toil.job.FileStore object at 0x7f46d6d0fe10>> ignored
2016-05-23 17:30:40,317 WARNING:toil.leader: The jobWrapper seems to have left a log file, indicating failure: b/8/jobuYbpXG
2016-05-23 17:30:40,318 WARNING:toil.leader: Reporting file: b/8/jobuYbpXG
2016-05-23 17:30:40,318 WARNING:toil.leader: b/8/jobuYbpXG:     ---TOIL WORKER OUTPUT LOG---
2016-05-23 17:30:40,318 WARNING:toil.leader: b/8/jobuYbpXG:     INFO:rdflib:RDFLib Version: 4.2.1
2016-05-23 17:30:40,319 WARNING:toil.leader: b/8/jobuYbpXG:     Traceback (most recent call last):
2016-05-23 17:30:40,319 WARNING:toil.leader: b/8/jobuYbpXG:       File "/home/samuel/.pyenv/versions/2.7.11/lib/python2.7/site-packages/toil/worker.py", line 327, in main
2016-05-23 17:30:40,319 WARNING:toil.leader: b/8/jobuYbpXG:         fileStore=fileStore)
2016-05-23 17:30:40,319 WARNING:toil.leader: b/8/jobuYbpXG:       File "/home/samuel/.pyenv/versions/2.7.11/lib/python2.7/site-packages/toil/job.py", line 1384, in _execute
2016-05-23 17:30:40,319 WARNING:toil.leader: b/8/jobuYbpXG:         returnValues = self._run(jobWrapper, fileStore)
2016-05-23 17:30:40,319 WARNING:toil.leader: b/8/jobuYbpXG:       File "/home/samuel/.pyenv/versions/2.7.11/lib/python2.7/site-packages/toil/job.py", line 1373, in _run
2016-05-23 17:30:40,319 WARNING:toil.leader: b/8/jobuYbpXG:         return self.run(fileStore)
2016-05-23 17:30:40,319 WARNING:toil.leader: b/8/jobuYbpXG:       File "/home/samuel/.pyenv/versions/2.7.11/lib/python2.7/site-packages/toil/cwl/cwltoil.py", line 198, in run
2016-05-23 17:30:40,320 WARNING:toil.leader: b/8/jobuYbpXG:         fillInDefaults(self.cwltool.tool["inputs"], cwljob)
2016-05-23 17:30:40,320 WARNING:toil.leader: b/8/jobuYbpXG:       File "/home/samuel/.pyenv/versions/2.7.11/lib/python2.7/site-packages/cwltool/process.py", line 202, in fillInDefaults
2016-05-23 17:30:40,320 WARNING:toil.leader: b/8/jobuYbpXG:         raise validate.ValidationException("Missing input parameter `%s`" % shortname(inp["id"]))
2016-05-23 17:30:40,320 WARNING:toil.leader: b/8/jobuYbpXG:     ValidationException: Missing input parameter `output_filename`
2016-05-23 17:30:40,320 WARNING:toil.leader: b/8/jobuYbpXG:     Exiting the worker because of a failed jobWrapper on host samuell
2016-05-23 17:30:40,320 WARNING:toil.leader: b/8/jobuYbpXG:     ERROR:toil.worker:Exiting the worker because of a failed jobWrapper on host samuell
2016-05-23 17:30:40,320 WARNING:toil.leader: b/8/jobuYbpXG:     WARNING:toil.jobWrapper:Due to failure we are reducing the remaining retry count of job b/8/jobuYbpXG to 0
2016-05-23 17:30:40,320 WARNING:toil.leader: b/8/jobuYbpXG:     WARNING:toil.jobWrapper:We have increased the default memory of the failed job to 2147483648 bytes
2016-05-23 17:30:40,321 WARNING:toil.leader: Job: b/8/jobuYbpXG is completely failed
```
