#!/bin/bash

SBATCH --mail-user=dhairyakarna@cs.uchicago.edu
SBATCH --mail-type=ALL
SBATCH --job-name=proj1_benchmark 
SBATCH --output=./slurm/out/%j.%N.stdout
SBATCH --error=./slurm/out/%j.%N.stderr
SBATCH --chdir=/home/dhairyakarna/ParallelProgramming/project-1-DhairyaKarna/proj1/benchmark
SBATCH --partition=debug 
SBATCH --nodes=1
SBATCH --ntasks=1
SBATCH --cpus-per-task=16
SBATCH --mem-per-cpu=900
SBATCH --exclusive
SBATCH --time=5:00


module load golang/1.19
# Your command here
python3 benchmark.py
python3 graph.py
