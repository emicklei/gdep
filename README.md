# gdep

Tool for exploring dependencies between resources on the Google Cloud Platform

## bigquery

    gdep -o g.dot bigquery PROJECT:DATASET.VIEW1 PROJECT.DATASET.VIEW2 && cat g.dot | dot -Tpng > gdep-bigquery.png && open gdep-bigquery.png

The fully qualified name of the view can have a `:` or a `.` to separate the project from the dataset.
Multiple names are separated by white space.

The `dot` tool is part of Graphviz and needs to be installed separately.

&copy; 2020, MIT Licensed. http://ernestmicklei.com