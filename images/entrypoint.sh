#!/bin/bash

# Validate conf
validate=true
[[ -z "${DATER_ACTION+x}" ]] \
    && echo "set the DATER_ACTION is required" \
    && validate=false

[[ -z "${FILE_URL+x}" && -z "${FILE_PATH+x}" ]] \
    && echo "FILE_URL or FILE_PATH required" \
    && validate=false

[[ $validate == false ]] && exit 1

# Run dater
exec dater xunit status \
            --fileUrl ${FILE_URL} \
	          --filePath ${FILE_PATH}
    
