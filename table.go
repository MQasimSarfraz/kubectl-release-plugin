package kubectlreleaseplugin

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

func FormatAndPrintTable(out io.Writer, headers []string, rows [][]string) error {
	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, strings.Join(headers, "\t"))
	fmt.Fprintln(w)
	for _, values := range rows {
		fmt.Fprintf(w, strings.Join(values, "\t"))
		fmt.Fprintln(w)
	}
	return w.Flush()
}
