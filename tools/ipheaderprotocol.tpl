// This is generated by xyun
// Do not modify here.

package const

var ipHeaderProtocol = make(map[int]string)

func init() {
	{{range $key, $value := .}}
	ipHeaderProtocol[{{$key}}] = "{{$value}}"{{end}}
}

func GetIpHeaderProtocol(id int) string{
	return ipHeaderProtocol[id]
}