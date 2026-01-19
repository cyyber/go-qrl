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

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
repo="$script_dir"
module_path="github.com/theQRL/go-zond"

: "${GOPATH:=$(go env GOPATH)}"
export GOPATH
export PATH="$GOPATH/bin:$PATH"

: "${OUT:=$repo/.oss-fuzz-out}"
mkdir -p "$OUT"
export OUT

: "${SANITIZER:=fuzzer}"
: "${CXX:=clang++}"
: "${CXXFLAGS:=-O1 -fno-omit-frame-pointer -gline-tables-only}"
export SANITIZER CXX CXXFLAGS

EXTRA_LDFLAGS=""
if [[ "$(uname -s)" == "Darwin" ]]; then
	EXTRA_LDFLAGS="-framework CoreFoundation -framework Security"
fi

function ensure_lib_fuzzing_engine() {
	if [[ -n "${LIB_FUZZING_ENGINE:-}" ]]; then
		return 0
	fi

	local sys
	sys="$(uname -s)"
	if [[ "$sys" == "Darwin" ]]; then
		local resource_dir
		resource_dir="$($CXX -print-resource-dir 2>/dev/null || true)"
		if [[ -n "$resource_dir" && -f "$resource_dir/lib/darwin/libclang_rt.fuzzer_osx.a" ]]; then
			local main="$resource_dir/lib/darwin/libclang_rt.fuzzer_osx.a"
			local interceptors="$resource_dir/lib/darwin/libclang_rt.fuzzer_interceptors_osx.a"
			if [[ -f "$interceptors" ]]; then
				LIB_FUZZING_ENGINE="$main $interceptors"
			else
				LIB_FUZZING_ENGINE="$main"
			fi
			export LIB_FUZZING_ENGINE
			return 0
		fi

		local brew_main=""
		local brew_interceptors=""
		for base in /opt/homebrew/Cellar/llvm/*/lib/clang/*/lib/darwin /usr/local/Cellar/llvm/*/lib/clang/*/lib/darwin; do
			if [[ -f "$base/libclang_rt.fuzzer_osx.a" ]]; then
				brew_main="$base/libclang_rt.fuzzer_osx.a"
				if [[ -f "$base/libclang_rt.fuzzer_interceptors_osx.a" ]]; then
					brew_interceptors="$base/libclang_rt.fuzzer_interceptors_osx.a"
				else
					brew_interceptors=""
				fi
			fi
		done
		if [[ -n "$brew_main" ]]; then
			if [[ -n "$brew_interceptors" ]]; then
				LIB_FUZZING_ENGINE="$brew_main $brew_interceptors"
			else
				LIB_FUZZING_ENGINE="$brew_main"
			fi
			export LIB_FUZZING_ENGINE
			return 0
		fi
	fi

	LIB_FUZZING_ENGINE="-fsanitize=fuzzer"
	export LIB_FUZZING_ENGINE
}

ensure_lib_fuzzing_engine

function pkg_dir() {
	local pkg="$1"
	local gopath_dir="$GOPATH/src/$pkg"
	if [[ -d "$gopath_dir" ]]; then
		echo "$gopath_dir"
		return 0
	fi

	if [[ "$pkg" == "$module_path" ]]; then
		echo "$repo"
		return 0
	fi
	if [[ "$pkg" == "$module_path/"* ]]; then
		echo "$repo/${pkg#"$module_path/"}"
		return 0
	fi

	if [[ -d "$repo/$pkg" ]]; then
		echo "$repo/$pkg"
		return 0
	fi

	echo "Unable to resolve package directory for: $pkg" >&2
	return 1
}

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

  path="$(pkg_dir "$package")"

  echo "Building $fuzzer"
  cd $path

  # Install build dependencies
  go mod tidy
  go get github.com/holiman/gofuzz-shim/testing

	if [[ $SANITIZER == *coverage* ]]; then
		coverbuild $path $function $fuzzer
	else
	  gofuzz-shim --func $function --package $package -f $file -o $fuzzer.a
		$CXX $CXXFLAGS $LIB_FUZZING_ENGINE $fuzzer.a $EXTRA_LDFLAGS -o "$OUT/$fuzzer"
		rm -f "$fuzzer.a" "$fuzzer.h" main.*.go
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
repo=$script_dir

compile_fuzzer github.com/theQRL/go-zond/accounts/abi \
  FuzzABI fuzzAbi \
  $repo/accounts/abi/abifuzzer_test.go

compile_fuzzer github.com/theQRL/go-zond/common/bitutil \
  FuzzEncoder fuzzBitutilEncoder \
  $repo/common/bitutil/compress_test.go

compile_fuzzer github.com/theQRL/go-zond/common/bitutil \
  FuzzDecoder fuzzBitutilDecoder \
  $repo/common/bitutil/compress_test.go

compile_fuzzer github.com/theQRL/go-zond/core/vm/runtime \
  FuzzVmRuntime fuzzVmRuntime\
  $repo/core/vm/runtime/runtime_fuzz_test.go

compile_fuzzer github.com/theQRL/go-zond/core/vm \
  FuzzPrecompiledContracts fuzzPrecompiledContracts\
  $repo/core/vm/contracts_fuzz_test.go,$repo/core/vm/contracts_test.go

compile_fuzzer github.com/theQRL/go-zond/core/types \
  FuzzRLP fuzzRlp \
  $repo/core/types/rlp_fuzzer_test.go

compile_fuzzer github.com/theQRL/go-zond/accounts/keystore \
  FuzzPassword fuzzKeystore \
  $repo/accounts/keystore/keystore_fuzzing_test.go

pkg=$repo/trie/
compile_fuzzer github.com/theQRL/go-zond/trie \
  FuzzTrie fuzzTrie \
  $pkg/trie_test.go,$pkg/database_test.go,$pkg/tracer_test.go,$pkg/proof_test.go,$pkg/iterator_test.go,$pkg/sync_test.go

compile_fuzzer github.com/theQRL/go-zond/trie \
  FuzzStackTrie fuzzStackTrie \
  $pkg/stacktrie_fuzzer_test.go,$pkg/iterator_test.go,$pkg/trie_test.go,$pkg/database_test.go,$pkg/tracer_test.go,$pkg/proof_test.go,$pkg/sync_test.go

#compile_fuzzer tests/fuzzers/snap  FuzzARange fuzz_account_range
compile_fuzzer github.com/theQRL/go-zond/qrl/protocols/snap \
  FuzzARange fuzz_account_range \
  $repo/qrl/protocols/snap/handler_fuzzing_test.go

compile_fuzzer github.com/theQRL/go-zond/qrl/protocols/snap \
  FuzzSRange fuzz_storage_range \
  $repo/qrl/protocols/snap/handler_fuzzing_test.go

compile_fuzzer github.com/theQRL/go-zond/qrl/protocols/snap \
  FuzzByteCodes fuzz_byte_codes \
  $repo/qrl/protocols/snap/handler_fuzzing_test.go

compile_fuzzer github.com/theQRL/go-zond/qrl/protocols/snap \
  FuzzTrieNodes fuzz_trie_nodes\
  $repo/qrl/protocols/snap/handler_fuzzing_test.go

compile_fuzzer github.com/theQRL/go-zond/tests/fuzzers/txfetcher \
  Fuzz fuzzTxfetcher \
  $repo/tests/fuzzers/txfetcher/txfetcher_test.go

compile_fuzzer github.com/theQRL/go-zond/tests/fuzzers/secp256k1 \
  Fuzz fuzzSecp256k1\
  $repo/tests/fuzzers/secp256k1/secp_test.go