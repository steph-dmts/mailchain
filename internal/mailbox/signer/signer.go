// Copyright 2019 Finobo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate mockgen -source=signer.go -package=signertest -destination=./signertest/signer_mock.go
package signer

// Signer return a signed transaction
type Signer interface {
	Sign(opts Options) (signedTransaction interface{}, err error)
}

// Options options related to different signers
type Options interface{}
