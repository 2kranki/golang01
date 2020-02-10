                Go Language Projects

I am currently learning Golang and its associated libraries in
conjunction with Docker.  So, I decided to create this repository
with code that I create.  Everything in this directory is public
domain unless stated otherwise within the file.  

Most of the projects here are operational.  Some are in-progress
and not gauranteed to work correctly.  Those are:
    console04
    console05
    gen


The 'html1' and 'html2' directories give an example of parsing HTML 
with Go using its experimental library. I wrote it to get an example,
because I need to parse html for regression testing in genapp. I have
updated the 'html2' main.go to perform the validation that I wanted.

'main.go' in 'html1' and html2' use the html tokenizer to parse the
html input.  This process is a single pass type system without a reset
option.  It worked fine for a general scan of the html.  'html1' is
an example of creating a dump of an html file.  'html2' is 'html1' with
the addition of some  validation needed for genapp.

'html3' and 'html4' use the html parser which creates a tree of the
parsed html elements.  This process allows for multiple scans of the
tree if needed. 'html3' parses the html and then prints the tree in
'preorder' or depth-first search mode. 'html4' does the same, but with
the addition of some validation needed for genapp.

WARNING: 'html4' is a work in progress and definitely is not complete.



