```sh
go run main.go -f foo.yaml
error converting YAML to JSON: yaml: unmarshal errors:
  line 8: key "name" already set in map
exit status 1
```

change the yaml import
```sh
go run main.go -f foo.yaml
main.YamlToGo{ArrayOfThings:[]main.Thing{main.Thing{Name:"thing-one", AttributeOne:"foo", AttributeTwo:"bar"}, main.Thing{Name:"thing-
two", AttributeOne:"foo", AttributeTwo:"bar"}}}
```
