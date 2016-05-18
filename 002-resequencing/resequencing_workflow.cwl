cwlVersion: cwl:draft-3
class: Workflow

inputs:
  - id: ref
    type: File
  - id: bwa_index_algo
    type: string

outputs:
  - id: fasta_index
    type: File
    source: "#create_seq_dict/fasta_index"

steps:
  - id: "create_index"
    run: tools/bwa-index.cwl
    inputs:
      - id: "input"
        source: "#ref"
      - id: "a"
        source: "#bwa_index_algo"
    outputs:
      - id: bwa_outputs

  - id: "create_seq_dict"
    run: tools/samtools-faidx.cwl
    inputs:
      - id: "input"
        source: "#ref"
    outputs:
      - id: fasta_index
