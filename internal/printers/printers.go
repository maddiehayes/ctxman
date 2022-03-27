package printers

import (
	"io"

	"github.com/liggitt/tabwriter"
)

const (
	tabwriterMinWidth = 6
	tabwriterWidth    = 4
	tabwriterPadding  = 3
	tabwriterPadChar  = ' '
	tabwriterFlags    = tabwriter.RememberWidths
)

func NewTabWriter(output io.Writer) *tabwriter.Writer {
	return tabwriter.NewWriter(output, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)
}

// func main() {
// 	// Observe that the third line has no trailing tab,
// 	// so its final cell is not part of an aligned column.
// 	const padding = 3
// 	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, '-', tabwriter.AlignRight|tabwriter.Debug)
// 	fmt.Fprintln(w, "a\tb\taligned\t")
// 	fmt.Fprintln(w, "aa\tbb\taligned\t")
// 	fmt.Fprintln(w, "aaa\tbbb\tunaligned") // no trailing tab
// 	fmt.Fprintln(w, "aaaa\tbbbb\taligned\t")
// 	w.Flush()

// }
