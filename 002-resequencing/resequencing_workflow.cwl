cwlVersion: cwl:draft-3
class: Workflow

inputs:
  - id: ref
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

outputs:
  - id: bwa_outputs
    type: { type: array, items: File }
    source: "#create_index/output"
  - id: fasta_index
    type: File
    source: "#create_seq_dict/index"
