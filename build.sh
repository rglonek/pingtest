rm -rf bin
mkdir -p bin
for i in linux darwin windows
do
  for j in amd64 386 arm arm64
  do
    fn="bin/pingtest.${i}.${j}"
    if [ "${i}" == "windows" ]
    then
      fn="${fn}.exe"
    fi
    echo "========== Building ${fn} =========="
    GOOS=${i} GOARCH=${j} go build -o ${fn} . && upx ${fn}
  done
done
