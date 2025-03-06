# Possible Use Cases

## Advantages

A function like this would be useful for when you really, really want to delete something. Using data recovery tools, it is possible to restore deleted files. Overwriting the bytes of the file first makes it much more difficult to restore the original memory within that file. Another advantage is that the helper function is file system agnostic, so it could be used on different OS platforms and give you a similar behavior.

## Drawbacks

For really large files, this is a lot less efficient than just deleting the file. You are also not completely guaranteeing the erasure of the information, due to different automated backup solutions. 