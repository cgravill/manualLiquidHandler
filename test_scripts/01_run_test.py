#!/usr/bin/env python
import subprocess
import os
import sys

basedir = os.path.dirname(sys.argv[0])
if basedir:
    os.chdir(basedir)
subprocess.check_call(['antharun', '--workflow', '01_workflow.json', '--parameters', '01_parameters.json', '--driver', 'localhost:50051'])
