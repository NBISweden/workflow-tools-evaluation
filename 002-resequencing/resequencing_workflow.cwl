cwlVersion: cwl:draft-3
class: Workflow

inputs:
  - id: ref
    type: File
  - id: reads
    type: File
  - id: bwa_index_algo
    type: string

steps:
  - id: "create_index"
    run: tools/bwa-index.cwl
    inputs:
      - id: "input"
        source: "#ref"
      - id: "a"
        source: "#bwa_index_algo"
    outputs:
      - id: output

  - id: "create_seq_dict"
    run: tools/samtools-faidx.cwl
    inputs:
      - id: "input"
        source: "#ref"
    outputs:
      - id: index

  - id: "align_single_reads"
    run: tools/bwa-aln.cwl
    inputs:
      - id: "prefix"
        source: "#ref"
      - id: "input"
        source: "#reads"
    outputs:
      - id: output

outputs:
  - id: bwa_outputs
    type: File
    source: "#align_single_reads/output"
#    type: { type: array, items: File }
  - id: fasta_index
    type: File
    source: "#create_seq_dict/index"
