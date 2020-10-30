# gdep

Tool for exploring dependencies between resources on the Google Cloud Platform

## bigquery

    gdep bigquery PROJECT:DATASET.VIEW1 PROJECT.DATASET.VIEW2 | dot -Tpng  > gdep-bigquery.png && open gdep-bigquery.png

The fully qualified name of the view can have a `:` or a `.` to separte the project from the dataset.
Multiple names are separated by white space.

&copy; 2020, MIT Licensed. http://ernestmicklei.com