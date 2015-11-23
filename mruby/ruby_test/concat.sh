#!/bin/bash
rm super.ruby; for f in `ls *.rb`; do cat $f | grep -v require; echo -e '\n'; done >> super.ruby
