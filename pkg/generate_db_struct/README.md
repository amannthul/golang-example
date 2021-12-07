1. 根据数据库在 store 层编写好相应的 struct
2. 运行生成代码将会生成 domain 中的所有代码和 db 相关代码的实现

generate_get_func.go 使用方法

1. go build ./generate_get_func.go
2. ./generate_get_func [filePath...] example:(./generate_get_func ../../store/report/report.go ./store/page/page.go)

<!-- 用于生成gorm的column -->
gorm_column_generate.go 使用方法与enerate_get_func.go 使用方法一致
