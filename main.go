package main

import (
	"bytes"
	"github.com/nathanburkett/diff_table/algorithm"
	"github.com/nathanburkett/diff_table/data_transfer"
	"github.com/nathanburkett/diff_table/input"
	"github.com/nathanburkett/diff_table/output"
	"github.com/nathanburkett/diff_table/transform"
	"os"
	"sort"
)

func main() {
	// parse cli flags
	// determine which reader
	// determine which prioritizing algo
	// determine which output type

	test := []byte(`7       1       app/assets/javascripts/admin.js
9       0       app/assets/javascripts/admin/articles.js
5       0       app/assets/stylesheets/admin.css.scss
1       0       app/models/book.rb
1       1       app/views/admin/books/_edit_form.html.erb
0       1       app/views/layouts/admin.html.erb
1       1       config/schedule.production.rb
1       1       lib/tasks/wishlists.rake
1       0       vendor/assets/javascripts/redactor/plugins.js
97      0       vendor/assets/javascripts/redactor/plugins/imagetag.js
`)

	reader := input.NewCliDiffNumstatReader(new(data_transfer.Diff))
	diff, err := reader.Read(bytes.NewBuffer(test))
	if err != nil {
		panic(err)
	}

	reducer := transform.NewDirectoryDiffMapReducer(1, []byte("/"))
	diff = reducer.Reduce(diff)

	sort.Sort(algorithm.TotalDeltaDescendingSorter{
		Diff: diff,
	})

	writer := output.NewMarkdownWriter(os.Stdout)
	if err := writer.Write(diff); err != nil {
		panic(err)
	}
}
