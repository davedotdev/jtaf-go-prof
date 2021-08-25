// Copyright (c) 2017-2021, Juniper Networks Inc. All rights reserved.
//
// License: Apache 2.0
//
// THIS SOFTWARE IS PROVIDED BY Juniper Networks, Inc. ''AS IS'' AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL Juniper Networks, Inc. BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/rs/xid"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	fileId := xid.New()

	f1, err := os.Create("./" + fileId.String() + "_cpuprof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f1.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f1); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	// -----------------------------------------------------
	// TERRAFORM CODE BEGIN
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return Provider()
		},
	})
	// TERRAFORM CODE END
	// -----------------------------------------------------

	f2, err := os.Create("./" + fileId.String() + "_memprof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f2.Close() // error handling omitted for example
	runtime.GC()     // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f2); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}
