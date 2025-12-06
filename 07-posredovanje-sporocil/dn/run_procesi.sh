#!/bin/bash
#SBATCH --nodes=1
#SBATCH --array=0-4
#SBATCH --reservation=fri
#SBATCH --output=proces-%a.txt

#Zazeni s sbatch run_proces.sh na gruci Arnes
path=.

module load Go
go build $path/procesi.go
srun procesi --id $SLURM_ARRAY_TASK_ID --N $SLURM_ARRAY_TASK_COUNT --root 9000