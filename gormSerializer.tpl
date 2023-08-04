func (x *{{.EnumName}}) Scan(v any) (err error) {
 *x = {{.EnumName}}(v.(int64))
 return nil
}

func (x {{.EnumName}}) Value() (v driver.Value, err error) {
 return int64(x), nil
}
