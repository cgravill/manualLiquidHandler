#!/usr/bin/env python
import subprocess
import os
import sys

basedir = os.path.dirname(sys.argv[0])
if basedir:
    os.chdir(basedir)
# DDN: Disable for now because planner requires head adaptors not supported by this driver
#subprocess.check_call(['antharun', '--workflow', '01_workflow.json', '--parameters', '01_parameters.json', '--driver', 'localhost:50051'])
