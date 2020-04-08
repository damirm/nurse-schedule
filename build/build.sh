GOOS=${GOOS:-darwin}
GO=${GO:-$(which go)}

BUILD_ENTRY=${BUILD_ENTRY:-$1}
OUT=${OUT:-$2}

${GO} build -o ${OUT} ${BUILD_ENTRY}
