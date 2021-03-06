/*
Copyright © 2021 Bedag Informatik AG

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package printer

import (
	"github.com/olekukonko/tablewriter"
)

type Format int

const (
	FormatJSON Format = iota
	FormatYAML
	FormatTable
	FormatSingle
	InvalidFormat
)

type Printer interface {
	Print()
}

type Printable interface {
	PrinterData(fields []string) map[string]interface{}
}

type Options struct {
	// Some printers can wrap the resulting elements
	// in a list. This option controls if the printer
	// should also wrap a single result element in a list.
	ListWrapSingleItem bool
}

type Queue []Printable

// DataFn defines the function signature needed to
// implement the Printable interface
type DataFn func(fields []string) map[string]interface{}

type job struct {
	dataFn DataFn
}

type structPrinter struct {
	data               []map[string]interface{}
	listWrapSingleItem bool
	marshal            func(interface{}) ([]byte, error)
}

type listPrinter struct {
	data []interface{}
}

type singlePrinter listPrinter

type tablePrinter struct {
	table *tablewriter.Table
}
