# NACsv
A simple utility that turns JSON results (from the blackbaud API) in student directory searches into a CSV for backup.

## How to use it
1. Place the json result inside `in.json`, following the following schema:
```json
{
    "students": [{students...}]
}
```
2. Run the program (you will need to compile it, figure it out)