#!/bin/bash -eu
# Copyright 2022 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
################################################################################

# This sets the -coverpgk for the coverage report when the corpus is executed through go test
coverpkg="github.com/theQRL/go-zond/..."

function coverbuild {
  path=$1
  function=$2
  fuzzer=$3
  tags=""

  if [[ $#  -eq 4 ]]; then
    tags="-tags $4"
  fi
  cd $path
  fuzzed_package=`pwd | rev | cut -d'/' -f 1 | rev`
  cp $GOPATH/ossfuzz_coverage_runner.go ./"${function,,}"_test.go
  sed -i -e 's/FuzzFunction/'$function'/' ./"${function,,}"_test.go
  sed -i -e 's/mypackagebeingfuzzed/'$fuzzed_package'/' ./"${function,,}"_test.go
  sed -i -e 's/TestFuzzCorpus/Test'$function'Corpus/' ./"${function,,}"_test.go

cat << DOG > $OUT/$fuzzer
#/bin/sh

  cd $OUT/$path
  go test -run Test${function}Corpus -v $tags -coverprofile \$1 -coverpkg $coverpkg

DOG

  chmod +x $OUT/$fuzzer
  #echo "Built script $OUT/$fuzzer"
  #cat $OUT/$fuzzer
  cd -
}

function compile_fuzzer() {
  package=$1
  function=$2
  fuzzer=$3
  file=$4

  path=$GOPATH/src/$package

  echo "Building $fuzzer"
  cd $path

  # Install build dependencies
  go mod tidy
  go get github.com/holiman/gofuzz-shim/testing

	if [[ $SANITIZER == *coverage* ]]; then
		coverbuild $path $function $fuzzer $coverpkg
	else
	  gofuzz-shim --func $function --package $package -f $file -o $fuzzer.a
		$CXX $CXXFLAGS $LIB_FUZZING_ENGINE $fuzzer.a -o $OUT/$fuzzer
	fi

  ## Check if there exists a seed corpus file
  corpusfile="${path}/testdata/${fuzzer}_seed_corpus.zip"
  if [ -f $corpusfile ]
  then
    cp $corpusfile $OUT/
    echo "Found seed corpus: $corpusfile"
  fi
  cd -
}

go install github.com/holiman/gofuzz-shim@latest
repo=$GOPATH/src/github.com/theQRL/go-zond
compile_fuzzer tests/fuzzers/bitutil  Fuzz      fuzzBitutilCompress
compile_fuzzer tests/fuzzers/runtime  Fuzz      fuzzVmRuntime
compile_fuzzer tests/fuzzers/keystore   Fuzz fuzzKeystore
compile_fuzzer tests/fuzzers/txfetcher  Fuzz fuzzTxfetcher
compile_fuzzer tests/fuzzers/rlp        Fuzz fuzzRlp
compile_fuzzer tests/fuzzers/trie       Fuzz fuzzTrie
compile_fuzzer tests/fuzzers/stacktrie  Fuzz fuzzStackTrie
compile_fuzzer tests/fuzzers/abi        Fuzz fuzzAbi
compile_fuzzer tests/fuzzers/secp256k1  Fuzz fuzzSecp256k1

compile_fuzzer tests/fuzzers/snap  FuzzARange fuzz_account_range
compile_fuzzer tests/fuzzers/snap  FuzzSRange fuzz_storage_range
compile_fuzzer tests/fuzzers/snap  FuzzByteCodes fuzz_byte_codes
compile_fuzzer tests/fuzzers/snap  FuzzTrieNodes fuzz_trie_nodes
